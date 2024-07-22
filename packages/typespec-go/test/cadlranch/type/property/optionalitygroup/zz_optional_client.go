// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package optionalitygroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// OptionalClient - Illustrates models with optional properties.
// Don't use this type directly, use a constructor function instead.
type OptionalClient struct {
	internal *azcore.Client
}

// NewOptionalBooleanLiteralClient creates a new instance of [OptionalBooleanLiteralClient].
func (client *OptionalClient) NewOptionalBooleanLiteralClient() *OptionalBooleanLiteralClient {
	return &OptionalBooleanLiteralClient{
		internal: client.internal,
	}
}

// NewOptionalBytesClient creates a new instance of [OptionalBytesClient].
func (client *OptionalClient) NewOptionalBytesClient() *OptionalBytesClient {
	return &OptionalBytesClient{
		internal: client.internal,
	}
}

// NewOptionalCollectionsByteClient creates a new instance of [OptionalCollectionsByteClient].
func (client *OptionalClient) NewOptionalCollectionsByteClient() *OptionalCollectionsByteClient {
	return &OptionalCollectionsByteClient{
		internal: client.internal,
	}
}

// NewOptionalCollectionsModelClient creates a new instance of [OptionalCollectionsModelClient].
func (client *OptionalClient) NewOptionalCollectionsModelClient() *OptionalCollectionsModelClient {
	return &OptionalCollectionsModelClient{
		internal: client.internal,
	}
}

// NewOptionalDatetimeClient creates a new instance of [OptionalDatetimeClient].
func (client *OptionalClient) NewOptionalDatetimeClient() *OptionalDatetimeClient {
	return &OptionalDatetimeClient{
		internal: client.internal,
	}
}

// NewOptionalDurationClient creates a new instance of [OptionalDurationClient].
func (client *OptionalClient) NewOptionalDurationClient() *OptionalDurationClient {
	return &OptionalDurationClient{
		internal: client.internal,
	}
}

// NewOptionalFloatLiteralClient creates a new instance of [OptionalFloatLiteralClient].
func (client *OptionalClient) NewOptionalFloatLiteralClient() *OptionalFloatLiteralClient {
	return &OptionalFloatLiteralClient{
		internal: client.internal,
	}
}

// NewOptionalIntLiteralClient creates a new instance of [OptionalIntLiteralClient].
func (client *OptionalClient) NewOptionalIntLiteralClient() *OptionalIntLiteralClient {
	return &OptionalIntLiteralClient{
		internal: client.internal,
	}
}

// NewOptionalPlainDateClient creates a new instance of [OptionalPlainDateClient].
func (client *OptionalClient) NewOptionalPlainDateClient() *OptionalPlainDateClient {
	return &OptionalPlainDateClient{
		internal: client.internal,
	}
}

// NewOptionalPlainTimeClient creates a new instance of [OptionalPlainTimeClient].
func (client *OptionalClient) NewOptionalPlainTimeClient() *OptionalPlainTimeClient {
	return &OptionalPlainTimeClient{
		internal: client.internal,
	}
}

// NewOptionalRequiredAndOptionalClient creates a new instance of [OptionalRequiredAndOptionalClient].
func (client *OptionalClient) NewOptionalRequiredAndOptionalClient() *OptionalRequiredAndOptionalClient {
	return &OptionalRequiredAndOptionalClient{
		internal: client.internal,
	}
}

// NewOptionalStringClient creates a new instance of [OptionalStringClient].
func (client *OptionalClient) NewOptionalStringClient() *OptionalStringClient {
	return &OptionalStringClient{
		internal: client.internal,
	}
}

// NewOptionalStringLiteralClient creates a new instance of [OptionalStringLiteralClient].
func (client *OptionalClient) NewOptionalStringLiteralClient() *OptionalStringLiteralClient {
	return &OptionalStringLiteralClient{
		internal: client.internal,
	}
}

// NewOptionalUnionFloatLiteralClient creates a new instance of [OptionalUnionFloatLiteralClient].
func (client *OptionalClient) NewOptionalUnionFloatLiteralClient() *OptionalUnionFloatLiteralClient {
	return &OptionalUnionFloatLiteralClient{
		internal: client.internal,
	}
}

// NewOptionalUnionIntLiteralClient creates a new instance of [OptionalUnionIntLiteralClient].
func (client *OptionalClient) NewOptionalUnionIntLiteralClient() *OptionalUnionIntLiteralClient {
	return &OptionalUnionIntLiteralClient{
		internal: client.internal,
	}
}

// NewOptionalUnionStringLiteralClient creates a new instance of [OptionalUnionStringLiteralClient].
func (client *OptionalClient) NewOptionalUnionStringLiteralClient() *OptionalUnionStringLiteralClient {
	return &OptionalUnionStringLiteralClient{
		internal: client.internal,
	}
}
