// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/custombaseurlgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PathsServer is a fake server for instances of the custombaseurlgroup.PathsClient type.
type PathsServer struct {
	// GetEmpty is the fake for method PathsClient.GetEmpty
	// HTTP status codes to indicate success: http.StatusOK
	GetEmpty func(ctx context.Context, host string, options *custombaseurlgroup.PathsClientGetEmptyOptions) (resp azfake.Responder[custombaseurlgroup.PathsClientGetEmptyResponse], errResp azfake.ErrorResponder)
}

// NewPathsServerTransport creates a new instance of PathsServerTransport with the provided implementation.
// The returned PathsServerTransport instance is connected to an instance of custombaseurlgroup.PathsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPathsServerTransport(srv *PathsServer) *PathsServerTransport {
	return &PathsServerTransport{srv: srv}
}

// PathsServerTransport connects instances of custombaseurlgroup.PathsClient to instances of PathsServer.
// Don't use this type directly, use NewPathsServerTransport instead.
type PathsServerTransport struct {
	srv *PathsServer
}

// Do implements the policy.Transporter interface for PathsServerTransport.
func (p *PathsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PathsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "PathsClient.GetEmpty":
			res.resp, res.err = p.dispatchGetEmpty(req)
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

func (p *PathsServerTransport) dispatchGetEmpty(req *http.Request) (*http.Response, error) {
	if p.srv.GetEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEmpty not implemented")}
	}
	respr, errRespr := p.srv.GetEmpty(req.Context(), req.URL.Host, nil)
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
