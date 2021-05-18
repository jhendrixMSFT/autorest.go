// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SecurityPartnerProvidersClient contains the methods for the SecurityPartnerProviders group.
// Don't use this type directly, use NewSecurityPartnerProvidersClient() instead.
type SecurityPartnerProvidersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewSecurityPartnerProvidersClient creates a new instance of SecurityPartnerProvidersClient with the specified values.
func NewSecurityPartnerProvidersClient(con *armcore.Connection, subscriptionID string) *SecurityPartnerProvidersClient {
	return &SecurityPartnerProvidersClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates the specified Security Partner Provider.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityPartnerProvidersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider, options *SecurityPartnerProvidersBeginCreateOrUpdateOptions) (SecurityPartnerProviderPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, securityPartnerProviderName, parameters, options)
	if err != nil {
		return SecurityPartnerProviderPollerResponse{}, err
	}
	result := SecurityPartnerProviderPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("SecurityPartnerProvidersClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return SecurityPartnerProviderPollerResponse{}, err
	}
	poller := &securityPartnerProviderPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SecurityPartnerProviderResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new SecurityPartnerProviderPoller from the specified resume token.
// token - The value must come from a previous call to SecurityPartnerProviderPoller.ResumeToken().
func (client *SecurityPartnerProvidersClient) ResumeCreateOrUpdate(ctx context.Context, token string) (SecurityPartnerProviderPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("SecurityPartnerProvidersClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return SecurityPartnerProviderPollerResponse{}, err
	}
	poller := &securityPartnerProviderPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SecurityPartnerProviderPollerResponse{}, err
	}
	result := SecurityPartnerProviderPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SecurityPartnerProviderResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified Security Partner Provider.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityPartnerProvidersClient) createOrUpdate(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider, options *SecurityPartnerProvidersBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, securityPartnerProviderName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityPartnerProvidersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider, options *SecurityPartnerProvidersBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityPartnerProviderName == "" {
		return nil, errors.New("parameter securityPartnerProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SecurityPartnerProvidersClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes the specified Security Partner Provider.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityPartnerProvidersClient) BeginDelete(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, securityPartnerProviderName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("SecurityPartnerProvidersClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *SecurityPartnerProvidersClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("SecurityPartnerProvidersClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified Security Partner Provider.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityPartnerProvidersClient) deleteOperation(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, securityPartnerProviderName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecurityPartnerProvidersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityPartnerProviderName == "" {
		return nil, errors.New("parameter securityPartnerProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SecurityPartnerProvidersClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the specified Security Partner Provider.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityPartnerProvidersClient) Get(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersGetOptions) (SecurityPartnerProviderResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityPartnerProviderName, options)
	if err != nil {
		return SecurityPartnerProviderResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SecurityPartnerProviderResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SecurityPartnerProviderResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityPartnerProvidersClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityPartnerProviderName == "" {
		return nil, errors.New("parameter securityPartnerProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecurityPartnerProvidersClient) getHandleResponse(resp *azcore.Response) (SecurityPartnerProviderResponse, error) {
	var val *SecurityPartnerProvider
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecurityPartnerProviderResponse{}, err
	}
	return SecurityPartnerProviderResponse{RawResponse: resp.Response, SecurityPartnerProvider: val}, nil
}

// getHandleError handles the Get error response.
func (client *SecurityPartnerProvidersClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all the Security Partner Providers in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityPartnerProvidersClient) List(options *SecurityPartnerProvidersListOptions) SecurityPartnerProviderListResultPager {
	return &securityPartnerProviderListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp SecurityPartnerProviderListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SecurityPartnerProviderListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *SecurityPartnerProvidersClient) listCreateRequest(ctx context.Context, options *SecurityPartnerProvidersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/securityPartnerProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecurityPartnerProvidersClient) listHandleResponse(resp *azcore.Response) (SecurityPartnerProviderListResultResponse, error) {
	var val *SecurityPartnerProviderListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecurityPartnerProviderListResultResponse{}, err
	}
	return SecurityPartnerProviderListResultResponse{RawResponse: resp.Response, SecurityPartnerProviderListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *SecurityPartnerProvidersClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - Lists all Security Partner Providers in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityPartnerProvidersClient) ListByResourceGroup(resourceGroupName string, options *SecurityPartnerProvidersListByResourceGroupOptions) SecurityPartnerProviderListResultPager {
	return &securityPartnerProviderListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp SecurityPartnerProviderListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SecurityPartnerProviderListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SecurityPartnerProvidersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SecurityPartnerProvidersListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SecurityPartnerProvidersClient) listByResourceGroupHandleResponse(resp *azcore.Response) (SecurityPartnerProviderListResultResponse, error) {
	var val *SecurityPartnerProviderListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecurityPartnerProviderListResultResponse{}, err
	}
	return SecurityPartnerProviderListResultResponse{RawResponse: resp.Response, SecurityPartnerProviderListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *SecurityPartnerProvidersClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UpdateTags - Updates tags of a Security Partner Provider resource.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityPartnerProvidersClient) UpdateTags(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters TagsObject, options *SecurityPartnerProvidersUpdateTagsOptions) (SecurityPartnerProviderResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, securityPartnerProviderName, parameters, options)
	if err != nil {
		return SecurityPartnerProviderResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SecurityPartnerProviderResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SecurityPartnerProviderResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SecurityPartnerProvidersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters TagsObject, options *SecurityPartnerProvidersUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityPartnerProviderName == "" {
		return nil, errors.New("parameter securityPartnerProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SecurityPartnerProvidersClient) updateTagsHandleResponse(resp *azcore.Response) (SecurityPartnerProviderResponse, error) {
	var val *SecurityPartnerProvider
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecurityPartnerProviderResponse{}, err
	}
	return SecurityPartnerProviderResponse{RawResponse: resp.Response, SecurityPartnerProvider: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *SecurityPartnerProvidersClient) updateTagsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
