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
	"naminggroup"
	"net/http"
)

// NamingUnionEnumServer is a fake server for instances of the naminggroup.NamingUnionEnumClient type.
type NamingUnionEnumServer struct {
	// UnionEnumMemberName is the fake for method NamingUnionEnumClient.UnionEnumMemberName
	// HTTP status codes to indicate success: http.StatusNoContent
	UnionEnumMemberName func(ctx context.Context, body naminggroup.ExtensibleEnum, options *naminggroup.NamingUnionEnumClientUnionEnumMemberNameOptions) (resp azfake.Responder[naminggroup.NamingUnionEnumClientUnionEnumMemberNameResponse], errResp azfake.ErrorResponder)

	// UnionEnumName is the fake for method NamingUnionEnumClient.UnionEnumName
	// HTTP status codes to indicate success: http.StatusNoContent
	UnionEnumName func(ctx context.Context, body naminggroup.ClientExtensibleEnum, options *naminggroup.NamingUnionEnumClientUnionEnumNameOptions) (resp azfake.Responder[naminggroup.NamingUnionEnumClientUnionEnumNameResponse], errResp azfake.ErrorResponder)
}

// NewNamingUnionEnumServerTransport creates a new instance of NamingUnionEnumServerTransport with the provided implementation.
// The returned NamingUnionEnumServerTransport instance is connected to an instance of naminggroup.NamingUnionEnumClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNamingUnionEnumServerTransport(srv *NamingUnionEnumServer) *NamingUnionEnumServerTransport {
	return &NamingUnionEnumServerTransport{srv: srv}
}

// NamingUnionEnumServerTransport connects instances of naminggroup.NamingUnionEnumClient to instances of NamingUnionEnumServer.
// Don't use this type directly, use NewNamingUnionEnumServerTransport instead.
type NamingUnionEnumServerTransport struct {
	srv *NamingUnionEnumServer
}

// Do implements the policy.Transporter interface for NamingUnionEnumServerTransport.
func (n *NamingUnionEnumServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NamingUnionEnumServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "NamingUnionEnumClient.UnionEnumMemberName":
			res.resp, res.err = n.dispatchUnionEnumMemberName(req)
		case "NamingUnionEnumClient.UnionEnumName":
			res.resp, res.err = n.dispatchUnionEnumName(req)
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

func (n *NamingUnionEnumServerTransport) dispatchUnionEnumMemberName(req *http.Request) (*http.Response, error) {
	if n.srv.UnionEnumMemberName == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnionEnumMemberName not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[naminggroup.ExtensibleEnum](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.UnionEnumMemberName(req.Context(), body, nil)
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

func (n *NamingUnionEnumServerTransport) dispatchUnionEnumName(req *http.Request) (*http.Response, error) {
	if n.srv.UnionEnumName == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnionEnumName not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[naminggroup.ClientExtensibleEnum](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.UnionEnumName(req.Context(), body, nil)
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
