// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package uniongroup_test

import (
	"context"
	"testing"
	"uniongroup"

	"github.com/stretchr/testify/require"
)

func TestUnionFloatsOnlyClient_Get(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionFloatsOnlyClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.Equal(t, uniongroup.UnionFloatsOnlyClientGetCases22, *resp.Value)
}

func TestUnionFloatsOnlyClient_Send(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionFloatsOnlyClient().Send(context.Background(), uniongroup.UnionFloatsOnlyClientGetCases22, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
