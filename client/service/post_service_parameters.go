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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/archer/models"
)

// NewPostServiceParams creates a new PostServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostServiceParams() *PostServiceParams {
	return &PostServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostServiceParamsWithTimeout creates a new PostServiceParams object
// with the ability to set a timeout on a request.
func NewPostServiceParamsWithTimeout(timeout time.Duration) *PostServiceParams {
	return &PostServiceParams{
		timeout: timeout,
	}
}

// NewPostServiceParamsWithContext creates a new PostServiceParams object
// with the ability to set a context for a request.
func NewPostServiceParamsWithContext(ctx context.Context) *PostServiceParams {
	return &PostServiceParams{
		Context: ctx,
	}
}

// NewPostServiceParamsWithHTTPClient creates a new PostServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostServiceParamsWithHTTPClient(client *http.Client) *PostServiceParams {
	return &PostServiceParams{
		HTTPClient: client,
	}
}

/*
PostServiceParams contains all the parameters to send to the API endpoint

	for the post service operation.

	Typically these are written to a http.Request.
*/
type PostServiceParams struct {

	/* Body.

	   Service object that needs to be added to the catalog
	*/
	Body *models.Service

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostServiceParams) WithDefaults() *PostServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post service params
func (o *PostServiceParams) WithTimeout(timeout time.Duration) *PostServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post service params
func (o *PostServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post service params
func (o *PostServiceParams) WithContext(ctx context.Context) *PostServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post service params
func (o *PostServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post service params
func (o *PostServiceParams) WithHTTPClient(client *http.Client) *PostServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post service params
func (o *PostServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post service params
func (o *PostServiceParams) WithBody(body *models.Service) *PostServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post service params
func (o *PostServiceParams) SetBody(body *models.Service) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
