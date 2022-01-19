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

// ManagementClient contains the methods for the NetworkManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagementClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ManagementClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CheckDNSNameAvailability - Checks whether a domain name in the cloudapp.azure.com zone is available for use.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location of the domain name.
// domainNameLabel - The domain name to be verified. It must conform to the following regular expression: ^[a-z][a-z0-9-]{1,61}[a-z0-9]$.
// options - ManagementClientCheckDNSNameAvailabilityOptions contains the optional parameters for the ManagementClient.CheckDNSNameAvailability
// method.
func (client *ManagementClient) CheckDNSNameAvailability(ctx context.Context, location string, domainNameLabel string, options *ManagementClientCheckDNSNameAvailabilityOptions) (ManagementClientCheckDNSNameAvailabilityResponse, error) {
	req, err := client.checkDNSNameAvailabilityCreateRequest(ctx, location, domainNameLabel, options)
	if err != nil {
		return ManagementClientCheckDNSNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementClientCheckDNSNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientCheckDNSNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkDNSNameAvailabilityHandleResponse(resp)
}

// checkDNSNameAvailabilityCreateRequest creates the CheckDNSNameAvailability request.
func (client *ManagementClient) checkDNSNameAvailabilityCreateRequest(ctx context.Context, location string, domainNameLabel string, options *ManagementClientCheckDNSNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/CheckDnsNameAvailability"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("domainNameLabel", domainNameLabel)
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// checkDNSNameAvailabilityHandleResponse handles the CheckDNSNameAvailability response.
func (client *ManagementClient) checkDNSNameAvailabilityHandleResponse(resp *http.Response) (ManagementClientCheckDNSNameAvailabilityResponse, error) {
	result := ManagementClientCheckDNSNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DNSNameAvailabilityResult); err != nil {
		return ManagementClientCheckDNSNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginDeleteBastionShareableLink - Deletes the Bastion Shareable Links for all the VMs specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// bastionHostName - The name of the Bastion Host.
// bslRequest - Post request for all the Bastion Shareable Link endpoints.
// options - ManagementClientBeginDeleteBastionShareableLinkOptions contains the optional parameters for the ManagementClient.BeginDeleteBastionShareableLink
// method.
func (client *ManagementClient) BeginDeleteBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *ManagementClientBeginDeleteBastionShareableLinkOptions) (ManagementClientDeleteBastionShareableLinkPollerResponse, error) {
	resp, err := client.deleteBastionShareableLink(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return ManagementClientDeleteBastionShareableLinkPollerResponse{}, err
	}
	result := ManagementClientDeleteBastionShareableLinkPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagementClient.DeleteBastionShareableLink", "location", resp, client.pl)
	if err != nil {
		return ManagementClientDeleteBastionShareableLinkPollerResponse{}, err
	}
	result.Poller = &ManagementClientDeleteBastionShareableLinkPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteBastionShareableLink - Deletes the Bastion Shareable Links for all the VMs specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagementClient) deleteBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *ManagementClientBeginDeleteBastionShareableLinkOptions) (*http.Response, error) {
	req, err := client.deleteBastionShareableLinkCreateRequest(ctx, resourceGroupName, bastionHostName, bslRequest, options)
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

// deleteBastionShareableLinkCreateRequest creates the DeleteBastionShareableLink request.
func (client *ManagementClient) deleteBastionShareableLinkCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *ManagementClientBeginDeleteBastionShareableLinkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/deleteShareableLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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
	return req, runtime.MarshalAsJSON(req, bslRequest)
}

// DisconnectActiveSessions - Returns the list of currently active sessions on the Bastion.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// bastionHostName - The name of the Bastion Host.
// sessionIDs - The list of sessionids to disconnect.
// options - ManagementClientDisconnectActiveSessionsOptions contains the optional parameters for the ManagementClient.DisconnectActiveSessions
// method.
func (client *ManagementClient) DisconnectActiveSessions(resourceGroupName string, bastionHostName string, sessionIDs SessionIDs, options *ManagementClientDisconnectActiveSessionsOptions) *ManagementClientDisconnectActiveSessionsPager {
	return &ManagementClientDisconnectActiveSessionsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.disconnectActiveSessionsCreateRequest(ctx, resourceGroupName, bastionHostName, sessionIDs, options)
		},
		advancer: func(ctx context.Context, resp ManagementClientDisconnectActiveSessionsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BastionSessionDeleteResult.NextLink)
		},
	}
}

// disconnectActiveSessionsCreateRequest creates the DisconnectActiveSessions request.
func (client *ManagementClient) disconnectActiveSessionsCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, sessionIDs SessionIDs, options *ManagementClientDisconnectActiveSessionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/disconnectActiveSessions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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
	return req, runtime.MarshalAsJSON(req, sessionIDs)
}

// disconnectActiveSessionsHandleResponse handles the DisconnectActiveSessions response.
func (client *ManagementClient) disconnectActiveSessionsHandleResponse(resp *http.Response) (ManagementClientDisconnectActiveSessionsResponse, error) {
	result := ManagementClientDisconnectActiveSessionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BastionSessionDeleteResult); err != nil {
		return ManagementClientDisconnectActiveSessionsResponse{}, err
	}
	return result, nil
}

// BeginGeneratevirtualwanvpnserverconfigurationvpnprofile - Generates a unique VPN profile for P2S clients for VirtualWan
// and associated VpnServerConfiguration combination in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// virtualWANName - The name of the VirtualWAN whose associated VpnServerConfigurations is needed.
// vpnClientParams - Parameters supplied to the generate VirtualWan VPN profile generation operation.
// options - ManagementClientBeginGeneratevirtualwanvpnserverconfigurationvpnprofileOptions contains the optional parameters
// for the ManagementClient.BeginGeneratevirtualwanvpnserverconfigurationvpnprofile method.
func (client *ManagementClient) BeginGeneratevirtualwanvpnserverconfigurationvpnprofile(ctx context.Context, resourceGroupName string, virtualWANName string, vpnClientParams VirtualWanVPNProfileParameters, options *ManagementClientBeginGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (ManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofilePollerResponse, error) {
	resp, err := client.generatevirtualwanvpnserverconfigurationvpnprofile(ctx, resourceGroupName, virtualWANName, vpnClientParams, options)
	if err != nil {
		return ManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofilePollerResponse{}, err
	}
	result := ManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofilePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagementClient.Generatevirtualwanvpnserverconfigurationvpnprofile", "location", resp, client.pl)
	if err != nil {
		return ManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofilePollerResponse{}, err
	}
	result.Poller = &ManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofilePoller{
		pt: pt,
	}
	return result, nil
}

// Generatevirtualwanvpnserverconfigurationvpnprofile - Generates a unique VPN profile for P2S clients for VirtualWan and
// associated VpnServerConfiguration combination in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagementClient) generatevirtualwanvpnserverconfigurationvpnprofile(ctx context.Context, resourceGroupName string, virtualWANName string, vpnClientParams VirtualWanVPNProfileParameters, options *ManagementClientBeginGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (*http.Response, error) {
	req, err := client.generatevirtualwanvpnserverconfigurationvpnprofileCreateRequest(ctx, resourceGroupName, virtualWANName, vpnClientParams, options)
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

// generatevirtualwanvpnserverconfigurationvpnprofileCreateRequest creates the Generatevirtualwanvpnserverconfigurationvpnprofile request.
func (client *ManagementClient) generatevirtualwanvpnserverconfigurationvpnprofileCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, vpnClientParams VirtualWanVPNProfileParameters, options *ManagementClientBeginGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/GenerateVpnProfile"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWANName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnClientParams)
}

// BeginGetActiveSessions - Returns the list of currently active sessions on the Bastion.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// bastionHostName - The name of the Bastion Host.
// options - ManagementClientBeginGetActiveSessionsOptions contains the optional parameters for the ManagementClient.BeginGetActiveSessions
// method.
func (client *ManagementClient) BeginGetActiveSessions(ctx context.Context, resourceGroupName string, bastionHostName string, options *ManagementClientBeginGetActiveSessionsOptions) (ManagementClientGetActiveSessionsPollerResponse, error) {
	resp, err := client.getActiveSessions(ctx, resourceGroupName, bastionHostName, options)
	if err != nil {
		return ManagementClientGetActiveSessionsPollerResponse{}, err
	}
	result := ManagementClientGetActiveSessionsPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagementClient.GetActiveSessions", "location", resp, client.pl)
	if err != nil {
		return ManagementClientGetActiveSessionsPollerResponse{}, err
	}
	result.Poller = &ManagementClientGetActiveSessionsPoller{
		pt:     pt,
		client: client,
	}
	return result, nil
}

// GetActiveSessions - Returns the list of currently active sessions on the Bastion.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagementClient) getActiveSessions(ctx context.Context, resourceGroupName string, bastionHostName string, options *ManagementClientBeginGetActiveSessionsOptions) (*http.Response, error) {
	req, err := client.getActiveSessionsCreateRequest(ctx, resourceGroupName, bastionHostName, options)
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

// getActiveSessionsCreateRequest creates the GetActiveSessions request.
func (client *ManagementClient) getActiveSessionsCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, options *ManagementClientBeginGetActiveSessionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/getActiveSessions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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

// getActiveSessionsHandleResponse handles the GetActiveSessions response.
func (client *ManagementClient) getActiveSessionsHandleResponse(resp *http.Response) (ManagementClientGetActiveSessionsResponse, error) {
	result := ManagementClientGetActiveSessionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BastionActiveSessionListResult); err != nil {
		return ManagementClientGetActiveSessionsResponse{}, err
	}
	return result, nil
}

// GetBastionShareableLink - Return the Bastion Shareable Links for all the VMs specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// bastionHostName - The name of the Bastion Host.
// bslRequest - Post request for all the Bastion Shareable Link endpoints.
// options - ManagementClientGetBastionShareableLinkOptions contains the optional parameters for the ManagementClient.GetBastionShareableLink
// method.
func (client *ManagementClient) GetBastionShareableLink(resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *ManagementClientGetBastionShareableLinkOptions) *ManagementClientGetBastionShareableLinkPager {
	return &ManagementClientGetBastionShareableLinkPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getBastionShareableLinkCreateRequest(ctx, resourceGroupName, bastionHostName, bslRequest, options)
		},
		advancer: func(ctx context.Context, resp ManagementClientGetBastionShareableLinkResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BastionShareableLinkListResult.NextLink)
		},
	}
}

// getBastionShareableLinkCreateRequest creates the GetBastionShareableLink request.
func (client *ManagementClient) getBastionShareableLinkCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *ManagementClientGetBastionShareableLinkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/getShareableLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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
	return req, runtime.MarshalAsJSON(req, bslRequest)
}

// getBastionShareableLinkHandleResponse handles the GetBastionShareableLink response.
func (client *ManagementClient) getBastionShareableLinkHandleResponse(resp *http.Response) (ManagementClientGetBastionShareableLinkResponse, error) {
	result := ManagementClientGetBastionShareableLinkResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BastionShareableLinkListResult); err != nil {
		return ManagementClientGetBastionShareableLinkResponse{}, err
	}
	return result, nil
}

// BeginPutBastionShareableLink - Creates a Bastion Shareable Links for all the VMs specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// bastionHostName - The name of the Bastion Host.
// bslRequest - Post request for all the Bastion Shareable Link endpoints.
// options - ManagementClientBeginPutBastionShareableLinkOptions contains the optional parameters for the ManagementClient.BeginPutBastionShareableLink
// method.
func (client *ManagementClient) BeginPutBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *ManagementClientBeginPutBastionShareableLinkOptions) (ManagementClientPutBastionShareableLinkPollerResponse, error) {
	resp, err := client.putBastionShareableLink(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return ManagementClientPutBastionShareableLinkPollerResponse{}, err
	}
	result := ManagementClientPutBastionShareableLinkPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagementClient.PutBastionShareableLink", "location", resp, client.pl)
	if err != nil {
		return ManagementClientPutBastionShareableLinkPollerResponse{}, err
	}
	result.Poller = &ManagementClientPutBastionShareableLinkPoller{
		pt:     pt,
		client: client,
	}
	return result, nil
}

// PutBastionShareableLink - Creates a Bastion Shareable Links for all the VMs specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ManagementClient) putBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *ManagementClientBeginPutBastionShareableLinkOptions) (*http.Response, error) {
	req, err := client.putBastionShareableLinkCreateRequest(ctx, resourceGroupName, bastionHostName, bslRequest, options)
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

// putBastionShareableLinkCreateRequest creates the PutBastionShareableLink request.
func (client *ManagementClient) putBastionShareableLinkCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *ManagementClientBeginPutBastionShareableLinkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/createShareableLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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
	return req, runtime.MarshalAsJSON(req, bslRequest)
}

// putBastionShareableLinkHandleResponse handles the PutBastionShareableLink response.
func (client *ManagementClient) putBastionShareableLinkHandleResponse(resp *http.Response) (ManagementClientPutBastionShareableLinkResponse, error) {
	result := ManagementClientPutBastionShareableLinkResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BastionShareableLinkListResult); err != nil {
		return ManagementClientPutBastionShareableLinkResponse{}, err
	}
	return result, nil
}

// SupportedSecurityProviders - Gives the supported security providers for the virtual wan.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// virtualWANName - The name of the VirtualWAN for which supported security providers are needed.
// options - ManagementClientSupportedSecurityProvidersOptions contains the optional parameters for the ManagementClient.SupportedSecurityProviders
// method.
func (client *ManagementClient) SupportedSecurityProviders(ctx context.Context, resourceGroupName string, virtualWANName string, options *ManagementClientSupportedSecurityProvidersOptions) (ManagementClientSupportedSecurityProvidersResponse, error) {
	req, err := client.supportedSecurityProvidersCreateRequest(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return ManagementClientSupportedSecurityProvidersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementClientSupportedSecurityProvidersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientSupportedSecurityProvidersResponse{}, runtime.NewResponseError(resp)
	}
	return client.supportedSecurityProvidersHandleResponse(resp)
}

// supportedSecurityProvidersCreateRequest creates the SupportedSecurityProviders request.
func (client *ManagementClient) supportedSecurityProvidersCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, options *ManagementClientSupportedSecurityProvidersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/supportedSecurityProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWANName))
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

// supportedSecurityProvidersHandleResponse handles the SupportedSecurityProviders response.
func (client *ManagementClient) supportedSecurityProvidersHandleResponse(resp *http.Response) (ManagementClientSupportedSecurityProvidersResponse, error) {
	result := ManagementClientSupportedSecurityProvidersResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualWanSecurityProviders); err != nil {
		return ManagementClientSupportedSecurityProvidersResponse{}, err
	}
	return result, nil
}
