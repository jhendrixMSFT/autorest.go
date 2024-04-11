// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter

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

// ServicesClient contains the methods for the Microsoft.ApiCenter namespace.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	internal *arm.Client
}

// NewServicesClient creates a new instance of [ServicesClient] with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServicesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	servicesClient := &ServicesClient{
		internal: cl,
	}
	return servicesClient, nil
}

// CreateOrUpdate - Creates new or updates existing API.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - resource - Resource create parameters.
//   - options - ServicesClientCreateOrUpdateOptions contains the optional parameters for the ServicesClient.CreateOrUpdate method.
func (client *ServicesClient) CreateOrUpdate(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, resource Service, options *ServicesClientCreateOrUpdateOptions) (ServicesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ServicesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, resource, options)
	if err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServicesClient) createOrUpdateCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, resource Service, options *ServicesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServicesClient) createOrUpdateHandleResponse(resp *http.Response) (ServicesClientCreateOrUpdateResponse, error) {
	result := ServicesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specified service.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - options - ServicesClientDeleteOptions contains the optional parameters for the ServicesClient.Delete method.
func (client *ServicesClient) Delete(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, options *ServicesClientDeleteOptions) (ServicesClientDeleteResponse, error) {
	var err error
	const operationName = "ServicesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, options)
	if err != nil {
		return ServicesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientDeleteResponse{}, err
	}
	return ServicesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, options *ServicesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExportMetadataSchema - Exports the effective metadata schema.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - payload - The content of the action request
//   - options - ServicesClientExportMetadataSchemaOptions contains the optional parameters for the ServicesClient.ExportMetadataSchema
//     method.
func (client *ServicesClient) BeginExportMetadataSchema(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, payload MetadataSchemaExportRequest, options *ServicesClientExportMetadataSchemaOptions) (*runtime.Poller[ServicesClientExportMetadataSchemaResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportMetadataSchema(ctx, subscriptionID, resourceGroupName, serviceName, payload, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServicesClientExportMetadataSchemaResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServicesClientExportMetadataSchemaResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ExportMetadataSchema - Exports the effective metadata schema.
func (client *ServicesClient) exportMetadataSchema(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, payload MetadataSchemaExportRequest, options *ServicesClientExportMetadataSchemaOptions) (*http.Response, error) {
	var err error
	const operationName = "ServicesClient.BeginExportMetadataSchema"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportMetadataSchemaCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, payload, options)
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

// exportMetadataSchemaCreateRequest creates the ExportMetadataSchema request.
func (client *ServicesClient) exportMetadataSchemaCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, payload MetadataSchemaExportRequest, options *ServicesClientExportMetadataSchemaOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/exportMetadataSchema"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, payload); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Returns details of the service.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	var err error
	const operationName = "ServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, options *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Returns a collection of services within the resource group.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ServicesClientListByResourceGroupOptions contains the optional parameters for the ServicesClient.NewListByResourceGroupPager
//     method.
func (client *ServicesClient) NewListByResourceGroupPager(subscriptionID string, resourceGroupName string, options *ServicesClientListByResourceGroupOptions) *runtime.Pager[ServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListByResourceGroupResponse]{
		More: func(page ServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListByResourceGroupResponse) (ServicesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServicesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, subscriptionID, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ServicesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServicesClient) listByResourceGroupCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, options *ServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (ServicesClientListByResourceGroupResponse, error) {
	result := ServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceListResult); err != nil {
		return ServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists services within an Azure subscription.
//   - subscriptionID - The ID of the target subscription.
//   - options - ServicesClientListBySubscriptionOptions contains the optional parameters for the ServicesClient.NewListBySubscriptionPager
//     method.
func (client *ServicesClient) NewListBySubscriptionPager(subscriptionID string, options *ServicesClientListBySubscriptionOptions) *runtime.Pager[ServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListBySubscriptionResponse]{
		More: func(page ServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListBySubscriptionResponse) (ServicesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServicesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, subscriptionID, options)
			}, nil)
			if err != nil {
				return ServicesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ServicesClient) listBySubscriptionCreateRequest(ctx context.Context, subscriptionID string, options *ServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ApiCenter/services"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (ServicesClientListBySubscriptionResponse, error) {
	result := ServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceListResult); err != nil {
		return ServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates existing service.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - payload - The resource properties to be updated.
//   - options - ServicesClientUpdateOptions contains the optional parameters for the ServicesClient.Update method.
func (client *ServicesClient) Update(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, payload ServiceUpdate, options *ServicesClientUpdateOptions) (ServicesClientUpdateResponse, error) {
	var err error
	const operationName = "ServicesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, payload, options)
	if err != nil {
		return ServicesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ServicesClient) updateCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, payload ServiceUpdate, options *ServicesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, payload); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ServicesClient) updateHandleResponse(resp *http.Response) (ServicesClientUpdateResponse, error) {
	result := ServicesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientUpdateResponse{}, err
	}
	return result, nil
}
