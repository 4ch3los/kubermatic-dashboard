// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteApplicationInstallationReader is a Reader for the DeleteApplicationInstallation structure.
type DeleteApplicationInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApplicationInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApplicationInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteApplicationInstallationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteApplicationInstallationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteApplicationInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteApplicationInstallationOK creates a DeleteApplicationInstallationOK with default headers values
func NewDeleteApplicationInstallationOK() *DeleteApplicationInstallationOK {
	return &DeleteApplicationInstallationOK{}
}

/*
DeleteApplicationInstallationOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteApplicationInstallationOK struct {
}

// IsSuccess returns true when this delete application installation o k response has a 2xx status code
func (o *DeleteApplicationInstallationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete application installation o k response has a 3xx status code
func (o *DeleteApplicationInstallationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete application installation o k response has a 4xx status code
func (o *DeleteApplicationInstallationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete application installation o k response has a 5xx status code
func (o *DeleteApplicationInstallationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete application installation o k response a status code equal to that given
func (o *DeleteApplicationInstallationOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteApplicationInstallationOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] deleteApplicationInstallationOK ", 200)
}

func (o *DeleteApplicationInstallationOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] deleteApplicationInstallationOK ", 200)
}

func (o *DeleteApplicationInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationInstallationUnauthorized creates a DeleteApplicationInstallationUnauthorized with default headers values
func NewDeleteApplicationInstallationUnauthorized() *DeleteApplicationInstallationUnauthorized {
	return &DeleteApplicationInstallationUnauthorized{}
}

/*
DeleteApplicationInstallationUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteApplicationInstallationUnauthorized struct {
}

// IsSuccess returns true when this delete application installation unauthorized response has a 2xx status code
func (o *DeleteApplicationInstallationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete application installation unauthorized response has a 3xx status code
func (o *DeleteApplicationInstallationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete application installation unauthorized response has a 4xx status code
func (o *DeleteApplicationInstallationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete application installation unauthorized response has a 5xx status code
func (o *DeleteApplicationInstallationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete application installation unauthorized response a status code equal to that given
func (o *DeleteApplicationInstallationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteApplicationInstallationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] deleteApplicationInstallationUnauthorized ", 401)
}

func (o *DeleteApplicationInstallationUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] deleteApplicationInstallationUnauthorized ", 401)
}

func (o *DeleteApplicationInstallationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationInstallationForbidden creates a DeleteApplicationInstallationForbidden with default headers values
func NewDeleteApplicationInstallationForbidden() *DeleteApplicationInstallationForbidden {
	return &DeleteApplicationInstallationForbidden{}
}

/*
DeleteApplicationInstallationForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteApplicationInstallationForbidden struct {
}

// IsSuccess returns true when this delete application installation forbidden response has a 2xx status code
func (o *DeleteApplicationInstallationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete application installation forbidden response has a 3xx status code
func (o *DeleteApplicationInstallationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete application installation forbidden response has a 4xx status code
func (o *DeleteApplicationInstallationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete application installation forbidden response has a 5xx status code
func (o *DeleteApplicationInstallationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete application installation forbidden response a status code equal to that given
func (o *DeleteApplicationInstallationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteApplicationInstallationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] deleteApplicationInstallationForbidden ", 403)
}

func (o *DeleteApplicationInstallationForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] deleteApplicationInstallationForbidden ", 403)
}

func (o *DeleteApplicationInstallationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationInstallationDefault creates a DeleteApplicationInstallationDefault with default headers values
func NewDeleteApplicationInstallationDefault(code int) *DeleteApplicationInstallationDefault {
	return &DeleteApplicationInstallationDefault{
		_statusCode: code,
	}
}

/*
DeleteApplicationInstallationDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteApplicationInstallationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete application installation default response
func (o *DeleteApplicationInstallationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete application installation default response has a 2xx status code
func (o *DeleteApplicationInstallationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete application installation default response has a 3xx status code
func (o *DeleteApplicationInstallationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete application installation default response has a 4xx status code
func (o *DeleteApplicationInstallationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete application installation default response has a 5xx status code
func (o *DeleteApplicationInstallationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete application installation default response a status code equal to that given
func (o *DeleteApplicationInstallationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteApplicationInstallationDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] deleteApplicationInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteApplicationInstallationDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] deleteApplicationInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteApplicationInstallationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteApplicationInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}