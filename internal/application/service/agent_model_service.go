package service

import (
	"context"
	"gateway-ms/internal/domain/model"
	proto "gateway-ms/proto"
)

type AgentModelService interface {
	CreateAgentAI(ctx context.Context, agent model.AIAgent) (string, error)
	DeleteAgent(ctx context.Context, agentID, userUUID string) (string, error)
	GetAgent(ctx context.Context, agentID, userUUID string) (*proto.Agent, error)
	ListAgents(ctx context.Context, userUUID string) ([]*proto.Agent, error)
	UpdateAgentAI(ctx context.Context, agent model.AIAgent) (string, error)
	CreatePerfilModel(ctx context.Context, model model.ModelIA) (string, error)
	GetPerfilModels(ctx context.Context) ([]*proto.AIModel, error)
}

// AgentModelServiceImpl local orquestra as chamadas para os microserviços de agentes e modelos.
type AgentModelServiceImpl struct {
	agentClient proto.AIAgentServiceClient
	modelClient proto.AIModelServiceClient
}

// NewAgentService cria uma nova instância do serviço local.
func NewAgentModelService(agentClient proto.AIAgentServiceClient, modelClient proto.AIModelServiceClient) AgentModelService {
	return &AgentModelServiceImpl{
		agentClient: agentClient,
		modelClient: modelClient,
	}
}

func (s *AgentModelServiceImpl) ListAgents(ctx context.Context, userUUID string) ([]*proto.Agent, error) {
	req := &proto.ListAgentsRequest{UuidUser: userUUID}
	res, err := s.agentClient.ListAgents(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Agents, nil
}
func (s *AgentModelServiceImpl) GetAgent(ctx context.Context, agentID, userUUID string) (*proto.Agent, error) {
	req := &proto.GetAgentRequest{
		Id:       agentID,
		UuidUser: userUUID,
	}
	res, err := s.agentClient.GetAgent(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Agent, nil
}
func (s *AgentModelServiceImpl) UpdateAgentAI(ctx context.Context, agent model.AIAgent) (string, error) {
	req := &proto.UpdateAgentRequest{
		Id:                 agent.Id,
		Name:               agent.Name,
		CompanyName:        agent.CompanyName,
		CompanyDescription: agent.CompanyDescription,
		ModelId:            uint32(agent.ModelId),
		BehaviourIa:        agent.BehaviourIa,
		CompanyUrl:         agent.CompanyUrl,
		UuidUser:           agent.UUID_user,
	}
	res, err := s.agentClient.UpdateAgent(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetStatus(), nil
}
func (s *AgentModelServiceImpl) CreateAgentAI(ctx context.Context, agent model.AIAgent) (string, error) {
	req := &proto.CreateAgentRequest{
		Name:               agent.Name,
		CompanyName:        agent.CompanyName,
		CompanyDescription: agent.CompanyDescription,
		ModelId:            uint32(agent.ModelId),
		BehaviourIa:        agent.BehaviourIa,
		CompanyUrl:         agent.CompanyUrl,
		UuidUser:           agent.UUID_user,
	}
	res, err := s.agentClient.CreateAgent(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetStatus(), nil
}
func (s *AgentModelServiceImpl) DeleteAgent(ctx context.Context, agentID, userUUID string) (string, error) {
	req := &proto.DeleteAgentRequest{
		Id:       agentID,
		UuidUser: userUUID,
	}
	res, err := s.agentClient.DeleteAgent(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetStatus(), nil
}

// CreatePerfilModel implements AgentModelService.
func (s *AgentModelServiceImpl) CreatePerfilModel(ctx context.Context, model model.ModelIA) (string, error) {
	req := &proto.CreateAIModelRequest{
		PerfilName:   model.PerFilName,
		Instructions: model.Instructions,
	}
	res, err := s.modelClient.CreatePerfilModel(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetStatus(), nil
}

// GetPerfilModels implements AgentModelService.
func (s *AgentModelServiceImpl) GetPerfilModels(ctx context.Context) ([]*proto.AIModel, error) {
	req := &proto.GetPerfilModelsRequest{}

	res, err := s.modelClient.GetPerfilModels(ctx, req)
	if err != nil {
		return []*proto.AIModel{}, err
	}
	return res.Models, nil
}
