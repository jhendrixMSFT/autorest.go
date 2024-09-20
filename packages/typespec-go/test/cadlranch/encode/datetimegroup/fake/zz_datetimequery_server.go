// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"datetimegroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// DatetimeQueryServer is a fake server for instances of the datetimegroup.DatetimeQueryClient type.
type DatetimeQueryServer struct {
	// Default is the fake for method DatetimeQueryClient.Default
	// HTTP status codes to indicate success: http.StatusNoContent
	Default func(ctx context.Context, value time.Time, options *datetimegroup.DatetimeQueryClientDefaultOptions) (resp azfake.Responder[datetimegroup.DatetimeQueryClientDefaultResponse], errResp azfake.ErrorResponder)

	// RFC3339 is the fake for method DatetimeQueryClient.RFC3339
	// HTTP status codes to indicate success: http.StatusNoContent
	RFC3339 func(ctx context.Context, value time.Time, options *datetimegroup.DatetimeQueryClientRFC3339Options) (resp azfake.Responder[datetimegroup.DatetimeQueryClientRFC3339Response], errResp azfake.ErrorResponder)

	// RFC7231 is the fake for method DatetimeQueryClient.RFC7231
	// HTTP status codes to indicate success: http.StatusNoContent
	RFC7231 func(ctx context.Context, value time.Time, options *datetimegroup.DatetimeQueryClientRFC7231Options) (resp azfake.Responder[datetimegroup.DatetimeQueryClientRFC7231Response], errResp azfake.ErrorResponder)

	// UnixTimestamp is the fake for method DatetimeQueryClient.UnixTimestamp
	// HTTP status codes to indicate success: http.StatusNoContent
	UnixTimestamp func(ctx context.Context, value time.Time, options *datetimegroup.DatetimeQueryClientUnixTimestampOptions) (resp azfake.Responder[datetimegroup.DatetimeQueryClientUnixTimestampResponse], errResp azfake.ErrorResponder)

	// UnixTimestampArray is the fake for method DatetimeQueryClient.UnixTimestampArray
	// HTTP status codes to indicate success: http.StatusNoContent
	UnixTimestampArray func(ctx context.Context, value []time.Time, options *datetimegroup.DatetimeQueryClientUnixTimestampArrayOptions) (resp azfake.Responder[datetimegroup.DatetimeQueryClientUnixTimestampArrayResponse], errResp azfake.ErrorResponder)
}

// NewDatetimeQueryServerTransport creates a new instance of DatetimeQueryServerTransport with the provided implementation.
// The returned DatetimeQueryServerTransport instance is connected to an instance of datetimegroup.DatetimeQueryClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatetimeQueryServerTransport(srv *DatetimeQueryServer) *DatetimeQueryServerTransport {
	return &DatetimeQueryServerTransport{srv: srv}
}

// DatetimeQueryServerTransport connects instances of datetimegroup.DatetimeQueryClient to instances of DatetimeQueryServer.
// Don't use this type directly, use NewDatetimeQueryServerTransport instead.
type DatetimeQueryServerTransport struct {
	srv *DatetimeQueryServer
}

// Do implements the policy.Transporter interface for DatetimeQueryServerTransport.
func (d *DatetimeQueryServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DatetimeQueryServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "DatetimeQueryClient.Default":
			res.resp, res.err = d.dispatchDefault(req)
		case "DatetimeQueryClient.RFC3339":
			res.resp, res.err = d.dispatchRFC3339(req)
		case "DatetimeQueryClient.RFC7231":
			res.resp, res.err = d.dispatchRFC7231(req)
		case "DatetimeQueryClient.UnixTimestamp":
			res.resp, res.err = d.dispatchUnixTimestamp(req)
		case "DatetimeQueryClient.UnixTimestampArray":
			res.resp, res.err = d.dispatchUnixTimestampArray(req)
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

func (d *DatetimeQueryServerTransport) dispatchDefault(req *http.Request) (*http.Response, error) {
	if d.srv.Default == nil {
		return nil, &nonRetriableError{errors.New("fake for method Default not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueParam, err := time.Parse(time.RFC3339Nano, valueUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Default(req.Context(), valueParam, nil)
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

func (d *DatetimeQueryServerTransport) dispatchRFC3339(req *http.Request) (*http.Response, error) {
	if d.srv.RFC3339 == nil {
		return nil, &nonRetriableError{errors.New("fake for method RFC3339 not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueParam, err := time.Parse(time.RFC3339Nano, valueUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.RFC3339(req.Context(), valueParam, nil)
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

func (d *DatetimeQueryServerTransport) dispatchRFC7231(req *http.Request) (*http.Response, error) {
	if d.srv.RFC7231 == nil {
		return nil, &nonRetriableError{errors.New("fake for method RFC7231 not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueParam, err := time.Parse(time.RFC1123, valueUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.RFC7231(req.Context(), valueParam, nil)
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

func (d *DatetimeQueryServerTransport) dispatchUnixTimestamp(req *http.Request) (*http.Response, error) {
	if d.srv.UnixTimestamp == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnixTimestamp not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueParam, err := parseWithCast(valueUnescaped, func(v string) (time.Time, error) {
		p, parseErr := strconv.ParseInt(v, 10, 64)
		if parseErr != nil {
			return time.Time{}, parseErr
		}
		return time.Unix(p, 0).UTC(), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.UnixTimestamp(req.Context(), valueParam, nil)
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

func (d *DatetimeQueryServerTransport) dispatchUnixTimestampArray(req *http.Request) (*http.Response, error) {
	if d.srv.UnixTimestampArray == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnixTimestampArray not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueElements := splitHelper(valueUnescaped, ",")
	valueParam := make([]time.Time, len(valueElements))
	for i := 0; i < len(valueElements); i++ {
		p, parseErr := strconv.ParseInt(valueElements[i], 10, 64)
		if parseErr != nil {
			return nil, parseErr
		}
		parsedTimeUnix := time.Unix(p, 0).UTC()
		valueParam[i] = time.Time(parsedTimeUnix)
	}
	respr, errRespr := d.srv.UnixTimestampArray(req.Context(), valueParam, nil)
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
