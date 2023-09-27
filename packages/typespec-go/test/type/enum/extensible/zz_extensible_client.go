//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT. DO NOT EDIT.

package extensible

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ExtensibleClient contains the methods for the Type.Enum.Extensible group.
// Don't use this type directly, use a constructor function instead.
type ExtensibleClient struct {
	internal *azcore.Client
}

func (client *ExtensibleClient) GetKnownValue(ctx context.Context, options *ExtensibleClientGetKnownValueOptions) (ExtensibleClientGetKnownValueResponse, error) {
	var err error
	req, err := client.getKnownValueCreateRequest(ctx, options)
	if err != nil {
		return ExtensibleClientGetKnownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensibleClientGetKnownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtensibleClientGetKnownValueResponse{}, err
	}
	resp, err := client.getKnownValueHandleResponse(httpResp)
	return resp, err
}

// getKnownValueCreateRequest creates the GetKnownValue request.
func (client *ExtensibleClient) getKnownValueCreateRequest(ctx context.Context, options *ExtensibleClientGetKnownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/known-value"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// getKnownValueHandleResponse handles the GetKnownValue response.
func (client *ExtensibleClient) getKnownValueHandleResponse(resp *http.Response) (ExtensibleClientGetKnownValueResponse, error) {
	result := ExtensibleClientGetKnownValueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ExtensibleClientGetKnownValueResponse{}, err
	}
	return result, nil
}

func (client *ExtensibleClient) GetUnknownValue(ctx context.Context, options *ExtensibleClientGetUnknownValueOptions) (ExtensibleClientGetUnknownValueResponse, error) {
	var err error
	req, err := client.getUnknownValueCreateRequest(ctx, options)
	if err != nil {
		return ExtensibleClientGetUnknownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensibleClientGetUnknownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtensibleClientGetUnknownValueResponse{}, err
	}
	resp, err := client.getUnknownValueHandleResponse(httpResp)
	return resp, err
}

// getUnknownValueCreateRequest creates the GetUnknownValue request.
func (client *ExtensibleClient) getUnknownValueCreateRequest(ctx context.Context, options *ExtensibleClientGetUnknownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/unknown-value"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// getUnknownValueHandleResponse handles the GetUnknownValue response.
func (client *ExtensibleClient) getUnknownValueHandleResponse(resp *http.Response) (ExtensibleClientGetUnknownValueResponse, error) {
	result := ExtensibleClientGetUnknownValueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ExtensibleClientGetUnknownValueResponse{}, err
	}
	return result, nil
}

func (client *ExtensibleClient) PutKnownValue(ctx context.Context, body DaysOfWeekExtensibleEnum, options *ExtensibleClientPutKnownValueOptions) (ExtensibleClientPutKnownValueResponse, error) {
	var err error
	req, err := client.putKnownValueCreateRequest(ctx, body, options)
	if err != nil {
		return ExtensibleClientPutKnownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensibleClientPutKnownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ExtensibleClientPutKnownValueResponse{}, err
	}
	return ExtensibleClientPutKnownValueResponse{}, nil
}

// putKnownValueCreateRequest creates the PutKnownValue request.
func (client *ExtensibleClient) putKnownValueCreateRequest(ctx context.Context, body DaysOfWeekExtensibleEnum, options *ExtensibleClientPutKnownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/known-value"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

func (client *ExtensibleClient) PutUnknownValue(ctx context.Context, body DaysOfWeekExtensibleEnum, options *ExtensibleClientPutUnknownValueOptions) (ExtensibleClientPutUnknownValueResponse, error) {
	var err error
	req, err := client.putUnknownValueCreateRequest(ctx, body, options)
	if err != nil {
		return ExtensibleClientPutUnknownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensibleClientPutUnknownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ExtensibleClientPutUnknownValueResponse{}, err
	}
	return ExtensibleClientPutUnknownValueResponse{}, nil
}

// putUnknownValueCreateRequest creates the PutUnknownValue request.
func (client *ExtensibleClient) putUnknownValueCreateRequest(ctx context.Context, body DaysOfWeekExtensibleEnum, options *ExtensibleClientPutUnknownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/unknown-value"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
