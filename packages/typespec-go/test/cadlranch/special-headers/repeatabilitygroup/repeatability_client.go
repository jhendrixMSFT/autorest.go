// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package repeatabilitygroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewRepeatabilityClient(options *azcore.ClientOptions) (*RepeatabilityClient, error) {
	internal, err := azcore.NewClient("repeatabilitygroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &RepeatabilityClient{
		internal: internal,
	}, nil
}
