package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	ExitOk int = iota
)

type State int

const (
	StateClaim State = iota
	StateProvision
	StatePopulate
	StateBuild
	StateProgress
	StateExport
	StateDone

	StateError
	StateSignal
	StateTimeout
)

var (
	argJSON bool

	argBuilderHost string
	argBuilderPort int

	argTimeoutClaim     int
	argTimeoutProvision int
	argTimeoutPopulate  int
	argTimeoutBuild     int
	argTimeoutExport    int

	argBuildPath string
)

type BuildRequest struct {
	Pipelines    []string `json:"pipelines"`
	Environments []string `json:"environments"`
}

func init() {
	flag.BoolVar(&argJSON, "json", false, "Enable JSON output")

	flag.StringVar(&argBuilderHost, "builder-host", "localhost", "Hostname or IP where this program will listen on.")
	flag.IntVar(&argBuilderPort, "builder-port", 3333, "Port this program will listen on.")

	flag.IntVar(&argTimeoutClaim, "timeout-claim", 600, "Timeout before the claim phase needs to be completed in seconds.")
	flag.IntVar(&argTimeoutProvision, "timeout-provision", 30, "Timeout before the provision phase needs to be completed in seconds.")
	flag.IntVar(&argTimeoutPopulate, "timeout-populate", 30, "Timeout before the populate phase needs to be completed in seconds.")
	flag.IntVar(&argTimeoutBuild, "timeout-build", 3600, "Timeout before the build phase needs to be completed in seconds.")
	flag.IntVar(&argTimeoutExport, "timeout-export", 1800, "Timeout before the export phase needs to be completed in seconds.")

	flag.StringVar(&argBuildPath, "build-path", "/run/osbuild", "Path to use as a build directory.")

	flag.Parse()

	logrus.SetLevel(logrus.InfoLevel)

	if argJSON {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

type Agent struct {
	Host         string
	Port         int
	State        State
	StateLock    sync.Mutex
	StateChannel chan State
}

type BackgroundProcess struct {
	Process *exec.Cmd
	Stdout  io.ReadCloser
	Stderr  io.ReadCloser
	Done    bool
	Error   error
}

var (
	Build *BackgroundProcess
)

func (agent *Agent) SetState(state State) {
	agent.StateLock.Lock()
	defer agent.StateLock.Unlock()

	if state <= agent.State {
		agent.State = StateError
	} else {
		agent.State = state
	}

	agent.StateChannel <- agent.State
}

func (agent *Agent) GetState() State {
	agent.StateLock.Lock()
	defer agent.StateLock.Unlock()

	return agent.State
}

func (agent *Agent) GuardState(stateWanted State) {
	if stateCurrent := agent.GetState(); stateWanted != stateCurrent {
		logrus.Fatalf("Agent.GuardState: Requested guard for %d but we're in %d. Exit.", stateWanted, stateCurrent)
	}
}

func (agent *Agent) HandleClaim(w http.ResponseWriter, r *http.Request) {
	agent.GuardState(StateClaim)

	if r.Method != "POST" {
		logrus.WithFields(
			logrus.Fields{"method": r.Method},
		).Fatal("Agent.HandleClaim: unexpected request method")
	}

	fmt.Fprintf(w, "%s", "done")

	logrus.Info("Agent.HandleClaim: Done.")

	agent.SetState(StateProvision)
}

func (agent *Agent) HandleProvision(w http.ResponseWriter, r *http.Request) {
	agent.GuardState(StateProvision)

	if r.Method != "PUT" {
		logrus.WithFields(
			logrus.Fields{"method": r.Method},
		).Fatal("Agent.HandleProvision: Unexpected request method.")
	}

	logrus.WithFields(logrus.Fields{"argBuildPath": argBuildPath}).Debug("Agent.HandleProvision: Opening manifest.json.")

	dst, err := os.OpenFile(
		path.Join(argBuildPath, "manifest.json"),
		os.O_WRONLY|os.O_CREATE|os.O_EXCL,
		0400,
	)
	defer dst.Close()

	if err != nil {
		logrus.Fatal("Agent.HandleProvision: Failed to open manifest.json.")
	}

	logrus.Debug("Agent.HandleProvision: Writing manifest.json.")

	_, err = io.Copy(dst, r.Body)

	if err != nil {
		logrus.Fatal("Agent.HandleProvision: Failed to write manifest.json.")
	}

	w.WriteHeader(http.StatusCreated)

	if _, err := w.Write([]byte(`done`)); err != nil {
		logrus.Fatal("Agent.HandleProvision: Failed to write response.")
	}

	logrus.Info("Agent.HandleProvision: Done.")

	agent.SetState(StatePopulate)
}

func (agent *Agent) HandlePopulate(w http.ResponseWriter, r *http.Request) {
	agent.GuardState(StatePopulate)

	if r.Method != "POST" {
		logrus.WithFields(
			logrus.Fields{"method": r.Method},
		).Fatal("Agent.HandlePopulate: unexpected request method")
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(`done`)); err != nil {
		logrus.Fatal("Agent.HandlePopulate: Failed to write response.")
	}

	logrus.Info("Agent.HandlePopulate: Done.")

	agent.SetState(StateBuild)
}

func (agent *Agent) HandleBuild(w http.ResponseWriter, r *http.Request) {
	agent.GuardState(StateBuild)

	if r.Method != "POST" {
		logrus.WithFields(
			logrus.Fields{"method": r.Method},
		).Fatal("Agent.HandleBuild: Unexpected request method.")
	}

	var buildRequest BuildRequest

	var err error

	if err = json.NewDecoder(r.Body).Decode(&buildRequest); err != nil {
		logrus.Fatal("HandleBuild: Failed to decode body.")
	}

	if Build != nil {
		logrus.Fatal("HandleBuild: Build started but Build was non-nil.")
	}

	args := []string{
		"--store", path.Join(argBuildPath, "store"),
		"--cache-max-size", "unlimited",
		"--checkpoint", "*",
		"--output-directory", path.Join(argBuildPath, "export"),
	}

	for _, pipeline := range buildRequest.Pipelines {
		args = append(args, "--export")
		args = append(args, pipeline)
	}

	args = append(args, path.Join(argBuildPath, "manifest.json"))

	envs := os.Environ()
	envs = append(envs, buildRequest.Environments...)

	Build = &BackgroundProcess{}
	Build.Process = exec.Command(
		"/usr/bin/osbuild",
		args...,
	)
	Build.Process.Env = envs

	logrus.Infof("BackgroundProcess: Starting %s with %s", Build.Process, envs)

	Build.Stdout, err = Build.Process.StdoutPipe()

	if err != nil {
		logrus.Fatal(err)
	}

	Build.Stderr, err = Build.Process.StderrPipe()

	if err != nil {
		logrus.Fatal(err)
	}

	if err := Build.Process.Start(); err != nil {
		logrus.Fatal("BackgroundProcess: Failed to start process.")
	}

	go func() {
		Build.Error = Build.Process.Wait()
		Build.Done = true

		logrus.Info("BackgroundProcess: Exited.")
	}()

	go func() {
		scanner := bufio.NewScanner(Build.Stdout)
		for scanner.Scan() {
			m := scanner.Text()
			logrus.Infof("BackgroundProcess: Stdout: %s", m)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(Build.Stderr)
		for scanner.Scan() {
			m := scanner.Text()
			logrus.Infof("BackgroundProcess: Stderr: %s", m)
		}
	}()

	w.WriteHeader(http.StatusCreated)

	agent.SetState(StateProgress)
}

func (agent *Agent) HandleProgress(w http.ResponseWriter, r *http.Request) {
	agent.GuardState(StateProgress)

	if r.Method != "GET" {
		logrus.WithFields(
			logrus.Fields{"method": r.Method},
		).Fatal("Agent.HandleProgress: Unexpected request method.")
	}

	if Build == nil {
		logrus.Fatal("HandleProgress: Progress requested but Build was nil.")
	}

	if Build.Done {
		w.WriteHeader(http.StatusOK)

		if Build.Error != nil {
			logrus.Fatalf("Agent.HandleBuild: Buildprocess exited with error: %s", Build.Error)
		}

		agent.SetState(StateExport)
	} else {
		w.WriteHeader(http.StatusAccepted)
	}

	if _, err := w.Write([]byte(`done`)); err != nil {
		logrus.Fatal("Agent.HandleBuild: Failed to write response.")
	}

	logrus.Info("Agent.HandleBuild: Done.")
}

func (agent *Agent) HandleExport(w http.ResponseWriter, r *http.Request) {
	agent.GuardState(StateExport)

	if r.Method != "GET" {
		logrus.WithFields(
			logrus.Fields{"method": r.Method},
		).Fatal("Agent.HandleExport: unexpected request method")
	}

	exportPath := r.URL.Query().Get("path")

	if exportPath == "" {
		logrus.Fatal("Agent.HandleExport: Missing export.")
	}

	// XXX check subdir
	srcPath := path.Join(argBuildPath, "export", exportPath)

	src, err := os.Open(
		srcPath,
	)

	if err != nil {
		logrus.Fatalf("Agent.HandleExport: Failed to open source: %s.", err)
	}

	_, err = io.Copy(w, src)

	if err != nil {
		logrus.Fatalf("Agent.HandleExport: Failed to write response: %s.", err)
	}

	logrus.Info("Agent.HandleExport: Done.")

	agent.SetState(StateDone)
}

func (agent *Agent) Serve() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/claim", agent.HandleClaim)

	mux.HandleFunc("/provision", agent.HandleProvision)
	mux.HandleFunc("/populate", agent.HandlePopulate)

	mux.HandleFunc("/build", agent.HandleBuild)
	mux.HandleFunc("/progress", agent.HandleProgress)

	mux.HandleFunc("/export", agent.HandleExport)

	net := &http.Server{
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1800 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		Addr:              fmt.Sprintf("%s:%d", agent.Host, agent.Port),
		Handler:           mux,
	}

	return net.ListenAndServe()
}

func main() {
	logrus.WithFields(
		logrus.Fields{
			"argJSON":             argJSON,
			"argBuilderHost":      argBuilderHost,
			"argBuilderPort":      argBuilderPort,
			"argTimeoutClaim":     argTimeoutClaim,
			"argTimeoutProvision": argTimeoutProvision,
			"argTimeoutBuild":     argTimeoutBuild,
			"argTimeoutExport":    argTimeoutExport,
		}).Info("main: startup")

	agent := Agent{
		State:        StateClaim,
		StateChannel: make(chan State, 16),
		Host:         argBuilderHost,
		Port:         argBuilderPort,
	}
	go agent.Serve()

	for state := range agent.StateChannel {
		if state == StateDone {
			logrus.Info("main: Shutdown.")
			os.Exit(ExitOk)
		}
	}
}