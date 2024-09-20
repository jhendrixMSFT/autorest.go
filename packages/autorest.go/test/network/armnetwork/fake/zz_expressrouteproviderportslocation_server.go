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
	"net/http"
	"net/url"
	"regexp"
)

// ExpressRouteProviderPortsLocationServer is a fake server for instances of the armnetwork.ExpressRouteProviderPortsLocationClient type.
type ExpressRouteProviderPortsLocationServer struct {
	// List is the fake for method ExpressRouteProviderPortsLocationClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, options *armnetwork.ExpressRouteProviderPortsLocationClientListOptions) (resp azfake.Responder[armnetwork.ExpressRouteProviderPortsLocationClientListResponse], errResp azfake.ErrorResponder)
}

// NewExpressRouteProviderPortsLocationServerTransport creates a new instance of ExpressRouteProviderPortsLocationServerTransport with the provided implementation.
// The returned ExpressRouteProviderPortsLocationServerTransport instance is connected to an instance of armnetwork.ExpressRouteProviderPortsLocationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExpressRouteProviderPortsLocationServerTransport(srv *ExpressRouteProviderPortsLocationServer) *ExpressRouteProviderPortsLocationServerTransport {
	return &ExpressRouteProviderPortsLocationServerTransport{srv: srv}
}

// ExpressRouteProviderPortsLocationServerTransport connects instances of armnetwork.ExpressRouteProviderPortsLocationClient to instances of ExpressRouteProviderPortsLocationServer.
// Don't use this type directly, use NewExpressRouteProviderPortsLocationServerTransport instead.
type ExpressRouteProviderPortsLocationServerTransport struct {
	srv *ExpressRouteProviderPortsLocationServer
}

// Do implements the policy.Transporter interface for ExpressRouteProviderPortsLocationServerTransport.
func (e *ExpressRouteProviderPortsLocationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *ExpressRouteProviderPortsLocationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "ExpressRouteProviderPortsLocationClient.List":
			res.resp, res.err = e.dispatchList(req)
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

func (e *ExpressRouteProviderPortsLocationServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if e.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/expressRouteProviderPorts`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	var options *armnetwork.ExpressRouteProviderPortsLocationClientListOptions
	if filterParam != nil {
		options = &armnetwork.ExpressRouteProviderPortsLocationClientListOptions{
			Filter: filterParam,
		}
	}
	respr, errRespr := e.srv.List(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteProviderPortListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
