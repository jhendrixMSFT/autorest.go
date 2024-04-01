// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package lrolegacygroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// LegacyClient - Illustrates bodies templated with Azure Core with long-running operation
// Don't use this type directly, use a constructor function instead.
type LegacyClient struct {
	internal *azcore.Client
}

// NewLegacyCreateResourcePollViaOperationLocationClient creates a new instance of [LegacyCreateResourcePollViaOperationLocationClient].
func (client *LegacyClient) NewLegacyCreateResourcePollViaOperationLocationClient() *LegacyCreateResourcePollViaOperationLocationClient {
	return &LegacyCreateResourcePollViaOperationLocationClient{
		internal: client.internal,
	}
}
