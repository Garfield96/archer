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
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/archer/models"
)

// PostServiceReader is a Reader for the PostService structure.
type PostServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostServiceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostServiceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostServiceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostServiceCreated creates a PostServiceCreated with default headers values
func NewPostServiceCreated() *PostServiceCreated {
	return &PostServiceCreated{}
}

/*
PostServiceCreated describes a response with status code 201, with default header values.

Service
*/
type PostServiceCreated struct {

	/* The UUID of the created resource

	   Format: uuid
	*/
	XTargetID strfmt.UUID

	Payload *models.Service
}

// IsSuccess returns true when this post service created response has a 2xx status code
func (o *PostServiceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post service created response has a 3xx status code
func (o *PostServiceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post service created response has a 4xx status code
func (o *PostServiceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post service created response has a 5xx status code
func (o *PostServiceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post service created response a status code equal to that given
func (o *PostServiceCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post service created response
func (o *PostServiceCreated) Code() int {
	return 201
}

func (o *PostServiceCreated) Error() string {
	return fmt.Sprintf("[POST /service][%d] postServiceCreated  %+v", 201, o.Payload)
}

func (o *PostServiceCreated) String() string {
	return fmt.Sprintf("[POST /service][%d] postServiceCreated  %+v", 201, o.Payload)
}

func (o *PostServiceCreated) GetPayload() *models.Service {
	return o.Payload
}

func (o *PostServiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Target-Id
	hdrXTargetID := response.GetHeader("X-Target-Id")

	if hdrXTargetID != "" {
		valxTargetId, err := formats.Parse("uuid", hdrXTargetID)
		if err != nil {
			return errors.InvalidType("X-Target-Id", "header", "strfmt.UUID", hdrXTargetID)
		}
		o.XTargetID = *(valxTargetId.(*strfmt.UUID))
	}

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServiceUnauthorized creates a PostServiceUnauthorized with default headers values
func NewPostServiceUnauthorized() *PostServiceUnauthorized {
	return &PostServiceUnauthorized{}
}

/*
PostServiceUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostServiceUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this post service unauthorized response has a 2xx status code
func (o *PostServiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post service unauthorized response has a 3xx status code
func (o *PostServiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post service unauthorized response has a 4xx status code
func (o *PostServiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post service unauthorized response has a 5xx status code
func (o *PostServiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post service unauthorized response a status code equal to that given
func (o *PostServiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post service unauthorized response
func (o *PostServiceUnauthorized) Code() int {
	return 401
}

func (o *PostServiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /service][%d] postServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostServiceUnauthorized) String() string {
	return fmt.Sprintf("[POST /service][%d] postServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostServiceUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServiceForbidden creates a PostServiceForbidden with default headers values
func NewPostServiceForbidden() *PostServiceForbidden {
	return &PostServiceForbidden{}
}

/*
PostServiceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostServiceForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this post service forbidden response has a 2xx status code
func (o *PostServiceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post service forbidden response has a 3xx status code
func (o *PostServiceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post service forbidden response has a 4xx status code
func (o *PostServiceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post service forbidden response has a 5xx status code
func (o *PostServiceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post service forbidden response a status code equal to that given
func (o *PostServiceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post service forbidden response
func (o *PostServiceForbidden) Code() int {
	return 403
}

func (o *PostServiceForbidden) Error() string {
	return fmt.Sprintf("[POST /service][%d] postServiceForbidden  %+v", 403, o.Payload)
}

func (o *PostServiceForbidden) String() string {
	return fmt.Sprintf("[POST /service][%d] postServiceForbidden  %+v", 403, o.Payload)
}

func (o *PostServiceForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServiceConflict creates a PostServiceConflict with default headers values
func NewPostServiceConflict() *PostServiceConflict {
	return &PostServiceConflict{}
}

/*
PostServiceConflict describes a response with status code 409, with default header values.

Duplicate entry
*/
type PostServiceConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this post service conflict response has a 2xx status code
func (o *PostServiceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post service conflict response has a 3xx status code
func (o *PostServiceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post service conflict response has a 4xx status code
func (o *PostServiceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post service conflict response has a 5xx status code
func (o *PostServiceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post service conflict response a status code equal to that given
func (o *PostServiceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post service conflict response
func (o *PostServiceConflict) Code() int {
	return 409
}

func (o *PostServiceConflict) Error() string {
	return fmt.Sprintf("[POST /service][%d] postServiceConflict  %+v", 409, o.Payload)
}

func (o *PostServiceConflict) String() string {
	return fmt.Sprintf("[POST /service][%d] postServiceConflict  %+v", 409, o.Payload)
}

func (o *PostServiceConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostServiceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServiceUnprocessableEntity creates a PostServiceUnprocessableEntity with default headers values
func NewPostServiceUnprocessableEntity() *PostServiceUnprocessableEntity {
	return &PostServiceUnprocessableEntity{}
}

/*
PostServiceUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Content
*/
type PostServiceUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this post service unprocessable entity response has a 2xx status code
func (o *PostServiceUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post service unprocessable entity response has a 3xx status code
func (o *PostServiceUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post service unprocessable entity response has a 4xx status code
func (o *PostServiceUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post service unprocessable entity response has a 5xx status code
func (o *PostServiceUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post service unprocessable entity response a status code equal to that given
func (o *PostServiceUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the post service unprocessable entity response
func (o *PostServiceUnprocessableEntity) Code() int {
	return 422
}

func (o *PostServiceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /service][%d] postServiceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostServiceUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /service][%d] postServiceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostServiceUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostServiceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
