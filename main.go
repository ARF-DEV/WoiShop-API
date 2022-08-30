package main

import (
	"azura-lab-intern/study-case-1/controller"
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleAuthConfig *oauth2.Config
	OAuthStateString string
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	GoogleAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/api/v1/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	OAuthStateString, _ = helpers.RandomString(10)
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected!")

	helpers.InitDB(db)
	productRepo := repository.CreateProductRepository(db)
	categoryRepo := repository.CreateCategoryRepository(db)
	userRepo := repository.CreateUserRepository(db)

	r := controller.NewRouter(GoogleAuthConfig, OAuthStateString, categoryRepo, productRepo, userRepo)

	log.Println("Listening in port 8000")
	http.ListenAndServe(":8000", r)
}
