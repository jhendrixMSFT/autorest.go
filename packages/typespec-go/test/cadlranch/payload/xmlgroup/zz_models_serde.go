// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package xmlgroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ModelWithArrayOfModel.
func (m ModelWithArrayOfModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "items", m.Items)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ModelWithArrayOfModel.
func (m *ModelWithArrayOfModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "items":
			err = unpopulate(val, "Items", &m.Items)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ModelWithAttributes.
func (m ModelWithAttributes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "enabled", m.Enabled)
	populate(objectMap, "id1", m.Id1)
	populate(objectMap, "id2", m.Id2)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ModelWithAttributes.
func (m *ModelWithAttributes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enabled":
			err = unpopulate(val, "Enabled", &m.Enabled)
			delete(rawMsg, key)
		case "id1":
			err = unpopulate(val, "Id1", &m.Id1)
			delete(rawMsg, key)
		case "id2":
			err = unpopulate(val, "Id2", &m.Id2)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ModelWithOptionalField.
func (m ModelWithOptionalField) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "item", m.Item)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ModelWithOptionalField.
func (m *ModelWithOptionalField) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "item":
			err = unpopulate(val, "Item", &m.Item)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ModelWithRenamedFields.
func (m ModelWithRenamedFields) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "inputData", m.InputData)
	populate(objectMap, "outputData", m.OutputData)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ModelWithRenamedFields.
func (m *ModelWithRenamedFields) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "inputData":
			err = unpopulate(val, "InputData", &m.InputData)
			delete(rawMsg, key)
		case "outputData":
			err = unpopulate(val, "OutputData", &m.OutputData)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ModelWithSimpleArrays.
func (m ModelWithSimpleArrays) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "colors", m.Colors)
	populate(objectMap, "counts", m.Counts)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ModelWithSimpleArrays.
func (m *ModelWithSimpleArrays) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "colors":
			err = unpopulate(val, "Colors", &m.Colors)
			delete(rawMsg, key)
		case "counts":
			err = unpopulate(val, "Counts", &m.Counts)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ModelWithUnwrappedArray.
func (m ModelWithUnwrappedArray) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "colors", m.Colors)
	populate(objectMap, "counts", m.Counts)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ModelWithUnwrappedArray.
func (m *ModelWithUnwrappedArray) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "colors":
			err = unpopulate(val, "Colors", &m.Colors)
			delete(rawMsg, key)
		case "counts":
			err = unpopulate(val, "Counts", &m.Counts)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SimpleModel.
func (s SimpleModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "age", s.Age)
	populate(objectMap, "name", s.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SimpleModel.
func (s *SimpleModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, "Age", &s.Age)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
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
