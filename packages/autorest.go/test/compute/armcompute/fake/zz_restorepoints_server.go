//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armcompute"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"regexp"
)

// RestorePointsServer is a fake server for instances of the armcompute.RestorePointsClient type.
type RestorePointsServer struct {
	// BeginCreate is the fake for method RestorePointsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, parameters armcompute.RestorePoint, options *armcompute.RestorePointsClientBeginCreateOptions) (resp azfake.PollerResponder[armcompute.RestorePointsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RestorePointsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, options *armcompute.RestorePointsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.RestorePointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RestorePointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, options *armcompute.RestorePointsClientGetOptions) (resp azfake.Responder[armcompute.RestorePointsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewRestorePointsServerTransport creates a new instance of RestorePointsServerTransport with the provided implementation.
// The returned RestorePointsServerTransport instance is connected to an instance of armcompute.RestorePointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRestorePointsServerTransport(srv *RestorePointsServer) *RestorePointsServerTransport {
	return &RestorePointsServerTransport{srv: srv}
}

// RestorePointsServerTransport connects instances of armcompute.RestorePointsClient to instances of RestorePointsServer.
// Don't use this type directly, use NewRestorePointsServerTransport instead.
type RestorePointsServerTransport struct {
	srv         *RestorePointsServer
	beginCreate *azfake.PollerResponder[armcompute.RestorePointsClientCreateResponse]
	beginDelete *azfake.PollerResponder[armcompute.RestorePointsClientDeleteResponse]
}

// Do implements the policy.Transporter interface for RestorePointsServerTransport.
func (r *RestorePointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RestorePointsClient.BeginCreate":
		resp, err = r.dispatchBeginCreate(req)
	case "RestorePointsClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RestorePointsClient.Get":
		resp, err = r.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RestorePointsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	if r.beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints/(?P<restorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.RestorePoint](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		restorePointCollectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
		if err != nil {
			return nil, err
		}
		restorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreate(req.Context(), resourceGroupNameUnescaped, restorePointCollectionNameUnescaped, restorePointNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		r.beginCreate = &respr
	}

	resp, err := server.PollerResponderNext(r.beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(r.beginCreate) {
		r.beginCreate = nil
	}

	return resp, nil
}

func (r *RestorePointsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if r.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints/(?P<restorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		restorePointCollectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
		if err != nil {
			return nil, err
		}
		restorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, restorePointCollectionNameUnescaped, restorePointNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		r.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(r.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(r.beginDelete) {
		r.beginDelete = nil
	}

	return resp, nil
}

func (r *RestorePointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints/(?P<restorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	restorePointCollectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
	if err != nil {
		return nil, err
	}
	restorePointNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(armcompute.RestorePointExpandOptions(expandUnescaped))
	var options *armcompute.RestorePointsClientGetOptions
	if expandParam != nil {
		options = &armcompute.RestorePointsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameUnescaped, restorePointCollectionNameUnescaped, restorePointNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RestorePoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
