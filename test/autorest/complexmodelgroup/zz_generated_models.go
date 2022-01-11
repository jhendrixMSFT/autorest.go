//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexmodelgroup

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

type CatalogArray struct {
	// Array of products
	ProductArray []*Product `json:"productArray,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CatalogArray.
func (c CatalogArray) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "productArray", c.ProductArray)
	return json.Marshal(objectMap)
}

type CatalogArrayOfDictionary struct {
	// Array of dictionary of products
	ProductArrayOfDictionary []map[string]*Product `json:"productArrayOfDictionary,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CatalogArrayOfDictionary.
func (c CatalogArrayOfDictionary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "productArrayOfDictionary", c.ProductArrayOfDictionary)
	return json.Marshal(objectMap)
}

type CatalogDictionary struct {
	// Dictionary of products
	ProductDictionary map[string]*Product `json:"productDictionary,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CatalogDictionary.
func (c CatalogDictionary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "productDictionary", c.ProductDictionary)
	return json.Marshal(objectMap)
}

type CatalogDictionaryOfArray struct {
	// Dictionary of Array of product
	ProductDictionaryOfArray map[string][]*Product `json:"productDictionaryOfArray,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CatalogDictionaryOfArray.
func (c CatalogDictionaryOfArray) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "productDictionaryOfArray", c.ProductDictionaryOfArray)
	return json.Marshal(objectMap)
}

// ComplexModelClientCreateOptions contains the optional parameters for the ComplexModelClient.Create method.
type ComplexModelClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ComplexModelClientListOptions contains the optional parameters for the ComplexModelClient.List method.
type ComplexModelClientListOptions struct {
	// placeholder for future optional parameters
}

// ComplexModelClientUpdateOptions contains the optional parameters for the ComplexModelClient.Update method.
type ComplexModelClientUpdateOptions struct {
	// placeholder for future optional parameters
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Product - The product documentation.
type Product struct {
	// Capacity of product. For example, 4 people.
	Capacity *string `json:"capacity,omitempty"`

	// Description of product.
	Description *string `json:"description,omitempty"`

	// Display name of product.
	DisplayName *string `json:"display_name,omitempty"`

	// Image URL representing the product.
	Image *string `json:"image,omitempty"`

	// Unique identifier representing a specific product for a given latitude & longitude. For example, uberX in San Francisco
	// will have a different product_id than uberX in Los Angeles.
	ProductID *string `json:"product_id,omitempty"`
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
