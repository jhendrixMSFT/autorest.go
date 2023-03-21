//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/list.json
func ExampleBatchEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBatchEndpointsClient().NewListPager("test-rg", "my-aml-workspace", &armmachinelearningservices.BatchEndpointsClientListOptions{Count: to.Ptr[int32](1),
		Skip: nil,
	})
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
		// page.BatchEndpointTrackedResourceArmPaginatedResult = armmachinelearningservices.BatchEndpointTrackedResourceArmPaginatedResult{
		// 	Value: []*armmachinelearningservices.BatchEndpointData{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armmachinelearningservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("string"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmachinelearningservices.ManagedServiceIdentity{
		// 				Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
		// 					"string": &armmachinelearningservices.UserAssignedIdentity{
		// 						ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 						PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					},
		// 				},
		// 			},
		// 			Kind: to.Ptr("string"),
		// 			Properties: &armmachinelearningservices.BatchEndpointDetails{
		// 				Description: to.Ptr("string"),
		// 				AuthMode: to.Ptr(armmachinelearningservices.EndpointAuthModeAMLToken),
		// 				Properties: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				ScoringURI: to.Ptr("https://www.contoso.com/example"),
		// 				SwaggerURI: to.Ptr("https://www.contoso.com/example"),
		// 				Defaults: &armmachinelearningservices.BatchEndpointDefaults{
		// 					DeploymentName: to.Ptr("string"),
		// 				},
		// 			},
		// 			SKU: &armmachinelearningservices.SKU{
		// 				Name: to.Ptr("string"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Family: to.Ptr("string"),
		// 				Size: to.Ptr("string"),
		// 				Tier: to.Ptr(armmachinelearningservices.SKUTierFree),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/delete.json
func ExampleBatchEndpointsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBatchEndpointsClient().BeginDelete(ctx, "resourceGroup-1234", "testworkspace", "testBatchEndpoint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/get.json
func ExampleBatchEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBatchEndpointsClient().Get(ctx, "test-rg", "my-aml-workspace", "testEndpointName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BatchEndpointData = armmachinelearningservices.BatchEndpointData{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearningservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("string"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmachinelearningservices.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
	// 			"string": &armmachinelearningservices.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr("string"),
	// 	Properties: &armmachinelearningservices.BatchEndpointDetails{
	// 		Description: to.Ptr("string"),
	// 		AuthMode: to.Ptr(armmachinelearningservices.EndpointAuthModeAMLToken),
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		ScoringURI: to.Ptr("https://www.contoso.com/example"),
	// 		SwaggerURI: to.Ptr("https://www.contoso.com/example"),
	// 		Defaults: &armmachinelearningservices.BatchEndpointDefaults{
	// 			DeploymentName: to.Ptr("string"),
	// 		},
	// 	},
	// 	SKU: &armmachinelearningservices.SKU{
	// 		Name: to.Ptr("string"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("string"),
	// 		Size: to.Ptr("string"),
	// 		Tier: to.Ptr(armmachinelearningservices.SKUTierFree),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/update.json
func ExampleBatchEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBatchEndpointsClient().BeginUpdate(ctx, "test-rg", "my-aml-workspace", "testEndpointName", armmachinelearningservices.PartialBatchEndpointPartialTrackedResource{
		Identity: &armmachinelearningservices.PartialManagedServiceIdentity{
			Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
			UserAssignedIdentities: map[string]any{
				"string": map[string]any{},
			},
		},
		Kind:     to.Ptr("string"),
		Location: to.Ptr("string"),
		Properties: &armmachinelearningservices.PartialBatchEndpoint{
			Defaults: &armmachinelearningservices.BatchEndpointDefaults{
				DeploymentName: to.Ptr("string"),
			},
		},
		SKU: &armmachinelearningservices.PartialSKU{
			Name:     to.Ptr("string"),
			Capacity: to.Ptr[int32](1),
			Family:   to.Ptr("string"),
			Size:     to.Ptr("string"),
			Tier:     to.Ptr(armmachinelearningservices.SKUTierFree),
		},
		Tags: map[string]*string{},
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
	// res.BatchEndpointData = armmachinelearningservices.BatchEndpointData{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearningservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("string"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmachinelearningservices.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
	// 			"string": &armmachinelearningservices.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr("string"),
	// 	Properties: &armmachinelearningservices.BatchEndpointDetails{
	// 		Description: to.Ptr("string"),
	// 		AuthMode: to.Ptr(armmachinelearningservices.EndpointAuthModeAMLToken),
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		ScoringURI: to.Ptr("https://www.contoso.com/example"),
	// 		SwaggerURI: to.Ptr("https://www.contoso.com/example"),
	// 		Defaults: &armmachinelearningservices.BatchEndpointDefaults{
	// 			DeploymentName: to.Ptr("string"),
	// 		},
	// 		ProvisioningState: to.Ptr(armmachinelearningservices.EndpointProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armmachinelearningservices.SKU{
	// 		Name: to.Ptr("string"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("string"),
	// 		Size: to.Ptr("string"),
	// 		Tier: to.Ptr(armmachinelearningservices.SKUTierFree),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/createOrUpdate.json
func ExampleBatchEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBatchEndpointsClient().BeginCreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "testEndpointName", armmachinelearningservices.BatchEndpointData{
		Location: to.Ptr("string"),
		Tags:     map[string]*string{},
		Identity: &armmachinelearningservices.ManagedServiceIdentity{
			Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
			UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
				"string": &armmachinelearningservices.UserAssignedIdentity{},
			},
		},
		Kind: to.Ptr("string"),
		Properties: &armmachinelearningservices.BatchEndpointDetails{
			Description: to.Ptr("string"),
			AuthMode:    to.Ptr(armmachinelearningservices.EndpointAuthModeAMLToken),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Defaults: &armmachinelearningservices.BatchEndpointDefaults{
				DeploymentName: to.Ptr("string"),
			},
		},
		SKU: &armmachinelearningservices.SKU{
			Name:     to.Ptr("string"),
			Capacity: to.Ptr[int32](1),
			Family:   to.Ptr("string"),
			Size:     to.Ptr("string"),
			Tier:     to.Ptr(armmachinelearningservices.SKUTierFree),
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
	// res.BatchEndpointData = armmachinelearningservices.BatchEndpointData{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearningservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("string"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmachinelearningservices.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
	// 			"string": &armmachinelearningservices.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr("string"),
	// 	Properties: &armmachinelearningservices.BatchEndpointDetails{
	// 		Description: to.Ptr("string"),
	// 		AuthMode: to.Ptr(armmachinelearningservices.EndpointAuthModeAMLToken),
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		ScoringURI: to.Ptr("https://www.contoso.com/example"),
	// 		SwaggerURI: to.Ptr("https://www.contoso.com/example"),
	// 		Defaults: &armmachinelearningservices.BatchEndpointDefaults{
	// 			DeploymentName: to.Ptr("string"),
	// 		},
	// 	},
	// 	SKU: &armmachinelearningservices.SKU{
	// 		Name: to.Ptr("string"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("string"),
	// 		Size: to.Ptr("string"),
	// 		Tier: to.Ptr(armmachinelearningservices.SKUTierFree),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/listKeys.json
func ExampleBatchEndpointsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBatchEndpointsClient().ListKeys(ctx, "test-rg", "my-aml-workspace", "testEndpointName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EndpointAuthKeys = armmachinelearningservices.EndpointAuthKeys{
	// 	PrimaryKey: to.Ptr("string"),
	// 	SecondaryKey: to.Ptr("string"),
	// }
}
