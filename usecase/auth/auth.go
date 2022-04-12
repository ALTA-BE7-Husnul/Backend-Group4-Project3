package auth

import (
	"project3/repository/auth"
)

type AuthUseCase struct {
	authRepository auth.AuthRepositoryInterface
}

func NewAuthUseCase(authRepo auth.AuthRepositoryInterface) AuthUseCaseInterface {
	return &AuthUseCase{
		authRepository: authRepo,
	}
}

func (auc *AuthUseCase) Login(email string, password string) (string, error) {
	token, err := auc.authRepository.Login(email, password)
	return token, err
}
