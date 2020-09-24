// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

import (
	"fmt"
	"net/http"
)

// ByteArrayResponse is the response envelope for operations that return a []byte type.
type ByteArrayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The null byte value
	Value *[]byte
}

// ByteGetEmptyOptions contains the optional parameters for the Byte.GetEmpty method.
type ByteGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// ByteGetInvalidOptions contains the optional parameters for the Byte.GetInvalid method.
type ByteGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// ByteGetNonASCIIOptions contains the optional parameters for the Byte.GetNonASCII method.
type ByteGetNonASCIIOptions struct {
	// placeholder for future optional parameters
}

// ByteGetNullOptions contains the optional parameters for the Byte.GetNull method.
type ByteGetNullOptions struct {
	// placeholder for future optional parameters
}

// BytePutNonASCIIOptions contains the optional parameters for the Byte.PutNonASCII method.
type BytePutNonASCIIOptions struct {
	// placeholder for future optional parameters
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}
