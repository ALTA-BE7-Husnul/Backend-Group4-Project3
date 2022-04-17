package attendees

import (
	"errors"
	_entities "project3/entities"
	_attendeesRepository "project3/repository/attendees"
)

type AttendeesUseCase struct {
	attendeesRepository _attendeesRepository.AttendeesRepositoryInterface
}

func NewAttendeesUseCase(attendeesRepo _attendeesRepository.AttendeesRepositoryInterface) AttendeesUseCaseInterface {
	return &AttendeesUseCase{
		attendeesRepository: attendeesRepo,
	}
}

func (uuc *AttendeesUseCase) CreateAttendees(request _entities.Attendees) (_entities.Attendees, int, error) {
	attendees, rows, err := uuc.attendeesRepository.CreateAttendees(request)
	if request.EventID == 0 {
		return attendees, 1, errors.New("can't be empty")
	}
	if request.UserID == 0 {
		return attendees, 1, errors.New("can't be empty")
	}

	return attendees, rows, err
}

func (uuc *AttendeesUseCase) GetAttendees(request _entities.Attendees) ([]_entities.Attendees, error) {
	attendeess, err := uuc.attendeesRepository.GetAttendees(request)
	return attendeess, err
}

func (uuc *AttendeesUseCase) DeleteAttendees(idToken uint, idEvent uint) (uint, error) {
	rows, err := uuc.attendeesRepository.DeleteAttendees(idToken, idEvent)
	return rows, err
}

func (uuc *AttendeesUseCase) GetEventsByUserId(user_ID int) ([]_entities.Attendees, error) {
	attendees, err := uuc.attendeesRepository.GetEventsByUserId(user_ID)
	return attendees, err
}
