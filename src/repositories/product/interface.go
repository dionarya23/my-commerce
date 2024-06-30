package productrepository

import (
	"database/sql"

	"dionpamungkas.com/my-commerce/src/entities"
)

type sProductRepository struct {
	DB *sql.DB
}

type ProductRepository interface {
	FindAll(filters *entities.ProductSearchFilter) ([]*entities.Product, error)
}

func New(db *sql.DB) ProductRepository {
	return &sProductRepository{DB: db}
}
