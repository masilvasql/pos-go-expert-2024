package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/masilvasql/pos-go-expert-2024/APIs/configs"
	_ "github.com/masilvasql/pos-go-expert-2024/APIs/docs"
	"github.com/masilvasql/pos-go-expert-2024/APIs/internal/entity"
	"github.com/masilvasql/pos-go-expert-2024/APIs/internal/infra/database"
	"github.com/masilvasql/pos-go-expert-2024/APIs/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Marcelo Abrão da Silva
// @contact.url    https://www.linkedin.com/in/marceloabraodasilva/
// @contact.email  masilvasql@gmail.com

// @license.name   License
// @license.url    https://www.linkedin.com/in/marceloabraodasilva/

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// pathEnv := fmt.Sprintf(".", dir)
	config, err := configs.LoadConfig(".")

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
	userHandler := handlers.NewUserHandler(userDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", config.TokenAuth))
	router.Use(middleware.WithValue("JwtExpiresIn", config.JWTExpiresIn))
	// router.Use(LogRequest)

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

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", router)

}

/*
Exemplo de como funciona um middleware
*/
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
