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

// ValueTypesExtensibleEnumClient contains the methods for the Type.Property.ValueTypes namespace.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesExtensibleEnumClient] instead.
type ValueTypesExtensibleEnumClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - ValueTypesExtensibleEnumClientGetOptions contains the optional parameters for the ValueTypesExtensibleEnumClient.Get
//     method.
func (client *ValueTypesExtensibleEnumClient) Get(ctx context.Context, options *ValueTypesExtensibleEnumClientGetOptions) (ValueTypesExtensibleEnumClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesExtensibleEnumClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesExtensibleEnumClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesExtensibleEnumClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesExtensibleEnumClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesExtensibleEnumClient) getCreateRequest(ctx context.Context, options *ValueTypesExtensibleEnumClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/extensible-enum"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesExtensibleEnumClient) getHandleResponse(resp *http.Response) (ValueTypesExtensibleEnumClientGetResponse, error) {
	result := ValueTypesExtensibleEnumClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensibleEnumProperty); err != nil {
		return ValueTypesExtensibleEnumClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - ValueTypesExtensibleEnumClientPutOptions contains the optional parameters for the ValueTypesExtensibleEnumClient.Put
//     method.
func (client *ValueTypesExtensibleEnumClient) Put(ctx context.Context, body ExtensibleEnumProperty, options *ValueTypesExtensibleEnumClientPutOptions) (ValueTypesExtensibleEnumClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesExtensibleEnumClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesExtensibleEnumClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesExtensibleEnumClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesExtensibleEnumClientPutResponse{}, err
	}
	return ValueTypesExtensibleEnumClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesExtensibleEnumClient) putCreateRequest(ctx context.Context, body ExtensibleEnumProperty, options *ValueTypesExtensibleEnumClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/extensible-enum"
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
