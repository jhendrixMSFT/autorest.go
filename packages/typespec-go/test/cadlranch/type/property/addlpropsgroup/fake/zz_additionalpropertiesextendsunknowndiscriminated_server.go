// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"addlpropsgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AdditionalPropertiesExtendsUnknownDiscriminatedServer is a fake server for instances of the addlpropsgroup.AdditionalPropertiesExtendsUnknownDiscriminatedClient type.
type AdditionalPropertiesExtendsUnknownDiscriminatedServer struct {
	// Get is the fake for method AdditionalPropertiesExtendsUnknownDiscriminatedClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *addlpropsgroup.AdditionalPropertiesExtendsUnknownDiscriminatedClientGetOptions) (resp azfake.Responder[addlpropsgroup.AdditionalPropertiesExtendsUnknownDiscriminatedClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method AdditionalPropertiesExtendsUnknownDiscriminatedClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body addlpropsgroup.ExtendsUnknownAdditionalPropertiesDiscriminatedClassification, options *addlpropsgroup.AdditionalPropertiesExtendsUnknownDiscriminatedClientPutOptions) (resp azfake.Responder[addlpropsgroup.AdditionalPropertiesExtendsUnknownDiscriminatedClientPutResponse], errResp azfake.ErrorResponder)
}

// NewAdditionalPropertiesExtendsUnknownDiscriminatedServerTransport creates a new instance of AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport with the provided implementation.
// The returned AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport instance is connected to an instance of addlpropsgroup.AdditionalPropertiesExtendsUnknownDiscriminatedClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAdditionalPropertiesExtendsUnknownDiscriminatedServerTransport(srv *AdditionalPropertiesExtendsUnknownDiscriminatedServer) *AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport {
	return &AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport{srv: srv}
}

// AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport connects instances of addlpropsgroup.AdditionalPropertiesExtendsUnknownDiscriminatedClient to instances of AdditionalPropertiesExtendsUnknownDiscriminatedServer.
// Don't use this type directly, use NewAdditionalPropertiesExtendsUnknownDiscriminatedServerTransport instead.
type AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport struct {
	srv *AdditionalPropertiesExtendsUnknownDiscriminatedServer
}

// Do implements the policy.Transporter interface for AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport.
func (a *AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "AdditionalPropertiesExtendsUnknownDiscriminatedClient.Get":
		resp, err = a.dispatchGet(req)
	case "AdditionalPropertiesExtendsUnknownDiscriminatedClient.Put":
		resp, err = a.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (a *AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := a.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExtendsUnknownAdditionalPropertiesDiscriminatedClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AdditionalPropertiesExtendsUnknownDiscriminatedServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if a.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	raw, err := readRequestBody(req)
	if err != nil {
		return nil, err
	}
	body, err := unmarshalExtendsUnknownAdditionalPropertiesDiscriminatedClassification(raw)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Put(req.Context(), body, nil)
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