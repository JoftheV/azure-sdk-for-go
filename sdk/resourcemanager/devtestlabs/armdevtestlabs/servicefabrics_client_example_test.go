//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_List.json
func ExampleServiceFabricsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceFabricsClient().NewListPager("resourceGroupName", "{labName}", "{userName}", &armdevtestlabs.ServiceFabricsClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
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
		// page.ServiceFabricList = armdevtestlabs.ServiceFabricList{
		// 	Value: []*armdevtestlabs.ServiceFabric{
		// 		{
		// 			Name: to.Ptr("{serviceFabricName}"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}"),
		// 			Location: to.Ptr("{location}"),
		// 			Tags: map[string]*string{
		// 				"tagName1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armdevtestlabs.ServiceFabricProperties{
		// 				ApplicableSchedule: &armdevtestlabs.ApplicableSchedule{
		// 					Name: to.Ptr("{scheduleName}"),
		// 					Type: to.Ptr("{scheduleType}"),
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{scheduleName}"),
		// 					Location: to.Ptr("{location}"),
		// 					Tags: map[string]*string{
		// 						"tagName1": to.Ptr("tagValue1"),
		// 					},
		// 					Properties: &armdevtestlabs.ApplicableScheduleProperties{
		// 						LabVMsShutdown: &armdevtestlabs.Schedule{
		// 							Name: to.Ptr("{scheduleName}"),
		// 							Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
		// 							Location: to.Ptr("{location}"),
		// 							Tags: map[string]*string{
		// 								"tagName1": to.Ptr("tagValue1"),
		// 							},
		// 							Properties: &armdevtestlabs.ScheduleProperties{
		// 								CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
		// 								DailyRecurrence: &armdevtestlabs.DayDetails{
		// 									Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
		// 								},
		// 								HourlyRecurrence: &armdevtestlabs.HourDetails{
		// 									Minute: to.Ptr[int32](30),
		// 								},
		// 								NotificationSettings: &armdevtestlabs.NotificationSettings{
		// 									EmailRecipient: to.Ptr("{email}"),
		// 									NotificationLocale: to.Ptr("EN"),
		// 									Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
		// 									TimeInMinutes: to.Ptr[int32](15),
		// 									WebhookURL: to.Ptr("{webhookUrl}"),
		// 								},
		// 								ProvisioningState: to.Ptr("Succeeded"),
		// 								Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
		// 								TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
		// 								TaskType: to.Ptr("{myLabVmTaskType}"),
		// 								TimeZoneID: to.Ptr("Pacific Standard Time"),
		// 								UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 								WeeklyRecurrence: &armdevtestlabs.WeekDetails{
		// 									Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
		// 									Weekdays: []*string{
		// 										to.Ptr("Monday"),
		// 										to.Ptr("Wednesday"),
		// 										to.Ptr("Friday")},
		// 									},
		// 								},
		// 							},
		// 							LabVMsStartup: &armdevtestlabs.Schedule{
		// 								Name: to.Ptr("{scheduleName}"),
		// 								Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
		// 								Location: to.Ptr("{location}"),
		// 								Tags: map[string]*string{
		// 									"tagName1": to.Ptr("tagValue1"),
		// 								},
		// 								Properties: &armdevtestlabs.ScheduleProperties{
		// 									CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
		// 									DailyRecurrence: &armdevtestlabs.DayDetails{
		// 										Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
		// 									},
		// 									HourlyRecurrence: &armdevtestlabs.HourDetails{
		// 										Minute: to.Ptr[int32](30),
		// 									},
		// 									NotificationSettings: &armdevtestlabs.NotificationSettings{
		// 										EmailRecipient: to.Ptr("{email}"),
		// 										NotificationLocale: to.Ptr("EN"),
		// 										Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
		// 										TimeInMinutes: to.Ptr[int32](15),
		// 										WebhookURL: to.Ptr("{webhookUrl}"),
		// 									},
		// 									ProvisioningState: to.Ptr("Succeeded"),
		// 									Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
		// 									TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
		// 									TaskType: to.Ptr("{myLabVmTaskType}"),
		// 									TimeZoneID: to.Ptr("Pacific Standard Time"),
		// 									UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 									WeeklyRecurrence: &armdevtestlabs.WeekDetails{
		// 										Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
		// 										Weekdays: []*string{
		// 											to.Ptr("Monday"),
		// 											to.Ptr("Wednesday"),
		// 											to.Ptr("Friday")},
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						EnvironmentID: to.Ptr("{environmentId}"),
		// 						ExternalServiceFabricID: to.Ptr("{serviceFabricId}"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_Get.json
func ExampleServiceFabricsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceFabricsClient().Get(ctx, "resourceGroupName", "{labName}", "{userName}", "{serviceFabricName}", &armdevtestlabs.ServiceFabricsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceFabric = armdevtestlabs.ServiceFabric{
	// 	Name: to.Ptr("{serviceFabricName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ServiceFabricProperties{
	// 		ApplicableSchedule: &armdevtestlabs.ApplicableSchedule{
	// 			Name: to.Ptr("{scheduleName}"),
	// 			Type: to.Ptr("{scheduleType}"),
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{scheduleName}"),
	// 			Location: to.Ptr("{location}"),
	// 			Tags: map[string]*string{
	// 				"tagName1": to.Ptr("tagValue1"),
	// 			},
	// 			Properties: &armdevtestlabs.ApplicableScheduleProperties{
	// 				LabVMsShutdown: &armdevtestlabs.Schedule{
	// 					Name: to.Ptr("{scheduleName}"),
	// 					Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
	// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
	// 					Location: to.Ptr("{location}"),
	// 					Tags: map[string]*string{
	// 						"tagName1": to.Ptr("tagValue1"),
	// 					},
	// 					Properties: &armdevtestlabs.ScheduleProperties{
	// 						CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
	// 						DailyRecurrence: &armdevtestlabs.DayDetails{
	// 							Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
	// 						},
	// 						HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 							Minute: to.Ptr[int32](30),
	// 						},
	// 						NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 							EmailRecipient: to.Ptr("{email}"),
	// 							NotificationLocale: to.Ptr("EN"),
	// 							Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 							TimeInMinutes: to.Ptr[int32](15),
	// 							WebhookURL: to.Ptr("{webhookUrl}"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 						TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
	// 						TaskType: to.Ptr("{myLabVmTaskType}"),
	// 						TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 						UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 						WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 							Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
	// 							Weekdays: []*string{
	// 								to.Ptr("Monday"),
	// 								to.Ptr("Wednesday"),
	// 								to.Ptr("Friday")},
	// 							},
	// 						},
	// 					},
	// 					LabVMsStartup: &armdevtestlabs.Schedule{
	// 						Name: to.Ptr("{scheduleName}"),
	// 						Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
	// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
	// 						Location: to.Ptr("{location}"),
	// 						Tags: map[string]*string{
	// 							"tagName1": to.Ptr("tagValue1"),
	// 						},
	// 						Properties: &armdevtestlabs.ScheduleProperties{
	// 							CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
	// 							DailyRecurrence: &armdevtestlabs.DayDetails{
	// 								Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
	// 							},
	// 							HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 								Minute: to.Ptr[int32](30),
	// 							},
	// 							NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 								EmailRecipient: to.Ptr("{email}"),
	// 								NotificationLocale: to.Ptr("EN"),
	// 								Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 								TimeInMinutes: to.Ptr[int32](15),
	// 								WebhookURL: to.Ptr("{webhookUrl}"),
	// 							},
	// 							ProvisioningState: to.Ptr("Succeeded"),
	// 							Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 							TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
	// 							TaskType: to.Ptr("{myLabVmTaskType}"),
	// 							TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 							UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 							WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 								Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
	// 								Weekdays: []*string{
	// 									to.Ptr("Monday"),
	// 									to.Ptr("Wednesday"),
	// 									to.Ptr("Friday")},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				EnvironmentID: to.Ptr("{environmentId}"),
	// 				ExternalServiceFabricID: to.Ptr("{serviceFabricId}"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_CreateOrUpdate.json
func ExampleServiceFabricsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceFabricsClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "{labName}", "{userName}", "{serviceFabricName}", armdevtestlabs.ServiceFabric{
		Location: to.Ptr("{location}"),
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
		},
		Properties: &armdevtestlabs.ServiceFabricProperties{
			EnvironmentID:           to.Ptr("{environmentId}"),
			ExternalServiceFabricID: to.Ptr("{serviceFabricId}"),
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
	// res.ServiceFabric = armdevtestlabs.ServiceFabric{
	// 	Name: to.Ptr("{serviceFabricName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ServiceFabricProperties{
	// 		ApplicableSchedule: &armdevtestlabs.ApplicableSchedule{
	// 			Name: to.Ptr("{scheduleName}"),
	// 			Type: to.Ptr("{scheduleType}"),
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{scheduleName}"),
	// 			Location: to.Ptr("{location}"),
	// 			Tags: map[string]*string{
	// 				"tagName1": to.Ptr("tagValue1"),
	// 			},
	// 			Properties: &armdevtestlabs.ApplicableScheduleProperties{
	// 				LabVMsShutdown: &armdevtestlabs.Schedule{
	// 					Name: to.Ptr("{scheduleName}"),
	// 					Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
	// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
	// 					Location: to.Ptr("{location}"),
	// 					Tags: map[string]*string{
	// 						"tagName1": to.Ptr("tagValue1"),
	// 					},
	// 					Properties: &armdevtestlabs.ScheduleProperties{
	// 						CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
	// 						DailyRecurrence: &armdevtestlabs.DayDetails{
	// 							Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
	// 						},
	// 						HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 							Minute: to.Ptr[int32](30),
	// 						},
	// 						NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 							EmailRecipient: to.Ptr("{email}"),
	// 							NotificationLocale: to.Ptr("EN"),
	// 							Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 							TimeInMinutes: to.Ptr[int32](15),
	// 							WebhookURL: to.Ptr("{webhookUrl}"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 						TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
	// 						TaskType: to.Ptr("{myLabVmTaskType}"),
	// 						TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 						UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 						WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 							Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
	// 							Weekdays: []*string{
	// 								to.Ptr("Monday"),
	// 								to.Ptr("Wednesday"),
	// 								to.Ptr("Friday")},
	// 							},
	// 						},
	// 					},
	// 					LabVMsStartup: &armdevtestlabs.Schedule{
	// 						Name: to.Ptr("{scheduleName}"),
	// 						Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
	// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
	// 						Location: to.Ptr("{location}"),
	// 						Tags: map[string]*string{
	// 							"tagName1": to.Ptr("tagValue1"),
	// 						},
	// 						Properties: &armdevtestlabs.ScheduleProperties{
	// 							CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
	// 							DailyRecurrence: &armdevtestlabs.DayDetails{
	// 								Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
	// 							},
	// 							HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 								Minute: to.Ptr[int32](30),
	// 							},
	// 							NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 								EmailRecipient: to.Ptr("{email}"),
	// 								NotificationLocale: to.Ptr("EN"),
	// 								Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 								TimeInMinutes: to.Ptr[int32](15),
	// 								WebhookURL: to.Ptr("{webhookUrl}"),
	// 							},
	// 							ProvisioningState: to.Ptr("Succeeded"),
	// 							Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 							TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
	// 							TaskType: to.Ptr("{myLabVmTaskType}"),
	// 							TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 							UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 							WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 								Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
	// 								Weekdays: []*string{
	// 									to.Ptr("Monday"),
	// 									to.Ptr("Wednesday"),
	// 									to.Ptr("Friday")},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				EnvironmentID: to.Ptr("{environmentId}"),
	// 				ExternalServiceFabricID: to.Ptr("{serviceFabricId}"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_Delete.json
func ExampleServiceFabricsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceFabricsClient().BeginDelete(ctx, "resourceGroupName", "{labName}", "{userName}", "{serviceFabricName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_Update.json
func ExampleServiceFabricsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceFabricsClient().Update(ctx, "resourceGroupName", "{labName}", "{userName}", "{serviceFabricName}", armdevtestlabs.ServiceFabricFragment{
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceFabric = armdevtestlabs.ServiceFabric{
	// 	Name: to.Ptr("{serviceFabricName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/users/serviceFabrics"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ServiceFabricProperties{
	// 		ApplicableSchedule: &armdevtestlabs.ApplicableSchedule{
	// 			Name: to.Ptr("{scheduleName}"),
	// 			Type: to.Ptr("{scheduleType}"),
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{scheduleName}"),
	// 			Location: to.Ptr("{location}"),
	// 			Tags: map[string]*string{
	// 				"tagName1": to.Ptr("tagValue1"),
	// 			},
	// 			Properties: &armdevtestlabs.ApplicableScheduleProperties{
	// 				LabVMsShutdown: &armdevtestlabs.Schedule{
	// 					Name: to.Ptr("{scheduleName}"),
	// 					Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
	// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
	// 					Location: to.Ptr("{location}"),
	// 					Tags: map[string]*string{
	// 						"tagName1": to.Ptr("tagValue1"),
	// 					},
	// 					Properties: &armdevtestlabs.ScheduleProperties{
	// 						CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
	// 						DailyRecurrence: &armdevtestlabs.DayDetails{
	// 							Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
	// 						},
	// 						HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 							Minute: to.Ptr[int32](30),
	// 						},
	// 						NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 							EmailRecipient: to.Ptr("{email}"),
	// 							NotificationLocale: to.Ptr("EN"),
	// 							Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 							TimeInMinutes: to.Ptr[int32](15),
	// 							WebhookURL: to.Ptr("{webhookUrl}"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 						TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
	// 						TaskType: to.Ptr("{myLabVmTaskType}"),
	// 						TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 						UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 						WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 							Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
	// 							Weekdays: []*string{
	// 								to.Ptr("Monday"),
	// 								to.Ptr("Wednesday"),
	// 								to.Ptr("Friday")},
	// 							},
	// 						},
	// 					},
	// 					LabVMsStartup: &armdevtestlabs.Schedule{
	// 						Name: to.Ptr("{scheduleName}"),
	// 						Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
	// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
	// 						Location: to.Ptr("{location}"),
	// 						Tags: map[string]*string{
	// 							"tagName1": to.Ptr("tagValue1"),
	// 						},
	// 						Properties: &armdevtestlabs.ScheduleProperties{
	// 							CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
	// 							DailyRecurrence: &armdevtestlabs.DayDetails{
	// 								Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
	// 							},
	// 							HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 								Minute: to.Ptr[int32](30),
	// 							},
	// 							NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 								EmailRecipient: to.Ptr("{email}"),
	// 								NotificationLocale: to.Ptr("EN"),
	// 								Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 								TimeInMinutes: to.Ptr[int32](15),
	// 								WebhookURL: to.Ptr("{webhookUrl}"),
	// 							},
	// 							ProvisioningState: to.Ptr("Succeeded"),
	// 							Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 							TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
	// 							TaskType: to.Ptr("{myLabVmTaskType}"),
	// 							TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 							UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 							WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 								Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
	// 								Weekdays: []*string{
	// 									to.Ptr("Monday"),
	// 									to.Ptr("Wednesday"),
	// 									to.Ptr("Friday")},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				EnvironmentID: to.Ptr("{environmentId}"),
	// 				ExternalServiceFabricID: to.Ptr("{serviceFabricId}"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_ListApplicableSchedules.json
func ExampleServiceFabricsClient_ListApplicableSchedules() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceFabricsClient().ListApplicableSchedules(ctx, "resourceGroupName", "{labName}", "{userName}", "{serviceFabricName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicableSchedule = armdevtestlabs.ApplicableSchedule{
	// 	Name: to.Ptr("{scheduleName}"),
	// 	Type: to.Ptr("{scheduleType}"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{scheduleName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ApplicableScheduleProperties{
	// 		LabVMsShutdown: &armdevtestlabs.Schedule{
	// 			Name: to.Ptr("{scheduleName}"),
	// 			Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
	// 			Location: to.Ptr("{location}"),
	// 			Tags: map[string]*string{
	// 				"tagName1": to.Ptr("tagValue1"),
	// 			},
	// 			Properties: &armdevtestlabs.ScheduleProperties{
	// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
	// 				DailyRecurrence: &armdevtestlabs.DayDetails{
	// 					Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
	// 				},
	// 				HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 					Minute: to.Ptr[int32](30),
	// 				},
	// 				NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 					EmailRecipient: to.Ptr("{email}"),
	// 					NotificationLocale: to.Ptr("EN"),
	// 					Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 					TimeInMinutes: to.Ptr[int32](15),
	// 					WebhookURL: to.Ptr("{webhookUrl}"),
	// 				},
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 				TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
	// 				TaskType: to.Ptr("{myLabVmTaskType}"),
	// 				TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 				WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 					Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
	// 					Weekdays: []*string{
	// 						to.Ptr("Monday"),
	// 						to.Ptr("Wednesday"),
	// 						to.Ptr("Friday")},
	// 					},
	// 				},
	// 			},
	// 			LabVMsStartup: &armdevtestlabs.Schedule{
	// 				Name: to.Ptr("{scheduleName}"),
	// 				Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
	// 				Location: to.Ptr("{location}"),
	// 				Tags: map[string]*string{
	// 					"tagName1": to.Ptr("tagValue1"),
	// 				},
	// 				Properties: &armdevtestlabs.ScheduleProperties{
	// 					CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
	// 					DailyRecurrence: &armdevtestlabs.DayDetails{
	// 						Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
	// 					},
	// 					HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 						Minute: to.Ptr[int32](30),
	// 					},
	// 					NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 						EmailRecipient: to.Ptr("{email}"),
	// 						NotificationLocale: to.Ptr("EN"),
	// 						Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 						TimeInMinutes: to.Ptr[int32](15),
	// 						WebhookURL: to.Ptr("{webhookUrl}"),
	// 					},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 					TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
	// 					TaskType: to.Ptr("{myLabVmTaskType}"),
	// 					TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 					UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 					WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 						Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
	// 						Weekdays: []*string{
	// 							to.Ptr("Monday"),
	// 							to.Ptr("Wednesday"),
	// 							to.Ptr("Friday")},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_Start.json
func ExampleServiceFabricsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceFabricsClient().BeginStart(ctx, "resourceGroupName", "{labName}", "{userName}", "{serviceFabricName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_Stop.json
func ExampleServiceFabricsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceFabricsClient().BeginStop(ctx, "resourceGroupName", "{labName}", "{userName}", "{serviceFabricName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}