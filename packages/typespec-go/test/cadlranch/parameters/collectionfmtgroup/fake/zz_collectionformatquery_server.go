// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"collectionfmtgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
)

// CollectionFormatQueryServer is a fake server for instances of the collectionfmtgroup.CollectionFormatQueryClient type.
type CollectionFormatQueryServer struct {
	// CSV is the fake for method CollectionFormatQueryClient.CSV
	// HTTP status codes to indicate success: http.StatusNoContent
	CSV func(ctx context.Context, colors []string, options *collectionfmtgroup.CollectionFormatQueryClientCSVOptions) (resp azfake.Responder[collectionfmtgroup.CollectionFormatQueryClientCSVResponse], errResp azfake.ErrorResponder)

	// Multi is the fake for method CollectionFormatQueryClient.Multi
	// HTTP status codes to indicate success: http.StatusNoContent
	Multi func(ctx context.Context, colors []string, options *collectionfmtgroup.CollectionFormatQueryClientMultiOptions) (resp azfake.Responder[collectionfmtgroup.CollectionFormatQueryClientMultiResponse], errResp azfake.ErrorResponder)

	// Pipes is the fake for method CollectionFormatQueryClient.Pipes
	// HTTP status codes to indicate success: http.StatusNoContent
	Pipes func(ctx context.Context, colors []string, options *collectionfmtgroup.CollectionFormatQueryClientPipesOptions) (resp azfake.Responder[collectionfmtgroup.CollectionFormatQueryClientPipesResponse], errResp azfake.ErrorResponder)

	// Ssv is the fake for method CollectionFormatQueryClient.Ssv
	// HTTP status codes to indicate success: http.StatusNoContent
	Ssv func(ctx context.Context, colors []string, options *collectionfmtgroup.CollectionFormatQueryClientSsvOptions) (resp azfake.Responder[collectionfmtgroup.CollectionFormatQueryClientSsvResponse], errResp azfake.ErrorResponder)

	// Tsv is the fake for method CollectionFormatQueryClient.Tsv
	// HTTP status codes to indicate success: http.StatusNoContent
	Tsv func(ctx context.Context, colors []string, options *collectionfmtgroup.CollectionFormatQueryClientTsvOptions) (resp azfake.Responder[collectionfmtgroup.CollectionFormatQueryClientTsvResponse], errResp azfake.ErrorResponder)
}

// NewCollectionFormatQueryServerTransport creates a new instance of CollectionFormatQueryServerTransport with the provided implementation.
// The returned CollectionFormatQueryServerTransport instance is connected to an instance of collectionfmtgroup.CollectionFormatQueryClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCollectionFormatQueryServerTransport(srv *CollectionFormatQueryServer) *CollectionFormatQueryServerTransport {
	return &CollectionFormatQueryServerTransport{srv: srv}
}

// CollectionFormatQueryServerTransport connects instances of collectionfmtgroup.CollectionFormatQueryClient to instances of CollectionFormatQueryServer.
// Don't use this type directly, use NewCollectionFormatQueryServerTransport instead.
type CollectionFormatQueryServerTransport struct {
	srv *CollectionFormatQueryServer
}

// Do implements the policy.Transporter interface for CollectionFormatQueryServerTransport.
func (c *CollectionFormatQueryServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CollectionFormatQueryServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "CollectionFormatQueryClient.CSV":
			res.resp, res.err = c.dispatchCSV(req)
		case "CollectionFormatQueryClient.Multi":
			res.resp, res.err = c.dispatchMulti(req)
		case "CollectionFormatQueryClient.Pipes":
			res.resp, res.err = c.dispatchPipes(req)
		case "CollectionFormatQueryClient.Ssv":
			res.resp, res.err = c.dispatchSsv(req)
		case "CollectionFormatQueryClient.Tsv":
			res.resp, res.err = c.dispatchTsv(req)
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

func (c *CollectionFormatQueryServerTransport) dispatchCSV(req *http.Request) (*http.Response, error) {
	if c.srv.CSV == nil {
		return nil, &nonRetriableError{errors.New("fake for method CSV not implemented")}
	}
	qp := req.URL.Query()
	colorsParam, err := url.QueryUnescape(qp.Get("colors"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CSV(req.Context(), splitHelper(colorsParam, ","), nil)
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

func (c *CollectionFormatQueryServerTransport) dispatchMulti(req *http.Request) (*http.Response, error) {
	if c.srv.Multi == nil {
		return nil, &nonRetriableError{errors.New("fake for method Multi not implemented")}
	}
	qp := req.URL.Query()
	colorsEscaped := qp["colors"]
	colorsParam := make([]string, len(colorsEscaped))
	for i, v := range colorsEscaped {
		u, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return nil, unescapeErr
		}
		colorsParam[i] = u
	}
	respr, errRespr := c.srv.Multi(req.Context(), colorsParam, nil)
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

func (c *CollectionFormatQueryServerTransport) dispatchPipes(req *http.Request) (*http.Response, error) {
	if c.srv.Pipes == nil {
		return nil, &nonRetriableError{errors.New("fake for method Pipes not implemented")}
	}
	qp := req.URL.Query()
	colorsParam, err := url.QueryUnescape(qp.Get("colors"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Pipes(req.Context(), splitHelper(colorsParam, "|"), nil)
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

func (c *CollectionFormatQueryServerTransport) dispatchSsv(req *http.Request) (*http.Response, error) {
	if c.srv.Ssv == nil {
		return nil, &nonRetriableError{errors.New("fake for method Ssv not implemented")}
	}
	qp := req.URL.Query()
	colorsParam, err := url.QueryUnescape(qp.Get("colors"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Ssv(req.Context(), splitHelper(colorsParam, " "), nil)
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

func (c *CollectionFormatQueryServerTransport) dispatchTsv(req *http.Request) (*http.Response, error) {
	if c.srv.Tsv == nil {
		return nil, &nonRetriableError{errors.New("fake for method Tsv not implemented")}
	}
	qp := req.URL.Query()
	colorsParam, err := url.QueryUnescape(qp.Get("colors"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Tsv(req.Context(), splitHelper(colorsParam, "\t"), nil)
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
