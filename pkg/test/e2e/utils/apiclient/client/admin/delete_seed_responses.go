// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteSeedReader is a Reader for the DeleteSeed structure.
type DeleteSeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteSeedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSeedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSeedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSeedOK creates a DeleteSeedOK with default headers values
func NewDeleteSeedOK() *DeleteSeedOK {
	return &DeleteSeedOK{}
}

/*
DeleteSeedOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteSeedOK struct {
}

// IsSuccess returns true when this delete seed o k response has a 2xx status code
func (o *DeleteSeedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete seed o k response has a 3xx status code
func (o *DeleteSeedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete seed o k response has a 4xx status code
func (o *DeleteSeedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete seed o k response has a 5xx status code
func (o *DeleteSeedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete seed o k response a status code equal to that given
func (o *DeleteSeedOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteSeedOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeedOK ", 200)
}

func (o *DeleteSeedOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeedOK ", 200)
}

func (o *DeleteSeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSeedUnauthorized creates a DeleteSeedUnauthorized with default headers values
func NewDeleteSeedUnauthorized() *DeleteSeedUnauthorized {
	return &DeleteSeedUnauthorized{}
}

/*
DeleteSeedUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteSeedUnauthorized struct {
}

// IsSuccess returns true when this delete seed unauthorized response has a 2xx status code
func (o *DeleteSeedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete seed unauthorized response has a 3xx status code
func (o *DeleteSeedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete seed unauthorized response has a 4xx status code
func (o *DeleteSeedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete seed unauthorized response has a 5xx status code
func (o *DeleteSeedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete seed unauthorized response a status code equal to that given
func (o *DeleteSeedUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteSeedUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeedUnauthorized ", 401)
}

func (o *DeleteSeedUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeedUnauthorized ", 401)
}

func (o *DeleteSeedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSeedForbidden creates a DeleteSeedForbidden with default headers values
func NewDeleteSeedForbidden() *DeleteSeedForbidden {
	return &DeleteSeedForbidden{}
}

/*
DeleteSeedForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteSeedForbidden struct {
}

// IsSuccess returns true when this delete seed forbidden response has a 2xx status code
func (o *DeleteSeedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete seed forbidden response has a 3xx status code
func (o *DeleteSeedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete seed forbidden response has a 4xx status code
func (o *DeleteSeedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete seed forbidden response has a 5xx status code
func (o *DeleteSeedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete seed forbidden response a status code equal to that given
func (o *DeleteSeedForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteSeedForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeedForbidden ", 403)
}

func (o *DeleteSeedForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeedForbidden ", 403)
}

func (o *DeleteSeedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSeedDefault creates a DeleteSeedDefault with default headers values
func NewDeleteSeedDefault(code int) *DeleteSeedDefault {
	return &DeleteSeedDefault{
		_statusCode: code,
	}
}

/*
DeleteSeedDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteSeedDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete seed default response
func (o *DeleteSeedDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete seed default response has a 2xx status code
func (o *DeleteSeedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete seed default response has a 3xx status code
func (o *DeleteSeedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete seed default response has a 4xx status code
func (o *DeleteSeedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete seed default response has a 5xx status code
func (o *DeleteSeedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete seed default response a status code equal to that given
func (o *DeleteSeedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteSeedDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeed default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSeedDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeed default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSeedDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteSeedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}