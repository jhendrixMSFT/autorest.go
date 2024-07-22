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

// OptionalPlainDateClient contains the methods for the OptionalPlainDate group.
// Don't use this type directly, use [OptionalClient.NewOptionalPlainDateClient] instead.
type OptionalPlainDateClient struct {
	internal *azcore.Client
}

// GetAll - Get models that will return all properties in the model
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalPlainDateClientGetAllOptions contains the optional parameters for the OptionalPlainDateClient.GetAll
//     method.
func (client *OptionalPlainDateClient) GetAll(ctx context.Context, options *OptionalPlainDateClientGetAllOptions) (OptionalPlainDateClientGetAllResponse, error) {
	var err error
	const operationName = "OptionalPlainDateClient.GetAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAllCreateRequest(ctx, options)
	if err != nil {
		return OptionalPlainDateClientGetAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalPlainDateClientGetAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalPlainDateClientGetAllResponse{}, err
	}
	resp, err := client.getAllHandleResponse(httpResp)
	return resp, err
}

// getAllCreateRequest creates the GetAll request.
func (client *OptionalPlainDateClient) getAllCreateRequest(ctx context.Context, _ *OptionalPlainDateClientGetAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/plainDate/all"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllHandleResponse handles the GetAll response.
func (client *OptionalPlainDateClient) getAllHandleResponse(resp *http.Response) (OptionalPlainDateClientGetAllResponse, error) {
	result := OptionalPlainDateClientGetAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PlainDateProperty); err != nil {
		return OptionalPlainDateClientGetAllResponse{}, err
	}
	return result, nil
}

// GetDefault - Get models that will return the default object
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalPlainDateClientGetDefaultOptions contains the optional parameters for the OptionalPlainDateClient.GetDefault
//     method.
func (client *OptionalPlainDateClient) GetDefault(ctx context.Context, options *OptionalPlainDateClientGetDefaultOptions) (OptionalPlainDateClientGetDefaultResponse, error) {
	var err error
	const operationName = "OptionalPlainDateClient.GetDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDefaultCreateRequest(ctx, options)
	if err != nil {
		return OptionalPlainDateClientGetDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalPlainDateClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalPlainDateClientGetDefaultResponse{}, err
	}
	resp, err := client.getDefaultHandleResponse(httpResp)
	return resp, err
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *OptionalPlainDateClient) getDefaultCreateRequest(ctx context.Context, _ *OptionalPlainDateClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/plainDate/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *OptionalPlainDateClient) getDefaultHandleResponse(resp *http.Response) (OptionalPlainDateClientGetDefaultResponse, error) {
	result := OptionalPlainDateClientGetDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PlainDateProperty); err != nil {
		return OptionalPlainDateClientGetDefaultResponse{}, err
	}
	return result, nil
}

// PutAll - Put a body with all properties present.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalPlainDateClientPutAllOptions contains the optional parameters for the OptionalPlainDateClient.PutAll
//     method.
func (client *OptionalPlainDateClient) PutAll(ctx context.Context, body PlainDateProperty, options *OptionalPlainDateClientPutAllOptions) (OptionalPlainDateClientPutAllResponse, error) {
	var err error
	const operationName = "OptionalPlainDateClient.PutAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putAllCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalPlainDateClientPutAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalPlainDateClientPutAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalPlainDateClientPutAllResponse{}, err
	}
	return OptionalPlainDateClientPutAllResponse{}, nil
}

// putAllCreateRequest creates the PutAll request.
func (client *OptionalPlainDateClient) putAllCreateRequest(ctx context.Context, body PlainDateProperty, _ *OptionalPlainDateClientPutAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/plainDate/all"
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
//   - options - OptionalPlainDateClientPutDefaultOptions contains the optional parameters for the OptionalPlainDateClient.PutDefault
//     method.
func (client *OptionalPlainDateClient) PutDefault(ctx context.Context, body PlainDateProperty, options *OptionalPlainDateClientPutDefaultOptions) (OptionalPlainDateClientPutDefaultResponse, error) {
	var err error
	const operationName = "OptionalPlainDateClient.PutDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putDefaultCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalPlainDateClientPutDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalPlainDateClientPutDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalPlainDateClientPutDefaultResponse{}, err
	}
	return OptionalPlainDateClientPutDefaultResponse{}, nil
}

// putDefaultCreateRequest creates the PutDefault request.
func (client *OptionalPlainDateClient) putDefaultCreateRequest(ctx context.Context, body PlainDateProperty, _ *OptionalPlainDateClientPutDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/plainDate/default"
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
