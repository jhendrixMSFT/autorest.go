//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package stringgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"net/http"
)

// EnumClient contains the methods for the Enum group.
// Don't use this type directly, use a constructor function instead.
type EnumClient struct {
	internal *azcore.Client
}

// GetNotExpandable - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - EnumClientGetNotExpandableOptions contains the optional parameters for the EnumClient.GetNotExpandable method.
func (client *EnumClient) GetNotExpandable(ctx context.Context, options *EnumClientGetNotExpandableOptions) (result EnumClientGetNotExpandableResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "EnumClient.GetNotExpandable", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getNotExpandableCreateRequest(ctx, options)
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
	result, err = client.getNotExpandableHandleResponse(resp)
	return
}

// getNotExpandableCreateRequest creates the GetNotExpandable request.
func (client *EnumClient) getNotExpandableCreateRequest(ctx context.Context, options *EnumClientGetNotExpandableOptions) (*policy.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNotExpandableHandleResponse handles the GetNotExpandable response.
func (client *EnumClient) getNotExpandableHandleResponse(resp *http.Response) (result EnumClientGetNotExpandableResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		result = EnumClientGetNotExpandableResponse{}
		return
	}
	return result, nil
}

// GetReferenced - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - EnumClientGetReferencedOptions contains the optional parameters for the EnumClient.GetReferenced method.
func (client *EnumClient) GetReferenced(ctx context.Context, options *EnumClientGetReferencedOptions) (result EnumClientGetReferencedResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "EnumClient.GetReferenced", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getReferencedCreateRequest(ctx, options)
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
	result, err = client.getReferencedHandleResponse(resp)
	return
}

// getReferencedCreateRequest creates the GetReferenced request.
func (client *EnumClient) getReferencedCreateRequest(ctx context.Context, options *EnumClientGetReferencedOptions) (*policy.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getReferencedHandleResponse handles the GetReferenced response.
func (client *EnumClient) getReferencedHandleResponse(resp *http.Response) (result EnumClientGetReferencedResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		result = EnumClientGetReferencedResponse{}
		return
	}
	return result, nil
}

// GetReferencedConstant - Get value 'green-color' from the constant.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - EnumClientGetReferencedConstantOptions contains the optional parameters for the EnumClient.GetReferencedConstant
//     method.
func (client *EnumClient) GetReferencedConstant(ctx context.Context, options *EnumClientGetReferencedConstantOptions) (result EnumClientGetReferencedConstantResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "EnumClient.GetReferencedConstant", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getReferencedConstantCreateRequest(ctx, options)
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
	result, err = client.getReferencedConstantHandleResponse(resp)
	return
}

// getReferencedConstantCreateRequest creates the GetReferencedConstant request.
func (client *EnumClient) getReferencedConstantCreateRequest(ctx context.Context, options *EnumClientGetReferencedConstantOptions) (*policy.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getReferencedConstantHandleResponse handles the GetReferencedConstant response.
func (client *EnumClient) getReferencedConstantHandleResponse(resp *http.Response) (result EnumClientGetReferencedConstantResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.RefColorConstant); err != nil {
		result = EnumClientGetReferencedConstantResponse{}
		return
	}
	return result, nil
}

// PutNotExpandable - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - stringBody - string body
//   - options - EnumClientPutNotExpandableOptions contains the optional parameters for the EnumClient.PutNotExpandable method.
func (client *EnumClient) PutNotExpandable(ctx context.Context, stringBody Colors, options *EnumClientPutNotExpandableOptions) (result EnumClientPutNotExpandableResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "EnumClient.PutNotExpandable", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.putNotExpandableCreateRequest(ctx, stringBody, options)
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

// putNotExpandableCreateRequest creates the PutNotExpandable request.
func (client *EnumClient) putNotExpandableCreateRequest(ctx context.Context, stringBody Colors, options *EnumClientPutNotExpandableOptions) (*policy.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, stringBody)
}

// PutReferenced - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - enumStringBody - enum string body
//   - options - EnumClientPutReferencedOptions contains the optional parameters for the EnumClient.PutReferenced method.
func (client *EnumClient) PutReferenced(ctx context.Context, enumStringBody Colors, options *EnumClientPutReferencedOptions) (result EnumClientPutReferencedResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "EnumClient.PutReferenced", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.putReferencedCreateRequest(ctx, enumStringBody, options)
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

// putReferencedCreateRequest creates the PutReferenced request.
func (client *EnumClient) putReferencedCreateRequest(ctx context.Context, enumStringBody Colors, options *EnumClientPutReferencedOptions) (*policy.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, enumStringBody)
}

// PutReferencedConstant - Sends value 'green-color' from a constant
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - enumStringBody - enum string body
//   - options - EnumClientPutReferencedConstantOptions contains the optional parameters for the EnumClient.PutReferencedConstant
//     method.
func (client *EnumClient) PutReferencedConstant(ctx context.Context, enumStringBody RefColorConstant, options *EnumClientPutReferencedConstantOptions) (result EnumClientPutReferencedConstantResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "EnumClient.PutReferencedConstant", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.putReferencedConstantCreateRequest(ctx, enumStringBody, options)
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

// putReferencedConstantCreateRequest creates the PutReferencedConstant request.
func (client *EnumClient) putReferencedConstantCreateRequest(ctx context.Context, enumStringBody RefColorConstant, options *EnumClientPutReferencedConstantOptions) (*policy.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, enumStringBody)
}
