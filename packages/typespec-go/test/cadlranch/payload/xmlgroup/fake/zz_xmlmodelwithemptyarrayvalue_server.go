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
	"xmlgroup"
)

// XMLModelWithEmptyArrayValueServer is a fake server for instances of the xmlgroup.XMLModelWithEmptyArrayValueClient type.
type XMLModelWithEmptyArrayValueServer struct {
	// Get is the fake for method XMLModelWithEmptyArrayValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *xmlgroup.XMLModelWithEmptyArrayValueClientGetOptions) (resp azfake.Responder[xmlgroup.XMLModelWithEmptyArrayValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method XMLModelWithEmptyArrayValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, input xmlgroup.ModelWithEmptyArray, options *xmlgroup.XMLModelWithEmptyArrayValueClientPutOptions) (resp azfake.Responder[xmlgroup.XMLModelWithEmptyArrayValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewXMLModelWithEmptyArrayValueServerTransport creates a new instance of XMLModelWithEmptyArrayValueServerTransport with the provided implementation.
// The returned XMLModelWithEmptyArrayValueServerTransport instance is connected to an instance of xmlgroup.XMLModelWithEmptyArrayValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewXMLModelWithEmptyArrayValueServerTransport(srv *XMLModelWithEmptyArrayValueServer) *XMLModelWithEmptyArrayValueServerTransport {
	return &XMLModelWithEmptyArrayValueServerTransport{srv: srv}
}

// XMLModelWithEmptyArrayValueServerTransport connects instances of xmlgroup.XMLModelWithEmptyArrayValueClient to instances of XMLModelWithEmptyArrayValueServer.
// Don't use this type directly, use NewXMLModelWithEmptyArrayValueServerTransport instead.
type XMLModelWithEmptyArrayValueServerTransport struct {
	srv *XMLModelWithEmptyArrayValueServer
}

// Do implements the policy.Transporter interface for XMLModelWithEmptyArrayValueServerTransport.
func (x *XMLModelWithEmptyArrayValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return x.dispatchToMethodFake(req, method)
}

func (x *XMLModelWithEmptyArrayValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "XMLModelWithEmptyArrayValueClient.Get":
		resp, err = x.dispatchGet(req)
	case "XMLModelWithEmptyArrayValueClient.Put":
		resp, err = x.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (x *XMLModelWithEmptyArrayValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if x.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := x.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).ModelWithEmptyArray, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("content-type", "application/xml")
	}
	return resp, nil
}

func (x *XMLModelWithEmptyArrayValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if x.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsXML[xmlgroup.ModelWithEmptyArray](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := x.srv.Put(req.Context(), body, nil)
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
