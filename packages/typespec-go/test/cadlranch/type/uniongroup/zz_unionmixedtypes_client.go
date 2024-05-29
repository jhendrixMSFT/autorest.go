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

// UnionMixedTypesClient - Describe union of floats "a" | 2 | 3.3
// Don't use this type directly, use [UnionClient.NewUnionMixedTypesClient] instead.
type UnionMixedTypesClient struct {
	internal *azcore.Client
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - UnionMixedTypesClientGetOptions contains the optional parameters for the UnionMixedTypesClient.Get method.
func (client *UnionMixedTypesClient) Get(ctx context.Context, options *UnionMixedTypesClientGetOptions) (UnionMixedTypesClientGetResponse, error) {
	var err error
	const operationName = "UnionMixedTypesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return UnionMixedTypesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnionMixedTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UnionMixedTypesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *UnionMixedTypesClient) getCreateRequest(ctx context.Context, _ *UnionMixedTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/union/mixed-types"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UnionMixedTypesClient) getHandleResponse(resp *http.Response) (UnionMixedTypesClientGetResponse, error) {
	result := UnionMixedTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return UnionMixedTypesClientGetResponse{}, err
	}
	return result, nil
}

// Send -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - UnionMixedTypesClientSendOptions contains the optional parameters for the UnionMixedTypesClient.Send method.
func (client *UnionMixedTypesClient) Send(ctx context.Context, mixedTypesCases MixedTypesCases, options *UnionMixedTypesClientSendOptions) (UnionMixedTypesClientSendResponse, error) {
	var err error
	const operationName = "UnionMixedTypesClient.Send"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sendCreateRequest(ctx, mixedTypesCases, options)
	if err != nil {
		return UnionMixedTypesClientSendResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnionMixedTypesClientSendResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UnionMixedTypesClientSendResponse{}, err
	}
	return UnionMixedTypesClientSendResponse{}, nil
}

// sendCreateRequest creates the Send request.
func (client *UnionMixedTypesClient) sendCreateRequest(ctx context.Context, mixedTypesCases MixedTypesCases, _ *UnionMixedTypesClientSendOptions) (*policy.Request, error) {
	urlPath := "/type/union/mixed-types"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	body := struct {
		Prop MixedTypesCases `json:"prop"`
	}{
		Prop: mixedTypesCases,
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}