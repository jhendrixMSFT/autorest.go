//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// APIVersionDefaultClient contains the methods for the APIVersionDefault group.
// Don't use this type directly, use a constructor function instead.
type APIVersionDefaultClient struct {
	internal *azcore.Client
}

// GetMethodGlobalNotProvidedValid - GET method with api-version modeled in global settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - options - APIVersionDefaultClientGetMethodGlobalNotProvidedValidOptions contains the optional parameters for the APIVersionDefaultClient.GetMethodGlobalNotProvidedValid
//     method.
func (client *APIVersionDefaultClient) GetMethodGlobalNotProvidedValid(ctx context.Context, options *APIVersionDefaultClientGetMethodGlobalNotProvidedValidOptions) (APIVersionDefaultClientGetMethodGlobalNotProvidedValidResponse, error) {
	req, err := client.getMethodGlobalNotProvidedValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionDefaultClientGetMethodGlobalNotProvidedValidResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIVersionDefaultClientGetMethodGlobalNotProvidedValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionDefaultClientGetMethodGlobalNotProvidedValidResponse{}, runtime.NewResponseError(resp)
	}
	return APIVersionDefaultClientGetMethodGlobalNotProvidedValidResponse{}, nil
}

// getMethodGlobalNotProvidedValidCreateRequest creates the GetMethodGlobalNotProvidedValid request.
func (client *APIVersionDefaultClient) getMethodGlobalNotProvidedValidCreateRequest(ctx context.Context, options *APIVersionDefaultClientGetMethodGlobalNotProvidedValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/globalNotProvided/2015-07-01-preview"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetMethodGlobalValid - GET method with api-version modeled in global settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - options - APIVersionDefaultClientGetMethodGlobalValidOptions contains the optional parameters for the APIVersionDefaultClient.GetMethodGlobalValid
//     method.
func (client *APIVersionDefaultClient) GetMethodGlobalValid(ctx context.Context, options *APIVersionDefaultClientGetMethodGlobalValidOptions) (APIVersionDefaultClientGetMethodGlobalValidResponse, error) {
	req, err := client.getMethodGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionDefaultClientGetMethodGlobalValidResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIVersionDefaultClientGetMethodGlobalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionDefaultClientGetMethodGlobalValidResponse{}, runtime.NewResponseError(resp)
	}
	return APIVersionDefaultClientGetMethodGlobalValidResponse{}, nil
}

// getMethodGlobalValidCreateRequest creates the GetMethodGlobalValid request.
func (client *APIVersionDefaultClient) getMethodGlobalValidCreateRequest(ctx context.Context, options *APIVersionDefaultClientGetMethodGlobalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/global/2015-07-01-preview"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetPathGlobalValid - GET method with api-version modeled in global settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - options - APIVersionDefaultClientGetPathGlobalValidOptions contains the optional parameters for the APIVersionDefaultClient.GetPathGlobalValid
//     method.
func (client *APIVersionDefaultClient) GetPathGlobalValid(ctx context.Context, options *APIVersionDefaultClientGetPathGlobalValidOptions) (APIVersionDefaultClientGetPathGlobalValidResponse, error) {
	req, err := client.getPathGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionDefaultClientGetPathGlobalValidResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIVersionDefaultClientGetPathGlobalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionDefaultClientGetPathGlobalValidResponse{}, runtime.NewResponseError(resp)
	}
	return APIVersionDefaultClientGetPathGlobalValidResponse{}, nil
}

// getPathGlobalValidCreateRequest creates the GetPathGlobalValid request.
func (client *APIVersionDefaultClient) getPathGlobalValidCreateRequest(ctx context.Context, options *APIVersionDefaultClientGetPathGlobalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/path/string/none/query/global/2015-07-01-preview"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetSwaggerGlobalValid - GET method with api-version modeled in global settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - options - APIVersionDefaultClientGetSwaggerGlobalValidOptions contains the optional parameters for the APIVersionDefaultClient.GetSwaggerGlobalValid
//     method.
func (client *APIVersionDefaultClient) GetSwaggerGlobalValid(ctx context.Context, options *APIVersionDefaultClientGetSwaggerGlobalValidOptions) (APIVersionDefaultClientGetSwaggerGlobalValidResponse, error) {
	req, err := client.getSwaggerGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionDefaultClientGetSwaggerGlobalValidResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIVersionDefaultClientGetSwaggerGlobalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionDefaultClientGetSwaggerGlobalValidResponse{}, runtime.NewResponseError(resp)
	}
	return APIVersionDefaultClientGetSwaggerGlobalValidResponse{}, nil
}

// getSwaggerGlobalValidCreateRequest creates the GetSwaggerGlobalValid request.
func (client *APIVersionDefaultClient) getSwaggerGlobalValidCreateRequest(ctx context.Context, options *APIVersionDefaultClientGetSwaggerGlobalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/swagger/string/none/query/global/2015-07-01-preview"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
