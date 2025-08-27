package http_server

import (
	"gateway-ms/internal/application/handler"
	model "gateway-ms/internal/infrastructure/commons/models/routes"
	"net/http"
)

type UserRoute struct {
	uri     string
	Handler handler.UserHandler
}

func NewUserRoute(userHandle handler.UserHandler) *UserRoute {
	return &UserRoute{
		uri:     "/user",
		Handler: userHandle,
	}
}

func (route UserRoute) getRoutersModel() []model.RouteModel {
	return []model.RouteModel{
		{
			URI:              route.uri + "/create",
			Method:           http.MethodPost,
			Func:             route.Handler.InsertNewUser,
			HasAuthenticated: true,
		},
	}
}
