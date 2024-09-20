// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"bytesgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BytesPropertyServer is a fake server for instances of the bytesgroup.BytesPropertyClient type.
type BytesPropertyServer struct {
	// Base64 is the fake for method BytesPropertyClient.Base64
	// HTTP status codes to indicate success: http.StatusOK
	Base64 func(ctx context.Context, body bytesgroup.Base64BytesProperty, options *bytesgroup.BytesPropertyClientBase64Options) (resp azfake.Responder[bytesgroup.BytesPropertyClientBase64Response], errResp azfake.ErrorResponder)

	// Base64URL is the fake for method BytesPropertyClient.Base64URL
	// HTTP status codes to indicate success: http.StatusOK
	Base64URL func(ctx context.Context, body bytesgroup.Base64URLBytesProperty, options *bytesgroup.BytesPropertyClientBase64URLOptions) (resp azfake.Responder[bytesgroup.BytesPropertyClientBase64URLResponse], errResp azfake.ErrorResponder)

	// Base64URLArray is the fake for method BytesPropertyClient.Base64URLArray
	// HTTP status codes to indicate success: http.StatusOK
	Base64URLArray func(ctx context.Context, body bytesgroup.Base64URLArrayBytesProperty, options *bytesgroup.BytesPropertyClientBase64URLArrayOptions) (resp azfake.Responder[bytesgroup.BytesPropertyClientBase64URLArrayResponse], errResp azfake.ErrorResponder)

	// Default is the fake for method BytesPropertyClient.Default
	// HTTP status codes to indicate success: http.StatusOK
	Default func(ctx context.Context, body bytesgroup.DefaultBytesProperty, options *bytesgroup.BytesPropertyClientDefaultOptions) (resp azfake.Responder[bytesgroup.BytesPropertyClientDefaultResponse], errResp azfake.ErrorResponder)
}

// NewBytesPropertyServerTransport creates a new instance of BytesPropertyServerTransport with the provided implementation.
// The returned BytesPropertyServerTransport instance is connected to an instance of bytesgroup.BytesPropertyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBytesPropertyServerTransport(srv *BytesPropertyServer) *BytesPropertyServerTransport {
	return &BytesPropertyServerTransport{srv: srv}
}

// BytesPropertyServerTransport connects instances of bytesgroup.BytesPropertyClient to instances of BytesPropertyServer.
// Don't use this type directly, use NewBytesPropertyServerTransport instead.
type BytesPropertyServerTransport struct {
	srv *BytesPropertyServer
}

// Do implements the policy.Transporter interface for BytesPropertyServerTransport.
func (b *BytesPropertyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BytesPropertyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "BytesPropertyClient.Base64":
			res.resp, res.err = b.dispatchBase64(req)
		case "BytesPropertyClient.Base64URL":
			res.resp, res.err = b.dispatchBase64URL(req)
		case "BytesPropertyClient.Base64URLArray":
			res.resp, res.err = b.dispatchBase64URLArray(req)
		case "BytesPropertyClient.Default":
			res.resp, res.err = b.dispatchDefault(req)
		default:
			res.err = fmt.Errorf("unhandled API %s", method)
		}

		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (b *BytesPropertyServerTransport) dispatchBase64(req *http.Request) (*http.Response, error) {
	if b.srv.Base64 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bytesgroup.Base64BytesProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Base64(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Base64BytesProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BytesPropertyServerTransport) dispatchBase64URL(req *http.Request) (*http.Response, error) {
	if b.srv.Base64URL == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64URL not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bytesgroup.Base64URLBytesProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Base64URL(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Base64URLBytesProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BytesPropertyServerTransport) dispatchBase64URLArray(req *http.Request) (*http.Response, error) {
	if b.srv.Base64URLArray == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64URLArray not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bytesgroup.Base64URLArrayBytesProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Base64URLArray(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Base64URLArrayBytesProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BytesPropertyServerTransport) dispatchDefault(req *http.Request) (*http.Response, error) {
	if b.srv.Default == nil {
		return nil, &nonRetriableError{errors.New("fake for method Default not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bytesgroup.DefaultBytesProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Default(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DefaultBytesProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
