package userrepository

import (
	"database/sql"

	"dionpamungkas.com/my-commerce/src/entities"
)

type sUserRepository struct {
	DB *sql.DB
}

type UserRepository interface {
	Create(*ParamsCreateUser) (*entities.User, error)
	FindOne(*entities.ParamsCreateUser) (*entities.User, error)
	IsExists(*entities.ParamsCreateUser) (bool, error)
}

func New(db *sql.DB) UserRepository {
	return &sUserRepository{DB: db}
}
