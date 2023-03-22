// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 SAP SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteServiceServiceID(params *DeleteServiceServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServiceServiceIDAccepted, error)

	GetService(params *GetServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceOK, error)

	GetServiceServiceID(params *GetServiceServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceServiceIDOK, error)

	GetServiceServiceIDEndpoints(params *GetServiceServiceIDEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceServiceIDEndpointsOK, error)

	PostService(params *PostServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostServiceCreated, error)

	PutServiceServiceID(params *PutServiceServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutServiceServiceIDOK, error)

	PutServiceServiceIDAcceptEndpoints(params *PutServiceServiceIDAcceptEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutServiceServiceIDAcceptEndpointsOK, error)

	PutServiceServiceIDRejectEndpoints(params *PutServiceServiceIDRejectEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutServiceServiceIDRejectEndpointsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	DeleteServiceServiceID removes service from catalog

	Deletes this service. There **must** be no active associated endpoint for successfully deleting the service.

Active endpoints can be rejected by the service owner via the `/service/{service_id}/reject_endpoints` API.
*/
func (a *Client) DeleteServiceServiceID(params *DeleteServiceServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServiceServiceIDAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceServiceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteServiceServiceID",
		Method:             "DELETE",
		PathPattern:        "/service/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteServiceServiceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteServiceServiceIDAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteServiceServiceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetService lists services
*/
func (a *Client) GetService(params *GetServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetService",
		Method:             "GET",
		PathPattern:        "/service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetServiceServiceID shows details of an service
*/
func (a *Client) GetServiceServiceID(params *GetServiceServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceServiceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceServiceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetServiceServiceID",
		Method:             "GET",
		PathPattern:        "/service/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceServiceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceServiceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetServiceServiceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetServiceServiceIDEndpoints lists service endpoints consumers

	Provides a list of service consumers (endpoints).

This list can be used to accept or reject requests, or disable active endpoints.
Rejected endpoints will be cleaned up after a specific time.
*/
func (a *Client) GetServiceServiceIDEndpoints(params *GetServiceServiceIDEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceServiceIDEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceServiceIDEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetServiceServiceIDEndpoints",
		Method:             "GET",
		PathPattern:        "/service/{service_id}/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceServiceIDEndpointsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceServiceIDEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetServiceServiceIDEndpoints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostService adds a new service to the catalog
*/
func (a *Client) PostService(params *PostServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostServiceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostService",
		Method:             "POST",
		PathPattern:        "/service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostServiceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutServiceServiceID updates an existing service
*/
func (a *Client) PutServiceServiceID(params *PutServiceServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutServiceServiceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutServiceServiceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutServiceServiceID",
		Method:             "PUT",
		PathPattern:        "/service/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutServiceServiceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutServiceServiceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutServiceServiceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PutServiceServiceIDAcceptEndpoints accepts endpoints

	Specify a list of endpoint consumers (`endpoint_ids` and/or `project_ids`) whose endpoints should be accepted.

* Existing active endpoints will be untouched.
* Rejected endpoints will be accepted.
* Pending endpoints will be accepted.
*/
func (a *Client) PutServiceServiceIDAcceptEndpoints(params *PutServiceServiceIDAcceptEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutServiceServiceIDAcceptEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutServiceServiceIDAcceptEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutServiceServiceIDAcceptEndpoints",
		Method:             "PUT",
		PathPattern:        "/service/{service_id}/accept_endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutServiceServiceIDAcceptEndpointsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutServiceServiceIDAcceptEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutServiceServiceIDAcceptEndpoints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PutServiceServiceIDRejectEndpoints rejects endpoints

	Specify a list of consumers (`endpoint_ids` and/or `project_ids`) whose endpoints should be rejected.

* Existing active endpoints will be rejected.
* Rejected endpoints will be untouched.
* Pending endpoints will be rejected.
*/
func (a *Client) PutServiceServiceIDRejectEndpoints(params *PutServiceServiceIDRejectEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutServiceServiceIDRejectEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutServiceServiceIDRejectEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutServiceServiceIDRejectEndpoints",
		Method:             "PUT",
		PathPattern:        "/service/{service_id}/reject_endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutServiceServiceIDRejectEndpointsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutServiceServiceIDRejectEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutServiceServiceIDRejectEndpoints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
