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

// DictionaryOperations contains the methods for the Dictionary group.
type DictionaryOperations interface {
	// GetEmpty - Get complex types with dictionary property which is empty
	GetEmpty(ctx context.Context, options *DictionaryGetEmptyOptions) (*DictionaryWrapperResponse, error)
	// GetNotProvided - Get complex types with dictionary property while server doesn't provide a response payload
	GetNotProvided(ctx context.Context, options *DictionaryGetNotProvidedOptions) (*DictionaryWrapperResponse, error)
	// GetNull - Get complex types with dictionary property which is null
	GetNull(ctx context.Context, options *DictionaryGetNullOptions) (*DictionaryWrapperResponse, error)
	// GetValid - Get complex types with dictionary property
	GetValid(ctx context.Context, options *DictionaryGetValidOptions) (*DictionaryWrapperResponse, error)
	// PutEmpty - Put complex types with dictionary property which is empty
	PutEmpty(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryPutEmptyOptions) (*http.Response, error)
	// PutValid - Put complex types with dictionary property
	PutValid(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryPutValidOptions) (*http.Response, error)
}

// DictionaryClient implements the DictionaryOperations interface.
// Don't use this type directly, use NewDictionaryClient() instead.
type DictionaryClient struct {
	*Client
}

// NewDictionaryClient creates a new instance of DictionaryClient with the specified values.
func NewDictionaryClient(c *Client) DictionaryOperations {
	return &DictionaryClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *DictionaryClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetEmpty - Get complex types with dictionary property which is empty
func (client *DictionaryClient) GetEmpty(ctx context.Context, options *DictionaryGetEmptyOptions) (*DictionaryWrapperResponse, error) {
	req, err := client.GetEmptyCreateRequest(ctx, options)
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
	result, err := client.GetEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEmptyCreateRequest creates the GetEmpty request.
func (client *DictionaryClient) GetEmptyCreateRequest(ctx context.Context, options *DictionaryGetEmptyOptions) (*azcore.Request, error) {
	urlPath := "/complex/dictionary/typed/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (client *DictionaryClient) GetEmptyHandleResponse(resp *azcore.Response) (*DictionaryWrapperResponse, error) {
	result := DictionaryWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DictionaryWrapper)
}

// GetEmptyHandleError handles the GetEmpty error response.
func (client *DictionaryClient) GetEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotProvided - Get complex types with dictionary property while server doesn't provide a response payload
func (client *DictionaryClient) GetNotProvided(ctx context.Context, options *DictionaryGetNotProvidedOptions) (*DictionaryWrapperResponse, error) {
	req, err := client.GetNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNotProvidedHandleError(resp)
	}
	result, err := client.GetNotProvidedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNotProvidedCreateRequest creates the GetNotProvided request.
func (client *DictionaryClient) GetNotProvidedCreateRequest(ctx context.Context, options *DictionaryGetNotProvidedOptions) (*azcore.Request, error) {
	urlPath := "/complex/dictionary/typed/notprovided"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNotProvidedHandleResponse handles the GetNotProvided response.
func (client *DictionaryClient) GetNotProvidedHandleResponse(resp *azcore.Response) (*DictionaryWrapperResponse, error) {
	result := DictionaryWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DictionaryWrapper)
}

// GetNotProvidedHandleError handles the GetNotProvided error response.
func (client *DictionaryClient) GetNotProvidedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get complex types with dictionary property which is null
func (client *DictionaryClient) GetNull(ctx context.Context, options *DictionaryGetNullOptions) (*DictionaryWrapperResponse, error) {
	req, err := client.GetNullCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNullHandleError(resp)
	}
	result, err := client.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNullCreateRequest creates the GetNull request.
func (client *DictionaryClient) GetNullCreateRequest(ctx context.Context, options *DictionaryGetNullOptions) (*azcore.Request, error) {
	urlPath := "/complex/dictionary/typed/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNullHandleResponse handles the GetNull response.
func (client *DictionaryClient) GetNullHandleResponse(resp *azcore.Response) (*DictionaryWrapperResponse, error) {
	result := DictionaryWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DictionaryWrapper)
}

// GetNullHandleError handles the GetNull error response.
func (client *DictionaryClient) GetNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetValid - Get complex types with dictionary property
func (client *DictionaryClient) GetValid(ctx context.Context, options *DictionaryGetValidOptions) (*DictionaryWrapperResponse, error) {
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
func (client *DictionaryClient) GetValidCreateRequest(ctx context.Context, options *DictionaryGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/dictionary/typed/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetValidHandleResponse handles the GetValid response.
func (client *DictionaryClient) GetValidHandleResponse(resp *azcore.Response) (*DictionaryWrapperResponse, error) {
	result := DictionaryWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DictionaryWrapper)
}

// GetValidHandleError handles the GetValid error response.
func (client *DictionaryClient) GetValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutEmpty - Put complex types with dictionary property which is empty
func (client *DictionaryClient) PutEmpty(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryPutEmptyOptions) (*http.Response, error) {
	req, err := client.PutEmptyCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// PutEmptyCreateRequest creates the PutEmpty request.
func (client *DictionaryClient) PutEmptyCreateRequest(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryPutEmptyOptions) (*azcore.Request, error) {
	urlPath := "/complex/dictionary/typed/empty"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutEmptyHandleError handles the PutEmpty error response.
func (client *DictionaryClient) PutEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValid - Put complex types with dictionary property
func (client *DictionaryClient) PutValid(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryPutValidOptions) (*http.Response, error) {
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
func (client *DictionaryClient) PutValidCreateRequest(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryPutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/dictionary/typed/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutValidHandleError handles the PutValid error response.
func (client *DictionaryClient) PutValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
