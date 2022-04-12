//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Create.json
func ExampleDataFlowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<data-flow-name>",
		armdatafactory.DataFlowResource{
			Properties: &armdatafactory.MappingDataFlow{
				Type:        to.Ptr("<type>"),
				Description: to.Ptr("<description>"),
				TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
					Script: to.Ptr("<script>"),
					Sinks: []*armdatafactory.DataFlowSink{
						{
							Name: to.Ptr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("<reference-name>"),
							},
						},
						{
							Name: to.Ptr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("<reference-name>"),
							},
						}},
					Sources: []*armdatafactory.DataFlowSource{
						{
							Name: to.Ptr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("<reference-name>"),
							},
						},
						{
							Name: to.Ptr("<name>"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("<reference-name>"),
							},
						}},
				},
			},
		},
		&armdatafactory.DataFlowsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Get.json
func ExampleDataFlowsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<data-flow-name>",
		&armdatafactory.DataFlowsClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Delete.json
func ExampleDataFlowsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<data-flow-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_ListByFactory.json
func ExampleDataFlowsClient_ListByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByFactory("<resource-group-name>",
		"<factory-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
