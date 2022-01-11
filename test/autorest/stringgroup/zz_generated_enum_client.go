//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// EnumClient contains the methods for the Enum group.
// Don't use this type directly, use NewEnumClient() instead.
type EnumClient struct {
	pl runtime.Pipeline
}

// NewEnumClient creates a new instance of EnumClient with the specified values.
// options - pass nil to accept the default values.
func NewEnumClient(options *azcore.ClientOptions) *EnumClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &EnumClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// GetNotExpandable - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
// If the operation fails it returns an *azcore.ResponseError type.
// options - EnumClientGetNotExpandableOptions contains the optional parameters for the EnumClient.GetNotExpandable method.
func (client *EnumClient) GetNotExpandable(ctx context.Context, options *EnumClientGetNotExpandableOptions) (EnumClientGetNotExpandableResponse, error) {
	req, err := client.getNotExpandableCreateRequest(ctx, options)
	if err != nil {
		return EnumClientGetNotExpandableResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnumClientGetNotExpandableResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumClientGetNotExpandableResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNotExpandableHandleResponse(resp)
}

// getNotExpandableCreateRequest creates the GetNotExpandable request.
func (client *EnumClient) getNotExpandableCreateRequest(ctx context.Context, options *EnumClientGetNotExpandableOptions) (*policy.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNotExpandableHandleResponse handles the GetNotExpandable response.
func (client *EnumClient) getNotExpandableHandleResponse(resp *http.Response) (EnumClientGetNotExpandableResponse, error) {
	result := EnumClientGetNotExpandableResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return EnumClientGetNotExpandableResponse{}, err
	}
	return result, nil
}

// GetReferenced - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
// If the operation fails it returns an *azcore.ResponseError type.
// options - EnumClientGetReferencedOptions contains the optional parameters for the EnumClient.GetReferenced method.
func (client *EnumClient) GetReferenced(ctx context.Context, options *EnumClientGetReferencedOptions) (EnumClientGetReferencedResponse, error) {
	req, err := client.getReferencedCreateRequest(ctx, options)
	if err != nil {
		return EnumClientGetReferencedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnumClientGetReferencedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumClientGetReferencedResponse{}, runtime.NewResponseError(resp)
	}
	return client.getReferencedHandleResponse(resp)
}

// getReferencedCreateRequest creates the GetReferenced request.
func (client *EnumClient) getReferencedCreateRequest(ctx context.Context, options *EnumClientGetReferencedOptions) (*policy.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getReferencedHandleResponse handles the GetReferenced response.
func (client *EnumClient) getReferencedHandleResponse(resp *http.Response) (EnumClientGetReferencedResponse, error) {
	result := EnumClientGetReferencedResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return EnumClientGetReferencedResponse{}, err
	}
	return result, nil
}

// GetReferencedConstant - Get value 'green-color' from the constant.
// If the operation fails it returns an *azcore.ResponseError type.
// options - EnumClientGetReferencedConstantOptions contains the optional parameters for the EnumClient.GetReferencedConstant
// method.
func (client *EnumClient) GetReferencedConstant(ctx context.Context, options *EnumClientGetReferencedConstantOptions) (EnumClientGetReferencedConstantResponse, error) {
	req, err := client.getReferencedConstantCreateRequest(ctx, options)
	if err != nil {
		return EnumClientGetReferencedConstantResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnumClientGetReferencedConstantResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumClientGetReferencedConstantResponse{}, runtime.NewResponseError(resp)
	}
	return client.getReferencedConstantHandleResponse(resp)
}

// getReferencedConstantCreateRequest creates the GetReferencedConstant request.
func (client *EnumClient) getReferencedConstantCreateRequest(ctx context.Context, options *EnumClientGetReferencedConstantOptions) (*policy.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getReferencedConstantHandleResponse handles the GetReferencedConstant response.
func (client *EnumClient) getReferencedConstantHandleResponse(resp *http.Response) (EnumClientGetReferencedConstantResponse, error) {
	result := EnumClientGetReferencedConstantResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RefColorConstant); err != nil {
		return EnumClientGetReferencedConstantResponse{}, err
	}
	return result, nil
}

// PutNotExpandable - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
// If the operation fails it returns an *azcore.ResponseError type.
// stringBody - string body
// options - EnumClientPutNotExpandableOptions contains the optional parameters for the EnumClient.PutNotExpandable method.
func (client *EnumClient) PutNotExpandable(ctx context.Context, stringBody Colors, options *EnumClientPutNotExpandableOptions) (EnumClientPutNotExpandableResponse, error) {
	req, err := client.putNotExpandableCreateRequest(ctx, stringBody, options)
	if err != nil {
		return EnumClientPutNotExpandableResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnumClientPutNotExpandableResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumClientPutNotExpandableResponse{}, runtime.NewResponseError(resp)
	}
	return EnumClientPutNotExpandableResponse{RawResponse: resp}, nil
}

// putNotExpandableCreateRequest creates the PutNotExpandable request.
func (client *EnumClient) putNotExpandableCreateRequest(ctx context.Context, stringBody Colors, options *EnumClientPutNotExpandableOptions) (*policy.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, stringBody)
}

// PutReferenced - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
// If the operation fails it returns an *azcore.ResponseError type.
// enumStringBody - enum string body
// options - EnumClientPutReferencedOptions contains the optional parameters for the EnumClient.PutReferenced method.
func (client *EnumClient) PutReferenced(ctx context.Context, enumStringBody Colors, options *EnumClientPutReferencedOptions) (EnumClientPutReferencedResponse, error) {
	req, err := client.putReferencedCreateRequest(ctx, enumStringBody, options)
	if err != nil {
		return EnumClientPutReferencedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnumClientPutReferencedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumClientPutReferencedResponse{}, runtime.NewResponseError(resp)
	}
	return EnumClientPutReferencedResponse{RawResponse: resp}, nil
}

// putReferencedCreateRequest creates the PutReferenced request.
func (client *EnumClient) putReferencedCreateRequest(ctx context.Context, enumStringBody Colors, options *EnumClientPutReferencedOptions) (*policy.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, enumStringBody)
}

// PutReferencedConstant - Sends value 'green-color' from a constant
// If the operation fails it returns an *azcore.ResponseError type.
// enumStringBody - enum string body
// options - EnumClientPutReferencedConstantOptions contains the optional parameters for the EnumClient.PutReferencedConstant
// method.
func (client *EnumClient) PutReferencedConstant(ctx context.Context, enumStringBody RefColorConstant, options *EnumClientPutReferencedConstantOptions) (EnumClientPutReferencedConstantResponse, error) {
	req, err := client.putReferencedConstantCreateRequest(ctx, enumStringBody, options)
	if err != nil {
		return EnumClientPutReferencedConstantResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnumClientPutReferencedConstantResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumClientPutReferencedConstantResponse{}, runtime.NewResponseError(resp)
	}
	return EnumClientPutReferencedConstantResponse{RawResponse: resp}, nil
}

// putReferencedConstantCreateRequest creates the PutReferencedConstant request.
func (client *EnumClient) putReferencedConstantCreateRequest(ctx context.Context, enumStringBody RefColorConstant, options *EnumClientPutReferencedConstantOptions) (*policy.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, enumStringBody)
}
