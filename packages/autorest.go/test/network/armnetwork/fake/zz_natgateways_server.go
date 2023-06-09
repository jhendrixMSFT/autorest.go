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

// NatGatewaysServer is a fake server for instances of the armnetwork.NatGatewaysClient type.
type NatGatewaysServer struct {
	// BeginCreateOrUpdate is the fake for method NatGatewaysClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, natGatewayName string, parameters armnetwork.NatGateway, options *armnetwork.NatGatewaysClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.NatGatewaysClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NatGatewaysClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, natGatewayName string, options *armnetwork.NatGatewaysClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.NatGatewaysClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NatGatewaysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, natGatewayName string, options *armnetwork.NatGatewaysClientGetOptions) (resp azfake.Responder[armnetwork.NatGatewaysClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method NatGatewaysClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.NatGatewaysClientListOptions) (resp azfake.PagerResponder[armnetwork.NatGatewaysClientListResponse])

	// NewListAllPager is the fake for method NatGatewaysClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.NatGatewaysClientListAllOptions) (resp azfake.PagerResponder[armnetwork.NatGatewaysClientListAllResponse])

	// UpdateTags is the fake for method NatGatewaysClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, natGatewayName string, parameters armnetwork.TagsObject, options *armnetwork.NatGatewaysClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.NatGatewaysClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewNatGatewaysServerTransport creates a new instance of NatGatewaysServerTransport with the provided implementation.
// The returned NatGatewaysServerTransport instance is connected to an instance of armnetwork.NatGatewaysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNatGatewaysServerTransport(srv *NatGatewaysServer) *NatGatewaysServerTransport {
	return &NatGatewaysServerTransport{srv: srv}
}

// NatGatewaysServerTransport connects instances of armnetwork.NatGatewaysClient to instances of NatGatewaysServer.
// Don't use this type directly, use NewNatGatewaysServerTransport instead.
type NatGatewaysServerTransport struct {
	srv                 *NatGatewaysServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.NatGatewaysClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.NatGatewaysClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.NatGatewaysClientListResponse]
	newListAllPager     *azfake.PagerResponder[armnetwork.NatGatewaysClientListAllResponse]
}

// Do implements the policy.Transporter interface for NatGatewaysServerTransport.
func (n *NatGatewaysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NatGatewaysClient.BeginCreateOrUpdate":
		resp, err = n.dispatchBeginCreateOrUpdate(req)
	case "NatGatewaysClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NatGatewaysClient.Get":
		resp, err = n.dispatchGet(req)
	case "NatGatewaysClient.NewListPager":
		resp, err = n.dispatchNewListPager(req)
	case "NatGatewaysClient.NewListAllPager":
		resp, err = n.dispatchNewListAllPager(req)
	case "NatGatewaysClient.UpdateTags":
		resp, err = n.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NatGatewaysServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if n.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/natGateways/(?P<natGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.NatGateway](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		natGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("natGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, natGatewayNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		n.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(n.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(n.beginCreateOrUpdate) {
		n.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (n *NatGatewaysServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if n.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/natGateways/(?P<natGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		natGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("natGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, natGatewayNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		n.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(n.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(n.beginDelete) {
		n.beginDelete = nil
	}

	return resp, nil
}

func (n *NatGatewaysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/natGateways/(?P<natGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	natGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("natGatewayName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.NatGatewaysClientGetOptions
	if expandParam != nil {
		options = &armnetwork.NatGatewaysClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameUnescaped, natGatewayNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NatGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NatGatewaysServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if n.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/natGateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListPager(resourceGroupNameUnescaped, nil)
		n.newListPager = &resp
		server.PagerResponderInjectNextLinks(n.newListPager, req, func(page *armnetwork.NatGatewaysClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(n.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(n.newListPager) {
		n.newListPager = nil
	}
	return resp, nil
}

func (n *NatGatewaysServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	if n.newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/natGateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := n.srv.NewListAllPager(nil)
		n.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(n.newListAllPager, req, func(page *armnetwork.NatGatewaysClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(n.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(n.newListAllPager) {
		n.newListAllPager = nil
	}
	return resp, nil
}

func (n *NatGatewaysServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if n.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/natGateways/(?P<natGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	natGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("natGatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, natGatewayNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NatGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
