//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package twoopgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ServiceClient contains the methods for the Client.Structure.Service group.
// Don't use this type directly, use a constructor function instead.
type ServiceClient struct {
	internal *azcore.Client
}

// - options - ServiceClientOneOptions contains the optional parameters for the ServiceClient.One method.
func (client *ServiceClient) One(ctx context.Context, options *ServiceClientOneOptions) (ServiceClientOneResponse, error) {
	var err error
	req, err := client.oneCreateRequest(ctx, options)
	if err != nil {
		return ServiceClientOneResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientOneResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientOneResponse{}, err
	}
	return ServiceClientOneResponse{}, nil
}

// oneCreateRequest creates the One request.
func (client *ServiceClient) oneCreateRequest(ctx context.Context, options *ServiceClientOneOptions) (*policy.Request, error) {
	urlPath := "/one"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - ServiceClientTwoOptions contains the optional parameters for the ServiceClient.Two method.
func (client *ServiceClient) Two(ctx context.Context, options *ServiceClientTwoOptions) (ServiceClientTwoResponse, error) {
	var err error
	req, err := client.twoCreateRequest(ctx, options)
	if err != nil {
		return ServiceClientTwoResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientTwoResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientTwoResponse{}, err
	}
	return ServiceClientTwoResponse{}, nil
}

// twoCreateRequest creates the Two request.
func (client *ServiceClient) twoCreateRequest(ctx context.Context, options *ServiceClientTwoOptions) (*policy.Request, error) {
	urlPath := "/two"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
