package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/masilvasql/pos-go-expert-2024/APIs/configs"
	"github.com/masilvasql/pos-go-expert-2024/APIs/internal/entity"
	"github.com/masilvasql/pos-go-expert-2024/APIs/internal/infra/database"
	"github.com/masilvasql/pos-go-expert-2024/APIs/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pathEnv := fmt.Sprintf("%s\\.env", dir)
	_, err = configs.LoadConfig(pathEnv)

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Product{}, entity.User{})
	if err != nil {
		panic(err)
	}

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products/{id}", productHandler.GetProduct)
	router.Put("/products/{id}", productHandler.UpdateProduct)

	http.ListenAndServe(":8000", router)

}
