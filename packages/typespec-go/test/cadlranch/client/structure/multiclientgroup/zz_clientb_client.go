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

// ClientBClient contains the methods for the Client.Structure.Service namespace.
// Don't use this type directly, use NewClientBClientWithNoCredential() instead.
type ClientBClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// ClientBClientOptions contains the optional values for creating a [ClientBClient].
type ClientBClientOptions struct {
	azcore.ClientOptions
}

// NewClientBClientWithNoCredential creates a new instance of [ClientBClient] with the specified values.
//   - options - ClientBClientOptions contains the optional values for creating a [ClientBClient]
func NewClientBClientWithNoCredential(endpoint string, client ClientType, options *ClientBClientOptions) (*ClientBClient, error) {
	if options == nil {
		options = &ClientBClientOptions{}
	}
	cl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	clientBClient := &ClientBClient{
		endpoint: endpoint,
		client:   client,
		internal: cl,
	}
	return clientBClient, nil
}

// - options - ClientBClientRenamedFourOptions contains the optional parameters for the ClientBClient.RenamedFour method.
func (client *ClientBClient) RenamedFour(ctx context.Context, options *ClientBClientRenamedFourOptions) (ClientBClientRenamedFourResponse, error) {
	var err error
	const operationName = "ClientBClient.RenamedFour"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
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
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/four"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - ClientBClientRenamedSixOptions contains the optional parameters for the ClientBClient.RenamedSix method.
func (client *ClientBClient) RenamedSix(ctx context.Context, options *ClientBClientRenamedSixOptions) (ClientBClientRenamedSixResponse, error) {
	var err error
	const operationName = "ClientBClient.RenamedSix"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
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
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/six"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - ClientBClientRenamedTwoOptions contains the optional parameters for the ClientBClient.RenamedTwo method.
func (client *ClientBClient) RenamedTwo(ctx context.Context, options *ClientBClientRenamedTwoOptions) (ClientBClientRenamedTwoResponse, error) {
	var err error
	const operationName = "ClientBClient.RenamedTwo"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
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
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/two"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
