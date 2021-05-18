// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type sqlScriptClient struct {
	con *connection
}

// BeginCreateOrUpdateSQLScript - Creates or updates a Sql Script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) BeginCreateOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptBeginCreateOrUpdateSQLScriptOptions) (SQLScriptResourcePollerResponse, error) {
	resp, err := client.createOrUpdateSQLScript(ctx, sqlScriptName, sqlScript, options)
	if err != nil {
		return SQLScriptResourcePollerResponse{}, err
	}
	result := SQLScriptResourcePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("sqlScriptClient.CreateOrUpdateSQLScript", resp, client.con.Pipeline(), client.createOrUpdateSQLScriptHandleError)
	if err != nil {
		return SQLScriptResourcePollerResponse{}, err
	}
	poller := &sqlScriptResourcePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SQLScriptResourceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdateSQLScript creates a new SQLScriptResourcePoller from the specified resume token.
// token - The value must come from a previous call to SQLScriptResourcePoller.ResumeToken().
func (client *sqlScriptClient) ResumeCreateOrUpdateSQLScript(ctx context.Context, token string) (SQLScriptResourcePollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("sqlScriptClient.CreateOrUpdateSQLScript", token, client.con.Pipeline(), client.createOrUpdateSQLScriptHandleError)
	if err != nil {
		return SQLScriptResourcePollerResponse{}, err
	}
	poller := &sqlScriptResourcePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SQLScriptResourcePollerResponse{}, err
	}
	result := SQLScriptResourcePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SQLScriptResourceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdateSQLScript - Creates or updates a Sql Script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) createOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptBeginCreateOrUpdateSQLScriptOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateSQLScriptCreateRequest(ctx, sqlScriptName, sqlScript, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateSQLScriptHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateSQLScriptCreateRequest creates the CreateOrUpdateSQLScript request.
func (client *sqlScriptClient) createOrUpdateSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptBeginCreateOrUpdateSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sqlScript)
}

// createOrUpdateSQLScriptHandleError handles the CreateOrUpdateSQLScript error response.
func (client *sqlScriptClient) createOrUpdateSQLScriptHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDeleteSQLScript - Deletes a Sql Script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) BeginDeleteSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptBeginDeleteSQLScriptOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteSQLScript(ctx, sqlScriptName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("sqlScriptClient.DeleteSQLScript", resp, client.con.Pipeline(), client.deleteSQLScriptHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDeleteSQLScript creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *sqlScriptClient) ResumeDeleteSQLScript(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("sqlScriptClient.DeleteSQLScript", token, client.con.Pipeline(), client.deleteSQLScriptHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// DeleteSQLScript - Deletes a Sql Script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) deleteSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptBeginDeleteSQLScriptOptions) (*azcore.Response, error) {
	req, err := client.deleteSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteSQLScriptHandleError(resp)
	}
	return resp, nil
}

// deleteSQLScriptCreateRequest creates the DeleteSQLScript request.
func (client *sqlScriptClient) deleteSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SQLScriptBeginDeleteSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteSQLScriptHandleError handles the DeleteSQLScript error response.
func (client *sqlScriptClient) deleteSQLScriptHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetSQLScript - Gets a sql script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) GetSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptGetSQLScriptOptions) (SQLScriptResourceResponse, error) {
	req, err := client.getSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return SQLScriptResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SQLScriptResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return SQLScriptResourceResponse{}, client.getSQLScriptHandleError(resp)
	}
	return client.getSQLScriptHandleResponse(resp)
}

// getSQLScriptCreateRequest creates the GetSQLScript request.
func (client *sqlScriptClient) getSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SQLScriptGetSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSQLScriptHandleResponse handles the GetSQLScript response.
func (client *sqlScriptClient) getSQLScriptHandleResponse(resp *azcore.Response) (SQLScriptResourceResponse, error) {
	var val *SQLScriptResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SQLScriptResourceResponse{}, err
	}
	return SQLScriptResourceResponse{RawResponse: resp.Response, SQLScriptResource: val}, nil
}

// getSQLScriptHandleError handles the GetSQLScript error response.
func (client *sqlScriptClient) getSQLScriptHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetSQLScriptsByWorkspace - Lists sql scripts.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) GetSQLScriptsByWorkspace(options *SQLScriptGetSQLScriptsByWorkspaceOptions) SQLScriptsListResponsePager {
	return &sqlScriptsListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getSQLScriptsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.getSQLScriptsByWorkspaceHandleResponse,
		errorer:   client.getSQLScriptsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp SQLScriptsListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SQLScriptsListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getSQLScriptsByWorkspaceCreateRequest creates the GetSQLScriptsByWorkspace request.
func (client *sqlScriptClient) getSQLScriptsByWorkspaceCreateRequest(ctx context.Context, options *SQLScriptGetSQLScriptsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSQLScriptsByWorkspaceHandleResponse handles the GetSQLScriptsByWorkspace response.
func (client *sqlScriptClient) getSQLScriptsByWorkspaceHandleResponse(resp *azcore.Response) (SQLScriptsListResponseResponse, error) {
	var val *SQLScriptsListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SQLScriptsListResponseResponse{}, err
	}
	return SQLScriptsListResponseResponse{RawResponse: resp.Response, SQLScriptsListResponse: val}, nil
}

// getSQLScriptsByWorkspaceHandleError handles the GetSQLScriptsByWorkspace error response.
func (client *sqlScriptClient) getSQLScriptsByWorkspaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginRenameSQLScript - Renames a sqlScript.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) BeginRenameSQLScript(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptBeginRenameSQLScriptOptions) (HTTPPollerResponse, error) {
	resp, err := client.renameSQLScript(ctx, sqlScriptName, request, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("sqlScriptClient.RenameSQLScript", resp, client.con.Pipeline(), client.renameSQLScriptHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRenameSQLScript creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *sqlScriptClient) ResumeRenameSQLScript(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("sqlScriptClient.RenameSQLScript", token, client.con.Pipeline(), client.renameSQLScriptHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// RenameSQLScript - Renames a sqlScript.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) renameSQLScript(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptBeginRenameSQLScriptOptions) (*azcore.Response, error) {
	req, err := client.renameSQLScriptCreateRequest(ctx, sqlScriptName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameSQLScriptHandleError(resp)
	}
	return resp, nil
}

// renameSQLScriptCreateRequest creates the RenameSQLScript request.
func (client *sqlScriptClient) renameSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptBeginRenameSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}/rename"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renameSQLScriptHandleError handles the RenameSQLScript error response.
func (client *sqlScriptClient) renameSQLScriptHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
