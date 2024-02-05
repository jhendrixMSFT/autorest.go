//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package multiclientgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ClientBClient contains the methods for the Client.Structure.Service group.
// Don't use this type directly, use a constructor function instead.
type ClientBClient struct {
	internal *azcore.Client
	endpoint string
}

// - options - ClientBClientRenamedFourOptions contains the optional parameters for the ClientBClient.RenamedFour method.
func (client *ClientBClient) RenamedFour(ctx context.Context, options *ClientBClientRenamedFourOptions) (ClientBClientRenamedFourResponse, error) {
	var err error
	req, err := client.renamedFourCreateRequest(ctx, options)
	if err != nil {
		return ClientBClientRenamedFourResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientBClientRenamedFourResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientBClientRenamedFourResponse{}, err
	}
	return ClientBClientRenamedFourResponse{}, nil
}

// renamedFourCreateRequest creates the RenamedFour request.
func (client *ClientBClient) renamedFourCreateRequest(ctx context.Context, options *ClientBClientRenamedFourOptions) (*policy.Request, error) {
	urlPath := "/four"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - ClientBClientRenamedSixOptions contains the optional parameters for the ClientBClient.RenamedSix method.
func (client *ClientBClient) RenamedSix(ctx context.Context, options *ClientBClientRenamedSixOptions) (ClientBClientRenamedSixResponse, error) {
	var err error
	req, err := client.renamedSixCreateRequest(ctx, options)
	if err != nil {
		return ClientBClientRenamedSixResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientBClientRenamedSixResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientBClientRenamedSixResponse{}, err
	}
	return ClientBClientRenamedSixResponse{}, nil
}

// renamedSixCreateRequest creates the RenamedSix request.
func (client *ClientBClient) renamedSixCreateRequest(ctx context.Context, options *ClientBClientRenamedSixOptions) (*policy.Request, error) {
	urlPath := "/six"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - ClientBClientRenamedTwoOptions contains the optional parameters for the ClientBClient.RenamedTwo method.
func (client *ClientBClient) RenamedTwo(ctx context.Context, options *ClientBClientRenamedTwoOptions) (ClientBClientRenamedTwoResponse, error) {
	var err error
	req, err := client.renamedTwoCreateRequest(ctx, options)
	if err != nil {
		return ClientBClientRenamedTwoResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientBClientRenamedTwoResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientBClientRenamedTwoResponse{}, err
	}
	return ClientBClientRenamedTwoResponse{}, nil
}

// renamedTwoCreateRequest creates the RenamedTwo request.
func (client *ClientBClient) renamedTwoCreateRequest(ctx context.Context, options *ClientBClientRenamedTwoOptions) (*policy.Request, error) {
	urlPath := "/two"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
