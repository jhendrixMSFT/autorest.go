// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// VirtualMachineScaleSetVMExtensionsClient contains the methods for the VirtualMachineScaleSetVMExtensions group.
// Don't use this type directly, use NewVirtualMachineScaleSetVMExtensionsClient() instead.
type VirtualMachineScaleSetVMExtensionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineScaleSetVMExtensionsClient creates a new instance of VirtualMachineScaleSetVMExtensionsClient with the specified values.
func NewVirtualMachineScaleSetVMExtensionsClient(con *armcore.Connection, subscriptionID string) *VirtualMachineScaleSetVMExtensionsClient {
	return &VirtualMachineScaleSetVMExtensionsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - The operation to create or update the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineScaleSetVMExtensionsBeginCreateOrUpdateOptions) (VirtualMachineExtensionPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineExtensionPollerResponse{}, err
	}
	result := VirtualMachineExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualMachineExtensionPollerResponse{}, err
	}
	poller := &virtualMachineExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualMachineExtensionPoller from the specified resume token.
// token - The value must come from a previous call to VirtualMachineExtensionPoller.ResumeToken().
func (client *VirtualMachineScaleSetVMExtensionsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (VirtualMachineExtensionPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualMachineExtensionPollerResponse{}, err
	}
	poller := &virtualMachineExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualMachineExtensionPollerResponse{}, err
	}
	result := VirtualMachineExtensionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - The operation to create or update the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineScaleSetVMExtensionsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, extensionParameters, options)
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
func (client *VirtualMachineScaleSetVMExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineScaleSetVMExtensionsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
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
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - The operation to delete the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMExtensionsClient.Delete", "", resp, client.deleteHandleError)
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
func (client *VirtualMachineScaleSetVMExtensionsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.Delete", token, client.deleteHandleError)
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

// Delete - The operation to delete the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, options)
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
func (client *VirtualMachineScaleSetVMExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
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
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - The operation to get the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsGetOptions) (VirtualMachineExtensionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, options)
	if err != nil {
		return VirtualMachineExtensionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineExtensionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineScaleSetVMExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetVMExtensionsClient) getHandleResponse(resp *azcore.Response) (VirtualMachineExtensionResponse, error) {
	var val *VirtualMachineExtension
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineExtensionResponse{}, err
	}
	return VirtualMachineExtensionResponse{RawResponse: resp.Response, VirtualMachineExtension: val}, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) getHandleError(resp *azcore.Response) error {
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

// List - The operation to get all extensions of an instance in Virtual Machine Scaleset.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) List(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMExtensionsListOptions) (VirtualMachineExtensionsListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, options)
	if err != nil {
		return VirtualMachineExtensionsListResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionsListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineExtensionsListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineScaleSetVMExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMExtensionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineScaleSetVMExtensionsClient) listHandleResponse(resp *azcore.Response) (VirtualMachineExtensionsListResultResponse, error) {
	var val *VirtualMachineExtensionsListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineExtensionsListResultResponse{}, err
	}
	return VirtualMachineExtensionsListResultResponse{RawResponse: resp.Response, VirtualMachineExtensionsListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) listHandleError(resp *azcore.Response) error {
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

// BeginUpdate - The operation to update the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineScaleSetVMExtensionsBeginUpdateOptions) (VirtualMachineExtensionPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineExtensionPollerResponse{}, err
	}
	result := VirtualMachineExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMExtensionsClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return VirtualMachineExtensionPollerResponse{}, err
	}
	poller := &virtualMachineExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new VirtualMachineExtensionPoller from the specified resume token.
// token - The value must come from a previous call to VirtualMachineExtensionPoller.ResumeToken().
func (client *VirtualMachineScaleSetVMExtensionsClient) ResumeUpdate(ctx context.Context, token string) (VirtualMachineExtensionPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.Update", token, client.updateHandleError)
	if err != nil {
		return VirtualMachineExtensionPollerResponse{}, err
	}
	poller := &virtualMachineExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualMachineExtensionPollerResponse{}, err
	}
	result := VirtualMachineExtensionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - The operation to update the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) update(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineScaleSetVMExtensionsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetVMExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineScaleSetVMExtensionsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
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
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// updateHandleError handles the Update error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) updateHandleError(resp *azcore.Response) error {
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
