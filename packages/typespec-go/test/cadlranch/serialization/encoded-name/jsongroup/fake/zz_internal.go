// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"net/http"
	"sync"
)

type result struct {
	resp *http.Response
	err  error
}

type nonRetriableError struct {
	error
}

func (nonRetriableError) NonRetriable() {
	// marker method
}

func contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func initServer[T any](mu *sync.Mutex, dst **T, src func() *T) {
	mu.Lock()
	if *dst == nil {
		*dst = src()
	}
	mu.Unlock()
}
