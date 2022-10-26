// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
)

// ServiceExplainHandlerFunc turns a function with the right signature into a service explain handler
type ServiceExplainHandlerFunc func(ServiceExplainParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceExplainHandlerFunc) Handle(params ServiceExplainParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ServiceExplainHandler interface for that can handle valid service explain params
type ServiceExplainHandler interface {
	Handle(ServiceExplainParams, *models.Principal) middleware.Responder
}

// NewServiceExplain creates a new http.Handler for the service explain operation
func NewServiceExplain(ctx *middleware.Context, handler ServiceExplainHandler) *ServiceExplain {
	return &ServiceExplain{Context: ctx, Handler: handler}
}

/* ServiceExplain swagger:route GET /services/{ServiceID}/explain service serviceExplain

explain status of service

Explain status of service

*/
type ServiceExplain struct {
	Context *middleware.Context
	Handler ServiceExplainHandler
}

func (o *ServiceExplain) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewServiceExplainParams()
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