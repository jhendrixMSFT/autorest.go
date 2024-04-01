// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package corescalargroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// ScalarClient contains the methods for the _Specs_.Azure.Core.Scalar namespace.
// Don't use this type directly, use a constructor function instead.
type ScalarClient struct {
	internal *azcore.Client
}

// NewScalarAzureLocationScalarClient creates a new instance of [ScalarAzureLocationScalarClient].
func (client *ScalarClient) NewScalarAzureLocationScalarClient() *ScalarAzureLocationScalarClient {
	return &ScalarAzureLocationScalarClient{
		internal: client.internal,
	}
}