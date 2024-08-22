package product

import "database/sql"

type ProductRepositoriyInterface interface {
	GetProduct(id int) (Product, error)
}

type ProductRepositoriy struct {
	db *sql.DB
}

func NewProductRepositoriy(db *sql.DB) *ProductRepositoriy {
	return &ProductRepositoriy{db}
}

func (r *ProductRepositoriy) GetProduct(id int) (Product, error) {
	return Product{
		ID:   id,
		Name: "Product Name",
	}, nil
}
