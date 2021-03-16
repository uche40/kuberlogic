// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
)

// DatabaseRestoreHandlerFunc turns a function with the right signature into a database restore handler
type DatabaseRestoreHandlerFunc func(DatabaseRestoreParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DatabaseRestoreHandlerFunc) Handle(params DatabaseRestoreParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DatabaseRestoreHandler interface for that can handle valid database restore params
type DatabaseRestoreHandler interface {
	Handle(DatabaseRestoreParams, *models.Principal) middleware.Responder
}

// NewDatabaseRestore creates a new http.Handler for the database restore operation
func NewDatabaseRestore(ctx *middleware.Context, handler DatabaseRestoreHandler) *DatabaseRestore {
	return &DatabaseRestore{Context: ctx, Handler: handler}
}

/*DatabaseRestore swagger:route POST /services/{ServiceID}/restores service databaseRestore

DatabaseRestore database restore API

*/
type DatabaseRestore struct {
	Context *middleware.Context
	Handler DatabaseRestoreHandler
}

func (o *DatabaseRestore) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDatabaseRestoreParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DatabaseRestoreBody database restore body
//
// swagger:model DatabaseRestoreBody
type DatabaseRestoreBody struct {

	// database
	// Required: true
	Database *string `json:"database"`

	// key
	// Required: true
	Key *string `json:"key"`
}

// Validate validates this database restore body
func (o *DatabaseRestoreBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DatabaseRestoreBody) validateDatabase(formats strfmt.Registry) error {

	if err := validate.Required("restoreItem"+"."+"database", "body", o.Database); err != nil {
		return err
	}

	return nil
}

func (o *DatabaseRestoreBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("restoreItem"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DatabaseRestoreBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DatabaseRestoreBody) UnmarshalBinary(b []byte) error {
	var res DatabaseRestoreBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
