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

// NullableServer is a fake server for instances of the nullablegroup.NullableClient type.
type NullableServer struct {
	// NullableBytesServer contains the fakes for client NullableBytesClient
	NullableBytesServer NullableBytesServer

	// NullableCollectionsByteServer contains the fakes for client NullableCollectionsByteClient
	NullableCollectionsByteServer NullableCollectionsByteServer

	// NullableCollectionsModelServer contains the fakes for client NullableCollectionsModelClient
	NullableCollectionsModelServer NullableCollectionsModelServer

	// NullableDatetimeServer contains the fakes for client NullableDatetimeClient
	NullableDatetimeServer NullableDatetimeServer

	// NullableDurationServer contains the fakes for client NullableDurationClient
	NullableDurationServer NullableDurationServer

	// NullableStringServer contains the fakes for client NullableStringClient
	NullableStringServer NullableStringServer
}

// NewNullableServerTransport creates a new instance of NullableServerTransport with the provided implementation.
// The returned NullableServerTransport instance is connected to an instance of nullablegroup.NullableClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNullableServerTransport(srv *NullableServer) *NullableServerTransport {
	return &NullableServerTransport{srv: srv}
}

// NullableServerTransport connects instances of nullablegroup.NullableClient to instances of NullableServer.
// Don't use this type directly, use NewNullableServerTransport instead.
type NullableServerTransport struct {
	srv                              *NullableServer
	trMu                             sync.Mutex
	trNullableBytesServer            *NullableBytesServerTransport
	trNullableCollectionsByteServer  *NullableCollectionsByteServerTransport
	trNullableCollectionsModelServer *NullableCollectionsModelServerTransport
	trNullableDatetimeServer         *NullableDatetimeServerTransport
	trNullableDurationServer         *NullableDurationServerTransport
	trNullableStringServer           *NullableStringServerTransport
}

// Do implements the policy.Transporter interface for NullableServerTransport.
func (n *NullableServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (n *NullableServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "NullableBytesClient":
		initServer(&n.trMu, &n.trNullableBytesServer, func() *NullableBytesServerTransport {
			return NewNullableBytesServerTransport(&n.srv.NullableBytesServer)
		})
		resp, err = n.trNullableBytesServer.Do(req)
	case "NullableCollectionsByteClient":
		initServer(&n.trMu, &n.trNullableCollectionsByteServer, func() *NullableCollectionsByteServerTransport {
			return NewNullableCollectionsByteServerTransport(&n.srv.NullableCollectionsByteServer)
		})
		resp, err = n.trNullableCollectionsByteServer.Do(req)
	case "NullableCollectionsModelClient":
		initServer(&n.trMu, &n.trNullableCollectionsModelServer, func() *NullableCollectionsModelServerTransport {
			return NewNullableCollectionsModelServerTransport(&n.srv.NullableCollectionsModelServer)
		})
		resp, err = n.trNullableCollectionsModelServer.Do(req)
	case "NullableDatetimeClient":
		initServer(&n.trMu, &n.trNullableDatetimeServer, func() *NullableDatetimeServerTransport {
			return NewNullableDatetimeServerTransport(&n.srv.NullableDatetimeServer)
		})
		resp, err = n.trNullableDatetimeServer.Do(req)
	case "NullableDurationClient":
		initServer(&n.trMu, &n.trNullableDurationServer, func() *NullableDurationServerTransport {
			return NewNullableDurationServerTransport(&n.srv.NullableDurationServer)
		})
		resp, err = n.trNullableDurationServer.Do(req)
	case "NullableStringClient":
		initServer(&n.trMu, &n.trNullableStringServer, func() *NullableStringServerTransport {
			return NewNullableStringServerTransport(&n.srv.NullableStringServer)
		})
		resp, err = n.trNullableStringServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}
