package configuration

import (
	"fmt"
	"gateway-ms/internal/infrastructure/commons/logger"
	"gateway-ms/internal/infrastructure/database"
	"gateway-ms/internal/infrastructure/security"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	Porta  = 8080
	Origin = ""
)

func LoadEnv() {
	if erro := godotenv.Load(); erro != nil {
		panic("Error ao carregar as variáveis de ambiente!")
	}

	slog.Info("Variáveis de ambiente carregadas com sucesso!")
}

func GetMainEnvs() {
	security.GetSecretKeyConfig(GetSecret("SECRET_KEY"))
	Origin = os.Getenv("ORIGINS")
}

// LoadLogger apenas para carregar logs personalizados
func LoadLogger() {
	custom_log := slog.New(logger.NewHandler(nil))
	slog.SetDefault(custom_log)
	slog.Info("Logger Carregado com sucesso!")
}

func LoadServer(routers http.Handler) {
	slog.Info(fmt.Sprintf("Servidor iniciado na porta %d", Porta))
	if erro := http.ListenAndServe(fmt.Sprintf(":%d", Porta), routers); erro != nil {
		panic(fmt.Sprintf("Error ao iniciar servidor %s", erro.Error()))
	}
}

func LoadDatabase() {
	db_name, _ := os.LookupEnv("DATABASE_NAME")
	db_host := os.Getenv("DATABASE_HOST")
	db_port := os.Getenv("DATABASE_PORT")
	db_user := GetSecret("DATABASE_USER")
	db_password := GetSecret("DATABASE_PASSWORD")

	database.SetDatabaseEnv(db_name, db_host, db_port, db_user, db_password)

	slog.Info("Conexão com Banco de dados Estabelecida")

	database.InitalStrucuture()
}
