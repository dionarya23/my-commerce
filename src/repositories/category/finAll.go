package categoryrepository

import (
	"log"

	"dionpamungkas.com/my-commerce/src/entities"
)

func (i *sCategoryRepository) FindAll() ([]*entities.Category, error) {
	query := "SELECT id, name FROM categories"

	rows, err := i.DB.Query(query)
	if err != nil {
		log.Printf("Error finding categories: %s", err)
		return nil, err
	}
	defer rows.Close()

	categories := make([]*entities.Category, 0)
	for rows.Next() {
		category := new(entities.Category)
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
