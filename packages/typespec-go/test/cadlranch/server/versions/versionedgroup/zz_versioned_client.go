// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package versionedgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// VersionedClient - Illustrates versioned server.
// Don't use this type directly, use a constructor function instead.
type VersionedClient struct {
	internal   *azcore.Client
	endpoint   string
	apiVersion string
}

//   - options - VersionedClientWithPathAPIVersionOptions contains the optional parameters for the VersionedClient.WithPathAPIVersion
//     method.
func (client *VersionedClient) WithPathAPIVersion(ctx context.Context, options *VersionedClientWithPathAPIVersionOptions) (VersionedClientWithPathAPIVersionResponse, error) {
	var err error
	req, err := client.withPathAPIVersionCreateRequest(ctx, options)
	if err != nil {
		return VersionedClientWithPathAPIVersionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VersionedClientWithPathAPIVersionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VersionedClientWithPathAPIVersionResponse{}, err
	}
	return VersionedClientWithPathAPIVersionResponse{}, nil
}

// withPathAPIVersionCreateRequest creates the WithPathAPIVersion request.
func (client *VersionedClient) withPathAPIVersionCreateRequest(ctx context.Context, options *VersionedClientWithPathAPIVersionOptions) (*policy.Request, error) {
	host := "{endpoint}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	req, err := runtime.NewRequest(ctx, http.MethodHead, host)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

//   - options - VersionedClientWithQueryAPIVersionOptions contains the optional parameters for the VersionedClient.WithQueryAPIVersion
//     method.
func (client *VersionedClient) WithQueryAPIVersion(ctx context.Context, options *VersionedClientWithQueryAPIVersionOptions) (VersionedClientWithQueryAPIVersionResponse, error) {
	var err error
	req, err := client.withQueryAPIVersionCreateRequest(ctx, options)
	if err != nil {
		return VersionedClientWithQueryAPIVersionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VersionedClientWithQueryAPIVersionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VersionedClientWithQueryAPIVersionResponse{}, err
	}
	return VersionedClientWithQueryAPIVersionResponse{}, nil
}

// withQueryAPIVersionCreateRequest creates the WithQueryAPIVersion request.
func (client *VersionedClient) withQueryAPIVersionCreateRequest(ctx context.Context, options *VersionedClientWithQueryAPIVersionOptions) (*policy.Request, error) {
	host := "{endpoint}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	urlPath := "/server/versions/versioned/with-query-api-version"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

//   - options - VersionedClientWithoutAPIVersionOptions contains the optional parameters for the VersionedClient.WithoutAPIVersion
//     method.
func (client *VersionedClient) WithoutAPIVersion(ctx context.Context, options *VersionedClientWithoutAPIVersionOptions) (VersionedClientWithoutAPIVersionResponse, error) {
	var err error
	req, err := client.withoutAPIVersionCreateRequest(ctx, options)
	if err != nil {
		return VersionedClientWithoutAPIVersionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VersionedClientWithoutAPIVersionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VersionedClientWithoutAPIVersionResponse{}, err
	}
	return VersionedClientWithoutAPIVersionResponse{}, nil
}

// withoutAPIVersionCreateRequest creates the WithoutAPIVersion request.
func (client *VersionedClient) withoutAPIVersionCreateRequest(ctx context.Context, options *VersionedClientWithoutAPIVersionOptions) (*policy.Request, error) {
	host := "{endpoint}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	urlPath := "/server/versions/versioned/without-api-version"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}