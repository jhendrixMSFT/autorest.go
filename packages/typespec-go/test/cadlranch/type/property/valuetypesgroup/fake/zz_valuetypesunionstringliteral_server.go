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

// ValueTypesUnionStringLiteralServer is a fake server for instances of the valuetypesgroup.ValueTypesUnionStringLiteralClient type.
type ValueTypesUnionStringLiteralServer struct {
	// Get is the fake for method ValueTypesUnionStringLiteralClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *valuetypesgroup.ValueTypesUnionStringLiteralClientGetOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesUnionStringLiteralClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ValueTypesUnionStringLiteralClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body valuetypesgroup.UnionStringLiteralProperty, options *valuetypesgroup.ValueTypesUnionStringLiteralClientPutOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesUnionStringLiteralClientPutResponse], errResp azfake.ErrorResponder)
}

// NewValueTypesUnionStringLiteralServerTransport creates a new instance of ValueTypesUnionStringLiteralServerTransport with the provided implementation.
// The returned ValueTypesUnionStringLiteralServerTransport instance is connected to an instance of valuetypesgroup.ValueTypesUnionStringLiteralClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewValueTypesUnionStringLiteralServerTransport(srv *ValueTypesUnionStringLiteralServer) *ValueTypesUnionStringLiteralServerTransport {
	return &ValueTypesUnionStringLiteralServerTransport{srv: srv}
}

// ValueTypesUnionStringLiteralServerTransport connects instances of valuetypesgroup.ValueTypesUnionStringLiteralClient to instances of ValueTypesUnionStringLiteralServer.
// Don't use this type directly, use NewValueTypesUnionStringLiteralServerTransport instead.
type ValueTypesUnionStringLiteralServerTransport struct {
	srv *ValueTypesUnionStringLiteralServer
}

// Do implements the policy.Transporter interface for ValueTypesUnionStringLiteralServerTransport.
func (v *ValueTypesUnionStringLiteralServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *ValueTypesUnionStringLiteralServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ValueTypesUnionStringLiteralClient.Get":
		resp, err = v.dispatchGet(req)
	case "ValueTypesUnionStringLiteralClient.Put":
		resp, err = v.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (v *ValueTypesUnionStringLiteralServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UnionStringLiteralProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *ValueTypesUnionStringLiteralServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if v.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[valuetypesgroup.UnionStringLiteralProperty](req)
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
