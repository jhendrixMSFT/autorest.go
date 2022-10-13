//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armdataboxedge_test

import (
	"armdataboxedge"
	"encoding/json"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

var addon = armdataboxedge.Addon{
	Kind: to.Ptr(armdataboxedge.AddonTypeArcForKubernetes),
	ID:   to.Ptr("/some/kind/of/resource/that/should/be/a/long/string/longer/than/this/most/likely"),
	Name: to.Ptr("FooBarBaz"),
	SystemData: &armdataboxedge.SystemData{
		CreatedAt:          to.Ptr(time.Now().UTC()),
		CreatedBy:          to.Ptr("jhendrix"),
		CreatedByType:      to.Ptr(armdataboxedge.CreatedByTypeUser),
		LastModifiedAt:     to.Ptr(time.Now().UTC()),
		LastModifiedBy:     to.Ptr("jhendrix"),
		LastModifiedByType: to.Ptr(armdataboxedge.CreatedByTypeUser),
	},
	Type: to.Ptr("something/something"),
}

const asJSON = `{"kind":"ArcForKubernetes","id":"/some/kind/of/resource/that/should/be/a/long/string/longer/than/this/most/likely","name":"FooBarBaz","systemData":{"createdAt":"2022-10-13T13:22:46.3952267-07:00","createdBy":"jhendrix","createdByType":"User","lastModifiedAt":"2022-10-13T13:22:46.3952267-07:00","lastModifiedBy":"jhendrix","lastModifiedByType":"User"},"type":"something/something"}`

// go test -benchmem -bench .

func TestAllFieldsPopulated(t *testing.T) {
	var v armdataboxedge.Addon
	require.NoError(t, json.Unmarshal([]byte(asJSON), &v))
	require.NotNil(t, v.ID)
	require.NotNil(t, v.Kind)
	require.NotNil(t, v.Name)
	require.NotNil(t, v.SystemData)
	require.NotNil(t, v.SystemData.CreatedAt)
	require.NotNil(t, v.SystemData.CreatedBy)
	require.NotNil(t, v.SystemData.CreatedByType)
	require.NotNil(t, v.SystemData.LastModifiedAt)
	require.NotNil(t, v.SystemData.LastModifiedBy)
	require.NotNil(t, v.SystemData.LastModifiedByType)
	require.NotNil(t, v.Type)
}

func TestUnmarshalJSON2(t *testing.T) {
	v := armdataboxedge.Addon{}
	require.NoError(t, v.UnmarshalJSON2([]byte(asJSON)))
	require.NotNil(t, v.ID)
	require.NotNil(t, v.Kind)
	require.NotNil(t, v.Name)
	require.NotNil(t, v.SystemData)
	require.NotNil(t, v.SystemData.CreatedAt)
	require.NotNil(t, v.SystemData.CreatedBy)
	require.NotNil(t, v.SystemData.CreatedByType)
	require.NotNil(t, v.SystemData.LastModifiedAt)
	require.NotNil(t, v.SystemData.LastModifiedBy)
	require.NotNil(t, v.SystemData.LastModifiedByType)
	require.NotNil(t, v.Type)
}

func BenchmarkCurrentMarshaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(addon); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrototypeMarshaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := addon.MarshalJSON2(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCurrentUnmarshaller(b *testing.B) {
	var v armdataboxedge.Addon
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal([]byte(asJSON), &v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrototypeUnmarshaller(b *testing.B) {
	var v armdataboxedge.Addon
	for i := 0; i < b.N; i++ {
		if err := v.UnmarshalJSON2([]byte(asJSON)); err != nil {
			b.Fatal(err)
		}
	}
}
