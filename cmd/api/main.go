package main

import (
	"gateway-ms/internal/application/handler"
	"gateway-ms/internal/application/service"
	"gateway-ms/internal/infrastructure/configuration"
	"gateway-ms/internal/infrastructure/http_server"
	"log/slog"
)

func main() {
	configuration.LoadEnv()
	configuration.LoadLogger()

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(authService)

	var routes []http_server.RouterInterface
	routes = append(routes, http_server.NewUserRoute(*userHandler))
	routes = append(routes, http_server.NewAuthRoute(*authHandler))
	slog.Info("Rotas HTTP registradas com sucesso!")

	routerHandles := http_server.NewRouters(routes)

	configuration.LoadDatabase()
	configuration.LoadServer(routerHandles)

}
