// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

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
)

// NewListVMwareCloudDirectorComputePoliciesNoCredentialsParams creates a new ListVMwareCloudDirectorComputePoliciesNoCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVMwareCloudDirectorComputePoliciesNoCredentialsParams() *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	return &ListVMwareCloudDirectorComputePoliciesNoCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVMwareCloudDirectorComputePoliciesNoCredentialsParamsWithTimeout creates a new ListVMwareCloudDirectorComputePoliciesNoCredentialsParams object
// with the ability to set a timeout on a request.
func NewListVMwareCloudDirectorComputePoliciesNoCredentialsParamsWithTimeout(timeout time.Duration) *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	return &ListVMwareCloudDirectorComputePoliciesNoCredentialsParams{
		timeout: timeout,
	}
}

// NewListVMwareCloudDirectorComputePoliciesNoCredentialsParamsWithContext creates a new ListVMwareCloudDirectorComputePoliciesNoCredentialsParams object
// with the ability to set a context for a request.
func NewListVMwareCloudDirectorComputePoliciesNoCredentialsParamsWithContext(ctx context.Context) *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	return &ListVMwareCloudDirectorComputePoliciesNoCredentialsParams{
		Context: ctx,
	}
}

// NewListVMwareCloudDirectorComputePoliciesNoCredentialsParamsWithHTTPClient creates a new ListVMwareCloudDirectorComputePoliciesNoCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVMwareCloudDirectorComputePoliciesNoCredentialsParamsWithHTTPClient(client *http.Client) *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	return &ListVMwareCloudDirectorComputePoliciesNoCredentialsParams{
		HTTPClient: client,
	}
}

/*
ListVMwareCloudDirectorComputePoliciesNoCredentialsParams contains all the parameters to send to the API endpoint

	for the list v mware cloud director compute policies no credentials operation.

	Typically these are written to a http.Request.
*/
type ListVMwareCloudDirectorComputePoliciesNoCredentialsParams struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list v mware cloud director compute policies no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) WithDefaults() *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list v mware cloud director compute policies no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) WithTimeout(timeout time.Duration) *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) WithContext(ctx context.Context) *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) WithHTTPClient(client *http.Client) *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) WithClusterID(clusterID string) *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) WithProjectID(projectID string) *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list v mware cloud director compute policies no credentials params
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListVMwareCloudDirectorComputePoliciesNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}