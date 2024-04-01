// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package defaultgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ServiceClient - Test that we can use @client and @operationGroup decorators to customize client side code structure, such
// as:
// 1. have everything as default.
// 2. to rename client or operation group
// 3. one client can have more than one operations groups
// 4. split one interface into two clients
// 5. have two clients with operations come from different interfaces
// 6. have two clients with a hierarchy relation.
// Don't use this type directly, use a constructor function instead.
type ServiceClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// NewServiceBarClient creates a new instance of [ServiceBarClient].
func (client *ServiceClient) NewServiceBarClient() *ServiceBarClient {
	return &ServiceBarClient{
		internal: client.internal,
		endpoint: client.endpoint,
		client:   client.client,
	}
}

// NewServiceBazClient creates a new instance of [ServiceBazClient].
func (client *ServiceClient) NewServiceBazClient() *ServiceBazClient {
	return &ServiceBazClient{
		internal: client.internal,
		endpoint: client.endpoint,
		client:   client.client,
	}
}

// NewServiceFooClient creates a new instance of [ServiceFooClient].
func (client *ServiceClient) NewServiceFooClient() *ServiceFooClient {
	return &ServiceFooClient{
		internal: client.internal,
		endpoint: client.endpoint,
		client:   client.client,
	}
}

// NewServiceQuxClient creates a new instance of [ServiceQuxClient].
func (client *ServiceClient) NewServiceQuxClient() *ServiceQuxClient {
	return &ServiceQuxClient{
		internal: client.internal,
		endpoint: client.endpoint,
		client:   client.client,
	}
}

// - options - ServiceClientOneOptions contains the optional parameters for the ServiceClient.One method.
func (client *ServiceClient) One(ctx context.Context, options *ServiceClientOneOptions) (ServiceClientOneResponse, error) {
	var err error
	const operationName = "ServiceClient.One"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.oneCreateRequest(ctx, options)
	if err != nil {
		return ServiceClientOneResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientOneResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientOneResponse{}, err
	}
	return ServiceClientOneResponse{}, nil
}

// oneCreateRequest creates the One request.
func (client *ServiceClient) oneCreateRequest(ctx context.Context, options *ServiceClientOneOptions) (*policy.Request, error) {
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

// - options - ServiceClientTwoOptions contains the optional parameters for the ServiceClient.Two method.
func (client *ServiceClient) Two(ctx context.Context, options *ServiceClientTwoOptions) (ServiceClientTwoResponse, error) {
	var err error
	const operationName = "ServiceClient.Two"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.twoCreateRequest(ctx, options)
	if err != nil {
		return ServiceClientTwoResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientTwoResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientTwoResponse{}, err
	}
	return ServiceClientTwoResponse{}, nil
}

// twoCreateRequest creates the Two request.
func (client *ServiceClient) twoCreateRequest(ctx context.Context, options *ServiceClientTwoOptions) (*policy.Request, error) {
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
