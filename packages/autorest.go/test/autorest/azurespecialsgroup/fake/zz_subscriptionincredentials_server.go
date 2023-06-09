//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/azurespecialsgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
)

// SubscriptionInCredentialsServer is a fake server for instances of the azurespecialsgroup.SubscriptionInCredentialsClient type.
type SubscriptionInCredentialsServer struct {
	// PostMethodGlobalNotProvidedValid is the fake for method SubscriptionInCredentialsClient.PostMethodGlobalNotProvidedValid
	// HTTP status codes to indicate success: http.StatusOK
	PostMethodGlobalNotProvidedValid func(ctx context.Context, options *azurespecialsgroup.SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidOptions) (resp azfake.Responder[azurespecialsgroup.SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidResponse], errResp azfake.ErrorResponder)

	// PostMethodGlobalNull is the fake for method SubscriptionInCredentialsClient.PostMethodGlobalNull
	// HTTP status codes to indicate success: http.StatusOK
	PostMethodGlobalNull func(ctx context.Context, options *azurespecialsgroup.SubscriptionInCredentialsClientPostMethodGlobalNullOptions) (resp azfake.Responder[azurespecialsgroup.SubscriptionInCredentialsClientPostMethodGlobalNullResponse], errResp azfake.ErrorResponder)

	// PostMethodGlobalValid is the fake for method SubscriptionInCredentialsClient.PostMethodGlobalValid
	// HTTP status codes to indicate success: http.StatusOK
	PostMethodGlobalValid func(ctx context.Context, options *azurespecialsgroup.SubscriptionInCredentialsClientPostMethodGlobalValidOptions) (resp azfake.Responder[azurespecialsgroup.SubscriptionInCredentialsClientPostMethodGlobalValidResponse], errResp azfake.ErrorResponder)

	// PostPathGlobalValid is the fake for method SubscriptionInCredentialsClient.PostPathGlobalValid
	// HTTP status codes to indicate success: http.StatusOK
	PostPathGlobalValid func(ctx context.Context, options *azurespecialsgroup.SubscriptionInCredentialsClientPostPathGlobalValidOptions) (resp azfake.Responder[azurespecialsgroup.SubscriptionInCredentialsClientPostPathGlobalValidResponse], errResp azfake.ErrorResponder)

	// PostSwaggerGlobalValid is the fake for method SubscriptionInCredentialsClient.PostSwaggerGlobalValid
	// HTTP status codes to indicate success: http.StatusOK
	PostSwaggerGlobalValid func(ctx context.Context, options *azurespecialsgroup.SubscriptionInCredentialsClientPostSwaggerGlobalValidOptions) (resp azfake.Responder[azurespecialsgroup.SubscriptionInCredentialsClientPostSwaggerGlobalValidResponse], errResp azfake.ErrorResponder)
}

// NewSubscriptionInCredentialsServerTransport creates a new instance of SubscriptionInCredentialsServerTransport with the provided implementation.
// The returned SubscriptionInCredentialsServerTransport instance is connected to an instance of azurespecialsgroup.SubscriptionInCredentialsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSubscriptionInCredentialsServerTransport(srv *SubscriptionInCredentialsServer) *SubscriptionInCredentialsServerTransport {
	return &SubscriptionInCredentialsServerTransport{srv: srv}
}

// SubscriptionInCredentialsServerTransport connects instances of azurespecialsgroup.SubscriptionInCredentialsClient to instances of SubscriptionInCredentialsServer.
// Don't use this type directly, use NewSubscriptionInCredentialsServerTransport instead.
type SubscriptionInCredentialsServerTransport struct {
	srv *SubscriptionInCredentialsServer
}

// Do implements the policy.Transporter interface for SubscriptionInCredentialsServerTransport.
func (s *SubscriptionInCredentialsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SubscriptionInCredentialsClient.PostMethodGlobalNotProvidedValid":
		resp, err = s.dispatchPostMethodGlobalNotProvidedValid(req)
	case "SubscriptionInCredentialsClient.PostMethodGlobalNull":
		resp, err = s.dispatchPostMethodGlobalNull(req)
	case "SubscriptionInCredentialsClient.PostMethodGlobalValid":
		resp, err = s.dispatchPostMethodGlobalValid(req)
	case "SubscriptionInCredentialsClient.PostPathGlobalValid":
		resp, err = s.dispatchPostPathGlobalValid(req)
	case "SubscriptionInCredentialsClient.PostSwaggerGlobalValid":
		resp, err = s.dispatchPostSwaggerGlobalValid(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SubscriptionInCredentialsServerTransport) dispatchPostMethodGlobalNotProvidedValid(req *http.Request) (*http.Response, error) {
	if s.srv.PostMethodGlobalNotProvidedValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostMethodGlobalNotProvidedValid not implemented")}
	}
	const regexStr = `/azurespecials/subscriptionId/method/string/none/path/globalNotProvided/1234-5678-9012-3456/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.PostMethodGlobalNotProvidedValid(req.Context(), nil)
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

func (s *SubscriptionInCredentialsServerTransport) dispatchPostMethodGlobalNull(req *http.Request) (*http.Response, error) {
	if s.srv.PostMethodGlobalNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostMethodGlobalNull not implemented")}
	}
	const regexStr = `/azurespecials/subscriptionId/method/string/none/path/global/null/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.PostMethodGlobalNull(req.Context(), nil)
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

func (s *SubscriptionInCredentialsServerTransport) dispatchPostMethodGlobalValid(req *http.Request) (*http.Response, error) {
	if s.srv.PostMethodGlobalValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostMethodGlobalValid not implemented")}
	}
	const regexStr = `/azurespecials/subscriptionId/method/string/none/path/global/1234-5678-9012-3456/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.PostMethodGlobalValid(req.Context(), nil)
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

func (s *SubscriptionInCredentialsServerTransport) dispatchPostPathGlobalValid(req *http.Request) (*http.Response, error) {
	if s.srv.PostPathGlobalValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostPathGlobalValid not implemented")}
	}
	const regexStr = `/azurespecials/subscriptionId/path/string/none/path/global/1234-5678-9012-3456/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.PostPathGlobalValid(req.Context(), nil)
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

func (s *SubscriptionInCredentialsServerTransport) dispatchPostSwaggerGlobalValid(req *http.Request) (*http.Response, error) {
	if s.srv.PostSwaggerGlobalValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostSwaggerGlobalValid not implemented")}
	}
	const regexStr = `/azurespecials/subscriptionId/swagger/string/none/path/global/1234-5678-9012-3456/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.PostSwaggerGlobalValid(req.Context(), nil)
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
