package main

import (
	"azura-lab-intern/study-case-1/config"
	"azura-lab-intern/study-case-1/controller"
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleAuthConfig *oauth2.Config
	OAuthStateString string
)

func main() {
	GoogleAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/api/v1/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	OAuthStateString, _ = helpers.RandomString(10)
	db := config.ConfigDatabase()
	defer db.Close()
	helpers.MigrateDB(db)
	productRepo := repository.CreateProductRepository(db)
	categoryRepo := repository.CreateCategoryRepository(db)
	userRepo := repository.CreateUserRepository(db)
	api := controller.APIController{
		GoogleConfig:     GoogleAuthConfig,
		OAuthStateString: OAuthStateString,
		ProductRepo:      productRepo,
		CategoryRepo:     categoryRepo,
		UserRepo:         userRepo,
		CartRepo:         repository.NewCartRepository(db),
		OrderRepo:        repository.NewOrderRepository(db),
	}

	r := api.GetRouter()

	log.Println("Listening in port 8000")
	http.ListenAndServe(":8000", r)
}
