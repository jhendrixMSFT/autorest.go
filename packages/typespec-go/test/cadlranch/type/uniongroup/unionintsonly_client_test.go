// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package uniongroup_test

import (
	"context"
	"testing"
	"uniongroup"

	"github.com/stretchr/testify/require"
)

func TestUnionIntsOnlyClient_Get(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionIntsOnlyClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.Equal(t, uniongroup.UnionIntsOnlyClientGetCases2, *resp.Value)
}

func TestUnionIntsOnlyClient_Send(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionIntsOnlyClient().Send(context.Background(), uniongroup.UnionIntsOnlyClientGetCases2, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
