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

// InboundNatRulesClient contains the methods for the InboundNatRules group.
// Don't use this type directly, use NewInboundNatRulesClient() instead.
type InboundNatRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInboundNatRulesClient creates a new instance of InboundNatRulesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInboundNatRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InboundNatRulesClient, error) {
	cl, err := arm.NewClient("armnetwork.InboundNatRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InboundNatRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - inboundNatRuleName - The name of the inbound nat rule.
//   - inboundNatRuleParameters - Parameters supplied to the create or update inbound nat rule operation.
//   - options - InboundNatRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the InboundNatRulesClient.BeginCreateOrUpdate
//     method.
func (client *InboundNatRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, options *InboundNatRulesClientBeginCreateOrUpdateOptions) (result *runtime.Poller[InboundNatRulesClientCreateOrUpdateResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "InboundNatRulesClient.BeginCreateOrUpdate", &tracing.SpanOptions{
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
		resp, err = client.createOrUpdate(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, inboundNatRuleParameters, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InboundNatRulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken[InboundNatRulesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// CreateOrUpdate - Creates or updates a load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *InboundNatRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, options *InboundNatRulesClientBeginCreateOrUpdateOptions) (resp *http.Response, err error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, inboundNatRuleParameters, options)
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
func (client *InboundNatRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, options *InboundNatRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if inboundNatRuleName == "" {
		return nil, errors.New("parameter inboundNatRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inboundNatRuleName}", url.PathEscape(inboundNatRuleName))
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
	return req, runtime.MarshalAsJSON(req, inboundNatRuleParameters)
}

// BeginDelete - Deletes the specified load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - inboundNatRuleName - The name of the inbound nat rule.
//   - options - InboundNatRulesClientBeginDeleteOptions contains the optional parameters for the InboundNatRulesClient.BeginDelete
//     method.
func (client *InboundNatRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientBeginDeleteOptions) (result *runtime.Poller[InboundNatRulesClientDeleteResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "InboundNatRulesClient.BeginDelete", &tracing.SpanOptions{
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
		resp, err = client.deleteOperation(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InboundNatRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken[InboundNatRulesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// Delete - Deletes the specified load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *InboundNatRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientBeginDeleteOptions) (resp *http.Response, err error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, options)
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
func (client *InboundNatRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if inboundNatRuleName == "" {
		return nil, errors.New("parameter inboundNatRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inboundNatRuleName}", url.PathEscape(inboundNatRuleName))
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

// Get - Gets the specified load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - inboundNatRuleName - The name of the inbound nat rule.
//   - options - InboundNatRulesClientGetOptions contains the optional parameters for the InboundNatRulesClient.Get method.
func (client *InboundNatRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientGetOptions) (result InboundNatRulesClientGetResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "InboundNatRulesClient.Get", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, options)
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
func (client *InboundNatRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if inboundNatRuleName == "" {
		return nil, errors.New("parameter inboundNatRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inboundNatRuleName}", url.PathEscape(inboundNatRuleName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InboundNatRulesClient) getHandleResponse(resp *http.Response) (result InboundNatRulesClientGetResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.InboundNatRule); err != nil {
		result = InboundNatRulesClientGetResponse{}
		return
	}
	return result, nil
}

// NewListPager - Gets all the inbound nat rules in a load balancer.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - options - InboundNatRulesClientListOptions contains the optional parameters for the InboundNatRulesClient.NewListPager
//     method.
func (client *InboundNatRulesClient) NewListPager(resourceGroupName string, loadBalancerName string, options *InboundNatRulesClientListOptions) *runtime.Pager[InboundNatRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[InboundNatRulesClientListResponse]{
		More: func(page InboundNatRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InboundNatRulesClientListResponse) (result InboundNatRulesClientListResponse, err error) {
			ctx, span := client.internal.Tracer().Start(ctx, "InboundNatRulesClient.NewListPager", &tracing.SpanOptions{
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
				req, err = client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
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
func (client *InboundNatRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *InboundNatRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
func (client *InboundNatRulesClient) listHandleResponse(resp *http.Response) (result InboundNatRulesClientListResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.InboundNatRuleListResult); err != nil {
		result = InboundNatRulesClientListResponse{}
		return
	}
	return result, nil
}
