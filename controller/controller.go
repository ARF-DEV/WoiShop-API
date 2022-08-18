package controller

import (
	"azura-lab-intern/study-case-1/controller/routes"
	"azura-lab-intern/study-case-1/repository"
	"net/http"
)

func NewMux(categoryRepo *repository.CategoryRepository, productRepo *repository.ProductRepository) *http.ServeMux {
	m := http.NewServeMux()

	m.Handle("/api/v1/products", routes.GetAllProduct(productRepo))
	m.Handle("/api/v1/product/by/id", routes.GetProductByID(productRepo))
	m.Handle("/api/v1/product/by/category", routes.GetAllProductByCategory(productRepo))
	m.Handle("/api/v1/categories", routes.GetAllCategory(categoryRepo))

	return m
}
