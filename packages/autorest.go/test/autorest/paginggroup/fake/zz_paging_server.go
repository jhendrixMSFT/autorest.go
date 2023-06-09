//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/paginggroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PagingServer is a fake server for instances of the paginggroup.PagingClient type.
type PagingServer struct {
	// NewDuplicateParamsPager is the fake for method PagingClient.NewDuplicateParamsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewDuplicateParamsPager func(options *paginggroup.PagingClientDuplicateParamsOptions) (resp azfake.PagerResponder[paginggroup.PagingClientDuplicateParamsResponse])

	// NewFirstResponseEmptyPager is the fake for method PagingClient.NewFirstResponseEmptyPager
	// HTTP status codes to indicate success: http.StatusOK
	NewFirstResponseEmptyPager func(options *paginggroup.PagingClientFirstResponseEmptyOptions) (resp azfake.PagerResponder[paginggroup.PagingClientFirstResponseEmptyResponse])

	// NewGetMultiplePagesPager is the fake for method PagingClient.NewGetMultiplePagesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetMultiplePagesPager func(options *paginggroup.PagingClientGetMultiplePagesOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesResponse])

	// NewGetMultiplePagesFailurePager is the fake for method PagingClient.NewGetMultiplePagesFailurePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetMultiplePagesFailurePager func(options *paginggroup.PagingClientGetMultiplePagesFailureOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesFailureResponse])

	// NewGetMultiplePagesFailureURIPager is the fake for method PagingClient.NewGetMultiplePagesFailureURIPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetMultiplePagesFailureURIPager func(options *paginggroup.PagingClientGetMultiplePagesFailureURIOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesFailureURIResponse])

	// NewGetMultiplePagesFragmentNextLinkPager is the fake for method PagingClient.NewGetMultiplePagesFragmentNextLinkPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetMultiplePagesFragmentNextLinkPager func(apiVersion string, tenant string, options *paginggroup.PagingClientGetMultiplePagesFragmentNextLinkOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesFragmentNextLinkResponse])

	// NewGetMultiplePagesFragmentWithGroupingNextLinkPager is the fake for method PagingClient.NewGetMultiplePagesFragmentWithGroupingNextLinkPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetMultiplePagesFragmentWithGroupingNextLinkPager func(customParameterGroup paginggroup.CustomParameterGroup, options *paginggroup.PagingClientGetMultiplePagesFragmentWithGroupingNextLinkOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesFragmentWithGroupingNextLinkResponse])

	// BeginGetMultiplePagesLRO is the fake for method PagingClient.BeginGetMultiplePagesLRO
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginGetMultiplePagesLRO func(ctx context.Context, options *paginggroup.PagingClientBeginGetMultiplePagesLROOptions) (resp azfake.PollerResponder[azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesLROResponse]], errResp azfake.ErrorResponder)

	// NewGetMultiplePagesRetryFirstPager is the fake for method PagingClient.NewGetMultiplePagesRetryFirstPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetMultiplePagesRetryFirstPager func(options *paginggroup.PagingClientGetMultiplePagesRetryFirstOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesRetryFirstResponse])

	// NewGetMultiplePagesRetrySecondPager is the fake for method PagingClient.NewGetMultiplePagesRetrySecondPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetMultiplePagesRetrySecondPager func(options *paginggroup.PagingClientGetMultiplePagesRetrySecondOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesRetrySecondResponse])

	// NewGetMultiplePagesWithOffsetPager is the fake for method PagingClient.NewGetMultiplePagesWithOffsetPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetMultiplePagesWithOffsetPager func(options paginggroup.PagingClientGetMultiplePagesWithOffsetOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesWithOffsetResponse])

	// NewGetNoItemNamePagesPager is the fake for method PagingClient.NewGetNoItemNamePagesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetNoItemNamePagesPager func(options *paginggroup.PagingClientGetNoItemNamePagesOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetNoItemNamePagesResponse])

	// NewGetNullNextLinkNamePagesPager is the fake for method PagingClient.NewGetNullNextLinkNamePagesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetNullNextLinkNamePagesPager func(options *paginggroup.PagingClientGetNullNextLinkNamePagesOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetNullNextLinkNamePagesResponse])

	// NewGetODataMultiplePagesPager is the fake for method PagingClient.NewGetODataMultiplePagesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetODataMultiplePagesPager func(options *paginggroup.PagingClientGetODataMultiplePagesOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetODataMultiplePagesResponse])

	// NewGetPagingModelWithItemNameWithXMSClientNamePager is the fake for method PagingClient.NewGetPagingModelWithItemNameWithXMSClientNamePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetPagingModelWithItemNameWithXMSClientNamePager func(options *paginggroup.PagingClientGetPagingModelWithItemNameWithXMSClientNameOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetPagingModelWithItemNameWithXMSClientNameResponse])

	// NewGetSinglePagesPager is the fake for method PagingClient.NewGetSinglePagesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetSinglePagesPager func(options *paginggroup.PagingClientGetSinglePagesOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetSinglePagesResponse])

	// NewGetSinglePagesFailurePager is the fake for method PagingClient.NewGetSinglePagesFailurePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetSinglePagesFailurePager func(options *paginggroup.PagingClientGetSinglePagesFailureOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetSinglePagesFailureResponse])

	// NewGetWithQueryParamsPager is the fake for method PagingClient.NewGetWithQueryParamsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetWithQueryParamsPager func(requiredQueryParameter int32, options *paginggroup.PagingClientGetWithQueryParamsOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetWithQueryParamsResponse])
}

// NewPagingServerTransport creates a new instance of PagingServerTransport with the provided implementation.
// The returned PagingServerTransport instance is connected to an instance of paginggroup.PagingClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPagingServerTransport(srv *PagingServer) *PagingServerTransport {
	return &PagingServerTransport{srv: srv}
}

// PagingServerTransport connects instances of paginggroup.PagingClient to instances of PagingServer.
// Don't use this type directly, use NewPagingServerTransport instead.
type PagingServerTransport struct {
	srv                                                  *PagingServer
	newDuplicateParamsPager                              *azfake.PagerResponder[paginggroup.PagingClientDuplicateParamsResponse]
	newFirstResponseEmptyPager                           *azfake.PagerResponder[paginggroup.PagingClientFirstResponseEmptyResponse]
	newGetMultiplePagesPager                             *azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesResponse]
	newGetMultiplePagesFailurePager                      *azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesFailureResponse]
	newGetMultiplePagesFailureURIPager                   *azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesFailureURIResponse]
	newGetMultiplePagesFragmentNextLinkPager             *azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesFragmentNextLinkResponse]
	newGetMultiplePagesFragmentWithGroupingNextLinkPager *azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesFragmentWithGroupingNextLinkResponse]
	beginGetMultiplePagesLRO                             *azfake.PollerResponder[azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesLROResponse]]
	newGetMultiplePagesRetryFirstPager                   *azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesRetryFirstResponse]
	newGetMultiplePagesRetrySecondPager                  *azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesRetrySecondResponse]
	newGetMultiplePagesWithOffsetPager                   *azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesWithOffsetResponse]
	newGetNoItemNamePagesPager                           *azfake.PagerResponder[paginggroup.PagingClientGetNoItemNamePagesResponse]
	newGetNullNextLinkNamePagesPager                     *azfake.PagerResponder[paginggroup.PagingClientGetNullNextLinkNamePagesResponse]
	newGetODataMultiplePagesPager                        *azfake.PagerResponder[paginggroup.PagingClientGetODataMultiplePagesResponse]
	newGetPagingModelWithItemNameWithXMSClientNamePager  *azfake.PagerResponder[paginggroup.PagingClientGetPagingModelWithItemNameWithXMSClientNameResponse]
	newGetSinglePagesPager                               *azfake.PagerResponder[paginggroup.PagingClientGetSinglePagesResponse]
	newGetSinglePagesFailurePager                        *azfake.PagerResponder[paginggroup.PagingClientGetSinglePagesFailureResponse]
	newGetWithQueryParamsPager                           *azfake.PagerResponder[paginggroup.PagingClientGetWithQueryParamsResponse]
}

// Do implements the policy.Transporter interface for PagingServerTransport.
func (p *PagingServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PagingClient.NewDuplicateParamsPager":
		resp, err = p.dispatchNewDuplicateParamsPager(req)
	case "PagingClient.NewFirstResponseEmptyPager":
		resp, err = p.dispatchNewFirstResponseEmptyPager(req)
	case "PagingClient.NewGetMultiplePagesPager":
		resp, err = p.dispatchNewGetMultiplePagesPager(req)
	case "PagingClient.NewGetMultiplePagesFailurePager":
		resp, err = p.dispatchNewGetMultiplePagesFailurePager(req)
	case "PagingClient.NewGetMultiplePagesFailureURIPager":
		resp, err = p.dispatchNewGetMultiplePagesFailureURIPager(req)
	case "PagingClient.NewGetMultiplePagesFragmentNextLinkPager":
		resp, err = p.dispatchNewGetMultiplePagesFragmentNextLinkPager(req)
	case "PagingClient.NewGetMultiplePagesFragmentWithGroupingNextLinkPager":
		resp, err = p.dispatchNewGetMultiplePagesFragmentWithGroupingNextLinkPager(req)
	case "PagingClient.BeginGetMultiplePagesLRO":
		resp, err = p.dispatchBeginGetMultiplePagesLRO(req)
	case "PagingClient.NewGetMultiplePagesRetryFirstPager":
		resp, err = p.dispatchNewGetMultiplePagesRetryFirstPager(req)
	case "PagingClient.NewGetMultiplePagesRetrySecondPager":
		resp, err = p.dispatchNewGetMultiplePagesRetrySecondPager(req)
	case "PagingClient.NewGetMultiplePagesWithOffsetPager":
		resp, err = p.dispatchNewGetMultiplePagesWithOffsetPager(req)
	case "PagingClient.NewGetNoItemNamePagesPager":
		resp, err = p.dispatchNewGetNoItemNamePagesPager(req)
	case "PagingClient.NewGetNullNextLinkNamePagesPager":
		resp, err = p.dispatchNewGetNullNextLinkNamePagesPager(req)
	case "PagingClient.NewGetODataMultiplePagesPager":
		resp, err = p.dispatchNewGetODataMultiplePagesPager(req)
	case "PagingClient.NewGetPagingModelWithItemNameWithXMSClientNamePager":
		resp, err = p.dispatchNewGetPagingModelWithItemNameWithXMSClientNamePager(req)
	case "PagingClient.NewGetSinglePagesPager":
		resp, err = p.dispatchNewGetSinglePagesPager(req)
	case "PagingClient.NewGetSinglePagesFailurePager":
		resp, err = p.dispatchNewGetSinglePagesFailurePager(req)
	case "PagingClient.NewGetWithQueryParamsPager":
		resp, err = p.dispatchNewGetWithQueryParamsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PagingServerTransport) dispatchNewDuplicateParamsPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewDuplicateParamsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewDuplicateParamsPager not implemented")}
	}
	if p.newDuplicateParamsPager == nil {
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *paginggroup.PagingClientDuplicateParamsOptions
		if filterParam != nil {
			options = &paginggroup.PagingClientDuplicateParamsOptions{
				Filter: filterParam,
			}
		}
		resp := p.srv.NewDuplicateParamsPager(options)
		p.newDuplicateParamsPager = &resp
		server.PagerResponderInjectNextLinks(p.newDuplicateParamsPager, req, func(page *paginggroup.PagingClientDuplicateParamsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newDuplicateParamsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newDuplicateParamsPager) {
		p.newDuplicateParamsPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewFirstResponseEmptyPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewFirstResponseEmptyPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewFirstResponseEmptyPager not implemented")}
	}
	if p.newFirstResponseEmptyPager == nil {
		resp := p.srv.NewFirstResponseEmptyPager(nil)
		p.newFirstResponseEmptyPager = &resp
		server.PagerResponderInjectNextLinks(p.newFirstResponseEmptyPager, req, func(page *paginggroup.PagingClientFirstResponseEmptyResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newFirstResponseEmptyPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newFirstResponseEmptyPager) {
		p.newFirstResponseEmptyPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetMultiplePagesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetMultiplePagesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetMultiplePagesPager not implemented")}
	}
	if p.newGetMultiplePagesPager == nil {
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "client-request-id"))
		maxresultsParam, err := parseOptional(getHeaderValue(req.Header, "maxresults"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		timeoutParam, err := parseOptional(getHeaderValue(req.Header, "timeout"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *paginggroup.PagingClientGetMultiplePagesOptions
		if clientRequestIDParam != nil || maxresultsParam != nil || timeoutParam != nil {
			options = &paginggroup.PagingClientGetMultiplePagesOptions{
				ClientRequestID: clientRequestIDParam,
				Maxresults:      maxresultsParam,
				Timeout:         timeoutParam,
			}
		}
		resp := p.srv.NewGetMultiplePagesPager(options)
		p.newGetMultiplePagesPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetMultiplePagesPager, req, func(page *paginggroup.PagingClientGetMultiplePagesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetMultiplePagesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetMultiplePagesPager) {
		p.newGetMultiplePagesPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetMultiplePagesFailurePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetMultiplePagesFailurePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetMultiplePagesFailurePager not implemented")}
	}
	if p.newGetMultiplePagesFailurePager == nil {
		resp := p.srv.NewGetMultiplePagesFailurePager(nil)
		p.newGetMultiplePagesFailurePager = &resp
		server.PagerResponderInjectNextLinks(p.newGetMultiplePagesFailurePager, req, func(page *paginggroup.PagingClientGetMultiplePagesFailureResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetMultiplePagesFailurePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetMultiplePagesFailurePager) {
		p.newGetMultiplePagesFailurePager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetMultiplePagesFailureURIPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetMultiplePagesFailureURIPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetMultiplePagesFailureURIPager not implemented")}
	}
	if p.newGetMultiplePagesFailureURIPager == nil {
		resp := p.srv.NewGetMultiplePagesFailureURIPager(nil)
		p.newGetMultiplePagesFailureURIPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetMultiplePagesFailureURIPager, req, func(page *paginggroup.PagingClientGetMultiplePagesFailureURIResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetMultiplePagesFailureURIPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetMultiplePagesFailureURIPager) {
		p.newGetMultiplePagesFailureURIPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetMultiplePagesFragmentNextLinkPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetMultiplePagesFragmentNextLinkPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetMultiplePagesFragmentNextLinkPager not implemented")}
	}
	if p.newGetMultiplePagesFragmentNextLinkPager == nil {
		const regexStr = `/paging/multiple/fragment/(?P<tenant>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		apiVersionUnescaped, err := url.QueryUnescape(qp.Get("api_version"))
		if err != nil {
			return nil, err
		}
		tenantUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tenant")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewGetMultiplePagesFragmentNextLinkPager(apiVersionUnescaped, tenantUnescaped, nil)
		p.newGetMultiplePagesFragmentNextLinkPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetMultiplePagesFragmentNextLinkPager, req, func(page *paginggroup.PagingClientGetMultiplePagesFragmentNextLinkResponse, createLink func() string) {
			page.ODataNextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetMultiplePagesFragmentNextLinkPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetMultiplePagesFragmentNextLinkPager) {
		p.newGetMultiplePagesFragmentNextLinkPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetMultiplePagesFragmentWithGroupingNextLinkPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetMultiplePagesFragmentWithGroupingNextLinkPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetMultiplePagesFragmentWithGroupingNextLinkPager not implemented")}
	}
	if p.newGetMultiplePagesFragmentWithGroupingNextLinkPager == nil {
		const regexStr = `/paging/multiple/fragmentwithgrouping/(?P<tenant>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		aPIVersionUnescaped, err := url.QueryUnescape(qp.Get("api_version"))
		if err != nil {
			return nil, err
		}
		aPIVersionParam := aPIVersionUnescaped
		tenantUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tenant")])
		if err != nil {
			return nil, err
		}
		tenantParam := tenantUnescaped
		customParameterGroup := paginggroup.CustomParameterGroup{
			APIVersion: aPIVersionParam,
			Tenant:     tenantParam,
		}
		resp := p.srv.NewGetMultiplePagesFragmentWithGroupingNextLinkPager(customParameterGroup, nil)
		p.newGetMultiplePagesFragmentWithGroupingNextLinkPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetMultiplePagesFragmentWithGroupingNextLinkPager, req, func(page *paginggroup.PagingClientGetMultiplePagesFragmentWithGroupingNextLinkResponse, createLink func() string) {
			page.ODataNextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetMultiplePagesFragmentWithGroupingNextLinkPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetMultiplePagesFragmentWithGroupingNextLinkPager) {
		p.newGetMultiplePagesFragmentWithGroupingNextLinkPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchBeginGetMultiplePagesLRO(req *http.Request) (*http.Response, error) {
	if p.srv.BeginGetMultiplePagesLRO == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetMultiplePagesLRO not implemented")}
	}
	if p.beginGetMultiplePagesLRO == nil {
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "client-request-id"))
		maxresultsParam, err := parseOptional(getHeaderValue(req.Header, "maxresults"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		timeoutParam, err := parseOptional(getHeaderValue(req.Header, "timeout"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *paginggroup.PagingClientBeginGetMultiplePagesLROOptions
		if clientRequestIDParam != nil || maxresultsParam != nil || timeoutParam != nil {
			options = &paginggroup.PagingClientBeginGetMultiplePagesLROOptions{
				ClientRequestID: clientRequestIDParam,
				Maxresults:      maxresultsParam,
				Timeout:         timeoutParam,
			}
		}
		respr, errRespr := p.srv.BeginGetMultiplePagesLRO(req.Context(), options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginGetMultiplePagesLRO = &respr
	}

	resp, err := server.PollerResponderNext(p.beginGetMultiplePagesLRO, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginGetMultiplePagesLRO) {
		p.beginGetMultiplePagesLRO = nil
	}

	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetMultiplePagesRetryFirstPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetMultiplePagesRetryFirstPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetMultiplePagesRetryFirstPager not implemented")}
	}
	if p.newGetMultiplePagesRetryFirstPager == nil {
		resp := p.srv.NewGetMultiplePagesRetryFirstPager(nil)
		p.newGetMultiplePagesRetryFirstPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetMultiplePagesRetryFirstPager, req, func(page *paginggroup.PagingClientGetMultiplePagesRetryFirstResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetMultiplePagesRetryFirstPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetMultiplePagesRetryFirstPager) {
		p.newGetMultiplePagesRetryFirstPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetMultiplePagesRetrySecondPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetMultiplePagesRetrySecondPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetMultiplePagesRetrySecondPager not implemented")}
	}
	if p.newGetMultiplePagesRetrySecondPager == nil {
		resp := p.srv.NewGetMultiplePagesRetrySecondPager(nil)
		p.newGetMultiplePagesRetrySecondPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetMultiplePagesRetrySecondPager, req, func(page *paginggroup.PagingClientGetMultiplePagesRetrySecondResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetMultiplePagesRetrySecondPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetMultiplePagesRetrySecondPager) {
		p.newGetMultiplePagesRetrySecondPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetMultiplePagesWithOffsetPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetMultiplePagesWithOffsetPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetMultiplePagesWithOffsetPager not implemented")}
	}
	if p.newGetMultiplePagesWithOffsetPager == nil {
		const regexStr = `/paging/multiple/withpath/(?P<offset>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "client-request-id"))
		maxresultsParam, err := parseOptional(getHeaderValue(req.Header, "maxresults"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		offsetUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("offset")])
		if err != nil {
			return nil, err
		}
		offsetParam, err := parseWithCast(offsetUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		timeoutParam, err := parseOptional(getHeaderValue(req.Header, "timeout"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		options := paginggroup.PagingClientGetMultiplePagesWithOffsetOptions{
			ClientRequestID: clientRequestIDParam,
			Maxresults:      maxresultsParam,
			Offset:          offsetParam,
			Timeout:         timeoutParam,
		}
		resp := p.srv.NewGetMultiplePagesWithOffsetPager(options)
		p.newGetMultiplePagesWithOffsetPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetMultiplePagesWithOffsetPager, req, func(page *paginggroup.PagingClientGetMultiplePagesWithOffsetResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetMultiplePagesWithOffsetPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetMultiplePagesWithOffsetPager) {
		p.newGetMultiplePagesWithOffsetPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetNoItemNamePagesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetNoItemNamePagesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetNoItemNamePagesPager not implemented")}
	}
	if p.newGetNoItemNamePagesPager == nil {
		resp := p.srv.NewGetNoItemNamePagesPager(nil)
		p.newGetNoItemNamePagesPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetNoItemNamePagesPager, req, func(page *paginggroup.PagingClientGetNoItemNamePagesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetNoItemNamePagesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetNoItemNamePagesPager) {
		p.newGetNoItemNamePagesPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetNullNextLinkNamePagesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetNullNextLinkNamePagesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetNullNextLinkNamePagesPager not implemented")}
	}
	if p.newGetNullNextLinkNamePagesPager == nil {
		resp := p.srv.NewGetNullNextLinkNamePagesPager(nil)
		p.newGetNullNextLinkNamePagesPager = &resp
	}
	resp, err := server.PagerResponderNext(p.newGetNullNextLinkNamePagesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetNullNextLinkNamePagesPager) {
		p.newGetNullNextLinkNamePagesPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetODataMultiplePagesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetODataMultiplePagesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetODataMultiplePagesPager not implemented")}
	}
	if p.newGetODataMultiplePagesPager == nil {
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "client-request-id"))
		maxresultsParam, err := parseOptional(getHeaderValue(req.Header, "maxresults"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		timeoutParam, err := parseOptional(getHeaderValue(req.Header, "timeout"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *paginggroup.PagingClientGetODataMultiplePagesOptions
		if clientRequestIDParam != nil || maxresultsParam != nil || timeoutParam != nil {
			options = &paginggroup.PagingClientGetODataMultiplePagesOptions{
				ClientRequestID: clientRequestIDParam,
				Maxresults:      maxresultsParam,
				Timeout:         timeoutParam,
			}
		}
		resp := p.srv.NewGetODataMultiplePagesPager(options)
		p.newGetODataMultiplePagesPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetODataMultiplePagesPager, req, func(page *paginggroup.PagingClientGetODataMultiplePagesResponse, createLink func() string) {
			page.ODataNextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetODataMultiplePagesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetODataMultiplePagesPager) {
		p.newGetODataMultiplePagesPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetPagingModelWithItemNameWithXMSClientNamePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetPagingModelWithItemNameWithXMSClientNamePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetPagingModelWithItemNameWithXMSClientNamePager not implemented")}
	}
	if p.newGetPagingModelWithItemNameWithXMSClientNamePager == nil {
		resp := p.srv.NewGetPagingModelWithItemNameWithXMSClientNamePager(nil)
		p.newGetPagingModelWithItemNameWithXMSClientNamePager = &resp
		server.PagerResponderInjectNextLinks(p.newGetPagingModelWithItemNameWithXMSClientNamePager, req, func(page *paginggroup.PagingClientGetPagingModelWithItemNameWithXMSClientNameResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetPagingModelWithItemNameWithXMSClientNamePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetPagingModelWithItemNameWithXMSClientNamePager) {
		p.newGetPagingModelWithItemNameWithXMSClientNamePager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetSinglePagesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetSinglePagesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetSinglePagesPager not implemented")}
	}
	if p.newGetSinglePagesPager == nil {
		resp := p.srv.NewGetSinglePagesPager(nil)
		p.newGetSinglePagesPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetSinglePagesPager, req, func(page *paginggroup.PagingClientGetSinglePagesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetSinglePagesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetSinglePagesPager) {
		p.newGetSinglePagesPager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetSinglePagesFailurePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetSinglePagesFailurePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetSinglePagesFailurePager not implemented")}
	}
	if p.newGetSinglePagesFailurePager == nil {
		resp := p.srv.NewGetSinglePagesFailurePager(nil)
		p.newGetSinglePagesFailurePager = &resp
		server.PagerResponderInjectNextLinks(p.newGetSinglePagesFailurePager, req, func(page *paginggroup.PagingClientGetSinglePagesFailureResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetSinglePagesFailurePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetSinglePagesFailurePager) {
		p.newGetSinglePagesFailurePager = nil
	}
	return resp, nil
}

func (p *PagingServerTransport) dispatchNewGetWithQueryParamsPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetWithQueryParamsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetWithQueryParamsPager not implemented")}
	}
	if p.newGetWithQueryParamsPager == nil {
		qp := req.URL.Query()
		requiredQueryParameterUnescaped, err := url.QueryUnescape(qp.Get("requiredQueryParameter"))
		if err != nil {
			return nil, err
		}
		requiredQueryParameterParam, err := parseWithCast(requiredQueryParameterUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewGetWithQueryParamsPager(int32(requiredQueryParameterParam), nil)
		p.newGetWithQueryParamsPager = &resp
		server.PagerResponderInjectNextLinks(p.newGetWithQueryParamsPager, req, func(page *paginggroup.PagingClientGetWithQueryParamsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newGetWithQueryParamsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newGetWithQueryParamsPager) {
		p.newGetWithQueryParamsPager = nil
	}
	return resp, nil
}
