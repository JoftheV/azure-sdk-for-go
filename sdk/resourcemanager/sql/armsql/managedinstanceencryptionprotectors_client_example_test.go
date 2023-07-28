//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceEncryptionProtectorRevalidate.json
func ExampleManagedInstanceEncryptionProtectorsClient_BeginRevalidate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstanceEncryptionProtectorsClient().BeginRevalidate(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", armsql.EncryptionProtectorNameCurrent, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceEncryptionProtectorList.json
func ExampleManagedInstanceEncryptionProtectorsClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstanceEncryptionProtectorsClient().NewListByInstancePager("sqlcrudtest-7398", "sqlcrudtest-4645", nil)
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
		// page.ManagedInstanceEncryptionProtectorListResult = armsql.ManagedInstanceEncryptionProtectorListResult{
		// 	Value: []*armsql.ManagedInstanceEncryptionProtector{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/encryptionProtector"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/encryptionProtector/current"),
		// 			Kind: to.Ptr("azurekeyvault"),
		// 			Properties: &armsql.ManagedInstanceEncryptionProtectorProperties{
		// 				ServerKeyName: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
		// 				ServerKeyType: to.Ptr(armsql.ServerKeyTypeAzureKeyVault),
		// 				URI: to.Ptr("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceEncryptionProtectorGet.json
func ExampleManagedInstanceEncryptionProtectorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedInstanceEncryptionProtectorsClient().Get(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", armsql.EncryptionProtectorNameCurrent, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedInstanceEncryptionProtector = armsql.ManagedInstanceEncryptionProtector{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/encryptionProtector"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/encryptionProtector/current"),
	// 	Kind: to.Ptr("azurekeyvault"),
	// 	Properties: &armsql.ManagedInstanceEncryptionProtectorProperties{
	// 		ServerKeyName: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
	// 		ServerKeyType: to.Ptr(armsql.ServerKeyTypeAzureKeyVault),
	// 		URI: to.Ptr("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceEncryptionProtectorCreateOrUpdateKeyVault.json
func ExampleManagedInstanceEncryptionProtectorsClient_BeginCreateOrUpdate_updateTheEncryptionProtectorToKeyVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstanceEncryptionProtectorsClient().BeginCreateOrUpdate(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", armsql.EncryptionProtectorNameCurrent, armsql.ManagedInstanceEncryptionProtector{
		Properties: &armsql.ManagedInstanceEncryptionProtectorProperties{
			AutoRotationEnabled: to.Ptr(false),
			ServerKeyName:       to.Ptr("someVault_someKey_01234567890123456789012345678901"),
			ServerKeyType:       to.Ptr(armsql.ServerKeyTypeAzureKeyVault),
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
	// res.ManagedInstanceEncryptionProtector = armsql.ManagedInstanceEncryptionProtector{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/encryptionProtector"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/encryptionProtector/current"),
	// 	Kind: to.Ptr("azurekeyvault"),
	// 	Properties: &armsql.ManagedInstanceEncryptionProtectorProperties{
	// 		AutoRotationEnabled: to.Ptr(false),
	// 		ServerKeyName: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
	// 		ServerKeyType: to.Ptr(armsql.ServerKeyTypeAzureKeyVault),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceEncryptionProtectorCreateOrUpdateServiceManaged.json
func ExampleManagedInstanceEncryptionProtectorsClient_BeginCreateOrUpdate_updateTheEncryptionProtectorToServiceManaged() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstanceEncryptionProtectorsClient().BeginCreateOrUpdate(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", armsql.EncryptionProtectorNameCurrent, armsql.ManagedInstanceEncryptionProtector{
		Properties: &armsql.ManagedInstanceEncryptionProtectorProperties{
			ServerKeyName: to.Ptr("ServiceManaged"),
			ServerKeyType: to.Ptr(armsql.ServerKeyTypeServiceManaged),
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
	// res.ManagedInstanceEncryptionProtector = armsql.ManagedInstanceEncryptionProtector{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/encryptionProtector"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/encryptionProtector/current"),
	// 	Kind: to.Ptr("servicemanaged"),
	// 	Properties: &armsql.ManagedInstanceEncryptionProtectorProperties{
	// 		ServerKeyName: to.Ptr("ServiceManaged"),
	// 		ServerKeyType: to.Ptr(armsql.ServerKeyTypeServiceManaged),
	// 	},
	// }
}
