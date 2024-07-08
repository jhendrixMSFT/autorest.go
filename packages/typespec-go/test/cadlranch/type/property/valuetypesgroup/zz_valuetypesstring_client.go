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

// ValueTypesStringClient contains the methods for the ValueTypesString group.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesStringClient] instead.
type ValueTypesStringClient struct {
	internal *azcore.Client
}

// Get - Get call
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ValueTypesStringClientGetOptions contains the optional parameters for the ValueTypesStringClient.Get method.
func (client *ValueTypesStringClient) Get(ctx context.Context, options *ValueTypesStringClientGetOptions) (ValueTypesStringClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesStringClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesStringClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesStringClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesStringClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesStringClient) getCreateRequest(ctx context.Context, _ *ValueTypesStringClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/string"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesStringClient) getHandleResponse(resp *http.Response) (ValueTypesStringClientGetResponse, error) {
	result := ValueTypesStringClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StringProperty); err != nil {
		return ValueTypesStringClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
// If the operation fails it returns an *azcore.ResponseError type.
//   - body - body
//   - options - ValueTypesStringClientPutOptions contains the optional parameters for the ValueTypesStringClient.Put method.
func (client *ValueTypesStringClient) Put(ctx context.Context, body StringProperty, options *ValueTypesStringClientPutOptions) (ValueTypesStringClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesStringClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesStringClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesStringClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesStringClientPutResponse{}, err
	}
	return ValueTypesStringClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesStringClient) putCreateRequest(ctx context.Context, body StringProperty, _ *ValueTypesStringClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/string"
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
