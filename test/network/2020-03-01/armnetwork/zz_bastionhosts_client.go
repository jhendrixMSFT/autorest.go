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

// BastionHostsClient contains the methods for the BastionHosts group.
// Don't use this type directly, use NewBastionHostsClient() instead.
type BastionHostsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBastionHostsClient creates a new instance of BastionHostsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBastionHostsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BastionHostsClient, error) {
	cl, err := arm.NewClient("armnetwork.BastionHostsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BastionHostsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - bastionHostName - The name of the Bastion Host.
//   - parameters - Parameters supplied to the create or update Bastion Host operation.
//   - options - BastionHostsClientBeginCreateOrUpdateOptions contains the optional parameters for the BastionHostsClient.BeginCreateOrUpdate
//     method.
func (client *BastionHostsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, bastionHostName string, parameters BastionHost, options *BastionHostsClientBeginCreateOrUpdateOptions) (result *runtime.Poller[BastionHostsClientCreateOrUpdateResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "BastionHostsClient.BeginCreateOrUpdate", &tracing.SpanOptions{
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
		resp, err = client.createOrUpdate(ctx, resourceGroupName, bastionHostName, parameters, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BastionHostsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken[BastionHostsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// CreateOrUpdate - Creates or updates the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *BastionHostsClient) createOrUpdate(ctx context.Context, resourceGroupName string, bastionHostName string, parameters BastionHost, options *BastionHostsClientBeginCreateOrUpdateOptions) (resp *http.Response, err error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, bastionHostName, parameters, options)
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
func (client *BastionHostsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, parameters BastionHost, options *BastionHostsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - bastionHostName - The name of the Bastion Host.
//   - options - BastionHostsClientBeginDeleteOptions contains the optional parameters for the BastionHostsClient.BeginDelete
//     method.
func (client *BastionHostsClient) BeginDelete(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientBeginDeleteOptions) (result *runtime.Poller[BastionHostsClientDeleteResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "BastionHostsClient.BeginDelete", &tracing.SpanOptions{
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
		resp, err = client.deleteOperation(ctx, resourceGroupName, bastionHostName, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BastionHostsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken[BastionHostsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// Delete - Deletes the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *BastionHostsClient) deleteOperation(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientBeginDeleteOptions) (resp *http.Response, err error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, bastionHostName, options)
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
func (client *BastionHostsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// Get - Gets the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - bastionHostName - The name of the Bastion Host.
//   - options - BastionHostsClientGetOptions contains the optional parameters for the BastionHostsClient.Get method.
func (client *BastionHostsClient) Get(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientGetOptions) (result BastionHostsClientGetResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "BastionHostsClient.Get", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getCreateRequest(ctx, resourceGroupName, bastionHostName, options)
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
func (client *BastionHostsClient) getCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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

// getHandleResponse handles the Get response.
func (client *BastionHostsClient) getHandleResponse(resp *http.Response) (result BastionHostsClientGetResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.BastionHost); err != nil {
		result = BastionHostsClientGetResponse{}
		return
	}
	return result, nil
}

// NewListPager - Lists all Bastion Hosts in a subscription.
//
// Generated from API version 2020-03-01
//   - options - BastionHostsClientListOptions contains the optional parameters for the BastionHostsClient.NewListPager method.
func (client *BastionHostsClient) NewListPager(options *BastionHostsClientListOptions) *runtime.Pager[BastionHostsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BastionHostsClientListResponse]{
		More: func(page BastionHostsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BastionHostsClientListResponse) (result BastionHostsClientListResponse, err error) {
			ctx, span := client.internal.Tracer().Start(ctx, "BastionHostsClient.NewListPager", &tracing.SpanOptions{
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
func (client *BastionHostsClient) listCreateRequest(ctx context.Context, options *BastionHostsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/bastionHosts"
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
func (client *BastionHostsClient) listHandleResponse(resp *http.Response) (result BastionHostsClientListResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.BastionHostListResult); err != nil {
		result = BastionHostsClientListResponse{}
		return
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all Bastion Hosts in a resource group.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - options - BastionHostsClientListByResourceGroupOptions contains the optional parameters for the BastionHostsClient.NewListByResourceGroupPager
//     method.
func (client *BastionHostsClient) NewListByResourceGroupPager(resourceGroupName string, options *BastionHostsClientListByResourceGroupOptions) *runtime.Pager[BastionHostsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[BastionHostsClientListByResourceGroupResponse]{
		More: func(page BastionHostsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BastionHostsClientListByResourceGroupResponse) (result BastionHostsClientListByResourceGroupResponse, err error) {
			ctx, span := client.internal.Tracer().Start(ctx, "BastionHostsClient.NewListByResourceGroupPager", &tracing.SpanOptions{
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
func (client *BastionHostsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *BastionHostsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *BastionHostsClient) listByResourceGroupHandleResponse(resp *http.Response) (result BastionHostsClientListByResourceGroupResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.BastionHostListResult); err != nil {
		result = BastionHostsClientListByResourceGroupResponse{}
		return
	}
	return result, nil
}
