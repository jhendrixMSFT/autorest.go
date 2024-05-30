// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package uniongroup_test

import (
	"context"
	"encoding/json"
	"testing"
	"uniongroup"

	"github.com/stretchr/testify/require"
)

func TestUnionMixedLiteralsClient_Get(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionMixedLiteralsClient().Get(context.Background(), nil)
	require.NoError(t, err)

	// marshal values used for comparison
	boollit, err := json.Marshal(true)
	require.NoError(t, err)
	floatlit, err := json.Marshal(float32(3.3))
	require.NoError(t, err)
	intlit, err := json.Marshal(int32(2))
	require.NoError(t, err)
	stringlit, err := json.Marshal("a")
	require.NoError(t, err)

	require.Equal(t, uniongroup.MixedLiteralsCases{
		BooleanLiteral: boollit,
		FloatLiteral:   floatlit,
		IntLiteral:     intlit,
		StringLiteral:  stringlit,
	}, resp.MixedLiteralsCases)
}

func TestUnionMixedLiteralsClient_Send(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)

	// marshal values used to send
	boollit, err := json.Marshal(true)
	require.NoError(t, err)
	floatlit, err := json.Marshal(float32(3.3))
	require.NoError(t, err)
	intlit, err := json.Marshal(int32(2))
	require.NoError(t, err)
	stringlit, err := json.Marshal("a")
	require.NoError(t, err)

	resp, err := client.NewUnionMixedLiteralsClient().Send(context.Background(), uniongroup.MixedLiteralsCases{
		BooleanLiteral: boollit,
		FloatLiteral:   floatlit,
		IntLiteral:     intlit,
		StringLiteral:  stringlit,
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
