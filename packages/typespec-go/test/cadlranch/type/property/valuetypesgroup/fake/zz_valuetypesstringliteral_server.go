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

// ValueTypesStringLiteralServer is a fake server for instances of the valuetypesgroup.ValueTypesStringLiteralClient type.
type ValueTypesStringLiteralServer struct {
	// Get is the fake for method ValueTypesStringLiteralClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *valuetypesgroup.ValueTypesStringLiteralClientGetOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesStringLiteralClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ValueTypesStringLiteralClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body valuetypesgroup.StringLiteralProperty, options *valuetypesgroup.ValueTypesStringLiteralClientPutOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesStringLiteralClientPutResponse], errResp azfake.ErrorResponder)
}

// NewValueTypesStringLiteralServerTransport creates a new instance of ValueTypesStringLiteralServerTransport with the provided implementation.
// The returned ValueTypesStringLiteralServerTransport instance is connected to an instance of valuetypesgroup.ValueTypesStringLiteralClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewValueTypesStringLiteralServerTransport(srv *ValueTypesStringLiteralServer) *ValueTypesStringLiteralServerTransport {
	return &ValueTypesStringLiteralServerTransport{srv: srv}
}

// ValueTypesStringLiteralServerTransport connects instances of valuetypesgroup.ValueTypesStringLiteralClient to instances of ValueTypesStringLiteralServer.
// Don't use this type directly, use NewValueTypesStringLiteralServerTransport instead.
type ValueTypesStringLiteralServerTransport struct {
	srv *ValueTypesStringLiteralServer
}

// Do implements the policy.Transporter interface for ValueTypesStringLiteralServerTransport.
func (v *ValueTypesStringLiteralServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *ValueTypesStringLiteralServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ValueTypesStringLiteralClient.Get":
		resp, err = v.dispatchGet(req)
	case "ValueTypesStringLiteralClient.Put":
		resp, err = v.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (v *ValueTypesStringLiteralServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringLiteralProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *ValueTypesStringLiteralServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if v.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[valuetypesgroup.StringLiteralProperty](req)
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
