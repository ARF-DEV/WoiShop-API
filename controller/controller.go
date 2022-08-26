package controller

import (
	"azura-lab-intern/study-case-1/controller/routes"
	"azura-lab-intern/study-case-1/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(categoryRepo *repository.CategoryRepository, productRepo *repository.ProductRepository) http.Handler {
	r := chi.NewRouter()

	r.Get("/api/v1/products/{id}", routes.GetProductByID(productRepo))
	r.Get("/api/v1/products", routes.GetAllProduct(productRepo))
	r.Get("/api/v1/categories", routes.GetAllCategory(categoryRepo))
	return r
}
