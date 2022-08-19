package controller

import (
	"azura-lab-intern/study-case-1/controller/routes"
	"azura-lab-intern/study-case-1/middleware"
	"azura-lab-intern/study-case-1/repository"
	"net/http"
)

func NewMux(categoryRepo *repository.CategoryRepository, productRepo *repository.ProductRepository) *http.ServeMux {
	m := http.NewServeMux()

	m.Handle("/api/v1/products", middleware.Method("GET", routes.GetAllProduct(productRepo)))
	m.Handle("/api/v1/product", middleware.Method("GET", routes.GetProductByID(productRepo)))
	m.Handle("/api/v1/product/by/category", middleware.Method("GET", routes.GetAllProductByCategory(productRepo)))
	m.Handle("/api/v1/categories", middleware.Method("GET", routes.GetAllCategory(categoryRepo)))

	return m
}
