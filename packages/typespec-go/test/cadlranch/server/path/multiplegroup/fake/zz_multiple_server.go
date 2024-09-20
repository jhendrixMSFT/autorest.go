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
	"multiplegroup"
	"net/http"
	"net/url"
	"regexp"
)

// MultipleServer is a fake server for instances of the multiplegroup.MultipleClient type.
type MultipleServer struct {
	// NoOperationParams is the fake for method MultipleClient.NoOperationParams
	// HTTP status codes to indicate success: http.StatusNoContent
	NoOperationParams func(ctx context.Context, options *multiplegroup.MultipleClientNoOperationParamsOptions) (resp azfake.Responder[multiplegroup.MultipleClientNoOperationParamsResponse], errResp azfake.ErrorResponder)

	// WithOperationPathParam is the fake for method MultipleClient.WithOperationPathParam
	// HTTP status codes to indicate success: http.StatusNoContent
	WithOperationPathParam func(ctx context.Context, keyword string, options *multiplegroup.MultipleClientWithOperationPathParamOptions) (resp azfake.Responder[multiplegroup.MultipleClientWithOperationPathParamResponse], errResp azfake.ErrorResponder)
}

// NewMultipleServerTransport creates a new instance of MultipleServerTransport with the provided implementation.
// The returned MultipleServerTransport instance is connected to an instance of multiplegroup.MultipleClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMultipleServerTransport(srv *MultipleServer) *MultipleServerTransport {
	return &MultipleServerTransport{srv: srv}
}

// MultipleServerTransport connects instances of multiplegroup.MultipleClient to instances of MultipleServer.
// Don't use this type directly, use NewMultipleServerTransport instead.
type MultipleServerTransport struct {
	srv *MultipleServer
}

// Do implements the policy.Transporter interface for MultipleServerTransport.
func (m *MultipleServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MultipleServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "MultipleClient.NoOperationParams":
			res.resp, res.err = m.dispatchNoOperationParams(req)
		case "MultipleClient.WithOperationPathParam":
			res.resp, res.err = m.dispatchWithOperationPathParam(req)
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

func (m *MultipleServerTransport) dispatchNoOperationParams(req *http.Request) (*http.Response, error) {
	if m.srv.NoOperationParams == nil {
		return nil, &nonRetriableError{errors.New("fake for method NoOperationParams not implemented")}
	}
	respr, errRespr := m.srv.NoOperationParams(req.Context(), nil)
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

func (m *MultipleServerTransport) dispatchWithOperationPathParam(req *http.Request) (*http.Response, error) {
	if m.srv.WithOperationPathParam == nil {
		return nil, &nonRetriableError{errors.New("fake for method WithOperationPathParam not implemented")}
	}
	const regexStr = `/(?P<keyword>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	keywordParam, err := url.PathUnescape(matches[regex.SubexpIndex("keyword")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.WithOperationPathParam(req.Context(), keywordParam, nil)
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
