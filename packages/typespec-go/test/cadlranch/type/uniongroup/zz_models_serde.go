// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package uniongroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type EnumsOnlyCases.
func (e EnumsOnlyCases) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "lr", e.Lr)
	populate(objectMap, "ud", e.Ud)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnumsOnlyCases.
func (e *EnumsOnlyCases) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "lr":
			err = unpopulate(val, "Lr", &e.Lr)
			delete(rawMsg, key)
		case "ud":
			err = unpopulate(val, "Ud", &e.Ud)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GetResponse2.
func (g GetResponse2) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "prop", g.Prop)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GetResponse2.
func (g *GetResponse2) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "prop":
			err = unpopulate(val, "Prop", &g.Prop)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GetResponse4.
func (g GetResponse4) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "prop", json.RawMessage(g.Prop))
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GetResponse4.
func (g *GetResponse4) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "prop":
			if string(val) != "null" {
				g.Prop = val
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GetResponse7.
func (g GetResponse7) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "prop", g.Prop)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GetResponse7.
func (g *GetResponse7) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "prop":
			err = unpopulate(val, "Prop", &g.Prop)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GetResponse8.
func (g GetResponse8) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "prop", g.Prop)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GetResponse8.
func (g *GetResponse8) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "prop":
			err = unpopulate(val, "Prop", &g.Prop)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GetResponse9.
func (g GetResponse9) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "prop", g.Prop)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GetResponse9.
func (g *GetResponse9) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "prop":
			err = unpopulate(val, "Prop", &g.Prop)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MixedLiteralsCases.
func (m MixedLiteralsCases) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "booleanLiteral", m.BooleanLiteral)
	populate(objectMap, "floatLiteral", m.FloatLiteral)
	populate(objectMap, "intLiteral", m.IntLiteral)
	populate(objectMap, "stringLiteral", m.StringLiteral)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MixedLiteralsCases.
func (m *MixedLiteralsCases) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "booleanLiteral":
			err = unpopulate(val, "BooleanLiteral", &m.BooleanLiteral)
			delete(rawMsg, key)
		case "floatLiteral":
			err = unpopulate(val, "FloatLiteral", &m.FloatLiteral)
			delete(rawMsg, key)
		case "intLiteral":
			err = unpopulate(val, "IntLiteral", &m.IntLiteral)
			delete(rawMsg, key)
		case "stringLiteral":
			err = unpopulate(val, "StringLiteral", &m.StringLiteral)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MixedTypesCases.
func (m MixedTypesCases) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "array", m.Array)
	populate(objectMap, "boolean", json.RawMessage(m.Boolean))
	populate(objectMap, "int", json.RawMessage(m.Int))
	populate(objectMap, "literal", json.RawMessage(m.Literal))
	populate(objectMap, "model", json.RawMessage(m.Model))
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MixedTypesCases.
func (m *MixedTypesCases) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "array":
			err = unpopulate(val, "Array", &m.Array)
			delete(rawMsg, key)
		case "boolean":
			if string(val) != "null" {
				m.Boolean = val
			}
			delete(rawMsg, key)
		case "int":
			if string(val) != "null" {
				m.Int = val
			}
			delete(rawMsg, key)
		case "literal":
			if string(val) != "null" {
				m.Literal = val
			}
			delete(rawMsg, key)
		case "model":
			if string(val) != "null" {
				m.Model = val
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StringAndArrayCases.
func (s StringAndArrayCases) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "array", json.RawMessage(s.Array))
	populate(objectMap, "string", json.RawMessage(s.String))
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StringAndArrayCases.
func (s *StringAndArrayCases) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "array":
			if string(val) != "null" {
				s.Array = val
			}
			delete(rawMsg, key)
		case "string":
			if string(val) != "null" {
				s.String = val
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}