package service

import (
	"context"
	"gateway-ms/internal/domain/gateway"
	"gateway-ms/proto"
)

// WebhookService orquestra o processamento do webhook.
type WebhookService struct {
	client gateway.WebhookProcessorClient
}

// NewWebhookService cria uma nova instância de WebhookService.
func NewWebhookService(client gateway.WebhookProcessorClient) *WebhookService {
	return &WebhookService{
		client: client,
	}
}

// Process recebe o request e delega a chamada para o cliente.
func (s *WebhookService) Process(ctx context.Context, req *proto.ProcessWebhookRequest) (*proto.ProcessWebhookResponse, error) {
	return s.client.ProcessWebhook(ctx, req)
}
