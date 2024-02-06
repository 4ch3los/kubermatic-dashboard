// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DNSConfig DNSConfig contains a machine's DNS configuration.
//
// swagger:model DNSConfig
type DNSConfig struct {

	// servers
	Servers []string `json:"servers"`
}

// Validate validates this DNS config
func (m *DNSConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS config based on context it is used
func (m *DNSConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSConfig) UnmarshalBinary(b []byte) error {
	var res DNSConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}