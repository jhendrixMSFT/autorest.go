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

// P2SVPNGatewaysClient contains the methods for the P2SVPNGateways group.
// Don't use this type directly, use NewP2SVPNGatewaysClient() instead.
type P2SVPNGatewaysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewP2SVPNGatewaysClient creates a new instance of P2SVPNGatewaysClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewP2SVPNGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *P2SVPNGatewaysClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &P2SVPNGatewaysClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a virtual wan p2s vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the P2SVpnGateway.
// gatewayName - The name of the gateway.
// p2SVPNGatewayParameters - Parameters supplied to create or Update a virtual wan p2s vpn gateway.
// options - P2SVPNGatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the P2SVPNGatewaysClient.BeginCreateOrUpdate
// method.
func (client *P2SVPNGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters P2SVPNGateway, options *P2SVPNGatewaysClientBeginCreateOrUpdateOptions) (P2SVPNGatewaysClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, p2SVPNGatewayParameters, options)
	if err != nil {
		return P2SVPNGatewaysClientCreateOrUpdatePollerResponse{}, err
	}
	result := P2SVPNGatewaysClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return P2SVPNGatewaysClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a virtual wan p2s vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *P2SVPNGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters P2SVPNGateway, options *P2SVPNGatewaysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, p2SVPNGatewayParameters, options)
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
func (client *P2SVPNGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters P2SVPNGateway, options *P2SVPNGatewaysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, p2SVPNGatewayParameters)
}

// BeginDelete - Deletes a virtual wan p2s vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the P2SVpnGateway.
// gatewayName - The name of the gateway.
// options - P2SVPNGatewaysClientBeginDeleteOptions contains the optional parameters for the P2SVPNGatewaysClient.BeginDelete
// method.
func (client *P2SVPNGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysClientBeginDeleteOptions) (P2SVPNGatewaysClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return P2SVPNGatewaysClientDeletePollerResponse{}, err
	}
	result := P2SVPNGatewaysClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.Delete", "location", resp, client.pl)
	if err != nil {
		return P2SVPNGatewaysClientDeletePollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a virtual wan p2s vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *P2SVPNGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, options)
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
func (client *P2SVPNGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// BeginDisconnectP2SVPNConnections - Disconnect P2S vpn connections of the virtual wan P2SVpnGateway in the specified resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// p2SVPNGatewayName - The name of the P2S Vpn Gateway.
// request - The parameters are supplied to disconnect p2s vpn connections.
// options - P2SVPNGatewaysClientBeginDisconnectP2SVPNConnectionsOptions contains the optional parameters for the P2SVPNGatewaysClient.BeginDisconnectP2SVPNConnections
// method.
func (client *P2SVPNGatewaysClient) BeginDisconnectP2SVPNConnections(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request P2SVPNConnectionRequest, options *P2SVPNGatewaysClientBeginDisconnectP2SVPNConnectionsOptions) (P2SVPNGatewaysClientDisconnectP2SVPNConnectionsPollerResponse, error) {
	resp, err := client.disconnectP2SVPNConnections(ctx, resourceGroupName, p2SVPNGatewayName, request, options)
	if err != nil {
		return P2SVPNGatewaysClientDisconnectP2SVPNConnectionsPollerResponse{}, err
	}
	result := P2SVPNGatewaysClientDisconnectP2SVPNConnectionsPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.DisconnectP2SVPNConnections", "location", resp, client.pl)
	if err != nil {
		return P2SVPNGatewaysClientDisconnectP2SVPNConnectionsPollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysClientDisconnectP2SVPNConnectionsPoller{
		pt: pt,
	}
	return result, nil
}

// DisconnectP2SVPNConnections - Disconnect P2S vpn connections of the virtual wan P2SVpnGateway in the specified resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *P2SVPNGatewaysClient) disconnectP2SVPNConnections(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request P2SVPNConnectionRequest, options *P2SVPNGatewaysClientBeginDisconnectP2SVPNConnectionsOptions) (*http.Response, error) {
	req, err := client.disconnectP2SVPNConnectionsCreateRequest(ctx, resourceGroupName, p2SVPNGatewayName, request, options)
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

// disconnectP2SVPNConnectionsCreateRequest creates the DisconnectP2SVPNConnections request.
func (client *P2SVPNGatewaysClient) disconnectP2SVPNConnectionsCreateRequest(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request P2SVPNConnectionRequest, options *P2SVPNGatewaysClientBeginDisconnectP2SVPNConnectionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{p2sVpnGatewayName}/disconnectP2sVpnConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if p2SVPNGatewayName == "" {
		return nil, errors.New("parameter p2SVPNGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{p2sVpnGatewayName}", url.PathEscape(p2SVPNGatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// BeginGenerateVPNProfile - Generates VPN profile for P2S client of the P2SVpnGateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// gatewayName - The name of the P2SVpnGateway.
// parameters - Parameters supplied to the generate P2SVpnGateway VPN client package operation.
// options - P2SVPNGatewaysClientBeginGenerateVPNProfileOptions contains the optional parameters for the P2SVPNGatewaysClient.BeginGenerateVPNProfile
// method.
func (client *P2SVPNGatewaysClient) BeginGenerateVPNProfile(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVPNProfileParameters, options *P2SVPNGatewaysClientBeginGenerateVPNProfileOptions) (P2SVPNGatewaysClientGenerateVPNProfilePollerResponse, error) {
	resp, err := client.generateVPNProfile(ctx, resourceGroupName, gatewayName, parameters, options)
	if err != nil {
		return P2SVPNGatewaysClientGenerateVPNProfilePollerResponse{}, err
	}
	result := P2SVPNGatewaysClientGenerateVPNProfilePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.GenerateVPNProfile", "location", resp, client.pl)
	if err != nil {
		return P2SVPNGatewaysClientGenerateVPNProfilePollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysClientGenerateVPNProfilePoller{
		pt: pt,
	}
	return result, nil
}

// GenerateVPNProfile - Generates VPN profile for P2S client of the P2SVpnGateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *P2SVPNGatewaysClient) generateVPNProfile(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVPNProfileParameters, options *P2SVPNGatewaysClientBeginGenerateVPNProfileOptions) (*http.Response, error) {
	req, err := client.generateVPNProfileCreateRequest(ctx, resourceGroupName, gatewayName, parameters, options)
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

// generateVPNProfileCreateRequest creates the GenerateVPNProfile request.
func (client *P2SVPNGatewaysClient) generateVPNProfileCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVPNProfileParameters, options *P2SVPNGatewaysClientBeginGenerateVPNProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/generatevpnprofile"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// Get - Retrieves the details of a virtual wan p2s vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the P2SVpnGateway.
// gatewayName - The name of the gateway.
// options - P2SVPNGatewaysClientGetOptions contains the optional parameters for the P2SVPNGatewaysClient.Get method.
func (client *P2SVPNGatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysClientGetOptions) (P2SVPNGatewaysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return P2SVPNGatewaysClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return P2SVPNGatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return P2SVPNGatewaysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *P2SVPNGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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
func (client *P2SVPNGatewaysClient) getHandleResponse(resp *http.Response) (P2SVPNGatewaysClientGetResponse, error) {
	result := P2SVPNGatewaysClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.P2SVPNGateway); err != nil {
		return P2SVPNGatewaysClientGetResponse{}, err
	}
	return result, nil
}

// BeginGetP2SVPNConnectionHealth - Gets the connection health of P2S clients of the virtual wan P2SVpnGateway in the specified
// resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// gatewayName - The name of the P2SVpnGateway.
// options - P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthOptions contains the optional parameters for the P2SVPNGatewaysClient.BeginGetP2SVPNConnectionHealth
// method.
func (client *P2SVPNGatewaysClient) BeginGetP2SVPNConnectionHealth(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthOptions) (P2SVPNGatewaysClientGetP2SVPNConnectionHealthPollerResponse, error) {
	resp, err := client.getP2SVPNConnectionHealth(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return P2SVPNGatewaysClientGetP2SVPNConnectionHealthPollerResponse{}, err
	}
	result := P2SVPNGatewaysClientGetP2SVPNConnectionHealthPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.GetP2SVPNConnectionHealth", "location", resp, client.pl)
	if err != nil {
		return P2SVPNGatewaysClientGetP2SVPNConnectionHealthPollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysClientGetP2SVPNConnectionHealthPoller{
		pt: pt,
	}
	return result, nil
}

// GetP2SVPNConnectionHealth - Gets the connection health of P2S clients of the virtual wan P2SVpnGateway in the specified
// resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealth(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthOptions) (*http.Response, error) {
	req, err := client.getP2SVPNConnectionHealthCreateRequest(ctx, resourceGroupName, gatewayName, options)
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

// getP2SVPNConnectionHealthCreateRequest creates the GetP2SVPNConnectionHealth request.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/getP2sVpnConnectionHealth"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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
	return req, nil
}

// BeginGetP2SVPNConnectionHealthDetailed - Gets the sas url to get the connection health detail of P2S clients of the virtual
// wan P2SVpnGateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// gatewayName - The name of the P2SVpnGateway.
// request - Request parameters supplied to get p2s vpn connections detailed health.
// options - P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthDetailedOptions contains the optional parameters for the P2SVPNGatewaysClient.BeginGetP2SVPNConnectionHealthDetailed
// method.
func (client *P2SVPNGatewaysClient) BeginGetP2SVPNConnectionHealthDetailed(ctx context.Context, resourceGroupName string, gatewayName string, request P2SVPNConnectionHealthRequest, options *P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthDetailedOptions) (P2SVPNGatewaysClientGetP2SVPNConnectionHealthDetailedPollerResponse, error) {
	resp, err := client.getP2SVPNConnectionHealthDetailed(ctx, resourceGroupName, gatewayName, request, options)
	if err != nil {
		return P2SVPNGatewaysClientGetP2SVPNConnectionHealthDetailedPollerResponse{}, err
	}
	result := P2SVPNGatewaysClientGetP2SVPNConnectionHealthDetailedPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.GetP2SVPNConnectionHealthDetailed", "location", resp, client.pl)
	if err != nil {
		return P2SVPNGatewaysClientGetP2SVPNConnectionHealthDetailedPollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysClientGetP2SVPNConnectionHealthDetailedPoller{
		pt: pt,
	}
	return result, nil
}

// GetP2SVPNConnectionHealthDetailed - Gets the sas url to get the connection health detail of P2S clients of the virtual
// wan P2SVpnGateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthDetailed(ctx context.Context, resourceGroupName string, gatewayName string, request P2SVPNConnectionHealthRequest, options *P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthDetailedOptions) (*http.Response, error) {
	req, err := client.getP2SVPNConnectionHealthDetailedCreateRequest(ctx, resourceGroupName, gatewayName, request, options)
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

// getP2SVPNConnectionHealthDetailedCreateRequest creates the GetP2SVPNConnectionHealthDetailed request.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthDetailedCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, request P2SVPNConnectionHealthRequest, options *P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthDetailedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/getP2sVpnConnectionHealthDetailed"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// List - Lists all the P2SVpnGateways in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - P2SVPNGatewaysClientListOptions contains the optional parameters for the P2SVPNGatewaysClient.List method.
func (client *P2SVPNGatewaysClient) List(options *P2SVPNGatewaysClientListOptions) *P2SVPNGatewaysClientListPager {
	return &P2SVPNGatewaysClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp P2SVPNGatewaysClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListP2SVPNGatewaysResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *P2SVPNGatewaysClient) listCreateRequest(ctx context.Context, options *P2SVPNGatewaysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/p2svpnGateways"
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
func (client *P2SVPNGatewaysClient) listHandleResponse(resp *http.Response) (P2SVPNGatewaysClientListResponse, error) {
	result := P2SVPNGatewaysClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListP2SVPNGatewaysResult); err != nil {
		return P2SVPNGatewaysClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all the P2SVpnGateways in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the P2SVpnGateway.
// options - P2SVPNGatewaysClientListByResourceGroupOptions contains the optional parameters for the P2SVPNGatewaysClient.ListByResourceGroup
// method.
func (client *P2SVPNGatewaysClient) ListByResourceGroup(resourceGroupName string, options *P2SVPNGatewaysClientListByResourceGroupOptions) *P2SVPNGatewaysClientListByResourceGroupPager {
	return &P2SVPNGatewaysClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp P2SVPNGatewaysClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListP2SVPNGatewaysResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *P2SVPNGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *P2SVPNGatewaysClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *P2SVPNGatewaysClient) listByResourceGroupHandleResponse(resp *http.Response) (P2SVPNGatewaysClientListByResourceGroupResponse, error) {
	result := P2SVPNGatewaysClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListP2SVPNGatewaysResult); err != nil {
		return P2SVPNGatewaysClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates virtual wan p2s vpn gateway tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the P2SVpnGateway.
// gatewayName - The name of the gateway.
// p2SVPNGatewayParameters - Parameters supplied to update a virtual wan p2s vpn gateway tags.
// options - P2SVPNGatewaysClientUpdateTagsOptions contains the optional parameters for the P2SVPNGatewaysClient.UpdateTags
// method.
func (client *P2SVPNGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters TagsObject, options *P2SVPNGatewaysClientUpdateTagsOptions) (P2SVPNGatewaysClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, gatewayName, p2SVPNGatewayParameters, options)
	if err != nil {
		return P2SVPNGatewaysClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return P2SVPNGatewaysClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return P2SVPNGatewaysClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *P2SVPNGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters TagsObject, options *P2SVPNGatewaysClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, p2SVPNGatewayParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *P2SVPNGatewaysClient) updateTagsHandleResponse(resp *http.Response) (P2SVPNGatewaysClientUpdateTagsResponse, error) {
	result := P2SVPNGatewaysClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.P2SVPNGateway); err != nil {
		return P2SVPNGatewaysClientUpdateTagsResponse{}, err
	}
	return result, nil
}
