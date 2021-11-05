//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ArrayClient contains the methods for the Array group.
// Don't use this type directly, use NewArrayClient() instead.
type ArrayClient struct {
	pl runtime.Pipeline
}

// NewArrayClient creates a new instance of ArrayClient with the specified values.
// options - pass nil to accept the default values.
func NewArrayClient(options *azcore.ClientOptions) *ArrayClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &ArrayClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// GetEmpty - Get complex types with array property which is empty
// If the operation fails it returns the *Error error type.
// options - ArrayGetEmptyOptions contains the optional parameters for the Array.GetEmpty method.
func (client *ArrayClient) GetEmpty(ctx context.Context, options *ArrayGetEmptyOptions) (ArrayGetEmptyResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return ArrayGetEmptyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArrayGetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArrayGetEmptyResponse{}, client.getEmptyHandleError(resp)
	}
	return client.getEmptyHandleResponse(resp)
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *ArrayClient) getEmptyCreateRequest(ctx context.Context, options *ArrayGetEmptyOptions) (*policy.Request, error) {
	urlPath := "/complex/array/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *ArrayClient) getEmptyHandleResponse(resp *http.Response) (ArrayGetEmptyResponse, error) {
	result := ArrayGetEmptyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArrayWrapper); err != nil {
		return ArrayGetEmptyResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *ArrayClient) getEmptyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetNotProvided - Get complex types with array property while server doesn't provide a response payload
// If the operation fails it returns the *Error error type.
// options - ArrayGetNotProvidedOptions contains the optional parameters for the Array.GetNotProvided method.
func (client *ArrayClient) GetNotProvided(ctx context.Context, options *ArrayGetNotProvidedOptions) (ArrayGetNotProvidedResponse, error) {
	req, err := client.getNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return ArrayGetNotProvidedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArrayGetNotProvidedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArrayGetNotProvidedResponse{}, client.getNotProvidedHandleError(resp)
	}
	return client.getNotProvidedHandleResponse(resp)
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *ArrayClient) getNotProvidedCreateRequest(ctx context.Context, options *ArrayGetNotProvidedOptions) (*policy.Request, error) {
	urlPath := "/complex/array/notprovided"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *ArrayClient) getNotProvidedHandleResponse(resp *http.Response) (ArrayGetNotProvidedResponse, error) {
	result := ArrayGetNotProvidedResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArrayWrapper); err != nil {
		return ArrayGetNotProvidedResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getNotProvidedHandleError handles the GetNotProvided error response.
func (client *ArrayClient) getNotProvidedHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetValid - Get complex types with array property
// If the operation fails it returns the *Error error type.
// options - ArrayGetValidOptions contains the optional parameters for the Array.GetValid method.
func (client *ArrayClient) GetValid(ctx context.Context, options *ArrayGetValidOptions) (ArrayGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return ArrayGetValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArrayGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArrayGetValidResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *ArrayClient) getValidCreateRequest(ctx context.Context, options *ArrayGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/array/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *ArrayClient) getValidHandleResponse(resp *http.Response) (ArrayGetValidResponse, error) {
	result := ArrayGetValidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArrayWrapper); err != nil {
		return ArrayGetValidResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getValidHandleError handles the GetValid error response.
func (client *ArrayClient) getValidHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// PutEmpty - Put complex types with array property which is empty
// If the operation fails it returns the *Error error type.
// complexBody - Please put an empty array
// options - ArrayPutEmptyOptions contains the optional parameters for the Array.PutEmpty method.
func (client *ArrayClient) PutEmpty(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutEmptyOptions) (ArrayPutEmptyResponse, error) {
	req, err := client.putEmptyCreateRequest(ctx, complexBody, options)
	if err != nil {
		return ArrayPutEmptyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArrayPutEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArrayPutEmptyResponse{}, client.putEmptyHandleError(resp)
	}
	return ArrayPutEmptyResponse{RawResponse: resp}, nil
}

// putEmptyCreateRequest creates the PutEmpty request.
func (client *ArrayClient) putEmptyCreateRequest(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutEmptyOptions) (*policy.Request, error) {
	urlPath := "/complex/array/empty"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// putEmptyHandleError handles the PutEmpty error response.
func (client *ArrayClient) putEmptyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// PutValid - Put complex types with array property
// If the operation fails it returns the *Error error type.
// complexBody - Please put an array with 4 items: "1, 2, 3, 4", "", null, "&S#$(*Y", "The quick brown fox jumps over the
// lazy dog"
// options - ArrayPutValidOptions contains the optional parameters for the Array.PutValid method.
func (client *ArrayClient) PutValid(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutValidOptions) (ArrayPutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return ArrayPutValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArrayPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArrayPutValidResponse{}, client.putValidHandleError(resp)
	}
	return ArrayPutValidResponse{RawResponse: resp}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *ArrayClient) putValidCreateRequest(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/array/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *ArrayClient) putValidHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
