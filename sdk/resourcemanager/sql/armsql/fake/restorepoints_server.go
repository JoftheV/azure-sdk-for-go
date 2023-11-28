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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"net/http"
	"net/url"
	"regexp"
)

// RestorePointsServer is a fake server for instances of the armsql.RestorePointsClient type.
type RestorePointsServer struct {
	// BeginCreate is the fake for method RestorePointsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters armsql.CreateDatabaseRestorePointDefinition, options *armsql.RestorePointsClientBeginCreateOptions) (resp azfake.PollerResponder[armsql.RestorePointsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method RestorePointsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, restorePointName string, options *armsql.RestorePointsClientDeleteOptions) (resp azfake.Responder[armsql.RestorePointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RestorePointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, restorePointName string, options *armsql.RestorePointsClientGetOptions) (resp azfake.Responder[armsql.RestorePointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDatabasePager is the fake for method RestorePointsClient.NewListByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabasePager func(resourceGroupName string, serverName string, databaseName string, options *armsql.RestorePointsClientListByDatabaseOptions) (resp azfake.PagerResponder[armsql.RestorePointsClientListByDatabaseResponse])
}

// NewRestorePointsServerTransport creates a new instance of RestorePointsServerTransport with the provided implementation.
// The returned RestorePointsServerTransport instance is connected to an instance of armsql.RestorePointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRestorePointsServerTransport(srv *RestorePointsServer) *RestorePointsServerTransport {
	return &RestorePointsServerTransport{
		srv:                    srv,
		beginCreate:            newTracker[azfake.PollerResponder[armsql.RestorePointsClientCreateResponse]](),
		newListByDatabasePager: newTracker[azfake.PagerResponder[armsql.RestorePointsClientListByDatabaseResponse]](),
	}
}

// RestorePointsServerTransport connects instances of armsql.RestorePointsClient to instances of RestorePointsServer.
// Don't use this type directly, use NewRestorePointsServerTransport instead.
type RestorePointsServerTransport struct {
	srv                    *RestorePointsServer
	beginCreate            *tracker[azfake.PollerResponder[armsql.RestorePointsClientCreateResponse]]
	newListByDatabasePager *tracker[azfake.PagerResponder[armsql.RestorePointsClientListByDatabaseResponse]]
}

// Do implements the policy.Transporter interface for RestorePointsServerTransport.
func (r *RestorePointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RestorePointsClient.BeginCreate":
		resp, err = r.dispatchBeginCreate(req)
	case "RestorePointsClient.Delete":
		resp, err = r.dispatchDelete(req)
	case "RestorePointsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RestorePointsClient.NewListByDatabasePager":
		resp, err = r.dispatchNewListByDatabasePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RestorePointsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := r.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.CreateDatabaseRestorePointDefinition](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreate(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		r.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		r.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		r.beginCreate.remove(req)
	}

	return resp, nil
}

func (r *RestorePointsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if r.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints/(?P<restorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	restorePointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Delete(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, restorePointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RestorePointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints/(?P<restorePointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	restorePointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, restorePointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RestorePoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RestorePointsServerTransport) dispatchNewListByDatabasePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabasePager not implemented")}
	}
	newListByDatabasePager := r.newListByDatabasePager.get(req)
	if newListByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorePoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByDatabasePager(resourceGroupNameParam, serverNameParam, databaseNameParam, nil)
		newListByDatabasePager = &resp
		r.newListByDatabasePager.add(req, newListByDatabasePager)
		server.PagerResponderInjectNextLinks(newListByDatabasePager, req, func(page *armsql.RestorePointsClientListByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabasePager) {
		r.newListByDatabasePager.remove(req)
	}
	return resp, nil
}