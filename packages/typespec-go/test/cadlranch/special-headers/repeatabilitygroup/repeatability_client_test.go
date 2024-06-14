// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package repeatabilitygroup_test

import (
	"context"
	"repeatabilitygroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeatabilityClient_ImmediateSuccess(t *testing.T) {
	client, err := repeatabilitygroup.NewRepeatabilityClient(nil)
	require.NoError(t, err)
	resp, err := client.ImmediateSuccess(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.RepeatabilityResult)
	require.Equal(t, *resp.RepeatabilityResult, repeatabilitygroup.ImmediateSuccessResponseRepeatabilityResultAccepted)
}
