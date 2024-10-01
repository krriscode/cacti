/*
Hyperledger Cactus Plugin - Keychain Memory 

Contains/describes the Hyperledger Cacti Keychain Memory plugin.

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-keychain-memory

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiDeleteKeychainEntryV1Request struct {
	ctx context.Context
	ApiService *DefaultApiService
	deleteKeychainEntryRequestV1 *DeleteKeychainEntryRequestV1
}

// Request body to delete a keychain entry via its key
func (r ApiDeleteKeychainEntryV1Request) DeleteKeychainEntryRequestV1(deleteKeychainEntryRequestV1 DeleteKeychainEntryRequestV1) ApiDeleteKeychainEntryV1Request {
	r.deleteKeychainEntryRequestV1 = &deleteKeychainEntryRequestV1
	return r
}

func (r ApiDeleteKeychainEntryV1Request) Execute() (*DeleteKeychainEntryResponseV1, *http.Response, error) {
	return r.ApiService.DeleteKeychainEntryV1Execute(r)
}

/*
DeleteKeychainEntryV1 Deletes an entry under a key on the keychain backend.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteKeychainEntryV1Request
*/
func (a *DefaultApiService) DeleteKeychainEntryV1(ctx context.Context) ApiDeleteKeychainEntryV1Request {
	return ApiDeleteKeychainEntryV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DeleteKeychainEntryResponseV1
func (a *DefaultApiService) DeleteKeychainEntryV1Execute(r ApiDeleteKeychainEntryV1Request) (*DeleteKeychainEntryResponseV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteKeychainEntryResponseV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteKeychainEntryV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/delete-keychain-entry"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deleteKeychainEntryRequestV1 == nil {
		return localVarReturnValue, nil, reportError("deleteKeychainEntryRequestV1 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.deleteKeychainEntryRequestV1
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

type ApiGetKeychainEntryV1Request struct {
	ctx context.Context
	ApiService *DefaultApiService
	getKeychainEntryRequestV1 *GetKeychainEntryRequestV1
}

// Request body to obtain a keychain entry via its key
func (r ApiGetKeychainEntryV1Request) GetKeychainEntryRequestV1(getKeychainEntryRequestV1 GetKeychainEntryRequestV1) ApiGetKeychainEntryV1Request {
	r.getKeychainEntryRequestV1 = &getKeychainEntryRequestV1
	return r
}

func (r ApiGetKeychainEntryV1Request) Execute() (*GetKeychainEntryResponseV1, *http.Response, error) {
	return r.ApiService.GetKeychainEntryV1Execute(r)
}

/*
GetKeychainEntryV1 Retrieves the contents of a keychain entry from the backend.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetKeychainEntryV1Request
*/
func (a *DefaultApiService) GetKeychainEntryV1(ctx context.Context) ApiGetKeychainEntryV1Request {
	return ApiGetKeychainEntryV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetKeychainEntryResponseV1
func (a *DefaultApiService) GetKeychainEntryV1Execute(r ApiGetKeychainEntryV1Request) (*GetKeychainEntryResponseV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetKeychainEntryResponseV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetKeychainEntryV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/get-keychain-entry"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getKeychainEntryRequestV1 == nil {
		return localVarReturnValue, nil, reportError("getKeychainEntryRequestV1 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.getKeychainEntryRequestV1
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

type ApiGetPrometheusMetricsV1Request struct {
	ctx context.Context
	ApiService *DefaultApiService
}

func (r ApiGetPrometheusMetricsV1Request) Execute() (string, *http.Response, error) {
	return r.ApiService.GetPrometheusMetricsV1Execute(r)
}

/*
GetPrometheusMetricsV1 Get the Prometheus Metrics

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPrometheusMetricsV1Request
*/
func (a *DefaultApiService) GetPrometheusMetricsV1(ctx context.Context) ApiGetPrometheusMetricsV1Request {
	return ApiGetPrometheusMetricsV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *DefaultApiService) GetPrometheusMetricsV1Execute(r ApiGetPrometheusMetricsV1Request) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetPrometheusMetricsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/get-prometheus-exporter-metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

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

type ApiHasKeychainEntryV1Request struct {
	ctx context.Context
	ApiService *DefaultApiService
	hasKeychainEntryRequestV1 *HasKeychainEntryRequestV1
}

// Request body for checking a keychain entry via its key
func (r ApiHasKeychainEntryV1Request) HasKeychainEntryRequestV1(hasKeychainEntryRequestV1 HasKeychainEntryRequestV1) ApiHasKeychainEntryV1Request {
	r.hasKeychainEntryRequestV1 = &hasKeychainEntryRequestV1
	return r
}

func (r ApiHasKeychainEntryV1Request) Execute() (*HasKeychainEntryResponseV1, *http.Response, error) {
	return r.ApiService.HasKeychainEntryV1Execute(r)
}

/*
HasKeychainEntryV1 Checks that an entry exists under a key on the keychain backend

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHasKeychainEntryV1Request
*/
func (a *DefaultApiService) HasKeychainEntryV1(ctx context.Context) ApiHasKeychainEntryV1Request {
	return ApiHasKeychainEntryV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HasKeychainEntryResponseV1
func (a *DefaultApiService) HasKeychainEntryV1Execute(r ApiHasKeychainEntryV1Request) (*HasKeychainEntryResponseV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HasKeychainEntryResponseV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.HasKeychainEntryV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/has-keychain-entry"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.hasKeychainEntryRequestV1 == nil {
		return localVarReturnValue, nil, reportError("hasKeychainEntryRequestV1 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.hasKeychainEntryRequestV1
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

type ApiSetKeychainEntryV1Request struct {
	ctx context.Context
	ApiService *DefaultApiService
	setKeychainEntryRequestV1 *SetKeychainEntryRequestV1
}

// Request body to write/update a keychain entry via its key
func (r ApiSetKeychainEntryV1Request) SetKeychainEntryRequestV1(setKeychainEntryRequestV1 SetKeychainEntryRequestV1) ApiSetKeychainEntryV1Request {
	r.setKeychainEntryRequestV1 = &setKeychainEntryRequestV1
	return r
}

func (r ApiSetKeychainEntryV1Request) Execute() (*SetKeychainEntryResponseV1, *http.Response, error) {
	return r.ApiService.SetKeychainEntryV1Execute(r)
}

/*
SetKeychainEntryV1 Sets a value under a key on the keychain backend.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSetKeychainEntryV1Request
*/
func (a *DefaultApiService) SetKeychainEntryV1(ctx context.Context) ApiSetKeychainEntryV1Request {
	return ApiSetKeychainEntryV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SetKeychainEntryResponseV1
func (a *DefaultApiService) SetKeychainEntryV1Execute(r ApiSetKeychainEntryV1Request) (*SetKeychainEntryResponseV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SetKeychainEntryResponseV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.SetKeychainEntryV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/set-keychain-entry"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setKeychainEntryRequestV1 == nil {
		return localVarReturnValue, nil, reportError("setKeychainEntryRequestV1 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.setKeychainEntryRequestV1
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
