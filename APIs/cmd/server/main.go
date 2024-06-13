package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
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
	config, err := configs.LoadConfig(pathEnv)

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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/generate-token", userHandler.GetJWT)

	http.ListenAndServe(":8000", router)

}
