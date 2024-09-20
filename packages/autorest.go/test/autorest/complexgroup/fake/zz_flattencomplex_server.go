// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/complexgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FlattencomplexServer is a fake server for instances of the complexgroup.FlattencomplexClient type.
type FlattencomplexServer struct {
	// GetValid is the fake for method FlattencomplexClient.GetValid
	// HTTP status codes to indicate success: http.StatusOK
	GetValid func(ctx context.Context, options *complexgroup.FlattencomplexClientGetValidOptions) (resp azfake.Responder[complexgroup.FlattencomplexClientGetValidResponse], errResp azfake.ErrorResponder)
}

// NewFlattencomplexServerTransport creates a new instance of FlattencomplexServerTransport with the provided implementation.
// The returned FlattencomplexServerTransport instance is connected to an instance of complexgroup.FlattencomplexClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFlattencomplexServerTransport(srv *FlattencomplexServer) *FlattencomplexServerTransport {
	return &FlattencomplexServerTransport{srv: srv}
}

// FlattencomplexServerTransport connects instances of complexgroup.FlattencomplexClient to instances of FlattencomplexServer.
// Don't use this type directly, use NewFlattencomplexServerTransport instead.
type FlattencomplexServerTransport struct {
	srv *FlattencomplexServer
}

// Do implements the policy.Transporter interface for FlattencomplexServerTransport.
func (f *FlattencomplexServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FlattencomplexServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "FlattencomplexClient.GetValid":
			res.resp, res.err = f.dispatchGetValid(req)
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

func (f *FlattencomplexServerTransport) dispatchGetValid(req *http.Request) (*http.Response, error) {
	if f.srv.GetValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetValid not implemented")}
	}
	respr, errRespr := f.srv.GetValid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MyBaseTypeClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
