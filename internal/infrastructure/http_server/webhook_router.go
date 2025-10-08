package http_server

import (
	"gateway-ms/internal/application/handler"
	model "gateway-ms/internal/infrastructure/commons/models/routes"
	"net/http"
)

type WebHookRoute struct {
	Handler handler.WebhookHandler
}

func NewWebHookRoute(userHandle handler.WebhookHandler) *WebHookRoute {
	return &WebHookRoute{
		Handler: userHandle,
	}
}

func (route WebHookRoute) getRoutersModel() []model.RouteModel {
	return []model.RouteModel{
		{
			URI:              "/whatsapp-pf",
			Method:           http.MethodPost,
			Func:             route.Handler.HandleWebhook,
			HasAuthenticated: false,
		},
	}
}
