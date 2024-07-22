// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package spreadgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SpreadAliasClient contains the methods for the SpreadAlias group.
// Don't use this type directly, use [SpreadClient.NewSpreadAliasClient] instead.
type SpreadAliasClient struct {
	internal *azcore.Client
}

// SpreadAsRequestBody -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - SpreadAliasClientSpreadAsRequestBodyOptions contains the optional parameters for the SpreadAliasClient.SpreadAsRequestBody
//     method.
func (client *SpreadAliasClient) SpreadAsRequestBody(ctx context.Context, name string, options *SpreadAliasClientSpreadAsRequestBodyOptions) (SpreadAliasClientSpreadAsRequestBodyResponse, error) {
	var err error
	const operationName = "SpreadAliasClient.SpreadAsRequestBody"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadAsRequestBodyCreateRequest(ctx, name, options)
	if err != nil {
		return SpreadAliasClientSpreadAsRequestBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadAliasClientSpreadAsRequestBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadAliasClientSpreadAsRequestBodyResponse{}, err
	}
	return SpreadAliasClientSpreadAsRequestBodyResponse{}, nil
}

// spreadAsRequestBodyCreateRequest creates the SpreadAsRequestBody request.
func (client *SpreadAliasClient) spreadAsRequestBodyCreateRequest(ctx context.Context, name string, _ *SpreadAliasClientSpreadAsRequestBodyOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/alias/request-body"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	body := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// SpreadAsRequestParameter -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - SpreadAliasClientSpreadAsRequestParameterOptions contains the optional parameters for the SpreadAliasClient.SpreadAsRequestParameter
//     method.
func (client *SpreadAliasClient) SpreadAsRequestParameter(ctx context.Context, id string, xMSTestHeader string, name string, options *SpreadAliasClientSpreadAsRequestParameterOptions) (SpreadAliasClientSpreadAsRequestParameterResponse, error) {
	var err error
	const operationName = "SpreadAliasClient.SpreadAsRequestParameter"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadAsRequestParameterCreateRequest(ctx, id, xMSTestHeader, name, options)
	if err != nil {
		return SpreadAliasClientSpreadAsRequestParameterResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadAliasClientSpreadAsRequestParameterResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadAliasClientSpreadAsRequestParameterResponse{}, err
	}
	return SpreadAliasClientSpreadAsRequestParameterResponse{}, nil
}

// spreadAsRequestParameterCreateRequest creates the SpreadAsRequestParameter request.
func (client *SpreadAliasClient) spreadAsRequestParameterCreateRequest(ctx context.Context, id string, xMSTestHeader string, name string, _ *SpreadAliasClientSpreadAsRequestParameterOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/alias/request-parameter/{id}"
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	req.Raw().Header["x-ms-test-header"] = []string{xMSTestHeader}
	body := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// SpreadParameterWithInnerAlias - spread an alias with contains another alias property as body.
// If the operation fails it returns an *azcore.ResponseError type.
//   - name - name of the Thing
//   - age - age of the Thing
//   - options - SpreadAliasClientSpreadParameterWithInnerAliasOptions contains the optional parameters for the SpreadAliasClient.SpreadParameterWithInnerAlias
//     method.
func (client *SpreadAliasClient) SpreadParameterWithInnerAlias(ctx context.Context, id string, name string, age int32, xMSTestHeader string, options *SpreadAliasClientSpreadParameterWithInnerAliasOptions) (SpreadAliasClientSpreadParameterWithInnerAliasResponse, error) {
	var err error
	const operationName = "SpreadAliasClient.SpreadParameterWithInnerAlias"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadParameterWithInnerAliasCreateRequest(ctx, id, name, age, xMSTestHeader, options)
	if err != nil {
		return SpreadAliasClientSpreadParameterWithInnerAliasResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadAliasClientSpreadParameterWithInnerAliasResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadAliasClientSpreadParameterWithInnerAliasResponse{}, err
	}
	return SpreadAliasClientSpreadParameterWithInnerAliasResponse{}, nil
}

// spreadParameterWithInnerAliasCreateRequest creates the SpreadParameterWithInnerAlias request.
func (client *SpreadAliasClient) spreadParameterWithInnerAliasCreateRequest(ctx context.Context, id string, name string, age int32, xMSTestHeader string, _ *SpreadAliasClientSpreadParameterWithInnerAliasOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/alias/inner-alias-parameter/{id}"
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	req.Raw().Header["x-ms-test-header"] = []string{xMSTestHeader}
	body := struct {
		Name string `json:"name"`
		Age  int32  `json:"age"`
	}{
		Name: name,
		Age:  age,
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// SpreadParameterWithInnerModel -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - SpreadAliasClientSpreadParameterWithInnerModelOptions contains the optional parameters for the SpreadAliasClient.SpreadParameterWithInnerModel
//     method.
func (client *SpreadAliasClient) SpreadParameterWithInnerModel(ctx context.Context, id string, name string, xMSTestHeader string, options *SpreadAliasClientSpreadParameterWithInnerModelOptions) (SpreadAliasClientSpreadParameterWithInnerModelResponse, error) {
	var err error
	const operationName = "SpreadAliasClient.SpreadParameterWithInnerModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadParameterWithInnerModelCreateRequest(ctx, id, name, xMSTestHeader, options)
	if err != nil {
		return SpreadAliasClientSpreadParameterWithInnerModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadAliasClientSpreadParameterWithInnerModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadAliasClientSpreadParameterWithInnerModelResponse{}, err
	}
	return SpreadAliasClientSpreadParameterWithInnerModelResponse{}, nil
}

// spreadParameterWithInnerModelCreateRequest creates the SpreadParameterWithInnerModel request.
func (client *SpreadAliasClient) spreadParameterWithInnerModelCreateRequest(ctx context.Context, id string, name string, xMSTestHeader string, _ *SpreadAliasClientSpreadParameterWithInnerModelOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/alias/inner-model-parameter/{id}"
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	req.Raw().Header["x-ms-test-header"] = []string{xMSTestHeader}
	body := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// SpreadWithMultipleParameters -
// If the operation fails it returns an *azcore.ResponseError type.
//   - requiredString - required string
//   - optionalInt - optional int
//   - requiredIntList - required int
//   - optionalStringList - optional string
//   - options - SpreadAliasClientSpreadWithMultipleParametersOptions contains the optional parameters for the SpreadAliasClient.SpreadWithMultipleParameters
//     method.
func (client *SpreadAliasClient) SpreadWithMultipleParameters(ctx context.Context, id string, xMSTestHeader string, requiredString string, optionalInt int32, requiredIntList []int32, optionalStringList []string, options *SpreadAliasClientSpreadWithMultipleParametersOptions) (SpreadAliasClientSpreadWithMultipleParametersResponse, error) {
	var err error
	const operationName = "SpreadAliasClient.SpreadWithMultipleParameters"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadWithMultipleParametersCreateRequest(ctx, id, xMSTestHeader, requiredString, optionalInt, requiredIntList, optionalStringList, options)
	if err != nil {
		return SpreadAliasClientSpreadWithMultipleParametersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadAliasClientSpreadWithMultipleParametersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadAliasClientSpreadWithMultipleParametersResponse{}, err
	}
	return SpreadAliasClientSpreadWithMultipleParametersResponse{}, nil
}

// spreadWithMultipleParametersCreateRequest creates the SpreadWithMultipleParameters request.
func (client *SpreadAliasClient) spreadWithMultipleParametersCreateRequest(ctx context.Context, id string, xMSTestHeader string, requiredString string, optionalInt int32, requiredIntList []int32, optionalStringList []string, _ *SpreadAliasClientSpreadWithMultipleParametersOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/alias/multiple-parameters/{id}"
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	req.Raw().Header["x-ms-test-header"] = []string{xMSTestHeader}
	body := struct {
		RequiredString     string   `json:"requiredString"`
		OptionalInt        int32    `json:"optionalInt"`
		RequiredIntList    []int32  `json:"requiredIntList"`
		OptionalStringList []string `json:"optionalStringList"`
	}{
		RequiredString:     requiredString,
		OptionalInt:        optionalInt,
		RequiredIntList:    requiredIntList,
		OptionalStringList: optionalStringList,
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
