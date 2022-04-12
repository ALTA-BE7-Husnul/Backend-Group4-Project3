package main

import (
	"fmt"
	"log"
	"net/http"
	"project3/configs"

	_authHandler "project3/delivery/handler/auth"
	_middleware "project3/delivery/middlewares"
	_routes "project3/delivery/routes"
	_authRepository "project3/repository/auth"
	_authUseCase "project3/usecase/auth"
	"project3/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_userHandler "project3/delivery/handler/users"
	_userRepository "project3/repository/user"
	_userUseCase "project3/usecase/user"

	_routes "project3/delivery/routes"
	_utils "project3/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)
	
	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(_middleware.CustomLogger())
	
	_routes.RegisterUserPath(e, userHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
