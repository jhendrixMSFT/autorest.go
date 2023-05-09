//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package additionalpropsgroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CatAPTrue.
func (c CatAPTrue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "friendly", c.Friendly)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "status", c.Status)
	if c.AdditionalProperties != nil {
		for key, val := range c.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CatAPTrue.
func (c *CatAPTrue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "friendly":
			err = unpopulate(val, "Friendly", &c.Friendly)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &c.Status)
			delete(rawMsg, key)
		default:
			if c.AdditionalProperties == nil {
				c.AdditionalProperties = map[string]any{}
			}
			if val != nil {
				var aux any
				err = json.Unmarshal(val, &aux)
				c.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PetAPInProperties.
func (p PetAPInProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalProperties", p.AdditionalProperties)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPInProperties.
func (p *PetAPInProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalProperties":
			err = unpopulate(val, "AdditionalProperties", &p.AdditionalProperties)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &p.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PetAPInPropertiesWithAPString.
func (p PetAPInPropertiesWithAPString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalProperties", p.AdditionalProperties1)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "@odata.location", p.ODataLocation)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPInPropertiesWithAPString.
func (p *PetAPInPropertiesWithAPString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalProperties":
			err = unpopulate(val, "AdditionalProperties1", &p.AdditionalProperties1)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "@odata.location":
			err = unpopulate(val, "ODataLocation", &p.ODataLocation)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = map[string]*string{}
			}
			if val != nil {
				var aux string
				err = json.Unmarshal(val, &aux)
				p.AdditionalProperties[key] = &aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PetAPObject.
func (p PetAPObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPObject.
func (p *PetAPObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = map[string]any{}
			}
			if val != nil {
				var aux any
				err = json.Unmarshal(val, &aux)
				p.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PetAPString.
func (p PetAPString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPString.
func (p *PetAPString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = map[string]*string{}
			}
			if val != nil {
				var aux string
				err = json.Unmarshal(val, &aux)
				p.AdditionalProperties[key] = &aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PetAPTrue.
func (p PetAPTrue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPTrue.
func (p *PetAPTrue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = map[string]any{}
			}
			if val != nil {
				var aux any
				err = json.Unmarshal(val, &aux)
				p.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
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
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}