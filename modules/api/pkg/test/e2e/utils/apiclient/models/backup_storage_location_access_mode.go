// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// BackupStorageLocationAccessMode BackupStorageLocationAccessMode represents the permissions for a BackupStorageLocation.
//
// +kubebuilder:validation:Enum=ReadOnly;ReadWrite
//
// swagger:model BackupStorageLocationAccessMode
type BackupStorageLocationAccessMode string

// Validate validates this backup storage location access mode
func (m BackupStorageLocationAccessMode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this backup storage location access mode based on context it is used
func (m BackupStorageLocationAccessMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}