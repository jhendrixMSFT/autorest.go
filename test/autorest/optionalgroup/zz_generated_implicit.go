// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package optionalgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ImplicitOperations contains the methods for the Implicit group.
type ImplicitOperations interface {
	// GetOptionalGlobalQuery - Test implicitly optional query parameter
	GetOptionalGlobalQuery(ctx context.Context, options *ImplicitGetOptionalGlobalQueryOptions) (*http.Response, error)
	// GetRequiredGlobalPath - Test implicitly required path parameter
	GetRequiredGlobalPath(ctx context.Context, options *ImplicitGetRequiredGlobalPathOptions) (*http.Response, error)
	// GetRequiredGlobalQuery - Test implicitly required query parameter
	GetRequiredGlobalQuery(ctx context.Context, options *ImplicitGetRequiredGlobalQueryOptions) (*http.Response, error)
	// GetRequiredPath - Test implicitly required path parameter
	GetRequiredPath(ctx context.Context, pathParameter string, options *ImplicitGetRequiredPathOptions) (*http.Response, error)
	// PutOptionalBody - Test implicitly optional body parameter
	PutOptionalBody(ctx context.Context, options *ImplicitPutOptionalBodyOptions) (*http.Response, error)
	// PutOptionalHeader - Test implicitly optional header parameter
	PutOptionalHeader(ctx context.Context, options *ImplicitPutOptionalHeaderOptions) (*http.Response, error)
	// PutOptionalQuery - Test implicitly optional query parameter
	PutOptionalQuery(ctx context.Context, options *ImplicitPutOptionalQueryOptions) (*http.Response, error)
}

// ImplicitClient implements the ImplicitOperations interface.
// Don't use this type directly, use NewImplicitClient() instead.
type ImplicitClient struct {
	*Client
	requiredGlobalPath  string
	requiredGlobalQuery string
	optionalGlobalQuery *int32
}

// NewImplicitClient creates a new instance of ImplicitClient with the specified values.
func NewImplicitClient(c *Client, requiredGlobalPath string, requiredGlobalQuery string, optionalGlobalQuery *int32) ImplicitOperations {
	return &ImplicitClient{Client: c, requiredGlobalPath: requiredGlobalPath, requiredGlobalQuery: requiredGlobalQuery, optionalGlobalQuery: optionalGlobalQuery}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ImplicitClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetOptionalGlobalQuery - Test implicitly optional query parameter
func (client *ImplicitClient) GetOptionalGlobalQuery(ctx context.Context, options *ImplicitGetOptionalGlobalQueryOptions) (*http.Response, error) {
	req, err := client.GetOptionalGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetOptionalGlobalQueryHandleError(resp)
	}
	return resp.Response, nil
}

// GetOptionalGlobalQueryCreateRequest creates the GetOptionalGlobalQuery request.
func (client *ImplicitClient) GetOptionalGlobalQueryCreateRequest(ctx context.Context, options *ImplicitGetOptionalGlobalQueryOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/global/optional/query"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if client.optionalGlobalQuery != nil {
		query.Set("optional-global-query", strconv.FormatInt(int64(*client.optionalGlobalQuery), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetOptionalGlobalQueryHandleError handles the GetOptionalGlobalQuery error response.
func (client *ImplicitClient) GetOptionalGlobalQueryHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetRequiredGlobalPath - Test implicitly required path parameter
func (client *ImplicitClient) GetRequiredGlobalPath(ctx context.Context, options *ImplicitGetRequiredGlobalPathOptions) (*http.Response, error) {
	req, err := client.GetRequiredGlobalPathCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetRequiredGlobalPathHandleError(resp)
	}
	return resp.Response, nil
}

// GetRequiredGlobalPathCreateRequest creates the GetRequiredGlobalPath request.
func (client *ImplicitClient) GetRequiredGlobalPathCreateRequest(ctx context.Context, options *ImplicitGetRequiredGlobalPathOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/global/required/path/{required-global-path}"
	urlPath = strings.ReplaceAll(urlPath, "{required-global-path}", url.PathEscape(client.requiredGlobalPath))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetRequiredGlobalPathHandleError handles the GetRequiredGlobalPath error response.
func (client *ImplicitClient) GetRequiredGlobalPathHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetRequiredGlobalQuery - Test implicitly required query parameter
func (client *ImplicitClient) GetRequiredGlobalQuery(ctx context.Context, options *ImplicitGetRequiredGlobalQueryOptions) (*http.Response, error) {
	req, err := client.GetRequiredGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetRequiredGlobalQueryHandleError(resp)
	}
	return resp.Response, nil
}

// GetRequiredGlobalQueryCreateRequest creates the GetRequiredGlobalQuery request.
func (client *ImplicitClient) GetRequiredGlobalQueryCreateRequest(ctx context.Context, options *ImplicitGetRequiredGlobalQueryOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/global/required/query"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("required-global-query", client.requiredGlobalQuery)
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetRequiredGlobalQueryHandleError handles the GetRequiredGlobalQuery error response.
func (client *ImplicitClient) GetRequiredGlobalQueryHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetRequiredPath - Test implicitly required path parameter
func (client *ImplicitClient) GetRequiredPath(ctx context.Context, pathParameter string, options *ImplicitGetRequiredPathOptions) (*http.Response, error) {
	req, err := client.GetRequiredPathCreateRequest(ctx, pathParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetRequiredPathHandleError(resp)
	}
	return resp.Response, nil
}

// GetRequiredPathCreateRequest creates the GetRequiredPath request.
func (client *ImplicitClient) GetRequiredPathCreateRequest(ctx context.Context, pathParameter string, options *ImplicitGetRequiredPathOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/implicit/required/path/{pathParameter}"
	urlPath = strings.ReplaceAll(urlPath, "{pathParameter}", url.PathEscape(pathParameter))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetRequiredPathHandleError handles the GetRequiredPath error response.
func (client *ImplicitClient) GetRequiredPathHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutOptionalBody - Test implicitly optional body parameter
func (client *ImplicitClient) PutOptionalBody(ctx context.Context, options *ImplicitPutOptionalBodyOptions) (*http.Response, error) {
	req, err := client.PutOptionalBodyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutOptionalBodyHandleError(resp)
	}
	return resp.Response, nil
}

// PutOptionalBodyCreateRequest creates the PutOptionalBody request.
func (client *ImplicitClient) PutOptionalBodyCreateRequest(ctx context.Context, options *ImplicitPutOptionalBodyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/implicit/optional/body"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.BodyParameter)
	}
	return req, nil
}

// PutOptionalBodyHandleError handles the PutOptionalBody error response.
func (client *ImplicitClient) PutOptionalBodyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutOptionalHeader - Test implicitly optional header parameter
func (client *ImplicitClient) PutOptionalHeader(ctx context.Context, options *ImplicitPutOptionalHeaderOptions) (*http.Response, error) {
	req, err := client.PutOptionalHeaderCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutOptionalHeaderHandleError(resp)
	}
	return resp.Response, nil
}

// PutOptionalHeaderCreateRequest creates the PutOptionalHeader request.
func (client *ImplicitClient) PutOptionalHeaderCreateRequest(ctx context.Context, options *ImplicitPutOptionalHeaderOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/implicit/optional/header"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.QueryParameter != nil {
		req.Header.Set("queryParameter", *options.QueryParameter)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PutOptionalHeaderHandleError handles the PutOptionalHeader error response.
func (client *ImplicitClient) PutOptionalHeaderHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutOptionalQuery - Test implicitly optional query parameter
func (client *ImplicitClient) PutOptionalQuery(ctx context.Context, options *ImplicitPutOptionalQueryOptions) (*http.Response, error) {
	req, err := client.PutOptionalQueryCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutOptionalQueryHandleError(resp)
	}
	return resp.Response, nil
}

// PutOptionalQueryCreateRequest creates the PutOptionalQuery request.
func (client *ImplicitClient) PutOptionalQueryCreateRequest(ctx context.Context, options *ImplicitPutOptionalQueryOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/implicit/optional/query"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if options != nil && options.QueryParameter != nil {
		query.Set("queryParameter", *options.QueryParameter)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PutOptionalQueryHandleError handles the PutOptionalQuery error response.
func (client *ImplicitClient) PutOptionalQueryHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
