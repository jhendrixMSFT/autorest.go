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
	"rettypechangedfromgroup"
)

// ReturnTypeChangedFromServer is a fake server for instances of the rettypechangedfromgroup.ReturnTypeChangedFromClient type.
type ReturnTypeChangedFromServer struct {
	// Test is the fake for method ReturnTypeChangedFromClient.Test
	// HTTP status codes to indicate success: http.StatusOK
	Test func(ctx context.Context, body string, options *rettypechangedfromgroup.ReturnTypeChangedFromClientTestOptions) (resp azfake.Responder[rettypechangedfromgroup.ReturnTypeChangedFromClientTestResponse], errResp azfake.ErrorResponder)
}

// NewReturnTypeChangedFromServerTransport creates a new instance of ReturnTypeChangedFromServerTransport with the provided implementation.
// The returned ReturnTypeChangedFromServerTransport instance is connected to an instance of rettypechangedfromgroup.ReturnTypeChangedFromClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReturnTypeChangedFromServerTransport(srv *ReturnTypeChangedFromServer) *ReturnTypeChangedFromServerTransport {
	return &ReturnTypeChangedFromServerTransport{srv: srv}
}

// ReturnTypeChangedFromServerTransport connects instances of rettypechangedfromgroup.ReturnTypeChangedFromClient to instances of ReturnTypeChangedFromServer.
// Don't use this type directly, use NewReturnTypeChangedFromServerTransport instead.
type ReturnTypeChangedFromServerTransport struct {
	srv *ReturnTypeChangedFromServer
}

// Do implements the policy.Transporter interface for ReturnTypeChangedFromServerTransport.
func (r *ReturnTypeChangedFromServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ReturnTypeChangedFromServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "ReturnTypeChangedFromClient.Test":
			res.resp, res.err = r.dispatchTest(req)
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

func (r *ReturnTypeChangedFromServerTransport) dispatchTest(req *http.Request) (*http.Response, error) {
	if r.srv.Test == nil {
		return nil, &nonRetriableError{errors.New("fake for method Test not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[string](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Test(req.Context(), body, nil)
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
