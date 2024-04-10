// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package enumdiscgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// EnumDiscriminatorClient - Illustrates inheritance with enum discriminator.
// Don't use this type directly, use NewEnumDiscriminatorClientWithNoCredential() instead.
type EnumDiscriminatorClient struct {
	internal *azcore.Client
}

// EnumDiscriminatorClientOptions contains the optional values for creating a [EnumDiscriminatorClient].
type EnumDiscriminatorClientOptions struct {
	azcore.ClientOptions
}

// NewEnumDiscriminatorClientWithNoCredential creates a new instance of [EnumDiscriminatorClient] with the specified values.
//   - options - EnumDiscriminatorClientOptions contains the optional values for creating a [EnumDiscriminatorClient]
func NewEnumDiscriminatorClientWithNoCredential(options *EnumDiscriminatorClientOptions) (*EnumDiscriminatorClient, error) {
	if options == nil {
		options = &EnumDiscriminatorClientOptions{}
	}
	cl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	client := &EnumDiscriminatorClient{
		internal: cl,
	}
	return client, nil
}

// GetExtensibleModel - Receive model with extensible enum discriminator type.
//   - options - EnumDiscriminatorClientGetExtensibleModelOptions contains the optional parameters for the EnumDiscriminatorClient.GetExtensibleModel
//     method.
func (client *EnumDiscriminatorClient) GetExtensibleModel(ctx context.Context, options *EnumDiscriminatorClientGetExtensibleModelOptions) (EnumDiscriminatorClientGetExtensibleModelResponse, error) {
	var err error
	const operationName = "EnumDiscriminatorClient.GetExtensibleModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getExtensibleModelCreateRequest(ctx, options)
	if err != nil {
		return EnumDiscriminatorClientGetExtensibleModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumDiscriminatorClientGetExtensibleModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumDiscriminatorClientGetExtensibleModelResponse{}, err
	}
	resp, err := client.getExtensibleModelHandleResponse(httpResp)
	return resp, err
}

// getExtensibleModelCreateRequest creates the GetExtensibleModel request.
func (client *EnumDiscriminatorClient) getExtensibleModelCreateRequest(ctx context.Context, options *EnumDiscriminatorClientGetExtensibleModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/enum-discriminator/extensible-enum"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getExtensibleModelHandleResponse handles the GetExtensibleModel response.
func (client *EnumDiscriminatorClient) getExtensibleModelHandleResponse(resp *http.Response) (EnumDiscriminatorClientGetExtensibleModelResponse, error) {
	result := EnumDiscriminatorClientGetExtensibleModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EnumDiscriminatorClientGetExtensibleModelResponse{}, err
	}
	return result, nil
}

// GetExtensibleModelMissingDiscriminator - Get a model omitting the discriminator.
//   - options - EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorOptions contains the optional parameters for the
//     EnumDiscriminatorClient.GetExtensibleModelMissingDiscriminator method.
func (client *EnumDiscriminatorClient) GetExtensibleModelMissingDiscriminator(ctx context.Context, options *EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorOptions) (EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorResponse, error) {
	var err error
	const operationName = "EnumDiscriminatorClient.GetExtensibleModelMissingDiscriminator"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getExtensibleModelMissingDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorResponse{}, err
	}
	resp, err := client.getExtensibleModelMissingDiscriminatorHandleResponse(httpResp)
	return resp, err
}

// getExtensibleModelMissingDiscriminatorCreateRequest creates the GetExtensibleModelMissingDiscriminator request.
func (client *EnumDiscriminatorClient) getExtensibleModelMissingDiscriminatorCreateRequest(ctx context.Context, options *EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/enum-discriminator/extensible-enum/missingdiscriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getExtensibleModelMissingDiscriminatorHandleResponse handles the GetExtensibleModelMissingDiscriminator response.
func (client *EnumDiscriminatorClient) getExtensibleModelMissingDiscriminatorHandleResponse(resp *http.Response) (EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorResponse, error) {
	result := EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorResponse{}, err
	}
	return result, nil
}

// GetExtensibleModelWrongDiscriminator - Get a model containing discriminator value never defined.
//   - options - EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorOptions contains the optional parameters for the EnumDiscriminatorClient.GetExtensibleModelWrongDiscriminator
//     method.
func (client *EnumDiscriminatorClient) GetExtensibleModelWrongDiscriminator(ctx context.Context, options *EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorOptions) (EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorResponse, error) {
	var err error
	const operationName = "EnumDiscriminatorClient.GetExtensibleModelWrongDiscriminator"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getExtensibleModelWrongDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorResponse{}, err
	}
	resp, err := client.getExtensibleModelWrongDiscriminatorHandleResponse(httpResp)
	return resp, err
}

// getExtensibleModelWrongDiscriminatorCreateRequest creates the GetExtensibleModelWrongDiscriminator request.
func (client *EnumDiscriminatorClient) getExtensibleModelWrongDiscriminatorCreateRequest(ctx context.Context, options *EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/enum-discriminator/extensible-enum/wrongdiscriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getExtensibleModelWrongDiscriminatorHandleResponse handles the GetExtensibleModelWrongDiscriminator response.
func (client *EnumDiscriminatorClient) getExtensibleModelWrongDiscriminatorHandleResponse(resp *http.Response) (EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorResponse, error) {
	result := EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorResponse{}, err
	}
	return result, nil
}

// GetFixedModel - Receive model with fixed enum discriminator type.
//   - options - EnumDiscriminatorClientGetFixedModelOptions contains the optional parameters for the EnumDiscriminatorClient.GetFixedModel
//     method.
func (client *EnumDiscriminatorClient) GetFixedModel(ctx context.Context, options *EnumDiscriminatorClientGetFixedModelOptions) (EnumDiscriminatorClientGetFixedModelResponse, error) {
	var err error
	const operationName = "EnumDiscriminatorClient.GetFixedModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getFixedModelCreateRequest(ctx, options)
	if err != nil {
		return EnumDiscriminatorClientGetFixedModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumDiscriminatorClientGetFixedModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumDiscriminatorClientGetFixedModelResponse{}, err
	}
	resp, err := client.getFixedModelHandleResponse(httpResp)
	return resp, err
}

// getFixedModelCreateRequest creates the GetFixedModel request.
func (client *EnumDiscriminatorClient) getFixedModelCreateRequest(ctx context.Context, options *EnumDiscriminatorClientGetFixedModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/enum-discriminator/fixed-enum"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFixedModelHandleResponse handles the GetFixedModel response.
func (client *EnumDiscriminatorClient) getFixedModelHandleResponse(resp *http.Response) (EnumDiscriminatorClientGetFixedModelResponse, error) {
	result := EnumDiscriminatorClientGetFixedModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EnumDiscriminatorClientGetFixedModelResponse{}, err
	}
	return result, nil
}

// GetFixedModelMissingDiscriminator - Get a model omitting the discriminator.
//   - options - EnumDiscriminatorClientGetFixedModelMissingDiscriminatorOptions contains the optional parameters for the EnumDiscriminatorClient.GetFixedModelMissingDiscriminator
//     method.
func (client *EnumDiscriminatorClient) GetFixedModelMissingDiscriminator(ctx context.Context, options *EnumDiscriminatorClientGetFixedModelMissingDiscriminatorOptions) (EnumDiscriminatorClientGetFixedModelMissingDiscriminatorResponse, error) {
	var err error
	const operationName = "EnumDiscriminatorClient.GetFixedModelMissingDiscriminator"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getFixedModelMissingDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return EnumDiscriminatorClientGetFixedModelMissingDiscriminatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumDiscriminatorClientGetFixedModelMissingDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumDiscriminatorClientGetFixedModelMissingDiscriminatorResponse{}, err
	}
	resp, err := client.getFixedModelMissingDiscriminatorHandleResponse(httpResp)
	return resp, err
}

// getFixedModelMissingDiscriminatorCreateRequest creates the GetFixedModelMissingDiscriminator request.
func (client *EnumDiscriminatorClient) getFixedModelMissingDiscriminatorCreateRequest(ctx context.Context, options *EnumDiscriminatorClientGetFixedModelMissingDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/enum-discriminator/fixed-enum/missingdiscriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFixedModelMissingDiscriminatorHandleResponse handles the GetFixedModelMissingDiscriminator response.
func (client *EnumDiscriminatorClient) getFixedModelMissingDiscriminatorHandleResponse(resp *http.Response) (EnumDiscriminatorClientGetFixedModelMissingDiscriminatorResponse, error) {
	result := EnumDiscriminatorClientGetFixedModelMissingDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EnumDiscriminatorClientGetFixedModelMissingDiscriminatorResponse{}, err
	}
	return result, nil
}

// GetFixedModelWrongDiscriminator - Get a model containing discriminator value never defined.
//   - options - EnumDiscriminatorClientGetFixedModelWrongDiscriminatorOptions contains the optional parameters for the EnumDiscriminatorClient.GetFixedModelWrongDiscriminator
//     method.
func (client *EnumDiscriminatorClient) GetFixedModelWrongDiscriminator(ctx context.Context, options *EnumDiscriminatorClientGetFixedModelWrongDiscriminatorOptions) (EnumDiscriminatorClientGetFixedModelWrongDiscriminatorResponse, error) {
	var err error
	const operationName = "EnumDiscriminatorClient.GetFixedModelWrongDiscriminator"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getFixedModelWrongDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return EnumDiscriminatorClientGetFixedModelWrongDiscriminatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumDiscriminatorClientGetFixedModelWrongDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumDiscriminatorClientGetFixedModelWrongDiscriminatorResponse{}, err
	}
	resp, err := client.getFixedModelWrongDiscriminatorHandleResponse(httpResp)
	return resp, err
}

// getFixedModelWrongDiscriminatorCreateRequest creates the GetFixedModelWrongDiscriminator request.
func (client *EnumDiscriminatorClient) getFixedModelWrongDiscriminatorCreateRequest(ctx context.Context, options *EnumDiscriminatorClientGetFixedModelWrongDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/enum-discriminator/fixed-enum/wrongdiscriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFixedModelWrongDiscriminatorHandleResponse handles the GetFixedModelWrongDiscriminator response.
func (client *EnumDiscriminatorClient) getFixedModelWrongDiscriminatorHandleResponse(resp *http.Response) (EnumDiscriminatorClientGetFixedModelWrongDiscriminatorResponse, error) {
	result := EnumDiscriminatorClientGetFixedModelWrongDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EnumDiscriminatorClientGetFixedModelWrongDiscriminatorResponse{}, err
	}
	return result, nil
}

// PutExtensibleModel - Send model with extensible enum discriminator type.
//   - input - Dog to create
//   - options - EnumDiscriminatorClientPutExtensibleModelOptions contains the optional parameters for the EnumDiscriminatorClient.PutExtensibleModel
//     method.
func (client *EnumDiscriminatorClient) PutExtensibleModel(ctx context.Context, input DogClassification, options *EnumDiscriminatorClientPutExtensibleModelOptions) (EnumDiscriminatorClientPutExtensibleModelResponse, error) {
	var err error
	const operationName = "EnumDiscriminatorClient.PutExtensibleModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putExtensibleModelCreateRequest(ctx, input, options)
	if err != nil {
		return EnumDiscriminatorClientPutExtensibleModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumDiscriminatorClientPutExtensibleModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return EnumDiscriminatorClientPutExtensibleModelResponse{}, err
	}
	return EnumDiscriminatorClientPutExtensibleModelResponse{}, nil
}

// putExtensibleModelCreateRequest creates the PutExtensibleModel request.
func (client *EnumDiscriminatorClient) putExtensibleModelCreateRequest(ctx context.Context, input DogClassification, options *EnumDiscriminatorClientPutExtensibleModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/enum-discriminator/extensible-enum"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// PutFixedModel - Send model with fixed enum discriminator type.
//   - input - Snake to create
//   - options - EnumDiscriminatorClientPutFixedModelOptions contains the optional parameters for the EnumDiscriminatorClient.PutFixedModel
//     method.
func (client *EnumDiscriminatorClient) PutFixedModel(ctx context.Context, input SnakeClassification, options *EnumDiscriminatorClientPutFixedModelOptions) (EnumDiscriminatorClientPutFixedModelResponse, error) {
	var err error
	const operationName = "EnumDiscriminatorClient.PutFixedModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putFixedModelCreateRequest(ctx, input, options)
	if err != nil {
		return EnumDiscriminatorClientPutFixedModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumDiscriminatorClientPutFixedModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return EnumDiscriminatorClientPutFixedModelResponse{}, err
	}
	return EnumDiscriminatorClientPutFixedModelResponse{}, nil
}

// putFixedModelCreateRequest creates the PutFixedModel request.
func (client *EnumDiscriminatorClient) putFixedModelCreateRequest(ctx context.Context, input SnakeClassification, options *EnumDiscriminatorClientPutFixedModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/enum-discriminator/fixed-enum"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}
