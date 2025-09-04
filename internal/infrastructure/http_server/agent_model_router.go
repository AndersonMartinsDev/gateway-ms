package http_server

import (
	"gateway-ms/internal/application/handler"
	model "gateway-ms/internal/infrastructure/commons/models/routes"
	"net/http"
)

type AgentModelRoute struct {
	uri     string
	Handler handler.AgentModelHandler
}

func NewAgentModelRoute(handler handler.AgentModelHandler) *AgentModelRoute {
	return &AgentModelRoute{
		uri:     "/agent-model",
		Handler: handler,
	}
}

func (route AgentModelRoute) getRoutersModel() []model.RouteModel {
	return []model.RouteModel{
		{
			URI:              route.uri + "/create-perfil",
			Method:           http.MethodPost,
			HasAuthenticated: true,
			Func:             route.Handler.CreatePerfilModel,
		},
		{
			URI:              route.uri + "/get/models",
			Method:           http.MethodGet,
			HasAuthenticated: true,
			Func:             route.Handler.GetModels,
		},
		{
			URI:              route.uri + "/create",
			Method:           http.MethodPost,
			HasAuthenticated: true,
			Func:             route.Handler.CreateAgentAI,
		},
		{
			URI:              route.uri,
			Method:           http.MethodPut,
			HasAuthenticated: true,
			Func:             route.Handler.UpdateAgentAI,
		},
		{
			URI:              route.uri,
			Method:           http.MethodGet,
			HasAuthenticated: true,
			Func:             route.Handler.ListAgentIA,
		},
		{
			URI:              route.uri + "/get",
			Method:           http.MethodGet,
			HasAuthenticated: true,
			Func:             route.Handler.GetAgentIa,
		},
		{
			URI:              route.uri + "/models/behaviour",
			Method:           http.MethodGet,
			HasAuthenticated: true,
			Func:             route.Handler.GetBehaviorAgentIa,
		},
		{
			URI:              route.uri,
			Method:           http.MethodDelete,
			HasAuthenticated: true,
			Func:             route.Handler.RemoveAgentIa,
		},
	}
}
