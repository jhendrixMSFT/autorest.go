// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmongocluster

import "time"

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string

	// The resource type.
	Type *string
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is not available.
	Message *string

	// Indicates if the resource name is available.
	NameAvailable *bool

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason
}

// ConnectionString - Connection string for the mongo cluster
type ConnectionString struct {
	// READ-ONLY; Value of the connection string
	ConnectionString *string

	// READ-ONLY; Description of the connection string
	Description *string
}

// FirewallRule - Represents a mongo cluster firewall rule.
type FirewallRule struct {
	// The resource-specific properties for this resource.
	Properties *FirewallRuleProperties

	// READ-ONLY; The name of the mongo cluster firewall rule.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FirewallRuleListResult - The response of a FirewallRule list operation.
type FirewallRuleListResult struct {
	// REQUIRED; The FirewallRule items on this page
	Value []*FirewallRule

	// The link to the next page of items
	NextLink *string
}

// FirewallRuleProperties - The properties of a mongo cluster firewall rule.
type FirewallRuleProperties struct {
	// REQUIRED; The end IP address of the mongo cluster firewall rule. Must be IPv4 format.
	EndIPAddress *string

	// REQUIRED; The start IP address of the mongo cluster firewall rule. Must be IPv4 format.
	StartIPAddress *string

	// READ-ONLY; The provisioning state of the firewall rule.
	ProvisioningState *ProvisioningState
}

// ListConnectionStringsResult - The connection strings for the given mongo cluster.
type ListConnectionStringsResult struct {
	// READ-ONLY; An array that contains the connection strings for a mongo cluster.
	ConnectionStrings []*ConnectionString
}

// ListResult - The response of a MongoCluster list operation.
type ListResult struct {
	// REQUIRED; The MongoCluster items on this page
	Value []*MongoCluster

	// The link to the next page of items
	NextLink *string
}

// MongoCluster - Represents a mongo cluster resource.
type MongoCluster struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *Properties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The name of the mongo cluster.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// NodeGroupSpec - Specification for a node group.
type NodeGroupSpec struct {
	// The disk storage size for the node group in GB. Example values: 128, 256, 512, 1024.
	DiskSizeGB *int64

	// Whether high availability is enabled on the node group.
	EnableHa *bool

	// The node type deployed in the node group.
	Kind *NodeKind

	// The number of nodes in the node group.
	NodeCount *int32

	// The resource sku for the node group. This defines the size of CPU and memory that is provisioned for each node. Example
	// values: 'M30', 'M40'.
	SKU *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Localized display information for this particular operation.
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
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
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

// PrivateEndpoint - The Private Endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The resource identifier for private endpoint
	ID *string
}

// PrivateEndpointConnection - The private endpoint connection resource
type PrivateEndpointConnection struct {
	// The private endpoint connection properties
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionProperties - Properties of the private endpoint connection.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// The private endpoint resource.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; The group ids for the private endpoint resource.
	GroupIDs []*string

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState
}

// PrivateEndpointConnectionResource - Concrete proxy resource types can be created by aliasing this type using a specific
// property type.
type PrivateEndpointConnectionResource struct {
	// The resource-specific properties for this resource.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; The name of the private endpoint connection associated with the Azure resource.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionResourceListResult - The response of a PrivateEndpointConnectionResource list operation.
type PrivateEndpointConnectionResourceListResult struct {
	// REQUIRED; The PrivateEndpointConnectionResource items on this page
	Value []*PrivateEndpointConnectionResource

	// The link to the next page of items
	NextLink *string
}

// PrivateLinkResource - Concrete proxy resource types can be created by aliasing this type using a specific property type.
type PrivateLinkResource struct {
	// The resource-specific properties for this resource.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; The name of the private link associated with the Azure resource.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceListResult - The response of a PrivateLinkResource list operation.
type PrivateLinkResourceListResult struct {
	// REQUIRED; The PrivateLinkResource items on this page
	Value []*PrivateLinkResource

	// The link to the next page of items
	NextLink *string
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource private link DNS zone name.
	RequiredZoneNames []*string

	// READ-ONLY; The private link resource group id.
	GroupID *string

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer
// and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string

	// The reason for approval/rejection of the connection.
	Description *string

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus
}

// Properties - The properties of a mongo cluster.
type Properties struct {
	// The administrator's login for the mongo cluster.
	AdministratorLogin *string

	// The password of the administrator login.
	AdministratorLoginPassword *string

	// The mode to create a mongo cluster.
	CreateMode *CreateMode

	// The list of node group specs in the cluster.
	NodeGroupSpecs []*NodeGroupSpec

	// Whether or not public endpoint access is allowed for this mongo cluster.
	PublicNetworkAccess *PublicNetworkAccess

	// The parameters to create a point-in-time restore mongo cluster.
	RestoreParameters *RestoreParameters

	// The Mongo DB server version. Defaults to the latest available version if not specified.
	ServerVersion *string

	// READ-ONLY; The status of the mongo cluster.
	ClusterStatus *Status

	// READ-ONLY; The default mongo connection string for the cluster.
	ConnectionString *string

	// READ-ONLY; Earliest restore timestamp in UTC ISO8601 format.
	EarliestRestoreTime *string

	// READ-ONLY; List of private endpoint connections.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; The provisioning state of the mongo cluster.
	ProvisioningState *ProvisioningState
}

// RestoreParameters - Parameters used for restore operations
type RestoreParameters struct {
	// UTC point in time to restore a mongo cluster
	PointInTimeUTC *time.Time

	// Resource ID to locate the source cluster to restore
	SourceResourceID *string
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
