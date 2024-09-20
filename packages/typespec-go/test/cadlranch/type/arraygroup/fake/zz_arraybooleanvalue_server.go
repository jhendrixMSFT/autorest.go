// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"arraygroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ArrayBooleanValueServer is a fake server for instances of the arraygroup.ArrayBooleanValueClient type.
type ArrayBooleanValueServer struct {
	// Get is the fake for method ArrayBooleanValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *arraygroup.ArrayBooleanValueClientGetOptions) (resp azfake.Responder[arraygroup.ArrayBooleanValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ArrayBooleanValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body []bool, options *arraygroup.ArrayBooleanValueClientPutOptions) (resp azfake.Responder[arraygroup.ArrayBooleanValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewArrayBooleanValueServerTransport creates a new instance of ArrayBooleanValueServerTransport with the provided implementation.
// The returned ArrayBooleanValueServerTransport instance is connected to an instance of arraygroup.ArrayBooleanValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewArrayBooleanValueServerTransport(srv *ArrayBooleanValueServer) *ArrayBooleanValueServerTransport {
	return &ArrayBooleanValueServerTransport{srv: srv}
}

// ArrayBooleanValueServerTransport connects instances of arraygroup.ArrayBooleanValueClient to instances of ArrayBooleanValueServer.
// Don't use this type directly, use NewArrayBooleanValueServerTransport instead.
type ArrayBooleanValueServerTransport struct {
	srv *ArrayBooleanValueServer
}

// Do implements the policy.Transporter interface for ArrayBooleanValueServerTransport.
func (a *ArrayBooleanValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *ArrayBooleanValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "ArrayBooleanValueClient.Get":
			res.resp, res.err = a.dispatchGet(req)
		case "ArrayBooleanValueClient.Put":
			res.resp, res.err = a.dispatchPut(req)
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

func (a *ArrayBooleanValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArrayBooleanValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if a.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[[]bool](req)
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
