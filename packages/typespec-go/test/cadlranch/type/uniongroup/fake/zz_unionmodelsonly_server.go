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

// UnionModelsOnlyServer is a fake server for instances of the uniongroup.UnionModelsOnlyClient type.
type UnionModelsOnlyServer struct {
	// Get is the fake for method UnionModelsOnlyClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *uniongroup.UnionModelsOnlyClientGetOptions) (resp azfake.Responder[uniongroup.UnionModelsOnlyClientGetResponse], errResp azfake.ErrorResponder)

	// Send is the fake for method UnionModelsOnlyClient.Send
	// HTTP status codes to indicate success: http.StatusNoContent
	Send func(ctx context.Context, prop []byte, options *uniongroup.UnionModelsOnlyClientSendOptions) (resp azfake.Responder[uniongroup.UnionModelsOnlyClientSendResponse], errResp azfake.ErrorResponder)
}

// NewUnionModelsOnlyServerTransport creates a new instance of UnionModelsOnlyServerTransport with the provided implementation.
// The returned UnionModelsOnlyServerTransport instance is connected to an instance of uniongroup.UnionModelsOnlyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUnionModelsOnlyServerTransport(srv *UnionModelsOnlyServer) *UnionModelsOnlyServerTransport {
	return &UnionModelsOnlyServerTransport{srv: srv}
}

// UnionModelsOnlyServerTransport connects instances of uniongroup.UnionModelsOnlyClient to instances of UnionModelsOnlyServer.
// Don't use this type directly, use NewUnionModelsOnlyServerTransport instead.
type UnionModelsOnlyServerTransport struct {
	srv *UnionModelsOnlyServer
}

// Do implements the policy.Transporter interface for UnionModelsOnlyServerTransport.
func (u *UnionModelsOnlyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return u.dispatchToMethodFake(req, method)
}

func (u *UnionModelsOnlyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "UnionModelsOnlyClient.Get":
		resp, err = u.dispatchGet(req)
	case "UnionModelsOnlyClient.Send":
		resp, err = u.dispatchSend(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (u *UnionModelsOnlyServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GetResponse4, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UnionModelsOnlyServerTransport) dispatchSend(req *http.Request) (*http.Response, error) {
	if u.srv.Send == nil {
		return nil, &nonRetriableError{errors.New("fake for method Send not implemented")}
	}
	type partialBodyParams struct {
		Prop []byte `json:"prop"`
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
