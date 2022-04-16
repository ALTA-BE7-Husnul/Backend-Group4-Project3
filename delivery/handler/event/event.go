package event

import (
	"fmt"
	"net/http"
	"project3/delivery/handler/image"
	"project3/delivery/helper"
	_middlewares "project3/delivery/middlewares"
	_entities "project3/entities"
	_eventUseCase "project3/usecase/event"
	"strconv"
	"time"

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
		var eventRequest EventRequest
		errBind := c.Bind(&eventRequest)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error to bind data"))
		}
		// formatting time
		layoutFormat := "2006-01-02T15:04:05Z0700"
		dateFormat := fmt.Sprintf("%s:00+0700", eventRequest.Date)
		dateParse, err_date_parse := time.Parse(layoutFormat, dateFormat)
		if err_date_parse != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error to format time.Time"))
		}

		fmt.Println(eventRequest.Date) //debugging
		fmt.Println(dateParse)         //debugging
		//set eventRequest to event
		var event _entities.Event
		event.UserID = eventRequest.UserID
		event.CategoryID = eventRequest.CategoryID
		event.Name = eventRequest.Name
		event.Host = eventRequest.Host
		event.Date = dateParse
		event.Location = eventRequest.Location
		event.Details = eventRequest.Details
		event.Quota = eventRequest.Quota
		event.Participants = eventRequest.Participants
		event.Image = eventRequest.Image

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
		allEventResponse := []map[string]interface{}{}
		for i := range events {
			response := map[string]interface{}{
				"id":       events[i].ID,
				"name":     events[i].Name,
				"host":     events[i].Host,
				"date":     events[i].Date,
				"location": events[i].Location,
				"image":    events[i].Image,
				"details":  events[i].Details,
				"quota":    events[i].Quota,
			}
			allEventResponse = append(allEventResponse, response)
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all events", allEventResponse))
	}
}

func (eh *EventHandler) GetEventByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		event_ID, _ := strconv.Atoi(c.Param("id"))
		event, err := eh.eventUseCase.GetEventById(event_ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get event"))
		}
		var eventResponse GetEventResponse
		eventResponse.ID = event.ID
		eventResponse.Name = event.Name
		eventResponse.Host = event.Host
		eventResponse.Details = event.Details
		eventResponse.Date = event.Date
		eventResponse.Location = event.Location
		eventResponse.Quota = event.Quota
		eventResponse.Image = event.Image
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get event", eventResponse))
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
		// update event
		_, rows, err_event := eh.eventUseCase.UpdateEvent(event, id, idToken)
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
