// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package azurepagegroup

const host = "http://localhost:3000"

// ListItemInputExtensibleEnum - An extensible enum input parameter.
type ListItemInputExtensibleEnum string

const (
	// ListItemInputExtensibleEnumFirst - The first enum value.
	ListItemInputExtensibleEnumFirst ListItemInputExtensibleEnum = "First"
	// ListItemInputExtensibleEnumSecond - The second enum value.
	ListItemInputExtensibleEnumSecond ListItemInputExtensibleEnum = "Second"
)

// PossibleListItemInputExtensibleEnumValues returns the possible values for the ListItemInputExtensibleEnum const type.
func PossibleListItemInputExtensibleEnumValues() []ListItemInputExtensibleEnum {
	return []ListItemInputExtensibleEnum{
		ListItemInputExtensibleEnumFirst,
		ListItemInputExtensibleEnumSecond,
	}
}