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
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetEndpointParams creates a new GetEndpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEndpointParams() *GetEndpointParams {
	return &GetEndpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointParamsWithTimeout creates a new GetEndpointParams object
// with the ability to set a timeout on a request.
func NewGetEndpointParamsWithTimeout(timeout time.Duration) *GetEndpointParams {
	return &GetEndpointParams{
		timeout: timeout,
	}
}

// NewGetEndpointParamsWithContext creates a new GetEndpointParams object
// with the ability to set a context for a request.
func NewGetEndpointParamsWithContext(ctx context.Context) *GetEndpointParams {
	return &GetEndpointParams{
		Context: ctx,
	}
}

// NewGetEndpointParamsWithHTTPClient creates a new GetEndpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEndpointParamsWithHTTPClient(client *http.Client) *GetEndpointParams {
	return &GetEndpointParams{
		HTTPClient: client,
	}
}

/*
GetEndpointParams contains all the parameters to send to the API endpoint

	for the get endpoint operation.

	Typically these are written to a http.Request.
*/
type GetEndpointParams struct {

	/* Limit.

	   Sets the page size.
	*/
	Limit *int64

	/* Marker.

	   Pagination ID of the last item in the previous list.

	   Format: uuid
	*/
	Marker *strfmt.UUID

	/* NotTags.

	     Filter for resources not having tags, multiple not-tags are considered as logical AND.
	Should be provided in a comma separated list.

	*/
	NotTags []string

	/* NotTagsAny.

	     Filter for resources not having tags, multiple tags are considered as logical OR.
	Should be provided in a comma separated list.

	*/
	NotTagsAny []string

	/* PageReverse.

	   Sets the page direction.
	*/
	PageReverse *bool

	/* ProjectID.

	   Filter for resources belonging or accessible by a specific project.

	*/
	ProjectID *string

	/* Sort.

	   Comma-separated list of sort keys, optionally prefix with - to reverse sort order.
	*/
	Sort *string

	/* Tags.

	     Filter for tags, multiple tags are considered as logical AND.
	Should be provided in a comma separated list.

	*/
	Tags []string

	/* TagsAny.

	     Filter for tags, multiple tags are considered as logical OR.
	Should be provided in a comma separated list.

	*/
	TagsAny []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointParams) WithDefaults() *GetEndpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get endpoint params
func (o *GetEndpointParams) WithTimeout(timeout time.Duration) *GetEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoint params
func (o *GetEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoint params
func (o *GetEndpointParams) WithContext(ctx context.Context) *GetEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoint params
func (o *GetEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoint params
func (o *GetEndpointParams) WithHTTPClient(client *http.Client) *GetEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoint params
func (o *GetEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get endpoint params
func (o *GetEndpointParams) WithLimit(limit *int64) *GetEndpointParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get endpoint params
func (o *GetEndpointParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMarker adds the marker to the get endpoint params
func (o *GetEndpointParams) WithMarker(marker *strfmt.UUID) *GetEndpointParams {
	o.SetMarker(marker)
	return o
}

// SetMarker adds the marker to the get endpoint params
func (o *GetEndpointParams) SetMarker(marker *strfmt.UUID) {
	o.Marker = marker
}

// WithNotTags adds the notTags to the get endpoint params
func (o *GetEndpointParams) WithNotTags(notTags []string) *GetEndpointParams {
	o.SetNotTags(notTags)
	return o
}

// SetNotTags adds the notTags to the get endpoint params
func (o *GetEndpointParams) SetNotTags(notTags []string) {
	o.NotTags = notTags
}

// WithNotTagsAny adds the notTagsAny to the get endpoint params
func (o *GetEndpointParams) WithNotTagsAny(notTagsAny []string) *GetEndpointParams {
	o.SetNotTagsAny(notTagsAny)
	return o
}

// SetNotTagsAny adds the notTagsAny to the get endpoint params
func (o *GetEndpointParams) SetNotTagsAny(notTagsAny []string) {
	o.NotTagsAny = notTagsAny
}

// WithPageReverse adds the pageReverse to the get endpoint params
func (o *GetEndpointParams) WithPageReverse(pageReverse *bool) *GetEndpointParams {
	o.SetPageReverse(pageReverse)
	return o
}

// SetPageReverse adds the pageReverse to the get endpoint params
func (o *GetEndpointParams) SetPageReverse(pageReverse *bool) {
	o.PageReverse = pageReverse
}

// WithProjectID adds the projectID to the get endpoint params
func (o *GetEndpointParams) WithProjectID(projectID *string) *GetEndpointParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get endpoint params
func (o *GetEndpointParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithSort adds the sort to the get endpoint params
func (o *GetEndpointParams) WithSort(sort *string) *GetEndpointParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get endpoint params
func (o *GetEndpointParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTags adds the tags to the get endpoint params
func (o *GetEndpointParams) WithTags(tags []string) *GetEndpointParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get endpoint params
func (o *GetEndpointParams) SetTags(tags []string) {
	o.Tags = tags
}

// WithTagsAny adds the tagsAny to the get endpoint params
func (o *GetEndpointParams) WithTagsAny(tagsAny []string) *GetEndpointParams {
	o.SetTagsAny(tagsAny)
	return o
}

// SetTagsAny adds the tagsAny to the get endpoint params
func (o *GetEndpointParams) SetTagsAny(tagsAny []string) {
	o.TagsAny = tagsAny
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Marker != nil {

		// query param marker
		var qrMarker strfmt.UUID

		if o.Marker != nil {
			qrMarker = *o.Marker
		}
		qMarker := qrMarker.String()
		if qMarker != "" {

			if err := r.SetQueryParam("marker", qMarker); err != nil {
				return err
			}
		}
	}

	if o.NotTags != nil {

		// binding items for not-tags
		joinedNotTags := o.bindParamNotTags(reg)

		// query array param not-tags
		if err := r.SetQueryParam("not-tags", joinedNotTags...); err != nil {
			return err
		}
	}

	if o.NotTagsAny != nil {

		// binding items for not-tags-any
		joinedNotTagsAny := o.bindParamNotTagsAny(reg)

		// query array param not-tags-any
		if err := r.SetQueryParam("not-tags-any", joinedNotTagsAny...); err != nil {
			return err
		}
	}

	if o.PageReverse != nil {

		// query param page_reverse
		var qrPageReverse bool

		if o.PageReverse != nil {
			qrPageReverse = *o.PageReverse
		}
		qPageReverse := swag.FormatBool(qrPageReverse)
		if qPageReverse != "" {

			if err := r.SetQueryParam("page_reverse", qPageReverse); err != nil {
				return err
			}
		}
	}

	if o.ProjectID != nil {

		// query param project_id
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("project_id", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Tags != nil {

		// binding items for tags
		joinedTags := o.bindParamTags(reg)

		// query array param tags
		if err := r.SetQueryParam("tags", joinedTags...); err != nil {
			return err
		}
	}

	if o.TagsAny != nil {

		// binding items for tags-any
		joinedTagsAny := o.bindParamTagsAny(reg)

		// query array param tags-any
		if err := r.SetQueryParam("tags-any", joinedTagsAny...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetEndpoint binds the parameter not-tags
func (o *GetEndpointParams) bindParamNotTags(formats strfmt.Registry) []string {
	notTagsIR := o.NotTags

	var notTagsIC []string
	for _, notTagsIIR := range notTagsIR { // explode []string

		notTagsIIV := notTagsIIR // string as string
		notTagsIC = append(notTagsIC, notTagsIIV)
	}

	// items.CollectionFormat: ""
	notTagsIS := swag.JoinByFormat(notTagsIC, "")

	return notTagsIS
}

// bindParamGetEndpoint binds the parameter not-tags-any
func (o *GetEndpointParams) bindParamNotTagsAny(formats strfmt.Registry) []string {
	notTagsAnyIR := o.NotTagsAny

	var notTagsAnyIC []string
	for _, notTagsAnyIIR := range notTagsAnyIR { // explode []string

		notTagsAnyIIV := notTagsAnyIIR // string as string
		notTagsAnyIC = append(notTagsAnyIC, notTagsAnyIIV)
	}

	// items.CollectionFormat: ""
	notTagsAnyIS := swag.JoinByFormat(notTagsAnyIC, "")

	return notTagsAnyIS
}

// bindParamGetEndpoint binds the parameter tags
func (o *GetEndpointParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: ""
	tagsIS := swag.JoinByFormat(tagsIC, "")

	return tagsIS
}

// bindParamGetEndpoint binds the parameter tags-any
func (o *GetEndpointParams) bindParamTagsAny(formats strfmt.Registry) []string {
	tagsAnyIR := o.TagsAny

	var tagsAnyIC []string
	for _, tagsAnyIIR := range tagsAnyIR { // explode []string

		tagsAnyIIV := tagsAnyIIR // string as string
		tagsAnyIC = append(tagsAnyIC, tagsAnyIIV)
	}

	// items.CollectionFormat: ""
	tagsAnyIS := swag.JoinByFormat(tagsAnyIC, "")

	return tagsAnyIS
}
