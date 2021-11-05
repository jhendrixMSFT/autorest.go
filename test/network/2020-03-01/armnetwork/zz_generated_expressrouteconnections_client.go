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

// ExpressRouteConnectionsClient contains the methods for the ExpressRouteConnections group.
// Don't use this type directly, use NewExpressRouteConnectionsClient() instead.
type ExpressRouteConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteConnectionsClient creates a new instance of ExpressRouteConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExpressRouteConnectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &ExpressRouteConnectionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// expressRouteGatewayName - The name of the ExpressRoute gateway.
// connectionName - The name of the connection subresource.
// putExpressRouteConnectionParameters - Parameters required in an ExpressRouteConnection PUT operation.
// options - ExpressRouteConnectionsBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteConnections.BeginCreateOrUpdate
// method.
func (client *ExpressRouteConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection, options *ExpressRouteConnectionsBeginCreateOrUpdateOptions) (ExpressRouteConnectionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, expressRouteGatewayName, connectionName, putExpressRouteConnectionParameters, options)
	if err != nil {
		return ExpressRouteConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRouteConnectionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExpressRouteConnectionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection, options *ExpressRouteConnectionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, connectionName, putExpressRouteConnectionParameters, options)
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
func (client *ExpressRouteConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection, options *ExpressRouteConnectionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
	return req, runtime.MarshalAsJSON(req, putExpressRouteConnectionParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteConnectionsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes a connection to a ExpressRoute circuit.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// expressRouteGatewayName - The name of the ExpressRoute gateway.
// connectionName - The name of the connection subresource.
// options - ExpressRouteConnectionsBeginDeleteOptions contains the optional parameters for the ExpressRouteConnections.BeginDelete
// method.
func (client *ExpressRouteConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsBeginDeleteOptions) (ExpressRouteConnectionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, expressRouteGatewayName, connectionName, options)
	if err != nil {
		return ExpressRouteConnectionsDeletePollerResponse{}, err
	}
	result := ExpressRouteConnectionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteConnectionsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ExpressRouteConnectionsDeletePollerResponse{}, err
	}
	result.Poller = &ExpressRouteConnectionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a connection to a ExpressRoute circuit.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, connectionName, options)
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
func (client *ExpressRouteConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
func (client *ExpressRouteConnectionsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the specified ExpressRouteConnection.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// expressRouteGatewayName - The name of the ExpressRoute gateway.
// connectionName - The name of the ExpressRoute connection.
// options - ExpressRouteConnectionsGetOptions contains the optional parameters for the ExpressRouteConnections.Get method.
func (client *ExpressRouteConnectionsClient) Get(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsGetOptions) (ExpressRouteConnectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, connectionName, options)
	if err != nil {
		return ExpressRouteConnectionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteConnectionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteConnectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
func (client *ExpressRouteConnectionsClient) getHandleResponse(resp *http.Response) (ExpressRouteConnectionsGetResponse, error) {
	result := ExpressRouteConnectionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteConnection); err != nil {
		return ExpressRouteConnectionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteConnectionsClient) getHandleError(resp *http.Response) error {
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

// List - Lists ExpressRouteConnections.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// expressRouteGatewayName - The name of the ExpressRoute gateway.
// options - ExpressRouteConnectionsListOptions contains the optional parameters for the ExpressRouteConnections.List method.
func (client *ExpressRouteConnectionsClient) List(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteConnectionsListOptions) (ExpressRouteConnectionsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, options)
	if err != nil {
		return ExpressRouteConnectionsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteConnectionsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteConnectionsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ExpressRouteConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteConnectionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
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
func (client *ExpressRouteConnectionsClient) listHandleResponse(resp *http.Response) (ExpressRouteConnectionsListResponse, error) {
	result := ExpressRouteConnectionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteConnectionList); err != nil {
		return ExpressRouteConnectionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteConnectionsClient) listHandleError(resp *http.Response) error {
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
