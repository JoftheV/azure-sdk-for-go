//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// JobExecutionsClient contains the methods for the JobExecutions group.
// Don't use this type directly, use NewJobExecutionsClient() instead.
type JobExecutionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewJobExecutionsClient creates a new instance of JobExecutionsClient with the specified values.
func NewJobExecutionsClient(con *arm.Connection, subscriptionID string) *JobExecutionsClient {
	return &JobExecutionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Cancel - Requests cancellation of a job execution.
// If the operation fails it returns a generic error.
func (client *JobExecutionsClient) Cancel(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsCancelOptions) (JobExecutionsCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
	if err != nil {
		return JobExecutionsCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobExecutionsCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobExecutionsCancelResponse{}, client.cancelHandleError(resp)
	}
	return JobExecutionsCancelResponse{RawResponse: resp}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *JobExecutionsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobExecutionId}", url.PathEscape(jobExecutionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *JobExecutionsClient) cancelHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginCreate - Starts an elastic job execution.
// If the operation fails it returns a generic error.
func (client *JobExecutionsClient) BeginCreate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsBeginCreateOptions) (JobExecutionsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
	if err != nil {
		return JobExecutionsCreatePollerResponse{}, err
	}
	result := JobExecutionsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobExecutionsClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return JobExecutionsCreatePollerResponse{}, err
	}
	result.Poller = &JobExecutionsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Starts an elastic job execution.
// If the operation fails it returns a generic error.
func (client *JobExecutionsClient) create(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *JobExecutionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/start"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createHandleError handles the Create error response.
func (client *JobExecutionsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginCreateOrUpdate - Creates or updates a job execution.
// If the operation fails it returns a generic error.
func (client *JobExecutionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsBeginCreateOrUpdateOptions) (JobExecutionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
	if err != nil {
		return JobExecutionsCreateOrUpdatePollerResponse{}, err
	}
	result := JobExecutionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobExecutionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return JobExecutionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &JobExecutionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a job execution.
// If the operation fails it returns a generic error.
func (client *JobExecutionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobExecutionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobExecutionId}", url.PathEscape(jobExecutionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *JobExecutionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a job execution.
// If the operation fails it returns a generic error.
func (client *JobExecutionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsGetOptions) (JobExecutionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
	if err != nil {
		return JobExecutionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobExecutionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobExecutionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobExecutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobExecutionId}", url.PathEscape(jobExecutionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobExecutionsClient) getHandleResponse(resp *http.Response) (JobExecutionsGetResponse, error) {
	result := JobExecutionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecution); err != nil {
		return JobExecutionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *JobExecutionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByAgent - Lists all executions in a job agent.
// If the operation fails it returns a generic error.
func (client *JobExecutionsClient) ListByAgent(resourceGroupName string, serverName string, jobAgentName string, options *JobExecutionsListByAgentOptions) *JobExecutionsListByAgentPager {
	return &JobExecutionsListByAgentPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAgentCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, options)
		},
		advancer: func(ctx context.Context, resp JobExecutionsListByAgentResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobExecutionListResult.NextLink)
		},
	}
}

// listByAgentCreateRequest creates the ListByAgent request.
func (client *JobExecutionsClient) listByAgentCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, options *JobExecutionsListByAgentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/executions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CreateTimeMin != nil {
		reqQP.Set("createTimeMin", options.CreateTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.CreateTimeMax != nil {
		reqQP.Set("createTimeMax", options.CreateTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMin != nil {
		reqQP.Set("endTimeMin", options.EndTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMax != nil {
		reqQP.Set("endTimeMax", options.EndTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.IsActive != nil {
		reqQP.Set("isActive", strconv.FormatBool(*options.IsActive))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAgentHandleResponse handles the ListByAgent response.
func (client *JobExecutionsClient) listByAgentHandleResponse(resp *http.Response) (JobExecutionsListByAgentResponse, error) {
	result := JobExecutionsListByAgentResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecutionListResult); err != nil {
		return JobExecutionsListByAgentResponse{}, err
	}
	return result, nil
}

// listByAgentHandleError handles the ListByAgent error response.
func (client *JobExecutionsClient) listByAgentHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByJob - Lists a job's executions.
// If the operation fails it returns a generic error.
func (client *JobExecutionsClient) ListByJob(resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsListByJobOptions) *JobExecutionsListByJobPager {
	return &JobExecutionsListByJobPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByJobCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
		},
		advancer: func(ctx context.Context, resp JobExecutionsListByJobResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobExecutionListResult.NextLink)
		},
	}
}

// listByJobCreateRequest creates the ListByJob request.
func (client *JobExecutionsClient) listByJobCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsListByJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CreateTimeMin != nil {
		reqQP.Set("createTimeMin", options.CreateTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.CreateTimeMax != nil {
		reqQP.Set("createTimeMax", options.CreateTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMin != nil {
		reqQP.Set("endTimeMin", options.EndTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMax != nil {
		reqQP.Set("endTimeMax", options.EndTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.IsActive != nil {
		reqQP.Set("isActive", strconv.FormatBool(*options.IsActive))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByJobHandleResponse handles the ListByJob response.
func (client *JobExecutionsClient) listByJobHandleResponse(resp *http.Response) (JobExecutionsListByJobResponse, error) {
	result := JobExecutionsListByJobResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecutionListResult); err != nil {
		return JobExecutionsListByJobResponse{}, err
	}
	return result, nil
}

// listByJobHandleError handles the ListByJob error response.
func (client *JobExecutionsClient) listByJobHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
