//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package recursivegroup

// element
type Element struct {
	Extension []Extension
}

// extension
type Extension struct {
	// REQUIRED
	Level     *int32
	Extension []Extension
}
