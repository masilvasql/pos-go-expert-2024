//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/masilvsql/di/product"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepositoriy,
	wire.Bind(new(product.ProductRepositoriyInterface), new(*product.ProductRepositoriy)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
