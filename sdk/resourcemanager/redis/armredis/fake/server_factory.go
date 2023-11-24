//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armredis.ClientFactory type.
type ServerFactory struct {
	AccessPolicyAssignmentServer     AccessPolicyAssignmentServer
	AccessPolicyServer               AccessPolicyServer
	AsyncOperationStatusServer       AsyncOperationStatusServer
	Server                           Server
	FirewallRulesServer              FirewallRulesServer
	LinkedServerServer               LinkedServerServer
	OperationsServer                 OperationsServer
	PatchSchedulesServer             PatchSchedulesServer
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer       PrivateLinkResourcesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armredis.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armredis.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trAccessPolicyAssignmentServer     *AccessPolicyAssignmentServerTransport
	trAccessPolicyServer               *AccessPolicyServerTransport
	trAsyncOperationStatusServer       *AsyncOperationStatusServerTransport
	trServer                           *ServerTransport
	trFirewallRulesServer              *FirewallRulesServerTransport
	trLinkedServerServer               *LinkedServerServerTransport
	trOperationsServer                 *OperationsServerTransport
	trPatchSchedulesServer             *PatchSchedulesServerTransport
	trPrivateEndpointConnectionsServer *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer       *PrivateLinkResourcesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AccessPolicyAssignmentClient":
		initServer(s, &s.trAccessPolicyAssignmentServer, func() *AccessPolicyAssignmentServerTransport {
			return NewAccessPolicyAssignmentServerTransport(&s.srv.AccessPolicyAssignmentServer)
		})
		resp, err = s.trAccessPolicyAssignmentServer.Do(req)
	case "AccessPolicyClient":
		initServer(s, &s.trAccessPolicyServer, func() *AccessPolicyServerTransport { return NewAccessPolicyServerTransport(&s.srv.AccessPolicyServer) })
		resp, err = s.trAccessPolicyServer.Do(req)
	case "AsyncOperationStatusClient":
		initServer(s, &s.trAsyncOperationStatusServer, func() *AsyncOperationStatusServerTransport {
			return NewAsyncOperationStatusServerTransport(&s.srv.AsyncOperationStatusServer)
		})
		resp, err = s.trAsyncOperationStatusServer.Do(req)
	case "Client":
		initServer(s, &s.trServer, func() *ServerTransport { return NewServerTransport(&s.srv.Server) })
		resp, err = s.trServer.Do(req)
	case "FirewallRulesClient":
		initServer(s, &s.trFirewallRulesServer, func() *FirewallRulesServerTransport {
			return NewFirewallRulesServerTransport(&s.srv.FirewallRulesServer)
		})
		resp, err = s.trFirewallRulesServer.Do(req)
	case "LinkedServerClient":
		initServer(s, &s.trLinkedServerServer, func() *LinkedServerServerTransport { return NewLinkedServerServerTransport(&s.srv.LinkedServerServer) })
		resp, err = s.trLinkedServerServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PatchSchedulesClient":
		initServer(s, &s.trPatchSchedulesServer, func() *PatchSchedulesServerTransport {
			return NewPatchSchedulesServerTransport(&s.srv.PatchSchedulesServer)
		})
		resp, err = s.trPatchSchedulesServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}