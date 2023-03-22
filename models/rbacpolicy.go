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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Rbacpolicy rbacpolicy
//
// swagger:model rbacpolicy
type Rbacpolicy struct {

	// created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The ID of the resource.
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// project id
	ProjectID Project `json:"project_id"`

	// The ID of the service resource.
	// Required: true
	// Format: uuid
	ServiceID *strfmt.UUID `json:"service_id"`

	// The ID of the project to which the RBAC policy will be enforced.
	// Example: 666da95112694b37b3efb0913de3f499
	Target string `json:"target,omitempty"`

	// target type
	// Enum: [project_id domain_id]
	TargetType string `json:"target_type,omitempty"`

	// updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Validate validates this rbacpolicy
func (m *Rbacpolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rbacpolicy) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	return nil
}

func (m *Rbacpolicy) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Rbacpolicy) validateProjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := m.ProjectID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("project_id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("project_id")
		}
		return err
	}

	return nil
}

func (m *Rbacpolicy) validateServiceID(formats strfmt.Registry) error {

	if err := validate.Required("service_id", "body", m.ServiceID); err != nil {
		return err
	}

	if err := validate.FormatOf("service_id", "body", "uuid", m.ServiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

var rbacpolicyTypeTargetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["project_id","domain_id"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rbacpolicyTypeTargetTypePropEnum = append(rbacpolicyTypeTargetTypePropEnum, v)
	}
}

const (

	// RbacpolicyTargetTypeProjectID captures enum value "project_id"
	RbacpolicyTargetTypeProjectID string = "project_id"

	// RbacpolicyTargetTypeDomainID captures enum value "domain_id"
	RbacpolicyTargetTypeDomainID string = "domain_id"
)

// prop value enum
func (m *Rbacpolicy) validateTargetTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rbacpolicyTypeTargetTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Rbacpolicy) validateTargetType(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTargetTypeEnum("target_type", "body", m.TargetType); err != nil {
		return err
	}

	return nil
}

func (m *Rbacpolicy) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this rbacpolicy based on the context it is used
func (m *Rbacpolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rbacpolicy) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Rbacpolicy) contextValidateProjectID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ProjectID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("project_id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("project_id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rbacpolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rbacpolicy) UnmarshalBinary(b []byte) error {
	var res Rbacpolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
