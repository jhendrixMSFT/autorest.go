// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// RouteFiltersOperations contains the methods for the RouteFilters group.
type RouteFiltersOperations interface {
	// BeginCreateOrUpdate - Creates or updates a route filter in a specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersCreateOrUpdateOptions) (*RouteFilterPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (RouteFilterPoller, error)
	// BeginDelete - Deletes the specified route filter.
	BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified route filter.
	Get(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersGetOptions) (*RouteFilterResponse, error)
	// List - Gets all route filters in a subscription.
	List(options *RouteFiltersListOptions) RouteFilterListResultPager
	// ListByResourceGroup - Gets all route filters in a resource group.
	ListByResourceGroup(resourceGroupName string, options *RouteFiltersListByResourceGroupOptions) RouteFilterListResultPager
	// UpdateTags - Updates tags of a route filter.
	UpdateTags(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject, options *RouteFiltersUpdateTagsOptions) (*RouteFilterResponse, error)
}

// RouteFiltersClient implements the RouteFiltersOperations interface.
// Don't use this type directly, use NewRouteFiltersClient() instead.
type RouteFiltersClient struct {
	*Client
	subscriptionID string
}

// NewRouteFiltersClient creates a new instance of RouteFiltersClient with the specified values.
func NewRouteFiltersClient(c *Client, subscriptionID string) RouteFiltersOperations {
	return &RouteFiltersClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *RouteFiltersClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *RouteFiltersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersCreateOrUpdateOptions) (*RouteFilterPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, routeFilterName, routeFilterParameters, options)
	if err != nil {
		return nil, err
	}
	result := &RouteFilterPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("RouteFiltersClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &routeFilterPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*RouteFilterResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *RouteFiltersClient) ResumeCreateOrUpdate(token string) (RouteFilterPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RouteFiltersClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &routeFilterPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a route filter in a specified resource group.
func (client *RouteFiltersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, routeFilterName, routeFilterParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RouteFiltersClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(routeFilterParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RouteFiltersClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*RouteFilterResponse, error) {
	result := RouteFilterResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilter)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RouteFiltersClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *RouteFiltersClient) BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, routeFilterName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("RouteFiltersClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *RouteFiltersClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RouteFiltersClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified route filter.
func (client *RouteFiltersClient) Delete(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, routeFilterName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *RouteFiltersClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *RouteFiltersClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified route filter.
func (client *RouteFiltersClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersGetOptions) (*RouteFilterResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, routeFilterName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *RouteFiltersClient) GetCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *RouteFiltersClient) GetHandleResponse(resp *azcore.Response) (*RouteFilterResponse, error) {
	result := RouteFilterResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilter)
}

// GetHandleError handles the Get error response.
func (client *RouteFiltersClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all route filters in a subscription.
func (client *RouteFiltersClient) List(options *RouteFiltersListOptions) RouteFilterListResultPager {
	return &routeFilterListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *RouteFilterListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RouteFilterListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *RouteFiltersClient) ListCreateRequest(ctx context.Context, options *RouteFiltersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *RouteFiltersClient) ListHandleResponse(resp *azcore.Response) (*RouteFilterListResultResponse, error) {
	result := RouteFilterListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilterListResult)
}

// ListHandleError handles the List error response.
func (client *RouteFiltersClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Gets all route filters in a resource group.
func (client *RouteFiltersClient) ListByResourceGroup(resourceGroupName string, options *RouteFiltersListByResourceGroupOptions) RouteFilterListResultPager {
	return &routeFilterListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *RouteFilterListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RouteFilterListResult.NextLink)
		},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *RouteFiltersClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RouteFiltersListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *RouteFiltersClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*RouteFilterListResultResponse, error) {
	result := RouteFilterListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilterListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *RouteFiltersClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates tags of a route filter.
func (client *RouteFiltersClient) UpdateTags(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject, options *RouteFiltersUpdateTagsOptions) (*RouteFilterResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, routeFilterName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *RouteFiltersClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject, options *RouteFiltersUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *RouteFiltersClient) UpdateTagsHandleResponse(resp *azcore.Response) (*RouteFilterResponse, error) {
	result := RouteFilterResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilter)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *RouteFiltersClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
