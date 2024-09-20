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

// ValueTypesDurationServer is a fake server for instances of the valuetypesgroup.ValueTypesDurationClient type.
type ValueTypesDurationServer struct {
	// Get is the fake for method ValueTypesDurationClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *valuetypesgroup.ValueTypesDurationClientGetOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesDurationClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ValueTypesDurationClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body valuetypesgroup.DurationProperty, options *valuetypesgroup.ValueTypesDurationClientPutOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesDurationClientPutResponse], errResp azfake.ErrorResponder)
}

// NewValueTypesDurationServerTransport creates a new instance of ValueTypesDurationServerTransport with the provided implementation.
// The returned ValueTypesDurationServerTransport instance is connected to an instance of valuetypesgroup.ValueTypesDurationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewValueTypesDurationServerTransport(srv *ValueTypesDurationServer) *ValueTypesDurationServerTransport {
	return &ValueTypesDurationServerTransport{srv: srv}
}

// ValueTypesDurationServerTransport connects instances of valuetypesgroup.ValueTypesDurationClient to instances of ValueTypesDurationServer.
// Don't use this type directly, use NewValueTypesDurationServerTransport instead.
type ValueTypesDurationServerTransport struct {
	srv *ValueTypesDurationServer
}

// Do implements the policy.Transporter interface for ValueTypesDurationServerTransport.
func (v *ValueTypesDurationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *ValueTypesDurationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "ValueTypesDurationClient.Get":
			res.resp, res.err = v.dispatchGet(req)
		case "ValueTypesDurationClient.Put":
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

func (v *ValueTypesDurationServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DurationProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *ValueTypesDurationServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if v.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[valuetypesgroup.DurationProperty](req)
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
