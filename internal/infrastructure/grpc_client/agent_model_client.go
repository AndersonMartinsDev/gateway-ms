package grpc_client

import (
	"context"
	"gateway-ms/internal/domain/gateway"
	"gateway-ms/proto"

	"google.golang.org/grpc"
)

// WebhookClient implementa a interface domain/gateway.WebhookProcessorClient.
type AgentModelClient struct {
	client proto.AIAgentServiceClient
}

// NewWebhookClient cria uma nova instância do cliente gRPC.
func NewAgentModelClient(conn *grpc.ClientConn) gateway.AgentClient {
	return &AgentModelClient{
		client: proto.NewAIAgentServiceClient(conn),
	}
}

// CreateAgent implements gateway.AgentModelClient.
func (a *AgentModelClient) CreateAgent(ctx context.Context, in *proto.CreateAgentRequest, opts ...grpc.CallOption) (*proto.CreateAgentResponse, error) {
	return a.client.CreateAgent(ctx, in, opts...)
}

// DeleteAgent implements gateway.AgentModelClient.
func (a *AgentModelClient) DeleteAgent(ctx context.Context, in *proto.DeleteAgentRequest, opts ...grpc.CallOption) (*proto.DeleteAgentResponse, error) {
	return a.client.DeleteAgent(ctx, in, opts...)
}

// GetAgent implements gateway.AgentModelClient.
func (a *AgentModelClient) GetAgent(ctx context.Context, in *proto.GetAgentRequest, opts ...grpc.CallOption) (*proto.GetAgentResponse, error) {
	return a.client.GetAgent(ctx, in, opts...)
}

// ListAgents implements gateway.AgentModelClient.
func (a *AgentModelClient) ListAgents(ctx context.Context, in *proto.ListAgentsRequest, opts ...grpc.CallOption) (*proto.ListAgentsResponse, error) {
	return a.client.ListAgents(ctx, in, opts...)
}

// UpdateAgent implements gateway.AgentModelClient.
func (a *AgentModelClient) UpdateAgent(ctx context.Context, in *proto.UpdateAgentRequest, opts ...grpc.CallOption) (*proto.UpdateAgentResponse, error) {
	return a.client.UpdateAgent(ctx, in, opts...)
}
