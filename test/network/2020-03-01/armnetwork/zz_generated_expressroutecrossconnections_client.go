// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ExpressRouteCrossConnectionsClient contains the methods for the ExpressRouteCrossConnections group.
// Don't use this type directly, use NewExpressRouteCrossConnectionsClient() instead.
type ExpressRouteCrossConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCrossConnectionsClient creates a new instance of ExpressRouteCrossConnectionsClient with the specified values.
func NewExpressRouteCrossConnectionsClient(con *armcore.Connection, subscriptionID string) *ExpressRouteCrossConnectionsClient {
	return &ExpressRouteCrossConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Update the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsBeginCreateOrUpdateOptions) (ExpressRouteCrossConnectionPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return ExpressRouteCrossConnectionPollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionPollerResponse{}, err
	}
	poller := &expressRouteCrossConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCrossConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ExpressRouteCrossConnectionPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCrossConnectionPoller.ResumeToken().
func (client *ExpressRouteCrossConnectionsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ExpressRouteCrossConnectionPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionPollerResponse{}, err
	}
	poller := &expressRouteCrossConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteCrossConnectionPollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCrossConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Update the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteCrossConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCrossConnectionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets details about the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsGetOptions) (ExpressRouteCrossConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
	if err != nil {
		return ExpressRouteCrossConnectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteCrossConnectionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCrossConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteCrossConnectionsClient) getHandleResponse(resp *azcore.Response) (ExpressRouteCrossConnectionResponse, error) {
	var val *ExpressRouteCrossConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCrossConnectionResponse{}, err
	}
	return ExpressRouteCrossConnectionResponse{RawResponse: resp.Response, ExpressRouteCrossConnection: val}, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteCrossConnectionsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Retrieves all the ExpressRouteCrossConnections in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) List(options *ExpressRouteCrossConnectionsListOptions) ExpressRouteCrossConnectionListResultPager {
	return &expressRouteCrossConnectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ExpressRouteCrossConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCrossConnectionsClient) listCreateRequest(ctx context.Context, options *ExpressRouteCrossConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteCrossConnectionsClient) listHandleResponse(resp *azcore.Response) (ExpressRouteCrossConnectionListResultResponse, error) {
	var val *ExpressRouteCrossConnectionListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCrossConnectionListResultResponse{}, err
	}
	return ExpressRouteCrossConnectionListResultResponse{RawResponse: resp.Response, ExpressRouteCrossConnectionListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteCrossConnectionsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) BeginListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsBeginListArpTableOptions) (ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	resp, err := client.listArpTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return ExpressRouteCircuitsArpTableListResultPollerResponse{}, err
	}
	result := ExpressRouteCircuitsArpTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.ListArpTable", "location", resp, client.listArpTableHandleError)
	if err != nil {
		return ExpressRouteCircuitsArpTableListResultPollerResponse{}, err
	}
	poller := &expressRouteCircuitsArpTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCircuitsArpTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListArpTable creates a new ExpressRouteCircuitsArpTableListResultPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitsArpTableListResultPoller.ResumeToken().
func (client *ExpressRouteCrossConnectionsClient) ResumeListArpTable(ctx context.Context, token string) (ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.ListArpTable", token, client.listArpTableHandleError)
	if err != nil {
		return ExpressRouteCircuitsArpTableListResultPollerResponse{}, err
	}
	poller := &expressRouteCircuitsArpTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteCircuitsArpTableListResultPollerResponse{}, err
	}
	result := ExpressRouteCircuitsArpTableListResultPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCircuitsArpTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) listArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsBeginListArpTableOptions) (*azcore.Response, error) {
	req, err := client.listArpTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listArpTableHandleError(resp)
	}
	return resp, nil
}

// listArpTableCreateRequest creates the ListArpTable request.
func (client *ExpressRouteCrossConnectionsClient) listArpTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsBeginListArpTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/arpTables/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listArpTableHandleError handles the ListArpTable error response.
func (client *ExpressRouteCrossConnectionsClient) listArpTableHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - Retrieves all the ExpressRouteCrossConnections in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) ListByResourceGroup(resourceGroupName string, options *ExpressRouteCrossConnectionsListByResourceGroupOptions) ExpressRouteCrossConnectionListResultPager {
	return &expressRouteCrossConnectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ExpressRouteCrossConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRouteCrossConnectionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteCrossConnectionsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ExpressRouteCrossConnectionsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ExpressRouteCrossConnectionListResultResponse, error) {
	var val *ExpressRouteCrossConnectionListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCrossConnectionListResultResponse{}, err
	}
	return ExpressRouteCrossConnectionListResultResponse{RawResponse: resp.Response, ExpressRouteCrossConnectionListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ExpressRouteCrossConnectionsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsBeginListRoutesTableOptions) (ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	resp, err := client.listRoutesTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return ExpressRouteCircuitsRoutesTableListResultPollerResponse{}, err
	}
	result := ExpressRouteCircuitsRoutesTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.ListRoutesTable", "location", resp, client.listRoutesTableHandleError)
	if err != nil {
		return ExpressRouteCircuitsRoutesTableListResultPollerResponse{}, err
	}
	poller := &expressRouteCircuitsRoutesTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCircuitsRoutesTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListRoutesTable creates a new ExpressRouteCircuitsRoutesTableListResultPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitsRoutesTableListResultPoller.ResumeToken().
func (client *ExpressRouteCrossConnectionsClient) ResumeListRoutesTable(ctx context.Context, token string) (ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.ListRoutesTable", token, client.listRoutesTableHandleError)
	if err != nil {
		return ExpressRouteCircuitsRoutesTableListResultPollerResponse{}, err
	}
	poller := &expressRouteCircuitsRoutesTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteCircuitsRoutesTableListResultPollerResponse{}, err
	}
	result := ExpressRouteCircuitsRoutesTableListResultPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCircuitsRoutesTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsBeginListRoutesTableOptions) (*azcore.Response, error) {
	req, err := client.listRoutesTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listRoutesTableHandleError(resp)
	}
	return resp, nil
}

// listRoutesTableCreateRequest creates the ListRoutesTable request.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsBeginListRoutesTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTables/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listRoutesTableHandleError handles the ListRoutesTable error response.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsBeginListRoutesTableSummaryOptions) (ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse, error) {
	resp, err := client.listRoutesTableSummary(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.ListRoutesTableSummary", "location", resp, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse{}, err
	}
	poller := &expressRouteCrossConnectionsRoutesTableSummaryListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCrossConnectionsRoutesTableSummaryListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListRoutesTableSummary creates a new ExpressRouteCrossConnectionsRoutesTableSummaryListResultPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCrossConnectionsRoutesTableSummaryListResultPoller.ResumeToken().
func (client *ExpressRouteCrossConnectionsClient) ResumeListRoutesTableSummary(ctx context.Context, token string) (ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.ListRoutesTableSummary", token, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse{}, err
	}
	poller := &expressRouteCrossConnectionsRoutesTableSummaryListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCrossConnectionsRoutesTableSummaryListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsBeginListRoutesTableSummaryOptions) (*azcore.Response, error) {
	req, err := client.listRoutesTableSummaryCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listRoutesTableSummaryHandleError(resp)
	}
	return resp, nil
}

// listRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableSummaryCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsBeginListRoutesTableSummaryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listRoutesTableSummaryHandleError handles the ListRoutesTableSummary error response.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableSummaryHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UpdateTags - Updates an express route cross connection tags.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsUpdateTagsOptions) (ExpressRouteCrossConnectionResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, crossConnectionName, crossConnectionParameters, options)
	if err != nil {
		return ExpressRouteCrossConnectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteCrossConnectionResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ExpressRouteCrossConnectionsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(crossConnectionParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ExpressRouteCrossConnectionsClient) updateTagsHandleResponse(resp *azcore.Response) (ExpressRouteCrossConnectionResponse, error) {
	var val *ExpressRouteCrossConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCrossConnectionResponse{}, err
	}
	return ExpressRouteCrossConnectionResponse{RawResponse: resp.Response, ExpressRouteCrossConnection: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *ExpressRouteCrossConnectionsClient) updateTagsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
