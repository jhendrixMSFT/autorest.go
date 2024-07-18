// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlbasicgroup_test

import (
	"context"
	"testing"
	"xmlbasicgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestBasicClient_GetModelWithArrayOfModel(t *testing.T) {
	client, err := xmlbasicgroup.NewBasicClient(nil)
	require.NoError(t, err)
	resp, err := client.GetModelWithArrayOfModel(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlbasicgroup.ModelWithArrayOfModel{
		Items: []xmlbasicgroup.SimpleModel{
			{
				Name: to.Ptr("foo"),
				Age:  to.Ptr[int32](123),
			},
			{
				Name: to.Ptr("bar"),
				Age:  to.Ptr[int32](456),
			},
		},
	}, resp.ModelWithArrayOfModel)
}

func TestBasicClient_GetModelWithOptionalField(t *testing.T) {
	client, err := xmlbasicgroup.NewBasicClient(nil)
	require.NoError(t, err)
	resp, err := client.GetModelWithOptionalField(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlbasicgroup.ModelWithOptionalField{
		Item: to.Ptr("widget"),
	}, resp.ModelWithOptionalField)
}

func TestBasicClient_GetModelWithSimpleArrays(t *testing.T) {
	client, err := xmlbasicgroup.NewBasicClient(nil)
	require.NoError(t, err)
	resp, err := client.GetModelWithSimpleArrays(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlbasicgroup.ModelWithSimpleArrays{
		Colors: []string{"red", "green", "blue"},
		Counts: []int32{1, 2},
	}, resp.ModelWithSimpleArrays)
}

func TestBasicClient_GetSimpleModel(t *testing.T) {
	client, err := xmlbasicgroup.NewBasicClient(nil)
	require.NoError(t, err)
	resp, err := client.GetSimpleModel(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, xmlbasicgroup.SimpleModel{
		Age:  to.Ptr[int32](123),
		Name: to.Ptr("foo"),
	}, resp.SimpleModel)
}

func TestBasicClient_PutModelWithArrayOfModel(t *testing.T) {
	client, err := xmlbasicgroup.NewBasicClient(nil)
	require.NoError(t, err)
	resp, err := client.PutModelWithArrayOfModel(context.Background(), xmlbasicgroup.ModelWithArrayOfModel{
		Items: []xmlbasicgroup.SimpleModel{
			{
				Name: to.Ptr("foo"),
				Age:  to.Ptr[int32](123),
			},
			{
				Name: to.Ptr("bar"),
				Age:  to.Ptr[int32](456),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestBasicClient_PutModelWithOptionalField(t *testing.T) {
	client, err := xmlbasicgroup.NewBasicClient(nil)
	require.NoError(t, err)
	resp, err := client.PutModelWithOptionalField(context.Background(), xmlbasicgroup.ModelWithOptionalField{
		Item: to.Ptr("widget"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestBasicClient_PutModelWithSimpleArrays(t *testing.T) {
	client, err := xmlbasicgroup.NewBasicClient(nil)
	require.NoError(t, err)
	resp, err := client.PutSimpleModel(context.Background(), xmlbasicgroup.SimpleModel{
		Age:  to.Ptr[int32](123),
		Name: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestBasicClient_PutSimpleModel(t *testing.T) {
	client, err := xmlbasicgroup.NewBasicClient(nil)
	require.NoError(t, err)
	resp, err := client.PutSimpleModel(context.Background(), xmlbasicgroup.SimpleModel{
		Age:  to.Ptr[int32](123),
		Name: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
