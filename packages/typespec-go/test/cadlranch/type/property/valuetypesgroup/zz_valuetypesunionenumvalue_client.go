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

// ValueTypesUnionEnumValueClient contains the methods for the Type.Property.ValueTypes namespace.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesUnionEnumValueClient] instead.
type ValueTypesUnionEnumValueClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - ValueTypesUnionEnumValueClientGetOptions contains the optional parameters for the ValueTypesUnionEnumValueClient.Get
//     method.
func (client *ValueTypesUnionEnumValueClient) Get(ctx context.Context, options *ValueTypesUnionEnumValueClientGetOptions) (ValueTypesUnionEnumValueClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesUnionEnumValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesUnionEnumValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesUnionEnumValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesUnionEnumValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesUnionEnumValueClient) getCreateRequest(ctx context.Context, options *ValueTypesUnionEnumValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/union-enum-value"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesUnionEnumValueClient) getHandleResponse(resp *http.Response) (ValueTypesUnionEnumValueClientGetResponse, error) {
	result := ValueTypesUnionEnumValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnionEnumValueProperty); err != nil {
		return ValueTypesUnionEnumValueClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - ValueTypesUnionEnumValueClientPutOptions contains the optional parameters for the ValueTypesUnionEnumValueClient.Put
//     method.
func (client *ValueTypesUnionEnumValueClient) Put(ctx context.Context, body UnionEnumValueProperty, options *ValueTypesUnionEnumValueClientPutOptions) (ValueTypesUnionEnumValueClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesUnionEnumValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesUnionEnumValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesUnionEnumValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesUnionEnumValueClientPutResponse{}, err
	}
	return ValueTypesUnionEnumValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesUnionEnumValueClient) putCreateRequest(ctx context.Context, body UnionEnumValueProperty, options *ValueTypesUnionEnumValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/union-enum-value"
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
