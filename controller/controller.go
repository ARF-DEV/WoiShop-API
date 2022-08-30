package controller

import (
	"azura-lab-intern/study-case-1/controller/routes"
	"azura-lab-intern/study-case-1/middleware"
	"azura-lab-intern/study-case-1/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

func NewRouter(GoogleConfig *oauth2.Config, OAuthStateString string, categoryRepo *repository.CategoryRepository, productRepo *repository.ProductRepository, userRepo *repository.UserRepository) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authorization(GoogleConfig))
		r.Get("/api/v1/products", routes.GetAllProduct(productRepo))
		r.Get("/api/v1/categories", routes.GetAllCategory(categoryRepo))
		r.Get("/api/v1/products/{id}", routes.GetProductByID(productRepo))
	})
	r.Group(func(r chi.Router) {
		r.Get("/api/v1/login/oauth", routes.LoginOAuth(GoogleConfig, OAuthStateString))
		r.Get("/api/v1/callback", routes.HandleOAuthCallBack(GoogleConfig))
		r.Post("/api/v1/oauth/token", routes.GetAccessToken(GoogleConfig, OAuthStateString))
		r.Post("/api/v1/login", routes.Login(userRepo))
		r.Post("/api/v1/register", routes.Register(userRepo))
		r.Post("/api/v1/register/verify", routes.VerifyOTPRegister(userRepo))

	})
	return r
}
