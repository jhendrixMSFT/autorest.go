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
	"uniongroup"
)

// UnionStringAndArrayServer is a fake server for instances of the uniongroup.UnionStringAndArrayClient type.
type UnionStringAndArrayServer struct {
	// Get is the fake for method UnionStringAndArrayClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *uniongroup.UnionStringAndArrayClientGetOptions) (resp azfake.Responder[uniongroup.UnionStringAndArrayClientGetResponse], errResp azfake.ErrorResponder)

	// Send is the fake for method UnionStringAndArrayClient.Send
	// HTTP status codes to indicate success: http.StatusNoContent
	Send func(ctx context.Context, sendRequest2 uniongroup.SendRequest2, options *uniongroup.UnionStringAndArrayClientSendOptions) (resp azfake.Responder[uniongroup.UnionStringAndArrayClientSendResponse], errResp azfake.ErrorResponder)
}

// NewUnionStringAndArrayServerTransport creates a new instance of UnionStringAndArrayServerTransport with the provided implementation.
// The returned UnionStringAndArrayServerTransport instance is connected to an instance of uniongroup.UnionStringAndArrayClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUnionStringAndArrayServerTransport(srv *UnionStringAndArrayServer) *UnionStringAndArrayServerTransport {
	return &UnionStringAndArrayServerTransport{srv: srv}
}

// UnionStringAndArrayServerTransport connects instances of uniongroup.UnionStringAndArrayClient to instances of UnionStringAndArrayServer.
// Don't use this type directly, use NewUnionStringAndArrayServerTransport instead.
type UnionStringAndArrayServerTransport struct {
	srv *UnionStringAndArrayServer
}

// Do implements the policy.Transporter interface for UnionStringAndArrayServerTransport.
func (u *UnionStringAndArrayServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return u.dispatchToMethodFake(req, method)
}

func (u *UnionStringAndArrayServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "UnionStringAndArrayClient.Get":
		resp, err = u.dispatchGet(req)
	case "UnionStringAndArrayClient.Send":
		resp, err = u.dispatchSend(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (u *UnionStringAndArrayServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if u.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := u.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GetResponse2, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UnionStringAndArrayServerTransport) dispatchSend(req *http.Request) (*http.Response, error) {
	if u.srv.Send == nil {
		return nil, &nonRetriableError{errors.New("fake for method Send not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[uniongroup.SendRequest2](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Send(req.Context(), body, nil)
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
