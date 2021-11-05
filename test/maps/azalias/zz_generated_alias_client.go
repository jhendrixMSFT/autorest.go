//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

type aliasClient struct {
	endpoint   string
	apiVersion *string
	pl         runtime.Pipeline
}

// newAliasClient creates a new instance of aliasClient with the specified values.
// geography - This parameter specifies where the Azure Maps Creator resource is located. Valid values are us and eu.
// apiVersion - Api Version. Specifying any value will set the value to 2.0.
// pl - the pipeline used for sending requests and handling responses.
func newAliasClient(geography *Geography, apiVersion *string, pl runtime.Pipeline) *aliasClient {
	hostURL := "https://{geography}.atlas.microsoft.com"
	if geography == nil {
		defaultValue := GeographyUs
		geography = &defaultValue
	}
	hostURL = strings.ReplaceAll(hostURL, "{geography}", string(*geography))
	client := &aliasClient{
		endpoint:   hostURL,
		apiVersion: apiVersion,
		pl:         pl,
	}
	return client
}

// Create - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This
// [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps] article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to create an alias. You can also assign the alias during the create request. An alias can reference
// an ID generated by a creator service, but cannot reference another alias
// ID.
// SUBMIT CREATE REQUEST To create your alias, you will use a POST request. If you would like to assign the alias during the
// creation, you will pass the resourceId query parameter.
// CREATE ALIAS RESPONSE The Create API returns a HTTP 201 Created response with the alias resource in the body.
// A sample response from creating an alias:
// { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId":
// "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14", "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }
// If the operation fails it returns a generic error.
// options - AliasCreateOptions contains the optional parameters for the Alias.Create method.
func (client *aliasClient) Create(ctx context.Context, options *AliasCreateOptions) (AliasCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, options)
	if err != nil {
		return AliasCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AliasCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return AliasCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *aliasClient) createCreateRequest(ctx context.Context, options *AliasCreateOptions) (*policy.Request, error) {
	urlPath := "/aliases"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if client.apiVersion != nil {
		reqQP.Set("api-version", "2.0")
	}
	if options != nil && options.CreatorDataItemID != nil {
		reqQP.Set("creatorDataItemId", *options.CreatorDataItemID)
	}
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", fmt.Sprintf("%d", qv))
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *aliasClient) createHandleResponse(resp *http.Response) (AliasCreateResponse, error) {
	result := AliasCreateResponse{RawResponse: resp}
	if val := resp.Header.Get("Access-Control-Expose-Headers"); val != "" {
		result.AccessControlExposeHeaders = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AliasesCreateResponse); err != nil {
		return AliasCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *aliasClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This
// [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps] article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to fetch a list of all previously successfully created aliases.
// SUBMIT LIST REQUEST To list all your aliases, you will issue a GET request with no additional parameters.
// LIST DATA RESPONSE The List API returns the complete list of all aliases in json format. The response contains the following
// details for each alias resource:
// > createdTimestamp - The timestamp that the alias was created. Format yyyy-MM-ddTHH:mm:ss.sssZ aliasId - The id for the
// alias. creatorDataItemId - The id for the creator data item that this alias
// references (could be null if the alias has not been assigned). lastUpdatedTimestamp - The last time the alias was assigned
// to a resource. Format yyyy-MM-ddTHH:mm:ss.sssZ
// A sample response returning 2 alias resources:
// { "aliases": [ { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId":
// "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14", "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }, { "createdTimestamp": "2020-02-18T19:53:33.123Z", "aliasId": "1856dbfc-7a66-ee5a-bf8d-51dbfe1906f6",
// "creatorDataItemId": null, "lastUpdatedTimestamp":
// "2020-02-18T19:53:33.123Z" } ] }
// If the operation fails it returns a generic error.
// options - AliasListOptions contains the optional parameters for the Alias.List method.
func (client *aliasClient) List(options *AliasListOptions) *AliasListPager {
	return &AliasListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AliasListResponseEnvelope) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AliasListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *aliasClient) listCreateRequest(ctx context.Context, options *AliasListOptions) (*policy.Request, error) {
	urlPath := "/aliases"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if client.apiVersion != nil {
		reqQP.Set("api-version", "2.0")
	}
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", string(qv))
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *aliasClient) listHandleResponse(resp *http.Response) (AliasListResponseEnvelope, error) {
	result := AliasListResponseEnvelope{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AliasListResponse); err != nil {
		return AliasListResponseEnvelope{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *aliasClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
