package api

import (
	"github.com/go-chi/chi/v5"
)

func InitRoute(r chi.Router) {
	initTestHandler(r)
	initCaptchaRoute(r)
	initOauthRoute(r)
	InitUserRoute(r)
	initUserExtHandler(r)
	InitResourceRoute(r)
	InitRoleRoute(r)
	InitUser_with_roleRoute(r)
	InitUseroleRoute(r)
	InitRole_with_resource_idsRoute(r)
	InitRole_resource_operateRoute(r)
	InitRole_detailRoute(r)
	InitUser_with_menuRoute(r)
}
