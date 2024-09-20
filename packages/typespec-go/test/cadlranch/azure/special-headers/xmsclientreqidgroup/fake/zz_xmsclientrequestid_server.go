// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"xmsclientreqidgroup"
)

// XMSClientRequestIDServer is a fake server for instances of the xmsclientreqidgroup.XMSClientRequestIDClient type.
type XMSClientRequestIDServer struct {
	// Get is the fake for method XMSClientRequestIDClient.Get
	// HTTP status codes to indicate success: http.StatusNoContent
	Get func(ctx context.Context, options *xmsclientreqidgroup.XMSClientRequestIDClientGetOptions) (resp azfake.Responder[xmsclientreqidgroup.XMSClientRequestIDClientGetResponse], errResp azfake.ErrorResponder)
}

// NewXMSClientRequestIDServerTransport creates a new instance of XMSClientRequestIDServerTransport with the provided implementation.
// The returned XMSClientRequestIDServerTransport instance is connected to an instance of xmsclientreqidgroup.XMSClientRequestIDClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewXMSClientRequestIDServerTransport(srv *XMSClientRequestIDServer) *XMSClientRequestIDServerTransport {
	return &XMSClientRequestIDServerTransport{srv: srv}
}

// XMSClientRequestIDServerTransport connects instances of xmsclientreqidgroup.XMSClientRequestIDClient to instances of XMSClientRequestIDServer.
// Don't use this type directly, use NewXMSClientRequestIDServerTransport instead.
type XMSClientRequestIDServerTransport struct {
	srv *XMSClientRequestIDServer
}

// Do implements the policy.Transporter interface for XMSClientRequestIDServerTransport.
func (x *XMSClientRequestIDServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return x.dispatchToMethodFake(req, method)
}

func (x *XMSClientRequestIDServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "XMSClientRequestIDClient.Get":
			res.resp, res.err = x.dispatchGet(req)
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

func (x *XMSClientRequestIDServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if x.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var options *xmsclientreqidgroup.XMSClientRequestIDClientGetOptions
	if clientRequestIDParam != nil {
		options = &xmsclientreqidgroup.XMSClientRequestIDClientGetOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := x.srv.Get(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
