//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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
	"strconv"
	"strings"
)

// VirtualMachineImagesClient contains the methods for the VirtualMachineImages group.
// Don't use this type directly, use NewVirtualMachineImagesClient() instead.
type VirtualMachineImagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineImagesClient creates a new instance of VirtualMachineImagesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineImagesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &VirtualMachineImagesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// Get - Gets a virtual machine image.
// If the operation fails it returns a generic error.
// location - The name of a supported Azure region.
// publisherName - A valid image publisher.
// offer - A valid image publisher offer.
// skus - A valid image SKU.
// version - A valid image SKU version.
// options - VirtualMachineImagesGetOptions contains the optional parameters for the VirtualMachineImages.Get method.
func (client *VirtualMachineImagesClient) Get(ctx context.Context, location string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesGetOptions) (VirtualMachineImagesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, publisherName, offer, skus, version, options)
	if err != nil {
		return VirtualMachineImagesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineImagesClient) getCreateRequest(ctx context.Context, location string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineImagesClient) getHandleResponse(resp *http.Response) (VirtualMachineImagesGetResponse, error) {
	result := VirtualMachineImagesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImage); err != nil {
		return VirtualMachineImagesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineImagesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
// If the operation fails it returns a generic error.
// location - The name of a supported Azure region.
// publisherName - A valid image publisher.
// offer - A valid image publisher offer.
// skus - A valid image SKU.
// options - VirtualMachineImagesListOptions contains the optional parameters for the VirtualMachineImages.List method.
func (client *VirtualMachineImagesClient) List(ctx context.Context, location string, publisherName string, offer string, skus string, options *VirtualMachineImagesListOptions) (VirtualMachineImagesListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, publisherName, offer, skus, options)
	if err != nil {
		return VirtualMachineImagesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineImagesClient) listCreateRequest(ctx context.Context, location string, publisherName string, offer string, skus string, options *VirtualMachineImagesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineImagesClient) listHandleResponse(resp *http.Response) (VirtualMachineImagesListResponse, error) {
	result := VirtualMachineImagesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineImagesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListOffers - Gets a list of virtual machine image offers for the specified location and publisher.
// If the operation fails it returns a generic error.
// location - The name of a supported Azure region.
// publisherName - A valid image publisher.
// options - VirtualMachineImagesListOffersOptions contains the optional parameters for the VirtualMachineImages.ListOffers
// method.
func (client *VirtualMachineImagesClient) ListOffers(ctx context.Context, location string, publisherName string, options *VirtualMachineImagesListOffersOptions) (VirtualMachineImagesListOffersResponse, error) {
	req, err := client.listOffersCreateRequest(ctx, location, publisherName, options)
	if err != nil {
		return VirtualMachineImagesListOffersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesListOffersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesListOffersResponse{}, client.listOffersHandleError(resp)
	}
	return client.listOffersHandleResponse(resp)
}

// listOffersCreateRequest creates the ListOffers request.
func (client *VirtualMachineImagesClient) listOffersCreateRequest(ctx context.Context, location string, publisherName string, options *VirtualMachineImagesListOffersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOffersHandleResponse handles the ListOffers response.
func (client *VirtualMachineImagesClient) listOffersHandleResponse(resp *http.Response) (VirtualMachineImagesListOffersResponse, error) {
	result := VirtualMachineImagesListOffersResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesListOffersResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listOffersHandleError handles the ListOffers error response.
func (client *VirtualMachineImagesClient) listOffersHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListPublishers - Gets a list of virtual machine image publishers for the specified Azure location.
// If the operation fails it returns a generic error.
// location - The name of a supported Azure region.
// options - VirtualMachineImagesListPublishersOptions contains the optional parameters for the VirtualMachineImages.ListPublishers
// method.
func (client *VirtualMachineImagesClient) ListPublishers(ctx context.Context, location string, options *VirtualMachineImagesListPublishersOptions) (VirtualMachineImagesListPublishersResponse, error) {
	req, err := client.listPublishersCreateRequest(ctx, location, options)
	if err != nil {
		return VirtualMachineImagesListPublishersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesListPublishersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesListPublishersResponse{}, client.listPublishersHandleError(resp)
	}
	return client.listPublishersHandleResponse(resp)
}

// listPublishersCreateRequest creates the ListPublishers request.
func (client *VirtualMachineImagesClient) listPublishersCreateRequest(ctx context.Context, location string, options *VirtualMachineImagesListPublishersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listPublishersHandleResponse handles the ListPublishers response.
func (client *VirtualMachineImagesClient) listPublishersHandleResponse(resp *http.Response) (VirtualMachineImagesListPublishersResponse, error) {
	result := VirtualMachineImagesListPublishersResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesListPublishersResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listPublishersHandleError handles the ListPublishers error response.
func (client *VirtualMachineImagesClient) listPublishersHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListSKUs - Gets a list of virtual machine image SKUs for the specified location, publisher, and offer.
// If the operation fails it returns a generic error.
// location - The name of a supported Azure region.
// publisherName - A valid image publisher.
// offer - A valid image publisher offer.
// options - VirtualMachineImagesListSKUsOptions contains the optional parameters for the VirtualMachineImages.ListSKUs method.
func (client *VirtualMachineImagesClient) ListSKUs(ctx context.Context, location string, publisherName string, offer string, options *VirtualMachineImagesListSKUsOptions) (VirtualMachineImagesListSKUsResponse, error) {
	req, err := client.listSKUsCreateRequest(ctx, location, publisherName, offer, options)
	if err != nil {
		return VirtualMachineImagesListSKUsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesListSKUsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesListSKUsResponse{}, client.listSKUsHandleError(resp)
	}
	return client.listSKUsHandleResponse(resp)
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *VirtualMachineImagesClient) listSKUsCreateRequest(ctx context.Context, location string, publisherName string, offer string, options *VirtualMachineImagesListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *VirtualMachineImagesClient) listSKUsHandleResponse(resp *http.Response) (VirtualMachineImagesListSKUsResponse, error) {
	result := VirtualMachineImagesListSKUsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesListSKUsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listSKUsHandleError handles the ListSKUs error response.
func (client *VirtualMachineImagesClient) listSKUsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
