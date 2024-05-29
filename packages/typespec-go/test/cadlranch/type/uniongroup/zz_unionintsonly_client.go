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

// UnionIntsOnlyClient - Describe union of integer 1 | 2 | 3
// Don't use this type directly, use [UnionClient.NewUnionIntsOnlyClient] instead.
type UnionIntsOnlyClient struct {
	internal *azcore.Client
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - UnionIntsOnlyClientGetOptions contains the optional parameters for the UnionIntsOnlyClient.Get method.
func (client *UnionIntsOnlyClient) Get(ctx context.Context, options *UnionIntsOnlyClientGetOptions) (UnionIntsOnlyClientGetResponse, error) {
	var err error
	const operationName = "UnionIntsOnlyClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return UnionIntsOnlyClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnionIntsOnlyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UnionIntsOnlyClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *UnionIntsOnlyClient) getCreateRequest(ctx context.Context, _ *UnionIntsOnlyClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/union/ints-only"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UnionIntsOnlyClient) getHandleResponse(resp *http.Response) (UnionIntsOnlyClientGetResponse, error) {
	result := UnionIntsOnlyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetResponse6); err != nil {
		return UnionIntsOnlyClientGetResponse{}, err
	}
	return result, nil
}

// Send -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - UnionIntsOnlyClientSendOptions contains the optional parameters for the UnionIntsOnlyClient.Send method.
func (client *UnionIntsOnlyClient) Send(ctx context.Context, prop GetResponseProp3, options *UnionIntsOnlyClientSendOptions) (UnionIntsOnlyClientSendResponse, error) {
	var err error
	const operationName = "UnionIntsOnlyClient.Send"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sendCreateRequest(ctx, prop, options)
	if err != nil {
		return UnionIntsOnlyClientSendResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnionIntsOnlyClientSendResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UnionIntsOnlyClientSendResponse{}, err
	}
	return UnionIntsOnlyClientSendResponse{}, nil
}

// sendCreateRequest creates the Send request.
func (client *UnionIntsOnlyClient) sendCreateRequest(ctx context.Context, prop GetResponseProp3, _ *UnionIntsOnlyClientSendOptions) (*policy.Request, error) {
	urlPath := "/type/union/ints-only"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	body := struct {
		Prop GetResponseProp3 `json:"prop"`
	}{
		Prop: prop,
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
