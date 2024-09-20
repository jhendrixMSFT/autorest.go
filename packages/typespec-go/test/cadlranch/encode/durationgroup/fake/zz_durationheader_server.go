// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"durationgroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

// DurationHeaderServer is a fake server for instances of the durationgroup.DurationHeaderClient type.
type DurationHeaderServer struct {
	// Default is the fake for method DurationHeaderClient.Default
	// HTTP status codes to indicate success: http.StatusNoContent
	Default func(ctx context.Context, duration string, options *durationgroup.DurationHeaderClientDefaultOptions) (resp azfake.Responder[durationgroup.DurationHeaderClientDefaultResponse], errResp azfake.ErrorResponder)

	// Float64Seconds is the fake for method DurationHeaderClient.Float64Seconds
	// HTTP status codes to indicate success: http.StatusNoContent
	Float64Seconds func(ctx context.Context, duration float64, options *durationgroup.DurationHeaderClientFloat64SecondsOptions) (resp azfake.Responder[durationgroup.DurationHeaderClientFloat64SecondsResponse], errResp azfake.ErrorResponder)

	// FloatSeconds is the fake for method DurationHeaderClient.FloatSeconds
	// HTTP status codes to indicate success: http.StatusNoContent
	FloatSeconds func(ctx context.Context, duration float32, options *durationgroup.DurationHeaderClientFloatSecondsOptions) (resp azfake.Responder[durationgroup.DurationHeaderClientFloatSecondsResponse], errResp azfake.ErrorResponder)

	// ISO8601 is the fake for method DurationHeaderClient.ISO8601
	// HTTP status codes to indicate success: http.StatusNoContent
	ISO8601 func(ctx context.Context, duration string, options *durationgroup.DurationHeaderClientISO8601Options) (resp azfake.Responder[durationgroup.DurationHeaderClientISO8601Response], errResp azfake.ErrorResponder)

	// ISO8601Array is the fake for method DurationHeaderClient.ISO8601Array
	// HTTP status codes to indicate success: http.StatusNoContent
	ISO8601Array func(ctx context.Context, duration []string, options *durationgroup.DurationHeaderClientISO8601ArrayOptions) (resp azfake.Responder[durationgroup.DurationHeaderClientISO8601ArrayResponse], errResp azfake.ErrorResponder)

	// Int32Seconds is the fake for method DurationHeaderClient.Int32Seconds
	// HTTP status codes to indicate success: http.StatusNoContent
	Int32Seconds func(ctx context.Context, duration int32, options *durationgroup.DurationHeaderClientInt32SecondsOptions) (resp azfake.Responder[durationgroup.DurationHeaderClientInt32SecondsResponse], errResp azfake.ErrorResponder)
}

// NewDurationHeaderServerTransport creates a new instance of DurationHeaderServerTransport with the provided implementation.
// The returned DurationHeaderServerTransport instance is connected to an instance of durationgroup.DurationHeaderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDurationHeaderServerTransport(srv *DurationHeaderServer) *DurationHeaderServerTransport {
	return &DurationHeaderServerTransport{srv: srv}
}

// DurationHeaderServerTransport connects instances of durationgroup.DurationHeaderClient to instances of DurationHeaderServer.
// Don't use this type directly, use NewDurationHeaderServerTransport instead.
type DurationHeaderServerTransport struct {
	srv *DurationHeaderServer
}

// Do implements the policy.Transporter interface for DurationHeaderServerTransport.
func (d *DurationHeaderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DurationHeaderServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "DurationHeaderClient.Default":
			res.resp, res.err = d.dispatchDefault(req)
		case "DurationHeaderClient.Float64Seconds":
			res.resp, res.err = d.dispatchFloat64Seconds(req)
		case "DurationHeaderClient.FloatSeconds":
			res.resp, res.err = d.dispatchFloatSeconds(req)
		case "DurationHeaderClient.ISO8601":
			res.resp, res.err = d.dispatchISO8601(req)
		case "DurationHeaderClient.ISO8601Array":
			res.resp, res.err = d.dispatchISO8601Array(req)
		case "DurationHeaderClient.Int32Seconds":
			res.resp, res.err = d.dispatchInt32Seconds(req)
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

func (d *DurationHeaderServerTransport) dispatchDefault(req *http.Request) (*http.Response, error) {
	if d.srv.Default == nil {
		return nil, &nonRetriableError{errors.New("fake for method Default not implemented")}
	}
	respr, errRespr := d.srv.Default(req.Context(), getHeaderValue(req.Header, "duration"), nil)
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

func (d *DurationHeaderServerTransport) dispatchFloat64Seconds(req *http.Request) (*http.Response, error) {
	if d.srv.Float64Seconds == nil {
		return nil, &nonRetriableError{errors.New("fake for method Float64Seconds not implemented")}
	}
	durationParam, err := strconv.ParseFloat(getHeaderValue(req.Header, "duration"), 64)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Float64Seconds(req.Context(), durationParam, nil)
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

func (d *DurationHeaderServerTransport) dispatchFloatSeconds(req *http.Request) (*http.Response, error) {
	if d.srv.FloatSeconds == nil {
		return nil, &nonRetriableError{errors.New("fake for method FloatSeconds not implemented")}
	}
	durationParam, err := parseWithCast(getHeaderValue(req.Header, "duration"), func(v string) (float32, error) {
		p, parseErr := strconv.ParseFloat(v, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return float32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.FloatSeconds(req.Context(), durationParam, nil)
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

func (d *DurationHeaderServerTransport) dispatchISO8601(req *http.Request) (*http.Response, error) {
	if d.srv.ISO8601 == nil {
		return nil, &nonRetriableError{errors.New("fake for method ISO8601 not implemented")}
	}
	respr, errRespr := d.srv.ISO8601(req.Context(), getHeaderValue(req.Header, "duration"), nil)
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

func (d *DurationHeaderServerTransport) dispatchISO8601Array(req *http.Request) (*http.Response, error) {
	if d.srv.ISO8601Array == nil {
		return nil, &nonRetriableError{errors.New("fake for method ISO8601Array not implemented")}
	}
	respr, errRespr := d.srv.ISO8601Array(req.Context(), splitHelper(getHeaderValue(req.Header, "duration"), ","), nil)
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

func (d *DurationHeaderServerTransport) dispatchInt32Seconds(req *http.Request) (*http.Response, error) {
	if d.srv.Int32Seconds == nil {
		return nil, &nonRetriableError{errors.New("fake for method Int32Seconds not implemented")}
	}
	durationParam, err := parseWithCast(getHeaderValue(req.Header, "duration"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Int32Seconds(req.Context(), durationParam, nil)
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
