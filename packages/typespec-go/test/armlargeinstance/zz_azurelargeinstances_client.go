// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armlargeinstance

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

// AzureLargeInstancesClient contains the methods for the Microsoft.AzureLargeInstance namespace.
// Don't use this type directly, use NewAzureLargeInstancesClient() instead.
type AzureLargeInstancesClient struct {
	internal *arm.Client
}

// NewAzureLargeInstancesClient creates a new instance of [AzureLargeInstancesClient] with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureLargeInstancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureLargeInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureLargeInstancesClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets an Azure Large Instance for the specified subscription, resource group,
// and instance name.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeInstanceName - Name of the AzureLargeInstance.
//   - options - AzureLargeInstancesClientGetOptions contains the optional parameters for the AzureLargeInstancesClient.Get method.
func (client *AzureLargeInstancesClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientGetOptions) (AzureLargeInstancesClientGetResponse, error) {
	var err error
	const operationName = "AzureLargeInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, azureLargeInstanceName, options)
	if err != nil {
		return AzureLargeInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureLargeInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureLargeInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AzureLargeInstancesClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeInstances/{azureLargeInstanceName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeInstanceName == "" {
		return nil, errors.New("parameter azureLargeInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeInstanceName}", url.PathEscape(azureLargeInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureLargeInstancesClient) getHandleResponse(resp *http.Response) (AzureLargeInstancesClientGetResponse, error) {
	result := AzureLargeInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeInstance); err != nil {
		return AzureLargeInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of Azure Large Instances in the specified subscription and resource
// group. The operations returns various properties of each Azure Large Instance.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - AzureLargeInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureLargeInstancesClient.NewListByResourceGroupPager
//     method.
func (client *AzureLargeInstancesClient) NewListByResourceGroupPager(subscriptionID string, resourceGroupName string, options *AzureLargeInstancesClientListByResourceGroupOptions) *runtime.Pager[AzureLargeInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureLargeInstancesClientListByResourceGroupResponse]{
		More: func(page AzureLargeInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureLargeInstancesClientListByResourceGroupResponse) (AzureLargeInstancesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureLargeInstancesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, subscriptionID, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AzureLargeInstancesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AzureLargeInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, options *AzureLargeInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeInstances"
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
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AzureLargeInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (AzureLargeInstancesClientListByResourceGroupResponse, error) {
	result := AzureLargeInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return AzureLargeInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of Azure Large Instances in the specified subscription. The
// operations returns various properties of each Azure Large Instance.
//   - subscriptionID - The ID of the target subscription.
//   - options - AzureLargeInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureLargeInstancesClient.NewListBySubscriptionPager
//     method.
func (client *AzureLargeInstancesClient) NewListBySubscriptionPager(subscriptionID string, options *AzureLargeInstancesClientListBySubscriptionOptions) *runtime.Pager[AzureLargeInstancesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureLargeInstancesClientListBySubscriptionResponse]{
		More: func(page AzureLargeInstancesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureLargeInstancesClientListBySubscriptionResponse) (AzureLargeInstancesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureLargeInstancesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, subscriptionID, options)
			}, nil)
			if err != nil {
				return AzureLargeInstancesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AzureLargeInstancesClient) listBySubscriptionCreateRequest(ctx context.Context, subscriptionID string, options *AzureLargeInstancesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureLargeInstance/azureLargeInstances"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AzureLargeInstancesClient) listBySubscriptionHandleResponse(resp *http.Response) (AzureLargeInstancesClientListBySubscriptionResponse, error) {
	result := AzureLargeInstancesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return AzureLargeInstancesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginRestart - The operation to restart an Azure Large Instance (only for compute instances)
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeInstanceName - Name of the AzureLargeInstance.
//   - options - AzureLargeInstancesClientRestartOptions contains the optional parameters for the AzureLargeInstancesClient.Restart
//     method.
func (client *AzureLargeInstancesClient) BeginRestart(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientRestartOptions) (*runtime.Poller[AzureLargeInstancesClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, subscriptionID, resourceGroupName, azureLargeInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureLargeInstancesClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureLargeInstancesClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restart - The operation to restart an Azure Large Instance (only for compute instances)
func (client *AzureLargeInstancesClient) restart(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientRestartOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureLargeInstancesClient.BeginRestart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartCreateRequest(ctx, subscriptionID, resourceGroupName, azureLargeInstanceName, options)
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

// restartCreateRequest creates the Restart request.
func (client *AzureLargeInstancesClient) restartCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeInstances/{azureLargeInstanceName}/restart"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeInstanceName == "" {
		return nil, errors.New("parameter azureLargeInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeInstanceName}", url.PathEscape(azureLargeInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if options != nil && options.ForceParameter != nil {
		if err := runtime.MarshalAsJSON(req, *options.ForceParameter); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// BeginShutdown - The operation to shutdown an Azure Large Instance (only for compute instances)
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeInstanceName - Name of the AzureLargeInstance.
//   - options - AzureLargeInstancesClientShutdownOptions contains the optional parameters for the AzureLargeInstancesClient.Shutdown
//     method.
func (client *AzureLargeInstancesClient) BeginShutdown(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientShutdownOptions) (*runtime.Poller[AzureLargeInstancesClientShutdownResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.shutdown(ctx, subscriptionID, resourceGroupName, azureLargeInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureLargeInstancesClientShutdownResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureLargeInstancesClientShutdownResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Shutdown - The operation to shutdown an Azure Large Instance (only for compute instances)
func (client *AzureLargeInstancesClient) shutdown(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientShutdownOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureLargeInstancesClient.BeginShutdown"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.shutdownCreateRequest(ctx, subscriptionID, resourceGroupName, azureLargeInstanceName, options)
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

// shutdownCreateRequest creates the Shutdown request.
func (client *AzureLargeInstancesClient) shutdownCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientShutdownOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeInstances/{azureLargeInstanceName}/shutdown"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeInstanceName == "" {
		return nil, errors.New("parameter azureLargeInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeInstanceName}", url.PathEscape(azureLargeInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStart - The operation to start an Azure Large Instance (only for compute instances)
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeInstanceName - Name of the AzureLargeInstance.
//   - options - AzureLargeInstancesClientStartOptions contains the optional parameters for the AzureLargeInstancesClient.Start
//     method.
func (client *AzureLargeInstancesClient) BeginStart(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientStartOptions) (*runtime.Poller[AzureLargeInstancesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, subscriptionID, resourceGroupName, azureLargeInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureLargeInstancesClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureLargeInstancesClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - The operation to start an Azure Large Instance (only for compute instances)
func (client *AzureLargeInstancesClient) start(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientStartOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureLargeInstancesClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, subscriptionID, resourceGroupName, azureLargeInstanceName, options)
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

// startCreateRequest creates the Start request.
func (client *AzureLargeInstancesClient) startCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, options *AzureLargeInstancesClientStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeInstances/{azureLargeInstanceName}/start"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeInstanceName == "" {
		return nil, errors.New("parameter azureLargeInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeInstanceName}", url.PathEscape(azureLargeInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Patches the Tags field of an Azure Large Instance for the specified
// subscription, resource group, and instance name.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeInstanceName - Name of the AzureLargeInstance.
//   - properties - The resource properties to be updated.
//   - options - AzureLargeInstancesClientUpdateOptions contains the optional parameters for the AzureLargeInstancesClient.Update
//     method.
func (client *AzureLargeInstancesClient) Update(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, properties TagsUpdate, options *AzureLargeInstancesClientUpdateOptions) (AzureLargeInstancesClientUpdateResponse, error) {
	var err error
	const operationName = "AzureLargeInstancesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, subscriptionID, resourceGroupName, azureLargeInstanceName, properties, options)
	if err != nil {
		return AzureLargeInstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureLargeInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureLargeInstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AzureLargeInstancesClient) updateCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, azureLargeInstanceName string, properties TagsUpdate, options *AzureLargeInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeInstances/{azureLargeInstanceName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeInstanceName == "" {
		return nil, errors.New("parameter azureLargeInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeInstanceName}", url.PathEscape(azureLargeInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AzureLargeInstancesClient) updateHandleResponse(resp *http.Response) (AzureLargeInstancesClientUpdateResponse, error) {
	result := AzureLargeInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeInstance); err != nil {
		return AzureLargeInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
