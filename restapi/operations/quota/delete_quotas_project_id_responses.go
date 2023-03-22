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

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sapcc/archer/models"
)

// DeleteQuotasProjectIDNoContentCode is the HTTP code returned for type DeleteQuotasProjectIDNoContent
const DeleteQuotasProjectIDNoContentCode int = 204

/*
DeleteQuotasProjectIDNoContent Resource successfully reset

swagger:response deleteQuotasProjectIdNoContent
*/
type DeleteQuotasProjectIDNoContent struct {
}

// NewDeleteQuotasProjectIDNoContent creates DeleteQuotasProjectIDNoContent with default headers values
func NewDeleteQuotasProjectIDNoContent() *DeleteQuotasProjectIDNoContent {

	return &DeleteQuotasProjectIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteQuotasProjectIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteQuotasProjectIDUnauthorizedCode is the HTTP code returned for type DeleteQuotasProjectIDUnauthorized
const DeleteQuotasProjectIDUnauthorizedCode int = 401

/*
DeleteQuotasProjectIDUnauthorized Unauthorized

swagger:response deleteQuotasProjectIdUnauthorized
*/
type DeleteQuotasProjectIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteQuotasProjectIDUnauthorized creates DeleteQuotasProjectIDUnauthorized with default headers values
func NewDeleteQuotasProjectIDUnauthorized() *DeleteQuotasProjectIDUnauthorized {

	return &DeleteQuotasProjectIDUnauthorized{}
}

// WithPayload adds the payload to the delete quotas project Id unauthorized response
func (o *DeleteQuotasProjectIDUnauthorized) WithPayload(payload *models.Error) *DeleteQuotasProjectIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete quotas project Id unauthorized response
func (o *DeleteQuotasProjectIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteQuotasProjectIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteQuotasProjectIDForbiddenCode is the HTTP code returned for type DeleteQuotasProjectIDForbidden
const DeleteQuotasProjectIDForbiddenCode int = 403

/*
DeleteQuotasProjectIDForbidden Forbidden

swagger:response deleteQuotasProjectIdForbidden
*/
type DeleteQuotasProjectIDForbidden struct {
}

// NewDeleteQuotasProjectIDForbidden creates DeleteQuotasProjectIDForbidden with default headers values
func NewDeleteQuotasProjectIDForbidden() *DeleteQuotasProjectIDForbidden {

	return &DeleteQuotasProjectIDForbidden{}
}

// WriteResponse to the client
func (o *DeleteQuotasProjectIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DeleteQuotasProjectIDNotFoundCode is the HTTP code returned for type DeleteQuotasProjectIDNotFound
const DeleteQuotasProjectIDNotFoundCode int = 404

/*
DeleteQuotasProjectIDNotFound Not Found

swagger:response deleteQuotasProjectIdNotFound
*/
type DeleteQuotasProjectIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteQuotasProjectIDNotFound creates DeleteQuotasProjectIDNotFound with default headers values
func NewDeleteQuotasProjectIDNotFound() *DeleteQuotasProjectIDNotFound {

	return &DeleteQuotasProjectIDNotFound{}
}

// WithPayload adds the payload to the delete quotas project Id not found response
func (o *DeleteQuotasProjectIDNotFound) WithPayload(payload *models.Error) *DeleteQuotasProjectIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete quotas project Id not found response
func (o *DeleteQuotasProjectIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteQuotasProjectIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteQuotasProjectIDUnprocessableEntityCode is the HTTP code returned for type DeleteQuotasProjectIDUnprocessableEntity
const DeleteQuotasProjectIDUnprocessableEntityCode int = 422

/*
DeleteQuotasProjectIDUnprocessableEntity Unprocessable Content

swagger:response deleteQuotasProjectIdUnprocessableEntity
*/
type DeleteQuotasProjectIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteQuotasProjectIDUnprocessableEntity creates DeleteQuotasProjectIDUnprocessableEntity with default headers values
func NewDeleteQuotasProjectIDUnprocessableEntity() *DeleteQuotasProjectIDUnprocessableEntity {

	return &DeleteQuotasProjectIDUnprocessableEntity{}
}

// WithPayload adds the payload to the delete quotas project Id unprocessable entity response
func (o *DeleteQuotasProjectIDUnprocessableEntity) WithPayload(payload *models.Error) *DeleteQuotasProjectIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete quotas project Id unprocessable entity response
func (o *DeleteQuotasProjectIDUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteQuotasProjectIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
