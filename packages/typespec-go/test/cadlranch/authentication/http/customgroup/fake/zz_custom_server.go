// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"customgroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// CustomServer is a fake server for instances of the customgroup.CustomClient type.
type CustomServer struct {
	// Invalid is the fake for method CustomClient.Invalid
	// HTTP status codes to indicate success: http.StatusNoContent
	Invalid func(ctx context.Context, options *customgroup.CustomClientInvalidOptions) (resp azfake.Responder[customgroup.CustomClientInvalidResponse], errResp azfake.ErrorResponder)

	// Valid is the fake for method CustomClient.Valid
	// HTTP status codes to indicate success: http.StatusNoContent
	Valid func(ctx context.Context, options *customgroup.CustomClientValidOptions) (resp azfake.Responder[customgroup.CustomClientValidResponse], errResp azfake.ErrorResponder)
}

// NewCustomServerTransport creates a new instance of CustomServerTransport with the provided implementation.
// The returned CustomServerTransport instance is connected to an instance of customgroup.CustomClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCustomServerTransport(srv *CustomServer) *CustomServerTransport {
	return &CustomServerTransport{srv: srv}
}

// CustomServerTransport connects instances of customgroup.CustomClient to instances of CustomServer.
// Don't use this type directly, use NewCustomServerTransport instead.
type CustomServerTransport struct {
	srv *CustomServer
}

// Do implements the policy.Transporter interface for CustomServerTransport.
func (c *CustomServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CustomServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "CustomClient.Invalid":
		resp, err = c.dispatchInvalid(req)
	case "CustomClient.Valid":
		resp, err = c.dispatchValid(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (c *CustomServerTransport) dispatchInvalid(req *http.Request) (*http.Response, error) {
	if c.srv.Invalid == nil {
		return nil, &nonRetriableError{errors.New("fake for method Invalid not implemented")}
	}
	respr, errRespr := c.srv.Invalid(req.Context(), nil)
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

func (c *CustomServerTransport) dispatchValid(req *http.Request) (*http.Response, error) {
	if c.srv.Valid == nil {
		return nil, &nonRetriableError{errors.New("fake for method Valid not implemented")}
	}
	respr, errRespr := c.srv.Valid(req.Context(), nil)
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