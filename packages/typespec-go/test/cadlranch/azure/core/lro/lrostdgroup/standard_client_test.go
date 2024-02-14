//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrostdgroup_test

import (
	"context"
	"lrostdgroup"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestStandardClient_BeginCreateOrReplace(t *testing.T) {
	client, err := lrostdgroup.NewStandardClient(nil)
	require.NoError(t, err)
	poller, err := client.BeginCreateOrReplace(context.Background(), lrostdgroup.User{
		Name: to.Ptr("madge"),
		Role: to.Ptr("contributor"),
	}, "madge", nil)
	require.NoError(t, err)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: time.Second,
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestStandardClient_BeginDelete(t *testing.T) {}

func TestStandardClient_BeginExport(t *testing.T) {}
