package userusecase

import (
	"os"

	"dionpamungkas.com/my-commerce/src/entities"
	"dionpamungkas.com/my-commerce/src/helpers"

	userrepository "dionpamungkas.com/my-commerce/src/repositories/user"
)

type (
	ParamsCreateUser struct {
		Email    string
		Name     string
		Password string
	}
)

func (i *sUserUsecase) CreateUser(p *ParamsCreateUser) (*ResultLogin, error) {

	filters := entities.ParamsCreateUser{
		Email: p.Email,
	}

	checkEmail, _ := i.userRepository.IsExists(&filters)

	if checkEmail {
		return nil, ErrEmailAlreadyUsed
	}

	hashedPassword, _ := helpers.HashPassword(p.Password)
	data, err := i.userRepository.Create(&userrepository.ParamsCreateUser{
		Email:    p.Email,
		Name:     p.Name,
		Password: hashedPassword,
	})

	paramsGenerateJWTRegister := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		UserId:          data.ID,
		SecretKey:       os.Getenv("JWT_SECRET"),
	}

	accessToken, _, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTRegister)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	if err != nil {
		return nil, err
	}

	return &ResultLogin{
		Name:        p.Name,
		Email:       p.Email,
		AccessToken: accessToken,
	}, nil
}
