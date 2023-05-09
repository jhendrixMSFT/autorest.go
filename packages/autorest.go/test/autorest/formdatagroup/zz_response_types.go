//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package formdatagroup

import "io"

// FormdataClientUploadFileResponse contains the response from method FormdataClient.UploadFile.
type FormdataClientUploadFileResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// FormdataClientUploadFileViaBodyResponse contains the response from method FormdataClient.UploadFileViaBody.
type FormdataClientUploadFileViaBodyResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// FormdataClientUploadFilesResponse contains the response from method FormdataClient.UploadFiles.
type FormdataClientUploadFilesResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}