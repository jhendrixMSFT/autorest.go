// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ReadonlypropertyOperations contains the methods for the Readonlyproperty group.
type ReadonlypropertyOperations interface {
	// GetValid - Get complex types that have readonly properties
	GetValid(ctx context.Context, options *ReadonlypropertyGetValidOptions) (*ReadonlyObjResponse, error)
	// PutValid - Put complex types that have readonly properties
	PutValid(ctx context.Context, complexBody ReadonlyObj, options *ReadonlypropertyPutValidOptions) (*http.Response, error)
}

// ReadonlypropertyClient implements the ReadonlypropertyOperations interface.
// Don't use this type directly, use NewReadonlypropertyClient() instead.
type ReadonlypropertyClient struct {
	*Client
}

// NewReadonlypropertyClient creates a new instance of ReadonlypropertyClient with the specified values.
func NewReadonlypropertyClient(c *Client) ReadonlypropertyOperations {
	return &ReadonlypropertyClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ReadonlypropertyClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetValid - Get complex types that have readonly properties
func (client *ReadonlypropertyClient) GetValid(ctx context.Context, options *ReadonlypropertyGetValidOptions) (*ReadonlyObjResponse, error) {
	req, err := client.GetValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetValidHandleError(resp)
	}
	result, err := client.GetValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetValidCreateRequest creates the GetValid request.
func (client *ReadonlypropertyClient) GetValidCreateRequest(ctx context.Context, options *ReadonlypropertyGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/readonlyproperty/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetValidHandleResponse handles the GetValid response.
func (client *ReadonlypropertyClient) GetValidHandleResponse(resp *azcore.Response) (*ReadonlyObjResponse, error) {
	result := ReadonlyObjResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ReadonlyObj)
}

// GetValidHandleError handles the GetValid error response.
func (client *ReadonlypropertyClient) GetValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValid - Put complex types that have readonly properties
func (client *ReadonlypropertyClient) PutValid(ctx context.Context, complexBody ReadonlyObj, options *ReadonlypropertyPutValidOptions) (*http.Response, error) {
	req, err := client.PutValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutValidHandleError(resp)
	}
	return resp.Response, nil
}

// PutValidCreateRequest creates the PutValid request.
func (client *ReadonlypropertyClient) PutValidCreateRequest(ctx context.Context, complexBody ReadonlyObj, options *ReadonlypropertyPutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/readonlyproperty/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutValidHandleError handles the PutValid error response.
func (client *ReadonlypropertyClient) PutValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
