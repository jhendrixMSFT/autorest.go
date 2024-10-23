// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdevopsinfrastructure_test

import (
	"armdevopsinfrastructure"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2024-04-04-preview/ResourceDetails_ListByPool.json
func ExampleResourceDetailsClient_NewListByPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevopsinfrastructure.NewClientFactory("a2e95d27-c161-4b61-bda4-11512c14c2c2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceDetailsClient().NewListByPoolPager("my-resource-group", "my-dev-ops-pool", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armdevopsinfrastructure.ResourceDetailsClientListByPoolResponse{
		// 	ResourceDetailsObjectListResult: armdevopsinfrastructure.ResourceDetailsObjectListResult{
		// 		Value: []*armdevopsinfrastructure.ResourceDetailsObject{
		// 			{
		// 				ID: to.Ptr("/subscriptions/a2e95d27-c161-4b61-bda4-11512c14c2c2/resourceGroups/my-resource-group/providers/Microsoft.DevOpsInfrastructure/pools/my-devops-pool/resources/my-dev-ops-pool:04dcde21-626e-5a7e-8659-ce12f9284b29:dd8cc705c_0"),
		// 				Name: to.Ptr("dd8cc705c000000"),
		// 				Properties: &armdevopsinfrastructure.ResourceDetailsObjectProperties{
		// 					Image: to.Ptr("my-image"),
		// 					ImageVersion: to.Ptr("4.0.0"),
		// 					Status: to.Ptr(armdevopsinfrastructure.ResourceStatusReady),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/a2e95d27-c161-4b61-bda4-11512c14c2c2/resourceGroups/my-resource-group/providers/Microsoft.DevOpsInfrastructure/pools/my-devops-pool/resources/my-dev-ops-pool:04dcde21-626e-5a7e-8659-ce12f9284b29:dd8cc705c_1"),
		// 				Name: to.Ptr("dd8cc705c000001"),
		// 				Properties: &armdevopsinfrastructure.ResourceDetailsObjectProperties{
		// 					Image: to.Ptr("my-image"),
		// 					ImageVersion: to.Ptr("4.0.0"),
		// 					Status: to.Ptr(armdevopsinfrastructure.ResourceStatusAllocated),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
