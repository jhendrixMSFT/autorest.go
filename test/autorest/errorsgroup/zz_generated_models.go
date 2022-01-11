//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

type Animal struct {
	AniType *string `json:"aniType,omitempty"`
}

type AnimalNotFound struct {
	// REQUIRED
	WhatNotFound *string `json:"whatNotFound,omitempty"`
	Name         *string `json:"name,omitempty"`
	Reason       *string `json:"reason,omitempty"`
	SomeBaseProp *string `json:"someBaseProp,omitempty"`
}

// GetNotFoundErrorBase implements the NotFoundErrorBaseClassification interface for type AnimalNotFound.
func (a *AnimalNotFound) GetNotFoundErrorBase() *NotFoundErrorBase {
	return &NotFoundErrorBase{
		Reason:       a.Reason,
		WhatNotFound: a.WhatNotFound,
		SomeBaseProp: a.SomeBaseProp,
	}
}

// MarshalJSON implements the json.Marshaller interface for type AnimalNotFound.
func (a AnimalNotFound) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", a.Name)
	populate(objectMap, "reason", a.Reason)
	populate(objectMap, "someBaseProp", a.SomeBaseProp)
	objectMap["whatNotFound"] = "AnimalNotFound"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AnimalNotFound.
func (a *AnimalNotFound) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, &a.Name)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, &a.Reason)
			delete(rawMsg, key)
		case "someBaseProp":
			err = unpopulate(val, &a.SomeBaseProp)
			delete(rawMsg, key)
		case "whatNotFound":
			err = unpopulate(val, &a.WhatNotFound)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type BaseError struct {
	SomeBaseProp *string `json:"someBaseProp,omitempty"`
}

type LinkNotFound struct {
	// REQUIRED
	WhatNotFound   *string `json:"whatNotFound,omitempty"`
	Reason         *string `json:"reason,omitempty"`
	SomeBaseProp   *string `json:"someBaseProp,omitempty"`
	WhatSubAddress *string `json:"whatSubAddress,omitempty"`
}

// GetNotFoundErrorBase implements the NotFoundErrorBaseClassification interface for type LinkNotFound.
func (l *LinkNotFound) GetNotFoundErrorBase() *NotFoundErrorBase {
	return &NotFoundErrorBase{
		Reason:       l.Reason,
		WhatNotFound: l.WhatNotFound,
		SomeBaseProp: l.SomeBaseProp,
	}
}

// MarshalJSON implements the json.Marshaller interface for type LinkNotFound.
func (l LinkNotFound) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "reason", l.Reason)
	populate(objectMap, "someBaseProp", l.SomeBaseProp)
	objectMap["whatNotFound"] = "InvalidResourceLink"
	populate(objectMap, "whatSubAddress", l.WhatSubAddress)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LinkNotFound.
func (l *LinkNotFound) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "reason":
			err = unpopulate(val, &l.Reason)
			delete(rawMsg, key)
		case "someBaseProp":
			err = unpopulate(val, &l.SomeBaseProp)
			delete(rawMsg, key)
		case "whatNotFound":
			err = unpopulate(val, &l.WhatNotFound)
			delete(rawMsg, key)
		case "whatSubAddress":
			err = unpopulate(val, &l.WhatSubAddress)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// NotFoundErrorBaseClassification provides polymorphic access to related types.
// Call the interface's GetNotFoundErrorBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AnimalNotFound, *LinkNotFound, *NotFoundErrorBase
type NotFoundErrorBaseClassification interface {
	// GetNotFoundErrorBase returns the NotFoundErrorBase content of the underlying type.
	GetNotFoundErrorBase() *NotFoundErrorBase
}

type NotFoundErrorBase struct {
	// REQUIRED
	WhatNotFound *string `json:"whatNotFound,omitempty"`
	Reason       *string `json:"reason,omitempty"`
	SomeBaseProp *string `json:"someBaseProp,omitempty"`
}

// GetNotFoundErrorBase implements the NotFoundErrorBaseClassification interface for type NotFoundErrorBase.
func (n *NotFoundErrorBase) GetNotFoundErrorBase() *NotFoundErrorBase { return n }

type Pet struct {
	AniType *string `json:"aniType,omitempty"`

	// READ-ONLY; Gets the Pet by id.
	Name *string `json:"name,omitempty" azure:"ro"`
}

type PetAction struct {
	// action feedback
	ActionResponse *string `json:"actionResponse,omitempty"`
}

// PetActionErrorClassification provides polymorphic access to related types.
// Call the interface's GetPetActionError() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PetActionError, *PetHungryOrThirstyError, *PetSadError
type PetActionErrorClassification interface {
	// GetPetActionError returns the PetActionError content of the underlying type.
	GetPetActionError() *PetActionError
}

type PetActionError struct {
	// REQUIRED
	ErrorType *string `json:"errorType,omitempty"`

	// action feedback
	ActionResponse *string `json:"actionResponse,omitempty"`

	// the error message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// GetPetActionError implements the PetActionErrorClassification interface for type PetActionError.
func (p *PetActionError) GetPetActionError() *PetActionError { return p }

// PetClientDoSomethingOptions contains the optional parameters for the PetClient.DoSomething method.
type PetClientDoSomethingOptions struct {
	// placeholder for future optional parameters
}

// PetClientGetPetByIDOptions contains the optional parameters for the PetClient.GetPetByID method.
type PetClientGetPetByIDOptions struct {
	// placeholder for future optional parameters
}

// PetClientHasModelsParamOptions contains the optional parameters for the PetClient.HasModelsParam method.
type PetClientHasModelsParamOptions struct {
	// Make sure model deserialization doesn't conflict with this param name, which has input name 'models'. Use client default
	// value in call
	Models *string
}

type PetHungryOrThirstyError struct {
	// REQUIRED
	ErrorType *string `json:"errorType,omitempty"`

	// action feedback
	ActionResponse *string `json:"actionResponse,omitempty"`

	// the error message
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// is the pet hungry or thirsty or both
	HungryOrThirsty *string `json:"hungryOrThirsty,omitempty"`

	// why is the pet sad
	Reason *string `json:"reason,omitempty"`
}

// GetPetActionError implements the PetActionErrorClassification interface for type PetHungryOrThirstyError.
func (p *PetHungryOrThirstyError) GetPetActionError() *PetActionError {
	return &PetActionError{
		ErrorType:      p.ErrorType,
		ErrorMessage:   p.ErrorMessage,
		ActionResponse: p.ActionResponse,
	}
}

// GetPetSadError implements the PetSadErrorClassification interface for type PetHungryOrThirstyError.
func (p *PetHungryOrThirstyError) GetPetSadError() *PetSadError {
	return &PetSadError{
		Reason:         p.Reason,
		ErrorType:      p.ErrorType,
		ErrorMessage:   p.ErrorMessage,
		ActionResponse: p.ActionResponse,
	}
}

// MarshalJSON implements the json.Marshaller interface for type PetHungryOrThirstyError.
func (p PetHungryOrThirstyError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionResponse", p.ActionResponse)
	populate(objectMap, "errorMessage", p.ErrorMessage)
	objectMap["errorType"] = "PetHungryOrThirstyError"
	populate(objectMap, "hungryOrThirsty", p.HungryOrThirsty)
	populate(objectMap, "reason", p.Reason)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetHungryOrThirstyError.
func (p *PetHungryOrThirstyError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionResponse":
			err = unpopulate(val, &p.ActionResponse)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, &p.ErrorMessage)
			delete(rawMsg, key)
		case "errorType":
			err = unpopulate(val, &p.ErrorType)
			delete(rawMsg, key)
		case "hungryOrThirsty":
			err = unpopulate(val, &p.HungryOrThirsty)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, &p.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetSadErrorClassification provides polymorphic access to related types.
// Call the interface's GetPetSadError() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PetHungryOrThirstyError, *PetSadError
type PetSadErrorClassification interface {
	PetActionErrorClassification
	// GetPetSadError returns the PetSadError content of the underlying type.
	GetPetSadError() *PetSadError
}

type PetSadError struct {
	// REQUIRED
	ErrorType *string `json:"errorType,omitempty"`

	// action feedback
	ActionResponse *string `json:"actionResponse,omitempty"`

	// the error message
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// why is the pet sad
	Reason *string `json:"reason,omitempty"`
}

// GetPetActionError implements the PetActionErrorClassification interface for type PetSadError.
func (p *PetSadError) GetPetActionError() *PetActionError {
	return &PetActionError{
		ErrorType:      p.ErrorType,
		ErrorMessage:   p.ErrorMessage,
		ActionResponse: p.ActionResponse,
	}
}

// GetPetSadError implements the PetSadErrorClassification interface for type PetSadError.
func (p *PetSadError) GetPetSadError() *PetSadError { return p }

// MarshalJSON implements the json.Marshaller interface for type PetSadError.
func (p PetSadError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionResponse", p.ActionResponse)
	populate(objectMap, "errorMessage", p.ErrorMessage)
	objectMap["errorType"] = "PetSadError"
	populate(objectMap, "reason", p.Reason)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetSadError.
func (p *PetSadError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionResponse":
			err = unpopulate(val, &p.ActionResponse)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, &p.ErrorMessage)
			delete(rawMsg, key)
		case "errorType":
			err = unpopulate(val, &p.ErrorType)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, &p.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
