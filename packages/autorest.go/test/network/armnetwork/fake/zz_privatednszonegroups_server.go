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

// PrivateDNSZoneGroupsServer is a fake server for instances of the armnetwork.PrivateDNSZoneGroupsClient type.
type PrivateDNSZoneGroupsServer struct {
	// BeginCreateOrUpdate is the fake for method PrivateDNSZoneGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters armnetwork.PrivateDNSZoneGroup, options *armnetwork.PrivateDNSZoneGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.PrivateDNSZoneGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateDNSZoneGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *armnetwork.PrivateDNSZoneGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.PrivateDNSZoneGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateDNSZoneGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *armnetwork.PrivateDNSZoneGroupsClientGetOptions) (resp azfake.Responder[armnetwork.PrivateDNSZoneGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PrivateDNSZoneGroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(privateEndpointName string, resourceGroupName string, options *armnetwork.PrivateDNSZoneGroupsClientListOptions) (resp azfake.PagerResponder[armnetwork.PrivateDNSZoneGroupsClientListResponse])
}

// NewPrivateDNSZoneGroupsServerTransport creates a new instance of PrivateDNSZoneGroupsServerTransport with the provided implementation.
// The returned PrivateDNSZoneGroupsServerTransport instance is connected to an instance of armnetwork.PrivateDNSZoneGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateDNSZoneGroupsServerTransport(srv *PrivateDNSZoneGroupsServer) *PrivateDNSZoneGroupsServerTransport {
	return &PrivateDNSZoneGroupsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.PrivateDNSZoneGroupsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.PrivateDNSZoneGroupsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.PrivateDNSZoneGroupsClientListResponse]](),
	}
}

// PrivateDNSZoneGroupsServerTransport connects instances of armnetwork.PrivateDNSZoneGroupsClient to instances of PrivateDNSZoneGroupsServer.
// Don't use this type directly, use NewPrivateDNSZoneGroupsServerTransport instead.
type PrivateDNSZoneGroupsServerTransport struct {
	srv                 *PrivateDNSZoneGroupsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.PrivateDNSZoneGroupsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.PrivateDNSZoneGroupsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.PrivateDNSZoneGroupsClientListResponse]]
}

// Do implements the policy.Transporter interface for PrivateDNSZoneGroupsServerTransport.
func (p *PrivateDNSZoneGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateDNSZoneGroupsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "PrivateDNSZoneGroupsClient.BeginCreateOrUpdate":
			res.resp, res.err = p.dispatchBeginCreateOrUpdate(req)
		case "PrivateDNSZoneGroupsClient.BeginDelete":
			res.resp, res.err = p.dispatchBeginDelete(req)
		case "PrivateDNSZoneGroupsClient.Get":
			res.resp, res.err = p.dispatchGet(req)
		case "PrivateDNSZoneGroupsClient.NewListPager":
			res.resp, res.err = p.dispatchNewListPager(req)
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

func (p *PrivateDNSZoneGroupsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateEndpoints/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateDnsZoneGroups/(?P<privateDnsZoneGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.PrivateDNSZoneGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
		if err != nil {
			return nil, err
		}
		privateDNSZoneGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateDnsZoneGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, privateEndpointNameParam, privateDNSZoneGroupNameParam, body, nil)
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

func (p *PrivateDNSZoneGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateEndpoints/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateDnsZoneGroups/(?P<privateDnsZoneGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
		if err != nil {
			return nil, err
		}
		privateDNSZoneGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateDnsZoneGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, privateEndpointNameParam, privateDNSZoneGroupNameParam, nil)
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

func (p *PrivateDNSZoneGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateEndpoints/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateDnsZoneGroups/(?P<privateDnsZoneGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
	if err != nil {
		return nil, err
	}
	privateDNSZoneGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateDnsZoneGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, privateEndpointNameParam, privateDNSZoneGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateDNSZoneGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateDNSZoneGroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateEndpoints/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateDnsZoneGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(privateEndpointNameParam, resourceGroupNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.PrivateDNSZoneGroupsClientListResponse, createLink func() string) {
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
