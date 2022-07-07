// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
)

// BackupListReader is a Reader for the BackupList structure.
type BackupListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewBackupListUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewBackupListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupListOK creates a BackupListOK with default headers values
func NewBackupListOK() *BackupListOK {
	return &BackupListOK{}
}

/* BackupListOK describes a response with status code 200, with default header values.

search results matching criteria
*/
type BackupListOK struct {
	Payload models.Backups
}

func (o *BackupListOK) Error() string {
	return fmt.Sprintf("[GET /backups/][%d] backupListOK  %+v", 200, o.Payload)
}
func (o *BackupListOK) GetPayload() models.Backups {
	return o.Payload
}

func (o *BackupListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListBadRequest creates a BackupListBadRequest with default headers values
func NewBackupListBadRequest() *BackupListBadRequest {
	return &BackupListBadRequest{}
}

/* BackupListBadRequest describes a response with status code 400, with default header values.

bad input parameter
*/
type BackupListBadRequest struct {
	Payload *models.Error
}

func (o *BackupListBadRequest) Error() string {
	return fmt.Sprintf("[GET /backups/][%d] backupListBadRequest  %+v", 400, o.Payload)
}
func (o *BackupListBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *BackupListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListUnauthorized creates a BackupListUnauthorized with default headers values
func NewBackupListUnauthorized() *BackupListUnauthorized {
	return &BackupListUnauthorized{}
}

/* BackupListUnauthorized describes a response with status code 401, with default header values.

bad authentication
*/
type BackupListUnauthorized struct {
}

func (o *BackupListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /backups/][%d] backupListUnauthorized ", 401)
}

func (o *BackupListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBackupListForbidden creates a BackupListForbidden with default headers values
func NewBackupListForbidden() *BackupListForbidden {
	return &BackupListForbidden{}
}

/* BackupListForbidden describes a response with status code 403, with default header values.

bad permissions
*/
type BackupListForbidden struct {
}

func (o *BackupListForbidden) Error() string {
	return fmt.Sprintf("[GET /backups/][%d] backupListForbidden ", 403)
}

func (o *BackupListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBackupListUnprocessableEntity creates a BackupListUnprocessableEntity with default headers values
func NewBackupListUnprocessableEntity() *BackupListUnprocessableEntity {
	return &BackupListUnprocessableEntity{}
}

/* BackupListUnprocessableEntity describes a response with status code 422, with default header values.

bad validation
*/
type BackupListUnprocessableEntity struct {
	Payload *models.Error
}

func (o *BackupListUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /backups/][%d] backupListUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *BackupListUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *BackupListUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListServiceUnavailable creates a BackupListServiceUnavailable with default headers values
func NewBackupListServiceUnavailable() *BackupListServiceUnavailable {
	return &BackupListServiceUnavailable{}
}

/* BackupListServiceUnavailable describes a response with status code 503, with default header values.

internal server error
*/
type BackupListServiceUnavailable struct {
	Payload *models.Error
}

func (o *BackupListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /backups/][%d] backupListServiceUnavailable  %+v", 503, o.Payload)
}
func (o *BackupListServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *BackupListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
