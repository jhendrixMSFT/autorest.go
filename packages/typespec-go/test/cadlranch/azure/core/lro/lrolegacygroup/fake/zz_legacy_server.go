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

// LegacyServer is a fake server for instances of the lrolegacygroup.LegacyClient type.
type LegacyServer struct {
	// LegacyCreateResourcePollViaOperationLocationServer contains the fakes for client LegacyCreateResourcePollViaOperationLocationClient
	LegacyCreateResourcePollViaOperationLocationServer LegacyCreateResourcePollViaOperationLocationServer
}

// NewLegacyServerTransport creates a new instance of LegacyServerTransport with the provided implementation.
// The returned LegacyServerTransport instance is connected to an instance of lrolegacygroup.LegacyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLegacyServerTransport(srv *LegacyServer) *LegacyServerTransport {
	return &LegacyServerTransport{srv: srv}
}

// LegacyServerTransport connects instances of lrolegacygroup.LegacyClient to instances of LegacyServer.
// Don't use this type directly, use NewLegacyServerTransport instead.
type LegacyServerTransport struct {
	srv                                                  *LegacyServer
	trMu                                                 sync.Mutex
	trLegacyCreateResourcePollViaOperationLocationServer *LegacyCreateResourcePollViaOperationLocationServerTransport
}

// Do implements the policy.Transporter interface for LegacyServerTransport.
func (l *LegacyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (l *LegacyServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "LegacyCreateResourcePollViaOperationLocationClient":
		initServer(&l.trMu, &l.trLegacyCreateResourcePollViaOperationLocationServer, func() *LegacyCreateResourcePollViaOperationLocationServerTransport {
			return NewLegacyCreateResourcePollViaOperationLocationServerTransport(&l.srv.LegacyCreateResourcePollViaOperationLocationServer)
		})
		resp, err = l.trLegacyCreateResourcePollViaOperationLocationServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}
