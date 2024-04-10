// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package renamedopgroup

const (
	moduleName    = "renamedopgroup"
	moduleVersion = "v0.1.0"
)

type ClientType string

const (
	ClientTypeDefault           ClientType = "default"
	ClientTypeMultiClient       ClientType = "multi-client"
	ClientTypeRenamedOperation  ClientType = "renamed-operation"
	ClientTypeTwoOperationGroup ClientType = "two-operation-group"
)

// PossibleClientTypeValues returns the possible values for the ClientType const type.
func PossibleClientTypeValues() []ClientType {
	return []ClientType{
		ClientTypeDefault,
		ClientTypeMultiClient,
		ClientTypeRenamedOperation,
		ClientTypeTwoOperationGroup,
	}
}
