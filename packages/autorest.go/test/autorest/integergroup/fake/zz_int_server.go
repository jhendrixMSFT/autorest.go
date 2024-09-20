// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/integergroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// IntServer is a fake server for instances of the integergroup.IntClient type.
type IntServer struct {
	// GetInvalid is the fake for method IntClient.GetInvalid
	// HTTP status codes to indicate success: http.StatusOK
	GetInvalid func(ctx context.Context, options *integergroup.IntClientGetInvalidOptions) (resp azfake.Responder[integergroup.IntClientGetInvalidResponse], errResp azfake.ErrorResponder)

	// GetInvalidUnixTime is the fake for method IntClient.GetInvalidUnixTime
	// HTTP status codes to indicate success: http.StatusOK
	GetInvalidUnixTime func(ctx context.Context, options *integergroup.IntClientGetInvalidUnixTimeOptions) (resp azfake.Responder[integergroup.IntClientGetInvalidUnixTimeResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method IntClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *integergroup.IntClientGetNullOptions) (resp azfake.Responder[integergroup.IntClientGetNullResponse], errResp azfake.ErrorResponder)

	// GetNullUnixTime is the fake for method IntClient.GetNullUnixTime
	// HTTP status codes to indicate success: http.StatusOK
	GetNullUnixTime func(ctx context.Context, options *integergroup.IntClientGetNullUnixTimeOptions) (resp azfake.Responder[integergroup.IntClientGetNullUnixTimeResponse], errResp azfake.ErrorResponder)

	// GetOverflowInt32 is the fake for method IntClient.GetOverflowInt32
	// HTTP status codes to indicate success: http.StatusOK
	GetOverflowInt32 func(ctx context.Context, options *integergroup.IntClientGetOverflowInt32Options) (resp azfake.Responder[integergroup.IntClientGetOverflowInt32Response], errResp azfake.ErrorResponder)

	// GetOverflowInt64 is the fake for method IntClient.GetOverflowInt64
	// HTTP status codes to indicate success: http.StatusOK
	GetOverflowInt64 func(ctx context.Context, options *integergroup.IntClientGetOverflowInt64Options) (resp azfake.Responder[integergroup.IntClientGetOverflowInt64Response], errResp azfake.ErrorResponder)

	// GetUnderflowInt32 is the fake for method IntClient.GetUnderflowInt32
	// HTTP status codes to indicate success: http.StatusOK
	GetUnderflowInt32 func(ctx context.Context, options *integergroup.IntClientGetUnderflowInt32Options) (resp azfake.Responder[integergroup.IntClientGetUnderflowInt32Response], errResp azfake.ErrorResponder)

	// GetUnderflowInt64 is the fake for method IntClient.GetUnderflowInt64
	// HTTP status codes to indicate success: http.StatusOK
	GetUnderflowInt64 func(ctx context.Context, options *integergroup.IntClientGetUnderflowInt64Options) (resp azfake.Responder[integergroup.IntClientGetUnderflowInt64Response], errResp azfake.ErrorResponder)

	// GetUnixTime is the fake for method IntClient.GetUnixTime
	// HTTP status codes to indicate success: http.StatusOK
	GetUnixTime func(ctx context.Context, options *integergroup.IntClientGetUnixTimeOptions) (resp azfake.Responder[integergroup.IntClientGetUnixTimeResponse], errResp azfake.ErrorResponder)

	// PutMax32 is the fake for method IntClient.PutMax32
	// HTTP status codes to indicate success: http.StatusOK
	PutMax32 func(ctx context.Context, intBody int32, options *integergroup.IntClientPutMax32Options) (resp azfake.Responder[integergroup.IntClientPutMax32Response], errResp azfake.ErrorResponder)

	// PutMax64 is the fake for method IntClient.PutMax64
	// HTTP status codes to indicate success: http.StatusOK
	PutMax64 func(ctx context.Context, intBody int64, options *integergroup.IntClientPutMax64Options) (resp azfake.Responder[integergroup.IntClientPutMax64Response], errResp azfake.ErrorResponder)

	// PutMin32 is the fake for method IntClient.PutMin32
	// HTTP status codes to indicate success: http.StatusOK
	PutMin32 func(ctx context.Context, intBody int32, options *integergroup.IntClientPutMin32Options) (resp azfake.Responder[integergroup.IntClientPutMin32Response], errResp azfake.ErrorResponder)

	// PutMin64 is the fake for method IntClient.PutMin64
	// HTTP status codes to indicate success: http.StatusOK
	PutMin64 func(ctx context.Context, intBody int64, options *integergroup.IntClientPutMin64Options) (resp azfake.Responder[integergroup.IntClientPutMin64Response], errResp azfake.ErrorResponder)

	// PutUnixTimeDate is the fake for method IntClient.PutUnixTimeDate
	// HTTP status codes to indicate success: http.StatusOK
	PutUnixTimeDate func(ctx context.Context, intBody time.Time, options *integergroup.IntClientPutUnixTimeDateOptions) (resp azfake.Responder[integergroup.IntClientPutUnixTimeDateResponse], errResp azfake.ErrorResponder)
}

// NewIntServerTransport creates a new instance of IntServerTransport with the provided implementation.
// The returned IntServerTransport instance is connected to an instance of integergroup.IntClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIntServerTransport(srv *IntServer) *IntServerTransport {
	return &IntServerTransport{srv: srv}
}

// IntServerTransport connects instances of integergroup.IntClient to instances of IntServer.
// Don't use this type directly, use NewIntServerTransport instead.
type IntServerTransport struct {
	srv *IntServer
}

// Do implements the policy.Transporter interface for IntServerTransport.
func (i *IntServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *IntServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "IntClient.GetInvalid":
			res.resp, res.err = i.dispatchGetInvalid(req)
		case "IntClient.GetInvalidUnixTime":
			res.resp, res.err = i.dispatchGetInvalidUnixTime(req)
		case "IntClient.GetNull":
			res.resp, res.err = i.dispatchGetNull(req)
		case "IntClient.GetNullUnixTime":
			res.resp, res.err = i.dispatchGetNullUnixTime(req)
		case "IntClient.GetOverflowInt32":
			res.resp, res.err = i.dispatchGetOverflowInt32(req)
		case "IntClient.GetOverflowInt64":
			res.resp, res.err = i.dispatchGetOverflowInt64(req)
		case "IntClient.GetUnderflowInt32":
			res.resp, res.err = i.dispatchGetUnderflowInt32(req)
		case "IntClient.GetUnderflowInt64":
			res.resp, res.err = i.dispatchGetUnderflowInt64(req)
		case "IntClient.GetUnixTime":
			res.resp, res.err = i.dispatchGetUnixTime(req)
		case "IntClient.PutMax32":
			res.resp, res.err = i.dispatchPutMax32(req)
		case "IntClient.PutMax64":
			res.resp, res.err = i.dispatchPutMax64(req)
		case "IntClient.PutMin32":
			res.resp, res.err = i.dispatchPutMin32(req)
		case "IntClient.PutMin64":
			res.resp, res.err = i.dispatchPutMin64(req)
		case "IntClient.PutUnixTimeDate":
			res.resp, res.err = i.dispatchPutUnixTimeDate(req)
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

func (i *IntServerTransport) dispatchGetInvalid(req *http.Request) (*http.Response, error) {
	if i.srv.GetInvalid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInvalid not implemented")}
	}
	respr, errRespr := i.srv.GetInvalid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntServerTransport) dispatchGetInvalidUnixTime(req *http.Request) (*http.Response, error) {
	if i.srv.GetInvalidUnixTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInvalidUnixTime not implemented")}
	}
	respr, errRespr := i.srv.GetInvalidUnixTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*timeUnix)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if i.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNull not implemented")}
	}
	respr, errRespr := i.srv.GetNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntServerTransport) dispatchGetNullUnixTime(req *http.Request) (*http.Response, error) {
	if i.srv.GetNullUnixTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNullUnixTime not implemented")}
	}
	respr, errRespr := i.srv.GetNullUnixTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*timeUnix)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntServerTransport) dispatchGetOverflowInt32(req *http.Request) (*http.Response, error) {
	if i.srv.GetOverflowInt32 == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetOverflowInt32 not implemented")}
	}
	respr, errRespr := i.srv.GetOverflowInt32(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntServerTransport) dispatchGetOverflowInt64(req *http.Request) (*http.Response, error) {
	if i.srv.GetOverflowInt64 == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetOverflowInt64 not implemented")}
	}
	respr, errRespr := i.srv.GetOverflowInt64(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntServerTransport) dispatchGetUnderflowInt32(req *http.Request) (*http.Response, error) {
	if i.srv.GetUnderflowInt32 == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUnderflowInt32 not implemented")}
	}
	respr, errRespr := i.srv.GetUnderflowInt32(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntServerTransport) dispatchGetUnderflowInt64(req *http.Request) (*http.Response, error) {
	if i.srv.GetUnderflowInt64 == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUnderflowInt64 not implemented")}
	}
	respr, errRespr := i.srv.GetUnderflowInt64(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntServerTransport) dispatchGetUnixTime(req *http.Request) (*http.Response, error) {
	if i.srv.GetUnixTime == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUnixTime not implemented")}
	}
	respr, errRespr := i.srv.GetUnixTime(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, (*timeUnix)(server.GetResponse(respr).Value), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntServerTransport) dispatchPutMax32(req *http.Request) (*http.Response, error) {
	if i.srv.PutMax32 == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutMax32 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[int32](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.PutMax32(req.Context(), body, nil)
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

func (i *IntServerTransport) dispatchPutMax64(req *http.Request) (*http.Response, error) {
	if i.srv.PutMax64 == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutMax64 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[int64](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.PutMax64(req.Context(), body, nil)
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

func (i *IntServerTransport) dispatchPutMin32(req *http.Request) (*http.Response, error) {
	if i.srv.PutMin32 == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutMin32 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[int32](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.PutMin32(req.Context(), body, nil)
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

func (i *IntServerTransport) dispatchPutMin64(req *http.Request) (*http.Response, error) {
	if i.srv.PutMin64 == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutMin64 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[int64](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.PutMin64(req.Context(), body, nil)
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

func (i *IntServerTransport) dispatchPutUnixTimeDate(req *http.Request) (*http.Response, error) {
	if i.srv.PutUnixTimeDate == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutUnixTimeDate not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[timeUnix](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.PutUnixTimeDate(req.Context(), time.Time(body), nil)
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
