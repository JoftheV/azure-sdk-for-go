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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ProtectionPolicyOperationResultsServer is a fake server for instances of the armrecoveryservicesbackup.ProtectionPolicyOperationResultsClient type.
type ProtectionPolicyOperationResultsServer struct {
	// Get is the fake for method ProtectionPolicyOperationResultsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string, options *armrecoveryservicesbackup.ProtectionPolicyOperationResultsClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectionPolicyOperationResultsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewProtectionPolicyOperationResultsServerTransport creates a new instance of ProtectionPolicyOperationResultsServerTransport with the provided implementation.
// The returned ProtectionPolicyOperationResultsServerTransport instance is connected to an instance of armrecoveryservicesbackup.ProtectionPolicyOperationResultsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProtectionPolicyOperationResultsServerTransport(srv *ProtectionPolicyOperationResultsServer) *ProtectionPolicyOperationResultsServerTransport {
	return &ProtectionPolicyOperationResultsServerTransport{srv: srv}
}

// ProtectionPolicyOperationResultsServerTransport connects instances of armrecoveryservicesbackup.ProtectionPolicyOperationResultsClient to instances of ProtectionPolicyOperationResultsServer.
// Don't use this type directly, use NewProtectionPolicyOperationResultsServerTransport instead.
type ProtectionPolicyOperationResultsServerTransport struct {
	srv *ProtectionPolicyOperationResultsServer
}

// Do implements the policy.Transporter interface for ProtectionPolicyOperationResultsServerTransport.
func (p *ProtectionPolicyOperationResultsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProtectionPolicyOperationResultsClient.Get":
		resp, err = p.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProtectionPolicyOperationResultsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, policyNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProtectionPolicyResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}