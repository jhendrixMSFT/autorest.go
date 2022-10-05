//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

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

// SQLPoolsClient contains the methods for the SQLPools group.
// Don't use this type directly, use NewSQLPoolsClient() instead.
type SQLPoolsClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// NewSQLPoolsClient creates a new instance of SQLPoolsClient with the specified values.
//   - endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLPoolsClient(endpoint string, credential azcore.TokenCredential, options *SQLPoolsClientOptions) *SQLPoolsClient {
	if options == nil {
		options = &SQLPoolsClientOptions{}
	}
	authPolicy := runtime.NewBearerTokenPolicy(credential, []string{"https://dev.azuresynapse.net/.default"}, nil)
	pl := runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{PerRetry: []policy.Policy{authPolicy}}, &options.ClientOptions)
	client := &SQLPoolsClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Get - Get Sql Pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlPoolName - The Sql Pool name
//   - options - SQLPoolsClientGetOptions contains the optional parameters for the SQLPoolsClient.Get method.
func (client *SQLPoolsClient) Get(ctx context.Context, sqlPoolName string, options *SQLPoolsClientGetOptions) (SQLPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, sqlPoolName, options)
	if err != nil {
		return SQLPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolsClient) getCreateRequest(ctx context.Context, sqlPoolName string, options *SQLPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/sqlPools/{sqlPoolName}"
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLPoolsClient) getHandleResponse(resp *http.Response) (SQLPoolsClientGetResponse, error) {
	result := SQLPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPool); err != nil {
		return SQLPoolsClientGetResponse{}, err
	}
	return result, nil
}

// List - List Sql Pools
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - options - SQLPoolsClientListOptions contains the optional parameters for the SQLPoolsClient.List method.
func (client *SQLPoolsClient) List(ctx context.Context, options *SQLPoolsClientListOptions) (SQLPoolsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return SQLPoolsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SQLPoolsClient) listCreateRequest(ctx context.Context, options *SQLPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/sqlPools"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLPoolsClient) listHandleResponse(resp *http.Response) (SQLPoolsClientListResponse, error) {
	result := SQLPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPoolInfoListResult); err != nil {
		return SQLPoolsClientListResponse{}, err
	}
	return result, nil
}
