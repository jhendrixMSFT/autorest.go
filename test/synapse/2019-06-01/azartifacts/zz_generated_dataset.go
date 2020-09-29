// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type datasetClient struct {
	*client
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *datasetClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdateDataset - Creates or updates a dataset.
func (client *datasetClient) CreateOrUpdateDataset(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetCreateOrUpdateDatasetOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateDatasetCreateRequest(ctx, datasetName, dataset, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.CreateOrUpdateDatasetHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateDatasetCreateRequest creates the CreateOrUpdateDataset request.
func (client *datasetClient) CreateOrUpdateDatasetCreateRequest(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetCreateOrUpdateDatasetOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetName}"
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(dataset)
}

// CreateOrUpdateDatasetHandleResponse handles the CreateOrUpdateDataset response.
func (client *datasetClient) CreateOrUpdateDatasetHandleResponse(resp *azcore.Response) (*DatasetResourceResponse, error) {
	result := DatasetResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DatasetResource)
}

// CreateOrUpdateDatasetHandleError handles the CreateOrUpdateDataset error response.
func (client *datasetClient) CreateOrUpdateDatasetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteDataset - Deletes a dataset.
func (client *datasetClient) DeleteDataset(ctx context.Context, datasetName string, options *DatasetDeleteDatasetOptions) (*azcore.Response, error) {
	req, err := client.DeleteDatasetCreateRequest(ctx, datasetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteDatasetHandleError(resp)
	}
	return resp, nil
}

// DeleteDatasetCreateRequest creates the DeleteDataset request.
func (client *datasetClient) DeleteDatasetCreateRequest(ctx context.Context, datasetName string, options *DatasetDeleteDatasetOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetName}"
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteDatasetHandleError handles the DeleteDataset error response.
func (client *datasetClient) DeleteDatasetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDataset - Gets a dataset.
func (client *datasetClient) GetDataset(ctx context.Context, datasetName string, options *DatasetGetDatasetOptions) (*DatasetResourceResponse, error) {
	req, err := client.GetDatasetCreateRequest(ctx, datasetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return nil, client.GetDatasetHandleError(resp)
	}
	result, err := client.GetDatasetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDatasetCreateRequest creates the GetDataset request.
func (client *datasetClient) GetDatasetCreateRequest(ctx context.Context, datasetName string, options *DatasetGetDatasetOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetName}"
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDatasetHandleResponse handles the GetDataset response.
func (client *datasetClient) GetDatasetHandleResponse(resp *azcore.Response) (*DatasetResourceResponse, error) {
	result := DatasetResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DatasetResource)
}

// GetDatasetHandleError handles the GetDataset error response.
func (client *datasetClient) GetDatasetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDatasetsByWorkspace - Lists datasets.
func (client *datasetClient) GetDatasetsByWorkspace(options *DatasetGetDatasetsByWorkspaceOptions) DatasetListResponsePager {
	return &datasetListResponsePager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetDatasetsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.GetDatasetsByWorkspaceHandleResponse,
		errorer:   client.GetDatasetsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *DatasetListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DatasetListResponse.NextLink)
		},
	}
}

// GetDatasetsByWorkspaceCreateRequest creates the GetDatasetsByWorkspace request.
func (client *datasetClient) GetDatasetsByWorkspaceCreateRequest(ctx context.Context, options *DatasetGetDatasetsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/datasets"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDatasetsByWorkspaceHandleResponse handles the GetDatasetsByWorkspace response.
func (client *datasetClient) GetDatasetsByWorkspaceHandleResponse(resp *azcore.Response) (*DatasetListResponseResponse, error) {
	result := DatasetListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DatasetListResponse)
}

// GetDatasetsByWorkspaceHandleError handles the GetDatasetsByWorkspace error response.
func (client *datasetClient) GetDatasetsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}