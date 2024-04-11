// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package emptygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// EmptyClient - Illustrates usage of empty model used in operation's parameters and responses.
// Don't use this type directly, use NewEmptyClientWithNoCredential() instead.
type EmptyClient struct {
	internal *azcore.Client
}

// EmptyClientOptions contains the optional values for creating a [EmptyClient].
type EmptyClientOptions struct {
	azcore.ClientOptions
}

// NewEmptyClientWithNoCredential creates a new instance of [EmptyClient] with the specified values.
//   - options - EmptyClientOptions contains the optional values for creating a [EmptyClient]
func NewEmptyClientWithNoCredential(options *EmptyClientOptions) (*EmptyClient, error) {
	if options == nil {
		options = &EmptyClientOptions{}
	}
	cl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	emptyClient := &EmptyClient{
		internal: cl,
	}
	return emptyClient, nil
}

// - options - GetEmptyOptions contains the optional parameters for the EmptyClient.GetEmpty method.
func (client *EmptyClient) GetEmpty(ctx context.Context, options *GetEmptyOptions) (GetEmptyResponse, error) {
	var err error
	const operationName = "EmptyClient.GetEmpty"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return GetEmptyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GetEmptyResponse{}, err
	}
	resp, err := client.getEmptyHandleResponse(httpResp)
	return resp, err
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *EmptyClient) getEmptyCreateRequest(ctx context.Context, options *GetEmptyOptions) (*policy.Request, error) {
	urlPath := "/type/model/empty/alone"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *EmptyClient) getEmptyHandleResponse(resp *http.Response) (GetEmptyResponse, error) {
	result := GetEmptyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmptyOutput); err != nil {
		return GetEmptyResponse{}, err
	}
	return result, nil
}

// - options - PostRoundTripEmptyOptions contains the optional parameters for the EmptyClient.PostRoundTripEmpty method.
func (client *EmptyClient) PostRoundTripEmpty(ctx context.Context, body EmptyInputOutput, options *PostRoundTripEmptyOptions) (PostRoundTripEmptyResponse, error) {
	var err error
	const operationName = "EmptyClient.PostRoundTripEmpty"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postRoundTripEmptyCreateRequest(ctx, body, options)
	if err != nil {
		return PostRoundTripEmptyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PostRoundTripEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PostRoundTripEmptyResponse{}, err
	}
	resp, err := client.postRoundTripEmptyHandleResponse(httpResp)
	return resp, err
}

// postRoundTripEmptyCreateRequest creates the PostRoundTripEmpty request.
func (client *EmptyClient) postRoundTripEmptyCreateRequest(ctx context.Context, body EmptyInputOutput, options *PostRoundTripEmptyOptions) (*policy.Request, error) {
	urlPath := "/type/model/empty/round-trip"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// postRoundTripEmptyHandleResponse handles the PostRoundTripEmpty response.
func (client *EmptyClient) postRoundTripEmptyHandleResponse(resp *http.Response) (PostRoundTripEmptyResponse, error) {
	result := PostRoundTripEmptyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmptyInputOutput); err != nil {
		return PostRoundTripEmptyResponse{}, err
	}
	return result, nil
}

// - options - PutEmptyOptions contains the optional parameters for the EmptyClient.PutEmpty method.
func (client *EmptyClient) PutEmpty(ctx context.Context, input EmptyInput, options *PutEmptyOptions) (PutEmptyResponse, error) {
	var err error
	const operationName = "EmptyClient.PutEmpty"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putEmptyCreateRequest(ctx, input, options)
	if err != nil {
		return PutEmptyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PutEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PutEmptyResponse{}, err
	}
	return PutEmptyResponse{}, nil
}

// putEmptyCreateRequest creates the PutEmpty request.
func (client *EmptyClient) putEmptyCreateRequest(ctx context.Context, input EmptyInput, options *PutEmptyOptions) (*policy.Request, error) {
	urlPath := "/type/model/empty/alone"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}
