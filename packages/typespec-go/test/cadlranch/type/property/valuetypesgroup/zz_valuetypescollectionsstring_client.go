// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package valuetypesgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ValueTypesCollectionsStringClient contains the methods for the ValueTypesCollectionsString group.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesCollectionsStringClient] instead.
type ValueTypesCollectionsStringClient struct {
	internal *azcore.Client
}

// Get - Get call
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ValueTypesCollectionsStringClientGetOptions contains the optional parameters for the ValueTypesCollectionsStringClient.Get
//     method.
func (client *ValueTypesCollectionsStringClient) Get(ctx context.Context, options *ValueTypesCollectionsStringClientGetOptions) (ValueTypesCollectionsStringClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesCollectionsStringClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesCollectionsStringClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesCollectionsStringClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesCollectionsStringClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesCollectionsStringClient) getCreateRequest(ctx context.Context, _ *ValueTypesCollectionsStringClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/collections/string"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesCollectionsStringClient) getHandleResponse(resp *http.Response) (ValueTypesCollectionsStringClientGetResponse, error) {
	result := ValueTypesCollectionsStringClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CollectionsStringProperty); err != nil {
		return ValueTypesCollectionsStringClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
// If the operation fails it returns an *azcore.ResponseError type.
//   - body - body
//   - options - ValueTypesCollectionsStringClientPutOptions contains the optional parameters for the ValueTypesCollectionsStringClient.Put
//     method.
func (client *ValueTypesCollectionsStringClient) Put(ctx context.Context, body CollectionsStringProperty, options *ValueTypesCollectionsStringClientPutOptions) (ValueTypesCollectionsStringClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesCollectionsStringClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesCollectionsStringClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesCollectionsStringClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesCollectionsStringClientPutResponse{}, err
	}
	return ValueTypesCollectionsStringClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesCollectionsStringClient) putCreateRequest(ctx context.Context, body CollectionsStringProperty, _ *ValueTypesCollectionsStringClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/collections/string"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
