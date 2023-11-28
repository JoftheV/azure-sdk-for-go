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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkServiceDesignVersionsServer is a fake server for instances of the armhybridnetwork.NetworkServiceDesignVersionsClient type.
type NetworkServiceDesignVersionsServer struct {
	// BeginCreateOrUpdate is the fake for method NetworkServiceDesignVersionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters armhybridnetwork.NetworkServiceDesignVersion, options *armhybridnetwork.NetworkServiceDesignVersionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NetworkServiceDesignVersionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, options *armhybridnetwork.NetworkServiceDesignVersionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NetworkServiceDesignVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, options *armhybridnetwork.NetworkServiceDesignVersionsClientGetOptions) (resp azfake.Responder[armhybridnetwork.NetworkServiceDesignVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByNetworkServiceDesignGroupPager is the fake for method NetworkServiceDesignVersionsClient.NewListByNetworkServiceDesignGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByNetworkServiceDesignGroupPager func(resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *armhybridnetwork.NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupOptions) (resp azfake.PagerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse])

	// Update is the fake for method NetworkServiceDesignVersionsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters armhybridnetwork.TagsObject, options *armhybridnetwork.NetworkServiceDesignVersionsClientUpdateOptions) (resp azfake.Responder[armhybridnetwork.NetworkServiceDesignVersionsClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateState is the fake for method NetworkServiceDesignVersionsClient.BeginUpdateState
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateState func(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters armhybridnetwork.NetworkServiceDesignVersionUpdateState, options *armhybridnetwork.NetworkServiceDesignVersionsClientBeginUpdateStateOptions) (resp azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientUpdateStateResponse], errResp azfake.ErrorResponder)
}

// NewNetworkServiceDesignVersionsServerTransport creates a new instance of NetworkServiceDesignVersionsServerTransport with the provided implementation.
// The returned NetworkServiceDesignVersionsServerTransport instance is connected to an instance of armhybridnetwork.NetworkServiceDesignVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkServiceDesignVersionsServerTransport(srv *NetworkServiceDesignVersionsServer) *NetworkServiceDesignVersionsServerTransport {
	return &NetworkServiceDesignVersionsServerTransport{
		srv:                                     srv,
		beginCreateOrUpdate:                     newTracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientCreateOrUpdateResponse]](),
		beginDelete:                             newTracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientDeleteResponse]](),
		newListByNetworkServiceDesignGroupPager: newTracker[azfake.PagerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse]](),
		beginUpdateState:                        newTracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientUpdateStateResponse]](),
	}
}

// NetworkServiceDesignVersionsServerTransport connects instances of armhybridnetwork.NetworkServiceDesignVersionsClient to instances of NetworkServiceDesignVersionsServer.
// Don't use this type directly, use NewNetworkServiceDesignVersionsServerTransport instead.
type NetworkServiceDesignVersionsServerTransport struct {
	srv                                     *NetworkServiceDesignVersionsServer
	beginCreateOrUpdate                     *tracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientCreateOrUpdateResponse]]
	beginDelete                             *tracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientDeleteResponse]]
	newListByNetworkServiceDesignGroupPager *tracker[azfake.PagerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse]]
	beginUpdateState                        *tracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignVersionsClientUpdateStateResponse]]
}

// Do implements the policy.Transporter interface for NetworkServiceDesignVersionsServerTransport.
func (n *NetworkServiceDesignVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkServiceDesignVersionsClient.BeginCreateOrUpdate":
		resp, err = n.dispatchBeginCreateOrUpdate(req)
	case "NetworkServiceDesignVersionsClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NetworkServiceDesignVersionsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NetworkServiceDesignVersionsClient.NewListByNetworkServiceDesignGroupPager":
		resp, err = n.dispatchNewListByNetworkServiceDesignGroupPager(req)
	case "NetworkServiceDesignVersionsClient.Update":
		resp, err = n.dispatchUpdate(req)
	case "NetworkServiceDesignVersionsClient.BeginUpdateState":
		resp, err = n.dispatchBeginUpdateState(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkServiceDesignVersionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignVersions/(?P<networkServiceDesignVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.NetworkServiceDesignVersion](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
		if err != nil {
			return nil, err
		}
		networkServiceDesignVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, networkServiceDesignVersionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NetworkServiceDesignVersionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignVersions/(?P<networkServiceDesignVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
		if err != nil {
			return nil, err
		}
		networkServiceDesignVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, networkServiceDesignVersionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		n.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NetworkServiceDesignVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignVersions/(?P<networkServiceDesignVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
	if err != nil {
		return nil, err
	}
	networkServiceDesignVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignVersionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, networkServiceDesignVersionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkServiceDesignVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkServiceDesignVersionsServerTransport) dispatchNewListByNetworkServiceDesignGroupPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByNetworkServiceDesignGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByNetworkServiceDesignGroupPager not implemented")}
	}
	newListByNetworkServiceDesignGroupPager := n.newListByNetworkServiceDesignGroupPager.get(req)
	if newListByNetworkServiceDesignGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignVersions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByNetworkServiceDesignGroupPager(resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, nil)
		newListByNetworkServiceDesignGroupPager = &resp
		n.newListByNetworkServiceDesignGroupPager.add(req, newListByNetworkServiceDesignGroupPager)
		server.PagerResponderInjectNextLinks(newListByNetworkServiceDesignGroupPager, req, func(page *armhybridnetwork.NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByNetworkServiceDesignGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByNetworkServiceDesignGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByNetworkServiceDesignGroupPager) {
		n.newListByNetworkServiceDesignGroupPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkServiceDesignVersionsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignVersions/(?P<networkServiceDesignVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
	if err != nil {
		return nil, err
	}
	networkServiceDesignVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignVersionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Update(req.Context(), resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, networkServiceDesignVersionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkServiceDesignVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkServiceDesignVersionsServerTransport) dispatchBeginUpdateState(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpdateState == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateState not implemented")}
	}
	beginUpdateState := n.beginUpdateState.get(req)
	if beginUpdateState == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignVersions/(?P<networkServiceDesignVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateState`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.NetworkServiceDesignVersionUpdateState](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
		if err != nil {
			return nil, err
		}
		networkServiceDesignVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginUpdateState(req.Context(), resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, networkServiceDesignVersionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateState = &respr
		n.beginUpdateState.add(req, beginUpdateState)
	}

	resp, err := server.PollerResponderNext(beginUpdateState, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginUpdateState.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateState) {
		n.beginUpdateState.remove(req)
	}

	return resp, nil
}