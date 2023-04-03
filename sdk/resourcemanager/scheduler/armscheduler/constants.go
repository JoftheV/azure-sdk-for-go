//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armscheduler

const (
	moduleName    = "armscheduler"
	moduleVersion = "v1.1.0"
)

type DayOfWeek string

const (
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
)

// PossibleDayOfWeekValues returns the possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekSunday,
		DayOfWeekMonday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
		DayOfWeekThursday,
		DayOfWeekFriday,
		DayOfWeekSaturday,
	}
}

// HTTPAuthenticationType - Gets or sets the HTTP authentication type.
type HTTPAuthenticationType string

const (
	HTTPAuthenticationTypeNotSpecified         HTTPAuthenticationType = "NotSpecified"
	HTTPAuthenticationTypeClientCertificate    HTTPAuthenticationType = "ClientCertificate"
	HTTPAuthenticationTypeActiveDirectoryOAuth HTTPAuthenticationType = "ActiveDirectoryOAuth"
	HTTPAuthenticationTypeBasic                HTTPAuthenticationType = "Basic"
)

// PossibleHTTPAuthenticationTypeValues returns the possible values for the HTTPAuthenticationType const type.
func PossibleHTTPAuthenticationTypeValues() []HTTPAuthenticationType {
	return []HTTPAuthenticationType{
		HTTPAuthenticationTypeNotSpecified,
		HTTPAuthenticationTypeClientCertificate,
		HTTPAuthenticationTypeActiveDirectoryOAuth,
		HTTPAuthenticationTypeBasic,
	}
}

// JobActionType - Gets or sets the job action type.
type JobActionType string

const (
	JobActionTypeHTTP            JobActionType = "Http"
	JobActionTypeHTTPS           JobActionType = "Https"
	JobActionTypeStorageQueue    JobActionType = "StorageQueue"
	JobActionTypeServiceBusQueue JobActionType = "ServiceBusQueue"
	JobActionTypeServiceBusTopic JobActionType = "ServiceBusTopic"
)

// PossibleJobActionTypeValues returns the possible values for the JobActionType const type.
func PossibleJobActionTypeValues() []JobActionType {
	return []JobActionType{
		JobActionTypeHTTP,
		JobActionTypeHTTPS,
		JobActionTypeStorageQueue,
		JobActionTypeServiceBusQueue,
		JobActionTypeServiceBusTopic,
	}
}

// JobCollectionState - Gets or sets the state.
type JobCollectionState string

const (
	JobCollectionStateEnabled   JobCollectionState = "Enabled"
	JobCollectionStateDisabled  JobCollectionState = "Disabled"
	JobCollectionStateSuspended JobCollectionState = "Suspended"
	JobCollectionStateDeleted   JobCollectionState = "Deleted"
)

// PossibleJobCollectionStateValues returns the possible values for the JobCollectionState const type.
func PossibleJobCollectionStateValues() []JobCollectionState {
	return []JobCollectionState{
		JobCollectionStateEnabled,
		JobCollectionStateDisabled,
		JobCollectionStateSuspended,
		JobCollectionStateDeleted,
	}
}

// JobExecutionStatus - Gets the job execution status.
type JobExecutionStatus string

const (
	JobExecutionStatusCompleted JobExecutionStatus = "Completed"
	JobExecutionStatusFailed    JobExecutionStatus = "Failed"
	JobExecutionStatusPostponed JobExecutionStatus = "Postponed"
)

// PossibleJobExecutionStatusValues returns the possible values for the JobExecutionStatus const type.
func PossibleJobExecutionStatusValues() []JobExecutionStatus {
	return []JobExecutionStatus{
		JobExecutionStatusCompleted,
		JobExecutionStatusFailed,
		JobExecutionStatusPostponed,
	}
}

// JobHistoryActionName - Gets the job history action name.
type JobHistoryActionName string

const (
	JobHistoryActionNameMainAction  JobHistoryActionName = "MainAction"
	JobHistoryActionNameErrorAction JobHistoryActionName = "ErrorAction"
)

// PossibleJobHistoryActionNameValues returns the possible values for the JobHistoryActionName const type.
func PossibleJobHistoryActionNameValues() []JobHistoryActionName {
	return []JobHistoryActionName{
		JobHistoryActionNameMainAction,
		JobHistoryActionNameErrorAction,
	}
}

// JobScheduleDay - Gets or sets the day. Must be one of monday, tuesday, wednesday, thursday, friday, saturday, sunday.
type JobScheduleDay string

const (
	JobScheduleDayMonday    JobScheduleDay = "Monday"
	JobScheduleDayTuesday   JobScheduleDay = "Tuesday"
	JobScheduleDayWednesday JobScheduleDay = "Wednesday"
	JobScheduleDayThursday  JobScheduleDay = "Thursday"
	JobScheduleDayFriday    JobScheduleDay = "Friday"
	JobScheduleDaySaturday  JobScheduleDay = "Saturday"
	JobScheduleDaySunday    JobScheduleDay = "Sunday"
)

// PossibleJobScheduleDayValues returns the possible values for the JobScheduleDay const type.
func PossibleJobScheduleDayValues() []JobScheduleDay {
	return []JobScheduleDay{
		JobScheduleDayMonday,
		JobScheduleDayTuesday,
		JobScheduleDayWednesday,
		JobScheduleDayThursday,
		JobScheduleDayFriday,
		JobScheduleDaySaturday,
		JobScheduleDaySunday,
	}
}

// JobState - Gets or set the job state.
type JobState string

const (
	JobStateEnabled   JobState = "Enabled"
	JobStateDisabled  JobState = "Disabled"
	JobStateFaulted   JobState = "Faulted"
	JobStateCompleted JobState = "Completed"
)

// PossibleJobStateValues returns the possible values for the JobState const type.
func PossibleJobStateValues() []JobState {
	return []JobState{
		JobStateEnabled,
		JobStateDisabled,
		JobStateFaulted,
		JobStateCompleted,
	}
}

// RecurrenceFrequency - Gets or sets the frequency of recurrence (second, minute, hour, day, week, month).
type RecurrenceFrequency string

const (
	RecurrenceFrequencyMinute RecurrenceFrequency = "Minute"
	RecurrenceFrequencyHour   RecurrenceFrequency = "Hour"
	RecurrenceFrequencyDay    RecurrenceFrequency = "Day"
	RecurrenceFrequencyWeek   RecurrenceFrequency = "Week"
	RecurrenceFrequencyMonth  RecurrenceFrequency = "Month"
)

// PossibleRecurrenceFrequencyValues returns the possible values for the RecurrenceFrequency const type.
func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return []RecurrenceFrequency{
		RecurrenceFrequencyMinute,
		RecurrenceFrequencyHour,
		RecurrenceFrequencyDay,
		RecurrenceFrequencyWeek,
		RecurrenceFrequencyMonth,
	}
}

// RetryType - Gets or sets the retry strategy to be used.
type RetryType string

const (
	RetryTypeNone  RetryType = "None"
	RetryTypeFixed RetryType = "Fixed"
)

// PossibleRetryTypeValues returns the possible values for the RetryType const type.
func PossibleRetryTypeValues() []RetryType {
	return []RetryType{
		RetryTypeNone,
		RetryTypeFixed,
	}
}

// SKUDefinition - Gets or set the SKU.
type SKUDefinition string

const (
	SKUDefinitionStandard   SKUDefinition = "Standard"
	SKUDefinitionFree       SKUDefinition = "Free"
	SKUDefinitionP10Premium SKUDefinition = "P10Premium"
	SKUDefinitionP20Premium SKUDefinition = "P20Premium"
)

// PossibleSKUDefinitionValues returns the possible values for the SKUDefinition const type.
func PossibleSKUDefinitionValues() []SKUDefinition {
	return []SKUDefinition{
		SKUDefinitionStandard,
		SKUDefinitionFree,
		SKUDefinitionP10Premium,
		SKUDefinitionP20Premium,
	}
}

// ServiceBusAuthenticationType - Gets or sets the authentication type.
type ServiceBusAuthenticationType string

const (
	ServiceBusAuthenticationTypeNotSpecified    ServiceBusAuthenticationType = "NotSpecified"
	ServiceBusAuthenticationTypeSharedAccessKey ServiceBusAuthenticationType = "SharedAccessKey"
)

// PossibleServiceBusAuthenticationTypeValues returns the possible values for the ServiceBusAuthenticationType const type.
func PossibleServiceBusAuthenticationTypeValues() []ServiceBusAuthenticationType {
	return []ServiceBusAuthenticationType{
		ServiceBusAuthenticationTypeNotSpecified,
		ServiceBusAuthenticationTypeSharedAccessKey,
	}
}

// ServiceBusTransportType - Gets or sets the transport type.
type ServiceBusTransportType string

const (
	ServiceBusTransportTypeNotSpecified ServiceBusTransportType = "NotSpecified"
	ServiceBusTransportTypeNetMessaging ServiceBusTransportType = "NetMessaging"
	ServiceBusTransportTypeAMQP         ServiceBusTransportType = "AMQP"
)

// PossibleServiceBusTransportTypeValues returns the possible values for the ServiceBusTransportType const type.
func PossibleServiceBusTransportTypeValues() []ServiceBusTransportType {
	return []ServiceBusTransportType{
		ServiceBusTransportTypeNotSpecified,
		ServiceBusTransportTypeNetMessaging,
		ServiceBusTransportTypeAMQP,
	}
}