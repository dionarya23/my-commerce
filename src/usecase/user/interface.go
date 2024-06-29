package userusecase

import (
	user "dionpamungkas.com/my-commerce/src/repositories/user"
)

type sUserUsecase struct {
	userRepository user.UserRepository
}

type UserUsecase interface {
	CreateUser(*ParamsCreateUser) (*ResultLogin, error)
	Login(*ParamsLogin) (*ResultLogin, error)
}

func New(
	userRepository user.UserRepository,
) UserUsecase {
	return &sUserUsecase{
		userRepository: userRepository,
	}
}
