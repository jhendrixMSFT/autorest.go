// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmongocluster

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

// MongoClustersClient contains the methods for the MongoClusters group.
// Don't use this type directly, use NewMongoClustersClient() instead.
type MongoClustersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMongoClustersClient creates a new instance of MongoClustersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMongoClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MongoClustersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MongoClustersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Check if mongo cluster name is available for use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - location - The name of the Azure region.
//   - body - The CheckAvailability request
//   - options - MongoClustersClientCheckNameAvailabilityOptions contains the optional parameters for the MongoClustersClient.CheckNameAvailability
//     method.
func (client *MongoClustersClient) CheckNameAvailability(ctx context.Context, location string, body CheckNameAvailabilityRequest, options *MongoClustersClientCheckNameAvailabilityOptions) (MongoClustersClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "MongoClustersClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, body, options)
	if err != nil {
		return MongoClustersClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MongoClustersClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MongoClustersClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *MongoClustersClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, body CheckNameAvailabilityRequest, _ *MongoClustersClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/checkMongoClusterNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *MongoClustersClient) checkNameAvailabilityHandleResponse(resp *http.Response) (MongoClustersClientCheckNameAvailabilityResponse, error) {
	result := MongoClustersClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return MongoClustersClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create or update a mongo cluster. Update overwrites all properties for the resource. To only modify
// some of the properties, use PATCH.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mongoClusterName - The name of the mongo cluster.
//   - resource - Resource create parameters.
//   - options - MongoClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the MongoClustersClient.BeginCreateOrUpdate
//     method.
func (client *MongoClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, mongoClusterName string, resource MongoCluster, options *MongoClustersClientBeginCreateOrUpdateOptions) (*runtime.Poller[MongoClustersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, mongoClusterName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MongoClustersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MongoClustersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some
// of the properties, use PATCH.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *MongoClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, mongoClusterName string, resource MongoCluster, options *MongoClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MongoClustersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, mongoClusterName, resource, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MongoClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, mongoClusterName string, resource MongoCluster, _ *MongoClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{mongoClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mongoClusterName == "" {
		return nil, errors.New("parameter mongoClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoClusterName}", url.PathEscape(mongoClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a mongo cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mongoClusterName - The name of the mongo cluster.
//   - options - MongoClustersClientBeginDeleteOptions contains the optional parameters for the MongoClustersClient.BeginDelete
//     method.
func (client *MongoClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, mongoClusterName string, options *MongoClustersClientBeginDeleteOptions) (*runtime.Poller[MongoClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, mongoClusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MongoClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MongoClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a mongo cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *MongoClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, mongoClusterName string, options *MongoClustersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "MongoClustersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, mongoClusterName, options)
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
func (client *MongoClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, mongoClusterName string, _ *MongoClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{mongoClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mongoClusterName == "" {
		return nil, errors.New("parameter mongoClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoClusterName}", url.PathEscape(mongoClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about a mongo cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mongoClusterName - The name of the mongo cluster.
//   - options - MongoClustersClientGetOptions contains the optional parameters for the MongoClustersClient.Get method.
func (client *MongoClustersClient) Get(ctx context.Context, resourceGroupName string, mongoClusterName string, options *MongoClustersClientGetOptions) (MongoClustersClientGetResponse, error) {
	var err error
	const operationName = "MongoClustersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, mongoClusterName, options)
	if err != nil {
		return MongoClustersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MongoClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MongoClustersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MongoClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, mongoClusterName string, _ *MongoClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{mongoClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mongoClusterName == "" {
		return nil, errors.New("parameter mongoClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoClusterName}", url.PathEscape(mongoClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MongoClustersClient) getHandleResponse(resp *http.Response) (MongoClustersClientGetResponse, error) {
	result := MongoClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MongoCluster); err != nil {
		return MongoClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all the mongo clusters in a given subscription.
//
// Generated from API version 2024-07-01
//   - options - MongoClustersClientListOptions contains the optional parameters for the MongoClustersClient.NewListPager method.
func (client *MongoClustersClient) NewListPager(options *MongoClustersClientListOptions) *runtime.Pager[MongoClustersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MongoClustersClientListResponse]{
		More: func(page MongoClustersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MongoClustersClientListResponse) (MongoClustersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MongoClustersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return MongoClustersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MongoClustersClient) listCreateRequest(ctx context.Context, _ *MongoClustersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/mongoClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MongoClustersClient) listHandleResponse(resp *http.Response) (MongoClustersClientListResponse, error) {
	result := MongoClustersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return MongoClustersClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the mongo clusters in a given resource group.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - MongoClustersClientListByResourceGroupOptions contains the optional parameters for the MongoClustersClient.NewListByResourceGroupPager
//     method.
func (client *MongoClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *MongoClustersClientListByResourceGroupOptions) *runtime.Pager[MongoClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[MongoClustersClientListByResourceGroupResponse]{
		More: func(page MongoClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MongoClustersClientListByResourceGroupResponse) (MongoClustersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MongoClustersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return MongoClustersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MongoClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *MongoClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters"
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
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MongoClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (MongoClustersClientListByResourceGroupResponse, error) {
	result := MongoClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return MongoClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListConnectionStrings - List mongo cluster connection strings. This includes the default connection string using SCRAM-SHA-256,
// as well as other connection strings supported by the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mongoClusterName - The name of the mongo cluster.
//   - options - MongoClustersClientListConnectionStringsOptions contains the optional parameters for the MongoClustersClient.ListConnectionStrings
//     method.
func (client *MongoClustersClient) ListConnectionStrings(ctx context.Context, resourceGroupName string, mongoClusterName string, options *MongoClustersClientListConnectionStringsOptions) (MongoClustersClientListConnectionStringsResponse, error) {
	var err error
	const operationName = "MongoClustersClient.ListConnectionStrings"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listConnectionStringsCreateRequest(ctx, resourceGroupName, mongoClusterName, options)
	if err != nil {
		return MongoClustersClientListConnectionStringsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MongoClustersClientListConnectionStringsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MongoClustersClientListConnectionStringsResponse{}, err
	}
	resp, err := client.listConnectionStringsHandleResponse(httpResp)
	return resp, err
}

// listConnectionStringsCreateRequest creates the ListConnectionStrings request.
func (client *MongoClustersClient) listConnectionStringsCreateRequest(ctx context.Context, resourceGroupName string, mongoClusterName string, _ *MongoClustersClientListConnectionStringsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{mongoClusterName}/listConnectionStrings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mongoClusterName == "" {
		return nil, errors.New("parameter mongoClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoClusterName}", url.PathEscape(mongoClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listConnectionStringsHandleResponse handles the ListConnectionStrings response.
func (client *MongoClustersClient) listConnectionStringsHandleResponse(resp *http.Response) (MongoClustersClientListConnectionStringsResponse, error) {
	result := MongoClustersClientListConnectionStringsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListConnectionStringsResult); err != nil {
		return MongoClustersClientListConnectionStringsResponse{}, err
	}
	return result, nil
}

// BeginPromote - Promotes a replica mongo cluster to a primary role.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mongoClusterName - The name of the mongo cluster.
//   - body - The content of the action request
//   - options - MongoClustersClientBeginPromoteOptions contains the optional parameters for the MongoClustersClient.BeginPromote
//     method.
func (client *MongoClustersClient) BeginPromote(ctx context.Context, resourceGroupName string, mongoClusterName string, body PromoteReplicaRequest, options *MongoClustersClientBeginPromoteOptions) (*runtime.Poller[MongoClustersClientPromoteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.promote(ctx, resourceGroupName, mongoClusterName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MongoClustersClientPromoteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MongoClustersClientPromoteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Promote - Promotes a replica mongo cluster to a primary role.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *MongoClustersClient) promote(ctx context.Context, resourceGroupName string, mongoClusterName string, body PromoteReplicaRequest, options *MongoClustersClientBeginPromoteOptions) (*http.Response, error) {
	var err error
	const operationName = "MongoClustersClient.BeginPromote"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.promoteCreateRequest(ctx, resourceGroupName, mongoClusterName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// promoteCreateRequest creates the Promote request.
func (client *MongoClustersClient) promoteCreateRequest(ctx context.Context, resourceGroupName string, mongoClusterName string, body PromoteReplicaRequest, _ *MongoClustersClientBeginPromoteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{mongoClusterName}/promote"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mongoClusterName == "" {
		return nil, errors.New("parameter mongoClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoClusterName}", url.PathEscape(mongoClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Updates an existing mongo cluster. The request body can contain one to many of the properties present in
// the normal mongo cluster definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mongoClusterName - The name of the mongo cluster.
//   - properties - The resource properties to be updated.
//   - options - MongoClustersClientBeginUpdateOptions contains the optional parameters for the MongoClustersClient.BeginUpdate
//     method.
func (client *MongoClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, mongoClusterName string, properties Update, options *MongoClustersClientBeginUpdateOptions) (*runtime.Poller[MongoClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, mongoClusterName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MongoClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MongoClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates an existing mongo cluster. The request body can contain one to many of the properties present in the normal
// mongo cluster definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *MongoClustersClient) update(ctx context.Context, resourceGroupName string, mongoClusterName string, properties Update, options *MongoClustersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MongoClustersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, mongoClusterName, properties, options)
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
func (client *MongoClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, mongoClusterName string, properties Update, _ *MongoClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{mongoClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mongoClusterName == "" {
		return nil, errors.New("parameter mongoClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoClusterName}", url.PathEscape(mongoClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
