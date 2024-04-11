// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package durationgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DurationClient - Test for encode decorator on duration.
// Don't use this type directly, use NewDurationClientWithNoCredential() instead.
type DurationClient struct {
	internal *azcore.Client
}

// DurationClientOptions contains the optional values for creating a [DurationClient].
type DurationClientOptions struct {
	azcore.ClientOptions
}

// NewDurationClientWithNoCredential creates a new instance of [DurationClient] with the specified values.
//   - options - DurationClientOptions contains the optional values for creating a [DurationClient]
func NewDurationClientWithNoCredential(options *DurationClientOptions) (*DurationClient, error) {
	if options == nil {
		options = &DurationClientOptions{}
	}
	cl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	durationClient := &DurationClient{
		internal: cl,
	}
	return durationClient, nil
}

// NewDurationHeaderClient creates a new instance of [DurationHeaderClient].
func (client *DurationClient) NewDurationHeaderClient() *DurationHeaderClient {
	return &DurationHeaderClient{
		internal: client.internal,
	}
}

// NewDurationPropertyClient creates a new instance of [DurationPropertyClient].
func (client *DurationClient) NewDurationPropertyClient() *DurationPropertyClient {
	return &DurationPropertyClient{
		internal: client.internal,
	}
}

// NewDurationQueryClient creates a new instance of [DurationQueryClient].
func (client *DurationClient) NewDurationQueryClient() *DurationQueryClient {
	return &DurationQueryClient{
		internal: client.internal,
	}
}
