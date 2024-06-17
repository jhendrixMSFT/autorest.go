// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcommunitymanagement

import "time"

// CommunityTraining - A CommunityProviderHub resource
type CommunityTraining struct {
	// The resource-specific properties for this resource.
	Properties *CommunityTrainingProperties

	// The SKU (Stock Keeping Unit) assigned to this resource.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the Community Training Resource
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CommunityTrainingListResult - The response of a CommunityTraining list operation.
type CommunityTrainingListResult struct {
	// REQUIRED; The CommunityTraining items on this page
	Value []*CommunityTraining

	// The link to the next page of items
	NextLink *string
}

// CommunityTrainingProperties - Details of the Community CommunityTraining.
type CommunityTrainingProperties struct {
	// READ-ONLY; To indicate whether the Community Training instance has Disaster Recovery enabled
	DisasterRecoveryEnabled *bool

	// READ-ONLY; The identity configuration of the Community Training resource
	IdentityConfiguration *IdentityConfigurationProperties

	// READ-ONLY; The email address of the portal admin
	PortalAdminEmailAddress *string

	// READ-ONLY; The portal name (website name) of the Community Training instance
	PortalName *string

	// READ-ONLY; The email address of the portal owner. Will be used as the primary contact
	PortalOwnerEmailAddress *string

	// READ-ONLY; The organization name of the portal owner
	PortalOwnerOrganizationName *string

	// READ-ONLY; To indicate whether the Community Training instance has Zone Redundancy enabled
	ZoneRedundancyEnabled *bool

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// CommunityTrainingUpdate - The type used for update operations of the CommunityTraining.
type CommunityTrainingUpdate struct {
	Properties *CommunityTrainingUpdateProperties

	// The SKU (Stock Keeping Unit) assigned to this resource.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string
}

// CommunityTrainingUpdateProperties - The updatable properties of the CommunityTraining.
type CommunityTrainingUpdateProperties struct {
	// READ-ONLY; The identity configuration of the Community Training resource
	IdentityConfiguration *IdentityConfigurationProperties
}

// IdentityConfigurationProperties - Details of the Community CommunityTraining Identity Configuration
type IdentityConfigurationProperties struct {
	// READ-ONLY; The clientId of the application registered in the selected identity provider for the Community Training Resource
	ClientID *string

	// READ-ONLY; The client secret of the application registered in the selected identity provider for the Community Training
	// Resource
	ClientSecret *string

	// READ-ONLY; The domain name of the selected identity provider for the Community Training Resource
	DomainName *string

	// READ-ONLY; The identity type of the Community Training Resource
	IdentityType *string

	// READ-ONLY; The tenantId of the selected identity provider for the Community Training Resource
	TenantID *string

	// READ-ONLY; The name of the authentication policy registered in ADB2C for the Community Training Resource
	B2CAuthenticationPolicy *string

	// READ-ONLY; The name of the password reset policy registered in ADB2C for the Community Training Resource
	B2CPasswordResetPolicy *string

	// READ-ONLY; The custom login parameters for the Community Training Resource
	CustomLoginParameters *string

	// READ-ONLY; To indicate whether the Community Training Resource has Teams enabled
	TeamsEnabled *bool
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure
	// Resource Manager/control-plane operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for and operation.
type OperationDisplay struct {
	// The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual Machine",
	// "Restart Virtual Machine".
	Operation *string

	// The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft Compute".
	Provider *string

	// The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// REQUIRED; The Operation items on this page
	Value []*Operation

	// The link to the next page of items
	NextLink *string
}

// SKU - The resource model definition representing SKU
type SKU struct {
	// REQUIRED; The name of the SKU. Ex - P3. It is typically a letter+number code
	Name *string

	// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the
	// resource this may be omitted.
	Capacity *int32

	// If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string

	// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string

	// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required
	// on a PUT.
	Tier *SKUTier
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}
