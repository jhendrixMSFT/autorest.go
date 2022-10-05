//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GalleryImagesClient contains the methods for the GalleryImages group.
// Don't use this type directly, use NewGalleryImagesClient() instead.
type GalleryImagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGalleryImagesClient creates a new instance of GalleryImagesClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGalleryImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GalleryImagesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &GalleryImagesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a gallery Image Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-12-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Image Gallery in which the Image Definition is to be created.
//   - galleryImageName - The name of the gallery Image Definition to be created or updated. The allowed characters are alphabets
//     and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80
//     characters.
//   - galleryImage - Parameters supplied to the create or update gallery image operation.
//   - options - GalleryImagesClientBeginCreateOrUpdateOptions contains the optional parameters for the GalleryImagesClient.BeginCreateOrUpdate
//     method.
func (client *GalleryImagesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImage, options *GalleryImagesClientBeginCreateOrUpdateOptions) (*runtime.Poller[GalleryImagesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, galleryImageName, galleryImage, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GalleryImagesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GalleryImagesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a gallery Image Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-12-01
func (client *GalleryImagesClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImage, options *GalleryImagesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, galleryImage, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GalleryImagesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImage, options *GalleryImagesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, galleryImage)
}

// BeginDelete - Delete a gallery image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-12-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Image Gallery in which the Image Definition is to be deleted.
//   - galleryImageName - The name of the gallery Image Definition to be deleted.
//   - options - GalleryImagesClientBeginDeleteOptions contains the optional parameters for the GalleryImagesClient.BeginDelete
//     method.
func (client *GalleryImagesClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesClientBeginDeleteOptions) (*runtime.Poller[GalleryImagesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, galleryName, galleryImageName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GalleryImagesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GalleryImagesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a gallery image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-12-01
func (client *GalleryImagesClient) deleteOperation(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, options)
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
func (client *GalleryImagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about a gallery Image Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-12-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Image Gallery from which the Image Definitions are to be retrieved.
//   - galleryImageName - The name of the gallery Image Definition to be retrieved.
//   - options - GalleryImagesClientGetOptions contains the optional parameters for the GalleryImagesClient.Get method.
func (client *GalleryImagesClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesClientGetOptions) (GalleryImagesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, options)
	if err != nil {
		return GalleryImagesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GalleryImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GalleryImagesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GalleryImagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleryImagesClient) getHandleResponse(resp *http.Response) (GalleryImagesClientGetResponse, error) {
	result := GalleryImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryImage); err != nil {
		return GalleryImagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByGalleryPager - List gallery Image Definitions in a gallery.
//
// Generated from API version 2019-12-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Image Gallery from which Image Definitions are to be listed.
//   - options - GalleryImagesClientListByGalleryOptions contains the optional parameters for the GalleryImagesClient.ListByGallery
//     method.
func (client *GalleryImagesClient) NewListByGalleryPager(resourceGroupName string, galleryName string, options *GalleryImagesClientListByGalleryOptions) *runtime.Pager[GalleryImagesClientListByGalleryResponse] {
	return runtime.NewPager(runtime.PagingHandler[GalleryImagesClientListByGalleryResponse]{
		More: func(page GalleryImagesClientListByGalleryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GalleryImagesClientListByGalleryResponse) (GalleryImagesClientListByGalleryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByGalleryCreateRequest(ctx, resourceGroupName, galleryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GalleryImagesClientListByGalleryResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GalleryImagesClientListByGalleryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GalleryImagesClientListByGalleryResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByGalleryHandleResponse(resp)
		},
	})
}

// listByGalleryCreateRequest creates the ListByGallery request.
func (client *GalleryImagesClient) listByGalleryCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleryImagesClientListByGalleryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByGalleryHandleResponse handles the ListByGallery response.
func (client *GalleryImagesClient) listByGalleryHandleResponse(resp *http.Response) (GalleryImagesClientListByGalleryResponse, error) {
	result := GalleryImagesClientListByGalleryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryImageList); err != nil {
		return GalleryImagesClientListByGalleryResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a gallery Image Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-12-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Image Gallery in which the Image Definition is to be updated.
//   - galleryImageName - The name of the gallery Image Definition to be updated. The allowed characters are alphabets and numbers
//     with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
//   - galleryImage - Parameters supplied to the update gallery image operation.
//   - options - GalleryImagesClientBeginUpdateOptions contains the optional parameters for the GalleryImagesClient.BeginUpdate
//     method.
func (client *GalleryImagesClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImageUpdate, options *GalleryImagesClientBeginUpdateOptions) (*runtime.Poller[GalleryImagesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, galleryName, galleryImageName, galleryImage, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GalleryImagesClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GalleryImagesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a gallery Image Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-12-01
func (client *GalleryImagesClient) update(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImageUpdate, options *GalleryImagesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, galleryImage, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *GalleryImagesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImageUpdate, options *GalleryImagesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, galleryImage)
}
