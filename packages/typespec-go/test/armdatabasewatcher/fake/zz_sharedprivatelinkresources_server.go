// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"armdatabasewatcher"
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

// SharedPrivateLinkResourcesServer is a fake server for instances of the armdatabasewatcher.SharedPrivateLinkResourcesClient type.
type SharedPrivateLinkResourcesServer struct {
	// BeginCreate is the fake for method SharedPrivateLinkResourcesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, resource armdatabasewatcher.SharedPrivateLinkResource, options *armdatabasewatcher.SharedPrivateLinkResourcesClientBeginCreateOptions) (resp azfake.PollerResponder[armdatabasewatcher.SharedPrivateLinkResourcesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SharedPrivateLinkResourcesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, options *armdatabasewatcher.SharedPrivateLinkResourcesClientBeginDeleteOptions) (resp azfake.PollerResponder[armdatabasewatcher.SharedPrivateLinkResourcesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SharedPrivateLinkResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, options *armdatabasewatcher.SharedPrivateLinkResourcesClientGetOptions) (resp azfake.Responder[armdatabasewatcher.SharedPrivateLinkResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByWatcherPager is the fake for method SharedPrivateLinkResourcesClient.NewListByWatcherPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWatcherPager func(resourceGroupName string, watcherName string, options *armdatabasewatcher.SharedPrivateLinkResourcesClientListByWatcherOptions) (resp azfake.PagerResponder[armdatabasewatcher.SharedPrivateLinkResourcesClientListByWatcherResponse])
}

// NewSharedPrivateLinkResourcesServerTransport creates a new instance of SharedPrivateLinkResourcesServerTransport with the provided implementation.
// The returned SharedPrivateLinkResourcesServerTransport instance is connected to an instance of armdatabasewatcher.SharedPrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSharedPrivateLinkResourcesServerTransport(srv *SharedPrivateLinkResourcesServer) *SharedPrivateLinkResourcesServerTransport {
	return &SharedPrivateLinkResourcesServerTransport{
		srv:                   srv,
		beginCreate:           newTracker[azfake.PollerResponder[armdatabasewatcher.SharedPrivateLinkResourcesClientCreateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armdatabasewatcher.SharedPrivateLinkResourcesClientDeleteResponse]](),
		newListByWatcherPager: newTracker[azfake.PagerResponder[armdatabasewatcher.SharedPrivateLinkResourcesClientListByWatcherResponse]](),
	}
}

// SharedPrivateLinkResourcesServerTransport connects instances of armdatabasewatcher.SharedPrivateLinkResourcesClient to instances of SharedPrivateLinkResourcesServer.
// Don't use this type directly, use NewSharedPrivateLinkResourcesServerTransport instead.
type SharedPrivateLinkResourcesServerTransport struct {
	srv                   *SharedPrivateLinkResourcesServer
	beginCreate           *tracker[azfake.PollerResponder[armdatabasewatcher.SharedPrivateLinkResourcesClientCreateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armdatabasewatcher.SharedPrivateLinkResourcesClientDeleteResponse]]
	newListByWatcherPager *tracker[azfake.PagerResponder[armdatabasewatcher.SharedPrivateLinkResourcesClientListByWatcherResponse]]
}

// Do implements the policy.Transporter interface for SharedPrivateLinkResourcesServerTransport.
func (s *SharedPrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SharedPrivateLinkResourcesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "SharedPrivateLinkResourcesClient.BeginCreate":
			res.resp, res.err = s.dispatchBeginCreate(req)
		case "SharedPrivateLinkResourcesClient.BeginDelete":
			res.resp, res.err = s.dispatchBeginDelete(req)
		case "SharedPrivateLinkResourcesClient.Get":
			res.resp, res.err = s.dispatchGet(req)
		case "SharedPrivateLinkResourcesClient.NewListByWatcherPager":
			res.resp, res.err = s.dispatchNewListByWatcherPager(req)
		default:
			res.err = fmt.Errorf("unhandled API %s", method)
		}

		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (s *SharedPrivateLinkResourcesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := s.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources/(?P<sharedPrivateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatabasewatcher.SharedPrivateLinkResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
		if err != nil {
			return nil, err
		}
		sharedPrivateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sharedPrivateLinkResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreate(req.Context(), resourceGroupNameParam, watcherNameParam, sharedPrivateLinkResourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		s.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		s.beginCreate.remove(req)
	}

	return resp, nil
}

func (s *SharedPrivateLinkResourcesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources/(?P<sharedPrivateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
		if err != nil {
			return nil, err
		}
		sharedPrivateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sharedPrivateLinkResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, watcherNameParam, sharedPrivateLinkResourceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SharedPrivateLinkResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources/(?P<sharedPrivateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
	if err != nil {
		return nil, err
	}
	sharedPrivateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sharedPrivateLinkResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, watcherNameParam, sharedPrivateLinkResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SharedPrivateLinkResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SharedPrivateLinkResourcesServerTransport) dispatchNewListByWatcherPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByWatcherPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWatcherPager not implemented")}
	}
	newListByWatcherPager := s.newListByWatcherPager.get(req)
	if newListByWatcherPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByWatcherPager(resourceGroupNameParam, watcherNameParam, nil)
		newListByWatcherPager = &resp
		s.newListByWatcherPager.add(req, newListByWatcherPager)
		server.PagerResponderInjectNextLinks(newListByWatcherPager, req, func(page *armdatabasewatcher.SharedPrivateLinkResourcesClientListByWatcherResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByWatcherPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByWatcherPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByWatcherPager) {
		s.newListByWatcherPager.remove(req)
	}
	return resp, nil
}
