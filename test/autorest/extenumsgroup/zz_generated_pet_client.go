//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package extenumsgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PetClient contains the methods for the Pet group.
// Don't use this type directly, use NewPetClient() instead.
type PetClient struct {
	pl runtime.Pipeline
}

// NewPetClient creates a new instance of PetClient with the specified values.
// options - pass nil to accept the default values.
func NewPetClient(options *azcore.ClientOptions) *PetClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &PetClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// AddPet - add pet
// If the operation fails it returns an *azcore.ResponseError type.
// options - PetClientAddPetOptions contains the optional parameters for the PetClient.AddPet method.
func (client *PetClient) AddPet(ctx context.Context, options *PetClientAddPetOptions) (PetClientAddPetResponse, error) {
	req, err := client.addPetCreateRequest(ctx, options)
	if err != nil {
		return PetClientAddPetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PetClientAddPetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PetClientAddPetResponse{}, runtime.NewResponseError(resp)
	}
	return client.addPetHandleResponse(resp)
}

// addPetCreateRequest creates the AddPet request.
func (client *PetClient) addPetCreateRequest(ctx context.Context, options *PetClientAddPetOptions) (*policy.Request, error) {
	urlPath := "/extensibleenums/pet/addPet"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.PetParam != nil {
		return req, runtime.MarshalAsJSON(req, *options.PetParam)
	}
	return req, nil
}

// addPetHandleResponse handles the AddPet response.
func (client *PetClient) addPetHandleResponse(resp *http.Response) (PetClientAddPetResponse, error) {
	result := PetClientAddPetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pet); err != nil {
		return PetClientAddPetResponse{}, err
	}
	return result, nil
}

// GetByPetID - get pet by id
// If the operation fails it returns an *azcore.ResponseError type.
// petID - Pet id
// options - PetClientGetByPetIDOptions contains the optional parameters for the PetClient.GetByPetID method.
func (client *PetClient) GetByPetID(ctx context.Context, petID string, options *PetClientGetByPetIDOptions) (PetClientGetByPetIDResponse, error) {
	req, err := client.getByPetIDCreateRequest(ctx, petID, options)
	if err != nil {
		return PetClientGetByPetIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PetClientGetByPetIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PetClientGetByPetIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByPetIDHandleResponse(resp)
}

// getByPetIDCreateRequest creates the GetByPetID request.
func (client *PetClient) getByPetIDCreateRequest(ctx context.Context, petID string, options *PetClientGetByPetIDOptions) (*policy.Request, error) {
	urlPath := "/extensibleenums/pet/{petId}"
	if petID == "" {
		return nil, errors.New("parameter petID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{petId}", url.PathEscape(petID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByPetIDHandleResponse handles the GetByPetID response.
func (client *PetClient) getByPetIDHandleResponse(resp *http.Response) (PetClientGetByPetIDResponse, error) {
	result := PetClientGetByPetIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pet); err != nil {
		return PetClientGetByPetIDResponse{}, err
	}
	return result, nil
}
