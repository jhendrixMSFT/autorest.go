// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package uniongroup_test

import (
	"context"
	"encoding/json"
	"testing"
	"uniongroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestUnionMixedTypesClient_Get(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnionMixedTypesClient().Get(context.Background(), nil)
	require.NoError(t, err)

	// marshal values used for comparison
	boolean, err := json.Marshal(true)
	require.NoError(t, err)
	intv, err := json.Marshal(2)
	require.NoError(t, err)
	literal, err := json.Marshal("a")
	require.NoError(t, err)
	model, err := json.Marshal(uniongroup.Cat{
		Name: to.Ptr("test"),
	})
	require.NoError(t, err)
	array := []json.RawMessage{}
	array = append(array, model)
	array = append(array, literal)
	array = append(array, intv)
	array = append(array, boolean)
	arrayv, err := json.Marshal(array)
	require.NoError(t, err)
	require.Equal(t, uniongroup.MixedTypesCases{
		Array:   arrayv,
		Boolean: boolean,
		Int:     intv,
		Literal: literal,
		Model:   model,
	}, resp.MixedTypesCases)
}

func TestUnionMixedTypesClient_Send(t *testing.T) {
	client, err := uniongroup.NewUnionClient(nil)
	require.NoError(t, err)

	// marshal values to send
	boolean, err := json.Marshal(true)
	require.NoError(t, err)
	intv, err := json.Marshal(2)
	require.NoError(t, err)
	literal, err := json.Marshal("a")
	require.NoError(t, err)
	model, err := json.Marshal(uniongroup.Cat{
		Name: to.Ptr("test"),
	})
	require.NoError(t, err)
	array := []json.RawMessage{}
	array = append(array, model)
	array = append(array, literal)
	array = append(array, intv)
	array = append(array, boolean)
	arrayv, err := json.Marshal(array)
	require.NoError(t, err)

	resp, err := client.NewUnionMixedTypesClient().Send(context.Background(), uniongroup.MixedTypesCases{
		Array:   arrayv,
		Boolean: boolean,
		Int:     intv,
		Literal: literal,
		Model:   model,
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
