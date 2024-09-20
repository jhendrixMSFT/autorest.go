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
	"renamedopgroup"
)

// RenamedOperationGroupServer is a fake server for instances of the renamedopgroup.RenamedOperationGroupClient type.
type RenamedOperationGroupServer struct {
	// RenamedFour is the fake for method RenamedOperationGroupClient.RenamedFour
	// HTTP status codes to indicate success: http.StatusNoContent
	RenamedFour func(ctx context.Context, options *renamedopgroup.RenamedOperationGroupClientRenamedFourOptions) (resp azfake.Responder[renamedopgroup.RenamedOperationGroupClientRenamedFourResponse], errResp azfake.ErrorResponder)

	// RenamedSix is the fake for method RenamedOperationGroupClient.RenamedSix
	// HTTP status codes to indicate success: http.StatusNoContent
	RenamedSix func(ctx context.Context, options *renamedopgroup.RenamedOperationGroupClientRenamedSixOptions) (resp azfake.Responder[renamedopgroup.RenamedOperationGroupClientRenamedSixResponse], errResp azfake.ErrorResponder)

	// RenamedTwo is the fake for method RenamedOperationGroupClient.RenamedTwo
	// HTTP status codes to indicate success: http.StatusNoContent
	RenamedTwo func(ctx context.Context, options *renamedopgroup.RenamedOperationGroupClientRenamedTwoOptions) (resp azfake.Responder[renamedopgroup.RenamedOperationGroupClientRenamedTwoResponse], errResp azfake.ErrorResponder)
}

// NewRenamedOperationGroupServerTransport creates a new instance of RenamedOperationGroupServerTransport with the provided implementation.
// The returned RenamedOperationGroupServerTransport instance is connected to an instance of renamedopgroup.RenamedOperationGroupClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRenamedOperationGroupServerTransport(srv *RenamedOperationGroupServer) *RenamedOperationGroupServerTransport {
	return &RenamedOperationGroupServerTransport{srv: srv}
}

// RenamedOperationGroupServerTransport connects instances of renamedopgroup.RenamedOperationGroupClient to instances of RenamedOperationGroupServer.
// Don't use this type directly, use NewRenamedOperationGroupServerTransport instead.
type RenamedOperationGroupServerTransport struct {
	srv *RenamedOperationGroupServer
}

// Do implements the policy.Transporter interface for RenamedOperationGroupServerTransport.
func (r *RenamedOperationGroupServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RenamedOperationGroupServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "RenamedOperationGroupClient.RenamedFour":
			res.resp, res.err = r.dispatchRenamedFour(req)
		case "RenamedOperationGroupClient.RenamedSix":
			res.resp, res.err = r.dispatchRenamedSix(req)
		case "RenamedOperationGroupClient.RenamedTwo":
			res.resp, res.err = r.dispatchRenamedTwo(req)
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

func (r *RenamedOperationGroupServerTransport) dispatchRenamedFour(req *http.Request) (*http.Response, error) {
	if r.srv.RenamedFour == nil {
		return nil, &nonRetriableError{errors.New("fake for method RenamedFour not implemented")}
	}
	respr, errRespr := r.srv.RenamedFour(req.Context(), nil)
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

func (r *RenamedOperationGroupServerTransport) dispatchRenamedSix(req *http.Request) (*http.Response, error) {
	if r.srv.RenamedSix == nil {
		return nil, &nonRetriableError{errors.New("fake for method RenamedSix not implemented")}
	}
	respr, errRespr := r.srv.RenamedSix(req.Context(), nil)
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

func (r *RenamedOperationGroupServerTransport) dispatchRenamedTwo(req *http.Request) (*http.Response, error) {
	if r.srv.RenamedTwo == nil {
		return nil, &nonRetriableError{errors.New("fake for method RenamedTwo not implemented")}
	}
	respr, errRespr := r.srv.RenamedTwo(req.Context(), nil)
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
