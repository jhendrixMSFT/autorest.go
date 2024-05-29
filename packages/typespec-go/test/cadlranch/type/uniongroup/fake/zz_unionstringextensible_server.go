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

// UnionStringExtensibleServer is a fake server for instances of the uniongroup.UnionStringExtensibleClient type.
type UnionStringExtensibleServer struct {
	// Get is the fake for method UnionStringExtensibleClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *uniongroup.UnionStringExtensibleClientGetOptions) (resp azfake.Responder[uniongroup.UnionStringExtensibleClientGetResponse], errResp azfake.ErrorResponder)

	// Send is the fake for method UnionStringExtensibleClient.Send
	// HTTP status codes to indicate success: http.StatusNoContent
	Send func(ctx context.Context, prop uniongroup.GetResponseProp4, options *uniongroup.UnionStringExtensibleClientSendOptions) (resp azfake.Responder[uniongroup.UnionStringExtensibleClientSendResponse], errResp azfake.ErrorResponder)
}

// NewUnionStringExtensibleServerTransport creates a new instance of UnionStringExtensibleServerTransport with the provided implementation.
// The returned UnionStringExtensibleServerTransport instance is connected to an instance of uniongroup.UnionStringExtensibleClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUnionStringExtensibleServerTransport(srv *UnionStringExtensibleServer) *UnionStringExtensibleServerTransport {
	return &UnionStringExtensibleServerTransport{srv: srv}
}

// UnionStringExtensibleServerTransport connects instances of uniongroup.UnionStringExtensibleClient to instances of UnionStringExtensibleServer.
// Don't use this type directly, use NewUnionStringExtensibleServerTransport instead.
type UnionStringExtensibleServerTransport struct {
	srv *UnionStringExtensibleServer
}

// Do implements the policy.Transporter interface for UnionStringExtensibleServerTransport.
func (u *UnionStringExtensibleServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return u.dispatchToMethodFake(req, method)
}

func (u *UnionStringExtensibleServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "UnionStringExtensibleClient.Get":
		resp, err = u.dispatchGet(req)
	case "UnionStringExtensibleClient.Send":
		resp, err = u.dispatchSend(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (u *UnionStringExtensibleServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GetResponse8, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UnionStringExtensibleServerTransport) dispatchSend(req *http.Request) (*http.Response, error) {
	if u.srv.Send == nil {
		return nil, &nonRetriableError{errors.New("fake for method Send not implemented")}
	}
	type partialBodyParams struct {
		Prop GetResponseProp4 `json:"prop"`
	}
	body, err := server.UnmarshalRequestAsJSON[partialBodyParams](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Send(req.Context(), body.Prop, nil)
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
