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

// VirtualNetworkTapsClient contains the methods for the VirtualNetworkTaps group.
// Don't use this type directly, use NewVirtualNetworkTapsClient() instead.
type VirtualNetworkTapsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualNetworkTapsClient creates a new instance of VirtualNetworkTapsClient with the specified values.
func NewVirtualNetworkTapsClient(con *armcore.Connection, subscriptionID string) *VirtualNetworkTapsClient {
	return &VirtualNetworkTapsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a Virtual Network Tap.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkTapsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap, options *VirtualNetworkTapsBeginCreateOrUpdateOptions) (VirtualNetworkTapPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, tapName, parameters, options)
	if err != nil {
		return VirtualNetworkTapPollerResponse{}, err
	}
	result := VirtualNetworkTapPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkTapsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualNetworkTapPollerResponse{}, err
	}
	poller := &virtualNetworkTapPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualNetworkTapResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualNetworkTapPoller from the specified resume token.
// token - The value must come from a previous call to VirtualNetworkTapPoller.ResumeToken().
func (client *VirtualNetworkTapsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (VirtualNetworkTapPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkTapsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualNetworkTapPollerResponse{}, err
	}
	poller := &virtualNetworkTapPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualNetworkTapPollerResponse{}, err
	}
	result := VirtualNetworkTapPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualNetworkTapResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a Virtual Network Tap.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkTapsClient) createOrUpdate(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap, options *VirtualNetworkTapsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, tapName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkTapsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap, options *VirtualNetworkTapsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if tapName == "" {
		return nil, errors.New("parameter tapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
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
func (client *VirtualNetworkTapsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified virtual network tap.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkTapsClient) BeginDelete(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, tapName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkTapsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VirtualNetworkTapsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkTapsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified virtual network tap.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkTapsClient) deleteOperation(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, tapName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkTapsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if tapName == "" {
		return nil, errors.New("parameter tapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteHandleError handles the Delete error response.
func (client *VirtualNetworkTapsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets information about the specified virtual network tap.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkTapsClient) Get(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsGetOptions) (VirtualNetworkTapResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, tapName, options)
	if err != nil {
		return VirtualNetworkTapResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkTapResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualNetworkTapResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkTapsClient) getCreateRequest(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if tapName == "" {
		return nil, errors.New("parameter tapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
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
func (client *VirtualNetworkTapsClient) getHandleResponse(resp *azcore.Response) (VirtualNetworkTapResponse, error) {
	var val *VirtualNetworkTap
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualNetworkTapResponse{}, err
	}
	return VirtualNetworkTapResponse{RawResponse: resp.Response, VirtualNetworkTap: val}, nil
}

// getHandleError handles the Get error response.
func (client *VirtualNetworkTapsClient) getHandleError(resp *azcore.Response) error {
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

// ListAll - Gets all the VirtualNetworkTaps in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkTapsClient) ListAll(options *VirtualNetworkTapsListAllOptions) VirtualNetworkTapListResultPager {
	return &virtualNetworkTapListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp VirtualNetworkTapListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkTapListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *VirtualNetworkTapsClient) listAllCreateRequest(ctx context.Context, options *VirtualNetworkTapsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworkTaps"
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

// listAllHandleResponse handles the ListAll response.
func (client *VirtualNetworkTapsClient) listAllHandleResponse(resp *azcore.Response) (VirtualNetworkTapListResultResponse, error) {
	var val *VirtualNetworkTapListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualNetworkTapListResultResponse{}, err
	}
	return VirtualNetworkTapListResultResponse{RawResponse: resp.Response, VirtualNetworkTapListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client *VirtualNetworkTapsClient) listAllHandleError(resp *azcore.Response) error {
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

// ListByResourceGroup - Gets all the VirtualNetworkTaps in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkTapsClient) ListByResourceGroup(resourceGroupName string, options *VirtualNetworkTapsListByResourceGroupOptions) VirtualNetworkTapListResultPager {
	return &virtualNetworkTapListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp VirtualNetworkTapListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkTapListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualNetworkTapsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualNetworkTapsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps"
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
func (client *VirtualNetworkTapsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (VirtualNetworkTapListResultResponse, error) {
	var val *VirtualNetworkTapListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualNetworkTapListResultResponse{}, err
	}
	return VirtualNetworkTapListResultResponse{RawResponse: resp.Response, VirtualNetworkTapListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VirtualNetworkTapsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
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

// UpdateTags - Updates an VirtualNetworkTap tags.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkTapsClient) UpdateTags(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject, options *VirtualNetworkTapsUpdateTagsOptions) (VirtualNetworkTapResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, tapName, tapParameters, options)
	if err != nil {
		return VirtualNetworkTapResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkTapResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualNetworkTapResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualNetworkTapsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject, options *VirtualNetworkTapsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if tapName == "" {
		return nil, errors.New("parameter tapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
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
	return req, req.MarshalAsJSON(tapParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualNetworkTapsClient) updateTagsHandleResponse(resp *azcore.Response) (VirtualNetworkTapResponse, error) {
	var val *VirtualNetworkTap
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualNetworkTapResponse{}, err
	}
	return VirtualNetworkTapResponse{RawResponse: resp.Response, VirtualNetworkTap: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *VirtualNetworkTapsClient) updateTagsHandleError(resp *azcore.Response) error {
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
