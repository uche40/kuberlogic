// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
)

// NewBackupAddParams creates a new BackupAddParams object
//
// There are no default values defined in the spec.
func NewBackupAddParams() BackupAddParams {

	return BackupAddParams{}
}

// BackupAddParams contains all the bound params for the backup add operation
// typically these are obtained from a http.Request
//
// swagger:parameters backupAdd
type BackupAddParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*backup item
	  Required: true
	  In: body
	*/
	BackupItem *models.Backup
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewBackupAddParams() beforehand.
func (o *BackupAddParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Backup
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("backupItem", "body", ""))
			} else {
				res = append(res, errors.NewParseError("backupItem", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.BackupItem = &body
			}
		}
	} else {
		res = append(res, errors.Required("backupItem", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
