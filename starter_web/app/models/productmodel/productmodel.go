package productmodel

import (
	"go-starter-webapp/app/entities"
	"go-starter-webapp/config"
	"fmt"
)

func Getall() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			id, 
			display_name, 
			created_on,
			last_login FROM users limit 10
	`)

	if err != nil {
		fmt.Println("invalid query");
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}
