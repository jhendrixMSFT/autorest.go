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
	"net/http"
	"net/url"
	"strings"
)

// ServiceEndpointPoliciesClient contains the methods for the ServiceEndpointPolicies group.
// Don't use this type directly, use NewServiceEndpointPoliciesClient() instead.
type ServiceEndpointPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServiceEndpointPoliciesClient creates a new instance of ServiceEndpointPoliciesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceEndpointPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceEndpointPoliciesClient, error) {
	cl, err := arm.NewClient("armnetwork.ServiceEndpointPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceEndpointPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a service Endpoint Policies.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - serviceEndpointPolicyName - The name of the service endpoint policy.
//   - parameters - Parameters supplied to the create or update service endpoint policy operation.
//   - options - ServiceEndpointPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServiceEndpointPoliciesClient.BeginCreateOrUpdate
//     method.
func (client *ServiceEndpointPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServiceEndpointPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceEndpointPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServiceEndpointPoliciesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServiceEndpointPoliciesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a service Endpoint Policies.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *ServiceEndpointPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceEndpointPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
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

// BeginDelete - Deletes the specified service endpoint policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - serviceEndpointPolicyName - The name of the service endpoint policy.
//   - options - ServiceEndpointPoliciesClientBeginDeleteOptions contains the optional parameters for the ServiceEndpointPoliciesClient.BeginDelete
//     method.
func (client *ServiceEndpointPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesClientBeginDeleteOptions) (*runtime.Poller[ServiceEndpointPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceEndpointPolicyName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServiceEndpointPoliciesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServiceEndpointPoliciesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified service endpoint policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *ServiceEndpointPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceEndpointPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
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

// Get - Gets the specified service Endpoint Policies in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - serviceEndpointPolicyName - The name of the service endpoint policy.
//   - options - ServiceEndpointPoliciesClientGetOptions contains the optional parameters for the ServiceEndpointPoliciesClient.Get
//     method.
func (client *ServiceEndpointPoliciesClient) Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesClientGetOptions) (ServiceEndpointPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, options)
	if err != nil {
		return ServiceEndpointPoliciesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceEndpointPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceEndpointPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceEndpointPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
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
func (client *ServiceEndpointPoliciesClient) getHandleResponse(resp *http.Response) (ServiceEndpointPoliciesClientGetResponse, error) {
	result := ServiceEndpointPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceEndpointPolicy); err != nil {
		return ServiceEndpointPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the service endpoint policies in a subscription.
//
// Generated from API version 2020-03-01
//   - options - ServiceEndpointPoliciesClientListOptions contains the optional parameters for the ServiceEndpointPoliciesClient.NewListPager
//     method.
func (client *ServiceEndpointPoliciesClient) NewListPager(options *ServiceEndpointPoliciesClientListOptions) *runtime.Pager[ServiceEndpointPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceEndpointPoliciesClientListResponse]{
		More: func(page ServiceEndpointPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceEndpointPoliciesClientListResponse) (ServiceEndpointPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceEndpointPoliciesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceEndpointPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceEndpointPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServiceEndpointPoliciesClient) listCreateRequest(ctx context.Context, options *ServiceEndpointPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ServiceEndpointPolicies"
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
func (client *ServiceEndpointPoliciesClient) listHandleResponse(resp *http.Response) (ServiceEndpointPoliciesClientListResponse, error) {
	result := ServiceEndpointPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceEndpointPolicyListResult); err != nil {
		return ServiceEndpointPoliciesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all service endpoint Policies in a resource group.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - options - ServiceEndpointPoliciesClientListByResourceGroupOptions contains the optional parameters for the ServiceEndpointPoliciesClient.NewListByResourceGroupPager
//     method.
func (client *ServiceEndpointPoliciesClient) NewListByResourceGroupPager(resourceGroupName string, options *ServiceEndpointPoliciesClientListByResourceGroupOptions) *runtime.Pager[ServiceEndpointPoliciesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceEndpointPoliciesClientListByResourceGroupResponse]{
		More: func(page ServiceEndpointPoliciesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceEndpointPoliciesClientListByResourceGroupResponse) (ServiceEndpointPoliciesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceEndpointPoliciesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceEndpointPoliciesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceEndpointPoliciesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServiceEndpointPoliciesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServiceEndpointPoliciesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies"
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
func (client *ServiceEndpointPoliciesClient) listByResourceGroupHandleResponse(resp *http.Response) (ServiceEndpointPoliciesClientListByResourceGroupResponse, error) {
	result := ServiceEndpointPoliciesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceEndpointPolicyListResult); err != nil {
		return ServiceEndpointPoliciesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates tags of a service endpoint policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - serviceEndpointPolicyName - The name of the service endpoint policy.
//   - parameters - Parameters supplied to update service endpoint policy tags.
//   - options - ServiceEndpointPoliciesClientUpdateTagsOptions contains the optional parameters for the ServiceEndpointPoliciesClient.UpdateTags
//     method.
func (client *ServiceEndpointPoliciesClient) UpdateTags(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters TagsObject, options *ServiceEndpointPoliciesClientUpdateTagsOptions) (ServiceEndpointPoliciesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, parameters, options)
	if err != nil {
		return ServiceEndpointPoliciesClientUpdateTagsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceEndpointPoliciesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceEndpointPoliciesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ServiceEndpointPoliciesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters TagsObject, options *ServiceEndpointPoliciesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ServiceEndpointPoliciesClient) updateTagsHandleResponse(resp *http.Response) (ServiceEndpointPoliciesClientUpdateTagsResponse, error) {
	result := ServiceEndpointPoliciesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceEndpointPolicy); err != nil {
		return ServiceEndpointPoliciesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
