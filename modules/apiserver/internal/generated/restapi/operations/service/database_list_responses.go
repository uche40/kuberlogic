// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
)

// DatabaseListOKCode is the HTTP code returned for type DatabaseListOK
const DatabaseListOKCode int = 200

/*DatabaseListOK search results matching criteria

swagger:response databaseListOK
*/
type DatabaseListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Database `json:"body,omitempty"`
}

// NewDatabaseListOK creates DatabaseListOK with default headers values
func NewDatabaseListOK() *DatabaseListOK {

	return &DatabaseListOK{}
}

// WithPayload adds the payload to the database list o k response
func (o *DatabaseListOK) WithPayload(payload []*models.Database) *DatabaseListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the database list o k response
func (o *DatabaseListOK) SetPayload(payload []*models.Database) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DatabaseListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Database, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DatabaseListBadRequestCode is the HTTP code returned for type DatabaseListBadRequest
const DatabaseListBadRequestCode int = 400

/*DatabaseListBadRequest invalid input, object invalid

swagger:response databaseListBadRequest
*/
type DatabaseListBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDatabaseListBadRequest creates DatabaseListBadRequest with default headers values
func NewDatabaseListBadRequest() *DatabaseListBadRequest {

	return &DatabaseListBadRequest{}
}

// WithPayload adds the payload to the database list bad request response
func (o *DatabaseListBadRequest) WithPayload(payload *models.Error) *DatabaseListBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the database list bad request response
func (o *DatabaseListBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DatabaseListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DatabaseListUnauthorizedCode is the HTTP code returned for type DatabaseListUnauthorized
const DatabaseListUnauthorizedCode int = 401

/*DatabaseListUnauthorized bad authentication

swagger:response databaseListUnauthorized
*/
type DatabaseListUnauthorized struct {
}

// NewDatabaseListUnauthorized creates DatabaseListUnauthorized with default headers values
func NewDatabaseListUnauthorized() *DatabaseListUnauthorized {

	return &DatabaseListUnauthorized{}
}

// WriteResponse to the client
func (o *DatabaseListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DatabaseListForbiddenCode is the HTTP code returned for type DatabaseListForbidden
const DatabaseListForbiddenCode int = 403

/*DatabaseListForbidden bad permissions

swagger:response databaseListForbidden
*/
type DatabaseListForbidden struct {
}

// NewDatabaseListForbidden creates DatabaseListForbidden with default headers values
func NewDatabaseListForbidden() *DatabaseListForbidden {

	return &DatabaseListForbidden{}
}

// WriteResponse to the client
func (o *DatabaseListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DatabaseListServiceUnavailableCode is the HTTP code returned for type DatabaseListServiceUnavailable
const DatabaseListServiceUnavailableCode int = 503

/*DatabaseListServiceUnavailable internal server error

swagger:response databaseListServiceUnavailable
*/
type DatabaseListServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDatabaseListServiceUnavailable creates DatabaseListServiceUnavailable with default headers values
func NewDatabaseListServiceUnavailable() *DatabaseListServiceUnavailable {

	return &DatabaseListServiceUnavailable{}
}

// WithPayload adds the payload to the database list service unavailable response
func (o *DatabaseListServiceUnavailable) WithPayload(payload *models.Error) *DatabaseListServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the database list service unavailable response
func (o *DatabaseListServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DatabaseListServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}