// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package optionalitygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OptionalBytesClient contains the methods for the OptionalBytes group.
// Don't use this type directly, use [OptionalClient.NewOptionalBytesClient] instead.
type OptionalBytesClient struct {
	internal *azcore.Client
}

// GetAll - Get models that will return all properties in the model
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalBytesClientGetAllOptions contains the optional parameters for the OptionalBytesClient.GetAll method.
func (client *OptionalBytesClient) GetAll(ctx context.Context, options *OptionalBytesClientGetAllOptions) (OptionalBytesClientGetAllResponse, error) {
	var err error
	const operationName = "OptionalBytesClient.GetAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAllCreateRequest(ctx, options)
	if err != nil {
		return OptionalBytesClientGetAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalBytesClientGetAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalBytesClientGetAllResponse{}, err
	}
	resp, err := client.getAllHandleResponse(httpResp)
	return resp, err
}

// getAllCreateRequest creates the GetAll request.
func (client *OptionalBytesClient) getAllCreateRequest(ctx context.Context, _ *OptionalBytesClientGetAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/bytes/all"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllHandleResponse handles the GetAll response.
func (client *OptionalBytesClient) getAllHandleResponse(resp *http.Response) (OptionalBytesClientGetAllResponse, error) {
	result := OptionalBytesClientGetAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BytesProperty); err != nil {
		return OptionalBytesClientGetAllResponse{}, err
	}
	return result, nil
}

// GetDefault - Get models that will return the default object
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalBytesClientGetDefaultOptions contains the optional parameters for the OptionalBytesClient.GetDefault
//     method.
func (client *OptionalBytesClient) GetDefault(ctx context.Context, options *OptionalBytesClientGetDefaultOptions) (OptionalBytesClientGetDefaultResponse, error) {
	var err error
	const operationName = "OptionalBytesClient.GetDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDefaultCreateRequest(ctx, options)
	if err != nil {
		return OptionalBytesClientGetDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalBytesClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalBytesClientGetDefaultResponse{}, err
	}
	resp, err := client.getDefaultHandleResponse(httpResp)
	return resp, err
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *OptionalBytesClient) getDefaultCreateRequest(ctx context.Context, _ *OptionalBytesClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/bytes/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *OptionalBytesClient) getDefaultHandleResponse(resp *http.Response) (OptionalBytesClientGetDefaultResponse, error) {
	result := OptionalBytesClientGetDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BytesProperty); err != nil {
		return OptionalBytesClientGetDefaultResponse{}, err
	}
	return result, nil
}

// PutAll - Put a body with all properties present.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalBytesClientPutAllOptions contains the optional parameters for the OptionalBytesClient.PutAll method.
func (client *OptionalBytesClient) PutAll(ctx context.Context, body BytesProperty, options *OptionalBytesClientPutAllOptions) (OptionalBytesClientPutAllResponse, error) {
	var err error
	const operationName = "OptionalBytesClient.PutAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putAllCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalBytesClientPutAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalBytesClientPutAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalBytesClientPutAllResponse{}, err
	}
	return OptionalBytesClientPutAllResponse{}, nil
}

// putAllCreateRequest creates the PutAll request.
func (client *OptionalBytesClient) putAllCreateRequest(ctx context.Context, body BytesProperty, _ *OptionalBytesClientPutAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/bytes/all"
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

// PutDefault - Put a body with default properties.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalBytesClientPutDefaultOptions contains the optional parameters for the OptionalBytesClient.PutDefault
//     method.
func (client *OptionalBytesClient) PutDefault(ctx context.Context, body BytesProperty, options *OptionalBytesClientPutDefaultOptions) (OptionalBytesClientPutDefaultResponse, error) {
	var err error
	const operationName = "OptionalBytesClient.PutDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putDefaultCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalBytesClientPutDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalBytesClientPutDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalBytesClientPutDefaultResponse{}, err
	}
	return OptionalBytesClientPutDefaultResponse{}, nil
}

// putDefaultCreateRequest creates the PutDefault request.
func (client *OptionalBytesClient) putDefaultCreateRequest(ctx context.Context, body BytesProperty, _ *OptionalBytesClientPutDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/bytes/default"
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
