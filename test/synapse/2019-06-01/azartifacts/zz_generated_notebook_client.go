//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type notebookClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newNotebookClient creates a new instance of notebookClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newNotebookClient(endpoint string, pl runtime.Pipeline) *notebookClient {
	client := &notebookClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// BeginCreateOrUpdateNotebook - Creates or updates a Note Book.
// If the operation fails it returns the *CloudError error type.
// notebookName - The notebook name.
// notebook - Note book resource definition.
// options - NotebookBeginCreateOrUpdateNotebookOptions contains the optional parameters for the Notebook.BeginCreateOrUpdateNotebook
// method.
func (client *notebookClient) BeginCreateOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (NotebookCreateOrUpdateNotebookPollerResponse, error) {
	resp, err := client.createOrUpdateNotebook(ctx, notebookName, notebook, options)
	if err != nil {
		return NotebookCreateOrUpdateNotebookPollerResponse{}, err
	}
	result := NotebookCreateOrUpdateNotebookPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("notebookClient.CreateOrUpdateNotebook", resp, client.pl, client.createOrUpdateNotebookHandleError)
	if err != nil {
		return NotebookCreateOrUpdateNotebookPollerResponse{}, err
	}
	result.Poller = &NotebookCreateOrUpdateNotebookPoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdateNotebook - Creates or updates a Note Book.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) createOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (*http.Response, error) {
	req, err := client.createOrUpdateNotebookCreateRequest(ctx, notebookName, notebook, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateNotebookHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateNotebookCreateRequest creates the CreateOrUpdateNotebook request.
func (client *notebookClient) createOrUpdateNotebookCreateRequest(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (*policy.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, notebook)
}

// createOrUpdateNotebookHandleError handles the CreateOrUpdateNotebook error response.
func (client *notebookClient) createOrUpdateNotebookHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDeleteNotebook - Deletes a Note book.
// If the operation fails it returns the *CloudError error type.
// notebookName - The notebook name.
// options - NotebookBeginDeleteNotebookOptions contains the optional parameters for the Notebook.BeginDeleteNotebook method.
func (client *notebookClient) BeginDeleteNotebook(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (NotebookDeleteNotebookPollerResponse, error) {
	resp, err := client.deleteNotebook(ctx, notebookName, options)
	if err != nil {
		return NotebookDeleteNotebookPollerResponse{}, err
	}
	result := NotebookDeleteNotebookPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("notebookClient.DeleteNotebook", resp, client.pl, client.deleteNotebookHandleError)
	if err != nil {
		return NotebookDeleteNotebookPollerResponse{}, err
	}
	result.Poller = &NotebookDeleteNotebookPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteNotebook - Deletes a Note book.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) deleteNotebook(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (*http.Response, error) {
	req, err := client.deleteNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteNotebookHandleError(resp)
	}
	return resp, nil
}

// deleteNotebookCreateRequest creates the DeleteNotebook request.
func (client *notebookClient) deleteNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (*policy.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteNotebookHandleError handles the DeleteNotebook error response.
func (client *notebookClient) deleteNotebookHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetNotebook - Gets a Note Book.
// If the operation fails it returns the *CloudError error type.
// notebookName - The notebook name.
// options - NotebookGetNotebookOptions contains the optional parameters for the Notebook.GetNotebook method.
func (client *notebookClient) GetNotebook(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (NotebookGetNotebookResponse, error) {
	req, err := client.getNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return NotebookGetNotebookResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotebookGetNotebookResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return NotebookGetNotebookResponse{}, client.getNotebookHandleError(resp)
	}
	return client.getNotebookHandleResponse(resp)
}

// getNotebookCreateRequest creates the GetNotebook request.
func (client *notebookClient) getNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (*policy.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNotebookHandleResponse handles the GetNotebook response.
func (client *notebookClient) getNotebookHandleResponse(resp *http.Response) (NotebookGetNotebookResponse, error) {
	result := NotebookGetNotebookResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookResource); err != nil {
		return NotebookGetNotebookResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getNotebookHandleError handles the GetNotebook error response.
func (client *notebookClient) getNotebookHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetNotebookSummaryByWorkSpace - Lists a summary of Notebooks.
// If the operation fails it returns the *CloudError error type.
// options - NotebookGetNotebookSummaryByWorkSpaceOptions contains the optional parameters for the Notebook.GetNotebookSummaryByWorkSpace
// method.
func (client *notebookClient) GetNotebookSummaryByWorkSpace(options *NotebookGetNotebookSummaryByWorkSpaceOptions) *NotebookGetNotebookSummaryByWorkSpacePager {
	return &NotebookGetNotebookSummaryByWorkSpacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getNotebookSummaryByWorkSpaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp NotebookGetNotebookSummaryByWorkSpaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
	}
}

// getNotebookSummaryByWorkSpaceCreateRequest creates the GetNotebookSummaryByWorkSpace request.
func (client *notebookClient) getNotebookSummaryByWorkSpaceCreateRequest(ctx context.Context, options *NotebookGetNotebookSummaryByWorkSpaceOptions) (*policy.Request, error) {
	urlPath := "/notebooks/summary"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNotebookSummaryByWorkSpaceHandleResponse handles the GetNotebookSummaryByWorkSpace response.
func (client *notebookClient) getNotebookSummaryByWorkSpaceHandleResponse(resp *http.Response) (NotebookGetNotebookSummaryByWorkSpaceResponse, error) {
	result := NotebookGetNotebookSummaryByWorkSpaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookListResponse); err != nil {
		return NotebookGetNotebookSummaryByWorkSpaceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getNotebookSummaryByWorkSpaceHandleError handles the GetNotebookSummaryByWorkSpace error response.
func (client *notebookClient) getNotebookSummaryByWorkSpaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetNotebooksByWorkspace - Lists Notebooks.
// If the operation fails it returns the *CloudError error type.
// options - NotebookGetNotebooksByWorkspaceOptions contains the optional parameters for the Notebook.GetNotebooksByWorkspace
// method.
func (client *notebookClient) GetNotebooksByWorkspace(options *NotebookGetNotebooksByWorkspaceOptions) *NotebookGetNotebooksByWorkspacePager {
	return &NotebookGetNotebooksByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getNotebooksByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp NotebookGetNotebooksByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
	}
}

// getNotebooksByWorkspaceCreateRequest creates the GetNotebooksByWorkspace request.
func (client *notebookClient) getNotebooksByWorkspaceCreateRequest(ctx context.Context, options *NotebookGetNotebooksByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/notebooks"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNotebooksByWorkspaceHandleResponse handles the GetNotebooksByWorkspace response.
func (client *notebookClient) getNotebooksByWorkspaceHandleResponse(resp *http.Response) (NotebookGetNotebooksByWorkspaceResponse, error) {
	result := NotebookGetNotebooksByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookListResponse); err != nil {
		return NotebookGetNotebooksByWorkspaceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getNotebooksByWorkspaceHandleError handles the GetNotebooksByWorkspace error response.
func (client *notebookClient) getNotebooksByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRenameNotebook - Renames a notebook.
// If the operation fails it returns the *CloudError error type.
// notebookName - The notebook name.
// request - proposed new name.
// options - NotebookBeginRenameNotebookOptions contains the optional parameters for the Notebook.BeginRenameNotebook method.
func (client *notebookClient) BeginRenameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (NotebookRenameNotebookPollerResponse, error) {
	resp, err := client.renameNotebook(ctx, notebookName, request, options)
	if err != nil {
		return NotebookRenameNotebookPollerResponse{}, err
	}
	result := NotebookRenameNotebookPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("notebookClient.RenameNotebook", resp, client.pl, client.renameNotebookHandleError)
	if err != nil {
		return NotebookRenameNotebookPollerResponse{}, err
	}
	result.Poller = &NotebookRenameNotebookPoller{
		pt: pt,
	}
	return result, nil
}

// RenameNotebook - Renames a notebook.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) renameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (*http.Response, error) {
	req, err := client.renameNotebookCreateRequest(ctx, notebookName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.renameNotebookHandleError(resp)
	}
	return resp, nil
}

// renameNotebookCreateRequest creates the RenameNotebook request.
func (client *notebookClient) renameNotebookCreateRequest(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (*policy.Request, error) {
	urlPath := "/notebooks/{notebookName}/rename"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// renameNotebookHandleError handles the RenameNotebook error response.
func (client *notebookClient) renameNotebookHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
