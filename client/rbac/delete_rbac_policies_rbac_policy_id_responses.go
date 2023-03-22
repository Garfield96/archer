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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/archer/models"
)

// DeleteRbacPoliciesRbacPolicyIDReader is a Reader for the DeleteRbacPoliciesRbacPolicyID structure.
type DeleteRbacPoliciesRbacPolicyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRbacPoliciesRbacPolicyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRbacPoliciesRbacPolicyIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRbacPoliciesRbacPolicyIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRbacPoliciesRbacPolicyIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRbacPoliciesRbacPolicyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteRbacPoliciesRbacPolicyIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRbacPoliciesRbacPolicyIDNoContent creates a DeleteRbacPoliciesRbacPolicyIDNoContent with default headers values
func NewDeleteRbacPoliciesRbacPolicyIDNoContent() *DeleteRbacPoliciesRbacPolicyIDNoContent {
	return &DeleteRbacPoliciesRbacPolicyIDNoContent{}
}

/*
DeleteRbacPoliciesRbacPolicyIDNoContent describes a response with status code 204, with default header values.

RBAC policy successfully deleted.
*/
type DeleteRbacPoliciesRbacPolicyIDNoContent struct {
}

// IsSuccess returns true when this delete rbac policies rbac policy Id no content response has a 2xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete rbac policies rbac policy Id no content response has a 3xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rbac policies rbac policy Id no content response has a 4xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rbac policies rbac policy Id no content response has a 5xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rbac policies rbac policy Id no content response a status code equal to that given
func (o *DeleteRbacPoliciesRbacPolicyIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete rbac policies rbac policy Id no content response
func (o *DeleteRbacPoliciesRbacPolicyIDNoContent) Code() int {
	return 204
}

func (o *DeleteRbacPoliciesRbacPolicyIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdNoContent ", 204)
}

func (o *DeleteRbacPoliciesRbacPolicyIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdNoContent ", 204)
}

func (o *DeleteRbacPoliciesRbacPolicyIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRbacPoliciesRbacPolicyIDUnauthorized creates a DeleteRbacPoliciesRbacPolicyIDUnauthorized with default headers values
func NewDeleteRbacPoliciesRbacPolicyIDUnauthorized() *DeleteRbacPoliciesRbacPolicyIDUnauthorized {
	return &DeleteRbacPoliciesRbacPolicyIDUnauthorized{}
}

/*
DeleteRbacPoliciesRbacPolicyIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRbacPoliciesRbacPolicyIDUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete rbac policies rbac policy Id unauthorized response has a 2xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rbac policies rbac policy Id unauthorized response has a 3xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rbac policies rbac policy Id unauthorized response has a 4xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rbac policies rbac policy Id unauthorized response has a 5xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rbac policies rbac policy Id unauthorized response a status code equal to that given
func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete rbac policies rbac policy Id unauthorized response
func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRbacPoliciesRbacPolicyIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRbacPoliciesRbacPolicyIDForbidden creates a DeleteRbacPoliciesRbacPolicyIDForbidden with default headers values
func NewDeleteRbacPoliciesRbacPolicyIDForbidden() *DeleteRbacPoliciesRbacPolicyIDForbidden {
	return &DeleteRbacPoliciesRbacPolicyIDForbidden{}
}

/*
DeleteRbacPoliciesRbacPolicyIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRbacPoliciesRbacPolicyIDForbidden struct {
}

// IsSuccess returns true when this delete rbac policies rbac policy Id forbidden response has a 2xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rbac policies rbac policy Id forbidden response has a 3xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rbac policies rbac policy Id forbidden response has a 4xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rbac policies rbac policy Id forbidden response has a 5xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rbac policies rbac policy Id forbidden response a status code equal to that given
func (o *DeleteRbacPoliciesRbacPolicyIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete rbac policies rbac policy Id forbidden response
func (o *DeleteRbacPoliciesRbacPolicyIDForbidden) Code() int {
	return 403
}

func (o *DeleteRbacPoliciesRbacPolicyIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdForbidden ", 403)
}

func (o *DeleteRbacPoliciesRbacPolicyIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdForbidden ", 403)
}

func (o *DeleteRbacPoliciesRbacPolicyIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRbacPoliciesRbacPolicyIDNotFound creates a DeleteRbacPoliciesRbacPolicyIDNotFound with default headers values
func NewDeleteRbacPoliciesRbacPolicyIDNotFound() *DeleteRbacPoliciesRbacPolicyIDNotFound {
	return &DeleteRbacPoliciesRbacPolicyIDNotFound{}
}

/*
DeleteRbacPoliciesRbacPolicyIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteRbacPoliciesRbacPolicyIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete rbac policies rbac policy Id not found response has a 2xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rbac policies rbac policy Id not found response has a 3xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rbac policies rbac policy Id not found response has a 4xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rbac policies rbac policy Id not found response has a 5xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rbac policies rbac policy Id not found response a status code equal to that given
func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete rbac policies rbac policy Id not found response
func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) Code() int {
	return 404
}

func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRbacPoliciesRbacPolicyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRbacPoliciesRbacPolicyIDUnprocessableEntity creates a DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity with default headers values
func NewDeleteRbacPoliciesRbacPolicyIDUnprocessableEntity() *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity {
	return &DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity{}
}

/*
DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Content
*/
type DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete rbac policies rbac policy Id unprocessable entity response has a 2xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rbac policies rbac policy Id unprocessable entity response has a 3xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rbac policies rbac policy Id unprocessable entity response has a 4xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rbac policies rbac policy Id unprocessable entity response has a 5xx status code
func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rbac policies rbac policy Id unprocessable entity response a status code equal to that given
func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete rbac policies rbac policy Id unprocessable entity response
func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /rbac-policies/{rbac_policy_id}][%d] deleteRbacPoliciesRbacPolicyIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRbacPoliciesRbacPolicyIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
