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

// ExpressRouteCircuitsOperations contains the methods for the ExpressRouteCircuits group.
type ExpressRouteCircuitsOperations interface {
	// BeginCreateOrUpdate - Creates or updates an express route circuit.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsCreateOrUpdateOptions) (*ExpressRouteCircuitPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ExpressRouteCircuitPoller, error)
	// BeginDelete - Deletes the specified express route circuit.
	BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets information about the specified express route circuit.
	Get(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetOptions) (*ExpressRouteCircuitResponse, error)
	// GetPeeringStats - Gets all stats from an express route circuit in a resource group.
	GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitsGetPeeringStatsOptions) (*ExpressRouteCircuitStatsResponse, error)
	// GetStats - Gets all the stats from an express route circuit in a resource group.
	GetStats(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetStatsOptions) (*ExpressRouteCircuitStatsResponse, error)
	// List - Gets all the express route circuits in a resource group.
	List(resourceGroupName string, options *ExpressRouteCircuitsListOptions) ExpressRouteCircuitListResultPager
	// ListAll - Gets all the express route circuits in a subscription.
	ListAll(options *ExpressRouteCircuitsListAllOptions) ExpressRouteCircuitListResultPager
	// BeginListArpTable - Gets the currently advertised ARP table associated with the express route circuit in a resource group.
	BeginListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListArpTableOptions) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error)
	// ResumeListArpTable - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error)
	// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route circuit in a resource group.
	BeginListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableOptions) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error)
	// ResumeListRoutesTable - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error)
	// BeginListRoutesTableSummary - Gets the currently advertised routes table summary associated with the express route circuit in a resource group.
	BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableSummaryOptions) (*ExpressRouteCircuitsRoutesTableSummaryListResultPollerResponse, error)
	// ResumeListRoutesTableSummary - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListRoutesTableSummary(token string) (ExpressRouteCircuitsRoutesTableSummaryListResultPoller, error)
	// UpdateTags - Updates an express route circuit tags.
	UpdateTags(ctx context.Context, resourceGroupName string, circuitName string, parameters TagsObject, options *ExpressRouteCircuitsUpdateTagsOptions) (*ExpressRouteCircuitResponse, error)
}

// ExpressRouteCircuitsClient implements the ExpressRouteCircuitsOperations interface.
// Don't use this type directly, use NewExpressRouteCircuitsClient() instead.
type ExpressRouteCircuitsClient struct {
	*Client
	subscriptionID string
}

// NewExpressRouteCircuitsClient creates a new instance of ExpressRouteCircuitsClient with the specified values.
func NewExpressRouteCircuitsClient(c *Client, subscriptionID string) ExpressRouteCircuitsOperations {
	return &ExpressRouteCircuitsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ExpressRouteCircuitsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *ExpressRouteCircuitsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsCreateOrUpdateOptions) (*ExpressRouteCircuitPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, circuitName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCircuitsClient) ResumeCreateOrUpdate(token string) (ExpressRouteCircuitPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates an express route circuit.
func (client *ExpressRouteCircuitsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, circuitName, parameters, options)
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
func (client *ExpressRouteCircuitsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *ExpressRouteCircuitsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitResponse, error) {
	result := ExpressRouteCircuitResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuit)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCircuitsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRouteCircuitsClient) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, circuitName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *ExpressRouteCircuitsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified express route circuit.
func (client *ExpressRouteCircuitsClient) Delete(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, circuitName, options)
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
func (client *ExpressRouteCircuitsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *ExpressRouteCircuitsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets information about the specified express route circuit.
func (client *ExpressRouteCircuitsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetOptions) (*ExpressRouteCircuitResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, circuitName, options)
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
func (client *ExpressRouteCircuitsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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

// GetHandleResponse handles the Get response.
func (client *ExpressRouteCircuitsClient) GetHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitResponse, error) {
	result := ExpressRouteCircuitResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuit)
}

// GetHandleError handles the Get error response.
func (client *ExpressRouteCircuitsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetPeeringStats - Gets all stats from an express route circuit in a resource group.
func (client *ExpressRouteCircuitsClient) GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitsGetPeeringStatsOptions) (*ExpressRouteCircuitStatsResponse, error) {
	req, err := client.GetPeeringStatsCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetPeeringStatsHandleError(resp)
	}
	result, err := client.GetPeeringStatsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPeeringStatsCreateRequest creates the GetPeeringStats request.
func (client *ExpressRouteCircuitsClient) GetPeeringStatsCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitsGetPeeringStatsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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

// GetPeeringStatsHandleResponse handles the GetPeeringStats response.
func (client *ExpressRouteCircuitsClient) GetPeeringStatsHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitStatsResponse, error) {
	result := ExpressRouteCircuitStatsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitStats)
}

// GetPeeringStatsHandleError handles the GetPeeringStats error response.
func (client *ExpressRouteCircuitsClient) GetPeeringStatsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetStats - Gets all the stats from an express route circuit in a resource group.
func (client *ExpressRouteCircuitsClient) GetStats(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetStatsOptions) (*ExpressRouteCircuitStatsResponse, error) {
	req, err := client.GetStatsCreateRequest(ctx, resourceGroupName, circuitName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetStatsHandleError(resp)
	}
	result, err := client.GetStatsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStatsCreateRequest creates the GetStats request.
func (client *ExpressRouteCircuitsClient) GetStatsCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetStatsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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

// GetStatsHandleResponse handles the GetStats response.
func (client *ExpressRouteCircuitsClient) GetStatsHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitStatsResponse, error) {
	result := ExpressRouteCircuitStatsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitStats)
}

// GetStatsHandleError handles the GetStats error response.
func (client *ExpressRouteCircuitsClient) GetStatsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the express route circuits in a resource group.
func (client *ExpressRouteCircuitsClient) List(resourceGroupName string, options *ExpressRouteCircuitsListOptions) ExpressRouteCircuitListResultPager {
	return &expressRouteCircuitListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteCircuitListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCircuitListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *ExpressRouteCircuitsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteCircuitsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits"
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
func (client *ExpressRouteCircuitsClient) ListHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitListResultResponse, error) {
	result := ExpressRouteCircuitListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitListResult)
}

// ListHandleError handles the List error response.
func (client *ExpressRouteCircuitsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the express route circuits in a subscription.
func (client *ExpressRouteCircuitsClient) ListAll(options *ExpressRouteCircuitsListAllOptions) ExpressRouteCircuitListResultPager {
	return &expressRouteCircuitListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAllCreateRequest(ctx, options)
		},
		responder: client.ListAllHandleResponse,
		errorer:   client.ListAllHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteCircuitListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCircuitListResult.NextLink)
		},
	}
}

// ListAllCreateRequest creates the ListAll request.
func (client *ExpressRouteCircuitsClient) ListAllCreateRequest(ctx context.Context, options *ExpressRouteCircuitsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits"
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
func (client *ExpressRouteCircuitsClient) ListAllHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitListResultResponse, error) {
	result := ExpressRouteCircuitListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *ExpressRouteCircuitsClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRouteCircuitsClient) BeginListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListArpTableOptions) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	resp, err := client.ListArpTable(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsArpTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.ListArpTable", "location", resp, client.ListArpTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsArpTableListResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCircuitsClient) ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.ListArpTable", token, client.ListArpTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsArpTableListResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route circuit in a resource group.
func (client *ExpressRouteCircuitsClient) ListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListArpTableOptions) (*azcore.Response, error) {
	req, err := client.ListArpTableCreateRequest(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ListArpTableHandleError(resp)
	}
	return resp, nil
}

// ListArpTableCreateRequest creates the ListArpTable request.
func (client *ExpressRouteCircuitsClient) ListArpTableCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListArpTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/arpTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListArpTableHandleResponse handles the ListArpTable response.
func (client *ExpressRouteCircuitsClient) ListArpTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
	result := ExpressRouteCircuitsArpTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsArpTableListResult)
}

// ListArpTableHandleError handles the ListArpTable error response.
func (client *ExpressRouteCircuitsClient) ListArpTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRouteCircuitsClient) BeginListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableOptions) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	resp, err := client.ListRoutesTable(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsRoutesTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.ListRoutesTable", "location", resp, client.ListRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsRoutesTableListResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCircuitsClient) ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.ListRoutesTable", token, client.ListRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsRoutesTableListResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route circuit in a resource group.
func (client *ExpressRouteCircuitsClient) ListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableOptions) (*azcore.Response, error) {
	req, err := client.ListRoutesTableCreateRequest(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ListRoutesTableHandleError(resp)
	}
	return resp, nil
}

// ListRoutesTableCreateRequest creates the ListRoutesTable request.
func (client *ExpressRouteCircuitsClient) ListRoutesTableCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListRoutesTableHandleResponse handles the ListRoutesTable response.
func (client *ExpressRouteCircuitsClient) ListRoutesTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
	result := ExpressRouteCircuitsRoutesTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsRoutesTableListResult)
}

// ListRoutesTableHandleError handles the ListRoutesTable error response.
func (client *ExpressRouteCircuitsClient) ListRoutesTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRouteCircuitsClient) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableSummaryOptions) (*ExpressRouteCircuitsRoutesTableSummaryListResultPollerResponse, error) {
	resp, err := client.ListRoutesTableSummary(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsRoutesTableSummaryListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.ListRoutesTableSummary", "location", resp, client.ListRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsRoutesTableSummaryListResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsRoutesTableSummaryListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCircuitsClient) ResumeListRoutesTableSummary(token string) (ExpressRouteCircuitsRoutesTableSummaryListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.ListRoutesTableSummary", token, client.ListRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsRoutesTableSummaryListResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// ListRoutesTableSummary - Gets the currently advertised routes table summary associated with the express route circuit in a resource group.
func (client *ExpressRouteCircuitsClient) ListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableSummaryOptions) (*azcore.Response, error) {
	req, err := client.ListRoutesTableSummaryCreateRequest(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ListRoutesTableSummaryHandleError(resp)
	}
	return resp, nil
}

// ListRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client *ExpressRouteCircuitsClient) ListRoutesTableSummaryCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableSummaryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListRoutesTableSummaryHandleResponse handles the ListRoutesTableSummary response.
func (client *ExpressRouteCircuitsClient) ListRoutesTableSummaryHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsRoutesTableSummaryListResultResponse, error) {
	result := ExpressRouteCircuitsRoutesTableSummaryListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsRoutesTableSummaryListResult)
}

// ListRoutesTableSummaryHandleError handles the ListRoutesTableSummary error response.
func (client *ExpressRouteCircuitsClient) ListRoutesTableSummaryHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates an express route circuit tags.
func (client *ExpressRouteCircuitsClient) UpdateTags(ctx context.Context, resourceGroupName string, circuitName string, parameters TagsObject, options *ExpressRouteCircuitsUpdateTagsOptions) (*ExpressRouteCircuitResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, circuitName, parameters, options)
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
func (client *ExpressRouteCircuitsClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, parameters TagsObject, options *ExpressRouteCircuitsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *ExpressRouteCircuitsClient) UpdateTagsHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitResponse, error) {
	result := ExpressRouteCircuitResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuit)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *ExpressRouteCircuitsClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
