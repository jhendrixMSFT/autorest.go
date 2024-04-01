// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"generatortests/formdatagroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
)

// FormdataServer is a fake server for instances of the formdatagroup.FormdataClient type.
type FormdataServer struct {
	// UploadFile is the fake for method FormdataClient.UploadFile
	// HTTP status codes to indicate success: http.StatusOK
	UploadFile func(ctx context.Context, fileContent io.ReadSeekCloser, fileName string, options *formdatagroup.FormdataClientUploadFileOptions) (resp azfake.Responder[formdatagroup.FormdataClientUploadFileResponse], errResp azfake.ErrorResponder)

	// UploadFileViaBody is the fake for method FormdataClient.UploadFileViaBody
	// HTTP status codes to indicate success: http.StatusOK
	UploadFileViaBody func(ctx context.Context, fileContent io.ReadSeekCloser, options *formdatagroup.FormdataClientUploadFileViaBodyOptions) (resp azfake.Responder[formdatagroup.FormdataClientUploadFileViaBodyResponse], errResp azfake.ErrorResponder)

	// UploadFiles is the fake for method FormdataClient.UploadFiles
	// HTTP status codes to indicate success: http.StatusOK
	UploadFiles func(ctx context.Context, fileContent []io.ReadSeekCloser, options *formdatagroup.FormdataClientUploadFilesOptions) (resp azfake.Responder[formdatagroup.FormdataClientUploadFilesResponse], errResp azfake.ErrorResponder)
}

// NewFormdataServerTransport creates a new instance of FormdataServerTransport with the provided implementation.
// The returned FormdataServerTransport instance is connected to an instance of formdatagroup.FormdataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFormdataServerTransport(srv *FormdataServer) *FormdataServerTransport {
	return &FormdataServerTransport{srv: srv}
}

// FormdataServerTransport connects instances of formdatagroup.FormdataClient to instances of FormdataServer.
// Don't use this type directly, use NewFormdataServerTransport instead.
type FormdataServerTransport struct {
	srv *FormdataServer
}

// Do implements the policy.Transporter interface for FormdataServerTransport.
func (f *FormdataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FormdataServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "FormdataClient.UploadFile":
		resp, err = f.dispatchUploadFile(req)
	case "FormdataClient.UploadFileViaBody":
		resp, err = f.dispatchUploadFileViaBody(req)
	case "FormdataClient.UploadFiles":
		resp, err = f.dispatchUploadFiles(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (f *FormdataServerTransport) dispatchUploadFile(req *http.Request) (*http.Response, error) {
	if f.srv.UploadFile == nil {
		return nil, &nonRetriableError{errors.New("fake for method UploadFile not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var fileContent io.ReadSeekCloser
	var fileName string
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "fileContent":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			fileContent = streaming.NopCloser(bytes.NewReader(content))
		case "fileName":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			fileName = string(content)
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := f.srv.UploadFile(req.Context(), fileContent, fileName, nil)
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
	return resp, nil
}

func (f *FormdataServerTransport) dispatchUploadFileViaBody(req *http.Request) (*http.Response, error) {
	if f.srv.UploadFileViaBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method UploadFileViaBody not implemented")}
	}
	respr, errRespr := f.srv.UploadFileViaBody(req.Context(), req.Body.(io.ReadSeekCloser), nil)
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
	return resp, nil
}

func (f *FormdataServerTransport) dispatchUploadFiles(req *http.Request) (*http.Response, error) {
	if f.srv.UploadFiles == nil {
		return nil, &nonRetriableError{errors.New("fake for method UploadFiles not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var fileContent []io.ReadSeekCloser
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "fileContent":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			fileContent = append(fileContent, streaming.NopCloser(bytes.NewReader(content)))
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := f.srv.UploadFiles(req.Context(), fileContent, nil)
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
	return resp, nil
}
