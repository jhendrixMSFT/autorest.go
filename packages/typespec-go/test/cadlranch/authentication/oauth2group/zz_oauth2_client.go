//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package oauth2group

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OAuth2Client contains the methods for the Authentication.OAuth2 group.
// Don't use this type directly, use a constructor function instead.
type OAuth2Client struct {
	internal *azcore.Client
	endpoint string
}

// Invalid - Check whether client is authenticated. Will return an invalid bearer error.
//   - options - OAuth2ClientInvalidOptions contains the optional parameters for the OAuth2Client.Invalid method.
func (client *OAuth2Client) Invalid(ctx context.Context, options *OAuth2ClientInvalidOptions) (OAuth2ClientInvalidResponse, error) {
	var err error
	req, err := client.invalidCreateRequest(ctx, options)
	if err != nil {
		return OAuth2ClientInvalidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OAuth2ClientInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OAuth2ClientInvalidResponse{}, err
	}
	return OAuth2ClientInvalidResponse{}, nil
}

// invalidCreateRequest creates the Invalid request.
func (client *OAuth2Client) invalidCreateRequest(ctx context.Context, options *OAuth2ClientInvalidOptions) (*policy.Request, error) {
	urlPath := "/authentication/oauth2/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Valid - Check whether client is authenticated
//   - options - OAuth2ClientValidOptions contains the optional parameters for the OAuth2Client.Valid method.
func (client *OAuth2Client) Valid(ctx context.Context, options *OAuth2ClientValidOptions) (OAuth2ClientValidResponse, error) {
	var err error
	req, err := client.validCreateRequest(ctx, options)
	if err != nil {
		return OAuth2ClientValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OAuth2ClientValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OAuth2ClientValidResponse{}, err
	}
	return OAuth2ClientValidResponse{}, nil
}

// validCreateRequest creates the Valid request.
func (client *OAuth2Client) validCreateRequest(ctx context.Context, options *OAuth2ClientValidOptions) (*policy.Request, error) {
	urlPath := "/authentication/oauth2/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
