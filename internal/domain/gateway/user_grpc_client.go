package gateway

import (
	"context"
	"gateway-ms/internal/domain/model"
)

// UserClient é a interface que define a comunicação com o microserviço de usuários.
// A lógica de negócio do gateway vai depender dessa interface, não da implementação gRPC.
type UserClient interface {
	// GetUserByID busca um usuário no serviço de usuários via gRPC.
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	// CreateUser cria um novo usuário no serviço de usuários via gRPC.
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}
