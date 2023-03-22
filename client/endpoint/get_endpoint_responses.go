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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/sapcc/archer/models"
)

// GetEndpointReader is a Reader for the GetEndpoint structure.
type GetEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEndpointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEndpointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEndpointForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetEndpointUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEndpointOK creates a GetEndpointOK with default headers values
func NewGetEndpointOK() *GetEndpointOK {
	return &GetEndpointOK{}
}

/*
GetEndpointOK describes a response with status code 200, with default header values.

An array of endpoints.
*/
type GetEndpointOK struct {
	Payload *GetEndpointOKBody
}

// IsSuccess returns true when this get endpoint o k response has a 2xx status code
func (o *GetEndpointOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get endpoint o k response has a 3xx status code
func (o *GetEndpointOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint o k response has a 4xx status code
func (o *GetEndpointOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint o k response has a 5xx status code
func (o *GetEndpointOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint o k response a status code equal to that given
func (o *GetEndpointOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get endpoint o k response
func (o *GetEndpointOK) Code() int {
	return 200
}

func (o *GetEndpointOK) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointOK  %+v", 200, o.Payload)
}

func (o *GetEndpointOK) String() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointOK  %+v", 200, o.Payload)
}

func (o *GetEndpointOK) GetPayload() *GetEndpointOKBody {
	return o.Payload
}

func (o *GetEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEndpointOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointBadRequest creates a GetEndpointBadRequest with default headers values
func NewGetEndpointBadRequest() *GetEndpointBadRequest {
	return &GetEndpointBadRequest{}
}

/*
GetEndpointBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetEndpointBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get endpoint bad request response has a 2xx status code
func (o *GetEndpointBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint bad request response has a 3xx status code
func (o *GetEndpointBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint bad request response has a 4xx status code
func (o *GetEndpointBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint bad request response has a 5xx status code
func (o *GetEndpointBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint bad request response a status code equal to that given
func (o *GetEndpointBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get endpoint bad request response
func (o *GetEndpointBadRequest) Code() int {
	return 400
}

func (o *GetEndpointBadRequest) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointBadRequest  %+v", 400, o.Payload)
}

func (o *GetEndpointBadRequest) String() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointBadRequest  %+v", 400, o.Payload)
}

func (o *GetEndpointBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEndpointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointUnauthorized creates a GetEndpointUnauthorized with default headers values
func NewGetEndpointUnauthorized() *GetEndpointUnauthorized {
	return &GetEndpointUnauthorized{}
}

/*
GetEndpointUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetEndpointUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get endpoint unauthorized response has a 2xx status code
func (o *GetEndpointUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint unauthorized response has a 3xx status code
func (o *GetEndpointUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint unauthorized response has a 4xx status code
func (o *GetEndpointUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint unauthorized response has a 5xx status code
func (o *GetEndpointUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint unauthorized response a status code equal to that given
func (o *GetEndpointUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get endpoint unauthorized response
func (o *GetEndpointUnauthorized) Code() int {
	return 401
}

func (o *GetEndpointUnauthorized) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointUnauthorized  %+v", 401, o.Payload)
}

func (o *GetEndpointUnauthorized) String() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointUnauthorized  %+v", 401, o.Payload)
}

func (o *GetEndpointUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEndpointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointForbidden creates a GetEndpointForbidden with default headers values
func NewGetEndpointForbidden() *GetEndpointForbidden {
	return &GetEndpointForbidden{}
}

/*
GetEndpointForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetEndpointForbidden struct {
}

// IsSuccess returns true when this get endpoint forbidden response has a 2xx status code
func (o *GetEndpointForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint forbidden response has a 3xx status code
func (o *GetEndpointForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint forbidden response has a 4xx status code
func (o *GetEndpointForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint forbidden response has a 5xx status code
func (o *GetEndpointForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint forbidden response a status code equal to that given
func (o *GetEndpointForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get endpoint forbidden response
func (o *GetEndpointForbidden) Code() int {
	return 403
}

func (o *GetEndpointForbidden) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointForbidden ", 403)
}

func (o *GetEndpointForbidden) String() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointForbidden ", 403)
}

func (o *GetEndpointForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointUnprocessableEntity creates a GetEndpointUnprocessableEntity with default headers values
func NewGetEndpointUnprocessableEntity() *GetEndpointUnprocessableEntity {
	return &GetEndpointUnprocessableEntity{}
}

/*
GetEndpointUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Content
*/
type GetEndpointUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this get endpoint unprocessable entity response has a 2xx status code
func (o *GetEndpointUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint unprocessable entity response has a 3xx status code
func (o *GetEndpointUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint unprocessable entity response has a 4xx status code
func (o *GetEndpointUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint unprocessable entity response has a 5xx status code
func (o *GetEndpointUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint unprocessable entity response a status code equal to that given
func (o *GetEndpointUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get endpoint unprocessable entity response
func (o *GetEndpointUnprocessableEntity) Code() int {
	return 422
}

func (o *GetEndpointUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetEndpointUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetEndpointUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEndpointUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetEndpointOKBody get endpoint o k body
swagger:model GetEndpointOKBody
*/
type GetEndpointOKBody struct {

	// items
	Items []*models.Endpoint `json:"items"`

	// links
	Links []*models.Link `json:"links,omitempty"`
}

// Validate validates this get endpoint o k body
func (o *GetEndpointOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEndpointOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEndpointOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEndpointOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetEndpointOKBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	for i := 0; i < len(o.Links); i++ {
		if swag.IsZero(o.Links[i]) { // not required
			continue
		}

		if o.Links[i] != nil {
			if err := o.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEndpointOK" + "." + "links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEndpointOK" + "." + "links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get endpoint o k body based on the context it is used
func (o *GetEndpointOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEndpointOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEndpointOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEndpointOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetEndpointOKBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Links); i++ {

		if o.Links[i] != nil {
			if err := o.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEndpointOK" + "." + "links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEndpointOK" + "." + "links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEndpointOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEndpointOKBody) UnmarshalBinary(b []byte) error {
	var res GetEndpointOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
