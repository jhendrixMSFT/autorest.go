// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package unversionedgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// NotVersionedClient - Illustrates not-versioned server.
// Don't use this type directly, use NewNotVersionedClientWithNoCredential() instead.
type NotVersionedClient struct {
	internal   *azcore.Client
	endpoint   string
	apiVersion string
}

// NotVersionedClientOptions contains the optional values for creating a [NotVersionedClient].
type NotVersionedClientOptions struct {
	azcore.ClientOptions
}

// NewNotVersionedClientWithNoCredential creates a new instance of [NotVersionedClient] with the specified values.
//   - options - NotVersionedClientOptions contains the optional values for creating a [NotVersionedClient]
func NewNotVersionedClientWithNoCredential(endpoint string, apiVersion string, options *NotVersionedClientOptions) (*NotVersionedClient, error) {
	if options == nil {
		options = &NotVersionedClientOptions{}
	}
	cl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	notVersionedClient := &NotVersionedClient{
		endpoint:   endpoint,
		apiVersion: apiVersion,
		internal:   cl,
	}
	return notVersionedClient, nil
}

//   - options - NotVersionedClientWithPathAPIVersionOptions contains the optional parameters for the NotVersionedClient.WithPathAPIVersion
//     method.
func (client *NotVersionedClient) WithPathAPIVersion(ctx context.Context, options *NotVersionedClientWithPathAPIVersionOptions) (NotVersionedClientWithPathAPIVersionResponse, error) {
	var err error
	const operationName = "NotVersionedClient.WithPathAPIVersion"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.withPathAPIVersionCreateRequest(ctx, options)
	if err != nil {
		return NotVersionedClientWithPathAPIVersionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotVersionedClientWithPathAPIVersionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotVersionedClientWithPathAPIVersionResponse{}, err
	}
	return NotVersionedClientWithPathAPIVersionResponse{}, nil
}

// withPathAPIVersionCreateRequest creates the WithPathAPIVersion request.
func (client *NotVersionedClient) withPathAPIVersionCreateRequest(ctx context.Context, options *NotVersionedClientWithPathAPIVersionOptions) (*policy.Request, error) {
	host := "{endpoint}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	urlPath := "/server/versions/not-versioned/with-path-api-version/{apiVersion}"
	if client.apiVersion == "" {
		return nil, errors.New("parameter client.apiVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiVersion}", url.PathEscape(client.apiVersion))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

//   - options - NotVersionedClientWithQueryAPIVersionOptions contains the optional parameters for the NotVersionedClient.WithQueryAPIVersion
//     method.
func (client *NotVersionedClient) WithQueryAPIVersion(ctx context.Context, options *NotVersionedClientWithQueryAPIVersionOptions) (NotVersionedClientWithQueryAPIVersionResponse, error) {
	var err error
	const operationName = "NotVersionedClient.WithQueryAPIVersion"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.withQueryAPIVersionCreateRequest(ctx, options)
	if err != nil {
		return NotVersionedClientWithQueryAPIVersionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotVersionedClientWithQueryAPIVersionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotVersionedClientWithQueryAPIVersionResponse{}, err
	}
	return NotVersionedClientWithQueryAPIVersionResponse{}, nil
}

// withQueryAPIVersionCreateRequest creates the WithQueryAPIVersion request.
func (client *NotVersionedClient) withQueryAPIVersionCreateRequest(ctx context.Context, options *NotVersionedClientWithQueryAPIVersionOptions) (*policy.Request, error) {
	host := "{endpoint}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	urlPath := "/server/versions/not-versioned/with-query-api-version"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

//   - options - NotVersionedClientWithoutAPIVersionOptions contains the optional parameters for the NotVersionedClient.WithoutAPIVersion
//     method.
func (client *NotVersionedClient) WithoutAPIVersion(ctx context.Context, options *NotVersionedClientWithoutAPIVersionOptions) (NotVersionedClientWithoutAPIVersionResponse, error) {
	var err error
	const operationName = "NotVersionedClient.WithoutAPIVersion"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.withoutAPIVersionCreateRequest(ctx, options)
	if err != nil {
		return NotVersionedClientWithoutAPIVersionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotVersionedClientWithoutAPIVersionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotVersionedClientWithoutAPIVersionResponse{}, err
	}
	return NotVersionedClientWithoutAPIVersionResponse{}, nil
}

// withoutAPIVersionCreateRequest creates the WithoutAPIVersion request.
func (client *NotVersionedClient) withoutAPIVersionCreateRequest(ctx context.Context, options *NotVersionedClientWithoutAPIVersionOptions) (*policy.Request, error) {
	host := "{endpoint}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	urlPath := "/server/versions/not-versioned/without-api-version"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
