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

func TestUnionEnumsOnlyClient_Get(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionEnumsOnlyClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, uniongroup.EnumsOnlyCases{
		Lr: to.Ptr(uniongroup.EnumsOnlyCasesLrRight),
		Ud: to.Ptr(uniongroup.EnumsOnlyCasesUdUp),
	}, resp.EnumsOnlyCases)
}

func TestUnionEnumsOnlyClient_Send(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionEnumsOnlyClient().Send(context.Background(), uniongroup.EnumsOnlyCases{
		Lr: to.Ptr(uniongroup.EnumsOnlyCasesLrRight),
		Ud: to.Ptr(uniongroup.EnumsOnlyCasesUdUp),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
