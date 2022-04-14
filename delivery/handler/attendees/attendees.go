package attendees

import (
	"net/http"
	"project3/delivery/helper"
	_middlewares "project3/delivery/middlewares"
	_attendeesUseCase "project3/usecase/attendees"

	_entities "project3/entities"

	"github.com/labstack/echo/v4"
)

type AttendeesHandler struct {
	attendeesUseCase _attendeesUseCase.AttendeesUseCaseInterface
}

func NewAttendeesHandler(u _attendeesUseCase.AttendeesUseCaseInterface) AttendeesHandler {
	return AttendeesHandler{
		attendeesUseCase: u,
	}
}

func (uh *AttendeesHandler) CreateAttendeesHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var param _entities.Attendees

		errBind := c.Bind(&param)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errBind.Error()))
		}
		param.UserID = uint(idToken)

		_, rows, err := uh.attendeesUseCase.CreateAttendees(param)

		if rows == 1 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("fail to read event"))
		}

		if rows == 2 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("quota is full"))
		}

		if rows == 3 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("you have joined"))
		}

		if rows == 4 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("fail to read attendees"))
		}
		if rows == 6 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("event not found"))
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success join event"))
	}
}