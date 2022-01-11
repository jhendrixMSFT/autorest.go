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

// AvailableResourceGroupDelegationsClient contains the methods for the AvailableResourceGroupDelegations group.
// Don't use this type directly, use NewAvailableResourceGroupDelegationsClient() instead.
type AvailableResourceGroupDelegationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailableResourceGroupDelegationsClient creates a new instance of AvailableResourceGroupDelegationsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailableResourceGroupDelegationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AvailableResourceGroupDelegationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AvailableResourceGroupDelegationsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Gets all of the available subnet delegations for this resource group in this region.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location of the domain name.
// resourceGroupName - The name of the resource group.
// options - AvailableResourceGroupDelegationsClientListOptions contains the optional parameters for the AvailableResourceGroupDelegationsClient.List
// method.
func (client *AvailableResourceGroupDelegationsClient) List(location string, resourceGroupName string, options *AvailableResourceGroupDelegationsClientListOptions) *AvailableResourceGroupDelegationsClientListPager {
	return &AvailableResourceGroupDelegationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp AvailableResourceGroupDelegationsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailableDelegationsResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AvailableResourceGroupDelegationsClient) listCreateRequest(ctx context.Context, location string, resourceGroupName string, options *AvailableResourceGroupDelegationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableDelegations"
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

// listHandleResponse handles the List response.
func (client *AvailableResourceGroupDelegationsClient) listHandleResponse(resp *http.Response) (AvailableResourceGroupDelegationsClientListResponse, error) {
	result := AvailableResourceGroupDelegationsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableDelegationsResult); err != nil {
		return AvailableResourceGroupDelegationsClientListResponse{}, err
	}
	return result, nil
}
