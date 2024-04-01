// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package contentneggroup

import "io"

// ContentNegotiationDifferentBodyClientGetAvatarAsJSONResponse contains the response from method ContentNegotiationDifferentBodyClient.GetAvatarAsJSON.
type ContentNegotiationDifferentBodyClientGetAvatarAsJSONResponse struct {
	PNGImageAsJSON
}

// ContentNegotiationDifferentBodyClientGetAvatarAsPNGResponse contains the response from method ContentNegotiationDifferentBodyClient.GetAvatarAsPNG.
type ContentNegotiationDifferentBodyClientGetAvatarAsPNGResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// ContentNegotiationSameBodyClientGetAvatarAsJPEGResponse contains the response from method ContentNegotiationSameBodyClient.GetAvatarAsJPEG.
type ContentNegotiationSameBodyClientGetAvatarAsJPEGResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// ContentNegotiationSameBodyClientGetAvatarAsPNGResponse contains the response from method ContentNegotiationSameBodyClient.GetAvatarAsPNG.
type ContentNegotiationSameBodyClientGetAvatarAsPNGResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}
