// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package extensiblegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ExtensibleClient contains the methods for the Type.Enum.Extensible namespace.
// Don't use this type directly, use NewExtensibleClientWithNoCredential() instead.
type ExtensibleClient struct {
	internal *azcore.Client
}

// ExtensibleClientOptions contains the optional values for creating a [ExtensibleClient].
type ExtensibleClientOptions struct {
	azcore.ClientOptions
}

// NewExtensibleClientWithNoCredential creates a new instance of [ExtensibleClient] with the specified values.
//   - options - ExtensibleClientOptions contains the optional values for creating a [ExtensibleClient]
func NewExtensibleClientWithNoCredential(options *ExtensibleClientOptions) (*ExtensibleClient, error) {
	if options == nil {
		options = &ExtensibleClientOptions{}
	}
	cl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	extensibleClient := &ExtensibleClient{
		internal: cl,
	}
	return extensibleClient, nil
}

// NewExtensibleStringClient creates a new instance of [ExtensibleStringClient].
func (client *ExtensibleClient) NewExtensibleStringClient() *ExtensibleStringClient {
	return &ExtensibleStringClient{
		internal: client.internal,
	}
}
