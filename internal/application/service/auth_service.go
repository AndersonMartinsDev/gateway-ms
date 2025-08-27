package service

import (
	"fmt"
	"gateway-ms/internal/domain/model"
	repositories "gateway-ms/internal/domain/repository"
	"gateway-ms/internal/infrastructure/security"
)

type AuthService struct {
	Repository     repositories.UserRepository
	handlePassword security.HandlePassword
	handleToken    security.HandleToken
}

func NewAuthService() *AuthService {
	return &AuthService{
		Repository:     *repositories.NewUserRepository(),
		handlePassword: *security.NewHandlerPassword(),
		handleToken:    *security.NewHandlerToken(),
	}
}

func (serv AuthService) Login(usuario *model.User) error {
	login, erro := serv.Repository.UserLogin(usuario.Username)
	if erro != nil {
		return erro
	}

	erro = login.Check()
	if erro != nil {
		return fmt.Errorf("Unauthorized")
	}
	usuario.UUID = login.UUID
	return serv.handlePassword.VerificarSenha(login.Password, usuario.Password)
}

func (serv AuthService) CriarToken(usuario model.User) (string, error) {
	return serv.handleToken.CriarToken(usuario)
}
