package main

import (
	"fmt"
	"log"
	"project3/configs"
	_authHandler "project3/delivery/handler/auth"
	_middleware "project3/delivery/middlewares"
	_routes "project3/delivery/routes"
	_authRepository "project3/repository/auth"
	_authUseCase "project3/usecase/auth"
	"project3/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middleware.CustomLogger())

	_routes.RegisterAuthPath(e, authHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
