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
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReservationsDetailsClient contains the methods for the ReservationsDetails group.
// Don't use this type directly, use NewReservationsDetailsClient() instead.
type ReservationsDetailsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewReservationsDetailsClient creates a new instance of ReservationsDetailsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReservationsDetailsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReservationsDetailsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ReservationsDetailsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Lists the reservations details for the defined scope and provided date range.
//
// Generated from API version 2019-10-01
//   - scope - The scope associated with reservations details operations. This includes '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}'
//     for BillingAccount scope (legacy), and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile
//     scope (modern).
//   - options - ReservationsDetailsClientListOptions contains the optional parameters for the ReservationsDetailsClient.List
//     method.
func (client *ReservationsDetailsClient) NewListPager(scope string, options *ReservationsDetailsClientListOptions) *runtime.Pager[ReservationsDetailsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationsDetailsClientListResponse]{
		More: func(page ReservationsDetailsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationsDetailsClientListResponse) (ReservationsDetailsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReservationsDetailsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReservationsDetailsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReservationsDetailsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReservationsDetailsClient) listCreateRequest(ctx context.Context, scope string, options *ReservationsDetailsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/reservationDetails"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StartDate != nil {
		reqQP.Set("startDate", *options.StartDate)
	}
	if options != nil && options.EndDate != nil {
		reqQP.Set("endDate", *options.EndDate)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ReservationID != nil {
		reqQP.Set("reservationId", *options.ReservationID)
	}
	if options != nil && options.ReservationOrderID != nil {
		reqQP.Set("reservationOrderId", *options.ReservationOrderID)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReservationsDetailsClient) listHandleResponse(resp *http.Response) (ReservationsDetailsClientListResponse, error) {
	result := ReservationsDetailsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationDetailsListResult); err != nil {
		return ReservationsDetailsClientListResponse{}, err
	}
	return result, nil
}

// NewListByReservationOrderPager - Lists the reservations details for provided date range.
//
// Generated from API version 2019-10-01
//   - reservationOrderID - Order Id of the reservation
//   - filter - Filter reservation details by date range. The properties/UsageDate for start date and end date. The filter supports
//     'le' and 'ge'
//   - options - ReservationsDetailsClientListByReservationOrderOptions contains the optional parameters for the ReservationsDetailsClient.ListByReservationOrder
//     method.
func (client *ReservationsDetailsClient) NewListByReservationOrderPager(reservationOrderID string, filter string, options *ReservationsDetailsClientListByReservationOrderOptions) *runtime.Pager[ReservationsDetailsClientListByReservationOrderResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationsDetailsClientListByReservationOrderResponse]{
		More: func(page ReservationsDetailsClientListByReservationOrderResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationsDetailsClientListByReservationOrderResponse) (ReservationsDetailsClientListByReservationOrderResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReservationOrderCreateRequest(ctx, reservationOrderID, filter, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReservationsDetailsClientListByReservationOrderResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReservationsDetailsClientListByReservationOrderResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReservationsDetailsClientListByReservationOrderResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReservationOrderHandleResponse(resp)
		},
	})
}

// listByReservationOrderCreateRequest creates the ListByReservationOrder request.
func (client *ReservationsDetailsClient) listByReservationOrderCreateRequest(ctx context.Context, reservationOrderID string, filter string, options *ReservationsDetailsClientListByReservationOrderOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationDetails"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReservationOrderHandleResponse handles the ListByReservationOrder response.
func (client *ReservationsDetailsClient) listByReservationOrderHandleResponse(resp *http.Response) (ReservationsDetailsClientListByReservationOrderResponse, error) {
	result := ReservationsDetailsClientListByReservationOrderResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationDetailsListResult); err != nil {
		return ReservationsDetailsClientListByReservationOrderResponse{}, err
	}
	return result, nil
}

// NewListByReservationOrderAndReservationPager - Lists the reservations details for provided date range.
//
// Generated from API version 2019-10-01
//   - reservationOrderID - Order Id of the reservation
//   - reservationID - Id of the reservation
//   - filter - Filter reservation details by date range. The properties/UsageDate for start date and end date. The filter supports
//     'le' and 'ge'
//   - options - ReservationsDetailsClientListByReservationOrderAndReservationOptions contains the optional parameters for the
//     ReservationsDetailsClient.ListByReservationOrderAndReservation method.
func (client *ReservationsDetailsClient) NewListByReservationOrderAndReservationPager(reservationOrderID string, reservationID string, filter string, options *ReservationsDetailsClientListByReservationOrderAndReservationOptions) *runtime.Pager[ReservationsDetailsClientListByReservationOrderAndReservationResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationsDetailsClientListByReservationOrderAndReservationResponse]{
		More: func(page ReservationsDetailsClientListByReservationOrderAndReservationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationsDetailsClientListByReservationOrderAndReservationResponse) (ReservationsDetailsClientListByReservationOrderAndReservationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReservationOrderAndReservationCreateRequest(ctx, reservationOrderID, reservationID, filter, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReservationsDetailsClientListByReservationOrderAndReservationResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReservationsDetailsClientListByReservationOrderAndReservationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReservationsDetailsClientListByReservationOrderAndReservationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReservationOrderAndReservationHandleResponse(resp)
		},
	})
}

// listByReservationOrderAndReservationCreateRequest creates the ListByReservationOrderAndReservation request.
func (client *ReservationsDetailsClient) listByReservationOrderAndReservationCreateRequest(ctx context.Context, reservationOrderID string, reservationID string, filter string, options *ReservationsDetailsClientListByReservationOrderAndReservationOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationDetails"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReservationOrderAndReservationHandleResponse handles the ListByReservationOrderAndReservation response.
func (client *ReservationsDetailsClient) listByReservationOrderAndReservationHandleResponse(resp *http.Response) (ReservationsDetailsClientListByReservationOrderAndReservationResponse, error) {
	result := ReservationsDetailsClientListByReservationOrderAndReservationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationDetailsListResult); err != nil {
		return ReservationsDetailsClientListByReservationOrderAndReservationResponse{}, err
	}
	return result, nil
}
