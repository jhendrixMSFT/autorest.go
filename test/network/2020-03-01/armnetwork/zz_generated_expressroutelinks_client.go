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

// ExpressRouteLinksClient contains the methods for the ExpressRouteLinks group.
// Don't use this type directly, use NewExpressRouteLinksClient() instead.
type ExpressRouteLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteLinksClient creates a new instance of ExpressRouteLinksClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExpressRouteLinksClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &ExpressRouteLinksClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// Get - Retrieves the specified ExpressRouteLink resource.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of the ExpressRoutePort resource.
// linkName - The name of the ExpressRouteLink resource.
// options - ExpressRouteLinksGetOptions contains the optional parameters for the ExpressRouteLinks.Get method.
func (client *ExpressRouteLinksClient) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string, options *ExpressRouteLinksGetOptions) (ExpressRouteLinksGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRoutePortName, linkName, options)
	if err != nil {
		return ExpressRouteLinksGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteLinksGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteLinksGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string, options *ExpressRouteLinksGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links/{linkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	if linkName == "" {
		return nil, errors.New("parameter linkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkName}", url.PathEscape(linkName))
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
func (client *ExpressRouteLinksClient) getHandleResponse(resp *http.Response) (ExpressRouteLinksGetResponse, error) {
	result := ExpressRouteLinksGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteLink); err != nil {
		return ExpressRouteLinksGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteLinksClient) getHandleError(resp *http.Response) error {
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

// List - Retrieve the ExpressRouteLink sub-resources of the specified ExpressRoutePort resource.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of the ExpressRoutePort resource.
// options - ExpressRouteLinksListOptions contains the optional parameters for the ExpressRouteLinks.List method.
func (client *ExpressRouteLinksClient) List(resourceGroupName string, expressRoutePortName string, options *ExpressRouteLinksListOptions) *ExpressRouteLinksListPager {
	return &ExpressRouteLinksListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
		},
		advancer: func(ctx context.Context, resp ExpressRouteLinksListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteLinkListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRouteLinksListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
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
func (client *ExpressRouteLinksClient) listHandleResponse(resp *http.Response) (ExpressRouteLinksListResponse, error) {
	result := ExpressRouteLinksListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteLinkListResult); err != nil {
		return ExpressRouteLinksListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteLinksClient) listHandleError(resp *http.Response) error {
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
