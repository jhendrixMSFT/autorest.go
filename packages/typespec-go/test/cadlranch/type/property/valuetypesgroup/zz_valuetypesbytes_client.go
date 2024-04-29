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

// ValueTypesBytesClient contains the methods for the Type.Property.ValueTypes namespace.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesBytesClient] instead.
type ValueTypesBytesClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - ValueTypesBytesClientGetOptions contains the optional parameters for the ValueTypesBytesClient.Get method.
func (client *ValueTypesBytesClient) Get(ctx context.Context, options *ValueTypesBytesClientGetOptions) (ValueTypesBytesClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesBytesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesBytesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesBytesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesBytesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesBytesClient) getCreateRequest(ctx context.Context, options *ValueTypesBytesClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/bytes"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesBytesClient) getHandleResponse(resp *http.Response) (ValueTypesBytesClientGetResponse, error) {
	result := ValueTypesBytesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BytesProperty); err != nil {
		return ValueTypesBytesClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - ValueTypesBytesClientPutOptions contains the optional parameters for the ValueTypesBytesClient.Put method.
func (client *ValueTypesBytesClient) Put(ctx context.Context, body BytesProperty, options *ValueTypesBytesClientPutOptions) (ValueTypesBytesClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesBytesClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesBytesClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesBytesClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesBytesClientPutResponse{}, err
	}
	return ValueTypesBytesClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesBytesClient) putCreateRequest(ctx context.Context, body BytesProperty, options *ValueTypesBytesClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/bytes"
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