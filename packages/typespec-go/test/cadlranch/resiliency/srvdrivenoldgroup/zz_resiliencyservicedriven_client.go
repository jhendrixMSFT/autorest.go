// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package srvdrivenoldgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ResiliencyServiceDrivenClient - Test that we can grow up a service spec and service deployment into a multi-versioned service
// with full client support.
// Don't use this type directly, use a constructor function instead.
type ResiliencyServiceDrivenClient struct {
	internal                 *azcore.Client
	endpoint                 string
	serviceDeploymentVersion string
	apiVersion               string
}

// ResiliencyServiceDrivenClientOptions contains the optional values for creating a [ResiliencyServiceDrivenClient].
type ResiliencyServiceDrivenClientOptions struct {
	azcore.ClientOptions
}

// NewResiliencyServiceDrivenClientWithNoCredential creates a new [ResiliencyServiceDrivenClient].
//   - options - optional client configuration; pass nil to accept the default values
func NewResiliencyServiceDrivenClientWithNoCredential(options *ResiliencyServiceDrivenClientOptions) (*ResiliencyServiceDrivenClient, error) {
	if options == nil {
		options = &ResiliencyServiceDrivenClientOptions{}
	}
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	return &ResiliencyServiceDrivenClient{
		internal: internal,
	}, nil
}

// FromNone - Test that currently accepts no parameters, will be updated in next spec to accept a new optional parameter as
// well
//   - options - ResiliencyServiceDrivenClientFromNoneOptions contains the optional parameters for the ResiliencyServiceDrivenClient.FromNone
//     method.
func (client *ResiliencyServiceDrivenClient) FromNone(ctx context.Context, options *ResiliencyServiceDrivenClientFromNoneOptions) (ResiliencyServiceDrivenClientFromNoneResponse, error) {
	var err error
	const operationName = "ResiliencyServiceDrivenClient.FromNone"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fromNoneCreateRequest(ctx, options)
	if err != nil {
		return ResiliencyServiceDrivenClientFromNoneResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResiliencyServiceDrivenClientFromNoneResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ResiliencyServiceDrivenClientFromNoneResponse{}, err
	}
	return ResiliencyServiceDrivenClientFromNoneResponse{}, nil
}

// fromNoneCreateRequest creates the FromNone request.
func (client *ResiliencyServiceDrivenClient) fromNoneCreateRequest(ctx context.Context, options *ResiliencyServiceDrivenClientFromNoneOptions) (*policy.Request, error) {
	host := "{endpoint}/resiliency/service-driven/client:v1/service:{serviceDeploymentVersion}/api-version:{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{serviceDeploymentVersion}", client.serviceDeploymentVersion)
	host = strings.ReplaceAll(host, "{apiVersion}", client.apiVersion)
	urlPath := "/add-optional-param/from-none"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// FromOneOptional - Test that currently accepts one optional parameter, will be updated in next spec to accept a new optional
// parameter as well
//   - options - ResiliencyServiceDrivenClientFromOneOptionalOptions contains the optional parameters for the ResiliencyServiceDrivenClient.FromOneOptional
//     method.
func (client *ResiliencyServiceDrivenClient) FromOneOptional(ctx context.Context, options *ResiliencyServiceDrivenClientFromOneOptionalOptions) (ResiliencyServiceDrivenClientFromOneOptionalResponse, error) {
	var err error
	const operationName = "ResiliencyServiceDrivenClient.FromOneOptional"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fromOneOptionalCreateRequest(ctx, options)
	if err != nil {
		return ResiliencyServiceDrivenClientFromOneOptionalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResiliencyServiceDrivenClientFromOneOptionalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ResiliencyServiceDrivenClientFromOneOptionalResponse{}, err
	}
	return ResiliencyServiceDrivenClientFromOneOptionalResponse{}, nil
}

// fromOneOptionalCreateRequest creates the FromOneOptional request.
func (client *ResiliencyServiceDrivenClient) fromOneOptionalCreateRequest(ctx context.Context, options *ResiliencyServiceDrivenClientFromOneOptionalOptions) (*policy.Request, error) {
	host := "{endpoint}/resiliency/service-driven/client:v1/service:{serviceDeploymentVersion}/api-version:{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{serviceDeploymentVersion}", client.serviceDeploymentVersion)
	host = strings.ReplaceAll(host, "{apiVersion}", client.apiVersion)
	urlPath := "/add-optional-param/from-one-optional"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Parameter != nil {
		reqQP.Set("parameter", *options.Parameter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// FromOneRequired - Test that currently accepts one required parameter, will be updated in next spec to accept a new optional
// parameter as well
//   - parameter - I am a required parameter
//   - options - ResiliencyServiceDrivenClientFromOneRequiredOptions contains the optional parameters for the ResiliencyServiceDrivenClient.FromOneRequired
//     method.
func (client *ResiliencyServiceDrivenClient) FromOneRequired(ctx context.Context, parameter string, options *ResiliencyServiceDrivenClientFromOneRequiredOptions) (ResiliencyServiceDrivenClientFromOneRequiredResponse, error) {
	var err error
	const operationName = "ResiliencyServiceDrivenClient.FromOneRequired"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fromOneRequiredCreateRequest(ctx, parameter, options)
	if err != nil {
		return ResiliencyServiceDrivenClientFromOneRequiredResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResiliencyServiceDrivenClientFromOneRequiredResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ResiliencyServiceDrivenClientFromOneRequiredResponse{}, err
	}
	return ResiliencyServiceDrivenClientFromOneRequiredResponse{}, nil
}

// fromOneRequiredCreateRequest creates the FromOneRequired request.
func (client *ResiliencyServiceDrivenClient) fromOneRequiredCreateRequest(ctx context.Context, parameter string, options *ResiliencyServiceDrivenClientFromOneRequiredOptions) (*policy.Request, error) {
	host := "{endpoint}/resiliency/service-driven/client:v1/service:{serviceDeploymentVersion}/api-version:{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{serviceDeploymentVersion}", client.serviceDeploymentVersion)
	host = strings.ReplaceAll(host, "{apiVersion}", client.apiVersion)
	urlPath := "/add-optional-param/from-one-required"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("parameter", parameter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
