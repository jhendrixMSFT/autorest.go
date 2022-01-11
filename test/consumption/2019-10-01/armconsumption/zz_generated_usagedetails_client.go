//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// UsageDetailsClient contains the methods for the UsageDetails group.
// Don't use this type directly, use NewUsageDetailsClient() instead.
type UsageDetailsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewUsageDetailsClient creates a new instance of UsageDetailsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUsageDetailsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *UsageDetailsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &UsageDetailsClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Lists the usage details for the defined scope. Usage details are available via this API only for May 1, 2014 or
// later. For more information on using this API, including how to specify a date range,
// please see: https://docs.microsoft.com/en-us/azure/cost-management-billing/costs/manage-automation
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope associated with usage details operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for
// Billing Account scope, '/providers/Microsoft.Billing/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount
// scope and '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope. For subscription,
// billing account, department, enrollment account and management group, you
// can also add billing period to the scope using '/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'. For e.g.
// to specify billing period at department scope use
// '/providers/Microsoft.Billing/departments/{departmentId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'.
// Also, Modern Commerce Account scopes are
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for billingAccount scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for
// billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// options - UsageDetailsClientListOptions contains the optional parameters for the UsageDetailsClient.List method.
func (client *UsageDetailsClient) List(scope string, options *UsageDetailsClientListOptions) *UsageDetailsClientListPager {
	return &UsageDetailsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp UsageDetailsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.UsageDetailsListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *UsageDetailsClient) listCreateRequest(ctx context.Context, scope string, options *UsageDetailsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/usageDetails"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-10-01")
	if options != nil && options.Metric != nil {
		reqQP.Set("metric", string(*options.Metric))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsageDetailsClient) listHandleResponse(resp *http.Response) (UsageDetailsClientListResponse, error) {
	result := UsageDetailsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageDetailsListResult); err != nil {
		return UsageDetailsClientListResponse{}, err
	}
	return result, nil
}
