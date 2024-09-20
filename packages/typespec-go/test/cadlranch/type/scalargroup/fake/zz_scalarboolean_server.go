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
	"scalargroup"
)

// ScalarBooleanServer is a fake server for instances of the scalargroup.ScalarBooleanClient type.
type ScalarBooleanServer struct {
	// Get is the fake for method ScalarBooleanClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *scalargroup.ScalarBooleanClientGetOptions) (resp azfake.Responder[scalargroup.ScalarBooleanClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ScalarBooleanClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body bool, options *scalargroup.ScalarBooleanClientPutOptions) (resp azfake.Responder[scalargroup.ScalarBooleanClientPutResponse], errResp azfake.ErrorResponder)
}

// NewScalarBooleanServerTransport creates a new instance of ScalarBooleanServerTransport with the provided implementation.
// The returned ScalarBooleanServerTransport instance is connected to an instance of scalargroup.ScalarBooleanClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScalarBooleanServerTransport(srv *ScalarBooleanServer) *ScalarBooleanServerTransport {
	return &ScalarBooleanServerTransport{srv: srv}
}

// ScalarBooleanServerTransport connects instances of scalargroup.ScalarBooleanClient to instances of ScalarBooleanServer.
// Don't use this type directly, use NewScalarBooleanServerTransport instead.
type ScalarBooleanServerTransport struct {
	srv *ScalarBooleanServer
}

// Do implements the policy.Transporter interface for ScalarBooleanServerTransport.
func (s *ScalarBooleanServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ScalarBooleanServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "ScalarBooleanClient.Get":
			res.resp, res.err = s.dispatchGet(req)
		case "ScalarBooleanClient.Put":
			res.resp, res.err = s.dispatchPut(req)
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

func (s *ScalarBooleanServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := s.srv.Get(req.Context(), nil)
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

func (s *ScalarBooleanServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if s.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Put(req.Context(), body, nil)
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
