// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package contentneggroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ContentNegotiationDifferentBodyClient contains the methods for the ContentNegotiationDifferentBody group.
// Don't use this type directly, use [ContentNegotiationClient.NewContentNegotiationDifferentBodyClient] instead.
type ContentNegotiationDifferentBodyClient struct {
	internal *azcore.Client
}

// GetAvatarAsJSON -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ContentNegotiationDifferentBodyClientGetAvatarAsJSONOptions contains the optional parameters for the ContentNegotiationDifferentBodyClient.GetAvatarAsJSON
//     method.
func (client *ContentNegotiationDifferentBodyClient) GetAvatarAsJSON(ctx context.Context, options *ContentNegotiationDifferentBodyClientGetAvatarAsJSONOptions) (ContentNegotiationDifferentBodyClientGetAvatarAsJSONResponse, error) {
	var err error
	const operationName = "ContentNegotiationDifferentBodyClient.GetAvatarAsJSON"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAvatarAsJSONCreateRequest(ctx, options)
	if err != nil {
		return ContentNegotiationDifferentBodyClientGetAvatarAsJSONResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentNegotiationDifferentBodyClientGetAvatarAsJSONResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContentNegotiationDifferentBodyClientGetAvatarAsJSONResponse{}, err
	}
	resp, err := client.getAvatarAsJSONHandleResponse(httpResp)
	return resp, err
}

// getAvatarAsJSONCreateRequest creates the GetAvatarAsJSON request.
func (client *ContentNegotiationDifferentBodyClient) getAvatarAsJSONCreateRequest(ctx context.Context, _ *ContentNegotiationDifferentBodyClientGetAvatarAsJSONOptions) (*policy.Request, error) {
	urlPath := "/content-negotiation/different-body"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["accept"] = []string{"application/json"}
	return req, nil
}

// getAvatarAsJSONHandleResponse handles the GetAvatarAsJSON response.
func (client *ContentNegotiationDifferentBodyClient) getAvatarAsJSONHandleResponse(resp *http.Response) (ContentNegotiationDifferentBodyClientGetAvatarAsJSONResponse, error) {
	result := ContentNegotiationDifferentBodyClientGetAvatarAsJSONResponse{}
	if val := resp.Header.Get("content-type"); val != "" {
		result.ContentType = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PNGImageAsJSON); err != nil {
		return ContentNegotiationDifferentBodyClientGetAvatarAsJSONResponse{}, err
	}
	return result, nil
}

// GetAvatarAsPNG -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ContentNegotiationDifferentBodyClientGetAvatarAsPNGOptions contains the optional parameters for the ContentNegotiationDifferentBodyClient.GetAvatarAsPNG
//     method.
func (client *ContentNegotiationDifferentBodyClient) GetAvatarAsPNG(ctx context.Context, options *ContentNegotiationDifferentBodyClientGetAvatarAsPNGOptions) (ContentNegotiationDifferentBodyClientGetAvatarAsPNGResponse, error) {
	var err error
	const operationName = "ContentNegotiationDifferentBodyClient.GetAvatarAsPNG"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAvatarAsPNGCreateRequest(ctx, options)
	if err != nil {
		return ContentNegotiationDifferentBodyClientGetAvatarAsPNGResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentNegotiationDifferentBodyClientGetAvatarAsPNGResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContentNegotiationDifferentBodyClientGetAvatarAsPNGResponse{}, err
	}
	resp, err := client.getAvatarAsPNGHandleResponse(httpResp)
	return resp, err
}

// getAvatarAsPNGCreateRequest creates the GetAvatarAsPNG request.
func (client *ContentNegotiationDifferentBodyClient) getAvatarAsPNGCreateRequest(ctx context.Context, _ *ContentNegotiationDifferentBodyClientGetAvatarAsPNGOptions) (*policy.Request, error) {
	urlPath := "/content-negotiation/different-body"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header["accept"] = []string{"image/png"}
	return req, nil
}

// getAvatarAsPNGHandleResponse handles the GetAvatarAsPNG response.
func (client *ContentNegotiationDifferentBodyClient) getAvatarAsPNGHandleResponse(resp *http.Response) (ContentNegotiationDifferentBodyClientGetAvatarAsPNGResponse, error) {
	result := ContentNegotiationDifferentBodyClientGetAvatarAsPNGResponse{Body: resp.Body}
	if val := resp.Header.Get("content-type"); val != "" {
		result.ContentType = &val
	}
	return result, nil
}
