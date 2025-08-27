package security

import (
	"fmt"
	"gateway-ms/internal/domain/model"
	"log/slog"

	// "mineops/src/domain/security/entities"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type HandleToken struct {
	userKey string
}

func NewHandlerToken() *HandleToken {
	return &HandleToken{
		userKey: "user_uuid",
	}
}

// CriarToken token com as permissões de usuário
func (tk HandleToken) CriarToken(user model.User) (string, error) {
	//1234
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes[tk.userKey] = user.UUID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	token.Valid = true
	return token.SignedString([]byte(SECRET_KEY)) //secret
}

// ValidarToken verifica se o token na requisição é valido
func (tk HandleToken) ValidarToken(r *http.Request) error {
	tokenString := tk.extrairToken(r)
	token, erro := jwt.Parse(tokenString, tk.retornarChaveDeVerificacao)

	if erro != nil {
		slog.Error(fmt.Sprintf("handleToken: %s", erro))
		return fmt.Errorf("invalid token")
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_uuid := permissoes["user_uuid"].(string)
		r.Header.Set(tk.userKey, user_uuid)
		return nil
	}
	return fmt.Errorf("token inválido")
}

func (tk HandleToken) extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func (tk HandleToken) retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}
	return []byte(SECRET_KEY), nil
}
