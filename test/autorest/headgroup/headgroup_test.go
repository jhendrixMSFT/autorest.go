// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newHTTPSuccessClient(t *testing.T) *HTTPSuccessClient {
	client, err := NewHTTPSuccessClient(nil)
	require.NoError(t, err)
	return client
}

func NewHTTPSuccessClient(options *azcore.ClientOptions) (*HTTPSuccessClient, error) {
	client, err := azcore.NewClient("headgroup.HTTPSuccessClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &HTTPSuccessClient{internal: client}, nil
}

// Head200 - Return 200 status code if successful
func TestHead200(t *testing.T) {
	client := newHTTPSuccessClient(t)
	resp, err := client.Head200(context.Background(), nil)
	require.NoError(t, err)
	if !resp.Success {
		t.Fatal("expected success")
	}
}

// Head204 - Return 204 status code if successful
func TestHead204(t *testing.T) {
	client := newHTTPSuccessClient(t)
	resp, err := client.Head204(context.Background(), nil)
	require.NoError(t, err)
	if !resp.Success {
		t.Fatal("expected success")
	}
}

// Head404 - Return 404 status code if successful
func TestHead404(t *testing.T) {
	client := newHTTPSuccessClient(t)
	resp, err := client.Head404(context.Background(), nil)
	require.NoError(t, err)
	if resp.Success {
		t.Fatal("expected non-success")
	}
}
