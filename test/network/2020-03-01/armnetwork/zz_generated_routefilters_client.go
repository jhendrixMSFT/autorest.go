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

// RouteFiltersClient contains the methods for the RouteFilters group.
// Don't use this type directly, use NewRouteFiltersClient() instead.
type RouteFiltersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRouteFiltersClient creates a new instance of RouteFiltersClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRouteFiltersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RouteFiltersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &RouteFiltersClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a route filter in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// routeFilterName - The name of the route filter.
// routeFilterParameters - Parameters supplied to the create or update route filter operation.
// options - RouteFiltersClientBeginCreateOrUpdateOptions contains the optional parameters for the RouteFiltersClient.BeginCreateOrUpdate
// method.
func (client *RouteFiltersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersClientBeginCreateOrUpdateOptions) (RouteFiltersClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, routeFilterName, routeFilterParameters, options)
	if err != nil {
		return RouteFiltersClientCreateOrUpdatePollerResponse{}, err
	}
	result := RouteFiltersClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RouteFiltersClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return RouteFiltersClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &RouteFiltersClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a route filter in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RouteFiltersClient) createOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, routeFilterName, routeFilterParameters, options)
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
func (client *RouteFiltersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
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
	return req, runtime.MarshalAsJSON(req, routeFilterParameters)
}

// BeginDelete - Deletes the specified route filter.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// routeFilterName - The name of the route filter.
// options - RouteFiltersClientBeginDeleteOptions contains the optional parameters for the RouteFiltersClient.BeginDelete
// method.
func (client *RouteFiltersClient) BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientBeginDeleteOptions) (RouteFiltersClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, routeFilterName, options)
	if err != nil {
		return RouteFiltersClientDeletePollerResponse{}, err
	}
	result := RouteFiltersClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RouteFiltersClient.Delete", "location", resp, client.pl)
	if err != nil {
		return RouteFiltersClientDeletePollerResponse{}, err
	}
	result.Poller = &RouteFiltersClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified route filter.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RouteFiltersClient) deleteOperation(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, routeFilterName, options)
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
func (client *RouteFiltersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
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

// Get - Gets the specified route filter.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// routeFilterName - The name of the route filter.
// options - RouteFiltersClientGetOptions contains the optional parameters for the RouteFiltersClient.Get method.
func (client *RouteFiltersClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientGetOptions) (RouteFiltersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, routeFilterName, options)
	if err != nil {
		return RouteFiltersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RouteFiltersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RouteFiltersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RouteFiltersClient) getCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
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
func (client *RouteFiltersClient) getHandleResponse(resp *http.Response) (RouteFiltersClientGetResponse, error) {
	result := RouteFiltersClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilter); err != nil {
		return RouteFiltersClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all route filters in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - RouteFiltersClientListOptions contains the optional parameters for the RouteFiltersClient.List method.
func (client *RouteFiltersClient) List(options *RouteFiltersClientListOptions) *RouteFiltersClientListPager {
	return &RouteFiltersClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp RouteFiltersClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RouteFilterListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RouteFiltersClient) listCreateRequest(ctx context.Context, options *RouteFiltersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters"
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
func (client *RouteFiltersClient) listHandleResponse(resp *http.Response) (RouteFiltersClientListResponse, error) {
	result := RouteFiltersClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilterListResult); err != nil {
		return RouteFiltersClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets all route filters in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - RouteFiltersClientListByResourceGroupOptions contains the optional parameters for the RouteFiltersClient.ListByResourceGroup
// method.
func (client *RouteFiltersClient) ListByResourceGroup(resourceGroupName string, options *RouteFiltersClientListByResourceGroupOptions) *RouteFiltersClientListByResourceGroupPager {
	return &RouteFiltersClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp RouteFiltersClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RouteFilterListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *RouteFiltersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RouteFiltersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters"
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *RouteFiltersClient) listByResourceGroupHandleResponse(resp *http.Response) (RouteFiltersClientListByResourceGroupResponse, error) {
	result := RouteFiltersClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilterListResult); err != nil {
		return RouteFiltersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates tags of a route filter.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// routeFilterName - The name of the route filter.
// parameters - Parameters supplied to update route filter tags.
// options - RouteFiltersClientUpdateTagsOptions contains the optional parameters for the RouteFiltersClient.UpdateTags method.
func (client *RouteFiltersClient) UpdateTags(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject, options *RouteFiltersClientUpdateTagsOptions) (RouteFiltersClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, routeFilterName, parameters, options)
	if err != nil {
		return RouteFiltersClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RouteFiltersClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RouteFiltersClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *RouteFiltersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject, options *RouteFiltersClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
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
func (client *RouteFiltersClient) updateTagsHandleResponse(resp *http.Response) (RouteFiltersClientUpdateTagsResponse, error) {
	result := RouteFiltersClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilter); err != nil {
		return RouteFiltersClientUpdateTagsResponse{}, err
	}
	return result, nil
}
