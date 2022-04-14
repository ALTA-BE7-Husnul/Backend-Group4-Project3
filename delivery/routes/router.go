package routes

import (
	_attendeesHandler "project3/delivery/handler/attendees"
	_authHandler "project3/delivery/handler/auth"
	_categoryHandler "project3/delivery/handler/category"
	_commentHandler "project3/delivery/handler/comment"
	_eventHandler "project3/delivery/handler/event"
	_userHandler "project3/delivery/handler/users"
	_middlewares "project3/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh _userHandler.UserHandler) {
	e.GET("/users", uh.GetUserByIdHandler(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUserHandler())
	e.PUT("/users/:id", uh.UpdateUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
}

func RegisterCommentPath(e *echo.Echo, uh _commentHandler.CommentHandler) {
	e.GET("/event/comments", uh.GetAllHandler())
	e.POST("/event/comments", uh.CreateCommentHandler(), _middlewares.JWTMiddleware())
}

func RegisterEventPath(e *echo.Echo, eh *_eventHandler.EventHandler) {
	e.POST("/event", eh.CreateEventHandler(), _middlewares.JWTMiddleware())
}

func RegisterCategoryPath(e *echo.Echo, uh _categoryHandler.CategoryHandler) {
	e.GET("/category", uh.GetAllCategoryHandler(), _middlewares.JWTMiddleware())
}

func RegisterJoinPath(e *echo.Echo, uh _attendeesHandler.AttendeesHandler) {
	e.POST("/event/participations", uh.CreateAttendeesHandler(), _middlewares.JWTMiddleware())
}
