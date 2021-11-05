//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ForecastsClient contains the methods for the Forecasts group.
// Don't use this type directly, use NewForecastsClient() instead.
type ForecastsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewForecastsClient creates a new instance of ForecastsClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewForecastsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ForecastsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &ForecastsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// List - Lists the forecast charges for scope defined. Please note that this API is no longer actively under development.
// We recommend using our new Forecast API moving forward:
// https://docs.microsoft.com/en-us/rest/api/cost-management/forecast/usage.
// If the operation fails it returns the *ErrorResponse error type.
// options - ForecastsListOptions contains the optional parameters for the Forecasts.List method.
func (client *ForecastsClient) List(ctx context.Context, options *ForecastsListOptions) (ForecastsListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return ForecastsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ForecastsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ForecastsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ForecastsClient) listCreateRequest(ctx context.Context, options *ForecastsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/forecasts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ForecastsClient) listHandleResponse(resp *http.Response) (ForecastsListResponse, error) {
	result := ForecastsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ForecastsListResult); err != nil {
		return ForecastsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ForecastsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
