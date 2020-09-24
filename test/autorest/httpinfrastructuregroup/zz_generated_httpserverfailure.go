// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HTTPServerFailureOperations contains the methods for the HTTPServerFailure group.
type HTTPServerFailureOperations interface {
	// Delete505 - Return 505 status code - should be represented in the client as an error
	Delete505(ctx context.Context, options *HTTPServerFailureDelete505Options) (*http.Response, error)
	// Get501 - Return 501 status code - should be represented in the client as an error
	Get501(ctx context.Context, options *HTTPServerFailureGet501Options) (*http.Response, error)
	// Head501 - Return 501 status code - should be represented in the client as an error
	Head501(ctx context.Context, options *HTTPServerFailureHead501Options) (*http.Response, error)
	// Post505 - Return 505 status code - should be represented in the client as an error
	Post505(ctx context.Context, options *HTTPServerFailurePost505Options) (*http.Response, error)
}

// HTTPServerFailureClient implements the HTTPServerFailureOperations interface.
// Don't use this type directly, use NewHTTPServerFailureClient() instead.
type HTTPServerFailureClient struct {
	*Client
}

// NewHTTPServerFailureClient creates a new instance of HTTPServerFailureClient with the specified values.
func NewHTTPServerFailureClient(c *Client) HTTPServerFailureOperations {
	return &HTTPServerFailureClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *HTTPServerFailureClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Delete505 - Return 505 status code - should be represented in the client as an error
func (client *HTTPServerFailureClient) Delete505(ctx context.Context, options *HTTPServerFailureDelete505Options) (*http.Response, error) {
	req, err := client.Delete505CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.Delete505HandleError(resp)
	}
	return resp.Response, nil
}

// Delete505CreateRequest creates the Delete505 request.
func (client *HTTPServerFailureClient) Delete505CreateRequest(ctx context.Context, options *HTTPServerFailureDelete505Options) (*azcore.Request, error) {
	urlPath := "/http/failure/server/505"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Delete505HandleError handles the Delete505 error response.
func (client *HTTPServerFailureClient) Delete505HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get501 - Return 501 status code - should be represented in the client as an error
func (client *HTTPServerFailureClient) Get501(ctx context.Context, options *HTTPServerFailureGet501Options) (*http.Response, error) {
	req, err := client.Get501CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.Get501HandleError(resp)
	}
	return resp.Response, nil
}

// Get501CreateRequest creates the Get501 request.
func (client *HTTPServerFailureClient) Get501CreateRequest(ctx context.Context, options *HTTPServerFailureGet501Options) (*azcore.Request, error) {
	urlPath := "/http/failure/server/501"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Get501HandleError handles the Get501 error response.
func (client *HTTPServerFailureClient) Get501HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head501 - Return 501 status code - should be represented in the client as an error
func (client *HTTPServerFailureClient) Head501(ctx context.Context, options *HTTPServerFailureHead501Options) (*http.Response, error) {
	req, err := client.Head501CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.Head501HandleError(resp)
	}
	return resp.Response, nil
}

// Head501CreateRequest creates the Head501 request.
func (client *HTTPServerFailureClient) Head501CreateRequest(ctx context.Context, options *HTTPServerFailureHead501Options) (*azcore.Request, error) {
	urlPath := "/http/failure/server/501"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Head501HandleError handles the Head501 error response.
func (client *HTTPServerFailureClient) Head501HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post505 - Return 505 status code - should be represented in the client as an error
func (client *HTTPServerFailureClient) Post505(ctx context.Context, options *HTTPServerFailurePost505Options) (*http.Response, error) {
	req, err := client.Post505CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.Post505HandleError(resp)
	}
	return resp.Response, nil
}

// Post505CreateRequest creates the Post505 request.
func (client *HTTPServerFailureClient) Post505CreateRequest(ctx context.Context, options *HTTPServerFailurePost505Options) (*azcore.Request, error) {
	urlPath := "/http/failure/server/505"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Post505HandleError handles the Post505 error response.
func (client *HTTPServerFailureClient) Post505HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
