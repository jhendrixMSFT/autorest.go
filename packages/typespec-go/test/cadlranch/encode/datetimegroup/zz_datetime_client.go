// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package datetimegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DatetimeClient - Test for encode decorator on datetime.
// Don't use this type directly, use a constructor function instead.
type DatetimeClient struct {
	internal *azcore.Client
}

// DatetimeClientOptions contains the optional values for creating a [DatetimeClient].
type DatetimeClientOptions struct {
	azcore.ClientOptions
}

// NewDatetimeClientWithNoCredential creates a new [DatetimeClient].
//   - options - optional client configuration; pass nil to accept the default values
func NewDatetimeClientWithNoCredential(options *DatetimeClientOptions) (*DatetimeClient, error) {
	if options == nil {
		options = &DatetimeClientOptions{}
	}
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	return &DatetimeClient{
		internal: internal,
	}, nil
}

// NewDatetimeHeaderClient creates a new instance of [HeaderClient].
func (client *DatetimeClient) NewDatetimeHeaderClient() *DatetimeHeaderClient {
	return &DatetimeHeaderClient{
		internal: client.internal,
	}
}

// NewDatetimePropertyClient creates a new instance of [DatetimePropertyClient].
func (client *DatetimeClient) NewDatetimePropertyClient() *DatetimePropertyClient {
	return &DatetimePropertyClient{
		internal: client.internal,
	}
}

// NewDatetimeQueryClient creates a new instance of [DatetimeQueryClient].
func (client *DatetimeClient) NewDatetimeQueryClient() *DatetimeQueryClient {
	return &DatetimeQueryClient{
		internal: client.internal,
	}
}

// NewDatetimeResponseHeaderClient creates a new instance of [DatetimeResponseHeaderClient].
func (client *DatetimeClient) NewDatetimeResponseHeaderClient() *DatetimeResponseHeaderClient {
	return &DatetimeResponseHeaderClient{
		internal: client.internal,
	}
}
