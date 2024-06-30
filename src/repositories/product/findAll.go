package productrepository

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"dionpamungkas.com/my-commerce/src/entities"
)

func (i *sProductRepository) FindAll(filters *entities.ProductSearchFilter) ([]*entities.Product, error) {
	query := "SELECT id, category_id, name, description FROM products WHERE 1=1"
	params := []interface{}{}

	n := (&entities.ProductSearchFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}
		if filters.CategoryId != "" {
			conditions = append(conditions, "category_id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.CategoryId)
		}

		if len(conditions) > 0 {
			query += " AND "
		}
		query += strings.Join(conditions, " AND ")
	}

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding products: %s", err)
		return nil, err
	}
	defer rows.Close()

	products := make([]*entities.Product, 0)
	for rows.Next() {
		product := new(entities.Product)
		err := rows.Scan(&product.ID, &product.CategoryId, &product.Name, &product.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
