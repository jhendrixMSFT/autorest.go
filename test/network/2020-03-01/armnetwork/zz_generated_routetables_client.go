//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RouteTablesClient contains the methods for the RouteTables group.
// Don't use this type directly, use NewRouteTablesClient() instead.
type RouteTablesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRouteTablesClient creates a new instance of RouteTablesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRouteTablesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RouteTablesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &RouteTablesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Create or updates a route table in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// routeTableName - The name of the route table.
// parameters - Parameters supplied to the create or update route table operation.
// options - RouteTablesClientBeginCreateOrUpdateOptions contains the optional parameters for the RouteTablesClient.BeginCreateOrUpdate
// method.
func (client *RouteTablesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters RouteTable, options *RouteTablesClientBeginCreateOrUpdateOptions) (RouteTablesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, routeTableName, parameters, options)
	if err != nil {
		return RouteTablesClientCreateOrUpdatePollerResponse{}, err
	}
	result := RouteTablesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RouteTablesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return RouteTablesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &RouteTablesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or updates a route table in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RouteTablesClient) createOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters RouteTable, options *RouteTablesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, routeTableName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RouteTablesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, parameters RouteTable, options *RouteTablesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified route table.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// routeTableName - The name of the route table.
// options - RouteTablesClientBeginDeleteOptions contains the optional parameters for the RouteTablesClient.BeginDelete method.
func (client *RouteTablesClient) BeginDelete(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientBeginDeleteOptions) (RouteTablesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, routeTableName, options)
	if err != nil {
		return RouteTablesClientDeletePollerResponse{}, err
	}
	result := RouteTablesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RouteTablesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return RouteTablesClientDeletePollerResponse{}, err
	}
	result.Poller = &RouteTablesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified route table.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RouteTablesClient) deleteOperation(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, routeTableName, options)
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

// deleteCreateRequest creates the Delete request.
func (client *RouteTablesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified route table.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// routeTableName - The name of the route table.
// options - RouteTablesClientGetOptions contains the optional parameters for the RouteTablesClient.Get method.
func (client *RouteTablesClient) Get(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientGetOptions) (RouteTablesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, routeTableName, options)
	if err != nil {
		return RouteTablesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RouteTablesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RouteTablesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RouteTablesClient) getCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RouteTablesClient) getHandleResponse(resp *http.Response) (RouteTablesClientGetResponse, error) {
	result := RouteTablesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteTable); err != nil {
		return RouteTablesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all route tables in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - RouteTablesClientListOptions contains the optional parameters for the RouteTablesClient.List method.
func (client *RouteTablesClient) List(resourceGroupName string, options *RouteTablesClientListOptions) *RouteTablesClientListPager {
	return &RouteTablesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp RouteTablesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RouteTableListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RouteTablesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *RouteTablesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RouteTablesClient) listHandleResponse(resp *http.Response) (RouteTablesClientListResponse, error) {
	result := RouteTablesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteTableListResult); err != nil {
		return RouteTablesClientListResponse{}, err
	}
	return result, nil
}

// ListAll - Gets all route tables in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - RouteTablesClientListAllOptions contains the optional parameters for the RouteTablesClient.ListAll method.
func (client *RouteTablesClient) ListAll(options *RouteTablesClientListAllOptions) *RouteTablesClientListAllPager {
	return &RouteTablesClientListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp RouteTablesClientListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RouteTableListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *RouteTablesClient) listAllCreateRequest(ctx context.Context, options *RouteTablesClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeTables"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *RouteTablesClient) listAllHandleResponse(resp *http.Response) (RouteTablesClientListAllResponse, error) {
	result := RouteTablesClientListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteTableListResult); err != nil {
		return RouteTablesClientListAllResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a route table tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// routeTableName - The name of the route table.
// parameters - Parameters supplied to update route table tags.
// options - RouteTablesClientUpdateTagsOptions contains the optional parameters for the RouteTablesClient.UpdateTags method.
func (client *RouteTablesClient) UpdateTags(ctx context.Context, resourceGroupName string, routeTableName string, parameters TagsObject, options *RouteTablesClientUpdateTagsOptions) (RouteTablesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, routeTableName, parameters, options)
	if err != nil {
		return RouteTablesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RouteTablesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RouteTablesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *RouteTablesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, parameters TagsObject, options *RouteTablesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *RouteTablesClient) updateTagsHandleResponse(resp *http.Response) (RouteTablesClientUpdateTagsResponse, error) {
	result := RouteTablesClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteTable); err != nil {
		return RouteTablesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
