// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
)

// ServiceArchiveReader is a Reader for the ServiceArchive structure.
type ServiceArchiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceArchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceArchiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceArchiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceArchiveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceArchiveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceArchiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceArchiveUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewServiceArchiveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceArchiveOK creates a ServiceArchiveOK with default headers values
func NewServiceArchiveOK() *ServiceArchiveOK {
	return &ServiceArchiveOK{}
}

/* ServiceArchiveOK describes a response with status code 200, with default header values.

service request to archive is sent
*/
type ServiceArchiveOK struct {
}

func (o *ServiceArchiveOK) Error() string {
	return fmt.Sprintf("[POST /services/{ServiceID}/archive][%d] serviceArchiveOK ", 200)
}

func (o *ServiceArchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceArchiveBadRequest creates a ServiceArchiveBadRequest with default headers values
func NewServiceArchiveBadRequest() *ServiceArchiveBadRequest {
	return &ServiceArchiveBadRequest{}
}

/* ServiceArchiveBadRequest describes a response with status code 400, with default header values.

invalid input
*/
type ServiceArchiveBadRequest struct {
	Payload *models.Error
}

func (o *ServiceArchiveBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/{ServiceID}/archive][%d] serviceArchiveBadRequest  %+v", 400, o.Payload)
}
func (o *ServiceArchiveBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceArchiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceArchiveUnauthorized creates a ServiceArchiveUnauthorized with default headers values
func NewServiceArchiveUnauthorized() *ServiceArchiveUnauthorized {
	return &ServiceArchiveUnauthorized{}
}

/* ServiceArchiveUnauthorized describes a response with status code 401, with default header values.

bad authentication
*/
type ServiceArchiveUnauthorized struct {
}

func (o *ServiceArchiveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /services/{ServiceID}/archive][%d] serviceArchiveUnauthorized ", 401)
}

func (o *ServiceArchiveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceArchiveForbidden creates a ServiceArchiveForbidden with default headers values
func NewServiceArchiveForbidden() *ServiceArchiveForbidden {
	return &ServiceArchiveForbidden{}
}

/* ServiceArchiveForbidden describes a response with status code 403, with default header values.

bad permissions
*/
type ServiceArchiveForbidden struct {
}

func (o *ServiceArchiveForbidden) Error() string {
	return fmt.Sprintf("[POST /services/{ServiceID}/archive][%d] serviceArchiveForbidden ", 403)
}

func (o *ServiceArchiveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceArchiveNotFound creates a ServiceArchiveNotFound with default headers values
func NewServiceArchiveNotFound() *ServiceArchiveNotFound {
	return &ServiceArchiveNotFound{}
}

/* ServiceArchiveNotFound describes a response with status code 404, with default header values.

service not found
*/
type ServiceArchiveNotFound struct {
	Payload *models.Error
}

func (o *ServiceArchiveNotFound) Error() string {
	return fmt.Sprintf("[POST /services/{ServiceID}/archive][%d] serviceArchiveNotFound  %+v", 404, o.Payload)
}
func (o *ServiceArchiveNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceArchiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceArchiveUnprocessableEntity creates a ServiceArchiveUnprocessableEntity with default headers values
func NewServiceArchiveUnprocessableEntity() *ServiceArchiveUnprocessableEntity {
	return &ServiceArchiveUnprocessableEntity{}
}

/* ServiceArchiveUnprocessableEntity describes a response with status code 422, with default header values.

bad validation
*/
type ServiceArchiveUnprocessableEntity struct {
}

func (o *ServiceArchiveUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /services/{ServiceID}/archive][%d] serviceArchiveUnprocessableEntity ", 422)
}

func (o *ServiceArchiveUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceArchiveServiceUnavailable creates a ServiceArchiveServiceUnavailable with default headers values
func NewServiceArchiveServiceUnavailable() *ServiceArchiveServiceUnavailable {
	return &ServiceArchiveServiceUnavailable{}
}

/* ServiceArchiveServiceUnavailable describes a response with status code 503, with default header values.

internal service error
*/
type ServiceArchiveServiceUnavailable struct {
	Payload *models.Error
}

func (o *ServiceArchiveServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /services/{ServiceID}/archive][%d] serviceArchiveServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ServiceArchiveServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceArchiveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}