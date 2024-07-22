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

// XMLModelWithRenamedFieldsValueServer is a fake server for instances of the xmlgroup.XMLModelWithRenamedFieldsValueClient type.
type XMLModelWithRenamedFieldsValueServer struct {
	// Get is the fake for method XMLModelWithRenamedFieldsValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *xmlgroup.XMLModelWithRenamedFieldsValueClientGetOptions) (resp azfake.Responder[xmlgroup.XMLModelWithRenamedFieldsValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method XMLModelWithRenamedFieldsValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, input xmlgroup.ModelWithRenamedFields, options *xmlgroup.XMLModelWithRenamedFieldsValueClientPutOptions) (resp azfake.Responder[xmlgroup.XMLModelWithRenamedFieldsValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewXMLModelWithRenamedFieldsValueServerTransport creates a new instance of XMLModelWithRenamedFieldsValueServerTransport with the provided implementation.
// The returned XMLModelWithRenamedFieldsValueServerTransport instance is connected to an instance of xmlgroup.XMLModelWithRenamedFieldsValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewXMLModelWithRenamedFieldsValueServerTransport(srv *XMLModelWithRenamedFieldsValueServer) *XMLModelWithRenamedFieldsValueServerTransport {
	return &XMLModelWithRenamedFieldsValueServerTransport{srv: srv}
}

// XMLModelWithRenamedFieldsValueServerTransport connects instances of xmlgroup.XMLModelWithRenamedFieldsValueClient to instances of XMLModelWithRenamedFieldsValueServer.
// Don't use this type directly, use NewXMLModelWithRenamedFieldsValueServerTransport instead.
type XMLModelWithRenamedFieldsValueServerTransport struct {
	srv *XMLModelWithRenamedFieldsValueServer
}

// Do implements the policy.Transporter interface for XMLModelWithRenamedFieldsValueServerTransport.
func (x *XMLModelWithRenamedFieldsValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return x.dispatchToMethodFake(req, method)
}

func (x *XMLModelWithRenamedFieldsValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "XMLModelWithRenamedFieldsValueClient.Get":
		resp, err = x.dispatchGet(req)
	case "XMLModelWithRenamedFieldsValueClient.Put":
		resp, err = x.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (x *XMLModelWithRenamedFieldsValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).ModelWithRenamedFields, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("content-type", "application/xml")
	}
	return resp, nil
}

func (x *XMLModelWithRenamedFieldsValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if x.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsXML[xmlgroup.ModelWithRenamedFields](req)
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
