// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/numbergroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// NumberServer is a fake server for instances of the numbergroup.NumberClient type.
type NumberServer struct {
	// GetBigDecimal is the fake for method NumberClient.GetBigDecimal
	// HTTP status codes to indicate success: http.StatusOK
	GetBigDecimal func(ctx context.Context, options *numbergroup.NumberClientGetBigDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientGetBigDecimalResponse], errResp azfake.ErrorResponder)

	// GetBigDecimalNegativeDecimal is the fake for method NumberClient.GetBigDecimalNegativeDecimal
	// HTTP status codes to indicate success: http.StatusOK
	GetBigDecimalNegativeDecimal func(ctx context.Context, options *numbergroup.NumberClientGetBigDecimalNegativeDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientGetBigDecimalNegativeDecimalResponse], errResp azfake.ErrorResponder)

	// GetBigDecimalPositiveDecimal is the fake for method NumberClient.GetBigDecimalPositiveDecimal
	// HTTP status codes to indicate success: http.StatusOK
	GetBigDecimalPositiveDecimal func(ctx context.Context, options *numbergroup.NumberClientGetBigDecimalPositiveDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientGetBigDecimalPositiveDecimalResponse], errResp azfake.ErrorResponder)

	// GetBigDouble is the fake for method NumberClient.GetBigDouble
	// HTTP status codes to indicate success: http.StatusOK
	GetBigDouble func(ctx context.Context, options *numbergroup.NumberClientGetBigDoubleOptions) (resp azfake.Responder[numbergroup.NumberClientGetBigDoubleResponse], errResp azfake.ErrorResponder)

	// GetBigDoubleNegativeDecimal is the fake for method NumberClient.GetBigDoubleNegativeDecimal
	// HTTP status codes to indicate success: http.StatusOK
	GetBigDoubleNegativeDecimal func(ctx context.Context, options *numbergroup.NumberClientGetBigDoubleNegativeDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientGetBigDoubleNegativeDecimalResponse], errResp azfake.ErrorResponder)

	// GetBigDoublePositiveDecimal is the fake for method NumberClient.GetBigDoublePositiveDecimal
	// HTTP status codes to indicate success: http.StatusOK
	GetBigDoublePositiveDecimal func(ctx context.Context, options *numbergroup.NumberClientGetBigDoublePositiveDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientGetBigDoublePositiveDecimalResponse], errResp azfake.ErrorResponder)

	// GetBigFloat is the fake for method NumberClient.GetBigFloat
	// HTTP status codes to indicate success: http.StatusOK
	GetBigFloat func(ctx context.Context, options *numbergroup.NumberClientGetBigFloatOptions) (resp azfake.Responder[numbergroup.NumberClientGetBigFloatResponse], errResp azfake.ErrorResponder)

	// GetInvalidDecimal is the fake for method NumberClient.GetInvalidDecimal
	// HTTP status codes to indicate success: http.StatusOK
	GetInvalidDecimal func(ctx context.Context, options *numbergroup.NumberClientGetInvalidDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientGetInvalidDecimalResponse], errResp azfake.ErrorResponder)

	// GetInvalidDouble is the fake for method NumberClient.GetInvalidDouble
	// HTTP status codes to indicate success: http.StatusOK
	GetInvalidDouble func(ctx context.Context, options *numbergroup.NumberClientGetInvalidDoubleOptions) (resp azfake.Responder[numbergroup.NumberClientGetInvalidDoubleResponse], errResp azfake.ErrorResponder)

	// GetInvalidFloat is the fake for method NumberClient.GetInvalidFloat
	// HTTP status codes to indicate success: http.StatusOK
	GetInvalidFloat func(ctx context.Context, options *numbergroup.NumberClientGetInvalidFloatOptions) (resp azfake.Responder[numbergroup.NumberClientGetInvalidFloatResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method NumberClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *numbergroup.NumberClientGetNullOptions) (resp azfake.Responder[numbergroup.NumberClientGetNullResponse], errResp azfake.ErrorResponder)

	// GetSmallDecimal is the fake for method NumberClient.GetSmallDecimal
	// HTTP status codes to indicate success: http.StatusOK
	GetSmallDecimal func(ctx context.Context, options *numbergroup.NumberClientGetSmallDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientGetSmallDecimalResponse], errResp azfake.ErrorResponder)

	// GetSmallDouble is the fake for method NumberClient.GetSmallDouble
	// HTTP status codes to indicate success: http.StatusOK
	GetSmallDouble func(ctx context.Context, options *numbergroup.NumberClientGetSmallDoubleOptions) (resp azfake.Responder[numbergroup.NumberClientGetSmallDoubleResponse], errResp azfake.ErrorResponder)

	// GetSmallFloat is the fake for method NumberClient.GetSmallFloat
	// HTTP status codes to indicate success: http.StatusOK
	GetSmallFloat func(ctx context.Context, options *numbergroup.NumberClientGetSmallFloatOptions) (resp azfake.Responder[numbergroup.NumberClientGetSmallFloatResponse], errResp azfake.ErrorResponder)

	// PutBigDecimal is the fake for method NumberClient.PutBigDecimal
	// HTTP status codes to indicate success: http.StatusOK
	PutBigDecimal func(ctx context.Context, numberBody float64, options *numbergroup.NumberClientPutBigDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientPutBigDecimalResponse], errResp azfake.ErrorResponder)

	// PutBigDecimalNegativeDecimal is the fake for method NumberClient.PutBigDecimalNegativeDecimal
	// HTTP status codes to indicate success: http.StatusOK
	PutBigDecimalNegativeDecimal func(ctx context.Context, options *numbergroup.NumberClientPutBigDecimalNegativeDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientPutBigDecimalNegativeDecimalResponse], errResp azfake.ErrorResponder)

	// PutBigDecimalPositiveDecimal is the fake for method NumberClient.PutBigDecimalPositiveDecimal
	// HTTP status codes to indicate success: http.StatusOK
	PutBigDecimalPositiveDecimal func(ctx context.Context, options *numbergroup.NumberClientPutBigDecimalPositiveDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientPutBigDecimalPositiveDecimalResponse], errResp azfake.ErrorResponder)

	// PutBigDouble is the fake for method NumberClient.PutBigDouble
	// HTTP status codes to indicate success: http.StatusOK
	PutBigDouble func(ctx context.Context, numberBody float64, options *numbergroup.NumberClientPutBigDoubleOptions) (resp azfake.Responder[numbergroup.NumberClientPutBigDoubleResponse], errResp azfake.ErrorResponder)

	// PutBigDoubleNegativeDecimal is the fake for method NumberClient.PutBigDoubleNegativeDecimal
	// HTTP status codes to indicate success: http.StatusOK
	PutBigDoubleNegativeDecimal func(ctx context.Context, options *numbergroup.NumberClientPutBigDoubleNegativeDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientPutBigDoubleNegativeDecimalResponse], errResp azfake.ErrorResponder)

	// PutBigDoublePositiveDecimal is the fake for method NumberClient.PutBigDoublePositiveDecimal
	// HTTP status codes to indicate success: http.StatusOK
	PutBigDoublePositiveDecimal func(ctx context.Context, options *numbergroup.NumberClientPutBigDoublePositiveDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientPutBigDoublePositiveDecimalResponse], errResp azfake.ErrorResponder)

	// PutBigFloat is the fake for method NumberClient.PutBigFloat
	// HTTP status codes to indicate success: http.StatusOK
	PutBigFloat func(ctx context.Context, numberBody float32, options *numbergroup.NumberClientPutBigFloatOptions) (resp azfake.Responder[numbergroup.NumberClientPutBigFloatResponse], errResp azfake.ErrorResponder)

	// PutSmallDecimal is the fake for method NumberClient.PutSmallDecimal
	// HTTP status codes to indicate success: http.StatusOK
	PutSmallDecimal func(ctx context.Context, numberBody float64, options *numbergroup.NumberClientPutSmallDecimalOptions) (resp azfake.Responder[numbergroup.NumberClientPutSmallDecimalResponse], errResp azfake.ErrorResponder)

	// PutSmallDouble is the fake for method NumberClient.PutSmallDouble
	// HTTP status codes to indicate success: http.StatusOK
	PutSmallDouble func(ctx context.Context, numberBody float64, options *numbergroup.NumberClientPutSmallDoubleOptions) (resp azfake.Responder[numbergroup.NumberClientPutSmallDoubleResponse], errResp azfake.ErrorResponder)

	// PutSmallFloat is the fake for method NumberClient.PutSmallFloat
	// HTTP status codes to indicate success: http.StatusOK
	PutSmallFloat func(ctx context.Context, numberBody float32, options *numbergroup.NumberClientPutSmallFloatOptions) (resp azfake.Responder[numbergroup.NumberClientPutSmallFloatResponse], errResp azfake.ErrorResponder)
}

// NewNumberServerTransport creates a new instance of NumberServerTransport with the provided implementation.
// The returned NumberServerTransport instance is connected to an instance of numbergroup.NumberClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNumberServerTransport(srv *NumberServer) *NumberServerTransport {
	return &NumberServerTransport{srv: srv}
}

// NumberServerTransport connects instances of numbergroup.NumberClient to instances of NumberServer.
// Don't use this type directly, use NewNumberServerTransport instead.
type NumberServerTransport struct {
	srv *NumberServer
}

// Do implements the policy.Transporter interface for NumberServerTransport.
func (n *NumberServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NumberServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "NumberClient.GetBigDecimal":
			res.resp, res.err = n.dispatchGetBigDecimal(req)
		case "NumberClient.GetBigDecimalNegativeDecimal":
			res.resp, res.err = n.dispatchGetBigDecimalNegativeDecimal(req)
		case "NumberClient.GetBigDecimalPositiveDecimal":
			res.resp, res.err = n.dispatchGetBigDecimalPositiveDecimal(req)
		case "NumberClient.GetBigDouble":
			res.resp, res.err = n.dispatchGetBigDouble(req)
		case "NumberClient.GetBigDoubleNegativeDecimal":
			res.resp, res.err = n.dispatchGetBigDoubleNegativeDecimal(req)
		case "NumberClient.GetBigDoublePositiveDecimal":
			res.resp, res.err = n.dispatchGetBigDoublePositiveDecimal(req)
		case "NumberClient.GetBigFloat":
			res.resp, res.err = n.dispatchGetBigFloat(req)
		case "NumberClient.GetInvalidDecimal":
			res.resp, res.err = n.dispatchGetInvalidDecimal(req)
		case "NumberClient.GetInvalidDouble":
			res.resp, res.err = n.dispatchGetInvalidDouble(req)
		case "NumberClient.GetInvalidFloat":
			res.resp, res.err = n.dispatchGetInvalidFloat(req)
		case "NumberClient.GetNull":
			res.resp, res.err = n.dispatchGetNull(req)
		case "NumberClient.GetSmallDecimal":
			res.resp, res.err = n.dispatchGetSmallDecimal(req)
		case "NumberClient.GetSmallDouble":
			res.resp, res.err = n.dispatchGetSmallDouble(req)
		case "NumberClient.GetSmallFloat":
			res.resp, res.err = n.dispatchGetSmallFloat(req)
		case "NumberClient.PutBigDecimal":
			res.resp, res.err = n.dispatchPutBigDecimal(req)
		case "NumberClient.PutBigDecimalNegativeDecimal":
			res.resp, res.err = n.dispatchPutBigDecimalNegativeDecimal(req)
		case "NumberClient.PutBigDecimalPositiveDecimal":
			res.resp, res.err = n.dispatchPutBigDecimalPositiveDecimal(req)
		case "NumberClient.PutBigDouble":
			res.resp, res.err = n.dispatchPutBigDouble(req)
		case "NumberClient.PutBigDoubleNegativeDecimal":
			res.resp, res.err = n.dispatchPutBigDoubleNegativeDecimal(req)
		case "NumberClient.PutBigDoublePositiveDecimal":
			res.resp, res.err = n.dispatchPutBigDoublePositiveDecimal(req)
		case "NumberClient.PutBigFloat":
			res.resp, res.err = n.dispatchPutBigFloat(req)
		case "NumberClient.PutSmallDecimal":
			res.resp, res.err = n.dispatchPutSmallDecimal(req)
		case "NumberClient.PutSmallDouble":
			res.resp, res.err = n.dispatchPutSmallDouble(req)
		case "NumberClient.PutSmallFloat":
			res.resp, res.err = n.dispatchPutSmallFloat(req)
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

func (n *NumberServerTransport) dispatchGetBigDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.GetBigDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBigDecimal not implemented")}
	}
	respr, errRespr := n.srv.GetBigDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetBigDecimalNegativeDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.GetBigDecimalNegativeDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBigDecimalNegativeDecimal not implemented")}
	}
	respr, errRespr := n.srv.GetBigDecimalNegativeDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetBigDecimalPositiveDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.GetBigDecimalPositiveDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBigDecimalPositiveDecimal not implemented")}
	}
	respr, errRespr := n.srv.GetBigDecimalPositiveDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetBigDouble(req *http.Request) (*http.Response, error) {
	if n.srv.GetBigDouble == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBigDouble not implemented")}
	}
	respr, errRespr := n.srv.GetBigDouble(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetBigDoubleNegativeDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.GetBigDoubleNegativeDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBigDoubleNegativeDecimal not implemented")}
	}
	respr, errRespr := n.srv.GetBigDoubleNegativeDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetBigDoublePositiveDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.GetBigDoublePositiveDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBigDoublePositiveDecimal not implemented")}
	}
	respr, errRespr := n.srv.GetBigDoublePositiveDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetBigFloat(req *http.Request) (*http.Response, error) {
	if n.srv.GetBigFloat == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBigFloat not implemented")}
	}
	respr, errRespr := n.srv.GetBigFloat(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetInvalidDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.GetInvalidDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInvalidDecimal not implemented")}
	}
	respr, errRespr := n.srv.GetInvalidDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetInvalidDouble(req *http.Request) (*http.Response, error) {
	if n.srv.GetInvalidDouble == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInvalidDouble not implemented")}
	}
	respr, errRespr := n.srv.GetInvalidDouble(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetInvalidFloat(req *http.Request) (*http.Response, error) {
	if n.srv.GetInvalidFloat == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInvalidFloat not implemented")}
	}
	respr, errRespr := n.srv.GetInvalidFloat(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if n.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNull not implemented")}
	}
	respr, errRespr := n.srv.GetNull(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetSmallDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.GetSmallDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSmallDecimal not implemented")}
	}
	respr, errRespr := n.srv.GetSmallDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetSmallDouble(req *http.Request) (*http.Response, error) {
	if n.srv.GetSmallDouble == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSmallDouble not implemented")}
	}
	respr, errRespr := n.srv.GetSmallDouble(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchGetSmallFloat(req *http.Request) (*http.Response, error) {
	if n.srv.GetSmallFloat == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSmallFloat not implemented")}
	}
	respr, errRespr := n.srv.GetSmallFloat(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchPutBigDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.PutBigDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutBigDecimal not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[float64](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PutBigDecimal(req.Context(), body, nil)
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

func (n *NumberServerTransport) dispatchPutBigDecimalNegativeDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.PutBigDecimalNegativeDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutBigDecimalNegativeDecimal not implemented")}
	}
	respr, errRespr := n.srv.PutBigDecimalNegativeDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchPutBigDecimalPositiveDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.PutBigDecimalPositiveDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutBigDecimalPositiveDecimal not implemented")}
	}
	respr, errRespr := n.srv.PutBigDecimalPositiveDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchPutBigDouble(req *http.Request) (*http.Response, error) {
	if n.srv.PutBigDouble == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutBigDouble not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[float64](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PutBigDouble(req.Context(), body, nil)
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

func (n *NumberServerTransport) dispatchPutBigDoubleNegativeDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.PutBigDoubleNegativeDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutBigDoubleNegativeDecimal not implemented")}
	}
	respr, errRespr := n.srv.PutBigDoubleNegativeDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchPutBigDoublePositiveDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.PutBigDoublePositiveDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutBigDoublePositiveDecimal not implemented")}
	}
	respr, errRespr := n.srv.PutBigDoublePositiveDecimal(req.Context(), nil)
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

func (n *NumberServerTransport) dispatchPutBigFloat(req *http.Request) (*http.Response, error) {
	if n.srv.PutBigFloat == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutBigFloat not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[float32](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PutBigFloat(req.Context(), body, nil)
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

func (n *NumberServerTransport) dispatchPutSmallDecimal(req *http.Request) (*http.Response, error) {
	if n.srv.PutSmallDecimal == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutSmallDecimal not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[float64](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PutSmallDecimal(req.Context(), body, nil)
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

func (n *NumberServerTransport) dispatchPutSmallDouble(req *http.Request) (*http.Response, error) {
	if n.srv.PutSmallDouble == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutSmallDouble not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[float64](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PutSmallDouble(req.Context(), body, nil)
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

func (n *NumberServerTransport) dispatchPutSmallFloat(req *http.Request) (*http.Response, error) {
	if n.srv.PutSmallFloat == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutSmallFloat not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[float32](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PutSmallFloat(req.Context(), body, nil)
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
