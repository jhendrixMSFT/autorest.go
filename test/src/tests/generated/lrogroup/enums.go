package lrogroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ProvisioningStateValues enumerates the values for provisioning state values.
type ProvisioningStateValues string

const (
	// Accepted ...
	Accepted ProvisioningStateValues = "Accepted"
	// Canceled ...
	Canceled ProvisioningStateValues = "canceled"
	// Created ...
	Created ProvisioningStateValues = "Created"
	// Creating ...
	Creating ProvisioningStateValues = "Creating"
	// Deleted ...
	Deleted ProvisioningStateValues = "Deleted"
	// Deleting ...
	Deleting ProvisioningStateValues = "Deleting"
	// Failed ...
	Failed ProvisioningStateValues = "Failed"
	// OK ...
	OK ProvisioningStateValues = "OK"
	// Succeeded ...
	Succeeded ProvisioningStateValues = "Succeeded"
	// Updated ...
	Updated ProvisioningStateValues = "Updated"
	// Updating ...
	Updating ProvisioningStateValues = "Updating"
)

// PossibleProvisioningStateValuesValues returns an array of possible values for the ProvisioningStateValues const type.
func PossibleProvisioningStateValuesValues() []ProvisioningStateValues {
	return []ProvisioningStateValues{Accepted, Canceled, Created, Creating, Deleted, Deleting, Failed, OK, Succeeded, Updated, Updating}
}

// ProvisioningStateValues1 enumerates the values for provisioning state values 1.
type ProvisioningStateValues1 string

const (
	// ProvisioningStateValues1Accepted ...
	ProvisioningStateValues1Accepted ProvisioningStateValues1 = "Accepted"
	// ProvisioningStateValues1Canceled ...
	ProvisioningStateValues1Canceled ProvisioningStateValues1 = "canceled"
	// ProvisioningStateValues1Created ...
	ProvisioningStateValues1Created ProvisioningStateValues1 = "Created"
	// ProvisioningStateValues1Creating ...
	ProvisioningStateValues1Creating ProvisioningStateValues1 = "Creating"
	// ProvisioningStateValues1Deleted ...
	ProvisioningStateValues1Deleted ProvisioningStateValues1 = "Deleted"
	// ProvisioningStateValues1Deleting ...
	ProvisioningStateValues1Deleting ProvisioningStateValues1 = "Deleting"
	// ProvisioningStateValues1Failed ...
	ProvisioningStateValues1Failed ProvisioningStateValues1 = "Failed"
	// ProvisioningStateValues1OK ...
	ProvisioningStateValues1OK ProvisioningStateValues1 = "OK"
	// ProvisioningStateValues1Succeeded ...
	ProvisioningStateValues1Succeeded ProvisioningStateValues1 = "Succeeded"
	// ProvisioningStateValues1Updated ...
	ProvisioningStateValues1Updated ProvisioningStateValues1 = "Updated"
	// ProvisioningStateValues1Updating ...
	ProvisioningStateValues1Updating ProvisioningStateValues1 = "Updating"
)

// PossibleProvisioningStateValues1Values returns an array of possible values for the ProvisioningStateValues1 const type.
func PossibleProvisioningStateValues1Values() []ProvisioningStateValues1 {
	return []ProvisioningStateValues1{ProvisioningStateValues1Accepted, ProvisioningStateValues1Canceled, ProvisioningStateValues1Created, ProvisioningStateValues1Creating, ProvisioningStateValues1Deleted, ProvisioningStateValues1Deleting, ProvisioningStateValues1Failed, ProvisioningStateValues1OK, ProvisioningStateValues1Succeeded, ProvisioningStateValues1Updated, ProvisioningStateValues1Updating}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusAccepted ...
	StatusAccepted Status = "Accepted"
	// StatusCanceled ...
	StatusCanceled Status = "canceled"
	// StatusCreated ...
	StatusCreated Status = "Created"
	// StatusCreating ...
	StatusCreating Status = "Creating"
	// StatusDeleted ...
	StatusDeleted Status = "Deleted"
	// StatusDeleting ...
	StatusDeleting Status = "Deleting"
	// StatusFailed ...
	StatusFailed Status = "Failed"
	// StatusOK ...
	StatusOK Status = "OK"
	// StatusSucceeded ...
	StatusSucceeded Status = "Succeeded"
	// StatusUpdated ...
	StatusUpdated Status = "Updated"
	// StatusUpdating ...
	StatusUpdating Status = "Updating"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusAccepted, StatusCanceled, StatusCreated, StatusCreating, StatusDeleted, StatusDeleting, StatusFailed, StatusOK, StatusSucceeded, StatusUpdated, StatusUpdating}
}
