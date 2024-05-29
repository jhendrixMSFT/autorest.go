// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package uniongroup_test

import (
	"context"
	"testing"
	"uniongroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestUnionMixedLiteralsClient_Get(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionMixedLiteralsClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, uniongroup.MixedLiteralsCases{
		BooleanLiteral: to.Ptr(true),
		FloatLiteral:   to.Ptr[float32](3.3),
		IntLiteral:     to.Ptr[int32](2),
		StringLiteral:  to.Ptr("a"),
	}, resp.MixedLiteralsCases)
}

func TestUnionMixedLiteralsClient_Send(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionMixedLiteralsClient().Send(context.Background(), uniongroup.MixedLiteralsCases{
		BooleanLiteral: to.Ptr(true),
		FloatLiteral:   to.Ptr[float32](3.3),
		IntLiteral:     to.Ptr[int32](2),
		StringLiteral:  to.Ptr("a"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
