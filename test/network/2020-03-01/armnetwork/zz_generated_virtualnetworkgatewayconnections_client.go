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

// VirtualNetworkGatewayConnectionsClient contains the methods for the VirtualNetworkGatewayConnections group.
// Don't use this type directly, use NewVirtualNetworkGatewayConnectionsClient() instead.
type VirtualNetworkGatewayConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualNetworkGatewayConnectionsClient creates a new instance of VirtualNetworkGatewayConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualNetworkGatewayConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualNetworkGatewayConnectionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VirtualNetworkGatewayConnectionsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a virtual network gateway connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// virtualNetworkGatewayConnectionName - The name of the virtual network gateway connection.
// parameters - Parameters supplied to the create or update virtual network gateway connection operation.
// options - VirtualNetworkGatewayConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkGatewayConnectionsClient.BeginCreateOrUpdate
// method.
func (client *VirtualNetworkGatewayConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection, options *VirtualNetworkGatewayConnectionsClientBeginCreateOrUpdateOptions) (VirtualNetworkGatewayConnectionsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualNetworkGatewayConnectionsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkGatewayConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkGatewayConnectionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a virtual network gateway connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkGatewayConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection, options *VirtualNetworkGatewayConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
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
func (client *VirtualNetworkGatewayConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection, options *VirtualNetworkGatewayConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayConnectionName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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

// BeginDelete - Deletes the specified virtual network Gateway connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// virtualNetworkGatewayConnectionName - The name of the virtual network gateway connection.
// options - VirtualNetworkGatewayConnectionsClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkGatewayConnectionsClient.BeginDelete
// method.
func (client *VirtualNetworkGatewayConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientBeginDeleteOptions) (VirtualNetworkGatewayConnectionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientDeletePollerResponse{}, err
	}
	result := VirtualNetworkGatewayConnectionsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkGatewayConnectionsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientDeletePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkGatewayConnectionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified virtual network Gateway connection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkGatewayConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
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
func (client *VirtualNetworkGatewayConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayConnectionName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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

// Get - Gets the specified virtual network gateway connection by resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// virtualNetworkGatewayConnectionName - The name of the virtual network gateway connection.
// options - VirtualNetworkGatewayConnectionsClientGetOptions contains the optional parameters for the VirtualNetworkGatewayConnectionsClient.Get
// method.
func (client *VirtualNetworkGatewayConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientGetOptions) (VirtualNetworkGatewayConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkGatewayConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkGatewayConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayConnectionName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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

// getHandleResponse handles the Get response.
func (client *VirtualNetworkGatewayConnectionsClient) getHandleResponse(resp *http.Response) (VirtualNetworkGatewayConnectionsClientGetResponse, error) {
	result := VirtualNetworkGatewayConnectionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkGatewayConnection); err != nil {
		return VirtualNetworkGatewayConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// GetSharedKey - The Get VirtualNetworkGatewayConnectionSharedKey operation retrieves information about the specified virtual
// network gateway connection shared key through Network resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// virtualNetworkGatewayConnectionName - The virtual network gateway connection shared key name.
// options - VirtualNetworkGatewayConnectionsClientGetSharedKeyOptions contains the optional parameters for the VirtualNetworkGatewayConnectionsClient.GetSharedKey
// method.
func (client *VirtualNetworkGatewayConnectionsClient) GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientGetSharedKeyOptions) (VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse, error) {
	req, err := client.getSharedKeyCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSharedKeyHandleResponse(resp)
}

// getSharedKeyCreateRequest creates the GetSharedKey request.
func (client *VirtualNetworkGatewayConnectionsClient) getSharedKeyCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientGetSharedKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayConnectionName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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

// getSharedKeyHandleResponse handles the GetSharedKey response.
func (client *VirtualNetworkGatewayConnectionsClient) getSharedKeyHandleResponse(resp *http.Response) (VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse, error) {
	result := VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSharedKey); err != nil {
		return VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse{}, err
	}
	return result, nil
}

// List - The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - VirtualNetworkGatewayConnectionsClientListOptions contains the optional parameters for the VirtualNetworkGatewayConnectionsClient.List
// method.
func (client *VirtualNetworkGatewayConnectionsClient) List(resourceGroupName string, options *VirtualNetworkGatewayConnectionsClientListOptions) *VirtualNetworkGatewayConnectionsClientListPager {
	return &VirtualNetworkGatewayConnectionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp VirtualNetworkGatewayConnectionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkGatewayConnectionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualNetworkGatewayConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualNetworkGatewayConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections"
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
func (client *VirtualNetworkGatewayConnectionsClient) listHandleResponse(resp *http.Response) (VirtualNetworkGatewayConnectionsClientListResponse, error) {
	result := VirtualNetworkGatewayConnectionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkGatewayConnectionListResult); err != nil {
		return VirtualNetworkGatewayConnectionsClientListResponse{}, err
	}
	return result, nil
}

// BeginResetSharedKey - The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection
// shared key for passed virtual network gateway connection in the specified resource group
// through Network resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// virtualNetworkGatewayConnectionName - The virtual network gateway connection reset shared key Name.
// parameters - Parameters supplied to the begin reset virtual network gateway connection shared key operation through network
// resource provider.
// options - VirtualNetworkGatewayConnectionsClientBeginResetSharedKeyOptions contains the optional parameters for the VirtualNetworkGatewayConnectionsClient.BeginResetSharedKey
// method.
func (client *VirtualNetworkGatewayConnectionsClient) BeginResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey, options *VirtualNetworkGatewayConnectionsClientBeginResetSharedKeyOptions) (VirtualNetworkGatewayConnectionsClientResetSharedKeyPollerResponse, error) {
	resp, err := client.resetSharedKey(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientResetSharedKeyPollerResponse{}, err
	}
	result := VirtualNetworkGatewayConnectionsClientResetSharedKeyPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkGatewayConnectionsClient.ResetSharedKey", "location", resp, client.pl)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientResetSharedKeyPollerResponse{}, err
	}
	result.Poller = &VirtualNetworkGatewayConnectionsClientResetSharedKeyPoller{
		pt: pt,
	}
	return result, nil
}

// ResetSharedKey - The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection
// shared key for passed virtual network gateway connection in the specified resource group
// through Network resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkGatewayConnectionsClient) resetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey, options *VirtualNetworkGatewayConnectionsClientBeginResetSharedKeyOptions) (*http.Response, error) {
	req, err := client.resetSharedKeyCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
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

// resetSharedKeyCreateRequest creates the ResetSharedKey request.
func (client *VirtualNetworkGatewayConnectionsClient) resetSharedKeyCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey, options *VirtualNetworkGatewayConnectionsClientBeginResetSharedKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey/reset"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayConnectionName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginSetSharedKey - The Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection
// shared key for passed virtual network gateway connection in the specified resource group through
// Network resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// virtualNetworkGatewayConnectionName - The virtual network gateway connection name.
// parameters - Parameters supplied to the Begin Set Virtual Network Gateway connection Shared key operation throughNetwork
// resource provider.
// options - VirtualNetworkGatewayConnectionsClientBeginSetSharedKeyOptions contains the optional parameters for the VirtualNetworkGatewayConnectionsClient.BeginSetSharedKey
// method.
func (client *VirtualNetworkGatewayConnectionsClient) BeginSetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey, options *VirtualNetworkGatewayConnectionsClientBeginSetSharedKeyOptions) (VirtualNetworkGatewayConnectionsClientSetSharedKeyPollerResponse, error) {
	resp, err := client.setSharedKey(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientSetSharedKeyPollerResponse{}, err
	}
	result := VirtualNetworkGatewayConnectionsClientSetSharedKeyPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkGatewayConnectionsClient.SetSharedKey", "azure-async-operation", resp, client.pl)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientSetSharedKeyPollerResponse{}, err
	}
	result.Poller = &VirtualNetworkGatewayConnectionsClientSetSharedKeyPoller{
		pt: pt,
	}
	return result, nil
}

// SetSharedKey - The Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection shared
// key for passed virtual network gateway connection in the specified resource group through
// Network resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkGatewayConnectionsClient) setSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey, options *VirtualNetworkGatewayConnectionsClientBeginSetSharedKeyOptions) (*http.Response, error) {
	req, err := client.setSharedKeyCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
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

// setSharedKeyCreateRequest creates the SetSharedKey request.
func (client *VirtualNetworkGatewayConnectionsClient) setSharedKeyCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey, options *VirtualNetworkGatewayConnectionsClientBeginSetSharedKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayConnectionName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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

// BeginStartPacketCapture - Starts packet capture on virtual network gateway connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// virtualNetworkGatewayConnectionName - The name of the virtual network gateway connection.
// options - VirtualNetworkGatewayConnectionsClientBeginStartPacketCaptureOptions contains the optional parameters for the
// VirtualNetworkGatewayConnectionsClient.BeginStartPacketCapture method.
func (client *VirtualNetworkGatewayConnectionsClient) BeginStartPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientBeginStartPacketCaptureOptions) (VirtualNetworkGatewayConnectionsClientStartPacketCapturePollerResponse, error) {
	resp, err := client.startPacketCapture(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientStartPacketCapturePollerResponse{}, err
	}
	result := VirtualNetworkGatewayConnectionsClientStartPacketCapturePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkGatewayConnectionsClient.StartPacketCapture", "location", resp, client.pl)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientStartPacketCapturePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkGatewayConnectionsClientStartPacketCapturePoller{
		pt: pt,
	}
	return result, nil
}

// StartPacketCapture - Starts packet capture on virtual network gateway connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkGatewayConnectionsClient) startPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientBeginStartPacketCaptureOptions) (*http.Response, error) {
	req, err := client.startPacketCaptureCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
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

// startPacketCaptureCreateRequest creates the StartPacketCapture request.
func (client *VirtualNetworkGatewayConnectionsClient) startPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsClientBeginStartPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/startPacketCapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayConnectionName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// BeginStopPacketCapture - Stops packet capture on virtual network gateway connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// virtualNetworkGatewayConnectionName - The name of the virtual network gateway Connection.
// parameters - Virtual network gateway packet capture parameters supplied to stop packet capture on gateway connection.
// options - VirtualNetworkGatewayConnectionsClientBeginStopPacketCaptureOptions contains the optional parameters for the
// VirtualNetworkGatewayConnectionsClient.BeginStopPacketCapture method.
func (client *VirtualNetworkGatewayConnectionsClient) BeginStopPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VPNPacketCaptureStopParameters, options *VirtualNetworkGatewayConnectionsClientBeginStopPacketCaptureOptions) (VirtualNetworkGatewayConnectionsClientStopPacketCapturePollerResponse, error) {
	resp, err := client.stopPacketCapture(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientStopPacketCapturePollerResponse{}, err
	}
	result := VirtualNetworkGatewayConnectionsClientStopPacketCapturePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkGatewayConnectionsClient.StopPacketCapture", "location", resp, client.pl)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientStopPacketCapturePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkGatewayConnectionsClientStopPacketCapturePoller{
		pt: pt,
	}
	return result, nil
}

// StopPacketCapture - Stops packet capture on virtual network gateway connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkGatewayConnectionsClient) stopPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VPNPacketCaptureStopParameters, options *VirtualNetworkGatewayConnectionsClientBeginStopPacketCaptureOptions) (*http.Response, error) {
	req, err := client.stopPacketCaptureCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
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

// stopPacketCaptureCreateRequest creates the StopPacketCapture request.
func (client *VirtualNetworkGatewayConnectionsClient) stopPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VPNPacketCaptureStopParameters, options *VirtualNetworkGatewayConnectionsClientBeginStopPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/stopPacketCapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayConnectionName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginUpdateTags - Updates a virtual network gateway connection tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// virtualNetworkGatewayConnectionName - The name of the virtual network gateway connection.
// parameters - Parameters supplied to update virtual network gateway connection tags.
// options - VirtualNetworkGatewayConnectionsClientBeginUpdateTagsOptions contains the optional parameters for the VirtualNetworkGatewayConnectionsClient.BeginUpdateTags
// method.
func (client *VirtualNetworkGatewayConnectionsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject, options *VirtualNetworkGatewayConnectionsClientBeginUpdateTagsOptions) (VirtualNetworkGatewayConnectionsClientUpdateTagsPollerResponse, error) {
	resp, err := client.updateTags(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientUpdateTagsPollerResponse{}, err
	}
	result := VirtualNetworkGatewayConnectionsClientUpdateTagsPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkGatewayConnectionsClient.UpdateTags", "azure-async-operation", resp, client.pl)
	if err != nil {
		return VirtualNetworkGatewayConnectionsClientUpdateTagsPollerResponse{}, err
	}
	result.Poller = &VirtualNetworkGatewayConnectionsClientUpdateTagsPoller{
		pt: pt,
	}
	return result, nil
}

// UpdateTags - Updates a virtual network gateway connection tags.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkGatewayConnectionsClient) updateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject, options *VirtualNetworkGatewayConnectionsClientBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
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

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualNetworkGatewayConnectionsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject, options *VirtualNetworkGatewayConnectionsClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayConnectionName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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
