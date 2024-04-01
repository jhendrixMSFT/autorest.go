// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"bodyoptionalgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// BodyOptionalityServer is a fake server for instances of the bodyoptionalgroup.BodyOptionalityClient type.
type BodyOptionalityServer struct {
	// BodyOptionalityOptionalExplicitServer contains the fakes for client BodyOptionalityOptionalExplicitClient
	BodyOptionalityOptionalExplicitServer BodyOptionalityOptionalExplicitServer

	// RequiredExplicit is the fake for method BodyOptionalityClient.RequiredExplicit
	// HTTP status codes to indicate success: http.StatusNoContent
	RequiredExplicit func(ctx context.Context, body bodyoptionalgroup.BodyModel, options *bodyoptionalgroup.BodyOptionalityClientRequiredExplicitOptions) (resp azfake.Responder[bodyoptionalgroup.BodyOptionalityClientRequiredExplicitResponse], errResp azfake.ErrorResponder)

	// RequiredImplicit is the fake for method BodyOptionalityClient.RequiredImplicit
	// HTTP status codes to indicate success: http.StatusNoContent
	RequiredImplicit func(ctx context.Context, body bodyoptionalgroup.BodyModel, options *bodyoptionalgroup.BodyOptionalityClientRequiredImplicitOptions) (resp azfake.Responder[bodyoptionalgroup.BodyOptionalityClientRequiredImplicitResponse], errResp azfake.ErrorResponder)
}

// NewBodyOptionalityServerTransport creates a new instance of BodyOptionalityServerTransport with the provided implementation.
// The returned BodyOptionalityServerTransport instance is connected to an instance of bodyoptionalgroup.BodyOptionalityClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBodyOptionalityServerTransport(srv *BodyOptionalityServer) *BodyOptionalityServerTransport {
	return &BodyOptionalityServerTransport{srv: srv}
}

// BodyOptionalityServerTransport connects instances of bodyoptionalgroup.BodyOptionalityClient to instances of BodyOptionalityServer.
// Don't use this type directly, use NewBodyOptionalityServerTransport instead.
type BodyOptionalityServerTransport struct {
	srv                                     *BodyOptionalityServer
	trMu                                    sync.Mutex
	trBodyOptionalityOptionalExplicitServer *BodyOptionalityOptionalExplicitServerTransport
}

// Do implements the policy.Transporter interface for BodyOptionalityServerTransport.
func (b *BodyOptionalityServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	if client := method[:strings.Index(method, ".")]; client != "BodyOptionalityClient" {
		return b.dispatchToClientFake(req, client)
	}
	return b.dispatchToMethodFake(req, method)
}

func (b *BodyOptionalityServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "BodyOptionalityOptionalExplicitClient":
		initServer(&b.trMu, &b.trBodyOptionalityOptionalExplicitServer, func() *BodyOptionalityOptionalExplicitServerTransport {
			return NewBodyOptionalityOptionalExplicitServerTransport(&b.srv.BodyOptionalityOptionalExplicitServer)
		})
		resp, err = b.trBodyOptionalityOptionalExplicitServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

func (b *BodyOptionalityServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "BodyOptionalityClient.RequiredExplicit":
		resp, err = b.dispatchRequiredExplicit(req)
	case "BodyOptionalityClient.RequiredImplicit":
		resp, err = b.dispatchRequiredImplicit(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (b *BodyOptionalityServerTransport) dispatchRequiredExplicit(req *http.Request) (*http.Response, error) {
	if b.srv.RequiredExplicit == nil {
		return nil, &nonRetriableError{errors.New("fake for method RequiredExplicit not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bodyoptionalgroup.BodyModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.RequiredExplicit(req.Context(), body, nil)
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

func (b *BodyOptionalityServerTransport) dispatchRequiredImplicit(req *http.Request) (*http.Response, error) {
	if b.srv.RequiredImplicit == nil {
		return nil, &nonRetriableError{errors.New("fake for method RequiredImplicit not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bodyoptionalgroup.BodyModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.RequiredImplicit(req.Context(), body, nil)
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
