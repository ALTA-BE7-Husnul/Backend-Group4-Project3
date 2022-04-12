package routes

import (
	_userHandler "project3/delivery/handler/users"
	_middlewares "project3/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterUserPath(e *echo.Echo, uh _userHandler.UserHandler) {
	e.GET("/users", uh.GetUserByIdHandler(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUserHandler())
	e.PUT("/users/:id", uh.UpdateUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
}