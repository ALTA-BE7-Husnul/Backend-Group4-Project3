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
		eventLength, eventLenErr := eh.eventUseCase.GetEvents()
		if eventLenErr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("there is no event yet"))

		}
		lastID := eventLength[len(eventLength)-1].ID
		fileName := "events_" + strconv.Itoa(idToken) + "_" + strconv.Itoa(int(lastID)+1)
		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := image.UploadImage("events", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to upload file"))
		}
		// create event
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
		events, err := eh.eventUseCase.GetEvents()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get events"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get events", events))
	}
}

func (eh *EventHandler) DeleteEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		id, _ := strconv.Atoi(c.Param("id"))
		rows, err := eh.eventUseCase.DeleteEvent(id, idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to operate delete"))
		}
		if rows == 0 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("event not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to delete event"))
	}
}

func (eh *EventHandler) UpdateEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		id, _ := strconv.Atoi(c.Param("id"))
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
		fileName := "events_" + strconv.Itoa(idToken) + "_" + strconv.Itoa(id)
		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := image.UploadImage("events", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to upload file"))
		}
		// update event
		imageURL := theUrl
		_, rows, err_event := eh.eventUseCase.UpdateEvent(event, id, idToken, imageURL)
		if err_event != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to update event"))
		}
		if rows == 1 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("database failed to be updated"))
		}
		if rows == 2 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unexpected update field"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to update event"))
	}
}
