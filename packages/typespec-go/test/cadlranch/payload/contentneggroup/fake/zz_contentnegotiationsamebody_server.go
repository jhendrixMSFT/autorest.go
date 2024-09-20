// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"contentneggroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ContentNegotiationSameBodyServer is a fake server for instances of the contentneggroup.ContentNegotiationSameBodyClient type.
type ContentNegotiationSameBodyServer struct {
	// GetAvatarAsJPEG is the fake for method ContentNegotiationSameBodyClient.GetAvatarAsJPEG
	// HTTP status codes to indicate success: http.StatusOK
	GetAvatarAsJPEG func(ctx context.Context, options *contentneggroup.ContentNegotiationSameBodyClientGetAvatarAsJPEGOptions) (resp azfake.Responder[contentneggroup.ContentNegotiationSameBodyClientGetAvatarAsJPEGResponse], errResp azfake.ErrorResponder)

	// GetAvatarAsPNG is the fake for method ContentNegotiationSameBodyClient.GetAvatarAsPNG
	// HTTP status codes to indicate success: http.StatusOK
	GetAvatarAsPNG func(ctx context.Context, options *contentneggroup.ContentNegotiationSameBodyClientGetAvatarAsPNGOptions) (resp azfake.Responder[contentneggroup.ContentNegotiationSameBodyClientGetAvatarAsPNGResponse], errResp azfake.ErrorResponder)
}

// NewContentNegotiationSameBodyServerTransport creates a new instance of ContentNegotiationSameBodyServerTransport with the provided implementation.
// The returned ContentNegotiationSameBodyServerTransport instance is connected to an instance of contentneggroup.ContentNegotiationSameBodyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContentNegotiationSameBodyServerTransport(srv *ContentNegotiationSameBodyServer) *ContentNegotiationSameBodyServerTransport {
	return &ContentNegotiationSameBodyServerTransport{srv: srv}
}

// ContentNegotiationSameBodyServerTransport connects instances of contentneggroup.ContentNegotiationSameBodyClient to instances of ContentNegotiationSameBodyServer.
// Don't use this type directly, use NewContentNegotiationSameBodyServerTransport instead.
type ContentNegotiationSameBodyServerTransport struct {
	srv *ContentNegotiationSameBodyServer
}

// Do implements the policy.Transporter interface for ContentNegotiationSameBodyServerTransport.
func (c *ContentNegotiationSameBodyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ContentNegotiationSameBodyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "ContentNegotiationSameBodyClient.GetAvatarAsJPEG":
			res.resp, res.err = c.dispatchGetAvatarAsJPEG(req)
		case "ContentNegotiationSameBodyClient.GetAvatarAsPNG":
			res.resp, res.err = c.dispatchGetAvatarAsPNG(req)
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

func (c *ContentNegotiationSameBodyServerTransport) dispatchGetAvatarAsJPEG(req *http.Request) (*http.Response, error) {
	if c.srv.GetAvatarAsJPEG == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAvatarAsJPEG not implemented")}
	}
	respr, errRespr := c.srv.GetAvatarAsJPEG(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: req.Header.Get("Content-Type"),
	})
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("content-type", "image/jpeg")
	}
	return resp, nil
}

func (c *ContentNegotiationSameBodyServerTransport) dispatchGetAvatarAsPNG(req *http.Request) (*http.Response, error) {
	if c.srv.GetAvatarAsPNG == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAvatarAsPNG not implemented")}
	}
	respr, errRespr := c.srv.GetAvatarAsPNG(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: req.Header.Get("Content-Type"),
	})
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("content-type", "image/png")
	}
	return resp, nil
}
