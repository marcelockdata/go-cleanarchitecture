package main

import (
	"cleanarchitecture/internal/handler"
	authhmiddleware "cleanarchitecture/internal/middleware"
	"cleanarchitecture/internal/repository"
	"cleanarchitecture/internal/usecase"
	"log"
	"net/http"

	_ "cleanarchitecture/docs" // Importe o pacote docs gerado pelo swag

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Clean Architecture API
// @version 1.0
// @description This is a sample server for a clean architecture application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Ajuste para funcionar no Docker
	db, err := repository.InitDB("/app/products.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	productRepository := repository.NewProductRepository(db)
	userRepository := repository.NewUserRepository(db)

	productUsecase := usecase.NewProductUsecase(productRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)

	productHandler := handler.NewProductHandler(productUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/auth", userHandler.Authenticate)

	r.Route("/products", func(r chi.Router) {
		r.Use(authhmiddleware.AuthMiddleware)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.ListProducts)
		r.Get("/{id}", productHandler.GetProductById)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	// Swagger route
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), // URL do arquivo gerado
	))

	log.Println("Server running on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
