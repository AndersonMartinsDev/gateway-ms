package handler

import (
	"gateway-ms/internal/application/service"
	"gateway-ms/internal/domain/model"
	"gateway-ms/internal/infrastructure/request"
	"gateway-ms/internal/infrastructure/response"
	"net/http"
)

// UserHandler é o handler HTTP para as requisições de usuário.
// Ele recebe o serviço de usuário por injeção de dependência.
type AuthHandler struct {
	service *service.AuthService
}

// NewUserHandler cria uma nova instância de UserHandler.
func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (handler AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var usuario model.User

	if erro := request.Serialization(r.Body, &usuario); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if erro := handler.service.Login(&usuario); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := handler.service.CriarToken(usuario)

	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	response.Response(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: token,
	}, nil)
}
