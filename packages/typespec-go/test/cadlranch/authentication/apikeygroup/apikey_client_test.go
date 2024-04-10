// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apikeygroup_test

import (
	"apikeygroup"
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestApiKeyClient_Invalid(t *testing.T) {
	client, err := apikeygroup.NewApiKeyClientWithKeyCredential(azcore.NewKeyCredential("invalid-key"), &apikeygroup.ApiKeyClientOptions{
		ClientOptions: azcore.ClientOptions{
			InsecureAllowCredentialWithHTTP: true,
		},
	})
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), nil)
	var respErr *azcore.ResponseError
	require.ErrorAs(t, err, &respErr)
	require.EqualValues(t, http.StatusForbidden, respErr.StatusCode)
	require.Zero(t, resp)
}

func TestApiKeyClient_Valid(t *testing.T) {}
