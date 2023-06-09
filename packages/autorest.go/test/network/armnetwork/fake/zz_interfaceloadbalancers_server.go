//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
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

// InterfaceLoadBalancersServer is a fake server for instances of the armnetwork.InterfaceLoadBalancersClient type.
type InterfaceLoadBalancersServer struct {
	// NewListPager is the fake for method InterfaceLoadBalancersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkInterfaceName string, options *armnetwork.InterfaceLoadBalancersClientListOptions) (resp azfake.PagerResponder[armnetwork.InterfaceLoadBalancersClientListResponse])
}

// NewInterfaceLoadBalancersServerTransport creates a new instance of InterfaceLoadBalancersServerTransport with the provided implementation.
// The returned InterfaceLoadBalancersServerTransport instance is connected to an instance of armnetwork.InterfaceLoadBalancersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInterfaceLoadBalancersServerTransport(srv *InterfaceLoadBalancersServer) *InterfaceLoadBalancersServerTransport {
	return &InterfaceLoadBalancersServerTransport{srv: srv}
}

// InterfaceLoadBalancersServerTransport connects instances of armnetwork.InterfaceLoadBalancersClient to instances of InterfaceLoadBalancersServer.
// Don't use this type directly, use NewInterfaceLoadBalancersServerTransport instead.
type InterfaceLoadBalancersServerTransport struct {
	srv          *InterfaceLoadBalancersServer
	newListPager *azfake.PagerResponder[armnetwork.InterfaceLoadBalancersClientListResponse]
}

// Do implements the policy.Transporter interface for InterfaceLoadBalancersServerTransport.
func (i *InterfaceLoadBalancersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "InterfaceLoadBalancersClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *InterfaceLoadBalancersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if i.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/loadBalancers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListPager(resourceGroupNameUnescaped, networkInterfaceNameUnescaped, nil)
		i.newListPager = &resp
		server.PagerResponderInjectNextLinks(i.newListPager, req, func(page *armnetwork.InterfaceLoadBalancersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(i.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(i.newListPager) {
		i.newListPager = nil
	}
	return resp, nil
}
