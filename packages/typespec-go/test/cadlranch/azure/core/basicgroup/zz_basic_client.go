// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package basicgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// BasicClient - Illustrates bodies templated with Azure Core
// Don't use this type directly, use NewBasicClientWithNoCredential() instead.
type BasicClient struct {
	internal *azcore.Client
}

// BasicClientOptions contains the optional values for creating a [BasicClient].
type BasicClientOptions struct {
	azcore.ClientOptions
}

// NewBasicClientWithNoCredential creates a new instance of [BasicClient] with the specified values.
//   - options - BasicClientOptions contains the optional values for creating a [BasicClient]
func NewBasicClientWithNoCredential(options *BasicClientOptions) (*BasicClient, error) {
	if options == nil {
		options = &BasicClientOptions{}
	}
	cl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	basicClient := &BasicClient{
		internal: cl,
	}
	return basicClient, nil
}

// NewBasicTwoModelsAsPageItemClient creates a new instance of [BasicTwoModelsAsPageItemClient].
func (client *BasicClient) NewBasicTwoModelsAsPageItemClient() *BasicTwoModelsAsPageItemClient {
	return &BasicTwoModelsAsPageItemClient{
		internal: client.internal,
	}
}

// CreateOrReplace - Adds a user or replaces a user's fields.
//   - id - The user's id.
//   - resource - The resource instance.
//   - options - BasicClientCreateOrReplaceOptions contains the optional parameters for the BasicClient.CreateOrReplace method.
func (client *BasicClient) CreateOrReplace(ctx context.Context, id int32, resource User, options *BasicClientCreateOrReplaceOptions) (BasicClientCreateOrReplaceResponse, error) {
	var err error
	const operationName = "BasicClient.CreateOrReplace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrReplaceCreateRequest(ctx, id, resource, options)
	if err != nil {
		return BasicClientCreateOrReplaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientCreateOrReplaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientCreateOrReplaceResponse{}, err
	}
	resp, err := client.createOrReplaceHandleResponse(httpResp)
	return resp, err
}

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *BasicClient) createOrReplaceCreateRequest(ctx context.Context, id int32, resource User, options *BasicClientCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrReplaceHandleResponse handles the CreateOrReplace response.
func (client *BasicClient) createOrReplaceHandleResponse(resp *http.Response) (BasicClientCreateOrReplaceResponse, error) {
	result := BasicClientCreateOrReplaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return BasicClientCreateOrReplaceResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Adds a user or updates a user's fields.
//   - id - The user's id.
//   - resource - The resource instance.
//   - options - BasicClientCreateOrUpdateOptions contains the optional parameters for the BasicClient.CreateOrUpdate method.
func (client *BasicClient) CreateOrUpdate(ctx context.Context, id int32, resource User, options *BasicClientCreateOrUpdateOptions) (BasicClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "BasicClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, id, resource, options)
	if err != nil {
		return BasicClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BasicClient) createOrUpdateCreateRequest(ctx context.Context, id int32, resource User, options *BasicClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/merge-patch+json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BasicClient) createOrUpdateHandleResponse(resp *http.Response) (BasicClientCreateOrUpdateResponse, error) {
	result := BasicClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return BasicClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a user.
//   - id - The user's id.
//   - options - BasicClientDeleteOptions contains the optional parameters for the BasicClient.Delete method.
func (client *BasicClient) Delete(ctx context.Context, id int32, options *BasicClientDeleteOptions) (BasicClientDeleteResponse, error) {
	var err error
	const operationName = "BasicClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, id, options)
	if err != nil {
		return BasicClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientDeleteResponse{}, err
	}
	return BasicClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BasicClient) deleteCreateRequest(ctx context.Context, id int32, options *BasicClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Export - Exports a user.
//   - id - The user's id.
//   - formatParam - The format of the data.
//   - options - BasicClientExportOptions contains the optional parameters for the BasicClient.Export method.
func (client *BasicClient) Export(ctx context.Context, id int32, formatParam string, options *BasicClientExportOptions) (BasicClientExportResponse, error) {
	var err error
	const operationName = "BasicClient.Export"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportCreateRequest(ctx, id, formatParam, options)
	if err != nil {
		return BasicClientExportResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientExportResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientExportResponse{}, err
	}
	resp, err := client.exportHandleResponse(httpResp)
	return resp, err
}

// exportCreateRequest creates the Export request.
func (client *BasicClient) exportCreateRequest(ctx context.Context, id int32, formatParam string, options *BasicClientExportOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}:export"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	reqQP.Set("format", formatParam)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// exportHandleResponse handles the Export response.
func (client *BasicClient) exportHandleResponse(resp *http.Response) (BasicClientExportResponse, error) {
	result := BasicClientExportResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return BasicClientExportResponse{}, err
	}
	return result, nil
}

// Get - Gets a user.
//   - id - The user's id.
//   - options - BasicClientGetOptions contains the optional parameters for the BasicClient.Get method.
func (client *BasicClient) Get(ctx context.Context, id int32, options *BasicClientGetOptions) (BasicClientGetResponse, error) {
	var err error
	const operationName = "BasicClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, id, options)
	if err != nil {
		return BasicClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BasicClient) getCreateRequest(ctx context.Context, id int32, options *BasicClientGetOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BasicClient) getHandleResponse(resp *http.Response) (BasicClientGetResponse, error) {
	result := BasicClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return BasicClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all users.
//   - options - BasicClientListOptions contains the optional parameters for the BasicClient.NewListPager method.
func (client *BasicClient) NewListPager(options *BasicClientListOptions) *runtime.Pager[BasicClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BasicClientListResponse]{
		More: func(page BasicClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BasicClientListResponse) (BasicClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BasicClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return BasicClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *BasicClient) listCreateRequest(ctx context.Context, options *BasicClientListOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	if options != nil && options.Expand != nil {
		for _, qv := range options.Expand {
			reqQP.Add("expand", qv)
		}
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("maxpagesize", strconv.FormatInt(int64(*options.Maxpagesize), 10))
	}
	if options != nil && options.Orderby != nil {
		for _, qv := range options.Orderby {
			reqQP.Add("orderby", qv)
		}
	}
	if options != nil && options.SelectParam != nil {
		for _, qv := range options.SelectParam {
			reqQP.Add("select", qv)
		}
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BasicClient) listHandleResponse(resp *http.Response) (BasicClientListResponse, error) {
	result := BasicClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedUser); err != nil {
		return BasicClientListResponse{}, err
	}
	return result, nil
}

// NewListWithCustomPageModelPager - List with custom page model.
//   - options - BasicClientListWithCustomPageModelOptions contains the optional parameters for the BasicClient.NewListWithCustomPageModelPager
//     method.
func (client *BasicClient) NewListWithCustomPageModelPager(options *BasicClientListWithCustomPageModelOptions) *runtime.Pager[BasicClientListWithCustomPageModelResponse] {
	return runtime.NewPager(runtime.PagingHandler[BasicClientListWithCustomPageModelResponse]{
		More: func(page BasicClientListWithCustomPageModelResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BasicClientListWithCustomPageModelResponse) (BasicClientListWithCustomPageModelResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BasicClient.NewListWithCustomPageModelPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listWithCustomPageModelCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return BasicClientListWithCustomPageModelResponse{}, err
			}
			return client.listWithCustomPageModelHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listWithCustomPageModelCreateRequest creates the ListWithCustomPageModel request.
func (client *BasicClient) listWithCustomPageModelCreateRequest(ctx context.Context, options *BasicClientListWithCustomPageModelOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/custom-page"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWithCustomPageModelHandleResponse handles the ListWithCustomPageModel response.
func (client *BasicClient) listWithCustomPageModelHandleResponse(resp *http.Response) (BasicClientListWithCustomPageModelResponse, error) {
	result := BasicClientListWithCustomPageModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserListResults); err != nil {
		return BasicClientListWithCustomPageModelResponse{}, err
	}
	return result, nil
}

// NewListWithPagePager - List with Azure.Core.Page<>.
//   - options - BasicClientListWithPageOptions contains the optional parameters for the BasicClient.NewListWithPagePager method.
func (client *BasicClient) NewListWithPagePager(options *BasicClientListWithPageOptions) *runtime.Pager[BasicClientListWithPageResponse] {
	return runtime.NewPager(runtime.PagingHandler[BasicClientListWithPageResponse]{
		More: func(page BasicClientListWithPageResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BasicClientListWithPageResponse) (BasicClientListWithPageResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BasicClient.NewListWithPagePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listWithPageCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return BasicClientListWithPageResponse{}, err
			}
			return client.listWithPageHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listWithPageCreateRequest creates the ListWithPage request.
func (client *BasicClient) listWithPageCreateRequest(ctx context.Context, options *BasicClientListWithPageOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/page"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWithPageHandleResponse handles the ListWithPage response.
func (client *BasicClient) listWithPageHandleResponse(resp *http.Response) (BasicClientListWithPageResponse, error) {
	result := BasicClientListWithPageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedUser); err != nil {
		return BasicClientListWithPageResponse{}, err
	}
	return result, nil
}

// NewListWithParametersPager - List with extensible enum parameter Azure.Core.Page<>.
//   - bodyInput - The body of the input.
//   - options - BasicClientListWithParametersOptions contains the optional parameters for the BasicClient.NewListWithParametersPager
//     method.
func (client *BasicClient) NewListWithParametersPager(bodyInput ListItemInputBody, options *BasicClientListWithParametersOptions) *runtime.Pager[BasicClientListWithParametersResponse] {
	return runtime.NewPager(runtime.PagingHandler[BasicClientListWithParametersResponse]{
		More: func(page BasicClientListWithParametersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BasicClientListWithParametersResponse) (BasicClientListWithParametersResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BasicClient.NewListWithParametersPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listWithParametersCreateRequest(ctx, bodyInput, options)
			}, nil)
			if err != nil {
				return BasicClientListWithParametersResponse{}, err
			}
			return client.listWithParametersHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listWithParametersCreateRequest creates the ListWithParameters request.
func (client *BasicClient) listWithParametersCreateRequest(ctx context.Context, bodyInput ListItemInputBody, options *BasicClientListWithParametersOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/parameters"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Another != nil {
		reqQP.Set("another", string(*options.Another))
	}
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, bodyInput); err != nil {
		return nil, err
	}
	return req, nil
}

// listWithParametersHandleResponse handles the ListWithParameters response.
func (client *BasicClient) listWithParametersHandleResponse(resp *http.Response) (BasicClientListWithParametersResponse, error) {
	result := BasicClientListWithParametersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedUser); err != nil {
		return BasicClientListWithParametersResponse{}, err
	}
	return result, nil
}
