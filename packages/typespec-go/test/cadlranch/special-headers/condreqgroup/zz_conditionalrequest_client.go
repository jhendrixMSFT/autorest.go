// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package condreqgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ConditionalRequestClient - Illustrates conditional request headers
// Don't use this type directly, use NewConditionalRequestClientWithNoCredential() instead.
type ConditionalRequestClient struct {
	internal *azcore.Client
}

// ConditionalRequestClientOptions contains the optional values for creating a [ConditionalRequestClient].
type ConditionalRequestClientOptions struct {
	azcore.ClientOptions
}

// NewConditionalRequestClientWithNoCredential creates a new instance of [ConditionalRequestClient] with the specified values.
//   - options - ConditionalRequestClientOptions contains the optional values for creating a [ConditionalRequestClient]
func NewConditionalRequestClientWithNoCredential(options *ConditionalRequestClientOptions) (*ConditionalRequestClient, error) {
	if options == nil {
		options = &ConditionalRequestClientOptions{}
	}
	cl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	conditionalRequestClient := &ConditionalRequestClient{
		internal: cl,
	}
	return conditionalRequestClient, nil
}

// PostIfMatch - Check when only If-Match in header is defined.
//   - options - ConditionalRequestClientPostIfMatchOptions contains the optional parameters for the ConditionalRequestClient.PostIfMatch
//     method.
func (client *ConditionalRequestClient) PostIfMatch(ctx context.Context, options *ConditionalRequestClientPostIfMatchOptions) (ConditionalRequestClientPostIfMatchResponse, error) {
	var err error
	const operationName = "ConditionalRequestClient.PostIfMatch"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postIfMatchCreateRequest(ctx, options)
	if err != nil {
		return ConditionalRequestClientPostIfMatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConditionalRequestClientPostIfMatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConditionalRequestClientPostIfMatchResponse{}, err
	}
	return ConditionalRequestClientPostIfMatchResponse{}, nil
}

// postIfMatchCreateRequest creates the PostIfMatch request.
func (client *ConditionalRequestClient) postIfMatchCreateRequest(ctx context.Context, options *ConditionalRequestClientPostIfMatchOptions) (*policy.Request, error) {
	urlPath := "/special-headers/conditional-request/if-match"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	return req, nil
}

// PostIfNoneMatch - Check when only If-None-Match in header is defined.
//   - options - ConditionalRequestClientPostIfNoneMatchOptions contains the optional parameters for the ConditionalRequestClient.PostIfNoneMatch
//     method.
func (client *ConditionalRequestClient) PostIfNoneMatch(ctx context.Context, options *ConditionalRequestClientPostIfNoneMatchOptions) (ConditionalRequestClientPostIfNoneMatchResponse, error) {
	var err error
	const operationName = "ConditionalRequestClient.PostIfNoneMatch"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postIfNoneMatchCreateRequest(ctx, options)
	if err != nil {
		return ConditionalRequestClientPostIfNoneMatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConditionalRequestClientPostIfNoneMatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConditionalRequestClientPostIfNoneMatchResponse{}, err
	}
	return ConditionalRequestClientPostIfNoneMatchResponse{}, nil
}

// postIfNoneMatchCreateRequest creates the PostIfNoneMatch request.
func (client *ConditionalRequestClient) postIfNoneMatchCreateRequest(ctx context.Context, options *ConditionalRequestClientPostIfNoneMatchOptions) (*policy.Request, error) {
	urlPath := "/special-headers/conditional-request/if-none-match"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	return req, nil
}
