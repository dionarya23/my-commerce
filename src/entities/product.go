package entities

import (
	"time"
)

type Product struct {
	ID          string    `json:"id"`
	CategoryId  string    `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductSearchFilter struct {
	CategoryId string
}
