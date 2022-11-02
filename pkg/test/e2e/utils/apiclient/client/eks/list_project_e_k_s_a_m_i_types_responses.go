// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListProjectEKSAMITypesReader is a Reader for the ListProjectEKSAMITypes structure.
type ListProjectEKSAMITypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectEKSAMITypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectEKSAMITypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListProjectEKSAMITypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectEKSAMITypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectEKSAMITypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectEKSAMITypesOK creates a ListProjectEKSAMITypesOK with default headers values
func NewListProjectEKSAMITypesOK() *ListProjectEKSAMITypesOK {
	return &ListProjectEKSAMITypesOK{}
}

/*
ListProjectEKSAMITypesOK describes a response with status code 200, with default header values.

EKSAMITypeList
*/
type ListProjectEKSAMITypesOK struct {
	Payload models.EKSAMITypeList
}

// IsSuccess returns true when this list project e k s a m i types o k response has a 2xx status code
func (o *ListProjectEKSAMITypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project e k s a m i types o k response has a 3xx status code
func (o *ListProjectEKSAMITypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project e k s a m i types o k response has a 4xx status code
func (o *ListProjectEKSAMITypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project e k s a m i types o k response has a 5xx status code
func (o *ListProjectEKSAMITypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project e k s a m i types o k response a status code equal to that given
func (o *ListProjectEKSAMITypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectEKSAMITypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/amitypes][%d] listProjectEKSAMITypesOK  %+v", 200, o.Payload)
}

func (o *ListProjectEKSAMITypesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/amitypes][%d] listProjectEKSAMITypesOK  %+v", 200, o.Payload)
}

func (o *ListProjectEKSAMITypesOK) GetPayload() models.EKSAMITypeList {
	return o.Payload
}

func (o *ListProjectEKSAMITypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectEKSAMITypesUnauthorized creates a ListProjectEKSAMITypesUnauthorized with default headers values
func NewListProjectEKSAMITypesUnauthorized() *ListProjectEKSAMITypesUnauthorized {
	return &ListProjectEKSAMITypesUnauthorized{}
}

/*
ListProjectEKSAMITypesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListProjectEKSAMITypesUnauthorized struct {
}

// IsSuccess returns true when this list project e k s a m i types unauthorized response has a 2xx status code
func (o *ListProjectEKSAMITypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project e k s a m i types unauthorized response has a 3xx status code
func (o *ListProjectEKSAMITypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project e k s a m i types unauthorized response has a 4xx status code
func (o *ListProjectEKSAMITypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project e k s a m i types unauthorized response has a 5xx status code
func (o *ListProjectEKSAMITypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list project e k s a m i types unauthorized response a status code equal to that given
func (o *ListProjectEKSAMITypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListProjectEKSAMITypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/amitypes][%d] listProjectEKSAMITypesUnauthorized ", 401)
}

func (o *ListProjectEKSAMITypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/amitypes][%d] listProjectEKSAMITypesUnauthorized ", 401)
}

func (o *ListProjectEKSAMITypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectEKSAMITypesForbidden creates a ListProjectEKSAMITypesForbidden with default headers values
func NewListProjectEKSAMITypesForbidden() *ListProjectEKSAMITypesForbidden {
	return &ListProjectEKSAMITypesForbidden{}
}

/*
ListProjectEKSAMITypesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListProjectEKSAMITypesForbidden struct {
}

// IsSuccess returns true when this list project e k s a m i types forbidden response has a 2xx status code
func (o *ListProjectEKSAMITypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project e k s a m i types forbidden response has a 3xx status code
func (o *ListProjectEKSAMITypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project e k s a m i types forbidden response has a 4xx status code
func (o *ListProjectEKSAMITypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project e k s a m i types forbidden response has a 5xx status code
func (o *ListProjectEKSAMITypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project e k s a m i types forbidden response a status code equal to that given
func (o *ListProjectEKSAMITypesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListProjectEKSAMITypesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/amitypes][%d] listProjectEKSAMITypesForbidden ", 403)
}

func (o *ListProjectEKSAMITypesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/amitypes][%d] listProjectEKSAMITypesForbidden ", 403)
}

func (o *ListProjectEKSAMITypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectEKSAMITypesDefault creates a ListProjectEKSAMITypesDefault with default headers values
func NewListProjectEKSAMITypesDefault(code int) *ListProjectEKSAMITypesDefault {
	return &ListProjectEKSAMITypesDefault{
		_statusCode: code,
	}
}

/*
ListProjectEKSAMITypesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectEKSAMITypesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project e k s a m i types default response
func (o *ListProjectEKSAMITypesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project e k s a m i types default response has a 2xx status code
func (o *ListProjectEKSAMITypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project e k s a m i types default response has a 3xx status code
func (o *ListProjectEKSAMITypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project e k s a m i types default response has a 4xx status code
func (o *ListProjectEKSAMITypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project e k s a m i types default response has a 5xx status code
func (o *ListProjectEKSAMITypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project e k s a m i types default response a status code equal to that given
func (o *ListProjectEKSAMITypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectEKSAMITypesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/amitypes][%d] listProjectEKSAMITypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectEKSAMITypesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/amitypes][%d] listProjectEKSAMITypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectEKSAMITypesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectEKSAMITypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
