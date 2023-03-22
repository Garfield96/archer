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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutServiceServiceIDRejectEndpointsHandlerFunc turns a function with the right signature into a put service service ID reject endpoints handler
type PutServiceServiceIDRejectEndpointsHandlerFunc func(PutServiceServiceIDRejectEndpointsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutServiceServiceIDRejectEndpointsHandlerFunc) Handle(params PutServiceServiceIDRejectEndpointsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutServiceServiceIDRejectEndpointsHandler interface for that can handle valid put service service ID reject endpoints params
type PutServiceServiceIDRejectEndpointsHandler interface {
	Handle(PutServiceServiceIDRejectEndpointsParams, interface{}) middleware.Responder
}

// NewPutServiceServiceIDRejectEndpoints creates a new http.Handler for the put service service ID reject endpoints operation
func NewPutServiceServiceIDRejectEndpoints(ctx *middleware.Context, handler PutServiceServiceIDRejectEndpointsHandler) *PutServiceServiceIDRejectEndpoints {
	return &PutServiceServiceIDRejectEndpoints{Context: ctx, Handler: handler}
}

/*
	PutServiceServiceIDRejectEndpoints swagger:route PUT /service/{service_id}/reject_endpoints Service putServiceServiceIdRejectEndpoints

# Reject endpoints

Specify a list of consumers (`endpoint_ids` and/or `project_ids`) whose endpoints should be rejected.
* Existing active endpoints will be rejected.
* Rejected endpoints will be untouched.
* Pending endpoints will be rejected.
*/
type PutServiceServiceIDRejectEndpoints struct {
	Context *middleware.Context
	Handler PutServiceServiceIDRejectEndpointsHandler
}

func (o *PutServiceServiceIDRejectEndpoints) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutServiceServiceIDRejectEndpointsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
