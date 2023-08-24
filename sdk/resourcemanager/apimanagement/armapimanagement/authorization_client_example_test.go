//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListAuthorizationsAuthCode.json
func ExampleAuthorizationClient_NewListByAuthorizationProviderPager_apiManagementListAuthorizationsAuthCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizationClient().NewListByAuthorizationProviderPager("rg1", "apimService1", "aadwithauthcode", &armapimanagement.AuthorizationClientListByAuthorizationProviderOptions{Filter: nil,
		Top:  nil,
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
		// page.AuthorizationCollection = armapimanagement.AuthorizationCollection{
		// 	Value: []*armapimanagement.AuthorizationContract{
		// 		{
		// 			Name: to.Ptr("authz1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz1"),
		// 			Properties: &armapimanagement.AuthorizationContractProperties{
		// 				AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
		// 				OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
		// 				Status: to.Ptr("Connected"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("authz2"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz2"),
		// 			Properties: &armapimanagement.AuthorizationContractProperties{
		// 				AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
		// 				Error: &armapimanagement.AuthorizationError{
		// 					Code: to.Ptr("Unauthenticated"),
		// 					Message: to.Ptr("This connection is not authenticated."),
		// 				},
		// 				OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
		// 				Status: to.Ptr("Error"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListAuthorizationsClientCred.json
func ExampleAuthorizationClient_NewListByAuthorizationProviderPager_apiManagementListAuthorizationsClientCred() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizationClient().NewListByAuthorizationProviderPager("rg1", "apimService1", "aadwithclientcred", &armapimanagement.AuthorizationClientListByAuthorizationProviderOptions{Filter: nil,
		Top:  nil,
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
		// page.AuthorizationCollection = armapimanagement.AuthorizationCollection{
		// 	Value: []*armapimanagement.AuthorizationContract{
		// 		{
		// 			Name: to.Ptr("authz1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithclientcred/authorizations/authz1"),
		// 			Properties: &armapimanagement.AuthorizationContractProperties{
		// 				AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
		// 				OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeClientCredentials),
		// 				Parameters: map[string]*string{
		// 					"clientId": to.Ptr("53790925-fdd3-4b80-bc7a-4c3aaf25801d"),
		// 				},
		// 				Status: to.Ptr("Connected"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("authz2"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithclientcred/authorizations/authz2"),
		// 			Properties: &armapimanagement.AuthorizationContractProperties{
		// 				AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
		// 				Error: &armapimanagement.AuthorizationError{
		// 					Code: to.Ptr("Unauthorized"),
		// 					Message: to.Ptr("Failed to acquire access token for service using client credentials flow: IdentityProvider=aadwithclientcred. Correlation Id=6299d09b-03b7-46ed-b355-0453451d7e49, UTC TimeStamp=5/14/2022 4:53:19 PM, Error: Failed to exchange client credentials for token. Response code=Unauthorized, Details:\r\n{\"error\":\"invalid_client\",\"error_description\":\"AADSTS7000215: Invalid client secret provided. Ensure the secret being sent in the request is the client secret value, not the client secret ID, for a secret added to app '53790925-fdd3-4b80-bc7a-4c3aaf25801d'.\\r\\nTrace ID: 4a18d3cd-9ad5-4664-b3eb-daaa2f435f00\\r\\nCorrelation ID: dde60c16-35cb-4572-8226-bfb4233af8d7\\r\\nTimestamp: 2022-05-14 16:53:19Z\",\"error_codes\":[7000215],\"timestamp\":\"2022-05-14 16:53:19Z\",\"trace_id\":\"4a18d3cd-9ad5-4664-b3eb-daaa2f435f00\",\"correlation_id\":\"dde60c16-35cb-4572-8226-bfb4233af8d7\",\"error_uri\":\"https://login.windows.net/error?code=7000215\"}"),
		// 				},
		// 				OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeClientCredentials),
		// 				Parameters: map[string]*string{
		// 					"clientId": to.Ptr("53790925-fdd3-4b80-bc7a-4c3aaf25801d"),
		// 				},
		// 				Status: to.Ptr("Error"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetAuthorization.json
func ExampleAuthorizationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationClient().Get(ctx, "rg1", "apimService1", "aadwithauthcode", "authz1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationContract = armapimanagement.AuthorizationContract{
	// 	Name: to.Ptr("authz1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz1"),
	// 	Properties: &armapimanagement.AuthorizationContractProperties{
	// 		AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
	// 		OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
	// 		Status: to.Ptr("Connected"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationAADAuthCode.json
func ExampleAuthorizationClient_CreateOrUpdate_apiManagementCreateAuthorizationAadAuthCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationClient().CreateOrUpdate(ctx, "rg1", "apimService1", "aadwithauthcode", "authz2", armapimanagement.AuthorizationContract{
		Properties: &armapimanagement.AuthorizationContractProperties{
			AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
			OAuth2GrantType:   to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
		},
	}, &armapimanagement.AuthorizationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationContract = armapimanagement.AuthorizationContract{
	// 	Name: to.Ptr("authz2"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz2"),
	// 	Properties: &armapimanagement.AuthorizationContractProperties{
	// 		AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
	// 		Error: &armapimanagement.AuthorizationError{
	// 			Code: to.Ptr("Unauthenticated"),
	// 			Message: to.Ptr("This connection is not authenticated."),
	// 		},
	// 		OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
	// 		Status: to.Ptr("Error"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationAADClientCred.json
func ExampleAuthorizationClient_CreateOrUpdate_apiManagementCreateAuthorizationAadClientCred() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationClient().CreateOrUpdate(ctx, "rg1", "apimService1", "aadwithclientcred", "authz1", armapimanagement.AuthorizationContract{
		Properties: &armapimanagement.AuthorizationContractProperties{
			AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
			OAuth2GrantType:   to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
			Parameters: map[string]*string{
				"clientId":     to.Ptr("53790925-fdd3-4b80-bc7a-4c3aaf25801d"),
				"clientSecret": to.Ptr("FcJkQ3iPSaKAQRA7Ft8Q~fZ1X5vKmqzUAfJagcJ8"),
			},
		},
	}, &armapimanagement.AuthorizationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationContract = armapimanagement.AuthorizationContract{
	// 	Name: to.Ptr("authz1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithclientcred/authorizations/authz1"),
	// 	Properties: &armapimanagement.AuthorizationContractProperties{
	// 		AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
	// 		OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeClientCredentials),
	// 		Parameters: map[string]*string{
	// 			"clientId": to.Ptr("53790925-fdd3-4b80-bc7a-4c3aaf25801d"),
	// 		},
	// 		Status: to.Ptr("Connected"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteAuthorization.json
func ExampleAuthorizationClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAuthorizationClient().Delete(ctx, "rg1", "apimService1", "aadwithauthcode", "authz1", "*", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPostAuthorizationConfirmConsentCodeRequest.json
func ExampleAuthorizationClient_ConfirmConsentCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAuthorizationClient().ConfirmConsentCode(ctx, "rg1", "apimService1", "aadwithauthcode", "authz1", armapimanagement.AuthorizationConfirmConsentCodeRequestContract{
		ConsentCode: to.Ptr("theconsentcode"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}