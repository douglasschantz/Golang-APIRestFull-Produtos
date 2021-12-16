package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/schantz/web/go-api-produtos/backend/database"
	"github.com/schantz/web/go-api-produtos/backend/models/product"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	database.InitDB()

	var (
		productRepository = product.NewRepository(database.ReturnDB())
	)

	var (
		productService product.Service
	)

	productService = product.NewService(productRepository)

	r := chi.NewRouter()

	r.Mount("/products", product.MakeHttpHandler(productService))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))

	http.ListenAndServe(":3000", r)
}
