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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
)

// DiskEncryptionSetsServer is a fake server for instances of the armcompute.DiskEncryptionSetsClient type.
type DiskEncryptionSetsServer struct {
	// BeginCreateOrUpdate is the fake for method DiskEncryptionSetsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet armcompute.DiskEncryptionSet, options *armcompute.DiskEncryptionSetsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.DiskEncryptionSetsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DiskEncryptionSetsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *armcompute.DiskEncryptionSetsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.DiskEncryptionSetsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DiskEncryptionSetsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *armcompute.DiskEncryptionSetsClientGetOptions) (resp azfake.Responder[armcompute.DiskEncryptionSetsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DiskEncryptionSetsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcompute.DiskEncryptionSetsClientListOptions) (resp azfake.PagerResponder[armcompute.DiskEncryptionSetsClientListResponse])

	// NewListAssociatedResourcesPager is the fake for method DiskEncryptionSetsClient.NewListAssociatedResourcesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAssociatedResourcesPager func(resourceGroupName string, diskEncryptionSetName string, options *armcompute.DiskEncryptionSetsClientListAssociatedResourcesOptions) (resp azfake.PagerResponder[armcompute.DiskEncryptionSetsClientListAssociatedResourcesResponse])

	// NewListByResourceGroupPager is the fake for method DiskEncryptionSetsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcompute.DiskEncryptionSetsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.DiskEncryptionSetsClientListByResourceGroupResponse])

	// BeginUpdate is the fake for method DiskEncryptionSetsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet armcompute.DiskEncryptionSetUpdate, options *armcompute.DiskEncryptionSetsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.DiskEncryptionSetsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDiskEncryptionSetsServerTransport creates a new instance of DiskEncryptionSetsServerTransport with the provided implementation.
// The returned DiskEncryptionSetsServerTransport instance is connected to an instance of armcompute.DiskEncryptionSetsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDiskEncryptionSetsServerTransport(srv *DiskEncryptionSetsServer) *DiskEncryptionSetsServerTransport {
	return &DiskEncryptionSetsServerTransport{srv: srv}
}

// DiskEncryptionSetsServerTransport connects instances of armcompute.DiskEncryptionSetsClient to instances of DiskEncryptionSetsServer.
// Don't use this type directly, use NewDiskEncryptionSetsServerTransport instead.
type DiskEncryptionSetsServerTransport struct {
	srv                             *DiskEncryptionSetsServer
	beginCreateOrUpdate             *azfake.PollerResponder[armcompute.DiskEncryptionSetsClientCreateOrUpdateResponse]
	beginDelete                     *azfake.PollerResponder[armcompute.DiskEncryptionSetsClientDeleteResponse]
	newListPager                    *azfake.PagerResponder[armcompute.DiskEncryptionSetsClientListResponse]
	newListAssociatedResourcesPager *azfake.PagerResponder[armcompute.DiskEncryptionSetsClientListAssociatedResourcesResponse]
	newListByResourceGroupPager     *azfake.PagerResponder[armcompute.DiskEncryptionSetsClientListByResourceGroupResponse]
	beginUpdate                     *azfake.PollerResponder[armcompute.DiskEncryptionSetsClientUpdateResponse]
}

// Do implements the policy.Transporter interface for DiskEncryptionSetsServerTransport.
func (d *DiskEncryptionSetsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DiskEncryptionSetsClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DiskEncryptionSetsClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DiskEncryptionSetsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DiskEncryptionSetsClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DiskEncryptionSetsClient.NewListAssociatedResourcesPager":
		resp, err = d.dispatchNewListAssociatedResourcesPager(req)
	case "DiskEncryptionSetsClient.NewListByResourceGroupPager":
		resp, err = d.dispatchNewListByResourceGroupPager(req)
	case "DiskEncryptionSetsClient.BeginUpdate":
		resp, err = d.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DiskEncryptionSetsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if d.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/diskEncryptionSets/(?P<diskEncryptionSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.DiskEncryptionSet](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskEncryptionSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskEncryptionSetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, diskEncryptionSetNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(d.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginCreateOrUpdate) {
		d.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (d *DiskEncryptionSetsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if d.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/diskEncryptionSets/(?P<diskEncryptionSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskEncryptionSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskEncryptionSetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, diskEncryptionSetNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(d.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginDelete) {
		d.beginDelete = nil
	}

	return resp, nil
}

func (d *DiskEncryptionSetsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/diskEncryptionSets/(?P<diskEncryptionSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	diskEncryptionSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskEncryptionSetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameUnescaped, diskEncryptionSetNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiskEncryptionSet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiskEncryptionSetsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if d.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/diskEncryptionSets`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListPager(nil)
		d.newListPager = &resp
		server.PagerResponderInjectNextLinks(d.newListPager, req, func(page *armcompute.DiskEncryptionSetsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListPager) {
		d.newListPager = nil
	}
	return resp, nil
}

func (d *DiskEncryptionSetsServerTransport) dispatchNewListAssociatedResourcesPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListAssociatedResourcesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAssociatedResourcesPager not implemented")}
	}
	if d.newListAssociatedResourcesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/diskEncryptionSets/(?P<diskEncryptionSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/associatedResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskEncryptionSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskEncryptionSetName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListAssociatedResourcesPager(resourceGroupNameUnescaped, diskEncryptionSetNameUnescaped, nil)
		d.newListAssociatedResourcesPager = &resp
		server.PagerResponderInjectNextLinks(d.newListAssociatedResourcesPager, req, func(page *armcompute.DiskEncryptionSetsClientListAssociatedResourcesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListAssociatedResourcesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListAssociatedResourcesPager) {
		d.newListAssociatedResourcesPager = nil
	}
	return resp, nil
}

func (d *DiskEncryptionSetsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	if d.newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/diskEncryptionSets`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		d.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(d.newListByResourceGroupPager, req, func(page *armcompute.DiskEncryptionSetsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListByResourceGroupPager) {
		d.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (d *DiskEncryptionSetsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	if d.beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/diskEncryptionSets/(?P<diskEncryptionSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.DiskEncryptionSetUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskEncryptionSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskEncryptionSetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdate(req.Context(), resourceGroupNameUnescaped, diskEncryptionSetNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(d.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginUpdate) {
		d.beginUpdate = nil
	}

	return resp, nil
}
