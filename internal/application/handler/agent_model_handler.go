package handler

import (
	"context"
	"fmt"
	"gateway-ms/internal/application/service"
	"gateway-ms/internal/domain/model"
	"gateway-ms/internal/infrastructure/request"
	"gateway-ms/internal/infrastructure/response"
	"net/http"
	"time"
)

// UserHandler é o handler HTTP para as requisições de usuário.
type AgentModelHandler struct {
	service service.AgentModelService
}

// NewUserHandler cria uma nova instância de UserHandler.
func NewAgentModelHandler(s service.AgentModelService) *AgentModelHandler {
	return &AgentModelHandler{
		service: s,
	}
}

func (handler AgentModelHandler) CreatePerfilModel(w http.ResponseWriter, r *http.Request) {
	var modelIA model.ModelIA
	if erro := request.Serialization(r.Body, &modelIA); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, erro := handler.service.CreatePerfilModel(ctx, modelIA)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.Response(w, http.StatusCreated, result, nil)

}

func (handler AgentModelHandler) GetModels(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	models, erro := handler.service.GetPerfilModels(ctx)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.Response(w, http.StatusCreated, models, nil)
}

func (handler AgentModelHandler) GetBehaviorAgentIa(w http.ResponseWriter, r *http.Request) {
	behavioural, err := handler.service.GetBehaviorAgentIa(context.Background())
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, behavioural)
}

func (handler AgentModelHandler) CreateAgentAI(w http.ResponseWriter, r *http.Request) {
	var agent model.AIAgent
	if erro := request.Serialization(r.Body, &agent); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	agent.UUID_user = r.Header.Get("user_uuid")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := handler.service.CreateAgentAI(ctx, agent)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, fmt.Errorf("has a error to create agent: %v", err))
		return
	}
	response.JSON(w, http.StatusOK, res)

}

func (handler AgentModelHandler) UpdateAgentAI(w http.ResponseWriter, r *http.Request) {
	var agent model.AIAgent
	if erro := request.Serialization(r.Body, &agent); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	agent.UUID_user = r.Header.Get("user_uuid")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := handler.service.UpdateAgentAI(ctx, agent)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, fmt.Errorf("has a error to list agents: %v", err))
		return
	}
	response.JSON(w, http.StatusOK, res)
}

func (handler AgentModelHandler) ListAgentIA(w http.ResponseWriter, r *http.Request) {
	// Exemplo: Obter o UUID do usuário de um token de autenticação ou parâmetro.
	userUUID := r.Header.Get("user_uuid")
	if userUUID == "" {
		http.Error(w, "user_uuid é obrigatório", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := handler.service.ListAgents(ctx, userUUID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, fmt.Errorf("has a error to list agents: %v", err))
		return
	}
	response.JSON(w, http.StatusOK, res)
}

func (handler AgentModelHandler) GetAgentIa(w http.ResponseWriter, r *http.Request) {
	userUUID := r.Header.Get("user_uuid")
	agent_id := r.URL.Query().Get("agent_id")
	if userUUID == "" {
		response.Erro(w, http.StatusBadRequest, fmt.Errorf("user_uuid is not found"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	agent, err := handler.service.GetAgent(ctx, agent_id, userUUID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, fmt.Errorf("error to try get a agent: %v", err))
		return
	}
	response.JSON(w, http.StatusOK, agent)
}

func (handler AgentModelHandler) RemoveAgentIa(w http.ResponseWriter, r *http.Request) {
	userUUID := r.Header.Get("user_uuid")
	agent_id := r.URL.Query().Get("agent_id")
	if userUUID == "" {
		http.Error(w, "user_uuid é obrigatório", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := handler.service.DeleteAgent(ctx, agent_id, userUUID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, fmt.Errorf("error to try remove agent: %v", err))
		return
	}
	response.JSON(w, http.StatusOK, res)
}
