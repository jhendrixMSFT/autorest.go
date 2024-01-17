//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package basicgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// TwoModelsAsPageItemClient contains the methods for the _Specs_.Azure.Core.Basic group.
// Don't use this type directly, use a constructor function instead.
type TwoModelsAsPageItemClient struct {
	internal   *azcore.Client
	apiVersion string
}

// NewListFirstItemPager - Two operations with two different page item types should be successfully generated. Should generate
// model for FirstItem.
func (client *TwoModelsAsPageItemClient) NewListFirstItemPager(options *TwoModelsAsPageItemClientListFirstItemOptions) *runtime.Pager[TwoModelsAsPageItemClientListFirstItemResponse] {
	return runtime.NewPager(runtime.PagingHandler[TwoModelsAsPageItemClientListFirstItemResponse]{
		More: func(page TwoModelsAsPageItemClientListFirstItemResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TwoModelsAsPageItemClientListFirstItemResponse) (TwoModelsAsPageItemClientListFirstItemResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listFirstItemCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TwoModelsAsPageItemClientListFirstItemResponse{}, err
			}
			return client.listFirstItemHandleResponse(resp)
		},
	})
}

// listFirstItemCreateRequest creates the ListFirstItem request.
func (client *TwoModelsAsPageItemClient) listFirstItemCreateRequest(ctx context.Context, options *TwoModelsAsPageItemClientListFirstItemOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/first-item"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listFirstItemHandleResponse handles the ListFirstItem response.
func (client *TwoModelsAsPageItemClient) listFirstItemHandleResponse(resp *http.Response) (TwoModelsAsPageItemClientListFirstItemResponse, error) {
	result := TwoModelsAsPageItemClientListFirstItemResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedFirstItem); err != nil {
		return TwoModelsAsPageItemClientListFirstItemResponse{}, err
	}
	return result, nil
}

// NewListSecondItemPager - Two operations with two different page item types should be successfully generated. Should generate
// model for SecondItem.
func (client *TwoModelsAsPageItemClient) NewListSecondItemPager(options *TwoModelsAsPageItemClientListSecondItemOptions) *runtime.Pager[TwoModelsAsPageItemClientListSecondItemResponse] {
	return runtime.NewPager(runtime.PagingHandler[TwoModelsAsPageItemClientListSecondItemResponse]{
		More: func(page TwoModelsAsPageItemClientListSecondItemResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TwoModelsAsPageItemClientListSecondItemResponse) (TwoModelsAsPageItemClientListSecondItemResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSecondItemCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TwoModelsAsPageItemClientListSecondItemResponse{}, err
			}
			return client.listSecondItemHandleResponse(resp)
		},
	})
}

// listSecondItemCreateRequest creates the ListSecondItem request.
func (client *TwoModelsAsPageItemClient) listSecondItemCreateRequest(ctx context.Context, options *TwoModelsAsPageItemClientListSecondItemOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/second-item"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecondItemHandleResponse handles the ListSecondItem response.
func (client *TwoModelsAsPageItemClient) listSecondItemHandleResponse(resp *http.Response) (TwoModelsAsPageItemClientListSecondItemResponse, error) {
	result := TwoModelsAsPageItemClientListSecondItemResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedSecondItem); err != nil {
		return TwoModelsAsPageItemClientListSecondItemResponse{}, err
	}
	return result, nil
}
