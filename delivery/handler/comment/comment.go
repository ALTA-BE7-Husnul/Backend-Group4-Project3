package comment

import (
	"net/http"
	"project3/delivery/helper"
	_middlewares "project3/delivery/middlewares"
	_commentUseCase "project3/usecase/comment"

	_entities "project3/entities"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentUseCase _commentUseCase.CommentUseCaseInterface
}

func NewCommentHandler(u _commentUseCase.CommentUseCaseInterface) CommentHandler {
	return CommentHandler{
		commentUseCase: u,
	}
}

func (uh *CommentHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		comments, err := uh.commentUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseComment := []map[string]interface{}{}
		for i := 0; i < len(comments); i++ {
			response := map[string]interface{}{
				"id":       comments[i].ID,
				"user_id":  comments[i].UserID,
				"event_id": comments[i].EventID,
				"comment":  comments[i].Comment,
			}
			responseComment = append(responseComment, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all comments", responseComment))
	}
}

func (uh *CommentHandler) CreateCommentHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var param _entities.Comment

		errBind := c.Bind(&param)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errBind.Error()))
		}
		param.UserID = uint(idToken)

		_, err := uh.commentUseCase.CreateComment(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create comment"))
	}
}
