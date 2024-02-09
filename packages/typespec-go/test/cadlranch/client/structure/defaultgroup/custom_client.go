//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package defaultgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewBarClient(options *azcore.ClientOptions) (*BarClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &BarClient{
		internal: internal,
	}, nil
}

func NewFooClient(options *azcore.ClientOptions) (*FooClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &FooClient{
		internal: internal,
	}, nil
}

func NewQuxClient(options *azcore.ClientOptions) (*QuxClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &QuxClient{
		internal: internal,
	}, nil
}

func NewServiceClient(options *azcore.ClientOptions) (*ServiceClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ServiceClient{
		internal: internal,
	}, nil
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("defaultgroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
