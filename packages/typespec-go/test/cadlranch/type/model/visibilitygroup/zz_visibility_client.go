// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package visibilitygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// VisibilityClient - Illustrates models with visibility properties.
// Don't use this type directly, use a constructor function instead.
type VisibilityClient struct {
	internal *azcore.Client
}

// VisibilityClientOptions contains the optional values for creating a [VisibilityClient].
type VisibilityClientOptions struct {
	azcore.ClientOptions
}

// NewVisibilityClientWithNoCredential creates a new [VisibilityClient].
//   - options - optional client configuration; pass nil to accept the default values
func NewVisibilityClientWithNoCredential(options *VisibilityClientOptions) (*VisibilityClient, error) {
	if options == nil {
		options = &VisibilityClientOptions{}
	}
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	return &VisibilityClient{
		internal: internal,
	}, nil
}

// - options - VisibilityClientDeleteModelOptions contains the optional parameters for the VisibilityClient.DeleteModel method.
func (client *VisibilityClient) DeleteModel(ctx context.Context, input VisibilityModel, options *VisibilityClientDeleteModelOptions) (VisibilityClientDeleteModelResponse, error) {
	var err error
	const operationName = "VisibilityClient.DeleteModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteModelCreateRequest(ctx, input, options)
	if err != nil {
		return VisibilityClientDeleteModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VisibilityClientDeleteModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return VisibilityClientDeleteModelResponse{}, err
	}
	return VisibilityClientDeleteModelResponse{}, nil
}

// deleteModelCreateRequest creates the DeleteModel request.
func (client *VisibilityClient) deleteModelCreateRequest(ctx context.Context, input VisibilityModel, options *VisibilityClientDeleteModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/visibility"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// - options - VisibilityClientGetModelOptions contains the optional parameters for the VisibilityClient.GetModel method.
func (client *VisibilityClient) GetModel(ctx context.Context, input VisibilityModel, options *VisibilityClientGetModelOptions) (VisibilityClientGetModelResponse, error) {
	var err error
	const operationName = "VisibilityClient.GetModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getModelCreateRequest(ctx, input, options)
	if err != nil {
		return VisibilityClientGetModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VisibilityClientGetModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VisibilityClientGetModelResponse{}, err
	}
	resp, err := client.getModelHandleResponse(httpResp)
	return resp, err
}

// getModelCreateRequest creates the GetModel request.
func (client *VisibilityClient) getModelCreateRequest(ctx context.Context, input VisibilityModel, options *VisibilityClientGetModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/visibility"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// getModelHandleResponse handles the GetModel response.
func (client *VisibilityClient) getModelHandleResponse(resp *http.Response) (VisibilityClientGetModelResponse, error) {
	result := VisibilityClientGetModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VisibilityModel); err != nil {
		return VisibilityClientGetModelResponse{}, err
	}
	return result, nil
}

// - options - VisibilityClientHeadModelOptions contains the optional parameters for the VisibilityClient.HeadModel method.
func (client *VisibilityClient) HeadModel(ctx context.Context, input VisibilityModel, options *VisibilityClientHeadModelOptions) (VisibilityClientHeadModelResponse, error) {
	var err error
	const operationName = "VisibilityClient.HeadModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.headModelCreateRequest(ctx, input, options)
	if err != nil {
		return VisibilityClientHeadModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VisibilityClientHeadModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VisibilityClientHeadModelResponse{}, err
	}
	return VisibilityClientHeadModelResponse{}, nil
}

// headModelCreateRequest creates the HeadModel request.
func (client *VisibilityClient) headModelCreateRequest(ctx context.Context, input VisibilityModel, options *VisibilityClientHeadModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/visibility"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// - options - VisibilityClientPatchModelOptions contains the optional parameters for the VisibilityClient.PatchModel method.
func (client *VisibilityClient) PatchModel(ctx context.Context, input VisibilityModel, options *VisibilityClientPatchModelOptions) (VisibilityClientPatchModelResponse, error) {
	var err error
	const operationName = "VisibilityClient.PatchModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patchModelCreateRequest(ctx, input, options)
	if err != nil {
		return VisibilityClientPatchModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VisibilityClientPatchModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return VisibilityClientPatchModelResponse{}, err
	}
	return VisibilityClientPatchModelResponse{}, nil
}

// patchModelCreateRequest creates the PatchModel request.
func (client *VisibilityClient) patchModelCreateRequest(ctx context.Context, input VisibilityModel, options *VisibilityClientPatchModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/visibility"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// - options - VisibilityClientPostModelOptions contains the optional parameters for the VisibilityClient.PostModel method.
func (client *VisibilityClient) PostModel(ctx context.Context, input VisibilityModel, options *VisibilityClientPostModelOptions) (VisibilityClientPostModelResponse, error) {
	var err error
	const operationName = "VisibilityClient.PostModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postModelCreateRequest(ctx, input, options)
	if err != nil {
		return VisibilityClientPostModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VisibilityClientPostModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return VisibilityClientPostModelResponse{}, err
	}
	return VisibilityClientPostModelResponse{}, nil
}

// postModelCreateRequest creates the PostModel request.
func (client *VisibilityClient) postModelCreateRequest(ctx context.Context, input VisibilityModel, options *VisibilityClientPostModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/visibility"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// - options - VisibilityClientPutModelOptions contains the optional parameters for the VisibilityClient.PutModel method.
func (client *VisibilityClient) PutModel(ctx context.Context, input VisibilityModel, options *VisibilityClientPutModelOptions) (VisibilityClientPutModelResponse, error) {
	var err error
	const operationName = "VisibilityClient.PutModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putModelCreateRequest(ctx, input, options)
	if err != nil {
		return VisibilityClientPutModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VisibilityClientPutModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return VisibilityClientPutModelResponse{}, err
	}
	return VisibilityClientPutModelResponse{}, nil
}

// putModelCreateRequest creates the PutModel request.
func (client *VisibilityClient) putModelCreateRequest(ctx context.Context, input VisibilityModel, options *VisibilityClientPutModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/visibility"
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
