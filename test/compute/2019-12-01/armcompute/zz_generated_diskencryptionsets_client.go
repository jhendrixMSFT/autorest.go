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

// DiskEncryptionSetsClient contains the methods for the DiskEncryptionSets group.
// Don't use this type directly, use NewDiskEncryptionSetsClient() instead.
type DiskEncryptionSetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDiskEncryptionSetsClient creates a new instance of DiskEncryptionSetsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiskEncryptionSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DiskEncryptionSetsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &DiskEncryptionSetsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a disk encryption set
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// diskEncryptionSetName - The name of the disk encryption set that is being created. The name can't be changed after the
// disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum
// name length is 80 characters.
// diskEncryptionSet - disk encryption set object supplied in the body of the Put disk encryption set operation.
// options - DiskEncryptionSetsBeginCreateOrUpdateOptions contains the optional parameters for the DiskEncryptionSets.BeginCreateOrUpdate
// method.
func (client *DiskEncryptionSetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsBeginCreateOrUpdateOptions) (DiskEncryptionSetsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return DiskEncryptionSetsCreateOrUpdatePollerResponse{}, err
	}
	result := DiskEncryptionSetsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskEncryptionSetsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return DiskEncryptionSetsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DiskEncryptionSetsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a disk encryption set
// If the operation fails it returns the *CloudError error type.
func (client *DiskEncryptionSetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
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
func (client *DiskEncryptionSetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, diskEncryptionSet)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DiskEncryptionSetsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes a disk encryption set.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// diskEncryptionSetName - The name of the disk encryption set that is being created. The name can't be changed after the
// disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum
// name length is 80 characters.
// options - DiskEncryptionSetsBeginDeleteOptions contains the optional parameters for the DiskEncryptionSets.BeginDelete
// method.
func (client *DiskEncryptionSetsClient) BeginDelete(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsBeginDeleteOptions) (DiskEncryptionSetsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, diskEncryptionSetName, options)
	if err != nil {
		return DiskEncryptionSetsDeletePollerResponse{}, err
	}
	result := DiskEncryptionSetsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskEncryptionSetsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return DiskEncryptionSetsDeletePollerResponse{}, err
	}
	result.Poller = &DiskEncryptionSetsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a disk encryption set.
// If the operation fails it returns the *CloudError error type.
func (client *DiskEncryptionSetsClient) deleteOperation(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, options)
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
func (client *DiskEncryptionSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DiskEncryptionSetsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets information about a disk encryption set.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// diskEncryptionSetName - The name of the disk encryption set that is being created. The name can't be changed after the
// disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum
// name length is 80 characters.
// options - DiskEncryptionSetsGetOptions contains the optional parameters for the DiskEncryptionSets.Get method.
func (client *DiskEncryptionSetsClient) Get(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsGetOptions) (DiskEncryptionSetsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, options)
	if err != nil {
		return DiskEncryptionSetsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiskEncryptionSetsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiskEncryptionSetsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiskEncryptionSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiskEncryptionSetsClient) getHandleResponse(resp *http.Response) (DiskEncryptionSetsGetResponse, error) {
	result := DiskEncryptionSetsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskEncryptionSet); err != nil {
		return DiskEncryptionSetsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DiskEncryptionSetsClient) getHandleError(resp *http.Response) error {
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

// List - Lists all the disk encryption sets under a subscription.
// If the operation fails it returns the *CloudError error type.
// options - DiskEncryptionSetsListOptions contains the optional parameters for the DiskEncryptionSets.List method.
func (client *DiskEncryptionSetsClient) List(options *DiskEncryptionSetsListOptions) *DiskEncryptionSetsListPager {
	return &DiskEncryptionSetsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DiskEncryptionSetsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiskEncryptionSetList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DiskEncryptionSetsClient) listCreateRequest(ctx context.Context, options *DiskEncryptionSetsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskEncryptionSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiskEncryptionSetsClient) listHandleResponse(resp *http.Response) (DiskEncryptionSetsListResponse, error) {
	result := DiskEncryptionSetsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskEncryptionSetList); err != nil {
		return DiskEncryptionSetsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DiskEncryptionSetsClient) listHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Lists all the disk encryption sets under a resource group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// options - DiskEncryptionSetsListByResourceGroupOptions contains the optional parameters for the DiskEncryptionSets.ListByResourceGroup
// method.
func (client *DiskEncryptionSetsClient) ListByResourceGroup(resourceGroupName string, options *DiskEncryptionSetsListByResourceGroupOptions) *DiskEncryptionSetsListByResourceGroupPager {
	return &DiskEncryptionSetsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DiskEncryptionSetsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiskEncryptionSetList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DiskEncryptionSetsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DiskEncryptionSetsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DiskEncryptionSetsClient) listByResourceGroupHandleResponse(resp *http.Response) (DiskEncryptionSetsListByResourceGroupResponse, error) {
	result := DiskEncryptionSetsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskEncryptionSetList); err != nil {
		return DiskEncryptionSetsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DiskEncryptionSetsClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// BeginUpdate - Updates (patches) a disk encryption set.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// diskEncryptionSetName - The name of the disk encryption set that is being created. The name can't be changed after the
// disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum
// name length is 80 characters.
// diskEncryptionSet - disk encryption set object supplied in the body of the Patch disk encryption set operation.
// options - DiskEncryptionSetsBeginUpdateOptions contains the optional parameters for the DiskEncryptionSets.BeginUpdate
// method.
func (client *DiskEncryptionSetsClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsBeginUpdateOptions) (DiskEncryptionSetsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return DiskEncryptionSetsUpdatePollerResponse{}, err
	}
	result := DiskEncryptionSetsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskEncryptionSetsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return DiskEncryptionSetsUpdatePollerResponse{}, err
	}
	result.Poller = &DiskEncryptionSetsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates (patches) a disk encryption set.
// If the operation fails it returns the *CloudError error type.
func (client *DiskEncryptionSetsClient) update(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DiskEncryptionSetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, diskEncryptionSet)
}

// updateHandleError handles the Update error response.
func (client *DiskEncryptionSetsClient) updateHandleError(resp *http.Response) error {
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
