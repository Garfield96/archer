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

package rbac

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sapcc/archer/models"
)

// PostRbacPoliciesCreatedCode is the HTTP code returned for type PostRbacPoliciesCreated
const PostRbacPoliciesCreatedCode int = 201

/*
PostRbacPoliciesCreated RBAC policy

swagger:response postRbacPoliciesCreated
*/
type PostRbacPoliciesCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Rbacpolicy `json:"body,omitempty"`
}

// NewPostRbacPoliciesCreated creates PostRbacPoliciesCreated with default headers values
func NewPostRbacPoliciesCreated() *PostRbacPoliciesCreated {

	return &PostRbacPoliciesCreated{}
}

// WithPayload adds the payload to the post rbac policies created response
func (o *PostRbacPoliciesCreated) WithPayload(payload *models.Rbacpolicy) *PostRbacPoliciesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post rbac policies created response
func (o *PostRbacPoliciesCreated) SetPayload(payload *models.Rbacpolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRbacPoliciesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRbacPoliciesUnauthorizedCode is the HTTP code returned for type PostRbacPoliciesUnauthorized
const PostRbacPoliciesUnauthorizedCode int = 401

/*
PostRbacPoliciesUnauthorized Unauthorized

swagger:response postRbacPoliciesUnauthorized
*/
type PostRbacPoliciesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostRbacPoliciesUnauthorized creates PostRbacPoliciesUnauthorized with default headers values
func NewPostRbacPoliciesUnauthorized() *PostRbacPoliciesUnauthorized {

	return &PostRbacPoliciesUnauthorized{}
}

// WithPayload adds the payload to the post rbac policies unauthorized response
func (o *PostRbacPoliciesUnauthorized) WithPayload(payload *models.Error) *PostRbacPoliciesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post rbac policies unauthorized response
func (o *PostRbacPoliciesUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRbacPoliciesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRbacPoliciesForbiddenCode is the HTTP code returned for type PostRbacPoliciesForbidden
const PostRbacPoliciesForbiddenCode int = 403

/*
PostRbacPoliciesForbidden Forbidden

swagger:response postRbacPoliciesForbidden
*/
type PostRbacPoliciesForbidden struct {
}

// NewPostRbacPoliciesForbidden creates PostRbacPoliciesForbidden with default headers values
func NewPostRbacPoliciesForbidden() *PostRbacPoliciesForbidden {

	return &PostRbacPoliciesForbidden{}
}

// WriteResponse to the client
func (o *PostRbacPoliciesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PostRbacPoliciesNotFoundCode is the HTTP code returned for type PostRbacPoliciesNotFound
const PostRbacPoliciesNotFoundCode int = 404

/*
PostRbacPoliciesNotFound service_id not found

swagger:response postRbacPoliciesNotFound
*/
type PostRbacPoliciesNotFound struct {
}

// NewPostRbacPoliciesNotFound creates PostRbacPoliciesNotFound with default headers values
func NewPostRbacPoliciesNotFound() *PostRbacPoliciesNotFound {

	return &PostRbacPoliciesNotFound{}
}

// WriteResponse to the client
func (o *PostRbacPoliciesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PostRbacPoliciesConflictCode is the HTTP code returned for type PostRbacPoliciesConflict
const PostRbacPoliciesConflictCode int = 409

/*
PostRbacPoliciesConflict Duplicate RBAC Policy

swagger:response postRbacPoliciesConflict
*/
type PostRbacPoliciesConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostRbacPoliciesConflict creates PostRbacPoliciesConflict with default headers values
func NewPostRbacPoliciesConflict() *PostRbacPoliciesConflict {

	return &PostRbacPoliciesConflict{}
}

// WithPayload adds the payload to the post rbac policies conflict response
func (o *PostRbacPoliciesConflict) WithPayload(payload *models.Error) *PostRbacPoliciesConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post rbac policies conflict response
func (o *PostRbacPoliciesConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRbacPoliciesConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRbacPoliciesUnprocessableEntityCode is the HTTP code returned for type PostRbacPoliciesUnprocessableEntity
const PostRbacPoliciesUnprocessableEntityCode int = 422

/*
PostRbacPoliciesUnprocessableEntity Unprocessable Content

swagger:response postRbacPoliciesUnprocessableEntity
*/
type PostRbacPoliciesUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostRbacPoliciesUnprocessableEntity creates PostRbacPoliciesUnprocessableEntity with default headers values
func NewPostRbacPoliciesUnprocessableEntity() *PostRbacPoliciesUnprocessableEntity {

	return &PostRbacPoliciesUnprocessableEntity{}
}

// WithPayload adds the payload to the post rbac policies unprocessable entity response
func (o *PostRbacPoliciesUnprocessableEntity) WithPayload(payload *models.Error) *PostRbacPoliciesUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post rbac policies unprocessable entity response
func (o *PostRbacPoliciesUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRbacPoliciesUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
