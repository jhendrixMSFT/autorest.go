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
	"singlegroup"
)

// SingleServer is a fake server for instances of the singlegroup.SingleClient type.
type SingleServer struct {
	// MyOp is the fake for method SingleClient.MyOp
	// HTTP status codes to indicate success: http.StatusOK
	MyOp func(ctx context.Context, options *singlegroup.SingleClientMyOpOptions) (resp azfake.Responder[singlegroup.SingleClientMyOpResponse], errResp azfake.ErrorResponder)
}

// NewSingleServerTransport creates a new instance of SingleServerTransport with the provided implementation.
// The returned SingleServerTransport instance is connected to an instance of singlegroup.SingleClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSingleServerTransport(srv *SingleServer) *SingleServerTransport {
	return &SingleServerTransport{srv: srv}
}

// SingleServerTransport connects instances of singlegroup.SingleClient to instances of SingleServer.
// Don't use this type directly, use NewSingleServerTransport instead.
type SingleServerTransport struct {
	srv *SingleServer
}

// Do implements the policy.Transporter interface for SingleServerTransport.
func (s *SingleServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SingleServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "SingleClient.MyOp":
			res.resp, res.err = s.dispatchMyOp(req)
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

func (s *SingleServerTransport) dispatchMyOp(req *http.Request) (*http.Response, error) {
	if s.srv.MyOp == nil {
		return nil, &nonRetriableError{errors.New("fake for method MyOp not implemented")}
	}
	respr, errRespr := s.srv.MyOp(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
