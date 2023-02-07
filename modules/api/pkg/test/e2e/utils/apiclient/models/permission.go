// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Permission Permission represents the permissions (i.e. role and clusterRole) associated to an object.
//
// swagger:model Permission
type Permission struct {

	// Namespace on which the permission is given. Empty if scope is Cluster
	Namespace string `json:"namespace,omitempty"`

	// RoleRefName is the name of the clusterRole or Role.
	RoleRefName string `json:"roleRefName,omitempty"`

	// Scope of the permission. Either "Cluster" or "Namespace".
	Scope string `json:"scope,omitempty"`
}

// Validate validates this permission
func (m *Permission) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this permission based on context it is used
func (m *Permission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Permission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Permission) UnmarshalBinary(b []byte) error {
	var res Permission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}