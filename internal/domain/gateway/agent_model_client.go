package gateway

import (
	"context"

	"gateway-ms/proto"

	"google.golang.org/grpc"
)

// WebhookProcessorClient define a interface para o serviço de processamento de webhooks.
// A camada de aplicação usará essa interface, sem saber que a implementação é gRPC.
type AgentClient interface {
	CreateAgent(ctx context.Context, in *proto.CreateAgentRequest, opts ...grpc.CallOption) (*proto.CreateAgentResponse, error)
	UpdateAgent(ctx context.Context, in *proto.UpdateAgentRequest, opts ...grpc.CallOption) (*proto.UpdateAgentResponse, error)
	GetAgent(ctx context.Context, in *proto.GetAgentRequest, opts ...grpc.CallOption) (*proto.GetAgentResponse, error)
	ListAgents(ctx context.Context, in *proto.ListAgentsRequest, opts ...grpc.CallOption) (*proto.ListAgentsResponse, error)
	DeleteAgent(ctx context.Context, in *proto.DeleteAgentRequest, opts ...grpc.CallOption) (*proto.DeleteAgentResponse, error)
}
