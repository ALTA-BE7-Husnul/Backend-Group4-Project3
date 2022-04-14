package main

import (
	"fmt"
	"log"
	"net/http"
	"project3/configs"

	_authHandler "project3/delivery/handler/auth"
	_eventHandler "project3/delivery/handler/event"
	_middleware "project3/delivery/middlewares"
	_authRepository "project3/repository/auth"
	_eventRepository "project3/repository/event"
	_authUseCase "project3/usecase/auth"
	_eventUseCase "project3/usecase/event"

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

	_attendeesHandler "project3/delivery/handler/attendees"
	_attendeesRepository "project3/repository/attendees"
	_attendeesUseCase "project3/usecase/attendees"

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

	eventRepo := _eventRepository.NewEventRepository(db)
	eventUseCase := _eventUseCase.NewEventUseCase(eventRepo)
	eventHandler := _eventHandler.NewEventHandler(eventUseCase)

	categoryRepo := _categoryRepository.NewCategoryRepository(db)
	categoryUseCase := _categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryHandler := _categoryHandler.NewCategoryHandler(categoryUseCase)

	attendeesRepo := _attendeesRepository.NewAttendeesRepository(db)
	attendeesUseCase := _attendeesUseCase.NewAttendeesUseCase(attendeesRepo)
	attendeesHandler := _attendeesHandler.NewAttendeesHandler(attendeesUseCase)

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
	_routes.RegisterEventPath(e, eventHandler)
	_routes.RegisterCategoryPath(e, categoryHandler)
	_routes.RegisterJoinPath(e, attendeesHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
