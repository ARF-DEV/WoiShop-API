package controller

import (
	"azura-lab-intern/study-case-1/controller/routes"
	"azura-lab-intern/study-case-1/middleware"
	"azura-lab-intern/study-case-1/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

type APIController struct {
	GoogleConfig     *oauth2.Config
	OAuthStateString string
	CategoryRepo     *repository.CategoryRepository
	ProductRepo      *repository.ProductRepository
	UserRepo         *repository.UserRepository
	CartRepo         *repository.CartRepository
	OrderRepo        *repository.OrderRepository
}

func (a *APIController) GetRouter() http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authorization(a.GoogleConfig))
		r.Get("/api/v1/products", routes.GetAllProduct(a.ProductRepo))
		r.Get("/api/v1/categories", routes.GetAllCategory(a.CategoryRepo))
		r.Get("/api/v1/products/{id}", routes.GetProductByID(a.ProductRepo))
		r.Get("/api/v1/carts/{id}", routes.GetCartByID(a.CartRepo, a.OrderRepo))
		r.Get("/api/v1/carts", routes.GetAllCart(a.CartRepo, a.OrderRepo))
	})
	r.Group(func(r chi.Router) {
		r.Get("/api/v1/login/oauth", routes.LoginOAuth(a.GoogleConfig, a.OAuthStateString))
		r.Get("/api/v1/callback", routes.HandleOAuthCallBack(a.GoogleConfig))
		r.Post("/api/v1/oauth/token", routes.GetAccessToken(a.GoogleConfig, a.OAuthStateString))
		r.Post("/api/v1/login", routes.Login(a.UserRepo))
		r.Post("/api/v1/register", routes.Register(a.UserRepo))
		r.Post("/api/v1/register/verify", routes.VerifyOTPRegister(a.UserRepo))

	})
	return r
}
