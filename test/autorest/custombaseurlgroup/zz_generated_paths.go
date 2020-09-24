// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package custombaseurlgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strings"
)

// PathsOperations contains the methods for the Paths group.
type PathsOperations interface {
	// GetEmpty - Get a 200 to test a valid base uri
	GetEmpty(ctx context.Context, accountName string, options *PathsGetEmptyOptions) (*http.Response, error)
}

// PathsClient implements the PathsOperations interface.
// Don't use this type directly, use NewPathsClient() instead.
type PathsClient struct {
	*Client
}

// NewPathsClient creates a new instance of PathsClient with the specified values.
func NewPathsClient(c *Client) PathsOperations {
	return &PathsClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *PathsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetEmpty - Get a 200 to test a valid base uri
func (client *PathsClient) GetEmpty(ctx context.Context, accountName string, options *PathsGetEmptyOptions) (*http.Response, error) {
	req, err := client.GetEmptyCreateRequest(ctx, accountName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// GetEmptyCreateRequest creates the GetEmpty request.
func (client *PathsClient) GetEmptyCreateRequest(ctx context.Context, accountName string, options *PathsGetEmptyOptions) (*azcore.Request, error) {
	host := "http://{accountName}{host}"
	host = strings.ReplaceAll(host, "{host}", client.host)
	host = strings.ReplaceAll(host, "{accountName}", accountName)
	urlPath := "/customuri"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetEmptyHandleError handles the GetEmpty error response.
func (client *PathsClient) GetEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
