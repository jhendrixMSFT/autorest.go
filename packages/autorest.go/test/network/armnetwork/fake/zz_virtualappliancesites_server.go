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

// VirtualApplianceSitesServer is a fake server for instances of the armnetwork.VirtualApplianceSitesClient type.
type VirtualApplianceSitesServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualApplianceSitesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters armnetwork.VirtualApplianceSite, options *armnetwork.VirtualApplianceSitesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualApplianceSitesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualApplianceSitesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *armnetwork.VirtualApplianceSitesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualApplianceSitesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualApplianceSitesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *armnetwork.VirtualApplianceSitesClientGetOptions) (resp azfake.Responder[armnetwork.VirtualApplianceSitesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualApplianceSitesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkVirtualApplianceName string, options *armnetwork.VirtualApplianceSitesClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualApplianceSitesClientListResponse])
}

// NewVirtualApplianceSitesServerTransport creates a new instance of VirtualApplianceSitesServerTransport with the provided implementation.
// The returned VirtualApplianceSitesServerTransport instance is connected to an instance of armnetwork.VirtualApplianceSitesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualApplianceSitesServerTransport(srv *VirtualApplianceSitesServer) *VirtualApplianceSitesServerTransport {
	return &VirtualApplianceSitesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.VirtualApplianceSitesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.VirtualApplianceSitesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.VirtualApplianceSitesClientListResponse]](),
	}
}

// VirtualApplianceSitesServerTransport connects instances of armnetwork.VirtualApplianceSitesClient to instances of VirtualApplianceSitesServer.
// Don't use this type directly, use NewVirtualApplianceSitesServerTransport instead.
type VirtualApplianceSitesServerTransport struct {
	srv                 *VirtualApplianceSitesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.VirtualApplianceSitesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.VirtualApplianceSitesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.VirtualApplianceSitesClientListResponse]]
}

// Do implements the policy.Transporter interface for VirtualApplianceSitesServerTransport.
func (v *VirtualApplianceSitesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VirtualApplianceSitesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "VirtualApplianceSitesClient.BeginCreateOrUpdate":
			res.resp, res.err = v.dispatchBeginCreateOrUpdate(req)
		case "VirtualApplianceSitesClient.BeginDelete":
			res.resp, res.err = v.dispatchBeginDelete(req)
		case "VirtualApplianceSitesClient.Get":
			res.resp, res.err = v.dispatchGet(req)
		case "VirtualApplianceSitesClient.NewListPager":
			res.resp, res.err = v.dispatchNewListPager(req)
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

func (v *VirtualApplianceSitesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualApplianceSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualApplianceSite](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualApplianceSitesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualApplianceSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, siteNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualApplianceSitesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualApplianceSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, siteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualApplianceSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualApplianceSitesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualApplianceSites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameParam, networkVirtualApplianceNameParam, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualApplianceSitesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}
