package grpc_client

import (
	"context"
	"gateway-ms/internal/domain/gateway"
	"gateway-ms/proto"

	"google.golang.org/grpc"
)

// ModelClient implementa a interface domain/gateway.AIModelServiceClient.
type ModelClient struct {
	client proto.AIModelServiceClient
}

// NewModelClient cria uma nova instância do cliente gRPC.
func NewModelClient(conn *grpc.ClientConn) gateway.ModelClient {
	return &ModelClient{
		client: proto.NewAIModelServiceClient(conn),
	}
}

// CreateAIModel implements gateway.ModelClient.
func (m *ModelClient) CreatePerfilModel(ctx context.Context, in *proto.CreateAIModelRequest, opts ...grpc.CallOption) (*proto.GetAIModelsResponse, error) {
	return m.client.CreatePerfilModel(ctx, in, opts...)
}

// GetAIModels implements gateway.ModelClient.
func (m *ModelClient) GetPerfilModels(ctx context.Context, in *proto.GetPerfilModelsRequest, opts ...grpc.CallOption) (*proto.GetPerfilModelsResponse, error) {
	return m.client.GetPerfilModels(ctx, in, opts...)
}
