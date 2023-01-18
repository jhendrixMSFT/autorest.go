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

// SupportPackagesClient contains the methods for the SupportPackages group.
// Don't use this type directly, use NewSupportPackagesClient() instead.
type SupportPackagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSupportPackagesClient creates a new instance of SupportPackagesClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSupportPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SupportPackagesClient, error) {
	cl, err := arm.NewClient("armdataboxedge.SupportPackagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SupportPackagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginTriggerSupportPackage - Triggers support package on the device
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - triggerSupportPackageRequest - The trigger support package request object
//   - options - SupportPackagesClientBeginTriggerSupportPackageOptions contains the optional parameters for the SupportPackagesClient.BeginTriggerSupportPackage
//     method.
func (client *SupportPackagesClient) BeginTriggerSupportPackage(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest TriggerSupportPackageRequest, options *SupportPackagesClientBeginTriggerSupportPackageOptions) (result *runtime.Poller[SupportPackagesClientTriggerSupportPackageResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "SupportPackagesClient.BeginTriggerSupportPackage", &tracing.SpanOptions{
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
		resp, err = client.triggerSupportPackage(ctx, deviceName, resourceGroupName, triggerSupportPackageRequest, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller[SupportPackagesClientTriggerSupportPackageResponse](resp, client.internal.Pipeline(), nil)
	} else {
		result, err = runtime.NewPollerFromResumeToken[SupportPackagesClientTriggerSupportPackageResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// TriggerSupportPackage - Triggers support package on the device
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *SupportPackagesClient) triggerSupportPackage(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest TriggerSupportPackageRequest, options *SupportPackagesClientBeginTriggerSupportPackageOptions) (resp *http.Response, err error) {
	req, err := client.triggerSupportPackageCreateRequest(ctx, deviceName, resourceGroupName, triggerSupportPackageRequest, options)
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

// triggerSupportPackageCreateRequest creates the TriggerSupportPackage request.
func (client *SupportPackagesClient) triggerSupportPackageCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest TriggerSupportPackageRequest, options *SupportPackagesClientBeginTriggerSupportPackageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggerSupportPackage"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, triggerSupportPackageRequest)
}
