//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRouteConnectionsClient contains the methods for the ExpressRouteConnections group.
// Don't use this type directly, use NewExpressRouteConnectionsClient() instead.
type ExpressRouteConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExpressRouteConnectionsClient creates a new instance of ExpressRouteConnectionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExpressRouteConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRouteConnectionsClient, error) {
	cl, err := arm.NewClient("armnetwork.ExpressRouteConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExpressRouteConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - expressRouteGatewayName - The name of the ExpressRoute gateway.
//   - connectionName - The name of the connection subresource.
//   - putExpressRouteConnectionParameters - Parameters required in an ExpressRouteConnection PUT operation.
//   - options - ExpressRouteConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteConnectionsClient.BeginCreateOrUpdate
//     method.
func (client *ExpressRouteConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection, options *ExpressRouteConnectionsClientBeginCreateOrUpdateOptions) (result *runtime.Poller[ExpressRouteConnectionsClientCreateOrUpdateResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ExpressRouteConnectionsClient.BeginCreateOrUpdate", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	if options == nil || options.ResumeToken == "" {
		var resp *http.Response
		resp, err = client.createOrUpdate(ctx, resourceGroupName, expressRouteGatewayName, connectionName, putExpressRouteConnectionParameters, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExpressRouteConnectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken[ExpressRouteConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// CreateOrUpdate - Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *ExpressRouteConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection, options *ExpressRouteConnectionsClientBeginCreateOrUpdateOptions) (resp *http.Response, err error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, connectionName, putExpressRouteConnectionParameters, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection, options *ExpressRouteConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, putExpressRouteConnectionParameters)
}

// BeginDelete - Deletes a connection to a ExpressRoute circuit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - expressRouteGatewayName - The name of the ExpressRoute gateway.
//   - connectionName - The name of the connection subresource.
//   - options - ExpressRouteConnectionsClientBeginDeleteOptions contains the optional parameters for the ExpressRouteConnectionsClient.BeginDelete
//     method.
func (client *ExpressRouteConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsClientBeginDeleteOptions) (result *runtime.Poller[ExpressRouteConnectionsClientDeleteResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ExpressRouteConnectionsClient.BeginDelete", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	if options == nil || options.ResumeToken == "" {
		var resp *http.Response
		resp, err = client.deleteOperation(ctx, resourceGroupName, expressRouteGatewayName, connectionName, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExpressRouteConnectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken[ExpressRouteConnectionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// Delete - Deletes a connection to a ExpressRoute circuit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *ExpressRouteConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsClientBeginDeleteOptions) (resp *http.Response, err error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, connectionName, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// deleteCreateRequest creates the Delete request.
func (client *ExpressRouteConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified ExpressRouteConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - expressRouteGatewayName - The name of the ExpressRoute gateway.
//   - connectionName - The name of the ExpressRoute connection.
//   - options - ExpressRouteConnectionsClientGetOptions contains the optional parameters for the ExpressRouteConnectionsClient.Get
//     method.
func (client *ExpressRouteConnectionsClient) Get(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsClientGetOptions) (result ExpressRouteConnectionsClientGetResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ExpressRouteConnectionsClient.Get", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, connectionName, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	result, err = client.getHandleResponse(resp)
	return
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsClientGetOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteConnectionsClient) getHandleResponse(resp *http.Response) (result ExpressRouteConnectionsClientGetResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.ExpressRouteConnection); err != nil {
		result = ExpressRouteConnectionsClientGetResponse{}
		return
	}
	return result, nil
}

// List - Lists ExpressRouteConnections.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - expressRouteGatewayName - The name of the ExpressRoute gateway.
//   - options - ExpressRouteConnectionsClientListOptions contains the optional parameters for the ExpressRouteConnectionsClient.List
//     method.
func (client *ExpressRouteConnectionsClient) List(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteConnectionsClientListOptions) (result ExpressRouteConnectionsClientListResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ExpressRouteConnectionsClient.List", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.listCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	result, err = client.listHandleResponse(resp)
	return
}

// listCreateRequest creates the List request.
func (client *ExpressRouteConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteConnectionsClientListOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteConnectionsClient) listHandleResponse(resp *http.Response) (result ExpressRouteConnectionsClientListResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.ExpressRouteConnectionList); err != nil {
		result = ExpressRouteConnectionsClientListResponse{}
		return
	}
	return result, nil
}
