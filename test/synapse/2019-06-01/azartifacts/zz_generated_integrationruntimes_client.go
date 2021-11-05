//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type integrationRuntimesClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newIntegrationRuntimesClient creates a new instance of integrationRuntimesClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newIntegrationRuntimesClient(endpoint string, pl runtime.Pipeline) *integrationRuntimesClient {
	client := &integrationRuntimesClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Get - Get Integration Runtime
// If the operation fails it returns the *ErrorContract error type.
// integrationRuntimeName - The Integration Runtime name
// options - IntegrationRuntimesGetOptions contains the optional parameters for the IntegrationRuntimes.Get method.
func (client *integrationRuntimesClient) Get(ctx context.Context, integrationRuntimeName string, options *IntegrationRuntimesGetOptions) (IntegrationRuntimesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *integrationRuntimesClient) getCreateRequest(ctx context.Context, integrationRuntimeName string, options *IntegrationRuntimesGetOptions) (*policy.Request, error) {
	urlPath := "/integrationRuntimes/{integrationRuntimeName}"
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *integrationRuntimesClient) getHandleResponse(resp *http.Response) (IntegrationRuntimesGetResponse, error) {
	result := IntegrationRuntimesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeResource); err != nil {
		return IntegrationRuntimesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *integrationRuntimesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List Integration Runtimes
// If the operation fails it returns the *ErrorContract error type.
// options - IntegrationRuntimesListOptions contains the optional parameters for the IntegrationRuntimes.List method.
func (client *integrationRuntimesClient) List(ctx context.Context, options *IntegrationRuntimesListOptions) (IntegrationRuntimesListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return IntegrationRuntimesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *integrationRuntimesClient) listCreateRequest(ctx context.Context, options *IntegrationRuntimesListOptions) (*policy.Request, error) {
	urlPath := "/integrationRuntimes"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *integrationRuntimesClient) listHandleResponse(resp *http.Response) (IntegrationRuntimesListResponse, error) {
	result := IntegrationRuntimesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeListResponse); err != nil {
		return IntegrationRuntimesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *integrationRuntimesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
