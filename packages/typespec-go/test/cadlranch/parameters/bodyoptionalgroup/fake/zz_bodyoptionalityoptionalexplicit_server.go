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
	"reflect"
)

// BodyOptionalityOptionalExplicitServer is a fake server for instances of the bodyoptionalgroup.BodyOptionalityOptionalExplicitClient type.
type BodyOptionalityOptionalExplicitServer struct {
	// Omit is the fake for method BodyOptionalityOptionalExplicitClient.Omit
	// HTTP status codes to indicate success: http.StatusNoContent
	Omit func(ctx context.Context, options *bodyoptionalgroup.BodyOptionalityOptionalExplicitClientOmitOptions) (resp azfake.Responder[bodyoptionalgroup.BodyOptionalityOptionalExplicitClientOmitResponse], errResp azfake.ErrorResponder)

	// Set is the fake for method BodyOptionalityOptionalExplicitClient.Set
	// HTTP status codes to indicate success: http.StatusNoContent
	Set func(ctx context.Context, options *bodyoptionalgroup.BodyOptionalityOptionalExplicitClientSetOptions) (resp azfake.Responder[bodyoptionalgroup.BodyOptionalityOptionalExplicitClientSetResponse], errResp azfake.ErrorResponder)
}

// NewBodyOptionalityOptionalExplicitServerTransport creates a new instance of BodyOptionalityOptionalExplicitServerTransport with the provided implementation.
// The returned BodyOptionalityOptionalExplicitServerTransport instance is connected to an instance of bodyoptionalgroup.BodyOptionalityOptionalExplicitClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBodyOptionalityOptionalExplicitServerTransport(srv *BodyOptionalityOptionalExplicitServer) *BodyOptionalityOptionalExplicitServerTransport {
	return &BodyOptionalityOptionalExplicitServerTransport{srv: srv}
}

// BodyOptionalityOptionalExplicitServerTransport connects instances of bodyoptionalgroup.BodyOptionalityOptionalExplicitClient to instances of BodyOptionalityOptionalExplicitServer.
// Don't use this type directly, use NewBodyOptionalityOptionalExplicitServerTransport instead.
type BodyOptionalityOptionalExplicitServerTransport struct {
	srv *BodyOptionalityOptionalExplicitServer
}

// Do implements the policy.Transporter interface for BodyOptionalityOptionalExplicitServerTransport.
func (b *BodyOptionalityOptionalExplicitServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BodyOptionalityOptionalExplicitServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "BodyOptionalityOptionalExplicitClient.Omit":
			res.resp, res.err = b.dispatchOmit(req)
		case "BodyOptionalityOptionalExplicitClient.Set":
			res.resp, res.err = b.dispatchSet(req)
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

func (b *BodyOptionalityOptionalExplicitServerTransport) dispatchOmit(req *http.Request) (*http.Response, error) {
	if b.srv.Omit == nil {
		return nil, &nonRetriableError{errors.New("fake for method Omit not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bodyoptionalgroup.BodyModel](req)
	if err != nil {
		return nil, err
	}
	contentTypeParam := getOptional(getHeaderValue(req.Header, "Content-Type"))
	var options *bodyoptionalgroup.BodyOptionalityOptionalExplicitClientOmitOptions
	if !reflect.ValueOf(body).IsZero() || contentTypeParam != nil {
		options = &bodyoptionalgroup.BodyOptionalityOptionalExplicitClientOmitOptions{
			Body:        &body,
			ContentType: contentTypeParam,
		}
	}
	respr, errRespr := b.srv.Omit(req.Context(), options)
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

func (b *BodyOptionalityOptionalExplicitServerTransport) dispatchSet(req *http.Request) (*http.Response, error) {
	if b.srv.Set == nil {
		return nil, &nonRetriableError{errors.New("fake for method Set not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bodyoptionalgroup.BodyModel](req)
	if err != nil {
		return nil, err
	}
	contentTypeParam := getOptional(getHeaderValue(req.Header, "Content-Type"))
	var options *bodyoptionalgroup.BodyOptionalityOptionalExplicitClientSetOptions
	if !reflect.ValueOf(body).IsZero() || contentTypeParam != nil {
		options = &bodyoptionalgroup.BodyOptionalityOptionalExplicitClientSetOptions{
			Body:        &body,
			ContentType: contentTypeParam,
		}
	}
	respr, errRespr := b.srv.Set(req.Context(), options)
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
