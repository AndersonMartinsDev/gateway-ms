package main

import (
	"gateway-ms/cmd"
	"gateway-ms/internal/infrastructure/configuration"
	"gateway-ms/internal/infrastructure/grpc_client"
	"gateway-ms/internal/infrastructure/http_server"
	"log"
	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	configuration.LoadEnv()
	configuration.LoadLogger()
	configuration.LoadDatabase()

	router_compose := cmd.NewRouterCompose()

	conn, err := amqp.Dial(router_compose.RabbitMQURL)
	if err != nil {
		log.Fatalf("Falha ao conectar no RabbitMQ: %v", err)
	}
	webHookProcessorMSgrpcConn := grpc_client.GrcpConnection(router_compose.WebhookPrMsURL)
	agentModelMSgrpcConn := grpc_client.GrcpConnection(router_compose.AgentModelMsURL)

	userHandle := router_compose.HandlerUserConfiguration()
	authHandle := router_compose.HandlerAuthConfiguration()
	webhookHandle := router_compose.HandlerWebhookConfiguration(conn, webHookProcessorMSgrpcConn)
	agentModelHandler := router_compose.HandlerAgentModelConfiguration(agentModelMSgrpcConn)

	var routes []http_server.RouterInterface
	routes = append(routes, http_server.NewUserRoute(*userHandle))
	routes = append(routes, http_server.NewAuthRoute(*authHandle))
	routes = append(routes, http_server.NewWebHookRoute(*webhookHandle))
	routes = append(routes, http_server.NewAgentModelRoute(*agentModelHandler))
	slog.Info("Rotas HTTP registradas com sucesso!")

	defer conn.Close()
	defer webHookProcessorMSgrpcConn.Close()
	defer agentModelMSgrpcConn.Close()
	configuration.LoadServer(http_server.NewRouters(routes))

}
