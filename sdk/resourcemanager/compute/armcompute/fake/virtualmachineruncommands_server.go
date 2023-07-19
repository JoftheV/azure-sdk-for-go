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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualMachineRunCommandsServer is a fake server for instances of the armcompute.VirtualMachineRunCommandsClient type.
type VirtualMachineRunCommandsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualMachineRunCommandsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand armcompute.VirtualMachineRunCommand, options *armcompute.VirtualMachineRunCommandsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualMachineRunCommandsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *armcompute.VirtualMachineRunCommandsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineRunCommandsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, commandID string, options *armcompute.VirtualMachineRunCommandsClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineRunCommandsClientGetResponse], errResp azfake.ErrorResponder)

	// GetByVirtualMachine is the fake for method VirtualMachineRunCommandsClient.GetByVirtualMachine
	// HTTP status codes to indicate success: http.StatusOK
	GetByVirtualMachine func(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineOptions) (resp azfake.Responder[armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualMachineRunCommandsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armcompute.VirtualMachineRunCommandsClientListOptions) (resp azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListResponse])

	// NewListByVirtualMachinePager is the fake for method VirtualMachineRunCommandsClient.NewListByVirtualMachinePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVirtualMachinePager func(resourceGroupName string, vmName string, options *armcompute.VirtualMachineRunCommandsClientListByVirtualMachineOptions) (resp azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListByVirtualMachineResponse])

	// BeginUpdate is the fake for method VirtualMachineRunCommandsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand armcompute.VirtualMachineRunCommandUpdate, options *armcompute.VirtualMachineRunCommandsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineRunCommandsServerTransport creates a new instance of VirtualMachineRunCommandsServerTransport with the provided implementation.
// The returned VirtualMachineRunCommandsServerTransport instance is connected to an instance of armcompute.VirtualMachineRunCommandsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineRunCommandsServerTransport(srv *VirtualMachineRunCommandsServer) *VirtualMachineRunCommandsServerTransport {
	return &VirtualMachineRunCommandsServerTransport{
		srv:                          srv,
		beginCreateOrUpdate:          newTracker[azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientCreateOrUpdateResponse]](),
		beginDelete:                  newTracker[azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientDeleteResponse]](),
		newListPager:                 newTracker[azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListResponse]](),
		newListByVirtualMachinePager: newTracker[azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListByVirtualMachineResponse]](),
		beginUpdate:                  newTracker[azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientUpdateResponse]](),
	}
}

// VirtualMachineRunCommandsServerTransport connects instances of armcompute.VirtualMachineRunCommandsClient to instances of VirtualMachineRunCommandsServer.
// Don't use this type directly, use NewVirtualMachineRunCommandsServerTransport instead.
type VirtualMachineRunCommandsServerTransport struct {
	srv                          *VirtualMachineRunCommandsServer
	beginCreateOrUpdate          *tracker[azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientCreateOrUpdateResponse]]
	beginDelete                  *tracker[azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientDeleteResponse]]
	newListPager                 *tracker[azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListResponse]]
	newListByVirtualMachinePager *tracker[azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListByVirtualMachineResponse]]
	beginUpdate                  *tracker[azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for VirtualMachineRunCommandsServerTransport.
func (v *VirtualMachineRunCommandsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineRunCommandsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualMachineRunCommandsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualMachineRunCommandsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineRunCommandsClient.GetByVirtualMachine":
		resp, err = v.dispatchGetByVirtualMachine(req)
	case "VirtualMachineRunCommandsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualMachineRunCommandsClient.NewListByVirtualMachinePager":
		resp, err = v.dispatchNewListByVirtualMachinePager(req)
	case "VirtualMachineRunCommandsClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands/(?P<runCommandName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineRunCommand](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmName")])
		if err != nil {
			return nil, err
		}
		runCommandNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("runCommandName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, vmNameUnescaped, runCommandNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands/(?P<runCommandName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmName")])
		if err != nil {
			return nil, err
		}
		runCommandNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("runCommandName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, vmNameUnescaped, runCommandNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands/(?P<commandId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	commandIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("commandId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), locationUnescaped, commandIDUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RunCommandDocument, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchGetByVirtualMachine(req *http.Request) (*http.Response, error) {
	if v.srv.GetByVirtualMachine == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByVirtualMachine not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands/(?P<runCommandName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmName")])
	if err != nil {
		return nil, err
	}
	runCommandNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("runCommandName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineOptions
	if expandParam != nil {
		options = &armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.GetByVirtualMachine(req.Context(), resourceGroupNameUnescaped, vmNameUnescaped, runCommandNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineRunCommand, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(locationUnescaped, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcompute.VirtualMachineRunCommandsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchNewListByVirtualMachinePager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByVirtualMachinePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVirtualMachinePager not implemented")}
	}
	newListByVirtualMachinePager := v.newListByVirtualMachinePager.get(req)
	if newListByVirtualMachinePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmName")])
		if err != nil {
			return nil, err
		}
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armcompute.VirtualMachineRunCommandsClientListByVirtualMachineOptions
		if expandParam != nil {
			options = &armcompute.VirtualMachineRunCommandsClientListByVirtualMachineOptions{
				Expand: expandParam,
			}
		}
		resp := v.srv.NewListByVirtualMachinePager(resourceGroupNameUnescaped, vmNameUnescaped, options)
		newListByVirtualMachinePager = &resp
		v.newListByVirtualMachinePager.add(req, newListByVirtualMachinePager)
		server.PagerResponderInjectNextLinks(newListByVirtualMachinePager, req, func(page *armcompute.VirtualMachineRunCommandsClientListByVirtualMachineResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVirtualMachinePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByVirtualMachinePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVirtualMachinePager) {
		v.newListByVirtualMachinePager.remove(req)
	}
	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands/(?P<runCommandName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineRunCommandUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmName")])
		if err != nil {
			return nil, err
		}
		runCommandNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("runCommandName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameUnescaped, vmNameUnescaped, runCommandNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}
