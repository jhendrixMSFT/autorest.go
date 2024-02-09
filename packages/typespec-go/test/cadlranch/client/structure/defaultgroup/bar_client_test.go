//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package defaultgroup_test

import (
	"context"
	"defaultgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBarClient_Five(t *testing.T) {
	client, err := defaultgroup.NewBarClient(nil)
	require.NoError(t, err)
	resp, err := client.Five(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestBarClient_Nine(t *testing.T) {
	client, err := defaultgroup.NewBarClient(nil)
	require.NoError(t, err)
	resp, err := client.Nine(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestBarClient_Six(t *testing.T) {
	client, err := defaultgroup.NewBarClient(nil)
	require.NoError(t, err)
	resp, err := client.Six(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
