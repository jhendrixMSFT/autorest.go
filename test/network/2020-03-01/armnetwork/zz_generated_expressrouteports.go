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

// ExpressRoutePortsOperations contains the methods for the ExpressRoutePorts group.
type ExpressRoutePortsOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified ExpressRoutePort resource.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsCreateOrUpdateOptions) (*ExpressRoutePortPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ExpressRoutePortPoller, error)
	// BeginDelete - Deletes the specified ExpressRoutePort resource.
	BeginDelete(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves the requested ExpressRoutePort resource.
	Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsGetOptions) (*ExpressRoutePortResponse, error)
	// List - List all the ExpressRoutePort resources in the specified subscription.
	List(options *ExpressRoutePortsListOptions) ExpressRoutePortListResultPager
	// ListByResourceGroup - List all the ExpressRoutePort resources in the specified resource group.
	ListByResourceGroup(resourceGroupName string, options *ExpressRoutePortsListByResourceGroupOptions) ExpressRoutePortListResultPager
	// UpdateTags - Update ExpressRoutePort tags.
	UpdateTags(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters TagsObject, options *ExpressRoutePortsUpdateTagsOptions) (*ExpressRoutePortResponse, error)
}

// ExpressRoutePortsClient implements the ExpressRoutePortsOperations interface.
// Don't use this type directly, use NewExpressRoutePortsClient() instead.
type ExpressRoutePortsClient struct {
	*Client
	subscriptionID string
}

// NewExpressRoutePortsClient creates a new instance of ExpressRoutePortsClient with the specified values.
func NewExpressRoutePortsClient(c *Client, subscriptionID string) ExpressRoutePortsOperations {
	return &ExpressRoutePortsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ExpressRoutePortsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *ExpressRoutePortsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsCreateOrUpdateOptions) (*ExpressRoutePortPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, expressRoutePortName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRoutePortPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRoutePortsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRoutePortPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRoutePortResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRoutePortsClient) ResumeCreateOrUpdate(token string) (ExpressRoutePortPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRoutePortsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRoutePortPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified ExpressRoutePort resource.
func (client *ExpressRoutePortsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, expressRoutePortName, parameters, options)
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
func (client *ExpressRoutePortsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ExpressRoutePortsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRoutePortResponse, error) {
	result := ExpressRoutePortResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRoutePort)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRoutePortsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRoutePortsClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, expressRoutePortName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRoutePortsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *ExpressRoutePortsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRoutePortsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified ExpressRoutePort resource.
func (client *ExpressRoutePortsClient) Delete(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
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
func (client *ExpressRoutePortsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
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
func (client *ExpressRoutePortsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the requested ExpressRoutePort resource.
func (client *ExpressRoutePortsClient) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsGetOptions) (*ExpressRoutePortResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
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
func (client *ExpressRoutePortsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
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

// GetHandleResponse handles the Get response.
func (client *ExpressRoutePortsClient) GetHandleResponse(resp *azcore.Response) (*ExpressRoutePortResponse, error) {
	result := ExpressRoutePortResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRoutePort)
}

// GetHandleError handles the Get error response.
func (client *ExpressRoutePortsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - List all the ExpressRoutePort resources in the specified subscription.
func (client *ExpressRoutePortsClient) List(options *ExpressRoutePortsListOptions) ExpressRoutePortListResultPager {
	return &expressRoutePortListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ExpressRoutePortListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRoutePortListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *ExpressRoutePortsClient) ListCreateRequest(ctx context.Context, options *ExpressRoutePortsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePorts"
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
func (client *ExpressRoutePortsClient) ListHandleResponse(resp *azcore.Response) (*ExpressRoutePortListResultResponse, error) {
	result := ExpressRoutePortListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRoutePortListResult)
}

// ListHandleError handles the List error response.
func (client *ExpressRoutePortsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - List all the ExpressRoutePort resources in the specified resource group.
func (client *ExpressRoutePortsClient) ListByResourceGroup(resourceGroupName string, options *ExpressRoutePortsListByResourceGroupOptions) ExpressRoutePortListResultPager {
	return &expressRoutePortListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *ExpressRoutePortListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRoutePortListResult.NextLink)
		},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRoutePortsClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRoutePortsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *ExpressRoutePortsClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*ExpressRoutePortListResultResponse, error) {
	result := ExpressRoutePortListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRoutePortListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ExpressRoutePortsClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Update ExpressRoutePort tags.
func (client *ExpressRoutePortsClient) UpdateTags(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters TagsObject, options *ExpressRoutePortsUpdateTagsOptions) (*ExpressRoutePortResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, expressRoutePortName, parameters, options)
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
func (client *ExpressRoutePortsClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters TagsObject, options *ExpressRoutePortsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
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
func (client *ExpressRoutePortsClient) UpdateTagsHandleResponse(resp *azcore.Response) (*ExpressRoutePortResponse, error) {
	result := ExpressRoutePortResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRoutePort)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *ExpressRoutePortsClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
