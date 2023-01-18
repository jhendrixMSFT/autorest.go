//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package urlmultigroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"net/http"
)

// QueriesClient contains the methods for the Queries group.
// Don't use this type directly, use a constructor function instead.
type QueriesClient struct {
	internal *azcore.Client
}

// ArrayStringMultiEmpty - Get an empty array [] of string using the multi-array format
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - QueriesClientArrayStringMultiEmptyOptions contains the optional parameters for the QueriesClient.ArrayStringMultiEmpty
//     method.
func (client *QueriesClient) ArrayStringMultiEmpty(ctx context.Context, options *QueriesClientArrayStringMultiEmptyOptions) (result QueriesClientArrayStringMultiEmptyResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "QueriesClient.ArrayStringMultiEmpty", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.arrayStringMultiEmptyCreateRequest(ctx, options)
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

// arrayStringMultiEmptyCreateRequest creates the ArrayStringMultiEmpty request.
func (client *QueriesClient) arrayStringMultiEmptyCreateRequest(ctx context.Context, options *QueriesClientArrayStringMultiEmptyOptions) (*policy.Request, error) {
	urlPath := "/queries/array/multi/string/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ArrayQuery != nil {
		for _, qv := range options.ArrayQuery {
			reqQP.Add("arrayQuery", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ArrayStringMultiNull - Get a null array of string using the multi-array format
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - QueriesClientArrayStringMultiNullOptions contains the optional parameters for the QueriesClient.ArrayStringMultiNull
//     method.
func (client *QueriesClient) ArrayStringMultiNull(ctx context.Context, options *QueriesClientArrayStringMultiNullOptions) (result QueriesClientArrayStringMultiNullResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "QueriesClient.ArrayStringMultiNull", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.arrayStringMultiNullCreateRequest(ctx, options)
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

// arrayStringMultiNullCreateRequest creates the ArrayStringMultiNull request.
func (client *QueriesClient) arrayStringMultiNullCreateRequest(ctx context.Context, options *QueriesClientArrayStringMultiNullOptions) (*policy.Request, error) {
	urlPath := "/queries/array/multi/string/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ArrayQuery != nil {
		for _, qv := range options.ArrayQuery {
			reqQP.Add("arrayQuery", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ArrayStringMultiValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ”] using the mult-array
// format
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - QueriesClientArrayStringMultiValidOptions contains the optional parameters for the QueriesClient.ArrayStringMultiValid
//     method.
func (client *QueriesClient) ArrayStringMultiValid(ctx context.Context, options *QueriesClientArrayStringMultiValidOptions) (result QueriesClientArrayStringMultiValidResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "QueriesClient.ArrayStringMultiValid", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.arrayStringMultiValidCreateRequest(ctx, options)
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

// arrayStringMultiValidCreateRequest creates the ArrayStringMultiValid request.
func (client *QueriesClient) arrayStringMultiValidCreateRequest(ctx context.Context, options *QueriesClientArrayStringMultiValidOptions) (*policy.Request, error) {
	urlPath := "/queries/array/multi/string/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ArrayQuery != nil {
		for _, qv := range options.ArrayQuery {
			reqQP.Add("arrayQuery", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
