package main

import (
	"fmt"
	"log"
	"net/http"
	"project3/configs"

	_authHandler "project3/delivery/handler/auth"
	_middleware "project3/delivery/middlewares"
	_authRepository "project3/repository/auth"
	_authUseCase "project3/usecase/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_userHandler "project3/delivery/handler/users"
	_routes "project3/delivery/routes"
	_userRepository "project3/repository/user"
	_userUseCase "project3/usecase/user"

	_commentHandler "project3/delivery/handler/comment"
	_commentRepository "project3/repository/comment"
	_commentUseCase "project3/usecase/comment"

	_categoryHandler "project3/delivery/handler/category"
	_categoryRepository "project3/repository/category"
	_categoryUseCase "project3/usecase/category"

	_utils "project3/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	commentRepo := _commentRepository.NewCommentRepository(db)
	commentUseCase := _commentUseCase.NewCommentUseCase(commentRepo)
	commentHandler := _commentHandler.NewCommentHandler(commentUseCase)

	categoryRepo := _categoryRepository.NewCategoryRepository(db)
	categoryUseCase := _categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryHandler := _categoryHandler.NewCategoryHandler(categoryUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},

	}))
	e.Use(_middleware.CustomLogger())

	_routes.RegisterUserPath(e, userHandler)
	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterCommentPath(e, commentHandler)
	_routes.RegisterCategoryPath(e, categoryHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
