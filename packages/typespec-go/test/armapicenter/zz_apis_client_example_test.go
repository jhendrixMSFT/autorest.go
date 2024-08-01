// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter_test

import (
	"armapicenter"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: /tsp/ApiCenter.Management/examples/2024-03-15-preview/Apis_CreateOrUpdate.json
func ExampleApisClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApisClient().CreateOrUpdate(ctx, "contoso-resources", "contoso", "default", "echo-api", armapicenter.API{
		Properties: &armapicenter.APIProperties{
			Title:          to.Ptr("Echo API"),
			Description:    to.Ptr("A simple HTTP request/response service."),
			LifecycleStage: to.Ptr(armapicenter.LifecycleStageDesign),
			Kind:           to.Ptr(armapicenter.APIKindRest),
			TermsOfService: &armapicenter.TermsOfService{
				URL: to.Ptr("https://contoso.com/terms-of-service"),
			},
			License: &armapicenter.License{
				URL: to.Ptr("https://contoso.com/license"),
			},
			ExternalDocumentation: []*armapicenter.ExternalDocumentation{
				{
					Title: to.Ptr("Onboarding docs"),
					URL:   to.Ptr("https://docs.contoso.com"),
				},
			},
			CustomProperties: &armapicenter.CustomProperties{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapicenter.ApisClientCreateOrUpdateResponse{
	// 	API: &armapicenter.API{
	// 		Type: to.Ptr("Microsoft.ApiCenter/services/apis"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/apis/echo-api"),
	// 		Name: to.Ptr("echo-api"),
	// 		SystemData: &armapicenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128871Z"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.1288716Z"); return t}()),
	// 		},
	// 		Properties: &armapicenter.APIProperties{
	// 			Title: to.Ptr("Echo API"),
	// 			Description: to.Ptr("A simple HTTP request/response service."),
	// 			LifecycleStage: to.Ptr(armapicenter.LifecycleStageDesign),
	// 			Kind: to.Ptr(armapicenter.APIKindRest),
	// 			TermsOfService: &armapicenter.TermsOfService{
	// 				URL: to.Ptr("https://contoso.com/terms-of-service"),
	// 			},
	// 			License: &armapicenter.License{
	// 				URL: to.Ptr("https://contoso.com/license"),
	// 			},
	// 			ExternalDocumentation: []*armapicenter.ExternalDocumentation{
	// 				{
	// 					Title: to.Ptr("Onboarding docs"),
	// 					URL: to.Ptr("https://docs.contoso.com"),
	// 				},
	// 			},
	// 			CustomProperties: &armapicenter.CustomProperties{
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: /tsp/ApiCenter.Management/examples/2024-03-15-preview/Apis_Delete.json
func ExampleApisClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApisClient().Delete(ctx, "contoso-resources", "contoso", "default", "echo-api", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapicenter.ApisClientDeleteResponse{
	// }
}

// Generated from example definition: /tsp/ApiCenter.Management/examples/2024-03-15-preview/Apis_Get.json
func ExampleApisClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApisClient().Get(ctx, "contoso-resources", "contoso", "default", "echo-api", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapicenter.ApisClientGetResponse{
	// 	API: &armapicenter.API{
	// 		Type: to.Ptr("Microsoft.ApiCenter/services/apis"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/apis/echo-api"),
	// 		Name: to.Ptr("public"),
	// 		SystemData: &armapicenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128871Z"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.1288716Z"); return t}()),
	// 		},
	// 		Properties: &armapicenter.APIProperties{
	// 			Title: to.Ptr("Echo API"),
	// 			Description: to.Ptr("A simple HTTP request/response service."),
	// 			LifecycleStage: to.Ptr(armapicenter.LifecycleStageDesign),
	// 			Kind: to.Ptr(armapicenter.APIKindRest),
	// 			TermsOfService: &armapicenter.TermsOfService{
	// 				URL: to.Ptr("https://contoso.com/terms-of-service"),
	// 			},
	// 			License: &armapicenter.License{
	// 				URL: to.Ptr("https://contoso.com/license"),
	// 			},
	// 			ExternalDocumentation: []*armapicenter.ExternalDocumentation{
	// 				{
	// 					Title: to.Ptr("Onboarding docs"),
	// 					URL: to.Ptr("https://docs.contoso.com"),
	// 				},
	// 			},
	// 			CustomProperties: &armapicenter.CustomProperties{
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: /tsp/ApiCenter.Management/examples/2024-03-15-preview/Apis_Head.json
func ExampleApisClient_Head() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApisClient().Head(ctx, "contoso-resources", "contoso", "default", "echo-api", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapicenter.ApisClientHeadResponse{
	// }
}

// Generated from example definition: /tsp/ApiCenter.Management/examples/2024-03-15-preview/Apis_List.json
func ExampleApisClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApisClient().NewListPager("contoso-resources", "contoso", "default", nil)
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
		// page = armapicenter.ApisClientListResponse{
		// 	APIListResult: armapicenter.APIListResult{
		// 		Value: []*armapicenter.API{
		// 			{
		// 				Type: to.Ptr("Microsoft.ApiCenter/services/environments"),
		// 				ID: to.Ptr("/subscriptions/a200340d-6b82-494d-9dbf-687ba6e33f9e/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/apis/echo-api"),
		// 				Name: to.Ptr("echo-api"),
		// 				SystemData: &armapicenter.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128871Z"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.1288716Z"); return t}()),
		// 				},
		// 				Properties: &armapicenter.APIProperties{
		// 					Title: to.Ptr("Echo API"),
		// 					Description: to.Ptr("A simple HTTP request/response service."),
		// 					LifecycleStage: to.Ptr(armapicenter.LifecycleStageDesign),
		// 					Kind: to.Ptr(armapicenter.APIKindRest),
		// 					TermsOfService: &armapicenter.TermsOfService{
		// 						URL: to.Ptr("https://contoso.com/terms-of-service"),
		// 					},
		// 					License: &armapicenter.License{
		// 						URL: to.Ptr("https://contoso.com/license"),
		// 					},
		// 					ExternalDocumentation: []*armapicenter.ExternalDocumentation{
		// 						{
		// 							Title: to.Ptr("Onboarding docs"),
		// 							URL: to.Ptr("https://docs.contoso.com"),
		// 						},
		// 					},
		// 					CustomProperties: &armapicenter.CustomProperties{
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
