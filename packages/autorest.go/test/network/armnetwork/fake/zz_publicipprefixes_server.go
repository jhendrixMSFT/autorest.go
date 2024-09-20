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

// PublicIPPrefixesServer is a fake server for instances of the armnetwork.PublicIPPrefixesClient type.
type PublicIPPrefixesServer struct {
	// BeginCreateOrUpdate is the fake for method PublicIPPrefixesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters armnetwork.PublicIPPrefix, options *armnetwork.PublicIPPrefixesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.PublicIPPrefixesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PublicIPPrefixesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *armnetwork.PublicIPPrefixesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.PublicIPPrefixesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PublicIPPrefixesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *armnetwork.PublicIPPrefixesClientGetOptions) (resp azfake.Responder[armnetwork.PublicIPPrefixesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PublicIPPrefixesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.PublicIPPrefixesClientListOptions) (resp azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListResponse])

	// NewListAllPager is the fake for method PublicIPPrefixesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.PublicIPPrefixesClientListAllOptions) (resp azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListAllResponse])

	// UpdateTags is the fake for method PublicIPPrefixesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters armnetwork.TagsObject, options *armnetwork.PublicIPPrefixesClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.PublicIPPrefixesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewPublicIPPrefixesServerTransport creates a new instance of PublicIPPrefixesServerTransport with the provided implementation.
// The returned PublicIPPrefixesServerTransport instance is connected to an instance of armnetwork.PublicIPPrefixesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPublicIPPrefixesServerTransport(srv *PublicIPPrefixesServer) *PublicIPPrefixesServerTransport {
	return &PublicIPPrefixesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.PublicIPPrefixesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.PublicIPPrefixesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListResponse]](),
		newListAllPager:     newTracker[azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListAllResponse]](),
	}
}

// PublicIPPrefixesServerTransport connects instances of armnetwork.PublicIPPrefixesClient to instances of PublicIPPrefixesServer.
// Don't use this type directly, use NewPublicIPPrefixesServerTransport instead.
type PublicIPPrefixesServerTransport struct {
	srv                 *PublicIPPrefixesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.PublicIPPrefixesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.PublicIPPrefixesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListResponse]]
	newListAllPager     *tracker[azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListAllResponse]]
}

// Do implements the policy.Transporter interface for PublicIPPrefixesServerTransport.
func (p *PublicIPPrefixesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PublicIPPrefixesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "PublicIPPrefixesClient.BeginCreateOrUpdate":
			res.resp, res.err = p.dispatchBeginCreateOrUpdate(req)
		case "PublicIPPrefixesClient.BeginDelete":
			res.resp, res.err = p.dispatchBeginDelete(req)
		case "PublicIPPrefixesClient.Get":
			res.resp, res.err = p.dispatchGet(req)
		case "PublicIPPrefixesClient.NewListPager":
			res.resp, res.err = p.dispatchNewListPager(req)
		case "PublicIPPrefixesClient.NewListAllPager":
			res.resp, res.err = p.dispatchNewListAllPager(req)
		case "PublicIPPrefixesClient.UpdateTags":
			res.resp, res.err = p.dispatchUpdateTags(req)
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

func (p *PublicIPPrefixesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/publicIPPrefixes/(?P<publicIpPrefixName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.PublicIPPrefix](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publicIPPrefixNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpPrefixName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, publicIPPrefixNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/publicIPPrefixes/(?P<publicIpPrefixName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publicIPPrefixNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpPrefixName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, publicIPPrefixNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/publicIPPrefixes/(?P<publicIpPrefixName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publicIPPrefixNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpPrefixName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.PublicIPPrefixesClientGetOptions
	if expandParam != nil {
		options = &armnetwork.PublicIPPrefixesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, publicIPPrefixNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicIPPrefix, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/publicIPPrefixes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.PublicIPPrefixesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := p.newListAllPager.get(req)
	if newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/publicIPPrefixes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListAllPager(nil)
		newListAllPager = &resp
		p.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armnetwork.PublicIPPrefixesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		p.newListAllPager.remove(req)
	}
	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if p.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/publicIPPrefixes/(?P<publicIpPrefixName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publicIPPrefixNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpPrefixName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.UpdateTags(req.Context(), resourceGroupNameParam, publicIPPrefixNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicIPPrefix, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
