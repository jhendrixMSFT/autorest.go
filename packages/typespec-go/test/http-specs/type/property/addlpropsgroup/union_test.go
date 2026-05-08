// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package addlpropsgroup

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpreadRecordForUnion_Unmarshal(t *testing.T) {
	data := []byte(`{"flag": true, "prop1": "abc", "prop2": 43.125}`)
	var s SpreadRecordForUnion
	err := json.Unmarshal(data, &s)
	require.NoError(t, err)
	require.NotNil(t, s.Flag)
	require.True(t, *s.Flag)
	require.Len(t, s.AdditionalProperties, 2)

	// prop1 should be string
	v1, ok := s.AdditionalProperties["prop1"].(*StringOrFloat32ValueString)
	require.True(t, ok)
	require.Equal(t, "abc", v1.Value)

	// prop2 should be float32
	v2, ok := s.AdditionalProperties["prop2"].(*StringOrFloat32ValueFloat32)
	require.True(t, ok)
	require.InDelta(t, float32(43.125), v2.Value, 0.001)
}

func TestSpreadRecordForUnion_Marshal(t *testing.T) {
	flag := true
	s := SpreadRecordForUnion{
		Flag: &flag,
		AdditionalProperties: map[string]StringOrFloat32Value{
			"prop1": &StringOrFloat32ValueString{Value: "abc"},
			"prop2": &StringOrFloat32ValueFloat32{Value: 43.125},
		},
	}
	data, err := json.Marshal(s)
	require.NoError(t, err)

	// Unmarshal back to verify round-trip
	var m map[string]any
	err = json.Unmarshal(data, &m)
	require.NoError(t, err)
	require.Equal(t, true, m["flag"])
	require.Equal(t, "abc", m["prop1"])
	require.InDelta(t, 43.125, m["prop2"], 0.001)
}

func TestMultipleSpreadRecord_Unmarshal(t *testing.T) {
	data := []byte(`{"flag": true, "prop1": "abc", "prop2": 43.125}`)
	var m MultipleSpreadRecord
	err := json.Unmarshal(data, &m)
	require.NoError(t, err)
	require.NotNil(t, m.Flag)
	require.True(t, *m.Flag)
	require.Len(t, m.AdditionalProperties, 2)

	v1, ok := m.AdditionalProperties["prop1"].(*StringOrFloat32ValueString)
	require.True(t, ok)
	require.Equal(t, "abc", v1.Value)

	v2, ok := m.AdditionalProperties["prop2"].(*StringOrFloat32ValueFloat32)
	require.True(t, ok)
	require.InDelta(t, float32(43.125), v2.Value, 0.001)
}

func TestSpreadRecordForNonDiscriminatedUnion_Unmarshal(t *testing.T) {
	data := []byte(`{"name": "abc", "prop1": {"kind": "kind0", "fooProp": "abc"}, "prop2": {"kind": "kind1", "start": "2021-01-01T00:00:00Z", "end": "2021-01-02T00:00:00Z"}}`)
	var s SpreadRecordForNonDiscriminatedUnion
	err := json.Unmarshal(data, &s)
	require.NoError(t, err)
	require.NotNil(t, s.Name)
	require.Equal(t, "abc", *s.Name)
	require.Len(t, s.AdditionalProperties, 2)

	// prop1 should be WidgetData0
	w0, ok := s.AdditionalProperties["prop1"].(*WidgetData0)
	require.True(t, ok)
	require.Equal(t, "kind0", *w0.Kind)
	require.Equal(t, "abc", *w0.FooProp)

	// prop2 should be WidgetData1
	w1, ok := s.AdditionalProperties["prop2"].(*WidgetData1)
	require.True(t, ok)
	require.Equal(t, "kind1", *w1.Kind)
	require.Equal(t, "2021-01-01T00:00:00Z", *w1.Start)
	require.Equal(t, "2021-01-02T00:00:00Z", *w1.End)
}

func TestSpreadRecordForNonDiscriminatedUnion_Marshal(t *testing.T) {
	name := "abc"
	kind0 := "kind0"
	fooProp := "abc"
	kind1 := "kind1"
	start := "2021-01-01T00:00:00Z"
	end := "2021-01-02T00:00:00Z"
	s := SpreadRecordForNonDiscriminatedUnion{
		Name: &name,
		AdditionalProperties: map[string]WidgetData0Or1Classification{
			"prop1": &WidgetData0{Kind: &kind0, FooProp: &fooProp},
			"prop2": &WidgetData1{Kind: &kind1, Start: &start, End: &end},
		},
	}
	data, err := json.Marshal(s)
	require.NoError(t, err)

	var m map[string]any
	err = json.Unmarshal(data, &m)
	require.NoError(t, err)
	require.Equal(t, "abc", m["name"])

	prop1 := m["prop1"].(map[string]any)
	require.Equal(t, "kind0", prop1["kind"])
	require.Equal(t, "abc", prop1["fooProp"])

	prop2 := m["prop2"].(map[string]any)
	require.Equal(t, "kind1", prop2["kind"])
	require.Equal(t, "2021-01-01T00:00:00Z", prop2["start"])
	require.Equal(t, "2021-01-02T00:00:00Z", prop2["end"])
}

func TestSpreadRecordForNonDiscriminatedUnion2_Unmarshal(t *testing.T) {
	data := []byte(`{"name": "abc", "prop1": {"kind": "kind1", "start": "2021-01-01T00:00:00Z"}, "prop2": {"kind": "kind1", "start": "2021-01-01T00:00:00Z", "end": "2021-01-02T00:00:00Z"}}`)
	var s SpreadRecordForNonDiscriminatedUnion2
	err := json.Unmarshal(data, &s)
	require.NoError(t, err)
	require.NotNil(t, s.Name)
	require.Equal(t, "abc", *s.Name)
	require.Len(t, s.AdditionalProperties, 2)

	// prop1: has RFC3339 "start" but no "end" -> ambiguous.
	// Per our heuristic: RFC3339 parse succeeds -> WidgetData1
	w1, ok := s.AdditionalProperties["prop1"].(*WidgetData1)
	require.True(t, ok)
	require.Equal(t, "kind1", *w1.Kind)
	require.Equal(t, "2021-01-01T00:00:00Z", *w1.Start)

	// prop2: has "end" -> definitely WidgetData1
	w2, ok := s.AdditionalProperties["prop2"].(*WidgetData1)
	require.True(t, ok)
	require.Equal(t, "kind1", *w2.Kind)
	require.Equal(t, "2021-01-01T00:00:00Z", *w2.Start)
	require.Equal(t, "2021-01-02T00:00:00Z", *w2.End)
}

func TestSpreadRecordForNonDiscriminatedUnion3_Unmarshal(t *testing.T) {
	data := []byte(`{"name": "abc", "prop1": [{"kind": "kind1", "start": "2021-01-01T00:00:00Z"}, {"kind": "kind1", "start": "2021-01-01T00:00:00Z"}], "prop2": {"kind": "kind1", "start": "2021-01-01T00:00:00Z", "end": "2021-01-02T00:00:00Z"}}`)
	var s SpreadRecordForNonDiscriminatedUnion3
	err := json.Unmarshal(data, &s)
	require.NoError(t, err)
	require.NotNil(t, s.Name)
	require.Equal(t, "abc", *s.Name)
	require.Len(t, s.AdditionalProperties, 2)

	// prop1 is an array -> WidgetData2ArrayValue
	arrVal, ok := s.AdditionalProperties["prop1"].(*WidgetData2ArrayValue)
	require.True(t, ok)
	require.Len(t, arrVal.Value, 2)
	require.Equal(t, "kind1", *arrVal.Value[0].Kind)
	require.Equal(t, "2021-01-01T00:00:00Z", *arrVal.Value[0].Start)

	// prop2 is an object -> WidgetData1
	w1, ok := s.AdditionalProperties["prop2"].(*WidgetData1)
	require.True(t, ok)
	require.Equal(t, "kind1", *w1.Kind)
	require.Equal(t, "2021-01-01T00:00:00Z", *w1.Start)
	require.Equal(t, "2021-01-02T00:00:00Z", *w1.End)
}

func TestSpreadRecordForNonDiscriminatedUnion3_Marshal(t *testing.T) {
	name := "abc"
	kind1 := "kind1"
	start := "2021-01-01T00:00:00Z"
	end := "2021-01-02T00:00:00Z"
	s := SpreadRecordForNonDiscriminatedUnion3{
		Name: &name,
		AdditionalProperties: map[string]WidgetData2ArrayOrWidgetData1Classification{
			"prop1": &WidgetData2ArrayValue{
				Value: []*WidgetData2{
					{Kind: &kind1, Start: &start},
					{Kind: &kind1, Start: &start},
				},
			},
			"prop2": &WidgetData1{Kind: &kind1, Start: &start, End: &end},
		},
	}
	data, err := json.Marshal(s)
	require.NoError(t, err)

	var m map[string]any
	err = json.Unmarshal(data, &m)
	require.NoError(t, err)
	require.Equal(t, "abc", m["name"])

	// prop1 should be array
	prop1, ok := m["prop1"].([]any)
	require.True(t, ok)
	require.Len(t, prop1, 2)

	// prop2 should be object
	prop2, ok := m["prop2"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "kind1", prop2["kind"])
}
