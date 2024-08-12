// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package resources

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

// TopLevelTrackedResourcesClient contains the methods for the TopLevelTrackedResources group.
// Don't use this type directly, use NewTopLevelTrackedResourcesClient() instead.
type TopLevelTrackedResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTopLevelTrackedResourcesClient creates a new instance of TopLevelTrackedResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTopLevelTrackedResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TopLevelTrackedResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TopLevelTrackedResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrReplace - Create a TopLevelTrackedResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - topLevelTrackedResourceName - arm resource name for path
//   - resource - Resource create parameters.
//   - options - TopLevelTrackedResourcesClientBeginCreateOrReplaceOptions contains the optional parameters for the TopLevelTrackedResourcesClient.BeginCreateOrReplace
//     method.
func (client *TopLevelTrackedResourcesClient) BeginCreateOrReplace(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, resource TopLevelTrackedResource, options *TopLevelTrackedResourcesClientBeginCreateOrReplaceOptions) (*runtime.Poller[TopLevelTrackedResourcesClientCreateOrReplaceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrReplace(ctx, resourceGroupName, topLevelTrackedResourceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TopLevelTrackedResourcesClientCreateOrReplaceResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TopLevelTrackedResourcesClientCreateOrReplaceResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrReplace - Create a TopLevelTrackedResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
func (client *TopLevelTrackedResourcesClient) createOrReplace(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, resource TopLevelTrackedResource, options *TopLevelTrackedResourcesClientBeginCreateOrReplaceOptions) (*http.Response, error) {
	var err error
	const operationName = "TopLevelTrackedResourcesClient.BeginCreateOrReplace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrReplaceCreateRequest(ctx, resourceGroupName, topLevelTrackedResourceName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *TopLevelTrackedResourcesClient) createOrReplaceCreateRequest(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, resource TopLevelTrackedResource, _ *TopLevelTrackedResourcesClientBeginCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources/{topLevelTrackedResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topLevelTrackedResourceName == "" {
		return nil, errors.New("parameter topLevelTrackedResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topLevelTrackedResourceName}", url.PathEscape(topLevelTrackedResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a TopLevelTrackedResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - topLevelTrackedResourceName - arm resource name for path
//   - options - TopLevelTrackedResourcesClientBeginDeleteOptions contains the optional parameters for the TopLevelTrackedResourcesClient.BeginDelete
//     method.
func (client *TopLevelTrackedResourcesClient) BeginDelete(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, options *TopLevelTrackedResourcesClientBeginDeleteOptions) (*runtime.Poller[TopLevelTrackedResourcesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, topLevelTrackedResourceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TopLevelTrackedResourcesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TopLevelTrackedResourcesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a TopLevelTrackedResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
func (client *TopLevelTrackedResourcesClient) deleteOperation(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, options *TopLevelTrackedResourcesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "TopLevelTrackedResourcesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, topLevelTrackedResourceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TopLevelTrackedResourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, _ *TopLevelTrackedResourcesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources/{topLevelTrackedResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topLevelTrackedResourceName == "" {
		return nil, errors.New("parameter topLevelTrackedResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topLevelTrackedResourceName}", url.PathEscape(topLevelTrackedResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a TopLevelTrackedResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - topLevelTrackedResourceName - arm resource name for path
//   - options - TopLevelTrackedResourcesClientGetOptions contains the optional parameters for the TopLevelTrackedResourcesClient.Get
//     method.
func (client *TopLevelTrackedResourcesClient) Get(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, options *TopLevelTrackedResourcesClientGetOptions) (TopLevelTrackedResourcesClientGetResponse, error) {
	var err error
	const operationName = "TopLevelTrackedResourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, topLevelTrackedResourceName, options)
	if err != nil {
		return TopLevelTrackedResourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopLevelTrackedResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopLevelTrackedResourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TopLevelTrackedResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, _ *TopLevelTrackedResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources/{topLevelTrackedResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topLevelTrackedResourceName == "" {
		return nil, errors.New("parameter topLevelTrackedResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topLevelTrackedResourceName}", url.PathEscape(topLevelTrackedResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopLevelTrackedResourcesClient) getHandleResponse(resp *http.Response) (TopLevelTrackedResourcesClientGetResponse, error) {
	result := TopLevelTrackedResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopLevelTrackedResource); err != nil {
		return TopLevelTrackedResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List TopLevelTrackedResource resources by resource group
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - TopLevelTrackedResourcesClientListByResourceGroupOptions contains the optional parameters for the TopLevelTrackedResourcesClient.NewListByResourceGroupPager
//     method.
func (client *TopLevelTrackedResourcesClient) NewListByResourceGroupPager(resourceGroupName string, options *TopLevelTrackedResourcesClientListByResourceGroupOptions) *runtime.Pager[TopLevelTrackedResourcesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopLevelTrackedResourcesClientListByResourceGroupResponse]{
		More: func(page TopLevelTrackedResourcesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopLevelTrackedResourcesClientListByResourceGroupResponse) (TopLevelTrackedResourcesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopLevelTrackedResourcesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return TopLevelTrackedResourcesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *TopLevelTrackedResourcesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *TopLevelTrackedResourcesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *TopLevelTrackedResourcesClient) listByResourceGroupHandleResponse(resp *http.Response) (TopLevelTrackedResourcesClientListByResourceGroupResponse, error) {
	result := TopLevelTrackedResourcesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopLevelTrackedResourceListResult); err != nil {
		return TopLevelTrackedResourcesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List TopLevelTrackedResource resources by subscription ID
//
// Generated from API version 2023-12-01-preview
//   - options - TopLevelTrackedResourcesClientListBySubscriptionOptions contains the optional parameters for the TopLevelTrackedResourcesClient.NewListBySubscriptionPager
//     method.
func (client *TopLevelTrackedResourcesClient) NewListBySubscriptionPager(options *TopLevelTrackedResourcesClientListBySubscriptionOptions) *runtime.Pager[TopLevelTrackedResourcesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopLevelTrackedResourcesClientListBySubscriptionResponse]{
		More: func(page TopLevelTrackedResourcesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopLevelTrackedResourcesClientListBySubscriptionResponse) (TopLevelTrackedResourcesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopLevelTrackedResourcesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TopLevelTrackedResourcesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *TopLevelTrackedResourcesClient) listBySubscriptionCreateRequest(ctx context.Context, _ *TopLevelTrackedResourcesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *TopLevelTrackedResourcesClient) listBySubscriptionHandleResponse(resp *http.Response) (TopLevelTrackedResourcesClientListBySubscriptionResponse, error) {
	result := TopLevelTrackedResourcesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopLevelTrackedResourceListResult); err != nil {
		return TopLevelTrackedResourcesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a TopLevelTrackedResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - topLevelTrackedResourceName - arm resource name for path
//   - properties - The resource properties to be updated.
//   - options - TopLevelTrackedResourcesClientBeginUpdateOptions contains the optional parameters for the TopLevelTrackedResourcesClient.BeginUpdate
//     method.
func (client *TopLevelTrackedResourcesClient) BeginUpdate(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, properties TopLevelTrackedResource, options *TopLevelTrackedResourcesClientBeginUpdateOptions) (*runtime.Poller[TopLevelTrackedResourcesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, topLevelTrackedResourceName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TopLevelTrackedResourcesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TopLevelTrackedResourcesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a TopLevelTrackedResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
func (client *TopLevelTrackedResourcesClient) update(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, properties TopLevelTrackedResource, options *TopLevelTrackedResourcesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "TopLevelTrackedResourcesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, topLevelTrackedResourceName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *TopLevelTrackedResourcesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, topLevelTrackedResourceName string, properties TopLevelTrackedResource, _ *TopLevelTrackedResourcesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Azure.ResourceManager.Models.Resources/topLevelTrackedResources/{topLevelTrackedResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topLevelTrackedResourceName == "" {
		return nil, errors.New("parameter topLevelTrackedResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topLevelTrackedResourceName}", url.PathEscape(topLevelTrackedResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
