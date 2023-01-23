//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package lrogroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// LROsCustomHeaderClient contains the methods for the LROsCustomHeader group.
// Don't use this type directly, use a constructor function instead.
type LROsCustomHeaderClient struct {
	internal *azcore.Client
}

// BeginPost202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all
// requests. Long running post request, service returns a 202 to the initial request, with 'Location' and
// 'Retry-After' headers, Polls return a 200 with a response body after success
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - LROsCustomHeaderClientBeginPost202Retry200Options contains the optional parameters for the LROsCustomHeaderClient.BeginPost202Retry200
//     method.
func (client *LROsCustomHeaderClient) BeginPost202Retry200(ctx context.Context, options *LROsCustomHeaderClientBeginPost202Retry200Options) (result *runtime.Poller[LROsCustomHeaderClientPost202Retry200Response], err error) {
	if options == nil || options.ResumeToken == "" {
		ctx, endSpan := runtime.StartSpan(ctx, "LROsCustomHeaderClient.BeginPost202Retry200", client.internal.Tracer(), nil)
		defer func() { endSpan(err) }()
		var resp *http.Response
		resp, err = client.post202Retry200(ctx, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LROsCustomHeaderClientPost202Retry200Response]{
			Tracer: client.internal.Tracer(),
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LROsCustomHeaderClientPost202Retry200Response]{
			Tracer: client.internal.Tracer(),
		})
	}
	return
}

// Post202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests.
// Long running post request, service returns a 202 to the initial request, with 'Location' and
// 'Retry-After' headers, Polls return a 200 with a response body after success
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
func (client *LROsCustomHeaderClient) post202Retry200(ctx context.Context, options *LROsCustomHeaderClientBeginPost202Retry200Options) (resp *http.Response, err error) {
	req, err := client.post202Retry200CreateRequest(ctx, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *LROsCustomHeaderClient) post202Retry200CreateRequest(ctx context.Context, options *LROsCustomHeaderClientBeginPost202Retry200Options) (*policy.Request, error) {
	urlPath := "/lro/customheader/post/202/retry/200"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Product != nil {
		return req, runtime.MarshalAsJSON(req, *options.Product)
	}
	return req, nil
}

// BeginPostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header
// for all requests. Long running post request, service returns a 202 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions contains the optional parameters for the LROsCustomHeaderClient.BeginPostAsyncRetrySucceeded
//     method.
func (client *LROsCustomHeaderClient) BeginPostAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions) (result *runtime.Poller[LROsCustomHeaderClientPostAsyncRetrySucceededResponse], err error) {
	if options == nil || options.ResumeToken == "" {
		ctx, endSpan := runtime.StartSpan(ctx, "LROsCustomHeaderClient.BeginPostAsyncRetrySucceeded", client.internal.Tracer(), nil)
		defer func() { endSpan(err) }()
		var resp *http.Response
		resp, err = client.postAsyncRetrySucceeded(ctx, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LROsCustomHeaderClientPostAsyncRetrySucceededResponse]{
			Tracer: client.internal.Tracer(),
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LROsCustomHeaderClientPostAsyncRetrySucceededResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
	return
}

// PostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for
// all requests. Long running post request, service returns a 202 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
func (client *LROsCustomHeaderClient) postAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions) (resp *http.Response, err error) {
	req, err := client.postAsyncRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// postAsyncRetrySucceededCreateRequest creates the PostAsyncRetrySucceeded request.
func (client *LROsCustomHeaderClient) postAsyncRetrySucceededCreateRequest(ctx context.Context, options *LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/customheader/postasync/retry/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Product != nil {
		return req, runtime.MarshalAsJSON(req, *options.Product)
	}
	return req, nil
}

// BeginPut201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header
// for all requests. Long running put request, service returns a 201 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - product - Product to put
//   - options - LROsCustomHeaderClientBeginPut201CreatingSucceeded200Options contains the optional parameters for the LROsCustomHeaderClient.BeginPut201CreatingSucceeded200
//     method.
func (client *LROsCustomHeaderClient) BeginPut201CreatingSucceeded200(ctx context.Context, product Product, options *LROsCustomHeaderClientBeginPut201CreatingSucceeded200Options) (result *runtime.Poller[LROsCustomHeaderClientPut201CreatingSucceeded200Response], err error) {
	if options == nil || options.ResumeToken == "" {
		ctx, endSpan := runtime.StartSpan(ctx, "LROsCustomHeaderClient.BeginPut201CreatingSucceeded200", client.internal.Tracer(), nil)
		defer func() { endSpan(err) }()
		var resp *http.Response
		resp, err = client.put201CreatingSucceeded200(ctx, product, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LROsCustomHeaderClientPut201CreatingSucceeded200Response]{
			Tracer: client.internal.Tracer(),
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LROsCustomHeaderClientPut201CreatingSucceeded200Response]{
			Tracer: client.internal.Tracer(),
		})
	}
	return
}

// Put201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for
// all requests. Long running put request, service returns a 201 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
func (client *LROsCustomHeaderClient) put201CreatingSucceeded200(ctx context.Context, product Product, options *LROsCustomHeaderClientBeginPut201CreatingSucceeded200Options) (resp *http.Response, err error) {
	req, err := client.put201CreatingSucceeded200CreateRequest(ctx, product, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *LROsCustomHeaderClient) put201CreatingSucceeded200CreateRequest(ctx context.Context, product Product, options *LROsCustomHeaderClientBeginPut201CreatingSucceeded200Options) (*policy.Request, error) {
	urlPath := "/lro/customheader/put/201/creating/succeeded/200"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, product)
}

// BeginPutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header
// for all requests. Long running put request, service returns a 200 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - product - Product to put
//   - options - LROsCustomHeaderClientBeginPutAsyncRetrySucceededOptions contains the optional parameters for the LROsCustomHeaderClient.BeginPutAsyncRetrySucceeded
//     method.
func (client *LROsCustomHeaderClient) BeginPutAsyncRetrySucceeded(ctx context.Context, product Product, options *LROsCustomHeaderClientBeginPutAsyncRetrySucceededOptions) (result *runtime.Poller[LROsCustomHeaderClientPutAsyncRetrySucceededResponse], err error) {
	if options == nil || options.ResumeToken == "" {
		ctx, endSpan := runtime.StartSpan(ctx, "LROsCustomHeaderClient.BeginPutAsyncRetrySucceeded", client.internal.Tracer(), nil)
		defer func() { endSpan(err) }()
		var resp *http.Response
		resp, err = client.putAsyncRetrySucceeded(ctx, product, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LROsCustomHeaderClientPutAsyncRetrySucceededResponse]{
			Tracer: client.internal.Tracer(),
		})
	} else {
		result, err = runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LROsCustomHeaderClientPutAsyncRetrySucceededResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
	return
}

// PutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all
// requests. Long running put request, service returns a 200 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
func (client *LROsCustomHeaderClient) putAsyncRetrySucceeded(ctx context.Context, product Product, options *LROsCustomHeaderClientBeginPutAsyncRetrySucceededOptions) (resp *http.Response, err error) {
	req, err := client.putAsyncRetrySucceededCreateRequest(ctx, product, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// putAsyncRetrySucceededCreateRequest creates the PutAsyncRetrySucceeded request.
func (client *LROsCustomHeaderClient) putAsyncRetrySucceededCreateRequest(ctx context.Context, product Product, options *LROsCustomHeaderClientBeginPutAsyncRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/customheader/putasync/retry/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, product)
}
