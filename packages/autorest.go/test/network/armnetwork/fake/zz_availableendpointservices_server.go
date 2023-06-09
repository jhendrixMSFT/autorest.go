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

// AvailableEndpointServicesServer is a fake server for instances of the armnetwork.AvailableEndpointServicesClient type.
type AvailableEndpointServicesServer struct {
	// NewListPager is the fake for method AvailableEndpointServicesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armnetwork.AvailableEndpointServicesClientListOptions) (resp azfake.PagerResponder[armnetwork.AvailableEndpointServicesClientListResponse])
}

// NewAvailableEndpointServicesServerTransport creates a new instance of AvailableEndpointServicesServerTransport with the provided implementation.
// The returned AvailableEndpointServicesServerTransport instance is connected to an instance of armnetwork.AvailableEndpointServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAvailableEndpointServicesServerTransport(srv *AvailableEndpointServicesServer) *AvailableEndpointServicesServerTransport {
	return &AvailableEndpointServicesServerTransport{srv: srv}
}

// AvailableEndpointServicesServerTransport connects instances of armnetwork.AvailableEndpointServicesClient to instances of AvailableEndpointServicesServer.
// Don't use this type directly, use NewAvailableEndpointServicesServerTransport instead.
type AvailableEndpointServicesServerTransport struct {
	srv          *AvailableEndpointServicesServer
	newListPager *azfake.PagerResponder[armnetwork.AvailableEndpointServicesClientListResponse]
}

// Do implements the policy.Transporter interface for AvailableEndpointServicesServerTransport.
func (a *AvailableEndpointServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AvailableEndpointServicesClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AvailableEndpointServicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if a.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkAvailableEndpointServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(locationUnescaped, nil)
		a.newListPager = &resp
		server.PagerResponderInjectNextLinks(a.newListPager, req, func(page *armnetwork.AvailableEndpointServicesClientListResponse, createLink func() string) {
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
