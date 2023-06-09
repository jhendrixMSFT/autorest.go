//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
)

// ApplicationSecurityGroupsServer is a fake server for instances of the armnetwork.ApplicationSecurityGroupsClient type.
type ApplicationSecurityGroupsServer struct {
	// BeginCreateOrUpdate is the fake for method ApplicationSecurityGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters armnetwork.ApplicationSecurityGroup, options *armnetwork.ApplicationSecurityGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ApplicationSecurityGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ApplicationSecurityGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *armnetwork.ApplicationSecurityGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ApplicationSecurityGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationSecurityGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *armnetwork.ApplicationSecurityGroupsClientGetOptions) (resp azfake.Responder[armnetwork.ApplicationSecurityGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ApplicationSecurityGroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.ApplicationSecurityGroupsClientListOptions) (resp azfake.PagerResponder[armnetwork.ApplicationSecurityGroupsClientListResponse])

	// NewListAllPager is the fake for method ApplicationSecurityGroupsClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.ApplicationSecurityGroupsClientListAllOptions) (resp azfake.PagerResponder[armnetwork.ApplicationSecurityGroupsClientListAllResponse])

	// UpdateTags is the fake for method ApplicationSecurityGroupsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters armnetwork.TagsObject, options *armnetwork.ApplicationSecurityGroupsClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.ApplicationSecurityGroupsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewApplicationSecurityGroupsServerTransport creates a new instance of ApplicationSecurityGroupsServerTransport with the provided implementation.
// The returned ApplicationSecurityGroupsServerTransport instance is connected to an instance of armnetwork.ApplicationSecurityGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationSecurityGroupsServerTransport(srv *ApplicationSecurityGroupsServer) *ApplicationSecurityGroupsServerTransport {
	return &ApplicationSecurityGroupsServerTransport{srv: srv}
}

// ApplicationSecurityGroupsServerTransport connects instances of armnetwork.ApplicationSecurityGroupsClient to instances of ApplicationSecurityGroupsServer.
// Don't use this type directly, use NewApplicationSecurityGroupsServerTransport instead.
type ApplicationSecurityGroupsServerTransport struct {
	srv                 *ApplicationSecurityGroupsServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.ApplicationSecurityGroupsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.ApplicationSecurityGroupsClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.ApplicationSecurityGroupsClientListResponse]
	newListAllPager     *azfake.PagerResponder[armnetwork.ApplicationSecurityGroupsClientListAllResponse]
}

// Do implements the policy.Transporter interface for ApplicationSecurityGroupsServerTransport.
func (a *ApplicationSecurityGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationSecurityGroupsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "ApplicationSecurityGroupsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "ApplicationSecurityGroupsClient.Get":
		resp, err = a.dispatchGet(req)
	case "ApplicationSecurityGroupsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "ApplicationSecurityGroupsClient.NewListAllPager":
		resp, err = a.dispatchNewListAllPager(req)
	case "ApplicationSecurityGroupsClient.UpdateTags":
		resp, err = a.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationSecurityGroupsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if a.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/applicationSecurityGroups/(?P<applicationSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ApplicationSecurityGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationSecurityGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("applicationSecurityGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, applicationSecurityGroupNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(a.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginCreateOrUpdate) {
		a.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (a *ApplicationSecurityGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if a.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/applicationSecurityGroups/(?P<applicationSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationSecurityGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("applicationSecurityGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, applicationSecurityGroupNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(a.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginDelete) {
		a.beginDelete = nil
	}

	return resp, nil
}

func (a *ApplicationSecurityGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/applicationSecurityGroups/(?P<applicationSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	applicationSecurityGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("applicationSecurityGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameUnescaped, applicationSecurityGroupNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationSecurityGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationSecurityGroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if a.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/applicationSecurityGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameUnescaped, nil)
		a.newListPager = &resp
		server.PagerResponderInjectNextLinks(a.newListPager, req, func(page *armnetwork.ApplicationSecurityGroupsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListPager) {
		a.newListPager = nil
	}
	return resp, nil
}

func (a *ApplicationSecurityGroupsServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	if a.newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/applicationSecurityGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListAllPager(nil)
		a.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(a.newListAllPager, req, func(page *armnetwork.ApplicationSecurityGroupsClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListAllPager) {
		a.newListAllPager = nil
	}
	return resp, nil
}

func (a *ApplicationSecurityGroupsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if a.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/applicationSecurityGroups/(?P<applicationSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	applicationSecurityGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("applicationSecurityGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, applicationSecurityGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationSecurityGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
