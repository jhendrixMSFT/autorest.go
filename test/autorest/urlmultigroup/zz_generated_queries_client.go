//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlmultigroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// QueriesClient contains the methods for the Queries group.
// Don't use this type directly, use NewQueriesClient() instead.
type QueriesClient struct {
	pl runtime.Pipeline
}

// NewQueriesClient creates a new instance of QueriesClient with the specified values.
// options - pass nil to accept the default values.
func NewQueriesClient(options *azcore.ClientOptions) *QueriesClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &QueriesClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// ArrayStringMultiEmpty - Get an empty array [] of string using the multi-array format
// If the operation fails it returns an *azcore.ResponseError type.
// options - QueriesClientArrayStringMultiEmptyOptions contains the optional parameters for the QueriesClient.ArrayStringMultiEmpty
// method.
func (client *QueriesClient) ArrayStringMultiEmpty(ctx context.Context, options *QueriesClientArrayStringMultiEmptyOptions) (QueriesClientArrayStringMultiEmptyResponse, error) {
	req, err := client.arrayStringMultiEmptyCreateRequest(ctx, options)
	if err != nil {
		return QueriesClientArrayStringMultiEmptyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueriesClientArrayStringMultiEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueriesClientArrayStringMultiEmptyResponse{}, runtime.NewResponseError(resp)
	}
	return QueriesClientArrayStringMultiEmptyResponse{RawResponse: resp}, nil
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// ArrayStringMultiNull - Get a null array of string using the multi-array format
// If the operation fails it returns an *azcore.ResponseError type.
// options - QueriesClientArrayStringMultiNullOptions contains the optional parameters for the QueriesClient.ArrayStringMultiNull
// method.
func (client *QueriesClient) ArrayStringMultiNull(ctx context.Context, options *QueriesClientArrayStringMultiNullOptions) (QueriesClientArrayStringMultiNullResponse, error) {
	req, err := client.arrayStringMultiNullCreateRequest(ctx, options)
	if err != nil {
		return QueriesClientArrayStringMultiNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueriesClientArrayStringMultiNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueriesClientArrayStringMultiNullResponse{}, runtime.NewResponseError(resp)
	}
	return QueriesClientArrayStringMultiNullResponse{RawResponse: resp}, nil
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// ArrayStringMultiValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the mult-array
// format
// If the operation fails it returns an *azcore.ResponseError type.
// options - QueriesClientArrayStringMultiValidOptions contains the optional parameters for the QueriesClient.ArrayStringMultiValid
// method.
func (client *QueriesClient) ArrayStringMultiValid(ctx context.Context, options *QueriesClientArrayStringMultiValidOptions) (QueriesClientArrayStringMultiValidResponse, error) {
	req, err := client.arrayStringMultiValidCreateRequest(ctx, options)
	if err != nil {
		return QueriesClientArrayStringMultiValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueriesClientArrayStringMultiValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueriesClientArrayStringMultiValidResponse{}, runtime.NewResponseError(resp)
	}
	return QueriesClientArrayStringMultiValidResponse{RawResponse: resp}, nil
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
