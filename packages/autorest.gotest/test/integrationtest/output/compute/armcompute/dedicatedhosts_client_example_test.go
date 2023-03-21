//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/CreateOrUpdateADedicatedHost.json
func ExampleDedicatedHostsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDedicatedHostGroup", "myDedicatedHost", armcompute.DedicatedHost{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("HR"),
		},
		Properties: &armcompute.DedicatedHostProperties{
			PlatformFaultDomain: to.Ptr[int32](1),
		},
		SKU: &armcompute.SKU{
			Name: to.Ptr("DSv3-Type1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHost = armcompute.DedicatedHost{
	// 	Name: to.Ptr("myDedicatedHost"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myDedicatedHost"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("HR"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostProperties{
	// 		AutoReplaceOnFailure: to.Ptr(false),
	// 		HostID: to.Ptr("{GUID}"),
	// 		LicenseType: to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
	// 		PlatformFaultDomain: to.Ptr[int32](1),
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("DSv3-Type1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetADedicatedHost.json
func ExampleDedicatedHostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostsClient().Get(ctx, "myResourceGroup", "myDedicatedHostGroup", "myHost", &armcompute.DedicatedHostsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHost = armcompute.DedicatedHost{
	// 	Name: to.Ptr("myHost"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("HR"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostProperties{
	// 		AutoReplaceOnFailure: to.Ptr(true),
	// 		HostID: to.Ptr("{GUID}"),
	// 		InstanceView: &armcompute.DedicatedHostInstanceView{
	// 			AssetID: to.Ptr("eb3f58b8-b4e8-4882-b69f-301a01812407"),
	// 			AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 				AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 					{
	// 						Count: to.Ptr[float64](10),
	// 						VMSize: to.Ptr("Standard_A1"),
	// 				}},
	// 			},
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				},
	// 				{
	// 					Code: to.Ptr("HealthState/available"),
	// 					DisplayStatus: to.Ptr("Host available"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			}},
	// 		},
	// 		PlatformFaultDomain: to.Ptr[int32](1),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-27T01:02:38.3138469+00:00"); return t}()),
	// 		VirtualMachines: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("DSv3-Type1"),
	// 	},
	// }
}
