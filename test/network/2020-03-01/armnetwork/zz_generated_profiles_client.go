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

// ProfilesClient contains the methods for the NetworkProfiles group.
// Don't use this type directly, use NewProfilesClient() instead.
type ProfilesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProfilesClient creates a new instance of ProfilesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProfilesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ProfilesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates a network profile.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkProfileName - The name of the network profile.
// parameters - Parameters supplied to the create or update network profile operation.
// options - ProfilesClientCreateOrUpdateOptions contains the optional parameters for the ProfilesClient.CreateOrUpdate method.
func (client *ProfilesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkProfileName string, parameters Profile, options *ProfilesClientCreateOrUpdateOptions) (ProfilesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkProfileName, parameters, options)
	if err != nil {
		return ProfilesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProfilesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProfilesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkProfileName string, parameters Profile, options *ProfilesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkProfileName == "" {
		return nil, errors.New("parameter networkProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkProfileName}", url.PathEscape(networkProfileName))
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
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProfilesClient) createOrUpdateHandleResponse(resp *http.Response) (ProfilesClientCreateOrUpdateResponse, error) {
	result := ProfilesClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the specified network profile.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkProfileName - The name of the NetworkProfile.
// options - ProfilesClientBeginDeleteOptions contains the optional parameters for the ProfilesClient.BeginDelete method.
func (client *ProfilesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkProfileName string, options *ProfilesClientBeginDeleteOptions) (ProfilesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkProfileName, options)
	if err != nil {
		return ProfilesClientDeletePollerResponse{}, err
	}
	result := ProfilesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ProfilesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return ProfilesClientDeletePollerResponse{}, err
	}
	result.Poller = &ProfilesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified network profile.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ProfilesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkProfileName string, options *ProfilesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkProfileName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkProfileName string, options *ProfilesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkProfileName == "" {
		return nil, errors.New("parameter networkProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkProfileName}", url.PathEscape(networkProfileName))
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

// Get - Gets the specified network profile in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkProfileName - The name of the public IP prefix.
// options - ProfilesClientGetOptions contains the optional parameters for the ProfilesClient.Get method.
func (client *ProfilesClient) Get(ctx context.Context, resourceGroupName string, networkProfileName string, options *ProfilesClientGetOptions) (ProfilesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkProfileName, options)
	if err != nil {
		return ProfilesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkProfileName string, options *ProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkProfileName == "" {
		return nil, errors.New("parameter networkProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkProfileName}", url.PathEscape(networkProfileName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProfilesClient) getHandleResponse(resp *http.Response) (ProfilesClientGetResponse, error) {
	result := ProfilesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all network profiles in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - ProfilesClientListOptions contains the optional parameters for the ProfilesClient.List method.
func (client *ProfilesClient) List(resourceGroupName string, options *ProfilesClientListOptions) *ProfilesClientListPager {
	return &ProfilesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ProfilesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ProfileListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ProfilesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ProfilesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles"
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
func (client *ProfilesClient) listHandleResponse(resp *http.Response) (ProfilesClientListResponse, error) {
	result := ProfilesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileListResult); err != nil {
		return ProfilesClientListResponse{}, err
	}
	return result, nil
}

// ListAll - Gets all the network profiles in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ProfilesClientListAllOptions contains the optional parameters for the ProfilesClient.ListAll method.
func (client *ProfilesClient) ListAll(options *ProfilesClientListAllOptions) *ProfilesClientListAllPager {
	return &ProfilesClientListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProfilesClientListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ProfileListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *ProfilesClient) listAllCreateRequest(ctx context.Context, options *ProfilesClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkProfiles"
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

// listAllHandleResponse handles the ListAll response.
func (client *ProfilesClient) listAllHandleResponse(resp *http.Response) (ProfilesClientListAllResponse, error) {
	result := ProfilesClientListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileListResult); err != nil {
		return ProfilesClientListAllResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates network profile tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkProfileName - The name of the network profile.
// parameters - Parameters supplied to update network profile tags.
// options - ProfilesClientUpdateTagsOptions contains the optional parameters for the ProfilesClient.UpdateTags method.
func (client *ProfilesClient) UpdateTags(ctx context.Context, resourceGroupName string, networkProfileName string, parameters TagsObject, options *ProfilesClientUpdateTagsOptions) (ProfilesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkProfileName, parameters, options)
	if err != nil {
		return ProfilesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ProfilesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkProfileName string, parameters TagsObject, options *ProfilesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkProfileName == "" {
		return nil, errors.New("parameter networkProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkProfileName}", url.PathEscape(networkProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ProfilesClient) updateTagsHandleResponse(resp *http.Response) (ProfilesClientUpdateTagsResponse, error) {
	result := ProfilesClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
