// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationInstallationSpec application installation spec
//
// swagger:model ApplicationInstallationSpec
type ApplicationInstallationSpec struct {

	// application ref
	ApplicationRef *ApplicationRef `json:"applicationRef,omitempty"`

	// deploy options
	DeployOptions *DeployOptions `json:"deployOptions,omitempty"`

	// namespace
	Namespace *NamespaceSpec `json:"namespace,omitempty"`

	// reconciliation interval
	ReconciliationInterval Duration `json:"reconciliationInterval,omitempty"`

	// values
	Values RawExtension `json:"values,omitempty"`
}

// Validate validates this application installation spec
func (m *ApplicationInstallationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationInstallationSpec) validateApplicationRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationRef) { // not required
		return nil
	}

	if m.ApplicationRef != nil {
		if err := m.ApplicationRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationRef")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationInstallationSpec) validateDeployOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DeployOptions) { // not required
		return nil
	}

	if m.DeployOptions != nil {
		if err := m.DeployOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationInstallationSpec) validateNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	if m.Namespace != nil {
		if err := m.Namespace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application installation spec based on the context it is used
func (m *ApplicationInstallationSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeployOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationInstallationSpec) contextValidateApplicationRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationRef != nil {
		if err := m.ApplicationRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationRef")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationInstallationSpec) contextValidateDeployOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.DeployOptions != nil {
		if err := m.DeployOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationInstallationSpec) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

	if m.Namespace != nil {
		if err := m.Namespace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationInstallationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationInstallationSpec) UnmarshalBinary(b []byte) error {
	var res ApplicationInstallationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}