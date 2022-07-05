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

// ServiceEditReader is a Reader for the ServiceEdit structure.
type ServiceEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceEditUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewServiceEditServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceEditOK creates a ServiceEditOK with default headers values
func NewServiceEditOK() *ServiceEditOK {
	return &ServiceEditOK{}
}

/* ServiceEditOK describes a response with status code 200, with default header values.

item edited
*/
type ServiceEditOK struct {
	Payload *models.Service
}

func (o *ServiceEditOK) Error() string {
	return fmt.Sprintf("[PATCH /services/{ServiceID}/][%d] serviceEditOK  %+v", 200, o.Payload)
}
func (o *ServiceEditOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *ServiceEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceEditBadRequest creates a ServiceEditBadRequest with default headers values
func NewServiceEditBadRequest() *ServiceEditBadRequest {
	return &ServiceEditBadRequest{}
}

/* ServiceEditBadRequest describes a response with status code 400, with default header values.

invalid input, object invalid
*/
type ServiceEditBadRequest struct {
	Payload *models.Error
}

func (o *ServiceEditBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /services/{ServiceID}/][%d] serviceEditBadRequest  %+v", 400, o.Payload)
}
func (o *ServiceEditBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceEditUnauthorized creates a ServiceEditUnauthorized with default headers values
func NewServiceEditUnauthorized() *ServiceEditUnauthorized {
	return &ServiceEditUnauthorized{}
}

/* ServiceEditUnauthorized describes a response with status code 401, with default header values.

bad authentication
*/
type ServiceEditUnauthorized struct {
}

func (o *ServiceEditUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /services/{ServiceID}/][%d] serviceEditUnauthorized ", 401)
}

func (o *ServiceEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceEditForbidden creates a ServiceEditForbidden with default headers values
func NewServiceEditForbidden() *ServiceEditForbidden {
	return &ServiceEditForbidden{}
}

/* ServiceEditForbidden describes a response with status code 403, with default header values.

bad permissions
*/
type ServiceEditForbidden struct {
}

func (o *ServiceEditForbidden) Error() string {
	return fmt.Sprintf("[PATCH /services/{ServiceID}/][%d] serviceEditForbidden ", 403)
}

func (o *ServiceEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceEditNotFound creates a ServiceEditNotFound with default headers values
func NewServiceEditNotFound() *ServiceEditNotFound {
	return &ServiceEditNotFound{}
}

/* ServiceEditNotFound describes a response with status code 404, with default header values.

item not found
*/
type ServiceEditNotFound struct {
	Payload *models.Error
}

func (o *ServiceEditNotFound) Error() string {
	return fmt.Sprintf("[PATCH /services/{ServiceID}/][%d] serviceEditNotFound  %+v", 404, o.Payload)
}
func (o *ServiceEditNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceEditUnprocessableEntity creates a ServiceEditUnprocessableEntity with default headers values
func NewServiceEditUnprocessableEntity() *ServiceEditUnprocessableEntity {
	return &ServiceEditUnprocessableEntity{}
}

/* ServiceEditUnprocessableEntity describes a response with status code 422, with default header values.

bad validation
*/
type ServiceEditUnprocessableEntity struct {
	Payload *models.Error
}

func (o *ServiceEditUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /services/{ServiceID}/][%d] serviceEditUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ServiceEditUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceEditUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceEditServiceUnavailable creates a ServiceEditServiceUnavailable with default headers values
func NewServiceEditServiceUnavailable() *ServiceEditServiceUnavailable {
	return &ServiceEditServiceUnavailable{}
}

/* ServiceEditServiceUnavailable describes a response with status code 503, with default header values.

internal server error
*/
type ServiceEditServiceUnavailable struct {
	Payload *models.Error
}

func (o *ServiceEditServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /services/{ServiceID}/][%d] serviceEditServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ServiceEditServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceEditServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}