//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPServerFailureClient contains the methods for the HTTPServerFailure group.
// Don't use this type directly, use NewHTTPServerFailureClient() instead.
type HTTPServerFailureClient struct {
	pl runtime.Pipeline
}

// NewHTTPServerFailureClient creates a new instance of HTTPServerFailureClient with the specified values.
// options - pass nil to accept the default values.
func NewHTTPServerFailureClient(options *azcore.ClientOptions) *HTTPServerFailureClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &HTTPServerFailureClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// Delete505 - Return 505 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPServerFailureClientDelete505Options contains the optional parameters for the HTTPServerFailureClient.Delete505
// method.
func (client *HTTPServerFailureClient) Delete505(ctx context.Context, options *HTTPServerFailureClientDelete505Options) (HTTPServerFailureClientDelete505Response, error) {
	req, err := client.delete505CreateRequest(ctx, options)
	if err != nil {
		return HTTPServerFailureClientDelete505Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPServerFailureClientDelete505Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPServerFailureClientDelete505Response{}, runtime.NewResponseError(resp)
	}
	return HTTPServerFailureClientDelete505Response{RawResponse: resp}, nil
}

// delete505CreateRequest creates the Delete505 request.
func (client *HTTPServerFailureClient) delete505CreateRequest(ctx context.Context, options *HTTPServerFailureClientDelete505Options) (*policy.Request, error) {
	urlPath := "/http/failure/server/505"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Get501 - Return 501 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPServerFailureClientGet501Options contains the optional parameters for the HTTPServerFailureClient.Get501
// method.
func (client *HTTPServerFailureClient) Get501(ctx context.Context, options *HTTPServerFailureClientGet501Options) (HTTPServerFailureClientGet501Response, error) {
	req, err := client.get501CreateRequest(ctx, options)
	if err != nil {
		return HTTPServerFailureClientGet501Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPServerFailureClientGet501Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPServerFailureClientGet501Response{}, runtime.NewResponseError(resp)
	}
	return HTTPServerFailureClientGet501Response{RawResponse: resp}, nil
}

// get501CreateRequest creates the Get501 request.
func (client *HTTPServerFailureClient) get501CreateRequest(ctx context.Context, options *HTTPServerFailureClientGet501Options) (*policy.Request, error) {
	urlPath := "/http/failure/server/501"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Head501 - Return 501 status code - should be represented in the client as an error
// options - HTTPServerFailureClientHead501Options contains the optional parameters for the HTTPServerFailureClient.Head501
// method.
func (client *HTTPServerFailureClient) Head501(ctx context.Context, options *HTTPServerFailureClientHead501Options) (HTTPServerFailureClientHead501Response, error) {
	req, err := client.head501CreateRequest(ctx, options)
	if err != nil {
		return HTTPServerFailureClientHead501Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPServerFailureClientHead501Response{}, err
	}
	result := HTTPServerFailureClientHead501Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head501CreateRequest creates the Head501 request.
func (client *HTTPServerFailureClient) head501CreateRequest(ctx context.Context, options *HTTPServerFailureClientHead501Options) (*policy.Request, error) {
	urlPath := "/http/failure/server/501"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Post505 - Return 505 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPServerFailureClientPost505Options contains the optional parameters for the HTTPServerFailureClient.Post505
// method.
func (client *HTTPServerFailureClient) Post505(ctx context.Context, options *HTTPServerFailureClientPost505Options) (HTTPServerFailureClientPost505Response, error) {
	req, err := client.post505CreateRequest(ctx, options)
	if err != nil {
		return HTTPServerFailureClientPost505Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPServerFailureClientPost505Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPServerFailureClientPost505Response{}, runtime.NewResponseError(resp)
	}
	return HTTPServerFailureClientPost505Response{RawResponse: resp}, nil
}

// post505CreateRequest creates the Post505 request.
func (client *HTTPServerFailureClient) post505CreateRequest(ctx context.Context, options *HTTPServerFailureClientPost505Options) (*policy.Request, error) {
	urlPath := "/http/failure/server/505"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}
