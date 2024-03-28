// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"contentneggroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DifferentBodyServer is a fake server for instances of the contentneggroup.DifferentBodyClient type.
type DifferentBodyServer struct {
	// GetAvatarAsJSON is the fake for method DifferentBodyClient.GetAvatarAsJSON
	// HTTP status codes to indicate success: http.StatusOK
	GetAvatarAsJSON func(ctx context.Context, options *contentneggroup.DifferentBodyClientGetAvatarAsJSONOptions) (resp azfake.Responder[contentneggroup.DifferentBodyClientGetAvatarAsJSONResponse], errResp azfake.ErrorResponder)

	// GetAvatarAsPNG is the fake for method DifferentBodyClient.GetAvatarAsPNG
	// HTTP status codes to indicate success: http.StatusOK
	GetAvatarAsPNG func(ctx context.Context, options *contentneggroup.DifferentBodyClientGetAvatarAsPNGOptions) (resp azfake.Responder[contentneggroup.DifferentBodyClientGetAvatarAsPNGResponse], errResp azfake.ErrorResponder)
}

// NewDifferentBodyServerTransport creates a new instance of DifferentBodyServerTransport with the provided implementation.
// The returned DifferentBodyServerTransport instance is connected to an instance of contentneggroup.DifferentBodyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDifferentBodyServerTransport(srv *DifferentBodyServer) *DifferentBodyServerTransport {
	return &DifferentBodyServerTransport{srv: srv}
}

// DifferentBodyServerTransport connects instances of contentneggroup.DifferentBodyClient to instances of DifferentBodyServer.
// Don't use this type directly, use NewDifferentBodyServerTransport instead.
type DifferentBodyServerTransport struct {
	srv *DifferentBodyServer
}

// Do implements the policy.Transporter interface for DifferentBodyServerTransport.
func (d *DifferentBodyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DifferentBodyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "DifferentBodyClient.GetAvatarAsJSON":
		resp, err = d.dispatchGetAvatarAsJSON(req)
	case "DifferentBodyClient.GetAvatarAsPNG":
		resp, err = d.dispatchGetAvatarAsPNG(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (d *DifferentBodyServerTransport) dispatchGetAvatarAsJSON(req *http.Request) (*http.Response, error) {
	if d.srv.GetAvatarAsJSON == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAvatarAsJSON not implemented")}
	}
	respr, errRespr := d.srv.GetAvatarAsJSON(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PNGImageAsJSON, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DifferentBodyServerTransport) dispatchGetAvatarAsPNG(req *http.Request) (*http.Response, error) {
	if d.srv.GetAvatarAsPNG == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAvatarAsPNG not implemented")}
	}
	respr, errRespr := d.srv.GetAvatarAsPNG(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: req.Header.Get("Content-Type"),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
