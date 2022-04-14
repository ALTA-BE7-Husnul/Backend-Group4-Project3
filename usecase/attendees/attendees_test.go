package attendees

import (
	"fmt"
	_entities "project3/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetAttendees(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepository{})
		data, err := attendeesUseCase.GetAttendees(_entities.Attendees{EventID: 1})
		assert.Nil(t, err)
		assert.Equal(t, uint(1), data[0].EventID)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepositoryError{})
		data, err := attendeesUseCase.GetAttendees(_entities.Attendees{})
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestCreateAttendees(t *testing.T) {
	t.Run("TestCreateUserSuccess", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepository{})
		data, rows, err := attendeesUseCase.CreateAttendees(_entities.Attendees{EventID: 1})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.EventID)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestCreateUserError", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepositoryError{})
		data, rows, err := attendeesUseCase.CreateAttendees(_entities.Attendees{})
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), data.EventID)
		assert.Nil(t, nil, err)
		assert.Equal(t, 1, rows)
	})
}


func TestDeleteAttendees(t *testing.T) {
	t.Run("TestDeleteUserSuccess", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepository{})
		rows, err := attendeesUseCase.DeleteAttendees(1, 1)
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), rows)
		
	})

	t.Run("TestDeleteUserError", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepositoryError{})
		rows, err := attendeesUseCase.DeleteAttendees(1, 1)
		assert.NotNil(t, err)
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), rows)
		
	})
}

// === mock success ===
type mockAttendeesRepository struct{}

func (m mockAttendeesRepository) GetAttendees(request _entities.Attendees) ([]_entities.Attendees, error) {
	return []_entities.Attendees{
		{EventID: 1, UserID: 1},
	}, nil
}

func (m mockAttendeesRepository) CreateAttendees(request _entities.Attendees) (_entities.Attendees, int, error) {
	return _entities.Attendees{
		EventID: 1, UserID: 1,
	}, 1, nil
}

func (m mockAttendeesRepository) DeleteAttendees(idToken uint, idEvent uint) (uint, error) {
	return 1, nil
}

// === mock error ===

type mockAttendeesRepositoryError struct{}

func (m mockAttendeesRepositoryError) GetAttendees(request _entities.Attendees) ([]_entities.Attendees, error) {
	return nil, fmt.Errorf("error get data user")
}

func (m mockAttendeesRepositoryError) CreateAttendees(request _entities.Attendees) (_entities.Attendees, int, error) {
	return _entities.Attendees{}, 1, fmt.Errorf("error create data user")
}

func (m mockAttendeesRepositoryError) DeleteAttendees(idToken uint, idEvent uint) (uint, error) {
	return 1, fmt.Errorf("error update data user")
}