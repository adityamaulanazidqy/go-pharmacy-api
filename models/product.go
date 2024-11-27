package models

import (
	"go-pharmacy-api/config"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

// Get all products
func GetAllProducts() ([]Product, error) {
	rows, err := config.DB.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Create a new product
func CreateProduct(product *Product) error {
	query := "INSERT INTO products (name, description, price, stock) VALUES (?, ?, ?, ?)"
	result, err := config.DB.Exec(query, product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = int(id)
	return nil
}

// Update a product
func UpdateProduct(product *Product, id string) error {
	query := "UPDATE products SET name = ?, description = ?, price = ?, stock = ? WHERE id = ?"
	_, err := config.DB.Exec(query, product.Name, product.Description, product.Price, product.Stock, id)
	return err
}

// Delete a product
func DeleteProduct(id string) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := config.DB.Exec(query, id)
	return err
}

func ResetProducts() error {
	_, err := config.DB.Exec("DELETE FROM products")
	if err != nil {
		return err
	}
	_, err = config.DB.Exec("DELETE FROM sqlite_sequence WHERE name='products'")
	return err
}
