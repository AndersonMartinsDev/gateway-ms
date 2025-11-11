package gateway

import (
	"context"
	"gateway-ms/proto"
)

// WebhookProcessorClient define a interface para o serviço de processamento de webhooks.
// A camada de aplicação usará essa interface, sem saber que a implementação é gRPC.
type WebhookProcessorClient interface {
	ProcessWebhook(ctx context.Context, req *proto.ProcessWebhookRequest) (*proto.ProcessWebhookResponse, error)
}
