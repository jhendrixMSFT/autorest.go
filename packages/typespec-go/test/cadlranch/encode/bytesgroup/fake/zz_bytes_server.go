// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// BytesServer is a fake server for instances of the bytesgroup.BytesClient type.
type BytesServer struct {
	// BytesHeaderServer contains the fakes for client BytesHeaderClient
	BytesHeaderServer BytesHeaderServer

	// BytesPropertyServer contains the fakes for client BytesPropertyClient
	BytesPropertyServer BytesPropertyServer

	// BytesQueryServer contains the fakes for client BytesQueryClient
	BytesQueryServer BytesQueryServer

	// BytesRequestBodyServer contains the fakes for client BytesRequestBodyClient
	BytesRequestBodyServer BytesRequestBodyServer

	// BytesResponseBodyServer contains the fakes for client BytesResponseBodyClient
	BytesResponseBodyServer BytesResponseBodyServer
}

// NewBytesServerTransport creates a new instance of BytesServerTransport with the provided implementation.
// The returned BytesServerTransport instance is connected to an instance of bytesgroup.BytesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBytesServerTransport(srv *BytesServer) *BytesServerTransport {
	return &BytesServerTransport{srv: srv}
}

// BytesServerTransport connects instances of bytesgroup.BytesClient to instances of BytesServer.
// Don't use this type directly, use NewBytesServerTransport instead.
type BytesServerTransport struct {
	srv                       *BytesServer
	trMu                      sync.Mutex
	trBytesHeaderServer       *BytesHeaderServerTransport
	trBytesPropertyServer     *BytesPropertyServerTransport
	trBytesQueryServer        *BytesQueryServerTransport
	trBytesRequestBodyServer  *BytesRequestBodyServerTransport
	trBytesResponseBodyServer *BytesResponseBodyServerTransport
}

// Do implements the policy.Transporter interface for BytesServerTransport.
func (b *BytesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (b *BytesServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "BytesHeaderClient":
		initServer(&b.trMu, &b.trBytesHeaderServer, func() *BytesHeaderServerTransport {
			return NewBytesHeaderServerTransport(&b.srv.BytesHeaderServer)
		})
		resp, err = b.trBytesHeaderServer.Do(req)
	case "BytesPropertyClient":
		initServer(&b.trMu, &b.trBytesPropertyServer, func() *BytesPropertyServerTransport {
			return NewBytesPropertyServerTransport(&b.srv.BytesPropertyServer)
		})
		resp, err = b.trBytesPropertyServer.Do(req)
	case "BytesQueryClient":
		initServer(&b.trMu, &b.trBytesQueryServer, func() *BytesQueryServerTransport {
			return NewBytesQueryServerTransport(&b.srv.BytesQueryServer)
		})
		resp, err = b.trBytesQueryServer.Do(req)
	case "BytesRequestBodyClient":
		initServer(&b.trMu, &b.trBytesRequestBodyServer, func() *BytesRequestBodyServerTransport {
			return NewBytesRequestBodyServerTransport(&b.srv.BytesRequestBodyServer)
		})
		resp, err = b.trBytesRequestBodyServer.Do(req)
	case "BytesResponseBodyClient":
		initServer(&b.trMu, &b.trBytesResponseBodyServer, func() *BytesResponseBodyServerTransport {
			return NewBytesResponseBodyServerTransport(&b.srv.BytesResponseBodyServer)
		})
		resp, err = b.trBytesResponseBodyServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}