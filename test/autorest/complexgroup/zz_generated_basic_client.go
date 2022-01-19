//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BasicClient contains the methods for the Basic group.
// Don't use this type directly, use NewBasicClient() instead.
type BasicClient struct {
	pl runtime.Pipeline
}

// NewBasicClient creates a new instance of BasicClient with the specified values.
// options - pass nil to accept the default values.
func NewBasicClient(options *azcore.ClientOptions) *BasicClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &BasicClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetEmpty - Get a basic complex type that is empty
// If the operation fails it returns an *azcore.ResponseError type.
// options - BasicClientGetEmptyOptions contains the optional parameters for the BasicClient.GetEmpty method.
func (client *BasicClient) GetEmpty(ctx context.Context, options *BasicClientGetEmptyOptions) (BasicClientGetEmptyResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return BasicClientGetEmptyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BasicClientGetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicClientGetEmptyResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEmptyHandleResponse(resp)
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *BasicClient) getEmptyCreateRequest(ctx context.Context, options *BasicClientGetEmptyOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *BasicClient) getEmptyHandleResponse(resp *http.Response) (BasicClientGetEmptyResponse, error) {
	result := BasicClientGetEmptyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicClientGetEmptyResponse{}, err
	}
	return result, nil
}

// GetInvalid - Get a basic complex type that is invalid for the local strong type
// If the operation fails it returns an *azcore.ResponseError type.
// options - BasicClientGetInvalidOptions contains the optional parameters for the BasicClient.GetInvalid method.
func (client *BasicClient) GetInvalid(ctx context.Context, options *BasicClientGetInvalidOptions) (BasicClientGetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return BasicClientGetInvalidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BasicClientGetInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicClientGetInvalidResponse{}, runtime.NewResponseError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *BasicClient) getInvalidCreateRequest(ctx context.Context, options *BasicClientGetInvalidOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *BasicClient) getInvalidHandleResponse(resp *http.Response) (BasicClientGetInvalidResponse, error) {
	result := BasicClientGetInvalidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicClientGetInvalidResponse{}, err
	}
	return result, nil
}

// GetNotProvided - Get a basic complex type while the server doesn't provide a response payload
// If the operation fails it returns an *azcore.ResponseError type.
// options - BasicClientGetNotProvidedOptions contains the optional parameters for the BasicClient.GetNotProvided method.
func (client *BasicClient) GetNotProvided(ctx context.Context, options *BasicClientGetNotProvidedOptions) (BasicClientGetNotProvidedResponse, error) {
	req, err := client.getNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return BasicClientGetNotProvidedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BasicClientGetNotProvidedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicClientGetNotProvidedResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNotProvidedHandleResponse(resp)
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *BasicClient) getNotProvidedCreateRequest(ctx context.Context, options *BasicClientGetNotProvidedOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/notprovided"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *BasicClient) getNotProvidedHandleResponse(resp *http.Response) (BasicClientGetNotProvidedResponse, error) {
	result := BasicClientGetNotProvidedResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicClientGetNotProvidedResponse{}, err
	}
	return result, nil
}

// GetNull - Get a basic complex type whose properties are null
// If the operation fails it returns an *azcore.ResponseError type.
// options - BasicClientGetNullOptions contains the optional parameters for the BasicClient.GetNull method.
func (client *BasicClient) GetNull(ctx context.Context, options *BasicClientGetNullOptions) (BasicClientGetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return BasicClientGetNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BasicClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicClientGetNullResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *BasicClient) getNullCreateRequest(ctx context.Context, options *BasicClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *BasicClient) getNullHandleResponse(resp *http.Response) (BasicClientGetNullResponse, error) {
	result := BasicClientGetNullResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicClientGetNullResponse{}, err
	}
	return result, nil
}

// GetValid - Get complex type {id: 2, name: 'abc', color: 'YELLOW'}
// If the operation fails it returns an *azcore.ResponseError type.
// options - BasicClientGetValidOptions contains the optional parameters for the BasicClient.GetValid method.
func (client *BasicClient) GetValid(ctx context.Context, options *BasicClientGetValidOptions) (BasicClientGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return BasicClientGetValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BasicClientGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicClientGetValidResponse{}, runtime.NewResponseError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *BasicClient) getValidCreateRequest(ctx context.Context, options *BasicClientGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *BasicClient) getValidHandleResponse(resp *http.Response) (BasicClientGetValidResponse, error) {
	result := BasicClientGetValidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicClientGetValidResponse{}, err
	}
	return result, nil
}

// PutValid - Please put {id: 2, name: 'abc', color: 'Magenta'}
// If the operation fails it returns an *azcore.ResponseError type.
// complexBody - Please put {id: 2, name: 'abc', color: 'Magenta'}
// options - BasicClientPutValidOptions contains the optional parameters for the BasicClient.PutValid method.
func (client *BasicClient) PutValid(ctx context.Context, complexBody Basic, options *BasicClientPutValidOptions) (BasicClientPutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return BasicClientPutValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BasicClientPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicClientPutValidResponse{}, runtime.NewResponseError(resp)
	}
	return BasicClientPutValidResponse{RawResponse: resp}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *BasicClient) putValidCreateRequest(ctx context.Context, complexBody Basic, options *BasicClientPutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-02-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, complexBody)
}
