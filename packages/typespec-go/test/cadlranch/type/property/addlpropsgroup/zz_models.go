//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

// The model extends from Record<float32> type.
type ExtendsFloatAdditionalProperties struct {
	// REQUIRED; The id property
	ID                   *float32
	AdditionalProperties map[string]*float32
}

// The model extends from Record<ModelForRecord> type.
type ExtendsModelAdditionalProperties struct {
	AdditionalProperties map[string]*ModelForRecord
}

// The model extends from Record<ModelForRecord[]> type.
type ExtendsModelArrayAdditionalProperties struct {
	AdditionalProperties map[string][]*ModelForRecord
}

// The model extends from Record<string> type.
type ExtendsStringAdditionalProperties struct {
	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]*string
}

// The model extends from Record<unknown> type.
type ExtendsUnknownAdditionalProperties struct {
	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any
}

// The model is from Record<float32> type.
type IsFloatAdditionalProperties struct {
	// REQUIRED; The id property
	ID                   *float32
	AdditionalProperties map[string]*float32
}

// The model is from Record<ModelForRecord> type.
type IsModelAdditionalProperties struct {
	AdditionalProperties map[string]*ModelForRecord
}

// The model is from Record<ModelForRecord[]> type.
type IsModelArrayAdditionalProperties struct {
	AdditionalProperties map[string][]*ModelForRecord
}

// The model is from Record<string> type.
type IsStringAdditionalProperties struct {
	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]*string
}

// The model is from Record<unknown> type.
type IsUnknownAdditionalProperties struct {
	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any
}

// model for record
type ModelForRecord struct {
	// REQUIRED; The state property
	State *string
}