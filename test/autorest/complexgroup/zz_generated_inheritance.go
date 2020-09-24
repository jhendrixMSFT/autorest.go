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

// InheritanceOperations contains the methods for the Inheritance group.
type InheritanceOperations interface {
	// GetValid - Get complex types that extend others
	GetValid(ctx context.Context, options *InheritanceGetValidOptions) (*SiameseResponse, error)
	// PutValid - Put complex types that extend others
	PutValid(ctx context.Context, complexBody Siamese, options *InheritancePutValidOptions) (*http.Response, error)
}

// InheritanceClient implements the InheritanceOperations interface.
// Don't use this type directly, use NewInheritanceClient() instead.
type InheritanceClient struct {
	*Client
}

// NewInheritanceClient creates a new instance of InheritanceClient with the specified values.
func NewInheritanceClient(c *Client) InheritanceOperations {
	return &InheritanceClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *InheritanceClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetValid - Get complex types that extend others
func (client *InheritanceClient) GetValid(ctx context.Context, options *InheritanceGetValidOptions) (*SiameseResponse, error) {
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
func (client *InheritanceClient) GetValidCreateRequest(ctx context.Context, options *InheritanceGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/inheritance/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetValidHandleResponse handles the GetValid response.
func (client *InheritanceClient) GetValidHandleResponse(resp *azcore.Response) (*SiameseResponse, error) {
	result := SiameseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Siamese)
}

// GetValidHandleError handles the GetValid error response.
func (client *InheritanceClient) GetValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValid - Put complex types that extend others
func (client *InheritanceClient) PutValid(ctx context.Context, complexBody Siamese, options *InheritancePutValidOptions) (*http.Response, error) {
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
func (client *InheritanceClient) PutValidCreateRequest(ctx context.Context, complexBody Siamese, options *InheritancePutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/inheritance/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutValidHandleError handles the PutValid error response.
func (client *InheritanceClient) PutValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
