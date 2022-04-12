//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the Quota group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host string
	pl   runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the quota limit for the specified resource with the requested value. To update the
// quota, follow these steps:
// 1. Use the GET operation for quotas and usages to determine how much quota remains for the specific resource and to calculate
// the new quota limit. These steps are detailed in this example
// [https://techcommunity.microsoft.com/t5/azure-governance-and-management/using-the-new-quota-rest-api/ba-p/2183670].
// 2. Use this PUT operation to update the quota limit. Please check the URI in location header for the detailed status of
// the request.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceName - Resource name for a given resource provider. For example:
// * SKU name for Microsoft.Compute
// * SKU or TotalLowPriorityCores for Microsoft.MachineLearningServices For Microsoft.Network PublicIPAddresses.
// scope - The target Azure resource URI. For example, /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/qms-test/providers/Microsoft.Batch/batchAccounts/testAccount/.
// This is the target Azure
// resource URI for the List GET operation. If a {resourceName} is added after /quotas, then it's the target Azure resource
// URI in the GET operation for the specific resource.
// createQuotaRequest - Quota request payload.
// options - ClientBeginCreateOrUpdateOptions contains the optional parameters for the Client.BeginCreateOrUpdate method.
func (client *Client) BeginCreateOrUpdate(ctx context.Context, resourceName string, scope string, createQuotaRequest CurrentQuotaLimitBase, options *ClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceName, scope, createQuotaRequest, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaOriginalURI,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update the quota limit for the specified resource with the requested value. To update the quota,
// follow these steps:
// 1. Use the GET operation for quotas and usages to determine how much quota remains for the specific resource and to calculate
// the new quota limit. These steps are detailed in this example
// [https://techcommunity.microsoft.com/t5/azure-governance-and-management/using-the-new-quota-rest-api/ba-p/2183670].
// 2. Use this PUT operation to update the quota limit. Please check the URI in location header for the detailed status of
// the request.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) createOrUpdate(ctx context.Context, resourceName string, scope string, createQuotaRequest CurrentQuotaLimitBase, options *ClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceName, scope, createQuotaRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *Client) createOrUpdateCreateRequest(ctx context.Context, resourceName string, scope string, createQuotaRequest CurrentQuotaLimitBase, options *ClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Quota/quotas/{resourceName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, createQuotaRequest)
}

// Get - Get the quota limit of a resource. The response can be used to determine the remaining quota to calculate a new quota
// limit that can be submitted with a PUT request.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceName - Resource name for a given resource provider. For example:
// * SKU name for Microsoft.Compute
// * SKU or TotalLowPriorityCores for Microsoft.MachineLearningServices For Microsoft.Network PublicIPAddresses.
// scope - The target Azure resource URI. For example, /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/qms-test/providers/Microsoft.Batch/batchAccounts/testAccount/.
// This is the target Azure
// resource URI for the List GET operation. If a {resourceName} is added after /quotas, then it's the target Azure resource
// URI in the GET operation for the specific resource.
// options - ClientGetOptions contains the optional parameters for the Client.Get method.
func (client *Client) Get(ctx context.Context, resourceName string, scope string, options *ClientGetOptions) (ClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceName, scope, options)
	if err != nil {
		return ClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *Client) getCreateRequest(ctx context.Context, resourceName string, scope string, options *ClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Quota/quotas/{resourceName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Client) getHandleResponse(resp *http.Response) (ClientGetResponse, error) {
	result := ClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CurrentQuotaLimitBase); err != nil {
		return ClientGetResponse{}, err
	}
	return result, nil
}

// List - Get a list of current quota limits of all resources for the specified scope. The response from this GET operation
// can be leveraged to submit requests to update a quota.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The target Azure resource URI. For example, /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/qms-test/providers/Microsoft.Batch/batchAccounts/testAccount/.
// This is the target Azure
// resource URI for the List GET operation. If a {resourceName} is added after /quotas, then it's the target Azure resource
// URI in the GET operation for the specific resource.
// options - ClientListOptions contains the optional parameters for the Client.List method.
func (client *Client) List(scope string, options *ClientListOptions) *runtime.Pager[ClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ClientListResponse]{
		More: func(page ClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListResponse) (ClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *Client) listCreateRequest(ctx context.Context, scope string, options *ClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Quota/quotas"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *Client) listHandleResponse(resp *http.Response) (ClientListResponse, error) {
	result := ClientListResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Limits); err != nil {
		return ClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update the quota limit for a specific resource to the specified value:
// 1. Use the Usages-GET and Quota-GET operations to determine the remaining quota for the specific resource and to calculate
// the new quota limit. These steps are detailed in this example
// [https://techcommunity.microsoft.com/t5/azure-governance-and-management/using-the-new-quota-rest-api/ba-p/2183670].
// 2. Use this PUT operation to update the quota limit. Please check the URI in location header for the detailed status of
// the request.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceName - Resource name for a given resource provider. For example:
// * SKU name for Microsoft.Compute
// * SKU or TotalLowPriorityCores for Microsoft.MachineLearningServices For Microsoft.Network PublicIPAddresses.
// scope - The target Azure resource URI. For example, /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/qms-test/providers/Microsoft.Batch/batchAccounts/testAccount/.
// This is the target Azure
// resource URI for the List GET operation. If a {resourceName} is added after /quotas, then it's the target Azure resource
// URI in the GET operation for the specific resource.
// createQuotaRequest - Quota requests payload.
// options - ClientBeginUpdateOptions contains the optional parameters for the Client.BeginUpdate method.
func (client *Client) BeginUpdate(ctx context.Context, resourceName string, scope string, createQuotaRequest CurrentQuotaLimitBase, options *ClientBeginUpdateOptions) (*armruntime.Poller[ClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceName, scope, createQuotaRequest, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaOriginalURI,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update the quota limit for a specific resource to the specified value:
// 1. Use the Usages-GET and Quota-GET operations to determine the remaining quota for the specific resource and to calculate
// the new quota limit. These steps are detailed in this example
// [https://techcommunity.microsoft.com/t5/azure-governance-and-management/using-the-new-quota-rest-api/ba-p/2183670].
// 2. Use this PUT operation to update the quota limit. Please check the URI in location header for the detailed status of
// the request.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) update(ctx context.Context, resourceName string, scope string, createQuotaRequest CurrentQuotaLimitBase, options *ClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceName, scope, createQuotaRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *Client) updateCreateRequest(ctx context.Context, resourceName string, scope string, createQuotaRequest CurrentQuotaLimitBase, options *ClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Quota/quotas/{resourceName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, createQuotaRequest)
}
