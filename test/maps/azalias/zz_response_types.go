//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azalias

// ClientCreateResponse contains the response from method Client.Create.
type ClientCreateResponse struct {
	AliasesCreateResponse
	// AccessControlExposeHeaders contains the information returned from the Access-Control-Expose-Headers header response.
	AccessControlExposeHeaders *string
}

// ClientGetScriptResponse contains the response from method Client.GetScript.
type ClientGetScriptResponse struct {
	Value *string
}

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	ListResponse
}

// ClientPolicyAssignmentResponse contains the response from method Client.PolicyAssignment.
type ClientPolicyAssignmentResponse struct {
	PolicyAssignmentProperties
}
