// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewServiceExplainParams creates a new ServiceExplainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceExplainParams() *ServiceExplainParams {
	return &ServiceExplainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceExplainParamsWithTimeout creates a new ServiceExplainParams object
// with the ability to set a timeout on a request.
func NewServiceExplainParamsWithTimeout(timeout time.Duration) *ServiceExplainParams {
	return &ServiceExplainParams{
		timeout: timeout,
	}
}

// NewServiceExplainParamsWithContext creates a new ServiceExplainParams object
// with the ability to set a context for a request.
func NewServiceExplainParamsWithContext(ctx context.Context) *ServiceExplainParams {
	return &ServiceExplainParams{
		Context: ctx,
	}
}

// NewServiceExplainParamsWithHTTPClient creates a new ServiceExplainParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceExplainParamsWithHTTPClient(client *http.Client) *ServiceExplainParams {
	return &ServiceExplainParams{
		HTTPClient: client,
	}
}

/* ServiceExplainParams contains all the parameters to send to the API endpoint
   for the service explain operation.

   Typically these are written to a http.Request.
*/
type ServiceExplainParams struct {

	/* ServiceID.

	   service Resource ID
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service explain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceExplainParams) WithDefaults() *ServiceExplainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service explain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceExplainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service explain params
func (o *ServiceExplainParams) WithTimeout(timeout time.Duration) *ServiceExplainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service explain params
func (o *ServiceExplainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service explain params
func (o *ServiceExplainParams) WithContext(ctx context.Context) *ServiceExplainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service explain params
func (o *ServiceExplainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service explain params
func (o *ServiceExplainParams) WithHTTPClient(client *http.Client) *ServiceExplainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service explain params
func (o *ServiceExplainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the service explain params
func (o *ServiceExplainParams) WithServiceID(serviceID string) *ServiceExplainParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the service explain params
func (o *ServiceExplainParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceExplainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ServiceID
	if err := r.SetPathParam("ServiceID", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}