// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package uniongroup_test

import (
	"context"
	"testing"
	"uniongroup"

	"github.com/stretchr/testify/require"
)

func TestUnionMixedTypesClient_Get(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionMixedTypesClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, uniongroup.MixedTypesCases{}, resp.MixedTypesCases)
}

func TestUnionMixedTypesClient_Send(t *testing.T) {}
