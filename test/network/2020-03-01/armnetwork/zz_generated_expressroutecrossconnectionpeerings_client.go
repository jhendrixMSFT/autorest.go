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
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRouteCrossConnectionPeeringsClient contains the methods for the ExpressRouteCrossConnectionPeerings group.
// Don't use this type directly, use NewExpressRouteCrossConnectionPeeringsClient() instead.
type ExpressRouteCrossConnectionPeeringsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteCrossConnectionPeeringsClient creates a new instance of ExpressRouteCrossConnectionPeeringsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteCrossConnectionPeeringsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExpressRouteCrossConnectionPeeringsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &ExpressRouteCrossConnectionPeeringsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// peeringParameters - Parameters supplied to the create or update ExpressRouteCrossConnection peering operation.
// options - ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteCrossConnectionPeerings.BeginCreateOrUpdate
// method.
func (client *ExpressRouteCrossConnectionPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions) (ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExpressRouteCrossConnectionPeeringsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteCrossConnectionPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
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
	return req, runtime.MarshalAsJSON(req, peeringParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified peering from the ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// options - ExpressRouteCrossConnectionPeeringsBeginDeleteOptions contains the optional parameters for the ExpressRouteCrossConnectionPeerings.BeginDelete
// method.
func (client *ExpressRouteCrossConnectionPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsBeginDeleteOptions) (ExpressRouteCrossConnectionPeeringsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, crossConnectionName, peeringName, options)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsDeletePollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionPeeringsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCrossConnectionPeeringsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsDeletePollerResponse{}, err
	}
	result.Poller = &ExpressRouteCrossConnectionPeeringsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified peering from the ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionPeeringsClient) deleteOperation(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExpressRouteCrossConnectionPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
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

// deleteHandleError handles the Delete error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified peering for the ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// options - ExpressRouteCrossConnectionPeeringsGetOptions contains the optional parameters for the ExpressRouteCrossConnectionPeerings.Get
// method.
func (client *ExpressRouteCrossConnectionPeeringsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsGetOptions) (ExpressRouteCrossConnectionPeeringsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCrossConnectionPeeringsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCrossConnectionPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
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
func (client *ExpressRouteCrossConnectionPeeringsClient) getHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionPeeringsGetResponse, error) {
	result := ExpressRouteCrossConnectionPeeringsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionPeering); err != nil {
		return ExpressRouteCrossConnectionPeeringsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all peerings in a specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// options - ExpressRouteCrossConnectionPeeringsListOptions contains the optional parameters for the ExpressRouteCrossConnectionPeerings.List
// method.
func (client *ExpressRouteCrossConnectionPeeringsClient) List(resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsListOptions) *ExpressRouteCrossConnectionPeeringsListPager {
	return &ExpressRouteCrossConnectionPeeringsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
		},
		advancer: func(ctx context.Context, resp ExpressRouteCrossConnectionPeeringsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionPeeringList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCrossConnectionPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings"
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
func (client *ExpressRouteCrossConnectionPeeringsClient) listHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionPeeringsListResponse, error) {
	result := ExpressRouteCrossConnectionPeeringsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionPeeringList); err != nil {
		return ExpressRouteCrossConnectionPeeringsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
