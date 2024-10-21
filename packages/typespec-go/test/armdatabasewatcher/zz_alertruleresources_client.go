// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdatabasewatcher

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AlertRuleResourcesClient contains the methods for the AlertRuleResources group.
// Don't use this type directly, use NewAlertRuleResourcesClient() instead.
type AlertRuleResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAlertRuleResourcesClient creates a new instance of AlertRuleResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAlertRuleResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertRuleResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AlertRuleResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a AlertRuleResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - alertRuleResourceName - The alert rule proxy resource name.
//   - resource - Resource create parameters.
//   - options - AlertRuleResourcesClientCreateOrUpdateOptions contains the optional parameters for the AlertRuleResourcesClient.CreateOrUpdate
//     method.
func (client *AlertRuleResourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, watcherName string, alertRuleResourceName string, resource AlertRuleResource, options *AlertRuleResourcesClientCreateOrUpdateOptions) (AlertRuleResourcesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "AlertRuleResourcesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, watcherName, alertRuleResourceName, resource, options)
	if err != nil {
		return AlertRuleResourcesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertRuleResourcesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return AlertRuleResourcesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AlertRuleResourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, alertRuleResourceName string, resource AlertRuleResource, _ *AlertRuleResourcesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/alertRuleResources/{alertRuleResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if alertRuleResourceName == "" {
		return nil, errors.New("parameter alertRuleResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertRuleResourceName}", url.PathEscape(alertRuleResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AlertRuleResourcesClient) createOrUpdateHandleResponse(resp *http.Response) (AlertRuleResourcesClientCreateOrUpdateResponse, error) {
	result := AlertRuleResourcesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleResource); err != nil {
		return AlertRuleResourcesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a AlertRuleResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - alertRuleResourceName - The alert rule proxy resource name.
//   - options - AlertRuleResourcesClientDeleteOptions contains the optional parameters for the AlertRuleResourcesClient.Delete
//     method.
func (client *AlertRuleResourcesClient) Delete(ctx context.Context, resourceGroupName string, watcherName string, alertRuleResourceName string, options *AlertRuleResourcesClientDeleteOptions) (AlertRuleResourcesClientDeleteResponse, error) {
	var err error
	const operationName = "AlertRuleResourcesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, watcherName, alertRuleResourceName, options)
	if err != nil {
		return AlertRuleResourcesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertRuleResourcesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AlertRuleResourcesClientDeleteResponse{}, err
	}
	return AlertRuleResourcesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AlertRuleResourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, alertRuleResourceName string, _ *AlertRuleResourcesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/alertRuleResources/{alertRuleResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if alertRuleResourceName == "" {
		return nil, errors.New("parameter alertRuleResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertRuleResourceName}", url.PathEscape(alertRuleResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a AlertRuleResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - alertRuleResourceName - The alert rule proxy resource name.
//   - options - AlertRuleResourcesClientGetOptions contains the optional parameters for the AlertRuleResourcesClient.Get method.
func (client *AlertRuleResourcesClient) Get(ctx context.Context, resourceGroupName string, watcherName string, alertRuleResourceName string, options *AlertRuleResourcesClientGetOptions) (AlertRuleResourcesClientGetResponse, error) {
	var err error
	const operationName = "AlertRuleResourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, watcherName, alertRuleResourceName, options)
	if err != nil {
		return AlertRuleResourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertRuleResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AlertRuleResourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AlertRuleResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, alertRuleResourceName string, _ *AlertRuleResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/alertRuleResources/{alertRuleResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if alertRuleResourceName == "" {
		return nil, errors.New("parameter alertRuleResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertRuleResourceName}", url.PathEscape(alertRuleResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertRuleResourcesClient) getHandleResponse(resp *http.Response) (AlertRuleResourcesClientGetResponse, error) {
	result := AlertRuleResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleResource); err != nil {
		return AlertRuleResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByParentPager - List AlertRuleResource resources by Watcher
//
// Generated from API version 2024-07-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - options - AlertRuleResourcesClientListByParentOptions contains the optional parameters for the AlertRuleResourcesClient.NewListByParentPager
//     method.
func (client *AlertRuleResourcesClient) NewListByParentPager(resourceGroupName string, watcherName string, options *AlertRuleResourcesClientListByParentOptions) *runtime.Pager[AlertRuleResourcesClientListByParentResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertRuleResourcesClientListByParentResponse]{
		More: func(page AlertRuleResourcesClientListByParentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertRuleResourcesClientListByParentResponse) (AlertRuleResourcesClientListByParentResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AlertRuleResourcesClient.NewListByParentPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByParentCreateRequest(ctx, resourceGroupName, watcherName, options)
			}, nil)
			if err != nil {
				return AlertRuleResourcesClientListByParentResponse{}, err
			}
			return client.listByParentHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByParentCreateRequest creates the ListByParent request.
func (client *AlertRuleResourcesClient) listByParentCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, _ *AlertRuleResourcesClientListByParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/alertRuleResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByParentHandleResponse handles the ListByParent response.
func (client *AlertRuleResourcesClient) listByParentHandleResponse(resp *http.Response) (AlertRuleResourcesClientListByParentResponse, error) {
	result := AlertRuleResourcesClientListByParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleResourceListResult); err != nil {
		return AlertRuleResourcesClientListByParentResponse{}, err
	}
	return result, nil
}
