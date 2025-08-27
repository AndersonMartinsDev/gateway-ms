package http_server

import (
	"gateway-ms/internal/application/handler"
	model "gateway-ms/internal/infrastructure/commons/models/routes"
	"net/http"
)

type AuthRouter struct {
	uri     string
	Handler handler.AuthHandler
}

func NewAuthRoute(handler handler.AuthHandler) *AuthRouter {
	return &AuthRouter{
		uri:     "/auth",
		Handler: handler,
	}
}

func (route AuthRouter) getRoutersModel() []model.RouteModel {
	return []model.RouteModel{
		{
			URI:              route.uri + "/login",
			Method:           http.MethodPost,
			Func:             route.Handler.Login,
			HasAuthenticated: false,
		},
	}
}
