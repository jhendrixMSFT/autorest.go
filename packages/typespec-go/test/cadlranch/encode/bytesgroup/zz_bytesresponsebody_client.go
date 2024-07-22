// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bytesgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BytesResponseBodyClient contains the methods for the BytesResponseBody group.
// Don't use this type directly, use [BytesClient.NewBytesResponseBodyClient] instead.
type BytesResponseBodyClient struct {
	internal *azcore.Client
}

// Base64 -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - BytesResponseBodyClientBase64Options contains the optional parameters for the BytesResponseBodyClient.Base64
//     method.
func (client *BytesResponseBodyClient) Base64(ctx context.Context, options *BytesResponseBodyClientBase64Options) (BytesResponseBodyClientBase64Response, error) {
	var err error
	const operationName = "BytesResponseBodyClient.Base64"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.base64CreateRequest(ctx, options)
	if err != nil {
		return BytesResponseBodyClientBase64Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BytesResponseBodyClientBase64Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BytesResponseBodyClientBase64Response{}, err
	}
	resp, err := client.base64HandleResponse(httpResp)
	return resp, err
}

// base64CreateRequest creates the Base64 request.
func (client *BytesResponseBodyClient) base64CreateRequest(ctx context.Context, _ *BytesResponseBodyClientBase64Options) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/base64"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// base64HandleResponse handles the Base64 response.
func (client *BytesResponseBodyClient) base64HandleResponse(resp *http.Response) (BytesResponseBodyClientBase64Response, error) {
	result := BytesResponseBodyClientBase64Response{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return BytesResponseBodyClientBase64Response{}, err
	}
	return result, nil
}

// Base64URL -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - BytesResponseBodyClientBase64URLOptions contains the optional parameters for the BytesResponseBodyClient.Base64URL
//     method.
func (client *BytesResponseBodyClient) Base64URL(ctx context.Context, options *BytesResponseBodyClientBase64URLOptions) (BytesResponseBodyClientBase64URLResponse, error) {
	var err error
	const operationName = "BytesResponseBodyClient.Base64URL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.base64URLCreateRequest(ctx, options)
	if err != nil {
		return BytesResponseBodyClientBase64URLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BytesResponseBodyClientBase64URLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BytesResponseBodyClientBase64URLResponse{}, err
	}
	resp, err := client.base64URLHandleResponse(httpResp)
	return resp, err
}

// base64URLCreateRequest creates the Base64URL request.
func (client *BytesResponseBodyClient) base64URLCreateRequest(ctx context.Context, _ *BytesResponseBodyClientBase64URLOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/base64url"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// base64URLHandleResponse handles the Base64URL response.
func (client *BytesResponseBodyClient) base64URLHandleResponse(resp *http.Response) (BytesResponseBodyClientBase64URLResponse, error) {
	result := BytesResponseBodyClientBase64URLResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64URLFormat); err != nil {
		return BytesResponseBodyClientBase64URLResponse{}, err
	}
	return result, nil
}

// CustomContentType -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - BytesResponseBodyClientCustomContentTypeOptions contains the optional parameters for the BytesResponseBodyClient.CustomContentType
//     method.
func (client *BytesResponseBodyClient) CustomContentType(ctx context.Context, options *BytesResponseBodyClientCustomContentTypeOptions) (BytesResponseBodyClientCustomContentTypeResponse, error) {
	var err error
	const operationName = "BytesResponseBodyClient.CustomContentType"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.customContentTypeCreateRequest(ctx, options)
	if err != nil {
		return BytesResponseBodyClientCustomContentTypeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BytesResponseBodyClientCustomContentTypeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BytesResponseBodyClientCustomContentTypeResponse{}, err
	}
	resp, err := client.customContentTypeHandleResponse(httpResp)
	return resp, err
}

// customContentTypeCreateRequest creates the CustomContentType request.
func (client *BytesResponseBodyClient) customContentTypeCreateRequest(ctx context.Context, _ *BytesResponseBodyClientCustomContentTypeOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/custom-content-type"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Accept"] = []string{"image/png"}
	return req, nil
}

// customContentTypeHandleResponse handles the CustomContentType response.
func (client *BytesResponseBodyClient) customContentTypeHandleResponse(resp *http.Response) (BytesResponseBodyClientCustomContentTypeResponse, error) {
	result := BytesResponseBodyClientCustomContentTypeResponse{Body: resp.Body}
	if val := resp.Header.Get("content-type"); val != "" {
		result.ContentType = &val
	}
	return result, nil
}

// Default -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - BytesResponseBodyClientDefaultOptions contains the optional parameters for the BytesResponseBodyClient.Default
//     method.
func (client *BytesResponseBodyClient) Default(ctx context.Context, options *BytesResponseBodyClientDefaultOptions) (BytesResponseBodyClientDefaultResponse, error) {
	var err error
	const operationName = "BytesResponseBodyClient.Default"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.defaultCreateRequest(ctx, options)
	if err != nil {
		return BytesResponseBodyClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BytesResponseBodyClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BytesResponseBodyClientDefaultResponse{}, err
	}
	resp, err := client.defaultHandleResponse(httpResp)
	return resp, err
}

// defaultCreateRequest creates the Default request.
func (client *BytesResponseBodyClient) defaultCreateRequest(ctx context.Context, _ *BytesResponseBodyClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// defaultHandleResponse handles the Default response.
func (client *BytesResponseBodyClient) defaultHandleResponse(resp *http.Response) (BytesResponseBodyClientDefaultResponse, error) {
	result := BytesResponseBodyClientDefaultResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return BytesResponseBodyClientDefaultResponse{}, err
	}
	return result, nil
}

// OctetStream -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - BytesResponseBodyClientOctetStreamOptions contains the optional parameters for the BytesResponseBodyClient.OctetStream
//     method.
func (client *BytesResponseBodyClient) OctetStream(ctx context.Context, options *BytesResponseBodyClientOctetStreamOptions) (BytesResponseBodyClientOctetStreamResponse, error) {
	var err error
	const operationName = "BytesResponseBodyClient.OctetStream"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.octetStreamCreateRequest(ctx, options)
	if err != nil {
		return BytesResponseBodyClientOctetStreamResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BytesResponseBodyClientOctetStreamResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BytesResponseBodyClientOctetStreamResponse{}, err
	}
	resp, err := client.octetStreamHandleResponse(httpResp)
	return resp, err
}

// octetStreamCreateRequest creates the OctetStream request.
func (client *BytesResponseBodyClient) octetStreamCreateRequest(ctx context.Context, _ *BytesResponseBodyClientOctetStreamOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/octet-stream"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Accept"] = []string{"application/octet-stream"}
	return req, nil
}

// octetStreamHandleResponse handles the OctetStream response.
func (client *BytesResponseBodyClient) octetStreamHandleResponse(resp *http.Response) (BytesResponseBodyClientOctetStreamResponse, error) {
	result := BytesResponseBodyClientOctetStreamResponse{Body: resp.Body}
	if val := resp.Header.Get("content-type"); val != "" {
		result.ContentType = &val
	}
	return result, nil
}
