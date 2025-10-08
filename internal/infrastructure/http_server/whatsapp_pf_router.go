package http_server

import (
	"gateway-ms/internal/application/handler"
	model "gateway-ms/internal/infrastructure/commons/models/routes"
	"net/http"
)

type WhatsAppPFRoute struct {
	handler handler.WhatsappPFHandler
}

func NewWhatsAppPFRoute(handler handler.WhatsappPFHandler) *WhatsAppPFRoute {
	return &WhatsAppPFRoute{
		handler: handler,
	}
}

func (route WhatsAppPFRoute) getRoutersModel() []model.RouteModel {
	return []model.RouteModel{
		{
			URI:              "/whatsapp/qrcode",
			Method:           http.MethodGet,
			Func:             route.handler.HandleQrCode,
			HasAuthenticated: true,
		},
	}
}
