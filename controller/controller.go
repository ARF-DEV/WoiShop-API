package controller

import (
	"azura-lab-intern/study-case-1/controller/routes"
	"azura-lab-intern/study-case-1/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(categoryRepo *repository.CategoryRepository, productRepo *repository.ProductRepository, userRepo *repository.UserRepository) http.Handler {
	r := chi.NewRouter()

	r.Get("/api/v1/products/{id}", routes.GetProductByID(productRepo))
	r.Get("/api/v1/products", routes.GetAllProduct(productRepo))
	r.Get("/api/v1/categories", routes.GetAllCategory(categoryRepo))
	r.Post("/api/v1/login", routes.Login(userRepo))
	r.Post("/api/v1/register", routes.Register(userRepo))
	r.Post("/api/v1/register/verify", routes.VerifyOTPRegister(userRepo))
	return r
}
