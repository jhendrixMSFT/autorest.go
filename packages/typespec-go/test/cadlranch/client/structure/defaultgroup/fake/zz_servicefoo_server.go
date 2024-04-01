// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"defaultgroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ServiceFooServer is a fake server for instances of the defaultgroup.ServiceFooClient type.
type ServiceFooServer struct {
	// Four is the fake for method ServiceFooClient.Four
	// HTTP status codes to indicate success: http.StatusNoContent
	Four func(ctx context.Context, options *defaultgroup.ServiceFooClientFourOptions) (resp azfake.Responder[defaultgroup.ServiceFooClientFourResponse], errResp azfake.ErrorResponder)

	// Three is the fake for method ServiceFooClient.Three
	// HTTP status codes to indicate success: http.StatusNoContent
	Three func(ctx context.Context, options *defaultgroup.ServiceFooClientThreeOptions) (resp azfake.Responder[defaultgroup.ServiceFooClientThreeResponse], errResp azfake.ErrorResponder)
}

// NewServiceFooServerTransport creates a new instance of ServiceFooServerTransport with the provided implementation.
// The returned ServiceFooServerTransport instance is connected to an instance of defaultgroup.ServiceFooClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceFooServerTransport(srv *ServiceFooServer) *ServiceFooServerTransport {
	return &ServiceFooServerTransport{srv: srv}
}

// ServiceFooServerTransport connects instances of defaultgroup.ServiceFooClient to instances of ServiceFooServer.
// Don't use this type directly, use NewServiceFooServerTransport instead.
type ServiceFooServerTransport struct {
	srv *ServiceFooServer
}

// Do implements the policy.Transporter interface for ServiceFooServerTransport.
func (s *ServiceFooServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServiceFooServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ServiceFooClient.Four":
		resp, err = s.dispatchFour(req)
	case "ServiceFooClient.Three":
		resp, err = s.dispatchThree(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (s *ServiceFooServerTransport) dispatchFour(req *http.Request) (*http.Response, error) {
	if s.srv.Four == nil {
		return nil, &nonRetriableError{errors.New("fake for method Four not implemented")}
	}
	respr, errRespr := s.srv.Four(req.Context(), nil)
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

func (s *ServiceFooServerTransport) dispatchThree(req *http.Request) (*http.Response, error) {
	if s.srv.Three == nil {
		return nil, &nonRetriableError{errors.New("fake for method Three not implemented")}
	}
	respr, errRespr := s.srv.Three(req.Context(), nil)
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