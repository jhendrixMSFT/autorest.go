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
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// VirtualMachineImagesServer is a fake server for instances of the armcompute.VirtualMachineImagesClient type.
type VirtualMachineImagesServer struct {
	// Get is the fake for method VirtualMachineImagesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, publisherName string, offer string, skus string, version string, options *armcompute.VirtualMachineImagesClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method VirtualMachineImagesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, location string, publisherName string, offer string, skus string, options *armcompute.VirtualMachineImagesClientListOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesClientListResponse], errResp azfake.ErrorResponder)

	// ListOffers is the fake for method VirtualMachineImagesClient.ListOffers
	// HTTP status codes to indicate success: http.StatusOK
	ListOffers func(ctx context.Context, location string, publisherName string, options *armcompute.VirtualMachineImagesClientListOffersOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesClientListOffersResponse], errResp azfake.ErrorResponder)

	// ListPublishers is the fake for method VirtualMachineImagesClient.ListPublishers
	// HTTP status codes to indicate success: http.StatusOK
	ListPublishers func(ctx context.Context, location string, options *armcompute.VirtualMachineImagesClientListPublishersOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesClientListPublishersResponse], errResp azfake.ErrorResponder)

	// ListSKUs is the fake for method VirtualMachineImagesClient.ListSKUs
	// HTTP status codes to indicate success: http.StatusOK
	ListSKUs func(ctx context.Context, location string, publisherName string, offer string, options *armcompute.VirtualMachineImagesClientListSKUsOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesClientListSKUsResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineImagesServerTransport creates a new instance of VirtualMachineImagesServerTransport with the provided implementation.
// The returned VirtualMachineImagesServerTransport instance is connected to an instance of armcompute.VirtualMachineImagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineImagesServerTransport(srv *VirtualMachineImagesServer) *VirtualMachineImagesServerTransport {
	return &VirtualMachineImagesServerTransport{srv: srv}
}

// VirtualMachineImagesServerTransport connects instances of armcompute.VirtualMachineImagesClient to instances of VirtualMachineImagesServer.
// Don't use this type directly, use NewVirtualMachineImagesServerTransport instead.
type VirtualMachineImagesServerTransport struct {
	srv *VirtualMachineImagesServer
}

// Do implements the policy.Transporter interface for VirtualMachineImagesServerTransport.
func (v *VirtualMachineImagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VirtualMachineImagesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "VirtualMachineImagesClient.Get":
			res.resp, res.err = v.dispatchGet(req)
		case "VirtualMachineImagesClient.List":
			res.resp, res.err = v.dispatchList(req)
		case "VirtualMachineImagesClient.ListOffers":
			res.resp, res.err = v.dispatchListOffers(req)
		case "VirtualMachineImagesClient.ListPublishers":
			res.resp, res.err = v.dispatchListPublishers(req)
		case "VirtualMachineImagesClient.ListSKUs":
			res.resp, res.err = v.dispatchListSKUs(req)
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

func (v *VirtualMachineImagesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmimage/offers/(?P<offer>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus/(?P<skus>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	offerParam, err := url.PathUnescape(matches[regex.SubexpIndex("offer")])
	if err != nil {
		return nil, err
	}
	skusParam, err := url.PathUnescape(matches[regex.SubexpIndex("skus")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), locationParam, publisherNameParam, offerParam, skusParam, versionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImagesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if v.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmimage/offers/(?P<offer>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus/(?P<skus>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	offerParam, err := url.PathUnescape(matches[regex.SubexpIndex("offer")])
	if err != nil {
		return nil, err
	}
	skusParam, err := url.PathUnescape(matches[regex.SubexpIndex("skus")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
	if err != nil {
		return nil, err
	}
	topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderbyParam := getOptional(orderbyUnescaped)
	var options *armcompute.VirtualMachineImagesClientListOptions
	if expandParam != nil || topParam != nil || orderbyParam != nil {
		options = &armcompute.VirtualMachineImagesClientListOptions{
			Expand:  expandParam,
			Top:     topParam,
			Orderby: orderbyParam,
		}
	}
	respr, errRespr := v.srv.List(req.Context(), locationParam, publisherNameParam, offerParam, skusParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImageResourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImagesServerTransport) dispatchListOffers(req *http.Request) (*http.Response, error) {
	if v.srv.ListOffers == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListOffers not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmimage/offers`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.ListOffers(req.Context(), locationParam, publisherNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImageResourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImagesServerTransport) dispatchListPublishers(req *http.Request) (*http.Response, error) {
	if v.srv.ListPublishers == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListPublishers not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.ListPublishers(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImageResourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImagesServerTransport) dispatchListSKUs(req *http.Request) (*http.Response, error) {
	if v.srv.ListSKUs == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSKUs not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmimage/offers/(?P<offer>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	offerParam, err := url.PathUnescape(matches[regex.SubexpIndex("offer")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.ListSKUs(req.Context(), locationParam, publisherNameParam, offerParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImageResourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
