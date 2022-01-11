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

// SecurityGroupsClient contains the methods for the NetworkSecurityGroups group.
// Don't use this type directly, use NewSecurityGroupsClient() instead.
type SecurityGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSecurityGroupsClient creates a new instance of SecurityGroupsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSecurityGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SecurityGroupsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SecurityGroupsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a network security group in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// parameters - Parameters supplied to the create or update network security group operation.
// options - SecurityGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the SecurityGroupsClient.BeginCreateOrUpdate
// method.
func (client *SecurityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *SecurityGroupsClientBeginCreateOrUpdateOptions) (SecurityGroupsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return SecurityGroupsClientCreateOrUpdatePollerResponse{}, err
	}
	result := SecurityGroupsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SecurityGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return SecurityGroupsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SecurityGroupsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a network security group in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SecurityGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *SecurityGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *SecurityGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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

// BeginDelete - Deletes the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// options - SecurityGroupsClientBeginDeleteOptions contains the optional parameters for the SecurityGroupsClient.BeginDelete
// method.
func (client *SecurityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientBeginDeleteOptions) (SecurityGroupsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return SecurityGroupsClientDeletePollerResponse{}, err
	}
	result := SecurityGroupsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SecurityGroupsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return SecurityGroupsClientDeletePollerResponse{}, err
	}
	result.Poller = &SecurityGroupsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SecurityGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
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
func (client *SecurityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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

// Get - Gets the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// options - SecurityGroupsClientGetOptions contains the optional parameters for the SecurityGroupsClient.Get method.
func (client *SecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientGetOptions) (SecurityGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return SecurityGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecurityGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecurityGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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
func (client *SecurityGroupsClient) getHandleResponse(resp *http.Response) (SecurityGroupsClientGetResponse, error) {
	result := SecurityGroupsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroup); err != nil {
		return SecurityGroupsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all network security groups in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - SecurityGroupsClientListOptions contains the optional parameters for the SecurityGroupsClient.List method.
func (client *SecurityGroupsClient) List(resourceGroupName string, options *SecurityGroupsClientListOptions) *SecurityGroupsClientListPager {
	return &SecurityGroupsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SecurityGroupsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecurityGroupListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SecurityGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *SecurityGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups"
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
func (client *SecurityGroupsClient) listHandleResponse(resp *http.Response) (SecurityGroupsClientListResponse, error) {
	result := SecurityGroupsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroupListResult); err != nil {
		return SecurityGroupsClientListResponse{}, err
	}
	return result, nil
}

// ListAll - Gets all network security groups in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - SecurityGroupsClientListAllOptions contains the optional parameters for the SecurityGroupsClient.ListAll method.
func (client *SecurityGroupsClient) ListAll(options *SecurityGroupsClientListAllOptions) *SecurityGroupsClientListAllPager {
	return &SecurityGroupsClientListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SecurityGroupsClientListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecurityGroupListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *SecurityGroupsClient) listAllCreateRequest(ctx context.Context, options *SecurityGroupsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups"
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
func (client *SecurityGroupsClient) listAllHandleResponse(resp *http.Response) (SecurityGroupsClientListAllResponse, error) {
	result := SecurityGroupsClientListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroupListResult); err != nil {
		return SecurityGroupsClientListAllResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a network security group tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// parameters - Parameters supplied to update network security group tags.
// options - SecurityGroupsClientUpdateTagsOptions contains the optional parameters for the SecurityGroupsClient.UpdateTags
// method.
func (client *SecurityGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *SecurityGroupsClientUpdateTagsOptions) (SecurityGroupsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return SecurityGroupsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecurityGroupsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecurityGroupsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SecurityGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *SecurityGroupsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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
func (client *SecurityGroupsClient) updateTagsHandleResponse(resp *http.Response) (SecurityGroupsClientUpdateTagsResponse, error) {
	result := SecurityGroupsClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroup); err != nil {
		return SecurityGroupsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
