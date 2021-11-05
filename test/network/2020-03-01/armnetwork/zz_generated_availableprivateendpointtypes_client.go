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

// AvailablePrivateEndpointTypesClient contains the methods for the AvailablePrivateEndpointTypes group.
// Don't use this type directly, use NewAvailablePrivateEndpointTypesClient() instead.
type AvailablePrivateEndpointTypesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailablePrivateEndpointTypesClient creates a new instance of AvailablePrivateEndpointTypesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailablePrivateEndpointTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AvailablePrivateEndpointTypesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &AvailablePrivateEndpointTypesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// List - Returns all of the resource types that can be linked to a Private Endpoint in this subscription in this region.
// If the operation fails it returns the *CloudError error type.
// location - The location of the domain name.
// options - AvailablePrivateEndpointTypesListOptions contains the optional parameters for the AvailablePrivateEndpointTypes.List
// method.
func (client *AvailablePrivateEndpointTypesClient) List(location string, options *AvailablePrivateEndpointTypesListOptions) *AvailablePrivateEndpointTypesListPager {
	return &AvailablePrivateEndpointTypesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp AvailablePrivateEndpointTypesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailablePrivateEndpointTypesResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AvailablePrivateEndpointTypesClient) listCreateRequest(ctx context.Context, location string, options *AvailablePrivateEndpointTypesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes"
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
func (client *AvailablePrivateEndpointTypesClient) listHandleResponse(resp *http.Response) (AvailablePrivateEndpointTypesListResponse, error) {
	result := AvailablePrivateEndpointTypesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailablePrivateEndpointTypesResult); err != nil {
		return AvailablePrivateEndpointTypesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AvailablePrivateEndpointTypesClient) listHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Returns all of the resource types that can be linked to a Private Endpoint in this subscription in
// this region.
// If the operation fails it returns the *CloudError error type.
// location - The location of the domain name.
// resourceGroupName - The name of the resource group.
// options - AvailablePrivateEndpointTypesListByResourceGroupOptions contains the optional parameters for the AvailablePrivateEndpointTypes.ListByResourceGroup
// method.
func (client *AvailablePrivateEndpointTypesClient) ListByResourceGroup(location string, resourceGroupName string, options *AvailablePrivateEndpointTypesListByResourceGroupOptions) *AvailablePrivateEndpointTypesListByResourceGroupPager {
	return &AvailablePrivateEndpointTypesListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, location, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp AvailablePrivateEndpointTypesListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailablePrivateEndpointTypesResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AvailablePrivateEndpointTypesClient) listByResourceGroupCreateRequest(ctx context.Context, location string, resourceGroupName string, options *AvailablePrivateEndpointTypesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AvailablePrivateEndpointTypesClient) listByResourceGroupHandleResponse(resp *http.Response) (AvailablePrivateEndpointTypesListByResourceGroupResponse, error) {
	result := AvailablePrivateEndpointTypesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailablePrivateEndpointTypesResult); err != nil {
		return AvailablePrivateEndpointTypesListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *AvailablePrivateEndpointTypesClient) listByResourceGroupHandleError(resp *http.Response) error {
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
