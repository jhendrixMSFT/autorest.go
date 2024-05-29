// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package uniongroup

import "encoding/json"

// UnionEnumsOnlyClientGetResponse contains the response from method UnionEnumsOnlyClient.Get.
type UnionEnumsOnlyClientGetResponse struct {
	EnumsOnlyCases
}

func (u *UnionEnumsOnlyClientGetResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return json.Unmarshal(rawMsg["prop"], &u.EnumsOnlyCases)
}

// UnionEnumsOnlyClientSendResponse contains the response from method UnionEnumsOnlyClient.Send.
type UnionEnumsOnlyClientSendResponse struct {
	// placeholder for future response values
}

// UnionFloatsOnlyClientGetResponse contains the response from method UnionFloatsOnlyClient.Get.
type UnionFloatsOnlyClientGetResponse struct {
	Value *UnionFloatsOnlyClientGetCases
}

func (u *UnionFloatsOnlyClientGetResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return json.Unmarshal(rawMsg["prop"], &u.Value)
}

// UnionFloatsOnlyClientSendResponse contains the response from method UnionFloatsOnlyClient.Send.
type UnionFloatsOnlyClientSendResponse struct {
	// placeholder for future response values
}

// UnionIntsOnlyClientGetResponse contains the response from method UnionIntsOnlyClient.Get.
type UnionIntsOnlyClientGetResponse struct {
	Value *UnionIntsOnlyClientGetCases
}

func (u *UnionIntsOnlyClientGetResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return json.Unmarshal(rawMsg["prop"], &u.Value)
}

// UnionIntsOnlyClientSendResponse contains the response from method UnionIntsOnlyClient.Send.
type UnionIntsOnlyClientSendResponse struct {
	// placeholder for future response values
}

// UnionMixedLiteralsClientGetResponse contains the response from method UnionMixedLiteralsClient.Get.
type UnionMixedLiteralsClientGetResponse struct {
	MixedLiteralsCases
}

func (u *UnionMixedLiteralsClientGetResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return json.Unmarshal(rawMsg["prop"], &u.MixedLiteralsCases)
}

// UnionMixedLiteralsClientSendResponse contains the response from method UnionMixedLiteralsClient.Send.
type UnionMixedLiteralsClientSendResponse struct {
	// placeholder for future response values
}

// UnionMixedTypesClientGetResponse contains the response from method UnionMixedTypesClient.Get.
type UnionMixedTypesClientGetResponse struct {
	MixedTypesCases
}

func (u *UnionMixedTypesClientGetResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return json.Unmarshal(rawMsg["prop"], &u.MixedTypesCases)
}

// UnionMixedTypesClientSendResponse contains the response from method UnionMixedTypesClient.Send.
type UnionMixedTypesClientSendResponse struct {
	// placeholder for future response values
}

// UnionModelsOnlyClientGetResponse contains the response from method UnionModelsOnlyClient.Get.
type UnionModelsOnlyClientGetResponse struct {
	GetResponse4
}

// UnionModelsOnlyClientSendResponse contains the response from method UnionModelsOnlyClient.Send.
type UnionModelsOnlyClientSendResponse struct {
	// placeholder for future response values
}

// UnionStringAndArrayClientGetResponse contains the response from method UnionStringAndArrayClient.Get.
type UnionStringAndArrayClientGetResponse struct {
	GetResponse2
}

// UnionStringAndArrayClientSendResponse contains the response from method UnionStringAndArrayClient.Send.
type UnionStringAndArrayClientSendResponse struct {
	// placeholder for future response values
}

// UnionStringExtensibleClientGetResponse contains the response from method UnionStringExtensibleClient.Get.
type UnionStringExtensibleClientGetResponse struct {
	GetResponse8
}

// UnionStringExtensibleClientSendResponse contains the response from method UnionStringExtensibleClient.Send.
type UnionStringExtensibleClientSendResponse struct {
	// placeholder for future response values
}

// UnionStringExtensibleNamedClientGetResponse contains the response from method UnionStringExtensibleNamedClient.Get.
type UnionStringExtensibleNamedClientGetResponse struct {
	GetResponse7
}

// UnionStringExtensibleNamedClientSendResponse contains the response from method UnionStringExtensibleNamedClient.Send.
type UnionStringExtensibleNamedClientSendResponse struct {
	// placeholder for future response values
}

// UnionStringsOnlyClientGetResponse contains the response from method UnionStringsOnlyClient.Get.
type UnionStringsOnlyClientGetResponse struct {
	GetResponse9
}

// UnionStringsOnlyClientSendResponse contains the response from method UnionStringsOnlyClient.Send.
type UnionStringsOnlyClientSendResponse struct {
	// placeholder for future response values
}