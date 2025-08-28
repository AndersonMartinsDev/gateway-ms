package service

import (
	"gateway-ms/internal/domain/model"
	repositories "gateway-ms/internal/domain/repository"
	"gateway-ms/internal/infrastructure/security"
)

type UserService struct {
	repository     repositories.UserRepository
	handlePassword security.HandlePassword
}

func NewUserService() *UserService {
	return &UserService{
		repository:     *repositories.NewUserRepository(),
		handlePassword: *security.NewHandlerPassword(),
	}
}

func (service UserService) NewUser(user model.User) error {
	if erro := user.Check(); erro != nil {
		return erro
	}
	senhaComHash, erro := service.handlePassword.Hash(user.Password)
	if erro != nil {
		return erro
	}
	user.Password = string(senhaComHash)
	return service.repository.InsertNewUser(user)
}
