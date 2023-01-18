//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// NotebookClient contains the methods for the Notebook group.
// Don't use this type directly, use a constructor function instead.
type NotebookClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateOrUpdateNotebook - Creates or updates a Note Book.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - notebookName - The notebook name.
//   - notebook - Note book resource definition.
//   - options - NotebookClientBeginCreateOrUpdateNotebookOptions contains the optional parameters for the NotebookClient.BeginCreateOrUpdateNotebook
//     method.
func (client *NotebookClient) BeginCreateOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookClientBeginCreateOrUpdateNotebookOptions) (*runtime.Poller[NotebookClientCreateOrUpdateNotebookResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateNotebook(ctx, notebookName, notebook, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[NotebookClientCreateOrUpdateNotebookResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[NotebookClientCreateOrUpdateNotebookResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdateNotebook - Creates or updates a Note Book.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *NotebookClient) createOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookClientBeginCreateOrUpdateNotebookOptions) (*http.Response, error) {
	req, err := client.createOrUpdateNotebookCreateRequest(ctx, notebookName, notebook, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateNotebookCreateRequest creates the CreateOrUpdateNotebook request.
func (client *NotebookClient) createOrUpdateNotebookCreateRequest(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookClientBeginCreateOrUpdateNotebookOptions) (*policy.Request, error) {
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
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, notebook)
}

// BeginDeleteNotebook - Deletes a Note book.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - notebookName - The notebook name.
//   - options - NotebookClientBeginDeleteNotebookOptions contains the optional parameters for the NotebookClient.BeginDeleteNotebook
//     method.
func (client *NotebookClient) BeginDeleteNotebook(ctx context.Context, notebookName string, options *NotebookClientBeginDeleteNotebookOptions) (*runtime.Poller[NotebookClientDeleteNotebookResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteNotebook(ctx, notebookName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[NotebookClientDeleteNotebookResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[NotebookClientDeleteNotebookResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// DeleteNotebook - Deletes a Note book.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *NotebookClient) deleteNotebook(ctx context.Context, notebookName string, options *NotebookClientBeginDeleteNotebookOptions) (*http.Response, error) {
	req, err := client.deleteNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteNotebookCreateRequest creates the DeleteNotebook request.
func (client *NotebookClient) deleteNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookClientBeginDeleteNotebookOptions) (*policy.Request, error) {
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetNotebook - Gets a Note Book.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - notebookName - The notebook name.
//   - options - NotebookClientGetNotebookOptions contains the optional parameters for the NotebookClient.GetNotebook method.
func (client *NotebookClient) GetNotebook(ctx context.Context, notebookName string, options *NotebookClientGetNotebookOptions) (NotebookClientGetNotebookResponse, error) {
	req, err := client.getNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return NotebookClientGetNotebookResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotebookClientGetNotebookResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return NotebookClientGetNotebookResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNotebookHandleResponse(resp)
}

// getNotebookCreateRequest creates the GetNotebook request.
func (client *NotebookClient) getNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookClientGetNotebookOptions) (*policy.Request, error) {
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
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNotebookHandleResponse handles the GetNotebook response.
func (client *NotebookClient) getNotebookHandleResponse(resp *http.Response) (NotebookClientGetNotebookResponse, error) {
	result := NotebookClientGetNotebookResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookResource); err != nil {
		return NotebookClientGetNotebookResponse{}, err
	}
	return result, nil
}

// NewGetNotebookSummaryByWorkSpacePager - Lists a summary of Notebooks.
//
// Generated from API version 2019-06-01-preview
//   - options - NotebookClientGetNotebookSummaryByWorkSpaceOptions contains the optional parameters for the NotebookClient.NewGetNotebookSummaryByWorkSpacePager
//     method.
func (client *NotebookClient) NewGetNotebookSummaryByWorkSpacePager(options *NotebookClientGetNotebookSummaryByWorkSpaceOptions) *runtime.Pager[NotebookClientGetNotebookSummaryByWorkSpaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[NotebookClientGetNotebookSummaryByWorkSpaceResponse]{
		More: func(page NotebookClientGetNotebookSummaryByWorkSpaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NotebookClientGetNotebookSummaryByWorkSpaceResponse) (NotebookClientGetNotebookSummaryByWorkSpaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getNotebookSummaryByWorkSpaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NotebookClientGetNotebookSummaryByWorkSpaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NotebookClientGetNotebookSummaryByWorkSpaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NotebookClientGetNotebookSummaryByWorkSpaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getNotebookSummaryByWorkSpaceHandleResponse(resp)
		},
	})
}

// getNotebookSummaryByWorkSpaceCreateRequest creates the GetNotebookSummaryByWorkSpace request.
func (client *NotebookClient) getNotebookSummaryByWorkSpaceCreateRequest(ctx context.Context, options *NotebookClientGetNotebookSummaryByWorkSpaceOptions) (*policy.Request, error) {
	urlPath := "/notebooks/summary"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNotebookSummaryByWorkSpaceHandleResponse handles the GetNotebookSummaryByWorkSpace response.
func (client *NotebookClient) getNotebookSummaryByWorkSpaceHandleResponse(resp *http.Response) (NotebookClientGetNotebookSummaryByWorkSpaceResponse, error) {
	result := NotebookClientGetNotebookSummaryByWorkSpaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookListResponse); err != nil {
		return NotebookClientGetNotebookSummaryByWorkSpaceResponse{}, err
	}
	return result, nil
}

// NewGetNotebooksByWorkspacePager - Lists Notebooks.
//
// Generated from API version 2019-06-01-preview
//   - options - NotebookClientGetNotebooksByWorkspaceOptions contains the optional parameters for the NotebookClient.NewGetNotebooksByWorkspacePager
//     method.
func (client *NotebookClient) NewGetNotebooksByWorkspacePager(options *NotebookClientGetNotebooksByWorkspaceOptions) *runtime.Pager[NotebookClientGetNotebooksByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[NotebookClientGetNotebooksByWorkspaceResponse]{
		More: func(page NotebookClientGetNotebooksByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NotebookClientGetNotebooksByWorkspaceResponse) (NotebookClientGetNotebooksByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getNotebooksByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NotebookClientGetNotebooksByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NotebookClientGetNotebooksByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NotebookClientGetNotebooksByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getNotebooksByWorkspaceHandleResponse(resp)
		},
	})
}

// getNotebooksByWorkspaceCreateRequest creates the GetNotebooksByWorkspace request.
func (client *NotebookClient) getNotebooksByWorkspaceCreateRequest(ctx context.Context, options *NotebookClientGetNotebooksByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/notebooks"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNotebooksByWorkspaceHandleResponse handles the GetNotebooksByWorkspace response.
func (client *NotebookClient) getNotebooksByWorkspaceHandleResponse(resp *http.Response) (NotebookClientGetNotebooksByWorkspaceResponse, error) {
	result := NotebookClientGetNotebooksByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookListResponse); err != nil {
		return NotebookClientGetNotebooksByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameNotebook - Renames a notebook.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - notebookName - The notebook name.
//   - request - proposed new name.
//   - options - NotebookClientBeginRenameNotebookOptions contains the optional parameters for the NotebookClient.BeginRenameNotebook
//     method.
func (client *NotebookClient) BeginRenameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookClientBeginRenameNotebookOptions) (*runtime.Poller[NotebookClientRenameNotebookResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renameNotebook(ctx, notebookName, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[NotebookClientRenameNotebookResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[NotebookClientRenameNotebookResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RenameNotebook - Renames a notebook.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *NotebookClient) renameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookClientBeginRenameNotebookOptions) (*http.Response, error) {
	req, err := client.renameNotebookCreateRequest(ctx, notebookName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// renameNotebookCreateRequest creates the RenameNotebook request.
func (client *NotebookClient) renameNotebookCreateRequest(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookClientBeginRenameNotebookOptions) (*policy.Request, error) {
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}
