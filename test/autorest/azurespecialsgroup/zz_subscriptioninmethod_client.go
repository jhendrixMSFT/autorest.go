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
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionInMethodClient contains the methods for the SubscriptionInMethod group.
// Don't use this type directly, use a constructor function instead.
type SubscriptionInMethodClient struct {
	internal *azcore.Client
}

// PostMethodLocalNull - POST method with subscriptionId modeled in the method. pass in subscription id = null, client-side
// validation should prevent you from making this call
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - subscriptionID - This should appear as a method parameter, use value null, client-side validation should prvenet the call
//   - options - SubscriptionInMethodClientPostMethodLocalNullOptions contains the optional parameters for the SubscriptionInMethodClient.PostMethodLocalNull
//     method.
func (client *SubscriptionInMethodClient) PostMethodLocalNull(ctx context.Context, subscriptionID string, options *SubscriptionInMethodClientPostMethodLocalNullOptions) (SubscriptionInMethodClientPostMethodLocalNullResponse, error) {
	req, err := client.postMethodLocalNullCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return SubscriptionInMethodClientPostMethodLocalNullResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionInMethodClientPostMethodLocalNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionInMethodClientPostMethodLocalNullResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionInMethodClientPostMethodLocalNullResponse{}, nil
}

// postMethodLocalNullCreateRequest creates the PostMethodLocalNull request.
func (client *SubscriptionInMethodClient) postMethodLocalNullCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionInMethodClientPostMethodLocalNullOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/local/null/{subscriptionId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// PostMethodLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456'
// to succeed
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - subscriptionID - This should appear as a method parameter, use value '1234-5678-9012-3456'
//   - options - SubscriptionInMethodClientPostMethodLocalValidOptions contains the optional parameters for the SubscriptionInMethodClient.PostMethodLocalValid
//     method.
func (client *SubscriptionInMethodClient) PostMethodLocalValid(ctx context.Context, subscriptionID string, options *SubscriptionInMethodClientPostMethodLocalValidOptions) (SubscriptionInMethodClientPostMethodLocalValidResponse, error) {
	req, err := client.postMethodLocalValidCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return SubscriptionInMethodClientPostMethodLocalValidResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionInMethodClientPostMethodLocalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionInMethodClientPostMethodLocalValidResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionInMethodClientPostMethodLocalValidResponse{}, nil
}

// postMethodLocalValidCreateRequest creates the PostMethodLocalValid request.
func (client *SubscriptionInMethodClient) postMethodLocalValidCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionInMethodClientPostMethodLocalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// PostPathLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456'
// to succeed
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - subscriptionID - Should appear as a method parameter -use value '1234-5678-9012-3456'
//   - options - SubscriptionInMethodClientPostPathLocalValidOptions contains the optional parameters for the SubscriptionInMethodClient.PostPathLocalValid
//     method.
func (client *SubscriptionInMethodClient) PostPathLocalValid(ctx context.Context, subscriptionID string, options *SubscriptionInMethodClientPostPathLocalValidOptions) (SubscriptionInMethodClientPostPathLocalValidResponse, error) {
	req, err := client.postPathLocalValidCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return SubscriptionInMethodClientPostPathLocalValidResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionInMethodClientPostPathLocalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionInMethodClientPostPathLocalValidResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionInMethodClientPostPathLocalValidResponse{}, nil
}

// postPathLocalValidCreateRequest creates the PostPathLocalValid request.
func (client *SubscriptionInMethodClient) postPathLocalValidCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionInMethodClientPostPathLocalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/subscriptionId/path/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// PostSwaggerLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456'
// to succeed
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - subscriptionID - The subscriptionId, which appears in the path, the value is always '1234-5678-9012-3456'
//   - options - SubscriptionInMethodClientPostSwaggerLocalValidOptions contains the optional parameters for the SubscriptionInMethodClient.PostSwaggerLocalValid
//     method.
func (client *SubscriptionInMethodClient) PostSwaggerLocalValid(ctx context.Context, subscriptionID string, options *SubscriptionInMethodClientPostSwaggerLocalValidOptions) (SubscriptionInMethodClientPostSwaggerLocalValidResponse, error) {
	req, err := client.postSwaggerLocalValidCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return SubscriptionInMethodClientPostSwaggerLocalValidResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionInMethodClientPostSwaggerLocalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionInMethodClientPostSwaggerLocalValidResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionInMethodClientPostSwaggerLocalValidResponse{}, nil
}

// postSwaggerLocalValidCreateRequest creates the PostSwaggerLocalValid request.
func (client *SubscriptionInMethodClient) postSwaggerLocalValidCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionInMethodClientPostSwaggerLocalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/subscriptionId/swagger/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
