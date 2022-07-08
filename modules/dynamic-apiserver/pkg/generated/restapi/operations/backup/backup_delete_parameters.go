// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewBackupDeleteParams creates a new BackupDeleteParams object
//
// There are no default values defined in the spec.
func NewBackupDeleteParams() BackupDeleteParams {

	return BackupDeleteParams{}
}

// BackupDeleteParams contains all the bound params for the backup delete operation
// typically these are obtained from a http.Request
//
// swagger:parameters backupDelete
type BackupDeleteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*backup Resource ID
	  Required: true
	  Max Length: 63
	  Min Length: 3
	  Pattern: [a-z0-9]([-a-z0-9]*[a-z0-9])?
	  In: path
	*/
	BackupID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewBackupDeleteParams() beforehand.
func (o *BackupDeleteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBackupID, rhkBackupID, _ := route.Params.GetOK("BackupID")
	if err := o.bindBackupID(rBackupID, rhkBackupID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBackupID binds and validates parameter BackupID from path.
func (o *BackupDeleteParams) bindBackupID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.BackupID = raw

	if err := o.validateBackupID(formats); err != nil {
		return err
	}

	return nil
}

// validateBackupID carries on validations for parameter BackupID
func (o *BackupDeleteParams) validateBackupID(formats strfmt.Registry) error {

	if err := validate.MinLength("BackupID", "path", o.BackupID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("BackupID", "path", o.BackupID, 63); err != nil {
		return err
	}

	if err := validate.Pattern("BackupID", "path", o.BackupID, `[a-z0-9]([-a-z0-9]*[a-z0-9])?`); err != nil {
		return err
	}

	return nil
}