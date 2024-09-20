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
	"valuetypesgroup"
)

// ValueTypesBytesServer is a fake server for instances of the valuetypesgroup.ValueTypesBytesClient type.
type ValueTypesBytesServer struct {
	// Get is the fake for method ValueTypesBytesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *valuetypesgroup.ValueTypesBytesClientGetOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesBytesClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ValueTypesBytesClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body valuetypesgroup.BytesProperty, options *valuetypesgroup.ValueTypesBytesClientPutOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesBytesClientPutResponse], errResp azfake.ErrorResponder)
}

// NewValueTypesBytesServerTransport creates a new instance of ValueTypesBytesServerTransport with the provided implementation.
// The returned ValueTypesBytesServerTransport instance is connected to an instance of valuetypesgroup.ValueTypesBytesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewValueTypesBytesServerTransport(srv *ValueTypesBytesServer) *ValueTypesBytesServerTransport {
	return &ValueTypesBytesServerTransport{srv: srv}
}

// ValueTypesBytesServerTransport connects instances of valuetypesgroup.ValueTypesBytesClient to instances of ValueTypesBytesServer.
// Don't use this type directly, use NewValueTypesBytesServerTransport instead.
type ValueTypesBytesServerTransport struct {
	srv *ValueTypesBytesServer
}

// Do implements the policy.Transporter interface for ValueTypesBytesServerTransport.
func (v *ValueTypesBytesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *ValueTypesBytesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "ValueTypesBytesClient.Get":
			res.resp, res.err = v.dispatchGet(req)
		case "ValueTypesBytesClient.Put":
			res.resp, res.err = v.dispatchPut(req)
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

func (v *ValueTypesBytesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := v.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BytesProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *ValueTypesBytesServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if v.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[valuetypesgroup.BytesProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Put(req.Context(), body, nil)
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
