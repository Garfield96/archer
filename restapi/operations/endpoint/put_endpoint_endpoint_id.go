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

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PutEndpointEndpointIDHandlerFunc turns a function with the right signature into a put endpoint endpoint ID handler
type PutEndpointEndpointIDHandlerFunc func(PutEndpointEndpointIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutEndpointEndpointIDHandlerFunc) Handle(params PutEndpointEndpointIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutEndpointEndpointIDHandler interface for that can handle valid put endpoint endpoint ID params
type PutEndpointEndpointIDHandler interface {
	Handle(PutEndpointEndpointIDParams, interface{}) middleware.Responder
}

// NewPutEndpointEndpointID creates a new http.Handler for the put endpoint endpoint ID operation
func NewPutEndpointEndpointID(ctx *middleware.Context, handler PutEndpointEndpointIDHandler) *PutEndpointEndpointID {
	return &PutEndpointEndpointID{Context: ctx, Handler: handler}
}

/*
	PutEndpointEndpointID swagger:route PUT /endpoint/{endpoint_id} Endpoint putEndpointEndpointId

Update an existing endpoint
*/
type PutEndpointEndpointID struct {
	Context *middleware.Context
	Handler PutEndpointEndpointIDHandler
}

func (o *PutEndpointEndpointID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutEndpointEndpointIDParams()
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

// PutEndpointEndpointIDBody put endpoint endpoint ID body
//
// swagger:model PutEndpointEndpointIDBody
type PutEndpointEndpointIDBody struct {

	// Description of the endpoint.
	// Example: An example of an endpoint.
	// Max Length: 255
	Description *string `json:"description,omitempty"`

	// Name of the endpoint.
	// Example: Example endpoint.
	// Max Length: 64
	Name *string `json:"name,omitempty"`

	// The list of tags on the resource.
	Tags []string `json:"tags"`
}

// Validate validates this put endpoint endpoint ID body
func (o *PutEndpointEndpointIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutEndpointEndpointIDBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("body"+"."+"description", "body", *o.Description, 255); err != nil {
		return err
	}

	return nil
}

func (o *PutEndpointEndpointIDBody) validateName(formats strfmt.Registry) error {
	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("body"+"."+"name", "body", *o.Name, 64); err != nil {
		return err
	}

	return nil
}

func (o *PutEndpointEndpointIDBody) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(o.Tags) { // not required
		return nil
	}

	for i := 0; i < len(o.Tags); i++ {

		if err := validate.MaxLength("body"+"."+"tags"+"."+strconv.Itoa(i), "body", o.Tags[i], 64); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this put endpoint endpoint ID body based on context it is used
func (o *PutEndpointEndpointIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutEndpointEndpointIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutEndpointEndpointIDBody) UnmarshalBinary(b []byte) error {
	var res PutEndpointEndpointIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
