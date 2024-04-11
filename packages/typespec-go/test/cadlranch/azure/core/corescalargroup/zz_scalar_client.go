// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package corescalargroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ScalarClient contains the methods for the _Specs_.Azure.Core.Scalar namespace.
// Don't use this type directly, use NewScalarClientWithNoCredential() instead.
type ScalarClient struct {
	internal *azcore.Client
}

// ScalarClientOptions contains the optional values for creating a [ScalarClient].
type ScalarClientOptions struct {
	azcore.ClientOptions
}

// NewScalarClientWithNoCredential creates a new instance of [ScalarClient] with the specified values.
//   - options - ScalarClientOptions contains the optional values for creating a [ScalarClient]
func NewScalarClientWithNoCredential(options *ScalarClientOptions) (*ScalarClient, error) {
	if options == nil {
		options = &ScalarClientOptions{}
	}
	cl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	scalarClient := &ScalarClient{
		internal: cl,
	}
	return scalarClient, nil
}

// NewScalarAzureLocationScalarClient creates a new instance of [ScalarAzureLocationScalarClient].
func (client *ScalarClient) NewScalarAzureLocationScalarClient() *ScalarAzureLocationScalarClient {
	return &ScalarAzureLocationScalarClient{
		internal: client.internal,
	}
}
