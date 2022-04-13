package event

import (
	"net/http"
	"project3/delivery/handler/image"
	"project3/delivery/helper"
	_middlewares "project3/delivery/middlewares"
	_entities "project3/entities"
	_eventUseCase "project3/usecase/event"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventUseCase _eventUseCase.EventUseCaseInterface
}

func NewEventHandler(event _eventUseCase.EventUseCaseInterface) *EventHandler {
	return &EventHandler{
		eventUseCase: event,
	}
}

func (eh *EventHandler) CreateEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// binding data
		var event _entities.Event
		errBind := c.Bind(&event)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error binding data"))
		}
		// binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image")
		if err_binding_image != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind image"))
		}
		// check file CheckFileExtension
		_, err_check_extension := image.CheckFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error checking file extension"))
		}
		// check file size
		err_check_size := image.CheckFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error checking file size"))
		}
		fileName := "events_" + strconv.Itoa(idToken) + "_" + strconv.Itoa(int(event.ID))
		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := image.UploadImage("events", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to upload file"))
		}
		// create certificate
		imageURL := theUrl
		err_event := eh.eventUseCase.CreateEvent(idToken, event, imageURL)
		if err_event != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to create event"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to create event"))
	}
}

func (eh *EventHandler) GetEventsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var events []_entities.Event
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get events", events))
	}
}
