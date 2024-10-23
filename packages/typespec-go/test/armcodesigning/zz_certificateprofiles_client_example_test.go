// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcodesigning_test

import (
	"armcodesigning"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
	"time"
)

// Generated from example definition: 2024-09-30-preview/CertificateProfiles_Create.json
func ExampleCertificateProfilesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcodesigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCertificateProfilesClient().BeginCreate(ctx, "MyResourceGroup", "MyAccount", "profileA", armcodesigning.CertificateProfile{
		Properties: &armcodesigning.CertificateProfileProperties{
			ProfileType:          to.Ptr(armcodesigning.ProfileTypePublicTrust),
			IdentityValidationID: to.Ptr("00000000-1234-5678-3333-444444444444"),
			IncludePostalCode:    to.Ptr(true),
			IncludeStreetAddress: to.Ptr(false),
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
	// res = armcodesigning.CertificateProfilesClientCreateResponse{
	// 	CertificateProfile: &armcodesigning.CertificateProfile{
	// 		Name: to.Ptr("profileA"),
	// 		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts/certificateProfiles"),
	// 		ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount/certificateProfiles/profileA"),
	// 		Properties: &armcodesigning.CertificateProfileProperties{
	// 			Certificates: []*armcodesigning.Certificate{
	// 				{
	// 					CreatedDate: to.Ptr("3/14/2023 5:27:49 PM"),
	// 					ExpiryDate: to.Ptr("3/17/2023 5:27:49 PM"),
	// 					EnhancedKeyUsage: to.Ptr("1.3.6.1.4.1.311.yy.xxxxxxxx.xxxxxxxx.xxxxxxxxx.xxxxxxxx"),
	// 					SerialNumber: to.Ptr("xxxxxxxxxxxxxxxxxx"),
	// 					Status: to.Ptr(armcodesigning.CertificateStatusActive),
	// 					SubjectName: to.Ptr("CN=Contoso Inc, O=Contoso Inc, L=New York, S=New York, C=US"),
	// 					Thumbprint: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
	// 				},
	// 			},
	// 			IdentityValidationID: to.Ptr("00000000-1234-5678-3333-444444444444"),
	// 			IncludeCity: to.Ptr(false),
	// 			IncludeCountry: to.Ptr(false),
	// 			IncludePostalCode: to.Ptr(true),
	// 			IncludeState: to.Ptr(false),
	// 			IncludeStreetAddress: to.Ptr(false),
	// 			ProfileType: to.Ptr(armcodesigning.ProfileTypePublicTrust),
	// 			ProvisioningState: to.Ptr(armcodesigning.ProvisioningStateSucceeded),
	// 			Status: to.Ptr(armcodesigning.CertificateProfileStatusActive),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-09-30-preview/CertificateProfiles_Delete.json
func ExampleCertificateProfilesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcodesigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCertificateProfilesClient().BeginDelete(ctx, "MyResourceGroup", "MyAccount", "profileA", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2024-09-30-preview/CertificateProfiles_Get.json
func ExampleCertificateProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcodesigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateProfilesClient().Get(ctx, "MyResourceGroup", "MyAccount", "profileA", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcodesigning.CertificateProfilesClientGetResponse{
	// 	CertificateProfile: &armcodesigning.CertificateProfile{
	// 		Name: to.Ptr("profileA"),
	// 		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts/certificateProfiles"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount/certificateProfiles/profileA"),
	// 		Properties: &armcodesigning.CertificateProfileProperties{
	// 			Certificates: []*armcodesigning.Certificate{
	// 				{
	// 					CreatedDate: to.Ptr("3/14/2023 5:27:49 PM"),
	// 					ExpiryDate: to.Ptr("3/17/2023 5:27:49 PM"),
	// 					EnhancedKeyUsage: to.Ptr("1.3.6.1.4.1.311.yy.xxxxxxxx.xxxxxxxx.xxxxxxxxx.xxxxxxxx"),
	// 					SerialNumber: to.Ptr("xxxxxxxxxxxxxxxxxx"),
	// 					Status: to.Ptr(armcodesigning.CertificateStatusActive),
	// 					SubjectName: to.Ptr("CN=Contoso Inc, O=Contoso Inc, L=New York, S=New York, C=US"),
	// 					Thumbprint: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
	// 				},
	// 			},
	// 			IdentityValidationID: to.Ptr("123456"),
	// 			IncludeCity: to.Ptr(false),
	// 			IncludeCountry: to.Ptr(false),
	// 			IncludePostalCode: to.Ptr(true),
	// 			IncludeState: to.Ptr(false),
	// 			IncludeStreetAddress: to.Ptr(false),
	// 			ProfileType: to.Ptr(armcodesigning.ProfileTypePublicTrust),
	// 			ProvisioningState: to.Ptr(armcodesigning.ProvisioningStateSucceeded),
	// 			Status: to.Ptr(armcodesigning.CertificateProfileStatusActive),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-09-30-preview/CertificateProfiles_ListByCodeSigningAccount.json
func ExampleCertificateProfilesClient_NewListByCodeSigningAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcodesigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificateProfilesClient().NewListByCodeSigningAccountPager("MyResourceGroup", "MyAccount", nil)
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
		// page = armcodesigning.CertificateProfilesClientListByCodeSigningAccountResponse{
		// 	CertificateProfileListResult: armcodesigning.CertificateProfileListResult{
		// 		Value: []*armcodesigning.CertificateProfile{
		// 			{
		// 				Name: to.Ptr("profileA"),
		// 				Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts/certificateProfiles"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/profileA"),
		// 				Properties: &armcodesigning.CertificateProfileProperties{
		// 					Certificates: []*armcodesigning.Certificate{
		// 						{
		// 							CreatedDate: to.Ptr("3/14/2023 5:27:49 PM"),
		// 							ExpiryDate: to.Ptr("3/17/2023 5:27:49 PM"),
		// 							EnhancedKeyUsage: to.Ptr("1.3.6.1.4.1.311.yy.xxxxxxxx.xxxxxxxx.xxxxxxxxx.xxxxxxxx"),
		// 							SerialNumber: to.Ptr("xxxxxxxxxxxxxxxxxx"),
		// 							Status: to.Ptr(armcodesigning.CertificateStatusActive),
		// 							SubjectName: to.Ptr("CN=Contoso Inc, O=Contoso Inc, L=New York, S=New York, C=US"),
		// 							Thumbprint: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		// 						},
		// 					},
		// 					IdentityValidationID: to.Ptr("123456"),
		// 					IncludeCity: to.Ptr(false),
		// 					IncludeCountry: to.Ptr(false),
		// 					IncludePostalCode: to.Ptr(true),
		// 					IncludeState: to.Ptr(false),
		// 					IncludeStreetAddress: to.Ptr(false),
		// 					ProfileType: to.Ptr(armcodesigning.ProfileTypePublicTrust),
		// 					ProvisioningState: to.Ptr(armcodesigning.ProvisioningStateSucceeded),
		// 					Status: to.Ptr(armcodesigning.CertificateProfileStatusActive),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2024-09-30-preview/CertificateProfiles_RevokeCertificate.json
func ExampleCertificateProfilesClient_RevokeCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcodesigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCertificateProfilesClient().RevokeCertificate(ctx, "MyResourceGroup", "MyAccount", "profileA", armcodesigning.RevokeCertificate{
		EffectiveAt:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-12T23:40:25+00:00"); return t }()),
		Reason:       to.Ptr("KeyCompromised"),
		Remarks:      to.Ptr("test"),
		SerialNumber: to.Ptr("xxxxxxxxxxxxxxxxxx"),
		Thumbprint:   to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
