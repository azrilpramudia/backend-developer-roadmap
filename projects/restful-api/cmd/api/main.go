package main

import (
	"log"
	"net/http"

	"github.com/azrilpramudia/go-restful-api/internal/config"
	"github.com/azrilpramudia/go-restful-api/internal/docs"
	"github.com/azrilpramudia/go-restful-api/internal/product"
	"github.com/azrilpramudia/go-restful-api/internal/repository/postgres"
	"github.com/azrilpramudia/go-restful-api/pkg/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresConnection(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer db.Close()

	productRepo := postgres.NewProductRepository(db)
	productUsecase := product.NewUsecase(productRepo)
	productHandler := product.NewHandler(productUsecase)

	mux := http.NewServeMux()
	productHandler.RegisterRoutes(mux)

	docsHandler := docs.NewHandler()
	docsHandler.RegisterRoutes(mux)

	log.Printf("server running on :%s", cfg.AppPort)
	if err := http.ListenAndServe(":"+cfg.AppPort, mux); err != nil {
		log.Fatal(err)
	}
}