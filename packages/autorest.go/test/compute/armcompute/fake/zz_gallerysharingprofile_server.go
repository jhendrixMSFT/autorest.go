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
)

// GallerySharingProfileServer is a fake server for instances of the armcompute.GallerySharingProfileClient type.
type GallerySharingProfileServer struct {
	// BeginUpdate is the fake for method GallerySharingProfileClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate armcompute.SharingUpdate, options *armcompute.GallerySharingProfileClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.GallerySharingProfileClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGallerySharingProfileServerTransport creates a new instance of GallerySharingProfileServerTransport with the provided implementation.
// The returned GallerySharingProfileServerTransport instance is connected to an instance of armcompute.GallerySharingProfileClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGallerySharingProfileServerTransport(srv *GallerySharingProfileServer) *GallerySharingProfileServerTransport {
	return &GallerySharingProfileServerTransport{
		srv:         srv,
		beginUpdate: newTracker[azfake.PollerResponder[armcompute.GallerySharingProfileClientUpdateResponse]](),
	}
}

// GallerySharingProfileServerTransport connects instances of armcompute.GallerySharingProfileClient to instances of GallerySharingProfileServer.
// Don't use this type directly, use NewGallerySharingProfileServerTransport instead.
type GallerySharingProfileServerTransport struct {
	srv         *GallerySharingProfileServer
	beginUpdate *tracker[azfake.PollerResponder[armcompute.GallerySharingProfileClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for GallerySharingProfileServerTransport.
func (g *GallerySharingProfileServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GallerySharingProfileServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "GallerySharingProfileClient.BeginUpdate":
			res.resp, res.err = g.dispatchBeginUpdate(req)
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

func (g *GallerySharingProfileServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := g.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/share`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.SharingUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginUpdate(req.Context(), resourceGroupNameParam, galleryNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		g.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		g.beginUpdate.remove(req)
	}

	return resp, nil
}
