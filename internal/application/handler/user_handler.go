package handler

import (
	"gateway-ms/internal/application/service"
	"gateway-ms/internal/domain/model"
	"gateway-ms/internal/infrastructure/request"
	"gateway-ms/internal/infrastructure/response"
	"net/http"
)

// UserHandler é o handler HTTP para as requisições de usuário.
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler cria uma nova instância de UserHandler.
func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

// InsertNewUser rota para inserção de novo usuario
func (handle UserHandler) InsertNewUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if erro := request.Serialization(r.Body, &user); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := handle.userService.NewUser(user); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.Response(w, http.StatusCreated, "Success to create user !", nil)
}
