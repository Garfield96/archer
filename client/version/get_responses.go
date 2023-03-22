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

package version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/archer/models"
)

// GetReader is a Reader for the Get structure.
type GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOK creates a GetOK with default headers values
func NewGetOK() *GetOK {
	return &GetOK{}
}

/*
GetOK describes a response with status code 200, with default header values.

Version
*/
type GetOK struct {
	Payload *models.Version
}

// IsSuccess returns true when this get o k response has a 2xx status code
func (o *GetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get o k response has a 3xx status code
func (o *GetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get o k response has a 4xx status code
func (o *GetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get o k response has a 5xx status code
func (o *GetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get o k response a status code equal to that given
func (o *GetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get o k response
func (o *GetOK) Code() int {
	return 200
}

func (o *GetOK) Error() string {
	return fmt.Sprintf("[GET /][%d] getOK  %+v", 200, o.Payload)
}

func (o *GetOK) String() string {
	return fmt.Sprintf("[GET /][%d] getOK  %+v", 200, o.Payload)
}

func (o *GetOK) GetPayload() *models.Version {
	return o.Payload
}

func (o *GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Version)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
