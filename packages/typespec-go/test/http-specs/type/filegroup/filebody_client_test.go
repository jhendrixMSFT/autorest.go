// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package filegroup_test

import (
	"context"
	"filegroup"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/require"
)

func TestFileBodyClient_DownloadFileSpecificContentType(t *testing.T) {
	client, err := filegroup.NewFileClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewFileBodyClient().DownloadFileSpecificContentType(context.Background(), nil)
	require.NoError(t, err)
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	pngFile, err := os.ReadFile("../../../../node_modules/@typespec/http-specs/assets/image.png")
	require.NoError(t, err)
	require.EqualValues(t, pngFile, respBody)
	require.NotNil(t, resp.ContentType)
	require.Equal(t, "image/png", *resp.ContentType)
}

func TestFileBodyClient_DownloadFileJSONContentType(t *testing.T) {
	client, err := filegroup.NewFileClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewFileBodyClient().DownloadFileJSONContentType(context.Background(), nil)
	require.NoError(t, err)
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.JSONEq(t, `{"message":"test file content"}`, string(respBody))
	require.NotNil(t, resp.ContentType)
	require.Equal(t, "application/json; charset=utf-8", *resp.ContentType)
}

func TestFileBodyClient_DownloadFileMultipleContentTypes(t *testing.T) {
	client, err := filegroup.NewFileClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewFileBodyClient().DownloadFileMultipleContentTypes(context.Background(), nil)
	require.NoError(t, err)
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	pngFile, err := os.ReadFile("../../../../node_modules/@typespec/http-specs/assets/image.png")
	require.NoError(t, err)
	require.EqualValues(t, pngFile, respBody)
	require.NotNil(t, resp.ContentType)
	require.Equal(t, filegroup.DownloadFileMultipleContentTypesResponseContentTypeImagePNG, *resp.ContentType)
}

func TestFileBodyClient_DownloadFileDefaultContentType(t *testing.T) {
	client, err := filegroup.NewFileClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewFileBodyClient().DownloadFileDefaultContentType(context.Background(), nil)
	require.NoError(t, err)
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	pngFile, err := os.ReadFile("../../../../node_modules/@typespec/http-specs/assets/image.png")
	require.NoError(t, err)
	require.EqualValues(t, pngFile, respBody)
	require.NotNil(t, resp.ContentType)
	require.Equal(t, "image/png", *resp.ContentType)
}

func TestFileBodyClient_UploadFileSpecificContentType(t *testing.T) {
	client, err := filegroup.NewFileClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	file, err := os.Open("../../../../node_modules/@typespec/http-specs/assets/image.png")
	require.NoError(t, err)
	defer file.Close()
	_, err = client.NewFileBodyClient().UploadFileSpecificContentType(context.Background(), file, nil)
	require.NoError(t, err)
}

func TestFileBodyClient_UploadFileJSONContentType(t *testing.T) {
	client, err := filegroup.NewFileClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	body := `{"message":"test file content"}`
	file := streaming.NopCloser(strings.NewReader(body))
	_, err = client.NewFileBodyClient().UploadFileJSONContentType(context.Background(), file, nil)
	require.NoError(t, err)
}

func TestFileBodyClient_UploadFileMultipleContentTypes(t *testing.T) {
	client, err := filegroup.NewFileClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	file, err := os.Open("../../../../node_modules/@typespec/http-specs/assets/image.png")
	require.NoError(t, err)
	defer file.Close()
	_, err = client.NewFileBodyClient().UploadFileMultipleContentTypes(context.Background(), file, filegroup.FileContentType2ImagePNG, nil)
	require.NoError(t, err)
}

func TestFileBodyClient_UploadFileDefaultContentType(t *testing.T) {
	client, err := filegroup.NewFileClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	file, err := os.Open("../../../../node_modules/@typespec/http-specs/assets/image.png")
	require.NoError(t, err)
	defer file.Close()
	ct := "image/png"
	_, err = client.NewFileBodyClient().UploadFileDefaultContentType(context.Background(), file, &filegroup.FileBodyClientUploadFileDefaultContentTypeOptions{
		ContentType: &ct,
	})
	require.NoError(t, err)
}
