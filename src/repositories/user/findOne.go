package userrepository

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"dionpamungkas.com/my-commerce/src/entities"
)

func (i *sUserRepository) FindOne(filters *entities.ParamsCreateUser) (*entities.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE "
	params := []interface{}{}
	conditions := []string{}

	if filters.ID != 0 {
		conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filters.ID)
	}
	if filters.Name != "" {
		conditions = append(conditions, "name = $"+strconv.Itoa(len(params)+1))
		params = append(params, filters.Name)
	}
	if filters.Email != "" {
		conditions = append(conditions, "email = $"+strconv.Itoa(len(params)+1))
		params = append(params, filters.Email)
	}

	query += strings.Join(conditions, " AND ")

	query += " LIMIT 1"

	row := i.DB.QueryRow(query, params...)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		log.Printf("Error find user: %s", err)
		if err == sql.ErrNoRows {
			return nil, nil // Return nil for both user and error
		}
		return nil, err
	}

	return &user, nil
}
