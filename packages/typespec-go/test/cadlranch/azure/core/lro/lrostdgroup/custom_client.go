//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrostdgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewStandardClient(options *azcore.ClientOptions) (*StandardClient, error) {
	internal, err := azcore.NewClient("lrostdgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &StandardClient{
		internal:   internal,
		apiVersion: "2022-12-01-preview",
	}, nil
}
