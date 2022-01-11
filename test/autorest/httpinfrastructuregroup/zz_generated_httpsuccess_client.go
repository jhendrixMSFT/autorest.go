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

// HTTPSuccessClient contains the methods for the HTTPSuccess group.
// Don't use this type directly, use NewHTTPSuccessClient() instead.
type HTTPSuccessClient struct {
	pl runtime.Pipeline
}

// NewHTTPSuccessClient creates a new instance of HTTPSuccessClient with the specified values.
// options - pass nil to accept the default values.
func NewHTTPSuccessClient(options *azcore.ClientOptions) *HTTPSuccessClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &HTTPSuccessClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Delete200 - Delete simple boolean value true returns 200
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientDelete200Options contains the optional parameters for the HTTPSuccessClient.Delete200 method.
func (client *HTTPSuccessClient) Delete200(ctx context.Context, options *HTTPSuccessClientDelete200Options) (HTTPSuccessClientDelete200Response, error) {
	req, err := client.delete200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientDelete200Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientDelete200Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPSuccessClientDelete200Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientDelete200Response{RawResponse: resp}, nil
}

// delete200CreateRequest creates the Delete200 request.
func (client *HTTPSuccessClient) delete200CreateRequest(ctx context.Context, options *HTTPSuccessClientDelete200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Delete202 - Delete true Boolean value in request returns 202 (accepted)
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientDelete202Options contains the optional parameters for the HTTPSuccessClient.Delete202 method.
func (client *HTTPSuccessClient) Delete202(ctx context.Context, options *HTTPSuccessClientDelete202Options) (HTTPSuccessClientDelete202Response, error) {
	req, err := client.delete202CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientDelete202Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientDelete202Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return HTTPSuccessClientDelete202Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientDelete202Response{RawResponse: resp}, nil
}

// delete202CreateRequest creates the Delete202 request.
func (client *HTTPSuccessClient) delete202CreateRequest(ctx context.Context, options *HTTPSuccessClientDelete202Options) (*policy.Request, error) {
	urlPath := "/http/success/202"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Delete204 - Delete true Boolean value in request returns 204 (no content)
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientDelete204Options contains the optional parameters for the HTTPSuccessClient.Delete204 method.
func (client *HTTPSuccessClient) Delete204(ctx context.Context, options *HTTPSuccessClientDelete204Options) (HTTPSuccessClientDelete204Response, error) {
	req, err := client.delete204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientDelete204Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientDelete204Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return HTTPSuccessClientDelete204Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientDelete204Response{RawResponse: resp}, nil
}

// delete204CreateRequest creates the Delete204 request.
func (client *HTTPSuccessClient) delete204CreateRequest(ctx context.Context, options *HTTPSuccessClientDelete204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Get200 - Get 200 success
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientGet200Options contains the optional parameters for the HTTPSuccessClient.Get200 method.
func (client *HTTPSuccessClient) Get200(ctx context.Context, options *HTTPSuccessClientGet200Options) (HTTPSuccessClientGet200Response, error) {
	req, err := client.get200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientGet200Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientGet200Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPSuccessClientGet200Response{}, runtime.NewResponseError(resp)
	}
	return client.get200HandleResponse(resp)
}

// get200CreateRequest creates the Get200 request.
func (client *HTTPSuccessClient) get200CreateRequest(ctx context.Context, options *HTTPSuccessClientGet200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// get200HandleResponse handles the Get200 response.
func (client *HTTPSuccessClient) get200HandleResponse(resp *http.Response) (HTTPSuccessClientGet200Response, error) {
	result := HTTPSuccessClientGet200Response{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPSuccessClientGet200Response{}, err
	}
	return result, nil
}

// Head200 - Return 200 status code if successful
// options - HTTPSuccessClientHead200Options contains the optional parameters for the HTTPSuccessClient.Head200 method.
func (client *HTTPSuccessClient) Head200(ctx context.Context, options *HTTPSuccessClientHead200Options) (HTTPSuccessClientHead200Response, error) {
	req, err := client.head200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead200Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientHead200Response{}, err
	}
	result := HTTPSuccessClientHead200Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head200CreateRequest creates the Head200 request.
func (client *HTTPSuccessClient) head200CreateRequest(ctx context.Context, options *HTTPSuccessClientHead200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Head204 - Return 204 status code if successful
// options - HTTPSuccessClientHead204Options contains the optional parameters for the HTTPSuccessClient.Head204 method.
func (client *HTTPSuccessClient) Head204(ctx context.Context, options *HTTPSuccessClientHead204Options) (HTTPSuccessClientHead204Response, error) {
	req, err := client.head204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead204Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientHead204Response{}, err
	}
	result := HTTPSuccessClientHead204Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head204CreateRequest creates the Head204 request.
func (client *HTTPSuccessClient) head204CreateRequest(ctx context.Context, options *HTTPSuccessClientHead204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Head404 - Return 404 status code
// options - HTTPSuccessClientHead404Options contains the optional parameters for the HTTPSuccessClient.Head404 method.
func (client *HTTPSuccessClient) Head404(ctx context.Context, options *HTTPSuccessClientHead404Options) (HTTPSuccessClientHead404Response, error) {
	req, err := client.head404CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead404Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientHead404Response{}, err
	}
	result := HTTPSuccessClientHead404Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head404CreateRequest creates the Head404 request.
func (client *HTTPSuccessClient) head404CreateRequest(ctx context.Context, options *HTTPSuccessClientHead404Options) (*policy.Request, error) {
	urlPath := "/http/success/404"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Options200 - Options 200 success
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientOptions200Options contains the optional parameters for the HTTPSuccessClient.Options200 method.
func (client *HTTPSuccessClient) Options200(ctx context.Context, options *HTTPSuccessClientOptions200Options) (HTTPSuccessClientOptions200Response, error) {
	req, err := client.options200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientOptions200Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientOptions200Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPSuccessClientOptions200Response{}, runtime.NewResponseError(resp)
	}
	return client.options200HandleResponse(resp)
}

// options200CreateRequest creates the Options200 request.
func (client *HTTPSuccessClient) options200CreateRequest(ctx context.Context, options *HTTPSuccessClientOptions200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// options200HandleResponse handles the Options200 response.
func (client *HTTPSuccessClient) options200HandleResponse(resp *http.Response) (HTTPSuccessClientOptions200Response, error) {
	result := HTTPSuccessClientOptions200Response{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPSuccessClientOptions200Response{}, err
	}
	return result, nil
}

// Patch200 - Patch true Boolean value in request returning 200
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPatch200Options contains the optional parameters for the HTTPSuccessClient.Patch200 method.
func (client *HTTPSuccessClient) Patch200(ctx context.Context, options *HTTPSuccessClientPatch200Options) (HTTPSuccessClientPatch200Response, error) {
	req, err := client.patch200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPatch200Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPatch200Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPSuccessClientPatch200Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPatch200Response{RawResponse: resp}, nil
}

// patch200CreateRequest creates the Patch200 request.
func (client *HTTPSuccessClient) patch200CreateRequest(ctx context.Context, options *HTTPSuccessClientPatch200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Patch202 - Patch true Boolean value in request returns 202
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPatch202Options contains the optional parameters for the HTTPSuccessClient.Patch202 method.
func (client *HTTPSuccessClient) Patch202(ctx context.Context, options *HTTPSuccessClientPatch202Options) (HTTPSuccessClientPatch202Response, error) {
	req, err := client.patch202CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPatch202Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPatch202Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return HTTPSuccessClientPatch202Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPatch202Response{RawResponse: resp}, nil
}

// patch202CreateRequest creates the Patch202 request.
func (client *HTTPSuccessClient) patch202CreateRequest(ctx context.Context, options *HTTPSuccessClientPatch202Options) (*policy.Request, error) {
	urlPath := "/http/success/202"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Patch204 - Patch true Boolean value in request returns 204 (no content)
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPatch204Options contains the optional parameters for the HTTPSuccessClient.Patch204 method.
func (client *HTTPSuccessClient) Patch204(ctx context.Context, options *HTTPSuccessClientPatch204Options) (HTTPSuccessClientPatch204Response, error) {
	req, err := client.patch204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPatch204Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPatch204Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return HTTPSuccessClientPatch204Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPatch204Response{RawResponse: resp}, nil
}

// patch204CreateRequest creates the Patch204 request.
func (client *HTTPSuccessClient) patch204CreateRequest(ctx context.Context, options *HTTPSuccessClientPatch204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Post200 - Post bollean value true in request that returns a 200
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPost200Options contains the optional parameters for the HTTPSuccessClient.Post200 method.
func (client *HTTPSuccessClient) Post200(ctx context.Context, options *HTTPSuccessClientPost200Options) (HTTPSuccessClientPost200Response, error) {
	req, err := client.post200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPost200Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPost200Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPSuccessClientPost200Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPost200Response{RawResponse: resp}, nil
}

// post200CreateRequest creates the Post200 request.
func (client *HTTPSuccessClient) post200CreateRequest(ctx context.Context, options *HTTPSuccessClientPost200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Post201 - Post true Boolean value in request returns 201 (Created)
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPost201Options contains the optional parameters for the HTTPSuccessClient.Post201 method.
func (client *HTTPSuccessClient) Post201(ctx context.Context, options *HTTPSuccessClientPost201Options) (HTTPSuccessClientPost201Response, error) {
	req, err := client.post201CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPost201Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPost201Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return HTTPSuccessClientPost201Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPost201Response{RawResponse: resp}, nil
}

// post201CreateRequest creates the Post201 request.
func (client *HTTPSuccessClient) post201CreateRequest(ctx context.Context, options *HTTPSuccessClientPost201Options) (*policy.Request, error) {
	urlPath := "/http/success/201"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Post202 - Post true Boolean value in request returns 202 (Accepted)
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPost202Options contains the optional parameters for the HTTPSuccessClient.Post202 method.
func (client *HTTPSuccessClient) Post202(ctx context.Context, options *HTTPSuccessClientPost202Options) (HTTPSuccessClientPost202Response, error) {
	req, err := client.post202CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPost202Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPost202Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return HTTPSuccessClientPost202Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPost202Response{RawResponse: resp}, nil
}

// post202CreateRequest creates the Post202 request.
func (client *HTTPSuccessClient) post202CreateRequest(ctx context.Context, options *HTTPSuccessClientPost202Options) (*policy.Request, error) {
	urlPath := "/http/success/202"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Post204 - Post true Boolean value in request returns 204 (no content)
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPost204Options contains the optional parameters for the HTTPSuccessClient.Post204 method.
func (client *HTTPSuccessClient) Post204(ctx context.Context, options *HTTPSuccessClientPost204Options) (HTTPSuccessClientPost204Response, error) {
	req, err := client.post204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPost204Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPost204Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return HTTPSuccessClientPost204Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPost204Response{RawResponse: resp}, nil
}

// post204CreateRequest creates the Post204 request.
func (client *HTTPSuccessClient) post204CreateRequest(ctx context.Context, options *HTTPSuccessClientPost204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Put200 - Put boolean value true returning 200 success
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPut200Options contains the optional parameters for the HTTPSuccessClient.Put200 method.
func (client *HTTPSuccessClient) Put200(ctx context.Context, options *HTTPSuccessClientPut200Options) (HTTPSuccessClientPut200Response, error) {
	req, err := client.put200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPut200Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPut200Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPSuccessClientPut200Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPut200Response{RawResponse: resp}, nil
}

// put200CreateRequest creates the Put200 request.
func (client *HTTPSuccessClient) put200CreateRequest(ctx context.Context, options *HTTPSuccessClientPut200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Put201 - Put true Boolean value in request returns 201
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPut201Options contains the optional parameters for the HTTPSuccessClient.Put201 method.
func (client *HTTPSuccessClient) Put201(ctx context.Context, options *HTTPSuccessClientPut201Options) (HTTPSuccessClientPut201Response, error) {
	req, err := client.put201CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPut201Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPut201Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return HTTPSuccessClientPut201Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPut201Response{RawResponse: resp}, nil
}

// put201CreateRequest creates the Put201 request.
func (client *HTTPSuccessClient) put201CreateRequest(ctx context.Context, options *HTTPSuccessClientPut201Options) (*policy.Request, error) {
	urlPath := "/http/success/201"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Put202 - Put true Boolean value in request returns 202 (Accepted)
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPut202Options contains the optional parameters for the HTTPSuccessClient.Put202 method.
func (client *HTTPSuccessClient) Put202(ctx context.Context, options *HTTPSuccessClientPut202Options) (HTTPSuccessClientPut202Response, error) {
	req, err := client.put202CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPut202Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPut202Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return HTTPSuccessClientPut202Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPut202Response{RawResponse: resp}, nil
}

// put202CreateRequest creates the Put202 request.
func (client *HTTPSuccessClient) put202CreateRequest(ctx context.Context, options *HTTPSuccessClientPut202Options) (*policy.Request, error) {
	urlPath := "/http/success/202"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// Put204 - Put true Boolean value in request returns 204 (no content)
// If the operation fails it returns an *azcore.ResponseError type.
// options - HTTPSuccessClientPut204Options contains the optional parameters for the HTTPSuccessClient.Put204 method.
func (client *HTTPSuccessClient) Put204(ctx context.Context, options *HTTPSuccessClientPut204Options) (HTTPSuccessClientPut204Response, error) {
	req, err := client.put204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPut204Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientPut204Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return HTTPSuccessClientPut204Response{}, runtime.NewResponseError(resp)
	}
	return HTTPSuccessClientPut204Response{RawResponse: resp}, nil
}

// put204CreateRequest creates the Put204 request.
func (client *HTTPSuccessClient) put204CreateRequest(ctx context.Context, options *HTTPSuccessClientPut204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}
