//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
	"net/http"
	"net/url"
	"regexp"
)

// AssetFiltersServer is a fake server for instances of the armmediaservices.AssetFiltersClient type.
type AssetFiltersServer struct {
	// CreateOrUpdate is the fake for method AssetFiltersClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, parameters armmediaservices.AssetFilter, options *armmediaservices.AssetFiltersClientCreateOrUpdateOptions) (resp azfake.Responder[armmediaservices.AssetFiltersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method AssetFiltersClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, options *armmediaservices.AssetFiltersClientDeleteOptions) (resp azfake.Responder[armmediaservices.AssetFiltersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AssetFiltersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, options *armmediaservices.AssetFiltersClientGetOptions) (resp azfake.Responder[armmediaservices.AssetFiltersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AssetFiltersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, assetName string, options *armmediaservices.AssetFiltersClientListOptions) (resp azfake.PagerResponder[armmediaservices.AssetFiltersClientListResponse])

	// Update is the fake for method AssetFiltersClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, parameters armmediaservices.AssetFilter, options *armmediaservices.AssetFiltersClientUpdateOptions) (resp azfake.Responder[armmediaservices.AssetFiltersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAssetFiltersServerTransport creates a new instance of AssetFiltersServerTransport with the provided implementation.
// The returned AssetFiltersServerTransport instance is connected to an instance of armmediaservices.AssetFiltersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAssetFiltersServerTransport(srv *AssetFiltersServer) *AssetFiltersServerTransport {
	return &AssetFiltersServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmediaservices.AssetFiltersClientListResponse]](),
	}
}

// AssetFiltersServerTransport connects instances of armmediaservices.AssetFiltersClient to instances of AssetFiltersServer.
// Don't use this type directly, use NewAssetFiltersServerTransport instead.
type AssetFiltersServerTransport struct {
	srv          *AssetFiltersServer
	newListPager *tracker[azfake.PagerResponder[armmediaservices.AssetFiltersClientListResponse]]
}

// Do implements the policy.Transporter interface for AssetFiltersServerTransport.
func (a *AssetFiltersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AssetFiltersClient.CreateOrUpdate":
		resp, err = a.dispatchCreateOrUpdate(req)
	case "AssetFiltersClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "AssetFiltersClient.Get":
		resp, err = a.dispatchGet(req)
	case "AssetFiltersClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AssetFiltersClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AssetFiltersServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assets/(?P<assetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assetFilters/(?P<filterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmediaservices.AssetFilter](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	assetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assetName")])
	if err != nil {
		return nil, err
	}
	filterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("filterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, accountNameParam, assetNameParam, filterNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssetFilter, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssetFiltersServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assets/(?P<assetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assetFilters/(?P<filterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	assetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assetName")])
	if err != nil {
		return nil, err
	}
	filterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("filterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, accountNameParam, assetNameParam, filterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssetFiltersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assets/(?P<assetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assetFilters/(?P<filterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	assetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assetName")])
	if err != nil {
		return nil, err
	}
	filterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("filterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, assetNameParam, filterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssetFilter, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssetFiltersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assets/(?P<assetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assetFilters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		assetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assetName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, accountNameParam, assetNameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmediaservices.AssetFiltersClientListResponse, createLink func() string) {
			page.ODataNextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

func (a *AssetFiltersServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assets/(?P<assetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assetFilters/(?P<filterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmediaservices.AssetFilter](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	assetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assetName")])
	if err != nil {
		return nil, err
	}
	filterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("filterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameParam, accountNameParam, assetNameParam, filterNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssetFilter, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}