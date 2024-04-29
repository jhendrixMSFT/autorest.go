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

// ValueTypesDatetimeClient contains the methods for the Type.Property.ValueTypes namespace.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesDatetimeClient] instead.
type ValueTypesDatetimeClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - ValueTypesDatetimeClientGetOptions contains the optional parameters for the ValueTypesDatetimeClient.Get method.
func (client *ValueTypesDatetimeClient) Get(ctx context.Context, options *ValueTypesDatetimeClientGetOptions) (ValueTypesDatetimeClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesDatetimeClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesDatetimeClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesDatetimeClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesDatetimeClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesDatetimeClient) getCreateRequest(ctx context.Context, options *ValueTypesDatetimeClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/datetime"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesDatetimeClient) getHandleResponse(resp *http.Response) (ValueTypesDatetimeClientGetResponse, error) {
	result := ValueTypesDatetimeClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatetimeProperty); err != nil {
		return ValueTypesDatetimeClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - ValueTypesDatetimeClientPutOptions contains the optional parameters for the ValueTypesDatetimeClient.Put method.
func (client *ValueTypesDatetimeClient) Put(ctx context.Context, body DatetimeProperty, options *ValueTypesDatetimeClientPutOptions) (ValueTypesDatetimeClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesDatetimeClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesDatetimeClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesDatetimeClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesDatetimeClientPutResponse{}, err
	}
	return ValueTypesDatetimeClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesDatetimeClient) putCreateRequest(ctx context.Context, body DatetimeProperty, options *ValueTypesDatetimeClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/datetime"
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
