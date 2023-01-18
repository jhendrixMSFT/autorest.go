//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconsumption

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// LotsClient contains the methods for the Lots group.
// Don't use this type directly, use NewLotsClient() instead.
type LotsClient struct {
	internal *arm.Client
}

// NewLotsClient creates a new instance of LotsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLotsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*LotsClient, error) {
	cl, err := arm.NewClient("armconsumption.LotsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LotsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Lists the lots by billingAccountId and billingProfileId.
//
// Generated from API version 2019-10-01
//   - scope - The scope associated with Lots operations. This includes '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfile/{billingProfileId}'
//     for Billing Profile scope, and
//     'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
//   - options - LotsClientListOptions contains the optional parameters for the LotsClient.NewListPager method.
func (client *LotsClient) NewListPager(scope string, options *LotsClientListOptions) *runtime.Pager[LotsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LotsClientListResponse]{
		More: func(page LotsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LotsClientListResponse) (LotsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LotsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LotsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LotsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LotsClient) listCreateRequest(ctx context.Context, scope string, options *LotsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/lots"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LotsClient) listHandleResponse(resp *http.Response) (LotsClientListResponse, error) {
	result := LotsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Lots); err != nil {
		return LotsClientListResponse{}, err
	}
	return result, nil
}
