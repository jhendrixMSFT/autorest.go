//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// VirtualMachineSizesClient contains the methods for the VirtualMachineSizes group.
// Don't use this type directly, use NewVirtualMachineSizesClient() instead.
type VirtualMachineSizesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineSizesClient creates a new instance of VirtualMachineSizesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineSizesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineSizesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VirtualMachineSizesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - This API is deprecated. Use Resources Skus [https://docs.microsoft.com/en-us/rest/api/compute/resourceskus/list]
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location upon which virtual-machine-sizes is queried.
// options - VirtualMachineSizesClientListOptions contains the optional parameters for the VirtualMachineSizesClient.List
// method.
func (client *VirtualMachineSizesClient) List(ctx context.Context, location string, options *VirtualMachineSizesClientListOptions) (VirtualMachineSizesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return VirtualMachineSizesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineSizesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineSizesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineSizesClient) listCreateRequest(ctx context.Context, location string, options *VirtualMachineSizesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/vmSizes"
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
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineSizesClient) listHandleResponse(resp *http.Response) (VirtualMachineSizesClientListResponse, error) {
	result := VirtualMachineSizesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineSizeListResult); err != nil {
		return VirtualMachineSizesClientListResponse{}, err
	}
	return result, nil
}
