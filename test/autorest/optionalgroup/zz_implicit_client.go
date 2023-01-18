//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package optionalgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ImplicitClient contains the methods for the Implicit group.
// Don't use this type directly, use a constructor function instead.
type ImplicitClient struct {
	internal            *azcore.Client
	requiredGlobalPath  string
	requiredGlobalQuery string
	optionalGlobalQuery *int32
}

// GetOptionalGlobalQuery - Test implicitly optional query parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientGetOptionalGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetOptionalGlobalQuery
//     method.
func (client *ImplicitClient) GetOptionalGlobalQuery(ctx context.Context, options *ImplicitClientGetOptionalGlobalQueryOptions) (result ImplicitClientGetOptionalGlobalQueryResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ImplicitClient.GetOptionalGlobalQuery", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getOptionalGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// getOptionalGlobalQueryCreateRequest creates the GetOptionalGlobalQuery request.
func (client *ImplicitClient) getOptionalGlobalQueryCreateRequest(ctx context.Context, options *ImplicitClientGetOptionalGlobalQueryOptions) (*policy.Request, error) {
	urlPath := "/reqopt/global/optional/query"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if client.optionalGlobalQuery != nil {
		reqQP.Set("optional-global-query", strconv.FormatInt(int64(*client.optionalGlobalQuery), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetRequiredGlobalPath - Test implicitly required path parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientGetRequiredGlobalPathOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalPath
//     method.
func (client *ImplicitClient) GetRequiredGlobalPath(ctx context.Context, options *ImplicitClientGetRequiredGlobalPathOptions) (result ImplicitClientGetRequiredGlobalPathResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ImplicitClient.GetRequiredGlobalPath", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getRequiredGlobalPathCreateRequest(ctx, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// getRequiredGlobalPathCreateRequest creates the GetRequiredGlobalPath request.
func (client *ImplicitClient) getRequiredGlobalPathCreateRequest(ctx context.Context, options *ImplicitClientGetRequiredGlobalPathOptions) (*policy.Request, error) {
	urlPath := "/reqopt/global/required/path/{required-global-path}"
	if client.requiredGlobalPath == "" {
		return nil, errors.New("parameter client.requiredGlobalPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{required-global-path}", url.PathEscape(client.requiredGlobalPath))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetRequiredGlobalQuery - Test implicitly required query parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientGetRequiredGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalQuery
//     method.
func (client *ImplicitClient) GetRequiredGlobalQuery(ctx context.Context, options *ImplicitClientGetRequiredGlobalQueryOptions) (result ImplicitClientGetRequiredGlobalQueryResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ImplicitClient.GetRequiredGlobalQuery", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getRequiredGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// getRequiredGlobalQueryCreateRequest creates the GetRequiredGlobalQuery request.
func (client *ImplicitClient) getRequiredGlobalQueryCreateRequest(ctx context.Context, options *ImplicitClientGetRequiredGlobalQueryOptions) (*policy.Request, error) {
	urlPath := "/reqopt/global/required/query"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("required-global-query", client.requiredGlobalQuery)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetRequiredPath - Test implicitly required path parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientGetRequiredPathOptions contains the optional parameters for the ImplicitClient.GetRequiredPath
//     method.
func (client *ImplicitClient) GetRequiredPath(ctx context.Context, pathParameter string, options *ImplicitClientGetRequiredPathOptions) (result ImplicitClientGetRequiredPathResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ImplicitClient.GetRequiredPath", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getRequiredPathCreateRequest(ctx, pathParameter, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// getRequiredPathCreateRequest creates the GetRequiredPath request.
func (client *ImplicitClient) getRequiredPathCreateRequest(ctx context.Context, pathParameter string, options *ImplicitClientGetRequiredPathOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/required/path/{pathParameter}"
	if pathParameter == "" {
		return nil, errors.New("parameter pathParameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pathParameter}", url.PathEscape(pathParameter))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// PutOptionalBinaryBody - Test implicitly optional body parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientPutOptionalBinaryBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBinaryBody
//     method.
func (client *ImplicitClient) PutOptionalBinaryBody(ctx context.Context, bodyParameter io.ReadSeekCloser, options *ImplicitClientPutOptionalBinaryBodyOptions) (result ImplicitClientPutOptionalBinaryBodyResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ImplicitClient.PutOptionalBinaryBody", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.putOptionalBinaryBodyCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// putOptionalBinaryBodyCreateRequest creates the PutOptionalBinaryBody request.
func (client *ImplicitClient) putOptionalBinaryBodyCreateRequest(ctx context.Context, bodyParameter io.ReadSeekCloser, options *ImplicitClientPutOptionalBinaryBodyOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/binary-body"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, req.SetBody(bodyParameter, "application/octet-stream")
}

// PutOptionalBody - Test implicitly optional body parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientPutOptionalBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBody
//     method.
func (client *ImplicitClient) PutOptionalBody(ctx context.Context, bodyParameter string, options *ImplicitClientPutOptionalBodyOptions) (result ImplicitClientPutOptionalBodyResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ImplicitClient.PutOptionalBody", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.putOptionalBodyCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// putOptionalBodyCreateRequest creates the PutOptionalBody request.
func (client *ImplicitClient) putOptionalBodyCreateRequest(ctx context.Context, bodyParameter string, options *ImplicitClientPutOptionalBodyOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/body"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	body := streaming.NopCloser(strings.NewReader(bodyParameter))
	return req, req.SetBody(body, "application/json")
}

// PutOptionalHeader - Test implicitly optional header parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientPutOptionalHeaderOptions contains the optional parameters for the ImplicitClient.PutOptionalHeader
//     method.
func (client *ImplicitClient) PutOptionalHeader(ctx context.Context, options *ImplicitClientPutOptionalHeaderOptions) (result ImplicitClientPutOptionalHeaderResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ImplicitClient.PutOptionalHeader", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.putOptionalHeaderCreateRequest(ctx, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// putOptionalHeaderCreateRequest creates the PutOptionalHeader request.
func (client *ImplicitClient) putOptionalHeaderCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalHeaderOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/header"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.QueryParameter != nil {
		req.Raw().Header["queryParameter"] = []string{*options.QueryParameter}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// PutOptionalQuery - Test implicitly optional query parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientPutOptionalQueryOptions contains the optional parameters for the ImplicitClient.PutOptionalQuery
//     method.
func (client *ImplicitClient) PutOptionalQuery(ctx context.Context, options *ImplicitClientPutOptionalQueryOptions) (result ImplicitClientPutOptionalQueryResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "ImplicitClient.PutOptionalQuery", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.putOptionalQueryCreateRequest(ctx, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// putOptionalQueryCreateRequest creates the PutOptionalQuery request.
func (client *ImplicitClient) putOptionalQueryCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalQueryOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/query"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.QueryParameter != nil {
		reqQP.Set("queryParameter", *options.QueryParameter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
