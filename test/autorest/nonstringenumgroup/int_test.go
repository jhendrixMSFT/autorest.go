// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newIntClient(t *testing.T) *IntClient {
	client, err := NewIntClient(nil)
	require.NoError(t, err)
	return client
}

func NewIntClient(options *azcore.ClientOptions) (*IntClient, error) {
	client, err := azcore.NewClient("nonstringenumgroup.IntClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &IntClient{internal: client}, nil
}

// Get - Get an int enum
func TestIntGet(t *testing.T) {
	client := newIntClient(t)
	result, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr(IntEnumFourHundredTwentyNine)); r != "" {
		t.Fatal(r)
	}
}

// Put - Put an int enum
func TestIntPut(t *testing.T) {
	client := newIntClient(t)
	result, err := client.Put(context.Background(), IntEnumTwoHundred, nil)
	require.NoError(t, err)
	if *result.Value != "Nice job posting an int enum" {
		t.Fatalf("unexpected value %s", *result.Value)
	}
}
