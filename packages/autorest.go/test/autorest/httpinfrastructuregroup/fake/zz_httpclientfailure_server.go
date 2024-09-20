// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/httpinfrastructuregroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// HTTPClientFailureServer is a fake server for instances of the httpinfrastructuregroup.HTTPClientFailureClient type.
type HTTPClientFailureServer struct {
	// Delete400 is the fake for method HTTPClientFailureClient.Delete400
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Delete400 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientDelete400Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientDelete400Response], errResp azfake.ErrorResponder)

	// Delete407 is the fake for method HTTPClientFailureClient.Delete407
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Delete407 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientDelete407Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientDelete407Response], errResp azfake.ErrorResponder)

	// Delete417 is the fake for method HTTPClientFailureClient.Delete417
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Delete417 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientDelete417Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientDelete417Response], errResp azfake.ErrorResponder)

	// Get400 is the fake for method HTTPClientFailureClient.Get400
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Get400 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientGet400Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientGet400Response], errResp azfake.ErrorResponder)

	// Get402 is the fake for method HTTPClientFailureClient.Get402
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Get402 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientGet402Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientGet402Response], errResp azfake.ErrorResponder)

	// Get403 is the fake for method HTTPClientFailureClient.Get403
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Get403 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientGet403Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientGet403Response], errResp azfake.ErrorResponder)

	// Get411 is the fake for method HTTPClientFailureClient.Get411
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Get411 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientGet411Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientGet411Response], errResp azfake.ErrorResponder)

	// Get412 is the fake for method HTTPClientFailureClient.Get412
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Get412 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientGet412Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientGet412Response], errResp azfake.ErrorResponder)

	// Get416 is the fake for method HTTPClientFailureClient.Get416
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Get416 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientGet416Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientGet416Response], errResp azfake.ErrorResponder)

	// Head400 is the fake for method HTTPClientFailureClient.Head400
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Head400 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientHead400Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientHead400Response], errResp azfake.ErrorResponder)

	// Head401 is the fake for method HTTPClientFailureClient.Head401
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Head401 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientHead401Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientHead401Response], errResp azfake.ErrorResponder)

	// Head410 is the fake for method HTTPClientFailureClient.Head410
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Head410 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientHead410Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientHead410Response], errResp azfake.ErrorResponder)

	// Head429 is the fake for method HTTPClientFailureClient.Head429
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Head429 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientHead429Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientHead429Response], errResp azfake.ErrorResponder)

	// Options400 is the fake for method HTTPClientFailureClient.Options400
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Options400 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientOptions400Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientOptions400Response], errResp azfake.ErrorResponder)

	// Options403 is the fake for method HTTPClientFailureClient.Options403
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Options403 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientOptions403Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientOptions403Response], errResp azfake.ErrorResponder)

	// Options412 is the fake for method HTTPClientFailureClient.Options412
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Options412 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientOptions412Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientOptions412Response], errResp azfake.ErrorResponder)

	// Patch400 is the fake for method HTTPClientFailureClient.Patch400
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Patch400 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPatch400Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPatch400Response], errResp azfake.ErrorResponder)

	// Patch405 is the fake for method HTTPClientFailureClient.Patch405
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Patch405 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPatch405Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPatch405Response], errResp azfake.ErrorResponder)

	// Patch414 is the fake for method HTTPClientFailureClient.Patch414
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Patch414 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPatch414Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPatch414Response], errResp azfake.ErrorResponder)

	// Post400 is the fake for method HTTPClientFailureClient.Post400
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Post400 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPost400Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPost400Response], errResp azfake.ErrorResponder)

	// Post406 is the fake for method HTTPClientFailureClient.Post406
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Post406 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPost406Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPost406Response], errResp azfake.ErrorResponder)

	// Post415 is the fake for method HTTPClientFailureClient.Post415
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Post415 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPost415Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPost415Response], errResp azfake.ErrorResponder)

	// Put400 is the fake for method HTTPClientFailureClient.Put400
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Put400 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPut400Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPut400Response], errResp azfake.ErrorResponder)

	// Put404 is the fake for method HTTPClientFailureClient.Put404
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Put404 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPut404Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPut404Response], errResp azfake.ErrorResponder)

	// Put409 is the fake for method HTTPClientFailureClient.Put409
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Put409 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPut409Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPut409Response], errResp azfake.ErrorResponder)

	// Put413 is the fake for method HTTPClientFailureClient.Put413
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Put413 func(ctx context.Context, options *httpinfrastructuregroup.HTTPClientFailureClientPut413Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPClientFailureClientPut413Response], errResp azfake.ErrorResponder)
}

// NewHTTPClientFailureServerTransport creates a new instance of HTTPClientFailureServerTransport with the provided implementation.
// The returned HTTPClientFailureServerTransport instance is connected to an instance of httpinfrastructuregroup.HTTPClientFailureClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHTTPClientFailureServerTransport(srv *HTTPClientFailureServer) *HTTPClientFailureServerTransport {
	return &HTTPClientFailureServerTransport{srv: srv}
}

// HTTPClientFailureServerTransport connects instances of httpinfrastructuregroup.HTTPClientFailureClient to instances of HTTPClientFailureServer.
// Don't use this type directly, use NewHTTPClientFailureServerTransport instead.
type HTTPClientFailureServerTransport struct {
	srv *HTTPClientFailureServer
}

// Do implements the policy.Transporter interface for HTTPClientFailureServerTransport.
func (h *HTTPClientFailureServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return h.dispatchToMethodFake(req, method)
}

func (h *HTTPClientFailureServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "HTTPClientFailureClient.Delete400":
			res.resp, res.err = h.dispatchDelete400(req)
		case "HTTPClientFailureClient.Delete407":
			res.resp, res.err = h.dispatchDelete407(req)
		case "HTTPClientFailureClient.Delete417":
			res.resp, res.err = h.dispatchDelete417(req)
		case "HTTPClientFailureClient.Get400":
			res.resp, res.err = h.dispatchGet400(req)
		case "HTTPClientFailureClient.Get402":
			res.resp, res.err = h.dispatchGet402(req)
		case "HTTPClientFailureClient.Get403":
			res.resp, res.err = h.dispatchGet403(req)
		case "HTTPClientFailureClient.Get411":
			res.resp, res.err = h.dispatchGet411(req)
		case "HTTPClientFailureClient.Get412":
			res.resp, res.err = h.dispatchGet412(req)
		case "HTTPClientFailureClient.Get416":
			res.resp, res.err = h.dispatchGet416(req)
		case "HTTPClientFailureClient.Head400":
			res.resp, res.err = h.dispatchHead400(req)
		case "HTTPClientFailureClient.Head401":
			res.resp, res.err = h.dispatchHead401(req)
		case "HTTPClientFailureClient.Head410":
			res.resp, res.err = h.dispatchHead410(req)
		case "HTTPClientFailureClient.Head429":
			res.resp, res.err = h.dispatchHead429(req)
		case "HTTPClientFailureClient.Options400":
			res.resp, res.err = h.dispatchOptions400(req)
		case "HTTPClientFailureClient.Options403":
			res.resp, res.err = h.dispatchOptions403(req)
		case "HTTPClientFailureClient.Options412":
			res.resp, res.err = h.dispatchOptions412(req)
		case "HTTPClientFailureClient.Patch400":
			res.resp, res.err = h.dispatchPatch400(req)
		case "HTTPClientFailureClient.Patch405":
			res.resp, res.err = h.dispatchPatch405(req)
		case "HTTPClientFailureClient.Patch414":
			res.resp, res.err = h.dispatchPatch414(req)
		case "HTTPClientFailureClient.Post400":
			res.resp, res.err = h.dispatchPost400(req)
		case "HTTPClientFailureClient.Post406":
			res.resp, res.err = h.dispatchPost406(req)
		case "HTTPClientFailureClient.Post415":
			res.resp, res.err = h.dispatchPost415(req)
		case "HTTPClientFailureClient.Put400":
			res.resp, res.err = h.dispatchPut400(req)
		case "HTTPClientFailureClient.Put404":
			res.resp, res.err = h.dispatchPut404(req)
		case "HTTPClientFailureClient.Put409":
			res.resp, res.err = h.dispatchPut409(req)
		case "HTTPClientFailureClient.Put413":
			res.resp, res.err = h.dispatchPut413(req)
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

func (h *HTTPClientFailureServerTransport) dispatchDelete400(req *http.Request) (*http.Response, error) {
	if h.srv.Delete400 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete400 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPClientFailureClientDelete400Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPClientFailureClientDelete400Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Delete400(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchDelete407(req *http.Request) (*http.Response, error) {
	if h.srv.Delete407 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete407 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPClientFailureClientDelete407Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPClientFailureClientDelete407Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Delete407(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchDelete417(req *http.Request) (*http.Response, error) {
	if h.srv.Delete417 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete417 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPClientFailureClientDelete417Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPClientFailureClientDelete417Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Delete417(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchGet400(req *http.Request) (*http.Response, error) {
	if h.srv.Get400 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get400 not implemented")}
	}
	respr, errRespr := h.srv.Get400(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchGet402(req *http.Request) (*http.Response, error) {
	if h.srv.Get402 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get402 not implemented")}
	}
	respr, errRespr := h.srv.Get402(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchGet403(req *http.Request) (*http.Response, error) {
	if h.srv.Get403 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get403 not implemented")}
	}
	respr, errRespr := h.srv.Get403(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchGet411(req *http.Request) (*http.Response, error) {
	if h.srv.Get411 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get411 not implemented")}
	}
	respr, errRespr := h.srv.Get411(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchGet412(req *http.Request) (*http.Response, error) {
	if h.srv.Get412 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get412 not implemented")}
	}
	respr, errRespr := h.srv.Get412(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchGet416(req *http.Request) (*http.Response, error) {
	if h.srv.Get416 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get416 not implemented")}
	}
	respr, errRespr := h.srv.Get416(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchHead400(req *http.Request) (*http.Response, error) {
	if h.srv.Head400 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head400 not implemented")}
	}
	respr, errRespr := h.srv.Head400(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchHead401(req *http.Request) (*http.Response, error) {
	if h.srv.Head401 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head401 not implemented")}
	}
	respr, errRespr := h.srv.Head401(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchHead410(req *http.Request) (*http.Response, error) {
	if h.srv.Head410 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head410 not implemented")}
	}
	respr, errRespr := h.srv.Head410(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchHead429(req *http.Request) (*http.Response, error) {
	if h.srv.Head429 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head429 not implemented")}
	}
	respr, errRespr := h.srv.Head429(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchOptions400(req *http.Request) (*http.Response, error) {
	if h.srv.Options400 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Options400 not implemented")}
	}
	respr, errRespr := h.srv.Options400(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchOptions403(req *http.Request) (*http.Response, error) {
	if h.srv.Options403 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Options403 not implemented")}
	}
	respr, errRespr := h.srv.Options403(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchOptions412(req *http.Request) (*http.Response, error) {
	if h.srv.Options412 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Options412 not implemented")}
	}
	respr, errRespr := h.srv.Options412(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPatch400(req *http.Request) (*http.Response, error) {
	if h.srv.Patch400 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch400 not implemented")}
	}
	respr, errRespr := h.srv.Patch400(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPatch405(req *http.Request) (*http.Response, error) {
	if h.srv.Patch405 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch405 not implemented")}
	}
	respr, errRespr := h.srv.Patch405(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPatch414(req *http.Request) (*http.Response, error) {
	if h.srv.Patch414 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch414 not implemented")}
	}
	respr, errRespr := h.srv.Patch414(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPost400(req *http.Request) (*http.Response, error) {
	if h.srv.Post400 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post400 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPClientFailureClientPost400Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPClientFailureClientPost400Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Post400(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPost406(req *http.Request) (*http.Response, error) {
	if h.srv.Post406 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post406 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPClientFailureClientPost406Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPClientFailureClientPost406Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Post406(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPost415(req *http.Request) (*http.Response, error) {
	if h.srv.Post415 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post415 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPClientFailureClientPost415Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPClientFailureClientPost415Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Post415(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPut400(req *http.Request) (*http.Response, error) {
	if h.srv.Put400 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put400 not implemented")}
	}
	respr, errRespr := h.srv.Put400(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPut404(req *http.Request) (*http.Response, error) {
	if h.srv.Put404 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put404 not implemented")}
	}
	respr, errRespr := h.srv.Put404(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPut409(req *http.Request) (*http.Response, error) {
	if h.srv.Put409 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put409 not implemented")}
	}
	respr, errRespr := h.srv.Put409(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPClientFailureServerTransport) dispatchPut413(req *http.Request) (*http.Response, error) {
	if h.srv.Put413 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put413 not implemented")}
	}
	respr, errRespr := h.srv.Put413(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
