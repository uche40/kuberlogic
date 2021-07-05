// Code generated by go-swagger; DO NOT EDIT.
package service

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/security"
	"github.com/kuberlogic/operator/modules/apiserver/internal/logging/posthog"
	"github.com/kuberlogic/operator/modules/apiserver/util"
)

func UserEditWrapper(srv Service, next UserEditHandlerFunc) (fn UserEditHandlerFunc) {
	return func(params UserEditParams, principal *models.Principal) middleware.Responder {

		log := srv.GetLogger()

		// check ServiceID param
		ns, name, err := util.SplitID(params.ServiceID)
		if err != nil {
			msg := "incorrect service id"
			log.Errorw(msg, "serviceId", params.ServiceID, "error", err)
			return NewUserEditBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		}

		// check auth
		authProvider := srv.GetAuthProvider()
		if authorized, err := authProvider.Authorize(principal, security.UserEditPermission, params.ServiceID); err != nil {
			msg := "auth bad request"
			log.Errorw(msg, "permission", security.UserEditPermission, "serviceId", params.ServiceID, "error", err)
			return NewUserEditBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		} else if !authorized {
			log.Errorw("auth forbidden", "permission", security.UserEditPermission, "serviceId", params.ServiceID)
			return NewUserEditForbidden()
		}

		// cluster should exists
		service, found, err := srv.LookupService(ns, name)
		if !found {
			msg := "service not found"
			log.Errorw(msg, "error", err)
			return NewUserEditBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		} else if err != nil {
			msg := "error getting service"
			log.Errorw(msg, "error", err)
			return NewUserEditServiceUnavailable().WithPayload(&models.Error{
				Message: msg,
			})
		}

		params.HTTPRequest = params.HTTPRequest.WithContext(
			context.WithValue(params.HTTPRequest.Context(), "service", service))

		// enqueue data to posthog
		posthogMsg := posthog.NewMessage("user-edit")
		posthogMsg.With("service-id", params.ServiceID)
		posthogMsg.With("user", params.User.Name)
		if perr := posthogMsg.Create(); perr != nil {
			msg := "could not enqueue posthog message"
			log.Errorw(msg, "error", perr)
			return NewUserEditServiceUnavailable().WithPayload(&models.Error{
				Message: msg,
			})
		}

		return next(params, principal)
	}
}