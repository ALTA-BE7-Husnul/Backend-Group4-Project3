package user

import (
	"fmt"
	_entities "project3/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetUserById(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.GetUserById(1)
		assert.Nil(t, err)
		assert.Equal(t, "odi", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, rows, err := userUseCase.GetUserById(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.User{}, data)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("TestCreateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreateUser(_entities.User{Name: "odi"})
		assert.Nil(t, nil, err)
		assert.Equal(t, "haudhi", data.Name)
	})

	t.Run("TestCreateUser2Success", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreateUser(_entities.User{Name: "odi", Email: "odi@mail.com"})
		assert.Nil(t, nil, err)
		assert.Equal(t, "haudhi", data.Name)
	})

	t.Run("TestCreateUser3Success", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreateUser(_entities.User{Name: "odi", Email: "odi@mail.com", Password: "lalala"})
		assert.Nil(t, nil, err)
		assert.Equal(t, "haudhi", data.Name)
	})

	t.Run("TestCreateUser4Success", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreateUser(_entities.User{Name: "odi", Email: "odi@mail.com"})
		assert.Nil(t, nil, err)
		assert.Equal(t, "haudhi", data.Name)
	})

	t.Run("TestCreateUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.CreateUser(_entities.User{Name: "odi"})
		assert.NotNil(t, err)
		assert.Equal(t, "", data.Name)
		assert.Nil(t, nil, err)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.UpdateUser(1, _entities.User{Name: "almas"})
		assert.Nil(t, err)
		assert.Equal(t, "almas", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.UpdateUser(1, _entities.User{Name: "almas", Email: "odi@mail.com"})
		assert.Nil(t, err)
		assert.Equal(t, "almas", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.UpdateUser(1, _entities.User{Name: "almas", Email: "odi@mail.com"})
		assert.Nil(t, err)
		assert.Equal(t, "almas", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, rows, err := userUseCase.UpdateUser(1, _entities.User{})
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Nil(t, nil, data.Name)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("TestDeleteUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		err := userUseCase.DeleteUser(1)
		assert.Nil(t, err)
		
	})

	t.Run("TestDeleteUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		err := userUseCase.DeleteUser(1)
		assert.NotNil(t, err)
		
	})
}

// === mock success ===
type mockUserRepository struct{}

func (m mockUserRepository) GetUserById(id int) (_entities.User, int, error) {
	return _entities.User{
		Name: "odi", Email: "odi@mail.com", Password: "lalala",
	}, 1, nil
}

func (m mockUserRepository) CreateUser(request _entities.User) (_entities.User, error) {
	return _entities.User{
		Name: "haudhi", Email: "odi@mail.com", Password: "lalala",
	}, nil
}

func (m mockUserRepository) UpdateUser(request _entities.User) (_entities.User, int, error) {
	return _entities.User{
		Name: "almas", Email: "odi@mail.com", Password: "lalala",
	}, 1, nil
}

func (m mockUserRepository) DeleteUser(id int) error {
	return nil
}


// === mock error ===

type mockUserRepositoryError struct{}

func (m mockUserRepositoryError) GetUserById(id int) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error get data user")
}

func (m mockUserRepositoryError) CreateUser(request _entities.User) (_entities.User, error) {
	return _entities.User{}, fmt.Errorf("error create data user")
}

func (m mockUserRepositoryError) UpdateUser(request _entities.User) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) DeleteUser(id int) error {
	return fmt.Errorf("error update data user")
}