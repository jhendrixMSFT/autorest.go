// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package nonstringenumgroup

import "net/http"

// FloatEnumResponse is the response envelope for operations that return a FloatEnum type.
type FloatEnumResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// List of float enums
	Value *FloatEnum
}

// FloatGetOptions contains the optional parameters for the Float.Get method.
type FloatGetOptions struct {
	// placeholder for future optional parameters
}

// FloatPutOptions contains the optional parameters for the Float.Put method.
type FloatPutOptions struct {
	// Input float enum.
	Input *FloatEnum
}

// IntEnumResponse is the response envelope for operations that return a IntEnum type.
type IntEnumResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// List of integer enums
	Value *IntEnum
}

// IntGetOptions contains the optional parameters for the Int.Get method.
type IntGetOptions struct {
	// placeholder for future optional parameters
}

// IntPutOptions contains the optional parameters for the Int.Put method.
type IntPutOptions struct {
	// Input int enum.
	Input *IntEnum
}

// StringResponse is the response envelope for operations that return a string type.
type StringResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Value       *string
}
