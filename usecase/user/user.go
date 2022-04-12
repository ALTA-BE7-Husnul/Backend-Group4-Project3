package user

import (
	"errors"
	"project3/delivery/helper"
	_entities "project3/entities"
	_userRepository "project3/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}


func (uuc *UserUseCase) CreateUser(request _entities.User) (_entities.User, error) {
	password, err := helper.HashPassword(request.Password)
	request.Password = password
	users, err := uuc.userRepository.CreateUser(request)
	
	if request.Name == "" {
		return users, errors.New("can't be empty")
	}
	if request.Email == "" {
		return users, errors.New("can't be empty")
	}
	if request.Password == "" {
		return users, errors.New("can't be empty")
	}
	
	return users, err
}

func (uuc *UserUseCase) UpdateUser(id int, request _entities.User) (_entities.User, int, error) {
	password, err := helper.HashPassword(request.Password)
	request.Password = password
	user, rows, err := uuc.userRepository.GetUserById(id)
	if err != nil {
		return user, 0, err
	}
	if rows == 0 {
		return user, 0, nil
	}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}

	users, rows, err := uuc.userRepository.UpdateUser(user)
	return users, rows, err
}

func (uuc *UserUseCase) DeleteUser(id int) error {
	err := uuc.userRepository.DeleteUser(id)
	return err
}

func (uuc *UserUseCase) GetUserById(id int) (_entities.User, int, error) {
	users, rows, err := uuc.userRepository.GetUserById(id)
	return users, rows, err
}
