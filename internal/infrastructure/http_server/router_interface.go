package http_server

import model "gateway-ms/internal/infrastructure/commons/models/routes"

type RouterInterface interface {
	getRoutersModel() []model.RouteModel
}
