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

// NatGatewaysOperations contains the methods for the NatGateways group.
type NatGatewaysOperations interface {
	// BeginCreateOrUpdate - Creates or updates a nat gateway.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysCreateOrUpdateOptions) (*NatGatewayPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (NatGatewayPoller, error)
	// BeginDelete - Deletes the specified nat gateway.
	BeginDelete(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified nat gateway in a specified resource group.
	Get(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysGetOptions) (*NatGatewayResponse, error)
	// List - Gets all nat gateways in a resource group.
	List(resourceGroupName string, options *NatGatewaysListOptions) NatGatewayListResultPager
	// ListAll - Gets all the Nat Gateways in a subscription.
	ListAll(options *NatGatewaysListAllOptions) NatGatewayListResultPager
	// UpdateTags - Updates nat gateway tags.
	UpdateTags(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject, options *NatGatewaysUpdateTagsOptions) (*NatGatewayResponse, error)
}

// NatGatewaysClient implements the NatGatewaysOperations interface.
// Don't use this type directly, use NewNatGatewaysClient() instead.
type NatGatewaysClient struct {
	*Client
	subscriptionID string
}

// NewNatGatewaysClient creates a new instance of NatGatewaysClient with the specified values.
func NewNatGatewaysClient(c *Client, subscriptionID string) NatGatewaysOperations {
	return &NatGatewaysClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *NatGatewaysClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *NatGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysCreateOrUpdateOptions) (*NatGatewayPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, natGatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &NatGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NatGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &natGatewayPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*NatGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *NatGatewaysClient) ResumeCreateOrUpdate(token string) (NatGatewayPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NatGatewaysClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &natGatewayPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a nat gateway.
func (client *NatGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, natGatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NatGatewaysClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *NatGatewaysClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*NatGatewayResponse, error) {
	result := NatGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NatGateway)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NatGatewaysClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *NatGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, natGatewayName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NatGatewaysClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *NatGatewaysClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NatGatewaysClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified nat gateway.
func (client *NatGatewaysClient) Delete(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, natGatewayName, options)
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
func (client *NatGatewaysClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
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
func (client *NatGatewaysClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified nat gateway in a specified resource group.
func (client *NatGatewaysClient) Get(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysGetOptions) (*NatGatewayResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, natGatewayName, options)
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
func (client *NatGatewaysClient) GetCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
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
func (client *NatGatewaysClient) GetHandleResponse(resp *azcore.Response) (*NatGatewayResponse, error) {
	result := NatGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NatGateway)
}

// GetHandleError handles the Get error response.
func (client *NatGatewaysClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all nat gateways in a resource group.
func (client *NatGatewaysClient) List(resourceGroupName string, options *NatGatewaysListOptions) NatGatewayListResultPager {
	return &natGatewayListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *NatGatewayListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NatGatewayListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *NatGatewaysClient) ListCreateRequest(ctx context.Context, resourceGroupName string, options *NatGatewaysListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways"
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

// ListHandleResponse handles the List response.
func (client *NatGatewaysClient) ListHandleResponse(resp *azcore.Response) (*NatGatewayListResultResponse, error) {
	result := NatGatewayListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NatGatewayListResult)
}

// ListHandleError handles the List error response.
func (client *NatGatewaysClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the Nat Gateways in a subscription.
func (client *NatGatewaysClient) ListAll(options *NatGatewaysListAllOptions) NatGatewayListResultPager {
	return &natGatewayListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAllCreateRequest(ctx, options)
		},
		responder: client.ListAllHandleResponse,
		errorer:   client.ListAllHandleError,
		advancer: func(ctx context.Context, resp *NatGatewayListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NatGatewayListResult.NextLink)
		},
	}
}

// ListAllCreateRequest creates the ListAll request.
func (client *NatGatewaysClient) ListAllCreateRequest(ctx context.Context, options *NatGatewaysListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/natGateways"
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

// ListAllHandleResponse handles the ListAll response.
func (client *NatGatewaysClient) ListAllHandleResponse(resp *azcore.Response) (*NatGatewayListResultResponse, error) {
	result := NatGatewayListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NatGatewayListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *NatGatewaysClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates nat gateway tags.
func (client *NatGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject, options *NatGatewaysUpdateTagsOptions) (*NatGatewayResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, natGatewayName, parameters, options)
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
func (client *NatGatewaysClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject, options *NatGatewaysUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
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
func (client *NatGatewaysClient) UpdateTagsHandleResponse(resp *azcore.Response) (*NatGatewayResponse, error) {
	result := NatGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NatGateway)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *NatGatewaysClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
