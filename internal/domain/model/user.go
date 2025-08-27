package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password,omitempty"`
	Email    string    `json:"email,omitempty"`
}

func (u *User) Check() error {
	if erro := u.validateUsuario(); erro != nil {
		return erro
	}
	if erro := u.formataUsuario(); erro != nil {
		return erro
	}
	return nil
}

func (u *User) formataUsuario() error {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
	return nil
}

func (u *User) validateUsuario() error {
	if erro := validated(u.Name, "nome"); erro != nil {
		return erro
	}

	if erro := validated(u.Username, "username"); erro != nil {
		return erro
	}

	if erro := validateEmail(u.Email); erro != nil {
		return erro
	}

	if erro := validated(u.Password, "senha"); erro != nil {
		return erro
	}
	return nil
}

func validated(valor, campo string) error {
	if valor == "" {
		return fmt.Errorf("o campo %s não pode ser vazio", campo)
	}
	return nil
}

func validateEmail(email string) error {
	if erro := validated(email, "email"); erro != nil {
		return erro
	}

	if erro := checkmail.ValidateFormat(email); erro != nil {
		return errors.New("email não é válido")
	}

	return nil
}
