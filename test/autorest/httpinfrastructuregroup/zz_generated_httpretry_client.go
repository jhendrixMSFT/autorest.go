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

// HTTPRetryClient contains the methods for the HTTPRetry group.
// Don't use this type directly, use NewHTTPRetryClient() instead.
type HTTPRetryClient struct {
	pl runtime.Pipeline
}

// NewHTTPRetryClient creates a new instance of HTTPRetryClient with the specified values.
// options - pass nil to accept the default values.
func NewHTTPRetryClient(options *azcore.ClientOptions) *HTTPRetryClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &HTTPRetryClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// Delete503 - Return 503 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPRetryClientDelete503Options contains the optional parameters for the HTTPRetryClient.Delete503 method.
func (client *HTTPRetryClient) Delete503(ctx context.Context, options *HTTPRetryClientDelete503Options) (HTTPRetryClientDelete503Response, error) {
	req, err := client.delete503CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientDelete503Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRetryClientDelete503Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryClientDelete503Response{}, runtime.NewResponseError(resp)
	}
	return HTTPRetryClientDelete503Response{RawResponse: resp}, nil
}

// delete503CreateRequest creates the Delete503 request.
func (client *HTTPRetryClient) delete503CreateRequest(ctx context.Context, options *HTTPRetryClientDelete503Options) (*policy.Request, error) {
	urlPath := "/http/retry/503"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Get502 - Return 502 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPRetryClientGet502Options contains the optional parameters for the HTTPRetryClient.Get502 method.
func (client *HTTPRetryClient) Get502(ctx context.Context, options *HTTPRetryClientGet502Options) (HTTPRetryClientGet502Response, error) {
	req, err := client.get502CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientGet502Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRetryClientGet502Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryClientGet502Response{}, runtime.NewResponseError(resp)
	}
	return HTTPRetryClientGet502Response{RawResponse: resp}, nil
}

// get502CreateRequest creates the Get502 request.
func (client *HTTPRetryClient) get502CreateRequest(ctx context.Context, options *HTTPRetryClientGet502Options) (*policy.Request, error) {
	urlPath := "/http/retry/502"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Head408 - Return 408 status code, then 200 after retry
// options - HTTPRetryClientHead408Options contains the optional parameters for the HTTPRetryClient.Head408 method.
func (client *HTTPRetryClient) Head408(ctx context.Context, options *HTTPRetryClientHead408Options) (HTTPRetryClientHead408Response, error) {
	req, err := client.head408CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientHead408Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRetryClientHead408Response{}, err
	}
	result := HTTPRetryClientHead408Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head408CreateRequest creates the Head408 request.
func (client *HTTPRetryClient) head408CreateRequest(ctx context.Context, options *HTTPRetryClientHead408Options) (*policy.Request, error) {
	urlPath := "/http/retry/408"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Options502 - Return 502 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPRetryClientOptions502Options contains the optional parameters for the HTTPRetryClient.Options502 method.
func (client *HTTPRetryClient) Options502(ctx context.Context, options *HTTPRetryClientOptions502Options) (HTTPRetryClientOptions502Response, error) {
	req, err := client.options502CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientOptions502Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRetryClientOptions502Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryClientOptions502Response{}, runtime.NewResponseError(resp)
	}
	return client.options502HandleResponse(resp)
}

// options502CreateRequest creates the Options502 request.
func (client *HTTPRetryClient) options502CreateRequest(ctx context.Context, options *HTTPRetryClientOptions502Options) (*policy.Request, error) {
	urlPath := "/http/retry/502"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// options502HandleResponse handles the Options502 response.
func (client *HTTPRetryClient) options502HandleResponse(resp *http.Response) (HTTPRetryClientOptions502Response, error) {
	result := HTTPRetryClientOptions502Response{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPRetryClientOptions502Response{}, err
	}
	return result, nil
}

// Patch500 - Return 500 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPRetryClientPatch500Options contains the optional parameters for the HTTPRetryClient.Patch500 method.
func (client *HTTPRetryClient) Patch500(ctx context.Context, options *HTTPRetryClientPatch500Options) (HTTPRetryClientPatch500Response, error) {
	req, err := client.patch500CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPatch500Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRetryClientPatch500Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryClientPatch500Response{}, runtime.NewResponseError(resp)
	}
	return HTTPRetryClientPatch500Response{RawResponse: resp}, nil
}

// patch500CreateRequest creates the Patch500 request.
func (client *HTTPRetryClient) patch500CreateRequest(ctx context.Context, options *HTTPRetryClientPatch500Options) (*policy.Request, error) {
	urlPath := "/http/retry/500"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Patch504 - Return 504 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPRetryClientPatch504Options contains the optional parameters for the HTTPRetryClient.Patch504 method.
func (client *HTTPRetryClient) Patch504(ctx context.Context, options *HTTPRetryClientPatch504Options) (HTTPRetryClientPatch504Response, error) {
	req, err := client.patch504CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPatch504Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRetryClientPatch504Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryClientPatch504Response{}, runtime.NewResponseError(resp)
	}
	return HTTPRetryClientPatch504Response{RawResponse: resp}, nil
}

// patch504CreateRequest creates the Patch504 request.
func (client *HTTPRetryClient) patch504CreateRequest(ctx context.Context, options *HTTPRetryClientPatch504Options) (*policy.Request, error) {
	urlPath := "/http/retry/504"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Post503 - Return 503 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPRetryClientPost503Options contains the optional parameters for the HTTPRetryClient.Post503 method.
func (client *HTTPRetryClient) Post503(ctx context.Context, options *HTTPRetryClientPost503Options) (HTTPRetryClientPost503Response, error) {
	req, err := client.post503CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPost503Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRetryClientPost503Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryClientPost503Response{}, runtime.NewResponseError(resp)
	}
	return HTTPRetryClientPost503Response{RawResponse: resp}, nil
}

// post503CreateRequest creates the Post503 request.
func (client *HTTPRetryClient) post503CreateRequest(ctx context.Context, options *HTTPRetryClientPost503Options) (*policy.Request, error) {
	urlPath := "/http/retry/503"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Put500 - Return 500 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPRetryClientPut500Options contains the optional parameters for the HTTPRetryClient.Put500 method.
func (client *HTTPRetryClient) Put500(ctx context.Context, options *HTTPRetryClientPut500Options) (HTTPRetryClientPut500Response, error) {
	req, err := client.put500CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPut500Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRetryClientPut500Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryClientPut500Response{}, runtime.NewResponseError(resp)
	}
	return HTTPRetryClientPut500Response{RawResponse: resp}, nil
}

// put500CreateRequest creates the Put500 request.
func (client *HTTPRetryClient) put500CreateRequest(ctx context.Context, options *HTTPRetryClientPut500Options) (*policy.Request, error) {
	urlPath := "/http/retry/500"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Put504 - Return 504 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPRetryClientPut504Options contains the optional parameters for the HTTPRetryClient.Put504 method.
func (client *HTTPRetryClient) Put504(ctx context.Context, options *HTTPRetryClientPut504Options) (HTTPRetryClientPut504Response, error) {
	req, err := client.put504CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPut504Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRetryClientPut504Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryClientPut504Response{}, runtime.NewResponseError(resp)
	}
	return HTTPRetryClientPut504Response{RawResponse: resp}, nil
}

// put504CreateRequest creates the Put504 request.
func (client *HTTPRetryClient) put504CreateRequest(ctx context.Context, options *HTTPRetryClientPut504Options) (*policy.Request, error) {
	urlPath := "/http/retry/504"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}
