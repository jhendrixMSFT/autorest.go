//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"net/http"
	"net/url"
	"strings"
)

// SQLScriptClient contains the methods for the SQLScript group.
// Don't use this type directly, use a constructor function instead.
type SQLScriptClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateOrUpdateSQLScript - Creates or updates a Sql Script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlScriptName - The sql script name.
//   - sqlScript - Sql Script resource definition.
//   - options - SQLScriptClientBeginCreateOrUpdateSQLScriptOptions contains the optional parameters for the SQLScriptClient.BeginCreateOrUpdateSQLScript
//     method.
func (client *SQLScriptClient) BeginCreateOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptClientBeginCreateOrUpdateSQLScriptOptions) (result *runtime.Poller[SQLScriptClientCreateOrUpdateSQLScriptResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "SQLScriptClient.BeginCreateOrUpdateSQLScript", &tracing.SpanOptions{
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
		resp, err = client.createOrUpdateSQLScript(ctx, sqlScriptName, sqlScript, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller[SQLScriptClientCreateOrUpdateSQLScriptResponse](resp, client.internal.Pipeline(), nil)
	} else {
		result, err = runtime.NewPollerFromResumeToken[SQLScriptClientCreateOrUpdateSQLScriptResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// CreateOrUpdateSQLScript - Creates or updates a Sql Script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *SQLScriptClient) createOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptClientBeginCreateOrUpdateSQLScriptOptions) (resp *http.Response, err error) {
	req, err := client.createOrUpdateSQLScriptCreateRequest(ctx, sqlScriptName, sqlScript, options)
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

// createOrUpdateSQLScriptCreateRequest creates the CreateOrUpdateSQLScript request.
func (client *SQLScriptClient) createOrUpdateSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptClientBeginCreateOrUpdateSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, sqlScript)
}

// BeginDeleteSQLScript - Deletes a Sql Script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlScriptName - The sql script name.
//   - options - SQLScriptClientBeginDeleteSQLScriptOptions contains the optional parameters for the SQLScriptClient.BeginDeleteSQLScript
//     method.
func (client *SQLScriptClient) BeginDeleteSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptClientBeginDeleteSQLScriptOptions) (result *runtime.Poller[SQLScriptClientDeleteSQLScriptResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "SQLScriptClient.BeginDeleteSQLScript", &tracing.SpanOptions{
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
		resp, err = client.deleteSQLScript(ctx, sqlScriptName, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller[SQLScriptClientDeleteSQLScriptResponse](resp, client.internal.Pipeline(), nil)
	} else {
		result, err = runtime.NewPollerFromResumeToken[SQLScriptClientDeleteSQLScriptResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// DeleteSQLScript - Deletes a Sql Script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *SQLScriptClient) deleteSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptClientBeginDeleteSQLScriptOptions) (resp *http.Response, err error) {
	req, err := client.deleteSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return
	}
	resp, err = client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(resp)
		return
	}
	return
}

// deleteSQLScriptCreateRequest creates the DeleteSQLScript request.
func (client *SQLScriptClient) deleteSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SQLScriptClientBeginDeleteSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetSQLScript - Gets a sql script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlScriptName - The sql script name.
//   - options - SQLScriptClientGetSQLScriptOptions contains the optional parameters for the SQLScriptClient.GetSQLScript method.
func (client *SQLScriptClient) GetSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptClientGetSQLScriptOptions) (result SQLScriptClientGetSQLScriptResponse, err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "SQLScriptClient.GetSQLScript", &tracing.SpanOptions{
		Kind: tracing.SpanKindInternal,
	})
	defer func() {
		if err != nil {
			span.AddError(err)
		}
		span.End()
	}()
	req, err := client.getSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		err = runtime.NewResponseError(resp)
		return
	}
	result, err = client.getSQLScriptHandleResponse(resp)
	return
}

// getSQLScriptCreateRequest creates the GetSQLScript request.
func (client *SQLScriptClient) getSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SQLScriptClientGetSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSQLScriptHandleResponse handles the GetSQLScript response.
func (client *SQLScriptClient) getSQLScriptHandleResponse(resp *http.Response) (result SQLScriptClientGetSQLScriptResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.SQLScriptResource); err != nil {
		result = SQLScriptClientGetSQLScriptResponse{}
		return
	}
	return result, nil
}

// NewGetSQLScriptsByWorkspacePager - Lists sql scripts.
//
// Generated from API version 2019-06-01-preview
//   - options - SQLScriptClientGetSQLScriptsByWorkspaceOptions contains the optional parameters for the SQLScriptClient.NewGetSQLScriptsByWorkspacePager
//     method.
func (client *SQLScriptClient) NewGetSQLScriptsByWorkspacePager(options *SQLScriptClientGetSQLScriptsByWorkspaceOptions) *runtime.Pager[SQLScriptClientGetSQLScriptsByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLScriptClientGetSQLScriptsByWorkspaceResponse]{
		More: func(page SQLScriptClientGetSQLScriptsByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLScriptClientGetSQLScriptsByWorkspaceResponse) (result SQLScriptClientGetSQLScriptsByWorkspaceResponse, err error) {
			ctx, span := client.internal.Tracer().Start(ctx, "SQLScriptClient.NewGetSQLScriptsByWorkspacePager", &tracing.SpanOptions{
				Kind: tracing.SpanKindInternal,
			})
			defer func() {
				if err != nil {
					span.AddError(err)
				}
				span.End()
			}()
			var req *policy.Request
			if page == nil {
				req, err = client.getSQLScriptsByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
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
			result, err = client.getSQLScriptsByWorkspaceHandleResponse(resp)
			return
		},
	})
}

// getSQLScriptsByWorkspaceCreateRequest creates the GetSQLScriptsByWorkspace request.
func (client *SQLScriptClient) getSQLScriptsByWorkspaceCreateRequest(ctx context.Context, options *SQLScriptClientGetSQLScriptsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSQLScriptsByWorkspaceHandleResponse handles the GetSQLScriptsByWorkspace response.
func (client *SQLScriptClient) getSQLScriptsByWorkspaceHandleResponse(resp *http.Response) (result SQLScriptClientGetSQLScriptsByWorkspaceResponse, err error) {
	if err = runtime.UnmarshalAsJSON(resp, &result.SQLScriptsListResponse); err != nil {
		result = SQLScriptClientGetSQLScriptsByWorkspaceResponse{}
		return
	}
	return result, nil
}

// BeginRenameSQLScript - Renames a sqlScript.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlScriptName - The sql script name.
//   - request - proposed new name.
//   - options - SQLScriptClientBeginRenameSQLScriptOptions contains the optional parameters for the SQLScriptClient.BeginRenameSQLScript
//     method.
func (client *SQLScriptClient) BeginRenameSQLScript(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptClientBeginRenameSQLScriptOptions) (result *runtime.Poller[SQLScriptClientRenameSQLScriptResponse], err error) {
	ctx, span := client.internal.Tracer().Start(ctx, "SQLScriptClient.BeginRenameSQLScript", &tracing.SpanOptions{
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
		resp, err = client.renameSQLScript(ctx, sqlScriptName, request, options)
		if err != nil {
			return
		}
		result, err = runtime.NewPoller[SQLScriptClientRenameSQLScriptResponse](resp, client.internal.Pipeline(), nil)
	} else {
		result, err = runtime.NewPollerFromResumeToken[SQLScriptClientRenameSQLScriptResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
	return
}

// RenameSQLScript - Renames a sqlScript.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *SQLScriptClient) renameSQLScript(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptClientBeginRenameSQLScriptOptions) (resp *http.Response, err error) {
	req, err := client.renameSQLScriptCreateRequest(ctx, sqlScriptName, request, options)
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

// renameSQLScriptCreateRequest creates the RenameSQLScript request.
func (client *SQLScriptClient) renameSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptClientBeginRenameSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}/rename"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}
