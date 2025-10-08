package cmd

import (
	"gateway-ms/internal/application/handler"
	"gateway-ms/internal/application/service"
	"gateway-ms/internal/infrastructure/grpc_client"
	"gateway-ms/internal/infrastructure/rabbitmq"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type RouterCompose struct {
	WebhookPrMsURL  string
	AgentModelMsURL string
	WhastAppPFUrl   string
	RabbitMQURL     string
}

func NewRouterCompose() *RouterCompose {
	return &RouterCompose{
		WebhookPrMsURL:  os.Getenv("WEBHOOK_PROCESSOR_MS_URL"),
		AgentModelMsURL: os.Getenv("AGENT_MODEL_MS_URL"),
		WhastAppPFUrl:   os.Getenv("WHATSAPP_PF_URL"),
		RabbitMQURL:     os.Getenv("RABBITMQ_URL"),
	}
}

func (manager RouterCompose) HandlerWhatsappPFConfiguration(grpcConn *grpc.ClientConn) *handler.WhatsappPFHandler {
	service := service.NewWhatsAppPFService(grpc_client.NewWhatsappPFClient(grpcConn))
	return handler.NewWhatsappPFHandler(service)
}

func (manager RouterCompose) HandlerWebhookConfiguration(conn *amqp.Connection, grpcConn *grpc.ClientConn) *handler.WebhookHandler {
	publisher, err := rabbitmq.NewPublisher(conn)
	if err != nil {
		log.Fatalf("Falha ao criar publicador RabbitMQ: %v", err)
	}
	userService := service.NewWebhookService(publisher)
	return handler.NewWebhookHandler(userService)
}

func (manager RouterCompose) HandlerAuthConfiguration() *handler.AuthHandler {
	authService := service.NewAuthService()
	return handler.NewAuthHandler(authService)
}

func (manager RouterCompose) HandlerUserConfiguration() *handler.UserHandler {
	userService := service.NewUserService()
	return handler.NewUserHandler(userService)
}

func (manager RouterCompose) HandlerAgentModelConfiguration(grpcConn *grpc.ClientConn) *handler.AgentModelHandler {
	agentClient := grpc_client.NewAgentModelClient(grpcConn)
	modelClient := grpc_client.NewModelClient(grpcConn)
	agentModelService := service.NewAgentModelService(agentClient, modelClient)
	return handler.NewAgentModelHandler(agentModelService)
}
