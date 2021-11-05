//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type triggerClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newTriggerClient creates a new instance of triggerClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newTriggerClient(endpoint string, pl runtime.Pipeline) *triggerClient {
	client := &triggerClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// BeginCreateOrUpdateTrigger - Creates or updates a trigger.
// If the operation fails it returns the *CloudError error type.
// triggerName - The trigger name.
// trigger - Trigger resource definition.
// options - TriggerBeginCreateOrUpdateTriggerOptions contains the optional parameters for the Trigger.BeginCreateOrUpdateTrigger
// method.
func (client *triggerClient) BeginCreateOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, options *TriggerBeginCreateOrUpdateTriggerOptions) (TriggerCreateOrUpdateTriggerPollerResponse, error) {
	resp, err := client.createOrUpdateTrigger(ctx, triggerName, trigger, options)
	if err != nil {
		return TriggerCreateOrUpdateTriggerPollerResponse{}, err
	}
	result := TriggerCreateOrUpdateTriggerPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("triggerClient.CreateOrUpdateTrigger", resp, client.pl, client.createOrUpdateTriggerHandleError)
	if err != nil {
		return TriggerCreateOrUpdateTriggerPollerResponse{}, err
	}
	result.Poller = &TriggerCreateOrUpdateTriggerPoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdateTrigger - Creates or updates a trigger.
// If the operation fails it returns the *CloudError error type.
func (client *triggerClient) createOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, options *TriggerBeginCreateOrUpdateTriggerOptions) (*http.Response, error) {
	req, err := client.createOrUpdateTriggerCreateRequest(ctx, triggerName, trigger, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateTriggerHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateTriggerCreateRequest creates the CreateOrUpdateTrigger request.
func (client *triggerClient) createOrUpdateTriggerCreateRequest(ctx context.Context, triggerName string, trigger TriggerResource, options *TriggerBeginCreateOrUpdateTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, trigger)
}

// createOrUpdateTriggerHandleError handles the CreateOrUpdateTrigger error response.
func (client *triggerClient) createOrUpdateTriggerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDeleteTrigger - Deletes a trigger.
// If the operation fails it returns the *CloudError error type.
// triggerName - The trigger name.
// options - TriggerBeginDeleteTriggerOptions contains the optional parameters for the Trigger.BeginDeleteTrigger method.
func (client *triggerClient) BeginDeleteTrigger(ctx context.Context, triggerName string, options *TriggerBeginDeleteTriggerOptions) (TriggerDeleteTriggerPollerResponse, error) {
	resp, err := client.deleteTrigger(ctx, triggerName, options)
	if err != nil {
		return TriggerDeleteTriggerPollerResponse{}, err
	}
	result := TriggerDeleteTriggerPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("triggerClient.DeleteTrigger", resp, client.pl, client.deleteTriggerHandleError)
	if err != nil {
		return TriggerDeleteTriggerPollerResponse{}, err
	}
	result.Poller = &TriggerDeleteTriggerPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteTrigger - Deletes a trigger.
// If the operation fails it returns the *CloudError error type.
func (client *triggerClient) deleteTrigger(ctx context.Context, triggerName string, options *TriggerBeginDeleteTriggerOptions) (*http.Response, error) {
	req, err := client.deleteTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteTriggerHandleError(resp)
	}
	return resp, nil
}

// deleteTriggerCreateRequest creates the DeleteTrigger request.
func (client *triggerClient) deleteTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginDeleteTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteTriggerHandleError handles the DeleteTrigger error response.
func (client *triggerClient) deleteTriggerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetEventSubscriptionStatus - Get a trigger's event subscription status.
// If the operation fails it returns the *CloudError error type.
// triggerName - The trigger name.
// options - TriggerGetEventSubscriptionStatusOptions contains the optional parameters for the Trigger.GetEventSubscriptionStatus
// method.
func (client *triggerClient) GetEventSubscriptionStatus(ctx context.Context, triggerName string, options *TriggerGetEventSubscriptionStatusOptions) (TriggerGetEventSubscriptionStatusResponse, error) {
	req, err := client.getEventSubscriptionStatusCreateRequest(ctx, triggerName, options)
	if err != nil {
		return TriggerGetEventSubscriptionStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerGetEventSubscriptionStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerGetEventSubscriptionStatusResponse{}, client.getEventSubscriptionStatusHandleError(resp)
	}
	return client.getEventSubscriptionStatusHandleResponse(resp)
}

// getEventSubscriptionStatusCreateRequest creates the GetEventSubscriptionStatus request.
func (client *triggerClient) getEventSubscriptionStatusCreateRequest(ctx context.Context, triggerName string, options *TriggerGetEventSubscriptionStatusOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/getEventSubscriptionStatus"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEventSubscriptionStatusHandleResponse handles the GetEventSubscriptionStatus response.
func (client *triggerClient) getEventSubscriptionStatusHandleResponse(resp *http.Response) (TriggerGetEventSubscriptionStatusResponse, error) {
	result := TriggerGetEventSubscriptionStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerSubscriptionOperationStatus); err != nil {
		return TriggerGetEventSubscriptionStatusResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getEventSubscriptionStatusHandleError handles the GetEventSubscriptionStatus error response.
func (client *triggerClient) getEventSubscriptionStatusHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetTrigger - Gets a trigger.
// If the operation fails it returns the *CloudError error type.
// triggerName - The trigger name.
// options - TriggerGetTriggerOptions contains the optional parameters for the Trigger.GetTrigger method.
func (client *triggerClient) GetTrigger(ctx context.Context, triggerName string, options *TriggerGetTriggerOptions) (TriggerGetTriggerResponse, error) {
	req, err := client.getTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return TriggerGetTriggerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerGetTriggerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return TriggerGetTriggerResponse{}, client.getTriggerHandleError(resp)
	}
	return client.getTriggerHandleResponse(resp)
}

// getTriggerCreateRequest creates the GetTrigger request.
func (client *triggerClient) getTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerGetTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getTriggerHandleResponse handles the GetTrigger response.
func (client *triggerClient) getTriggerHandleResponse(resp *http.Response) (TriggerGetTriggerResponse, error) {
	result := TriggerGetTriggerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerResource); err != nil {
		return TriggerGetTriggerResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getTriggerHandleError handles the GetTrigger error response.
func (client *triggerClient) getTriggerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetTriggersByWorkspace - Lists triggers.
// If the operation fails it returns the *CloudError error type.
// options - TriggerGetTriggersByWorkspaceOptions contains the optional parameters for the Trigger.GetTriggersByWorkspace
// method.
func (client *triggerClient) GetTriggersByWorkspace(options *TriggerGetTriggersByWorkspaceOptions) *TriggerGetTriggersByWorkspacePager {
	return &TriggerGetTriggersByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getTriggersByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp TriggerGetTriggersByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TriggerListResponse.NextLink)
		},
	}
}

// getTriggersByWorkspaceCreateRequest creates the GetTriggersByWorkspace request.
func (client *triggerClient) getTriggersByWorkspaceCreateRequest(ctx context.Context, options *TriggerGetTriggersByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/triggers"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getTriggersByWorkspaceHandleResponse handles the GetTriggersByWorkspace response.
func (client *triggerClient) getTriggersByWorkspaceHandleResponse(resp *http.Response) (TriggerGetTriggersByWorkspaceResponse, error) {
	result := TriggerGetTriggersByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerListResponse); err != nil {
		return TriggerGetTriggersByWorkspaceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getTriggersByWorkspaceHandleError handles the GetTriggersByWorkspace error response.
func (client *triggerClient) getTriggersByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginStartTrigger - Starts a trigger.
// If the operation fails it returns the *CloudError error type.
// triggerName - The trigger name.
// options - TriggerBeginStartTriggerOptions contains the optional parameters for the Trigger.BeginStartTrigger method.
func (client *triggerClient) BeginStartTrigger(ctx context.Context, triggerName string, options *TriggerBeginStartTriggerOptions) (TriggerStartTriggerPollerResponse, error) {
	resp, err := client.startTrigger(ctx, triggerName, options)
	if err != nil {
		return TriggerStartTriggerPollerResponse{}, err
	}
	result := TriggerStartTriggerPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("triggerClient.StartTrigger", resp, client.pl, client.startTriggerHandleError)
	if err != nil {
		return TriggerStartTriggerPollerResponse{}, err
	}
	result.Poller = &TriggerStartTriggerPoller{
		pt: pt,
	}
	return result, nil
}

// StartTrigger - Starts a trigger.
// If the operation fails it returns the *CloudError error type.
func (client *triggerClient) startTrigger(ctx context.Context, triggerName string, options *TriggerBeginStartTriggerOptions) (*http.Response, error) {
	req, err := client.startTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.startTriggerHandleError(resp)
	}
	return resp, nil
}

// startTriggerCreateRequest creates the StartTrigger request.
func (client *triggerClient) startTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginStartTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/start"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// startTriggerHandleError handles the StartTrigger error response.
func (client *triggerClient) startTriggerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginStopTrigger - Stops a trigger.
// If the operation fails it returns the *CloudError error type.
// triggerName - The trigger name.
// options - TriggerBeginStopTriggerOptions contains the optional parameters for the Trigger.BeginStopTrigger method.
func (client *triggerClient) BeginStopTrigger(ctx context.Context, triggerName string, options *TriggerBeginStopTriggerOptions) (TriggerStopTriggerPollerResponse, error) {
	resp, err := client.stopTrigger(ctx, triggerName, options)
	if err != nil {
		return TriggerStopTriggerPollerResponse{}, err
	}
	result := TriggerStopTriggerPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("triggerClient.StopTrigger", resp, client.pl, client.stopTriggerHandleError)
	if err != nil {
		return TriggerStopTriggerPollerResponse{}, err
	}
	result.Poller = &TriggerStopTriggerPoller{
		pt: pt,
	}
	return result, nil
}

// StopTrigger - Stops a trigger.
// If the operation fails it returns the *CloudError error type.
func (client *triggerClient) stopTrigger(ctx context.Context, triggerName string, options *TriggerBeginStopTriggerOptions) (*http.Response, error) {
	req, err := client.stopTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.stopTriggerHandleError(resp)
	}
	return resp, nil
}

// stopTriggerCreateRequest creates the StopTrigger request.
func (client *triggerClient) stopTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginStopTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/stop"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// stopTriggerHandleError handles the StopTrigger error response.
func (client *triggerClient) stopTriggerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginSubscribeTriggerToEvents - Subscribe event trigger to events.
// If the operation fails it returns the *CloudError error type.
// triggerName - The trigger name.
// options - TriggerBeginSubscribeTriggerToEventsOptions contains the optional parameters for the Trigger.BeginSubscribeTriggerToEvents
// method.
func (client *triggerClient) BeginSubscribeTriggerToEvents(ctx context.Context, triggerName string, options *TriggerBeginSubscribeTriggerToEventsOptions) (TriggerSubscribeTriggerToEventsPollerResponse, error) {
	resp, err := client.subscribeTriggerToEvents(ctx, triggerName, options)
	if err != nil {
		return TriggerSubscribeTriggerToEventsPollerResponse{}, err
	}
	result := TriggerSubscribeTriggerToEventsPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("triggerClient.SubscribeTriggerToEvents", resp, client.pl, client.subscribeTriggerToEventsHandleError)
	if err != nil {
		return TriggerSubscribeTriggerToEventsPollerResponse{}, err
	}
	result.Poller = &TriggerSubscribeTriggerToEventsPoller{
		pt: pt,
	}
	return result, nil
}

// SubscribeTriggerToEvents - Subscribe event trigger to events.
// If the operation fails it returns the *CloudError error type.
func (client *triggerClient) subscribeTriggerToEvents(ctx context.Context, triggerName string, options *TriggerBeginSubscribeTriggerToEventsOptions) (*http.Response, error) {
	req, err := client.subscribeTriggerToEventsCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.subscribeTriggerToEventsHandleError(resp)
	}
	return resp, nil
}

// subscribeTriggerToEventsCreateRequest creates the SubscribeTriggerToEvents request.
func (client *triggerClient) subscribeTriggerToEventsCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginSubscribeTriggerToEventsOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/subscribeToEvents"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// subscribeTriggerToEventsHandleError handles the SubscribeTriggerToEvents error response.
func (client *triggerClient) subscribeTriggerToEventsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUnsubscribeTriggerFromEvents - Unsubscribe event trigger from events.
// If the operation fails it returns the *CloudError error type.
// triggerName - The trigger name.
// options - TriggerBeginUnsubscribeTriggerFromEventsOptions contains the optional parameters for the Trigger.BeginUnsubscribeTriggerFromEvents
// method.
func (client *triggerClient) BeginUnsubscribeTriggerFromEvents(ctx context.Context, triggerName string, options *TriggerBeginUnsubscribeTriggerFromEventsOptions) (TriggerUnsubscribeTriggerFromEventsPollerResponse, error) {
	resp, err := client.unsubscribeTriggerFromEvents(ctx, triggerName, options)
	if err != nil {
		return TriggerUnsubscribeTriggerFromEventsPollerResponse{}, err
	}
	result := TriggerUnsubscribeTriggerFromEventsPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("triggerClient.UnsubscribeTriggerFromEvents", resp, client.pl, client.unsubscribeTriggerFromEventsHandleError)
	if err != nil {
		return TriggerUnsubscribeTriggerFromEventsPollerResponse{}, err
	}
	result.Poller = &TriggerUnsubscribeTriggerFromEventsPoller{
		pt: pt,
	}
	return result, nil
}

// UnsubscribeTriggerFromEvents - Unsubscribe event trigger from events.
// If the operation fails it returns the *CloudError error type.
func (client *triggerClient) unsubscribeTriggerFromEvents(ctx context.Context, triggerName string, options *TriggerBeginUnsubscribeTriggerFromEventsOptions) (*http.Response, error) {
	req, err := client.unsubscribeTriggerFromEventsCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.unsubscribeTriggerFromEventsHandleError(resp)
	}
	return resp, nil
}

// unsubscribeTriggerFromEventsCreateRequest creates the UnsubscribeTriggerFromEvents request.
func (client *triggerClient) unsubscribeTriggerFromEventsCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginUnsubscribeTriggerFromEventsOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/unsubscribeFromEvents"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// unsubscribeTriggerFromEventsHandleError handles the UnsubscribeTriggerFromEvents error response.
func (client *triggerClient) unsubscribeTriggerFromEventsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
