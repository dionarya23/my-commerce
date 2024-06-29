package categoryrepository

import (
	"database/sql"

	"dionpamungkas.com/my-commerce/src/entities"
)

type sCategoryRepository struct {
	DB *sql.DB
}

type CategoryRepository interface {
	FindAll() ([]*entities.Category, error)
}

func New(db *sql.DB) CategoryRepository {
	return &sCategoryRepository{DB: db}
}
