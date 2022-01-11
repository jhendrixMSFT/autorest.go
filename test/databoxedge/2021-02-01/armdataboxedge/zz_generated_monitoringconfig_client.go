//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MonitoringConfigClient contains the methods for the MonitoringConfig group.
// Don't use this type directly, use NewMonitoringConfigClient() instead.
type MonitoringConfigClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMonitoringConfigClient creates a new instance of MonitoringConfigClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMonitoringConfigClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MonitoringConfigClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &MonitoringConfigClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a new metric configuration or updates an existing one for a role.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// roleName - The role name.
// resourceGroupName - The resource group name.
// monitoringMetricConfiguration - The metric configuration.
// options - MonitoringConfigClientBeginCreateOrUpdateOptions contains the optional parameters for the MonitoringConfigClient.BeginCreateOrUpdate
// method.
func (client *MonitoringConfigClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, roleName string, resourceGroupName string, monitoringMetricConfiguration MonitoringMetricConfiguration, options *MonitoringConfigClientBeginCreateOrUpdateOptions) (MonitoringConfigClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, roleName, resourceGroupName, monitoringMetricConfiguration, options)
	if err != nil {
		return MonitoringConfigClientCreateOrUpdatePollerResponse{}, err
	}
	result := MonitoringConfigClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MonitoringConfigClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return MonitoringConfigClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &MonitoringConfigClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a new metric configuration or updates an existing one for a role.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MonitoringConfigClient) createOrUpdate(ctx context.Context, deviceName string, roleName string, resourceGroupName string, monitoringMetricConfiguration MonitoringMetricConfiguration, options *MonitoringConfigClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, roleName, resourceGroupName, monitoringMetricConfiguration, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MonitoringConfigClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, monitoringMetricConfiguration MonitoringMetricConfiguration, options *MonitoringConfigClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, monitoringMetricConfiguration)
}

// BeginDelete - deletes a new metric configuration for a role.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// roleName - The role name.
// resourceGroupName - The resource group name.
// options - MonitoringConfigClientBeginDeleteOptions contains the optional parameters for the MonitoringConfigClient.BeginDelete
// method.
func (client *MonitoringConfigClient) BeginDelete(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientBeginDeleteOptions) (MonitoringConfigClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, roleName, resourceGroupName, options)
	if err != nil {
		return MonitoringConfigClientDeletePollerResponse{}, err
	}
	result := MonitoringConfigClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MonitoringConfigClient.Delete", "", resp, client.pl)
	if err != nil {
		return MonitoringConfigClientDeletePollerResponse{}, err
	}
	result.Poller = &MonitoringConfigClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - deletes a new metric configuration for a role.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MonitoringConfigClient) deleteOperation(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MonitoringConfigClient) deleteCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a metric configuration of a role.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// roleName - The role name.
// resourceGroupName - The resource group name.
// options - MonitoringConfigClientGetOptions contains the optional parameters for the MonitoringConfigClient.Get method.
func (client *MonitoringConfigClient) Get(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientGetOptions) (MonitoringConfigClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
	if err != nil {
		return MonitoringConfigClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitoringConfigClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitoringConfigClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MonitoringConfigClient) getCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MonitoringConfigClient) getHandleResponse(resp *http.Response) (MonitoringConfigClientGetResponse, error) {
	result := MonitoringConfigClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringMetricConfiguration); err != nil {
		return MonitoringConfigClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists metric configurations in a role.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// roleName - The role name.
// resourceGroupName - The resource group name.
// options - MonitoringConfigClientListOptions contains the optional parameters for the MonitoringConfigClient.List method.
func (client *MonitoringConfigClient) List(deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientListOptions) *MonitoringConfigClientListPager {
	return &MonitoringConfigClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp MonitoringConfigClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MonitoringMetricConfigurationList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *MonitoringConfigClient) listCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MonitoringConfigClient) listHandleResponse(resp *http.Response) (MonitoringConfigClientListResponse, error) {
	result := MonitoringConfigClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringMetricConfigurationList); err != nil {
		return MonitoringConfigClientListResponse{}, err
	}
	return result, nil
}
