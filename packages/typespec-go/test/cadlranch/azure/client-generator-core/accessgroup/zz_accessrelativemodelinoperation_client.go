// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package accessgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AccessRelativeModelInOperationClient contains the methods for the AccessRelativeModelInOperation group.
// Don't use this type directly, use [AccessClient.NewAccessRelativeModelInOperationClient] instead.
type AccessRelativeModelInOperationClient struct {
	internal *azcore.Client
}

// discriminator - Expected query parameter: kind=<any string>
// Expected response body:
// ```json
// {
// "name": <any string>,
// "kind": "real"
// }
// ```
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - accessRelativeModelInOperationClientdiscriminatorOptions contains the optional parameters for the AccessRelativeModelInOperationClient.discriminator
//     method.
func (client *AccessRelativeModelInOperationClient) discriminator(ctx context.Context, kind string, options *accessRelativeModelInOperationClientdiscriminatorOptions) (accessRelativeModelInOperationClientdiscriminatorResponse, error) {
	var err error
	const operationName = "AccessRelativeModelInOperationClient.discriminator"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.discriminatorCreateRequest(ctx, kind, options)
	if err != nil {
		return accessRelativeModelInOperationClientdiscriminatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return accessRelativeModelInOperationClientdiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return accessRelativeModelInOperationClientdiscriminatorResponse{}, err
	}
	resp, err := client.discriminatorHandleResponse(httpResp)
	return resp, err
}

// discriminatorCreateRequest creates the discriminator request.
func (client *AccessRelativeModelInOperationClient) discriminatorCreateRequest(ctx context.Context, kind string, _ *accessRelativeModelInOperationClientdiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/relativeModelInOperation/discriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("kind", kind)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// discriminatorHandleResponse handles the discriminator response.
func (client *AccessRelativeModelInOperationClient) discriminatorHandleResponse(resp *http.Response) (accessRelativeModelInOperationClientdiscriminatorResponse, error) {
	result := accessRelativeModelInOperationClientdiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return accessRelativeModelInOperationClientdiscriminatorResponse{}, err
	}
	return result, nil
}

// operation - Expected query parameter: name=<any string>
// Expected response body:
// ```json
// {
// "name": <any string>,
// "inner":
// {
// "name": <any string>
// }
// }
// ```
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - accessRelativeModelInOperationClientoperationOptions contains the optional parameters for the AccessRelativeModelInOperationClient.operation
//     method.
func (client *AccessRelativeModelInOperationClient) operation(ctx context.Context, name string, options *accessRelativeModelInOperationClientoperationOptions) (accessRelativeModelInOperationClientoperationResponse, error) {
	var err error
	const operationName = "AccessRelativeModelInOperationClient.operation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.operationCreateRequest(ctx, name, options)
	if err != nil {
		return accessRelativeModelInOperationClientoperationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return accessRelativeModelInOperationClientoperationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return accessRelativeModelInOperationClientoperationResponse{}, err
	}
	resp, err := client.operationHandleResponse(httpResp)
	return resp, err
}

// operationCreateRequest creates the operation request.
func (client *AccessRelativeModelInOperationClient) operationCreateRequest(ctx context.Context, name string, _ *accessRelativeModelInOperationClientoperationOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/relativeModelInOperation/operation"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("name", name)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// operationHandleResponse handles the operation response.
func (client *AccessRelativeModelInOperationClient) operationHandleResponse(resp *http.Response) (accessRelativeModelInOperationClientoperationResponse, error) {
	result := accessRelativeModelInOperationClientoperationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.outerModel); err != nil {
		return accessRelativeModelInOperationClientoperationResponse{}, err
	}
	return result, nil
}
