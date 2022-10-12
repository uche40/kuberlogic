// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
)

// ServiceArchiveHandlerFunc turns a function with the right signature into a service archive handler
type ServiceArchiveHandlerFunc func(ServiceArchiveParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceArchiveHandlerFunc) Handle(params ServiceArchiveParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ServiceArchiveHandler interface for that can handle valid service archive params
type ServiceArchiveHandler interface {
	Handle(ServiceArchiveParams, *models.Principal) middleware.Responder
}

// NewServiceArchive creates a new http.Handler for the service archive operation
func NewServiceArchive(ctx *middleware.Context, handler ServiceArchiveHandler) *ServiceArchive {
	return &ServiceArchive{Context: ctx, Handler: handler}
}

/* ServiceArchive swagger:route POST /services/{ServiceID}/archive service serviceArchive

archive service

archive service (for example, if user subscription got cancelled)

*/
type ServiceArchive struct {
	Context *middleware.Context
	Handler ServiceArchiveHandler
}

func (o *ServiceArchive) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewServiceArchiveParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
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
