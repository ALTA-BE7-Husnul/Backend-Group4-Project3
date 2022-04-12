package user

import (
	_entities "project3/entities"
)

type UserUseCaseInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
	UpdateUser(id int, request _entities.User) (_entities.User, int, error)
	DeleteUser(id int) error
	GetUserById(id int) (_entities.User, int, error)
}
