package repository

import (
	"gateway-ms/internal/domain/model"
	"gateway-ms/internal/infrastructure/database"

	"github.com/google/uuid"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repository UserRepository) UserLogin(username string) (model.User, error) {
	db := database.GetConnectionDatabase()

	query := `SELECT *
			  FROM 
			  	tb_user
			  WHERE nickname = $1`
	row, erro := db.Query(query, username)

	if erro != nil {
		return model.User{}, erro
	}

	var login model.User
	if row.Next() {
		if erro := row.Scan(
			&login.UUID,
			&login.Name,
			&login.Username,
			&login.Password,
			&login.Email,
		); erro != nil {
			return model.User{}, erro
		}
	}
	defer row.Close()

	return login, nil
}

// NovoUsuario criação de novo usuario
func (repository UserRepository) InsertNewUser(user model.User) error {
	db := database.GetConnectionDatabase()
	insert := `
		INSERT INTO tb_user(
			uuid,
			username,
			nickname,
			pass_word,
			email) 
		VALUES($1, $2, $3, $4, $5)`
	statement, erro := db.Prepare(insert)

	if erro != nil {
		return erro
	}

	_, erro = statement.Exec(uuid.New(), user.Name, user.Username, user.Password, user.Email)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	return nil
}
