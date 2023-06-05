// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Protocol Protocol defines network protocols supported for things like container ports.
//
// +enum
//
// swagger:model Protocol
type Protocol string

// Validate validates this protocol
func (m Protocol) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this protocol based on context it is used
func (m Protocol) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}