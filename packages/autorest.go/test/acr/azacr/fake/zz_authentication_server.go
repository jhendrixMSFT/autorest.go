//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"azacr"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"reflect"
)

// AuthenticationServer is a fake server for instances of the azacr.AuthenticationClient type.
type AuthenticationServer struct {
	// ExchangeAADAccessTokenForAcrRefreshToken is the fake for method AuthenticationClient.ExchangeAADAccessTokenForAcrRefreshToken
	// HTTP status codes to indicate success: http.StatusOK
	ExchangeAADAccessTokenForAcrRefreshToken func(ctx context.Context, grantType azacr.PostContentSchemaGrantType, service string, options *azacr.AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenOptions) (resp azfake.Responder[azacr.AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenResponse], errResp azfake.ErrorResponder)

	// ExchangeAcrRefreshTokenForAcrAccessToken is the fake for method AuthenticationClient.ExchangeAcrRefreshTokenForAcrAccessToken
	// HTTP status codes to indicate success: http.StatusOK
	ExchangeAcrRefreshTokenForAcrAccessToken func(ctx context.Context, service string, scope string, refreshToken string, options *azacr.AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenOptions) (resp azfake.Responder[azacr.AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenResponse], errResp azfake.ErrorResponder)

	// GetAcrAccessTokenFromLogin is the fake for method AuthenticationClient.GetAcrAccessTokenFromLogin
	// HTTP status codes to indicate success: http.StatusOK
	GetAcrAccessTokenFromLogin func(ctx context.Context, service string, scope string, options *azacr.AuthenticationClientGetAcrAccessTokenFromLoginOptions) (resp azfake.Responder[azacr.AuthenticationClientGetAcrAccessTokenFromLoginResponse], errResp azfake.ErrorResponder)
}

// NewAuthenticationServerTransport creates a new instance of AuthenticationServerTransport with the provided implementation.
// The returned AuthenticationServerTransport instance is connected to an instance of azacr.AuthenticationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAuthenticationServerTransport(srv *AuthenticationServer) *AuthenticationServerTransport {
	return &AuthenticationServerTransport{srv: srv}
}

// AuthenticationServerTransport connects instances of azacr.AuthenticationClient to instances of AuthenticationServer.
// Don't use this type directly, use NewAuthenticationServerTransport instead.
type AuthenticationServerTransport struct {
	srv *AuthenticationServer
}

// Do implements the policy.Transporter interface for AuthenticationServerTransport.
func (a *AuthenticationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AuthenticationClient.ExchangeAADAccessTokenForAcrRefreshToken":
		resp, err = a.dispatchExchangeAADAccessTokenForAcrRefreshToken(req)
	case "AuthenticationClient.ExchangeAcrRefreshTokenForAcrAccessToken":
		resp, err = a.dispatchExchangeAcrRefreshTokenForAcrAccessToken(req)
	case "AuthenticationClient.GetAcrAccessTokenFromLogin":
		resp, err = a.dispatchGetAcrAccessTokenFromLogin(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AuthenticationServerTransport) dispatchExchangeAADAccessTokenForAcrRefreshToken(req *http.Request) (*http.Response, error) {
	if a.srv.ExchangeAADAccessTokenForAcrRefreshToken == nil {
		return nil, &nonRetriableError{errors.New("fake for method ExchangeAADAccessTokenForAcrRefreshToken not implemented")}
	}
	var grantType azacr.PostContentSchemaGrantType
	var service string
	var Tenant string
	var RefreshToken string
	var AccessToken string
	if err := req.ParseForm(); err != nil {
		return nil, &nonRetriableError{fmt.Errorf("failed parsing form data: %v", err)}
	}
	for key := range req.Form {
		switch key {
		case "grant_type":
			grantType = azacr.PostContentSchemaGrantType(req.FormValue(key))
		case "service":
			service = req.FormValue(key)
		case "tenant":
			Tenant = req.FormValue(key)
		case "refresh_token":
			RefreshToken = req.FormValue(key)
		case "access_token":
			AccessToken = req.FormValue(key)
		}
	}
	var options *azacr.AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenOptions
	if !reflect.ValueOf(Tenant).IsZero() || !reflect.ValueOf(RefreshToken).IsZero() || !reflect.ValueOf(AccessToken).IsZero() {
		options = &azacr.AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenOptions{
			Tenant:       &Tenant,
			RefreshToken: &RefreshToken,
			AccessToken:  &AccessToken,
		}
	}
	respr, errRespr := a.srv.ExchangeAADAccessTokenForAcrRefreshToken(req.Context(), grantType, service, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RefreshToken, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AuthenticationServerTransport) dispatchExchangeAcrRefreshTokenForAcrAccessToken(req *http.Request) (*http.Response, error) {
	if a.srv.ExchangeAcrRefreshTokenForAcrAccessToken == nil {
		return nil, &nonRetriableError{errors.New("fake for method ExchangeAcrRefreshTokenForAcrAccessToken not implemented")}
	}
	var service string
	var scope string
	var refreshToken string
	var grantType azacr.TokenGrantType
	if err := req.ParseForm(); err != nil {
		return nil, &nonRetriableError{fmt.Errorf("failed parsing form data: %v", err)}
	}
	for key := range req.Form {
		switch key {
		case "service":
			service = req.FormValue(key)
		case "scope":
			scope = req.FormValue(key)
		case "refresh_token":
			refreshToken = req.FormValue(key)
		case "grant_type":
			grantType = azacr.TokenGrantType(req.FormValue(key))
		}
	}
	var options *azacr.AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenOptions
	if !reflect.ValueOf(grantType).IsZero() {
		options = &azacr.AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenOptions{
			GrantType: &grantType,
		}
	}
	respr, errRespr := a.srv.ExchangeAcrRefreshTokenForAcrAccessToken(req.Context(), service, scope, refreshToken, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessToken, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AuthenticationServerTransport) dispatchGetAcrAccessTokenFromLogin(req *http.Request) (*http.Response, error) {
	if a.srv.GetAcrAccessTokenFromLogin == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAcrAccessTokenFromLogin not implemented")}
	}
	qp := req.URL.Query()
	serviceUnescaped, err := url.QueryUnescape(qp.Get("service"))
	if err != nil {
		return nil, err
	}
	scopeUnescaped, err := url.QueryUnescape(qp.Get("scope"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetAcrAccessTokenFromLogin(req.Context(), serviceUnescaped, scopeUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessToken, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
