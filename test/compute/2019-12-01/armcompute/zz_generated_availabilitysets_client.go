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

// AvailabilitySetsClient contains the methods for the AvailabilitySets group.
// Don't use this type directly, use NewAvailabilitySetsClient() instead.
type AvailabilitySetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailabilitySetsClient creates a new instance of AvailabilitySetsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailabilitySetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AvailabilitySetsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AvailabilitySetsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or update an availability set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// availabilitySetName - The name of the availability set.
// parameters - Parameters supplied to the Create Availability Set operation.
// options - AvailabilitySetsClientCreateOrUpdateOptions contains the optional parameters for the AvailabilitySetsClient.CreateOrUpdate
// method.
func (client *AvailabilitySetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySet, options *AvailabilitySetsClientCreateOrUpdateOptions) (AvailabilitySetsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, availabilitySetName, parameters, options)
	if err != nil {
		return AvailabilitySetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AvailabilitySetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AvailabilitySetsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AvailabilitySetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySet, options *AvailabilitySetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AvailabilitySetsClient) createOrUpdateHandleResponse(resp *http.Response) (AvailabilitySetsClientCreateOrUpdateResponse, error) {
	result := AvailabilitySetsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySet); err != nil {
		return AvailabilitySetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an availability set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// availabilitySetName - The name of the availability set.
// options - AvailabilitySetsClientDeleteOptions contains the optional parameters for the AvailabilitySetsClient.Delete method.
func (client *AvailabilitySetsClient) Delete(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientDeleteOptions) (AvailabilitySetsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, availabilitySetName, options)
	if err != nil {
		return AvailabilitySetsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AvailabilitySetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AvailabilitySetsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AvailabilitySetsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AvailabilitySetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Retrieves information about an availability set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// availabilitySetName - The name of the availability set.
// options - AvailabilitySetsClientGetOptions contains the optional parameters for the AvailabilitySetsClient.Get method.
func (client *AvailabilitySetsClient) Get(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientGetOptions) (AvailabilitySetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, availabilitySetName, options)
	if err != nil {
		return AvailabilitySetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AvailabilitySetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AvailabilitySetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AvailabilitySetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
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

// getHandleResponse handles the Get response.
func (client *AvailabilitySetsClient) getHandleResponse(resp *http.Response) (AvailabilitySetsClientGetResponse, error) {
	result := AvailabilitySetsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySet); err != nil {
		return AvailabilitySetsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all availability sets in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - AvailabilitySetsClientListOptions contains the optional parameters for the AvailabilitySetsClient.List method.
func (client *AvailabilitySetsClient) List(resourceGroupName string, options *AvailabilitySetsClientListOptions) *AvailabilitySetsClientListPager {
	return &AvailabilitySetsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp AvailabilitySetsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailabilitySetListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AvailabilitySetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *AvailabilitySetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets"
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
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailabilitySetsClient) listHandleResponse(resp *http.Response) (AvailabilitySetsClientListResponse, error) {
	result := AvailabilitySetsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySetListResult); err != nil {
		return AvailabilitySetsClientListResponse{}, err
	}
	return result, nil
}

// ListAvailableSizes - Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing
// availability set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// availabilitySetName - The name of the availability set.
// options - AvailabilitySetsClientListAvailableSizesOptions contains the optional parameters for the AvailabilitySetsClient.ListAvailableSizes
// method.
func (client *AvailabilitySetsClient) ListAvailableSizes(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientListAvailableSizesOptions) (AvailabilitySetsClientListAvailableSizesResponse, error) {
	req, err := client.listAvailableSizesCreateRequest(ctx, resourceGroupName, availabilitySetName, options)
	if err != nil {
		return AvailabilitySetsClientListAvailableSizesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AvailabilitySetsClientListAvailableSizesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AvailabilitySetsClientListAvailableSizesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAvailableSizesHandleResponse(resp)
}

// listAvailableSizesCreateRequest creates the ListAvailableSizes request.
func (client *AvailabilitySetsClient) listAvailableSizesCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientListAvailableSizesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
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

// listAvailableSizesHandleResponse handles the ListAvailableSizes response.
func (client *AvailabilitySetsClient) listAvailableSizesHandleResponse(resp *http.Response) (AvailabilitySetsClientListAvailableSizesResponse, error) {
	result := AvailabilitySetsClientListAvailableSizesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineSizeListResult); err != nil {
		return AvailabilitySetsClientListAvailableSizesResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Lists all availability sets in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AvailabilitySetsClientListBySubscriptionOptions contains the optional parameters for the AvailabilitySetsClient.ListBySubscription
// method.
func (client *AvailabilitySetsClient) ListBySubscription(options *AvailabilitySetsClientListBySubscriptionOptions) *AvailabilitySetsClientListBySubscriptionPager {
	return &AvailabilitySetsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AvailabilitySetsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailabilitySetListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AvailabilitySetsClient) listBySubscriptionCreateRequest(ctx context.Context, options *AvailabilitySetsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/availabilitySets"
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AvailabilitySetsClient) listBySubscriptionHandleResponse(resp *http.Response) (AvailabilitySetsClientListBySubscriptionResponse, error) {
	result := AvailabilitySetsClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySetListResult); err != nil {
		return AvailabilitySetsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an availability set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// availabilitySetName - The name of the availability set.
// parameters - Parameters supplied to the Update Availability Set operation.
// options - AvailabilitySetsClientUpdateOptions contains the optional parameters for the AvailabilitySetsClient.Update method.
func (client *AvailabilitySetsClient) Update(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate, options *AvailabilitySetsClientUpdateOptions) (AvailabilitySetsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, availabilitySetName, parameters, options)
	if err != nil {
		return AvailabilitySetsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AvailabilitySetsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AvailabilitySetsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AvailabilitySetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate, options *AvailabilitySetsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *AvailabilitySetsClient) updateHandleResponse(resp *http.Response) (AvailabilitySetsClientUpdateResponse, error) {
	result := AvailabilitySetsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySet); err != nil {
		return AvailabilitySetsClientUpdateResponse{}, err
	}
	return result, nil
}
