// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package addlpropsgroup

import (
	"encoding/json"
	"fmt"
	"time"
)

// marshalStringOrFloat32Value marshals the union value to its JSON representation.
func marshalStringOrFloat32Value(v StringOrFloat32Value) any {
	if v == nil {
		return nil
	}
	switch t := v.(type) {
	case *StringOrFloat32ValueString:
		return t.Value
	case *StringOrFloat32ValueFloat32:
		return t.Value
	default:
		return nil
	}
}

// unmarshalStringOrFloat32Value unmarshals a JSON value into a StringOrFloat32Value.
// Discrimination strategy: JSON token type (string starts with '"', number is digit/'-').
func unmarshalStringOrFloat32Value(raw json.RawMessage) (StringOrFloat32Value, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil
	}
	// A JSON string token always starts with '"'
	if raw[0] == '"' {
		var s string
		if err := json.Unmarshal(raw, &s); err != nil {
			return nil, fmt.Errorf("unmarshalling StringOrFloat32Value as string: %v", err)
		}
		return &StringOrFloat32ValueString{Value: s}, nil
	}
	// Otherwise try float32
	var f float32
	if err := json.Unmarshal(raw, &f); err != nil {
		return nil, fmt.Errorf("unmarshalling StringOrFloat32Value as float32: %v", err)
	}
	return &StringOrFloat32ValueFloat32{Value: f}, nil
}

// unmarshalWidgetData0Or1Classification unmarshals a JSON value into WidgetData0 or WidgetData1.
// Discrimination strategy: field value of "kind" ("kind0" -> WidgetData0, "kind1" -> WidgetData1).
func unmarshalWidgetData0Or1Classification(raw json.RawMessage) (WidgetData0Or1Classification, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil
	}
	var disc struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(raw, &disc); err != nil {
		return nil, fmt.Errorf("unmarshalling WidgetData0Or1Classification discriminator: %v", err)
	}
	switch disc.Kind {
	case "kind0":
		var v WidgetData0
		if err := json.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "kind1":
		var v WidgetData1
		if err := json.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unknown kind value for WidgetData0Or1Classification: %q", disc.Kind)
	}
}

// unmarshalWidgetData2Or1Classification unmarshals a JSON value into WidgetData2 or WidgetData1.
// Discrimination strategy: structural probing. Both have kind:"kind1" and "start".
// WidgetData1 has "start" as utcDateTime (RFC3339) and optional "end".
// WidgetData2 has "start" as plain string, no "end".
// We distinguish by: if "end" is present -> WidgetData1. Otherwise, try parsing "start" as time.
func unmarshalWidgetData2Or1Classification(raw json.RawMessage) (WidgetData2Or1Classification, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil
	}
	var probe map[string]json.RawMessage
	if err := json.Unmarshal(raw, &probe); err != nil {
		return nil, fmt.Errorf("unmarshalling WidgetData2Or1Classification: %v", err)
	}
	// If "end" field is present, it's definitely WidgetData1
	if _, hasEnd := probe["end"]; hasEnd {
		var v WidgetData1
		if err := json.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return &v, nil
	}
	// Try parsing "start" as RFC3339 time. If it parses, it's WidgetData1.
	if startRaw, ok := probe["start"]; ok {
		var s string
		if err := json.Unmarshal(startRaw, &s); err == nil {
			if _, err := time.Parse(time.RFC3339, s); err == nil {
				var v WidgetData1
				if err := json.Unmarshal(raw, &v); err != nil {
					return nil, err
				}
				return &v, nil
			}
		}
	}
	// Fall back to WidgetData2
	var v WidgetData2
	if err := json.Unmarshal(raw, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// unmarshalWidgetData2ArrayOrWidgetData1Classification unmarshals a JSON value into []*WidgetData2 or *WidgetData1.
// Discrimination strategy: JSON token type (array '[' -> WidgetData2[], object '{' -> WidgetData1).
func unmarshalWidgetData2ArrayOrWidgetData1Classification(raw json.RawMessage) (WidgetData2ArrayOrWidgetData1Classification, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil
	}
	// Array starts with '['
	if raw[0] == '[' {
		var arr []*WidgetData2
		if err := json.Unmarshal(raw, &arr); err != nil {
			return nil, fmt.Errorf("unmarshalling WidgetData2ArrayOrWidgetData1Classification as array: %v", err)
		}
		return &WidgetData2ArrayValue{Value: arr}, nil
	}
	// Otherwise it's an object -> WidgetData1
	var v WidgetData1
	if err := json.Unmarshal(raw, &v); err != nil {
		return nil, fmt.Errorf("unmarshalling WidgetData2ArrayOrWidgetData1Classification as WidgetData1: %v", err)
	}
	return &v, nil
}
