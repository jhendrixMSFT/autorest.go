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
	"strings"
)

// ClientAClient contains the methods for the Client.Structure.Service namespace.
// Don't use this type directly, use a constructor function instead.
type ClientAClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// ClientAClientOptions contains the optional values for creating a [ClientAClient].
type ClientAClientOptions struct {
	azcore.ClientOptions
}

// NewClientAClientWithNoCredential creates a new [ClientAClient].
//   - options - optional client configuration; pass nil to accept the default values
func NewClientAClientWithNoCredential(options *ClientAClientOptions) (*ClientAClient, error) {
	if options == nil {
		options = &ClientAClientOptions{}
	}
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	return &ClientAClient{
		internal: internal,
	}, nil
}

// - options - ClientAClientRenamedFiveOptions contains the optional parameters for the ClientAClient.RenamedFive method.
func (client *ClientAClient) RenamedFive(ctx context.Context, options *ClientAClientRenamedFiveOptions) (ClientAClientRenamedFiveResponse, error) {
	var err error
	const operationName = "ClientAClient.RenamedFive"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.renamedFiveCreateRequest(ctx, options)
	if err != nil {
		return ClientAClientRenamedFiveResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientAClientRenamedFiveResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientAClientRenamedFiveResponse{}, err
	}
	return ClientAClientRenamedFiveResponse{}, nil
}

// renamedFiveCreateRequest creates the RenamedFive request.
func (client *ClientAClient) renamedFiveCreateRequest(ctx context.Context, options *ClientAClientRenamedFiveOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/five"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - ClientAClientRenamedOneOptions contains the optional parameters for the ClientAClient.RenamedOne method.
func (client *ClientAClient) RenamedOne(ctx context.Context, options *ClientAClientRenamedOneOptions) (ClientAClientRenamedOneResponse, error) {
	var err error
	const operationName = "ClientAClient.RenamedOne"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.renamedOneCreateRequest(ctx, options)
	if err != nil {
		return ClientAClientRenamedOneResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientAClientRenamedOneResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientAClientRenamedOneResponse{}, err
	}
	return ClientAClientRenamedOneResponse{}, nil
}

// renamedOneCreateRequest creates the RenamedOne request.
func (client *ClientAClient) renamedOneCreateRequest(ctx context.Context, options *ClientAClientRenamedOneOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/one"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - ClientAClientRenamedThreeOptions contains the optional parameters for the ClientAClient.RenamedThree method.
func (client *ClientAClient) RenamedThree(ctx context.Context, options *ClientAClientRenamedThreeOptions) (ClientAClientRenamedThreeResponse, error) {
	var err error
	const operationName = "ClientAClient.RenamedThree"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.renamedThreeCreateRequest(ctx, options)
	if err != nil {
		return ClientAClientRenamedThreeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientAClientRenamedThreeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientAClientRenamedThreeResponse{}, err
	}
	return ClientAClientRenamedThreeResponse{}, nil
}

// renamedThreeCreateRequest creates the RenamedThree request.
func (client *ClientAClient) renamedThreeCreateRequest(ctx context.Context, options *ClientAClientRenamedThreeOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/three"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
