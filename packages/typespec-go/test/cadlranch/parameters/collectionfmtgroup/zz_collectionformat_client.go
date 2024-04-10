// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package collectionfmtgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// CollectionFormatClient - Test for collectionFormat.
// Don't use this type directly, use a constructor function instead.
type CollectionFormatClient struct {
	internal *azcore.Client
}

// CollectionFormatClientOptions contains the optional values for creating a [CollectionFormatClient].
type CollectionFormatClientOptions struct {
	azcore.ClientOptions
}

// NewCollectionFormatClientWithNoCredential creates a new [CollectionFormatClient].
//   - options - optional client configuration; pass nil to accept the default values
func NewCollectionFormatClientWithNoCredential(options *CollectionFormatClientOptions) (*CollectionFormatClient, error) {
	if options == nil {
		options = &CollectionFormatClientOptions{}
	}
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	return &CollectionFormatClient{
		internal: internal,
	}, nil
}

// NewCollectionFormatHeaderClient creates a new instance of [HeaderClient].
func (client *CollectionFormatClient) NewCollectionFormatHeaderClient() *CollectionFormatHeaderClient {
	return &CollectionFormatHeaderClient{
		internal: client.internal,
	}
}

// NewCollectionFormatQueryClient creates a new instance of [CollectionFormatQueryClient].
func (client *CollectionFormatClient) NewCollectionFormatQueryClient() *CollectionFormatQueryClient {
	return &CollectionFormatQueryClient{
		internal: client.internal,
	}
}
