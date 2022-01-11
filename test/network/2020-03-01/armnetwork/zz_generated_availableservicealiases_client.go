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

// AvailableServiceAliasesClient contains the methods for the AvailableServiceAliases group.
// Don't use this type directly, use NewAvailableServiceAliasesClient() instead.
type AvailableServiceAliasesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailableServiceAliasesClient creates a new instance of AvailableServiceAliasesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailableServiceAliasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AvailableServiceAliasesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AvailableServiceAliasesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Gets all available service aliases for this subscription in this region.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location.
// options - AvailableServiceAliasesClientListOptions contains the optional parameters for the AvailableServiceAliasesClient.List
// method.
func (client *AvailableServiceAliasesClient) List(location string, options *AvailableServiceAliasesClientListOptions) *AvailableServiceAliasesClientListPager {
	return &AvailableServiceAliasesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp AvailableServiceAliasesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailableServiceAliasesResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AvailableServiceAliasesClient) listCreateRequest(ctx context.Context, location string, options *AvailableServiceAliasesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableServiceAliases"
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
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailableServiceAliasesClient) listHandleResponse(resp *http.Response) (AvailableServiceAliasesClientListResponse, error) {
	result := AvailableServiceAliasesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableServiceAliasesResult); err != nil {
		return AvailableServiceAliasesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets all available service aliases for this resource group in this region.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// location - The location.
// options - AvailableServiceAliasesClientListByResourceGroupOptions contains the optional parameters for the AvailableServiceAliasesClient.ListByResourceGroup
// method.
func (client *AvailableServiceAliasesClient) ListByResourceGroup(resourceGroupName string, location string, options *AvailableServiceAliasesClientListByResourceGroupOptions) *AvailableServiceAliasesClientListByResourceGroupPager {
	return &AvailableServiceAliasesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, location, options)
		},
		advancer: func(ctx context.Context, resp AvailableServiceAliasesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailableServiceAliasesResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AvailableServiceAliasesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, location string, options *AvailableServiceAliasesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableServiceAliases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AvailableServiceAliasesClient) listByResourceGroupHandleResponse(resp *http.Response) (AvailableServiceAliasesClientListByResourceGroupResponse, error) {
	result := AvailableServiceAliasesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableServiceAliasesResult); err != nil {
		return AvailableServiceAliasesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}
