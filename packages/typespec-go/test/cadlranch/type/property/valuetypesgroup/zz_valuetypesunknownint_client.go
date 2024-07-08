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

// ValueTypesUnknownIntClient contains the methods for the ValueTypesUnknownInt group.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesUnknownIntClient] instead.
type ValueTypesUnknownIntClient struct {
	internal *azcore.Client
}

// Get - Get call
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ValueTypesUnknownIntClientGetOptions contains the optional parameters for the ValueTypesUnknownIntClient.Get
//     method.
func (client *ValueTypesUnknownIntClient) Get(ctx context.Context, options *ValueTypesUnknownIntClientGetOptions) (ValueTypesUnknownIntClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesUnknownIntClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesUnknownIntClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesUnknownIntClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesUnknownIntClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesUnknownIntClient) getCreateRequest(ctx context.Context, _ *ValueTypesUnknownIntClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/unknown/int"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesUnknownIntClient) getHandleResponse(resp *http.Response) (ValueTypesUnknownIntClientGetResponse, error) {
	result := ValueTypesUnknownIntClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnknownIntProperty); err != nil {
		return ValueTypesUnknownIntClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
// If the operation fails it returns an *azcore.ResponseError type.
//   - body - body
//   - options - ValueTypesUnknownIntClientPutOptions contains the optional parameters for the ValueTypesUnknownIntClient.Put
//     method.
func (client *ValueTypesUnknownIntClient) Put(ctx context.Context, body UnknownIntProperty, options *ValueTypesUnknownIntClientPutOptions) (ValueTypesUnknownIntClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesUnknownIntClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesUnknownIntClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesUnknownIntClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesUnknownIntClientPutResponse{}, err
	}
	return ValueTypesUnknownIntClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesUnknownIntClient) putCreateRequest(ctx context.Context, body UnknownIntProperty, _ *ValueTypesUnknownIntClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/unknown/int"
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
