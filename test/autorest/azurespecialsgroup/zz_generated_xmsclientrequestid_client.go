//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// XMSClientRequestIDClient contains the methods for the XMSClientRequestID group.
// Don't use this type directly, use NewXMSClientRequestIDClient() instead.
type XMSClientRequestIDClient struct {
	pl runtime.Pipeline
}

// NewXMSClientRequestIDClient creates a new instance of XMSClientRequestIDClient with the specified values.
// options - pass nil to accept the default values.
func NewXMSClientRequestIDClient(options *azcore.ClientOptions) *XMSClientRequestIDClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &XMSClientRequestIDClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// If the operation fails it returns an *azcore.ResponseError type.
// options - XMSClientRequestIDClientGetOptions contains the optional parameters for the XMSClientRequestIDClient.Get method.
func (client *XMSClientRequestIDClient) Get(ctx context.Context, options *XMSClientRequestIDClientGetOptions) (XMSClientRequestIDClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return XMSClientRequestIDClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return XMSClientRequestIDClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return XMSClientRequestIDClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return XMSClientRequestIDClientGetResponse{RawResponse: resp}, nil
}

// getCreateRequest creates the Get request.
func (client *XMSClientRequestIDClient) getCreateRequest(ctx context.Context, options *XMSClientRequestIDClientGetOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/method/"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// If the operation fails it returns an *azcore.ResponseError type.
// xmsClientRequestID - This should appear as a method parameter, use value '9C4D50EE-2D56-4CD3-8152-34347DC9F2B0'
// options - XMSClientRequestIDClientParamGetOptions contains the optional parameters for the XMSClientRequestIDClient.ParamGet
// method.
func (client *XMSClientRequestIDClient) ParamGet(ctx context.Context, xmsClientRequestID string, options *XMSClientRequestIDClientParamGetOptions) (XMSClientRequestIDClientParamGetResponse, error) {
	req, err := client.paramGetCreateRequest(ctx, xmsClientRequestID, options)
	if err != nil {
		return XMSClientRequestIDClientParamGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return XMSClientRequestIDClientParamGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return XMSClientRequestIDClientParamGetResponse{}, runtime.NewResponseError(resp)
	}
	return XMSClientRequestIDClientParamGetResponse{RawResponse: resp}, nil
}

// paramGetCreateRequest creates the ParamGet request.
func (client *XMSClientRequestIDClient) paramGetCreateRequest(ctx context.Context, xmsClientRequestID string, options *XMSClientRequestIDClientParamGetOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/via-param/method/"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("x-ms-client-request-id", xmsClientRequestID)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
