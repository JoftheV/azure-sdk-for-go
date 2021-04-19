package healthcareapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/healthcareapis/mgmt/2020-03-15/healthcareapis"

// CheckNameAvailabilityParameters input values.
type CheckNameAvailabilityParameters struct {
	// Name - The name of the service instance to check.
	Name *string `json:"name,omitempty"`
	// Type - The fully qualified resource type which includes provider namespace.
	Type *string `json:"type,omitempty"`
}

// ErrorDetails error details.
type ErrorDetails struct {
	// Error - Object containing error details.
	Error *ErrorDetailsInternal `json:"error,omitempty"`
}

// ErrorDetailsInternal error details.
type ErrorDetailsInternal struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The target of the particular error.
	Target *string `json:"target,omitempty"`
}

// Operation service REST API operation.
type Operation struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{read | write | action | delete}
	Name *string `json:"name,omitempty"`
	// Origin - READ-ONLY; Default value is 'user,system'.
	Origin *string `json:"origin,omitempty"`
	// Display - The information displayed about the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Display != nil {
		objectMap["display"] = o.Display
	}
	return json.Marshal(objectMap)
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - READ-ONLY; Service provider: Microsoft.HealthcareApis
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource Type: Services
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; Name of the operation
	Operation *string `json:"operation,omitempty"`
	// Description - READ-ONLY; Friendly description for the operation,
	Description *string `json:"description,omitempty"`
}

// OperationListResult a list of service operations. It contains a list of operations and a URL link to get
// the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of service description objects.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - READ-ONLY; A list of service operations supported by the Microsoft.HealthcareApis resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationListResult.
func (olr OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if olr.NextLink != nil {
		objectMap["nextLink"] = olr.NextLink
	}
	return json.Marshal(objectMap)
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{
		fn:  getNextPage,
		olr: cur,
	}
}

// OperationResultsDescription the properties indicating the operation result of an operation on a service.
type OperationResultsDescription struct {
	// ID - READ-ONLY; The ID of the operation returned.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the operation result.
	Name *string `json:"name,omitempty"`
	// Status - READ-ONLY; The status of the operation being performed. Possible values include: 'Canceled', 'Succeeded', 'Failed', 'Requested', 'Running'
	Status OperationResultStatus `json:"status,omitempty"`
	// StartTime - READ-ONLY; The time that the operation was started.
	StartTime *string `json:"startTime,omitempty"`
	// Properties - Additional properties of the operation result.
	Properties interface{} `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationResultsDescription.
func (ord OperationResultsDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ord.Properties != nil {
		objectMap["properties"] = ord.Properties
	}
	return json.Marshal(objectMap)
}

// Resource the common properties of a service.
type Resource struct {
	// ID - READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type.
	Type *string `json:"type,omitempty"`
	// Kind - The kind of the service. Possible values include: 'Fhir', 'FhirStu3', 'FhirR4'
	Kind Kind `json:"kind,omitempty"`
	// Location - The resource location.
	Location *string `json:"location,omitempty"`
	// Tags - The resource tags.
	Tags map[string]*string `json:"tags"`
	// Etag - An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `json:"etag,omitempty"`
	// Identity - Setting indicating whether the service has a managed identity associated with it.
	Identity *ResourceIdentity `json:"identity,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Kind != "" {
		objectMap["kind"] = r.Kind
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	if r.Etag != nil {
		objectMap["etag"] = r.Etag
	}
	if r.Identity != nil {
		objectMap["identity"] = r.Identity
	}
	return json.Marshal(objectMap)
}

// ResourceIdentity setting indicating whether the service has a managed identity associated with it.
type ResourceIdentity struct {
	// PrincipalID - READ-ONLY; The principal ID of the resource identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant ID of the resource.
	TenantID *string `json:"tenantId,omitempty"`
	// Type - Type of identity being specified, currently SystemAssigned and None are allowed. Possible values include: 'SystemAssigned', 'None'
	Type ManagedServiceIdentityType `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceIdentity.
func (r ResourceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Type != "" {
		objectMap["type"] = r.Type
	}
	return json.Marshal(objectMap)
}

// ServiceAccessPolicyEntry an access policy entry.
type ServiceAccessPolicyEntry struct {
	// ObjectID - An Azure AD object ID (User or Apps) that is allowed access to the FHIR service.
	ObjectID *string `json:"objectId,omitempty"`
}

// ServiceAuthenticationConfigurationInfo authentication configuration information
type ServiceAuthenticationConfigurationInfo struct {
	// Authority - The authority url for the service
	Authority *string `json:"authority,omitempty"`
	// Audience - The audience url for the service
	Audience *string `json:"audience,omitempty"`
	// SmartProxyEnabled - If the SMART on FHIR proxy is enabled
	SmartProxyEnabled *bool `json:"smartProxyEnabled,omitempty"`
}

// ServiceCorsConfigurationInfo the settings for the CORS configuration of the service instance.
type ServiceCorsConfigurationInfo struct {
	// Origins - The origins to be allowed via CORS.
	Origins *[]string `json:"origins,omitempty"`
	// Headers - The headers to be allowed via CORS.
	Headers *[]string `json:"headers,omitempty"`
	// Methods - The methods to be allowed via CORS.
	Methods *[]string `json:"methods,omitempty"`
	// MaxAge - The max age to be allowed via CORS.
	MaxAge *int32 `json:"maxAge,omitempty"`
	// AllowCredentials - If credentials are allowed via CORS.
	AllowCredentials *bool `json:"allowCredentials,omitempty"`
}

// ServiceCosmosDbConfigurationInfo the settings for the Cosmos DB database backing the service.
type ServiceCosmosDbConfigurationInfo struct {
	// OfferThroughput - The provisioned throughput for the backing database.
	OfferThroughput *int32 `json:"offerThroughput,omitempty"`
	// KeyVaultKeyURI - The URI of the customer-managed key for the backing database.
	KeyVaultKeyURI *string `json:"keyVaultKeyUri,omitempty"`
}

// ServiceExportConfigurationInfo export operation configuration information
type ServiceExportConfigurationInfo struct {
	// StorageAccountName - The name of the default export storage account.
	StorageAccountName *string `json:"storageAccountName,omitempty"`
}

// ServicesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ServicesCreateOrUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ServicesClient) (ServicesDescription, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *ServicesCreateOrUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for ServicesCreateOrUpdateFuture.Result.
func (future *ServicesCreateOrUpdateFuture) result(client ServicesClient) (sd ServicesDescription, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.ServicesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		sd.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("healthcareapis.ServicesCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if sd.Response.Response, err = future.GetResult(sender); err == nil && sd.Response.Response.StatusCode != http.StatusNoContent {
		sd, err = client.CreateOrUpdateResponder(sd.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "healthcareapis.ServicesCreateOrUpdateFuture", "Result", sd.Response.Response, "Failure responding to request")
		}
	}
	return
}

// ServicesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ServicesDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ServicesClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *ServicesDeleteFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for ServicesDeleteFuture.Result.
func (future *ServicesDeleteFuture) result(client ServicesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.ServicesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		ar.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("healthcareapis.ServicesDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// ServicesDescription the description of the service.
type ServicesDescription struct {
	autorest.Response `json:"-"`
	// Properties - The common properties of a service.
	Properties *ServicesProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type.
	Type *string `json:"type,omitempty"`
	// Kind - The kind of the service. Possible values include: 'Fhir', 'FhirStu3', 'FhirR4'
	Kind Kind `json:"kind,omitempty"`
	// Location - The resource location.
	Location *string `json:"location,omitempty"`
	// Tags - The resource tags.
	Tags map[string]*string `json:"tags"`
	// Etag - An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `json:"etag,omitempty"`
	// Identity - Setting indicating whether the service has a managed identity associated with it.
	Identity *ResourceIdentity `json:"identity,omitempty"`
}

// MarshalJSON is the custom marshaler for ServicesDescription.
func (sd ServicesDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sd.Properties != nil {
		objectMap["properties"] = sd.Properties
	}
	if sd.Kind != "" {
		objectMap["kind"] = sd.Kind
	}
	if sd.Location != nil {
		objectMap["location"] = sd.Location
	}
	if sd.Tags != nil {
		objectMap["tags"] = sd.Tags
	}
	if sd.Etag != nil {
		objectMap["etag"] = sd.Etag
	}
	if sd.Identity != nil {
		objectMap["identity"] = sd.Identity
	}
	return json.Marshal(objectMap)
}

// ServicesDescriptionListResult a list of service description objects with a next link.
type ServicesDescriptionListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of service description objects.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - A list of service description objects.
	Value *[]ServicesDescription `json:"value,omitempty"`
}

// ServicesDescriptionListResultIterator provides access to a complete listing of ServicesDescription
// values.
type ServicesDescriptionListResultIterator struct {
	i    int
	page ServicesDescriptionListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ServicesDescriptionListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesDescriptionListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ServicesDescriptionListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ServicesDescriptionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ServicesDescriptionListResultIterator) Response() ServicesDescriptionListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ServicesDescriptionListResultIterator) Value() ServicesDescription {
	if !iter.page.NotDone() {
		return ServicesDescription{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ServicesDescriptionListResultIterator type.
func NewServicesDescriptionListResultIterator(page ServicesDescriptionListResultPage) ServicesDescriptionListResultIterator {
	return ServicesDescriptionListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (sdlr ServicesDescriptionListResult) IsEmpty() bool {
	return sdlr.Value == nil || len(*sdlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (sdlr ServicesDescriptionListResult) hasNextLink() bool {
	return sdlr.NextLink != nil && len(*sdlr.NextLink) != 0
}

// servicesDescriptionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (sdlr ServicesDescriptionListResult) servicesDescriptionListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !sdlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(sdlr.NextLink)))
}

// ServicesDescriptionListResultPage contains a page of ServicesDescription values.
type ServicesDescriptionListResultPage struct {
	fn   func(context.Context, ServicesDescriptionListResult) (ServicesDescriptionListResult, error)
	sdlr ServicesDescriptionListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ServicesDescriptionListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesDescriptionListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.sdlr)
		if err != nil {
			return err
		}
		page.sdlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ServicesDescriptionListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ServicesDescriptionListResultPage) NotDone() bool {
	return !page.sdlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ServicesDescriptionListResultPage) Response() ServicesDescriptionListResult {
	return page.sdlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ServicesDescriptionListResultPage) Values() []ServicesDescription {
	if page.sdlr.IsEmpty() {
		return nil
	}
	return *page.sdlr.Value
}

// Creates a new instance of the ServicesDescriptionListResultPage type.
func NewServicesDescriptionListResultPage(cur ServicesDescriptionListResult, getNextPage func(context.Context, ServicesDescriptionListResult) (ServicesDescriptionListResult, error)) ServicesDescriptionListResultPage {
	return ServicesDescriptionListResultPage{
		fn:   getNextPage,
		sdlr: cur,
	}
}

// ServicesNameAvailabilityInfo the properties indicating whether a given service name is available.
type ServicesNameAvailabilityInfo struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; The value which indicates whether the provided name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - READ-ONLY; The reason for unavailability. Possible values include: 'Invalid', 'AlreadyExists'
	Reason ServiceNameUnavailabilityReason `json:"reason,omitempty"`
	// Message - The detailed reason message.
	Message *string `json:"message,omitempty"`
}

// MarshalJSON is the custom marshaler for ServicesNameAvailabilityInfo.
func (snai ServicesNameAvailabilityInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if snai.Message != nil {
		objectMap["message"] = snai.Message
	}
	return json.Marshal(objectMap)
}

// ServicesPatchDescription the description of the service.
type ServicesPatchDescription struct {
	// Tags - Instance tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ServicesPatchDescription.
func (spd ServicesPatchDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if spd.Tags != nil {
		objectMap["tags"] = spd.Tags
	}
	return json.Marshal(objectMap)
}

// ServicesProperties the properties of a service instance.
type ServicesProperties struct {
	// ProvisioningState - READ-ONLY; The provisioning state. Possible values include: 'ProvisioningStateDeleting', 'ProvisioningStateSucceeded', 'ProvisioningStateCreating', 'ProvisioningStateAccepted', 'ProvisioningStateVerifying', 'ProvisioningStateUpdating', 'ProvisioningStateFailed', 'ProvisioningStateCanceled', 'ProvisioningStateDeprovisioned'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// AccessPolicies - The access policies of the service instance.
	AccessPolicies *[]ServiceAccessPolicyEntry `json:"accessPolicies,omitempty"`
	// CosmosDbConfiguration - The settings for the Cosmos DB database backing the service.
	CosmosDbConfiguration *ServiceCosmosDbConfigurationInfo `json:"cosmosDbConfiguration,omitempty"`
	// AuthenticationConfiguration - The authentication configuration for the service instance.
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfo `json:"authenticationConfiguration,omitempty"`
	// CorsConfiguration - The settings for the CORS configuration of the service instance.
	CorsConfiguration *ServiceCorsConfigurationInfo `json:"corsConfiguration,omitempty"`
	// ExportConfiguration - The settings for the export operation of the service instance.
	ExportConfiguration *ServiceExportConfigurationInfo `json:"exportConfiguration,omitempty"`
}

// MarshalJSON is the custom marshaler for ServicesProperties.
func (sp ServicesProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sp.AccessPolicies != nil {
		objectMap["accessPolicies"] = sp.AccessPolicies
	}
	if sp.CosmosDbConfiguration != nil {
		objectMap["cosmosDbConfiguration"] = sp.CosmosDbConfiguration
	}
	if sp.AuthenticationConfiguration != nil {
		objectMap["authenticationConfiguration"] = sp.AuthenticationConfiguration
	}
	if sp.CorsConfiguration != nil {
		objectMap["corsConfiguration"] = sp.CorsConfiguration
	}
	if sp.ExportConfiguration != nil {
		objectMap["exportConfiguration"] = sp.ExportConfiguration
	}
	return json.Marshal(objectMap)
}

// ServicesUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ServicesUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ServicesClient) (ServicesDescription, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *ServicesUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for ServicesUpdateFuture.Result.
func (future *ServicesUpdateFuture) result(client ServicesClient) (sd ServicesDescription, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.ServicesUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		sd.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("healthcareapis.ServicesUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if sd.Response.Response, err = future.GetResult(sender); err == nil && sd.Response.Response.StatusCode != http.StatusNoContent {
		sd, err = client.UpdateResponder(sd.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "healthcareapis.ServicesUpdateFuture", "Result", sd.Response.Response, "Failure responding to request")
		}
	}
	return
}

// SetObject ...
type SetObject struct {
	autorest.Response `json:"-"`
	Value             interface{} `json:"value,omitempty"`
}
