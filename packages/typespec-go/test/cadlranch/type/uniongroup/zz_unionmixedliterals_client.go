// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package uniongroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// UnionMixedLiteralsClient - Describe union of floats "a" | 2 | 3.3
// Don't use this type directly, use [UnionClient.NewUnionMixedLiteralsClient] instead.
type UnionMixedLiteralsClient struct {
	internal *azcore.Client
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - UnionMixedLiteralsClientGetOptions contains the optional parameters for the UnionMixedLiteralsClient.Get method.
func (client *UnionMixedLiteralsClient) Get(ctx context.Context, options *UnionMixedLiteralsClientGetOptions) (UnionMixedLiteralsClientGetResponse, error) {
	var err error
	const operationName = "UnionMixedLiteralsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return UnionMixedLiteralsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnionMixedLiteralsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UnionMixedLiteralsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *UnionMixedLiteralsClient) getCreateRequest(ctx context.Context, _ *UnionMixedLiteralsClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/union/mixed-literals"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UnionMixedLiteralsClient) getHandleResponse(resp *http.Response) (UnionMixedLiteralsClientGetResponse, error) {
	result := UnionMixedLiteralsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetResponse1); err != nil {
		return UnionMixedLiteralsClientGetResponse{}, err
	}
	return result, nil
}

// Send -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - UnionMixedLiteralsClientSendOptions contains the optional parameters for the UnionMixedLiteralsClient.Send method.
func (client *UnionMixedLiteralsClient) Send(ctx context.Context, sendRequest1 SendRequest1, options *UnionMixedLiteralsClientSendOptions) (UnionMixedLiteralsClientSendResponse, error) {
	var err error
	const operationName = "UnionMixedLiteralsClient.Send"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sendCreateRequest(ctx, sendRequest1, options)
	if err != nil {
		return UnionMixedLiteralsClientSendResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnionMixedLiteralsClientSendResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UnionMixedLiteralsClientSendResponse{}, err
	}
	return UnionMixedLiteralsClientSendResponse{}, nil
}

// sendCreateRequest creates the Send request.
func (client *UnionMixedLiteralsClient) sendCreateRequest(ctx context.Context, sendRequest1 SendRequest1, _ *UnionMixedLiteralsClientSendOptions) (*policy.Request, error) {
	urlPath := "/type/union/mixed-literals"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sendRequest1); err != nil {
		return nil, err
	}
	return req, nil
}
