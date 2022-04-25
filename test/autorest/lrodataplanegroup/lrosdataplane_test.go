// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrodataplanegroup

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newLROsDataplaneClient() *LROsDataplaneClient {
	return NewLROsDataplaneClient(runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, nil))
}

func TestBeginPostOplocSucceeded(t *testing.T) {
	client := newLROsDataplaneClient()
	const resultID = "test result"
	poller, err := client.BeginPostOplocSucceeded(context.Background(), resultID, nil)
	require.NoError(t, err)
	tk, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = client.BeginPostOplocSucceeded(context.Background(), resultID, &LROsDataplaneClientBeginPostOplocSucceededOptions{
		ResumeToken: tk,
	})
	require.NoError(t, err)
	result, err := poller.PollUntilDone(context.Background(), 1*time.Second)
	require.NoError(t, err)
	require.Equal(t, "processing complete", *result.Value)
}
