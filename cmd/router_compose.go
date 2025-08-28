package cmd

import (
	"gateway-ms/internal/application/handler"
	"gateway-ms/internal/application/service"
	"gateway-ms/internal/infrastructure/grpc_client"

	"google.golang.org/grpc"
)

type RouterCompose struct {
	WebhookPrMsURL string
}

func NewRouterCompose() *RouterCompose {
	return &RouterCompose{
		WebhookPrMsURL: "localhost:50051",
	}
}

func (manager RouterCompose) HandlerWebhookConfiguration(grpcConn *grpc.ClientConn) *handler.WebhookHandler {
	webhookClient := grpc_client.NewWebhookClient(grpcConn)
	userService := service.NewWebhookService(webhookClient)
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
