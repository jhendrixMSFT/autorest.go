//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"net/http"
	"net/url"
	"strings"
)

// VirtualHubsClient contains the methods for the VirtualHubs group.
// Don't use this type directly, use NewVirtualHubsClient() instead.
type VirtualHubsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualHubsClient creates a new instance of VirtualHubsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualHubsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualHubsClient, error) {
	cl, err := arm.NewClient("armnetwork.VirtualHubsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualHubsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a VirtualHub resource if it doesn't exist else updates the existing VirtualHub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - virtualHubParameters - Parameters supplied to create or update VirtualHub.
//   - options - VirtualHubsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualHubsClient.BeginCreateOrUpdate
//     method.
func (client *VirtualHubsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters VirtualHub, options *VirtualHubsClientBeginCreateOrUpdateOptions) (result *runtime.Poller[VirtualHubsClientCreateOrUpdateResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "VirtualHubsClient.BeginCreateOrUpdate", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	if options == nil || options.ResumeToken == "" {
		var resp *http.Response
		resp, err = client.createOrUpdate(ctx, resourceGroupName, virtualHubName, virtualHubParameters, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualHubsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken[VirtualHubsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// CreateOrUpdate - Creates a VirtualHub resource if it doesn't exist else updates the existing VirtualHub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *VirtualHubsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters VirtualHub, options *VirtualHubsClientBeginCreateOrUpdateOptions) (resp *http.Response, err error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, virtualHubParameters, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualHubsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters VirtualHub, options *VirtualHubsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, virtualHubParameters)
}

// BeginDelete - Deletes a VirtualHub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - options - VirtualHubsClientBeginDeleteOptions contains the optional parameters for the VirtualHubsClient.BeginDelete method.
func (client *VirtualHubsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsClientBeginDeleteOptions) (result *runtime.Poller[VirtualHubsClientDeleteResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "VirtualHubsClient.BeginDelete", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	if options == nil || options.ResumeToken == "" {
		var resp *http.Response
		resp, err = client.deleteOperation(ctx, resourceGroupName, virtualHubName, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualHubsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken[VirtualHubsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// Delete - Deletes a VirtualHub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *VirtualHubsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsClientBeginDeleteOptions) (resp *http.Response, err error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualHubsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a VirtualHub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - options - VirtualHubsClientGetOptions contains the optional parameters for the VirtualHubsClient.Get method.
func (client *VirtualHubsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsClientGetOptions) (result VirtualHubsClientGetResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "VirtualHubsClient.Get", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	result, err = client.getHandleResponse(resp)
	return
}

// getCreateRequest creates the Get request.
func (client *VirtualHubsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualHubsClient) getHandleResponse(resp *http.Response) (result VirtualHubsClientGetResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.VirtualHub); err != nil {
		result = VirtualHubsClientGetResponse{}
		return
	}
	return result, nil
}

// NewListPager - Lists all the VirtualHubs in a subscription.
//
// Generated from API version 2020-03-01
//   - options - VirtualHubsClientListOptions contains the optional parameters for the VirtualHubsClient.NewListPager method.
func (client *VirtualHubsClient) NewListPager(options *VirtualHubsClientListOptions) *runtime.Pager[VirtualHubsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualHubsClientListResponse]{
		More: func(page VirtualHubsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualHubsClientListResponse) (result VirtualHubsClientListResponse, err error) {
			ctx, span := client.internal.Tracer().Start(ctx, "VirtualHubsClient.NewListPager", &tracing.SpanOptions{
				Kind: tracing.SpanKindInternal,
			})
			defer func() {
				if err != nil {
					span.AddError(err)
				}
				span.End()
			}()
			var req *policy.Request
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				err = runtime.NewResponseError(resp)
				return
			}
			result, err = client.listHandleResponse(resp)
			return
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualHubsClient) listCreateRequest(ctx context.Context, options *VirtualHubsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualHubs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualHubsClient) listHandleResponse(resp *http.Response) (result VirtualHubsClientListResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.ListVirtualHubsResult); err != nil {
		result = VirtualHubsClientListResponse{}
		return
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the VirtualHubs in a resource group.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - options - VirtualHubsClientListByResourceGroupOptions contains the optional parameters for the VirtualHubsClient.NewListByResourceGroupPager
//     method.
func (client *VirtualHubsClient) NewListByResourceGroupPager(resourceGroupName string, options *VirtualHubsClientListByResourceGroupOptions) *runtime.Pager[VirtualHubsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualHubsClientListByResourceGroupResponse]{
		More: func(page VirtualHubsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualHubsClientListByResourceGroupResponse) (result VirtualHubsClientListByResourceGroupResponse, err error) {
			ctx, span := client.internal.Tracer().Start(ctx, "VirtualHubsClient.NewListByResourceGroupPager", &tracing.SpanOptions{
				Kind: tracing.SpanKindInternal,
			})
			defer func() {
				if err != nil {
					span.AddError(err)
				}
				span.End()
			}()
			var req *policy.Request
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				err = runtime.NewResponseError(resp)
				return
			}
			result, err = client.listByResourceGroupHandleResponse(resp)
			return
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualHubsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualHubsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs"
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
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualHubsClient) listByResourceGroupHandleResponse(resp *http.Response) (result VirtualHubsClientListByResourceGroupResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.ListVirtualHubsResult); err != nil {
		result = VirtualHubsClientListByResourceGroupResponse{}
		return
	}
	return result, nil
}

// UpdateTags - Updates VirtualHub tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - virtualHubParameters - Parameters supplied to update VirtualHub tags.
//   - options - VirtualHubsClientUpdateTagsOptions contains the optional parameters for the VirtualHubsClient.UpdateTags method.
func (client *VirtualHubsClient) UpdateTags(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters TagsObject, options *VirtualHubsClientUpdateTagsOptions) (result VirtualHubsClientUpdateTagsResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "VirtualHubsClient.UpdateTags", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, virtualHubName, virtualHubParameters, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	result, err = client.updateTagsHandleResponse(resp)
	return
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualHubsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters TagsObject, options *VirtualHubsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, virtualHubParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualHubsClient) updateTagsHandleResponse(resp *http.Response) (result VirtualHubsClientUpdateTagsResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.VirtualHub); err != nil {
		result = VirtualHubsClientUpdateTagsResponse{}
		return
	}
	return result, nil
}
