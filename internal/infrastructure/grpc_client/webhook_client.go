package grpc_client

import (
	"context"
	"gateway-ms/internal/domain/gateway"
	"gateway-ms/proto"

	"google.golang.org/grpc"
)

// WebhookClient implementa a interface domain/gateway.WebhookProcessorClient.
type WebhookClient struct {
	client proto.WebhookProcessorServiceClient
}

// NewWebhookClient cria uma nova instância do cliente gRPC.
func NewWebhookClient(conn *grpc.ClientConn) gateway.WebhookProcessorClient {
	return &WebhookClient{
		client: proto.NewWebhookProcessorServiceClient(conn),
	}
}

// ProcessWebhook é a implementação do método da interface.
func (c *WebhookClient) ProcessWebhook(ctx context.Context, req *proto.ProcessWebhookRequest) (*proto.ProcessWebhookResponse, error) {
	return c.client.ProcessWebhook(ctx, req)
}
