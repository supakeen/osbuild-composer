/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
	"reflect"
)


// ContentCollectionSignaturesAPIService ContentCollectionSignaturesAPI service
type ContentCollectionSignaturesAPIService service

type ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest struct {
	ctx context.Context
	ApiService *ContentCollectionSignaturesAPIService
	file *os.File
	signedCollection *string
	repository *string
}

// An uploaded file that may be turned into the artifact of the content unit.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest) File(file *os.File) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest {
	r.file = file
	return r
}

// The content this signature is pointing to.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest) SignedCollection(signedCollection string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest {
	r.signedCollection = &signedCollection
	return r
}

// A URI of a repository the new content unit should be associated with.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest) Repository(repository string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest {
	r.repository = &repository
	return r
}

func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionSignaturesCreateExecute(r)
}

/*
ContentAnsibleCollectionSignaturesCreate Create a collection version signature

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest
*/
func (a *ContentCollectionSignaturesAPIService) ContentAnsibleCollectionSignaturesCreate(ctx context.Context) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest {
	return ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentCollectionSignaturesAPIService) ContentAnsibleCollectionSignaturesCreateExecute(r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionSignaturesAPIService.ContentAnsibleCollectionSignaturesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_signatures/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}
	if r.signedCollection == nil {
		return localVarReturnValue, nil, reportError("signedCollection is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "")
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_collection", r.signedCollection, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest struct {
	ctx context.Context
	ApiService *ContentCollectionSignaturesAPIService
	limit *int32
	offset *int32
	ordering *[]string
	pubkeyFingerprint *string
	pubkeyFingerprintIn *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	signedCollection *string
	signingService *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) Limit(limit int32) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) Offset(offset int32) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;data&#x60; - Data * &#x60;-data&#x60; - Data (descending) * &#x60;digest&#x60; - Digest * &#x60;-digest&#x60; - Digest (descending) * &#x60;pubkey_fingerprint&#x60; - Pubkey fingerprint * &#x60;-pubkey_fingerprint&#x60; - Pubkey fingerprint (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) Ordering(ordering []string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pubkey_fingerprint matches value
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) PubkeyFingerprint(pubkeyFingerprint string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.pubkeyFingerprint = &pubkeyFingerprint
	return r
}

// Filter results where pubkey_fingerprint is in a comma-separated list of values
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) PubkeyFingerprintIn(pubkeyFingerprintIn []string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.pubkeyFingerprintIn = &pubkeyFingerprintIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) PulpHrefIn(pulpHrefIn []string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) PulpIdIn(pulpIdIn []string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) RepositoryVersion(repositoryVersion string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter signatures for collection version
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) SignedCollection(signedCollection string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.signedCollection = &signedCollection
	return r
}

// Filter signatures produced by signature service
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) SigningService(signingService string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.signingService = &signingService
	return r
}

// A list of fields to include in the response.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) Fields(fields []string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) ExcludeFields(excludeFields []string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) Execute() (*PaginatedansibleCollectionVersionSignatureResponseList, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionSignaturesListExecute(r)
}

/*
ContentAnsibleCollectionSignaturesList List collection version signatures

ViewSet for looking at signature objects for CollectionVersion content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest
*/
func (a *ContentCollectionSignaturesAPIService) ContentAnsibleCollectionSignaturesList(ctx context.Context) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest {
	return ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedansibleCollectionVersionSignatureResponseList
func (a *ContentCollectionSignaturesAPIService) ContentAnsibleCollectionSignaturesListExecute(r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesListRequest) (*PaginatedansibleCollectionVersionSignatureResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedansibleCollectionVersionSignatureResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionSignaturesAPIService.ContentAnsibleCollectionSignaturesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_signatures/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pubkeyFingerprint != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pubkey_fingerprint", r.pubkeyFingerprint, "")
	}
	if r.pubkeyFingerprintIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pubkey_fingerprint__in", r.pubkeyFingerprintIn, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.signedCollection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "signed_collection", r.signedCollection, "")
	}
	if r.signingService != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "signing_service", r.signingService, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest struct {
	ctx context.Context
	ApiService *ContentCollectionSignaturesAPIService
	ansibleCollectionVersionSignatureHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest) Fields(fields []string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest) ExcludeFields(excludeFields []string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest) Execute() (*AnsibleCollectionVersionSignatureResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionSignaturesReadExecute(r)
}

/*
ContentAnsibleCollectionSignaturesRead Inspect a collection version signature

ViewSet for looking at signature objects for CollectionVersion content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionVersionSignatureHref
 @return ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest
*/
func (a *ContentCollectionSignaturesAPIService) ContentAnsibleCollectionSignaturesRead(ctx context.Context, ansibleCollectionVersionSignatureHref string) ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest {
	return ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionVersionSignatureHref: ansibleCollectionVersionSignatureHref,
	}
}

// Execute executes the request
//  @return AnsibleCollectionVersionSignatureResponse
func (a *ContentCollectionSignaturesAPIService) ContentAnsibleCollectionSignaturesReadExecute(r ContentCollectionSignaturesAPIContentAnsibleCollectionSignaturesReadRequest) (*AnsibleCollectionVersionSignatureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnsibleCollectionVersionSignatureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionSignaturesAPIService.ContentAnsibleCollectionSignaturesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{ansible_collection_version_signature_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_version_signature_href"+"}", parameterValueToString(r.ansibleCollectionVersionSignatureHref, "ansibleCollectionVersionSignatureHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}