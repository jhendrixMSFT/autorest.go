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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type linkedServiceClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newLinkedServiceClient creates a new instance of linkedServiceClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newLinkedServiceClient(endpoint string, pl runtime.Pipeline) *linkedServiceClient {
	client := &linkedServiceClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// BeginCreateOrUpdateLinkedService - Creates or updates a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
// linkedServiceName - The linked service name.
// linkedService - Linked service resource definition.
// options - linkedServiceClientBeginCreateOrUpdateLinkedServiceOptions contains the optional parameters for the linkedServiceClient.BeginCreateOrUpdateLinkedService
// method.
func (client *linkedServiceClient) BeginCreateOrUpdateLinkedService(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *linkedServiceClientBeginCreateOrUpdateLinkedServiceOptions) (linkedServiceClientCreateOrUpdateLinkedServicePollerResponse, error) {
	resp, err := client.createOrUpdateLinkedService(ctx, linkedServiceName, linkedService, options)
	if err != nil {
		return linkedServiceClientCreateOrUpdateLinkedServicePollerResponse{}, err
	}
	result := linkedServiceClientCreateOrUpdateLinkedServicePollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("linkedServiceClient.CreateOrUpdateLinkedService", resp, client.pl)
	if err != nil {
		return linkedServiceClientCreateOrUpdateLinkedServicePollerResponse{}, err
	}
	result.Poller = &linkedServiceClientCreateOrUpdateLinkedServicePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdateLinkedService - Creates or updates a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *linkedServiceClient) createOrUpdateLinkedService(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *linkedServiceClientBeginCreateOrUpdateLinkedServiceOptions) (*http.Response, error) {
	req, err := client.createOrUpdateLinkedServiceCreateRequest(ctx, linkedServiceName, linkedService, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateLinkedServiceCreateRequest creates the CreateOrUpdateLinkedService request.
func (client *linkedServiceClient) createOrUpdateLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *linkedServiceClientBeginCreateOrUpdateLinkedServiceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
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
	return req, runtime.MarshalAsJSON(req, linkedService)
}

// BeginDeleteLinkedService - Deletes a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
// linkedServiceName - The linked service name.
// options - linkedServiceClientBeginDeleteLinkedServiceOptions contains the optional parameters for the linkedServiceClient.BeginDeleteLinkedService
// method.
func (client *linkedServiceClient) BeginDeleteLinkedService(ctx context.Context, linkedServiceName string, options *linkedServiceClientBeginDeleteLinkedServiceOptions) (linkedServiceClientDeleteLinkedServicePollerResponse, error) {
	resp, err := client.deleteLinkedService(ctx, linkedServiceName, options)
	if err != nil {
		return linkedServiceClientDeleteLinkedServicePollerResponse{}, err
	}
	result := linkedServiceClientDeleteLinkedServicePollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("linkedServiceClient.DeleteLinkedService", resp, client.pl)
	if err != nil {
		return linkedServiceClientDeleteLinkedServicePollerResponse{}, err
	}
	result.Poller = &linkedServiceClientDeleteLinkedServicePoller{
		pt: pt,
	}
	return result, nil
}

// DeleteLinkedService - Deletes a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *linkedServiceClient) deleteLinkedService(ctx context.Context, linkedServiceName string, options *linkedServiceClientBeginDeleteLinkedServiceOptions) (*http.Response, error) {
	req, err := client.deleteLinkedServiceCreateRequest(ctx, linkedServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteLinkedServiceCreateRequest creates the DeleteLinkedService request.
func (client *linkedServiceClient) deleteLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, options *linkedServiceClientBeginDeleteLinkedServiceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
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

// GetLinkedService - Gets a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
// linkedServiceName - The linked service name.
// options - linkedServiceClientGetLinkedServiceOptions contains the optional parameters for the linkedServiceClient.GetLinkedService
// method.
func (client *linkedServiceClient) GetLinkedService(ctx context.Context, linkedServiceName string, options *linkedServiceClientGetLinkedServiceOptions) (linkedServiceClientGetLinkedServiceResponse, error) {
	req, err := client.getLinkedServiceCreateRequest(ctx, linkedServiceName, options)
	if err != nil {
		return linkedServiceClientGetLinkedServiceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return linkedServiceClientGetLinkedServiceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return linkedServiceClientGetLinkedServiceResponse{}, runtime.NewResponseError(resp)
	}
	return client.getLinkedServiceHandleResponse(resp)
}

// getLinkedServiceCreateRequest creates the GetLinkedService request.
func (client *linkedServiceClient) getLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, options *linkedServiceClientGetLinkedServiceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
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

// getLinkedServiceHandleResponse handles the GetLinkedService response.
func (client *linkedServiceClient) getLinkedServiceHandleResponse(resp *http.Response) (linkedServiceClientGetLinkedServiceResponse, error) {
	result := linkedServiceClientGetLinkedServiceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServiceResource); err != nil {
		return linkedServiceClientGetLinkedServiceResponse{}, err
	}
	return result, nil
}

// GetLinkedServicesByWorkspace - Lists linked services.
// If the operation fails it returns an *azcore.ResponseError type.
// options - linkedServiceClientGetLinkedServicesByWorkspaceOptions contains the optional parameters for the linkedServiceClient.GetLinkedServicesByWorkspace
// method.
func (client *linkedServiceClient) GetLinkedServicesByWorkspace(options *linkedServiceClientGetLinkedServicesByWorkspaceOptions) *linkedServiceClientGetLinkedServicesByWorkspacePager {
	return &linkedServiceClientGetLinkedServicesByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getLinkedServicesByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp linkedServiceClientGetLinkedServicesByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LinkedServiceListResponse.NextLink)
		},
	}
}

// getLinkedServicesByWorkspaceCreateRequest creates the GetLinkedServicesByWorkspace request.
func (client *linkedServiceClient) getLinkedServicesByWorkspaceCreateRequest(ctx context.Context, options *linkedServiceClientGetLinkedServicesByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices"
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

// getLinkedServicesByWorkspaceHandleResponse handles the GetLinkedServicesByWorkspace response.
func (client *linkedServiceClient) getLinkedServicesByWorkspaceHandleResponse(resp *http.Response) (linkedServiceClientGetLinkedServicesByWorkspaceResponse, error) {
	result := linkedServiceClientGetLinkedServicesByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServiceListResponse); err != nil {
		return linkedServiceClientGetLinkedServicesByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameLinkedService - Renames a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
// linkedServiceName - The linked service name.
// request - proposed new name.
// options - linkedServiceClientBeginRenameLinkedServiceOptions contains the optional parameters for the linkedServiceClient.BeginRenameLinkedService
// method.
func (client *linkedServiceClient) BeginRenameLinkedService(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *linkedServiceClientBeginRenameLinkedServiceOptions) (linkedServiceClientRenameLinkedServicePollerResponse, error) {
	resp, err := client.renameLinkedService(ctx, linkedServiceName, request, options)
	if err != nil {
		return linkedServiceClientRenameLinkedServicePollerResponse{}, err
	}
	result := linkedServiceClientRenameLinkedServicePollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("linkedServiceClient.RenameLinkedService", resp, client.pl)
	if err != nil {
		return linkedServiceClientRenameLinkedServicePollerResponse{}, err
	}
	result.Poller = &linkedServiceClientRenameLinkedServicePoller{
		pt: pt,
	}
	return result, nil
}

// RenameLinkedService - Renames a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *linkedServiceClient) renameLinkedService(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *linkedServiceClientBeginRenameLinkedServiceOptions) (*http.Response, error) {
	req, err := client.renameLinkedServiceCreateRequest(ctx, linkedServiceName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// renameLinkedServiceCreateRequest creates the RenameLinkedService request.
func (client *linkedServiceClient) renameLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *linkedServiceClientBeginRenameLinkedServiceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}/rename"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
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
