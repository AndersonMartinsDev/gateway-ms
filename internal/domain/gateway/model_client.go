package gateway

import (
	"context"

	"gateway-ms/proto"

	"google.golang.org/grpc"
)

// WebhookProcessorClient define a interface para o serviço de processamento de webhooks.
// A camada de aplicação usará essa interface, sem saber que a implementação é gRPC.
type ModelClient interface {
	CreatePerfilModel(ctx context.Context, in *proto.CreateAIModelRequest, opts ...grpc.CallOption) (*proto.GetAIModelsResponse, error)
	GetPerfilModels(ctx context.Context, in *proto.GetPerfilModelsRequest, opts ...grpc.CallOption) (*proto.GetPerfilModelsResponse, error)
}
