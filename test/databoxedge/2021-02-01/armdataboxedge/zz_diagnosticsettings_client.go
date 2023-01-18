//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"net/http"
	"net/url"
	"strings"
)

// DiagnosticSettingsClient contains the methods for the DiagnosticSettings group.
// Don't use this type directly, use NewDiagnosticSettingsClient() instead.
type DiagnosticSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDiagnosticSettingsClient creates a new instance of DiagnosticSettingsClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDiagnosticSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiagnosticSettingsClient, error) {
	cl, err := arm.NewClient("armdataboxedge.DiagnosticSettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DiagnosticSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetDiagnosticProactiveLogCollectionSettings - Gets the proactive log collection settings of the specified Data Box Edge/Data
// Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsOptions contains the optional parameters for
//     the DiagnosticSettingsClient.GetDiagnosticProactiveLogCollectionSettings method.
func (client *DiagnosticSettingsClient) GetDiagnosticProactiveLogCollectionSettings(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsOptions) (result DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "DiagnosticSettingsClient.GetDiagnosticProactiveLogCollectionSettings", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	result, err = client.getDiagnosticProactiveLogCollectionSettingsHandleResponse(resp)
	return
}

// getDiagnosticProactiveLogCollectionSettingsCreateRequest creates the GetDiagnosticProactiveLogCollectionSettings request.
func (client *DiagnosticSettingsClient) getDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticProactiveLogCollectionSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDiagnosticProactiveLogCollectionSettingsHandleResponse handles the GetDiagnosticProactiveLogCollectionSettings response.
func (client *DiagnosticSettingsClient) getDiagnosticProactiveLogCollectionSettingsHandleResponse(resp *http.Response) (result DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.DiagnosticProactiveLogCollectionSettings); err != nil {
		result = DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse{}
		return
	}
	return result, nil
}

// GetDiagnosticRemoteSupportSettings - Gets the diagnostic remote support settings of the specified Data Box Edge/Data Box
// Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsOptions contains the optional parameters for the DiagnosticSettingsClient.GetDiagnosticRemoteSupportSettings
//     method.
func (client *DiagnosticSettingsClient) GetDiagnosticRemoteSupportSettings(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsOptions) (result DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "DiagnosticSettingsClient.GetDiagnosticRemoteSupportSettings", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getDiagnosticRemoteSupportSettingsCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		err = runtime.NewResponseError(resp)
		return
	}
	result, err = client.getDiagnosticRemoteSupportSettingsHandleResponse(resp)
	return
}

// getDiagnosticRemoteSupportSettingsCreateRequest creates the GetDiagnosticRemoteSupportSettings request.
func (client *DiagnosticSettingsClient) getDiagnosticRemoteSupportSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticRemoteSupportSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDiagnosticRemoteSupportSettingsHandleResponse handles the GetDiagnosticRemoteSupportSettings response.
func (client *DiagnosticSettingsClient) getDiagnosticRemoteSupportSettingsHandleResponse(resp *http.Response) (result DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.DiagnosticRemoteSupportSettings); err != nil {
		result = DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse{}
		return
	}
	return result, nil
}

// BeginUpdateDiagnosticProactiveLogCollectionSettings - Updates the proactive log collection settings on a Data Box Edge/Data
// Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - proactiveLogCollectionSettings - The proactive log collection settings.
//   - options - DiagnosticSettingsClientBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions contains the optional parameters
//     for the DiagnosticSettingsClient.BeginUpdateDiagnosticProactiveLogCollectionSettings method.
func (client *DiagnosticSettingsClient) BeginUpdateDiagnosticProactiveLogCollectionSettings(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings DiagnosticProactiveLogCollectionSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (result *runtime.Poller[DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "DiagnosticSettingsClient.BeginUpdateDiagnosticProactiveLogCollectionSettings", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	if options == nil || options.ResumeToken == "" {
		var resp *http.Response
		resp, err = client.updateDiagnosticProactiveLogCollectionSettings(ctx, deviceName, resourceGroupName, proactiveLogCollectionSettings, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller[DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse](resp, client.internal.Pipeline(), nil)
	} else {
		result, err = runtime.NewPollerFromResumeToken[DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// UpdateDiagnosticProactiveLogCollectionSettings - Updates the proactive log collection settings on a Data Box Edge/Data
// Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *DiagnosticSettingsClient) updateDiagnosticProactiveLogCollectionSettings(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings DiagnosticProactiveLogCollectionSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (resp *http.Response, err error) {
	req, err := client.updateDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx, deviceName, resourceGroupName, proactiveLogCollectionSettings, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// updateDiagnosticProactiveLogCollectionSettingsCreateRequest creates the UpdateDiagnosticProactiveLogCollectionSettings request.
func (client *DiagnosticSettingsClient) updateDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings DiagnosticProactiveLogCollectionSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticProactiveLogCollectionSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, proactiveLogCollectionSettings)
}

// BeginUpdateDiagnosticRemoteSupportSettings - Updates the diagnostic remote support settings on a Data Box Edge/Data Box
// Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - diagnosticRemoteSupportSettings - The diagnostic remote support settings.
//   - options - DiagnosticSettingsClientBeginUpdateDiagnosticRemoteSupportSettingsOptions contains the optional parameters for
//     the DiagnosticSettingsClient.BeginUpdateDiagnosticRemoteSupportSettings method.
func (client *DiagnosticSettingsClient) BeginUpdateDiagnosticRemoteSupportSettings(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings DiagnosticRemoteSupportSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticRemoteSupportSettingsOptions) (result *runtime.Poller[DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "DiagnosticSettingsClient.BeginUpdateDiagnosticRemoteSupportSettings", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	if options == nil || options.ResumeToken == "" {
		var resp *http.Response
		resp, err = client.updateDiagnosticRemoteSupportSettings(ctx, deviceName, resourceGroupName, diagnosticRemoteSupportSettings, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller[DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse](resp, client.internal.Pipeline(), nil)
	} else {
		result, err = runtime.NewPollerFromResumeToken[DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// UpdateDiagnosticRemoteSupportSettings - Updates the diagnostic remote support settings on a Data Box Edge/Data Box Gateway
// device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *DiagnosticSettingsClient) updateDiagnosticRemoteSupportSettings(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings DiagnosticRemoteSupportSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticRemoteSupportSettingsOptions) (resp *http.Response, err error) {
	req, err := client.updateDiagnosticRemoteSupportSettingsCreateRequest(ctx, deviceName, resourceGroupName, diagnosticRemoteSupportSettings, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// updateDiagnosticRemoteSupportSettingsCreateRequest creates the UpdateDiagnosticRemoteSupportSettings request.
func (client *DiagnosticSettingsClient) updateDiagnosticRemoteSupportSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings DiagnosticRemoteSupportSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticRemoteSupportSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticRemoteSupportSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, diagnosticRemoteSupportSettings)
}
