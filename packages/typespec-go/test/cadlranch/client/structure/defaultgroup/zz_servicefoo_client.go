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

// ServiceFooClient contains the methods for the ServiceFoo group.
// Don't use this type directly, use [ServiceClient.NewServiceFooClient] instead.
type ServiceFooClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// Four -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ServiceFooClientFourOptions contains the optional parameters for the ServiceFooClient.Four method.
func (client *ServiceFooClient) Four(ctx context.Context, options *ServiceFooClientFourOptions) (ServiceFooClientFourResponse, error) {
	var err error
	const operationName = "ServiceFooClient.Four"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fourCreateRequest(ctx, options)
	if err != nil {
		return ServiceFooClientFourResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceFooClientFourResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceFooClientFourResponse{}, err
	}
	return ServiceFooClientFourResponse{}, nil
}

// fourCreateRequest creates the Four request.
func (client *ServiceFooClient) fourCreateRequest(ctx context.Context, _ *ServiceFooClientFourOptions) (*policy.Request, error) {
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

// Three -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ServiceFooClientThreeOptions contains the optional parameters for the ServiceFooClient.Three method.
func (client *ServiceFooClient) Three(ctx context.Context, options *ServiceFooClientThreeOptions) (ServiceFooClientThreeResponse, error) {
	var err error
	const operationName = "ServiceFooClient.Three"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.threeCreateRequest(ctx, options)
	if err != nil {
		return ServiceFooClientThreeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceFooClientThreeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceFooClientThreeResponse{}, err
	}
	return ServiceFooClientThreeResponse{}, nil
}

// threeCreateRequest creates the Three request.
func (client *ServiceFooClient) threeCreateRequest(ctx context.Context, _ *ServiceFooClientThreeOptions) (*policy.Request, error) {
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
