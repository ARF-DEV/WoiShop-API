package controller_test

import (
	"azura-lab-intern/study-case-1/controller"
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	_ "github.com/lib/pq"
)

var (
	GoogleAuthConfig  *oauth2.Config
	OAuthStateString  string
	GoogleRedirectURL string = "http://localhost:8000/api/v1/callback"
)

func TestGetAllProduct(t *testing.T) {

	err := godotenv.Load("./../.env")

	if err != nil {
		assert.NoError(t, err, "Cannot load .env file")
	}

	GoogleAuthConfig = &oauth2.Config{
		RedirectURL:  GoogleRedirectURL,
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
		assert.NoError(t, err, "Cannot open database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		assert.NoError(t, err, "Cannot ping database")
	}
	fmt.Println("Database Connected!")

	helpers.MigrateDB(db)
	productRepo := repository.CreateProductRepository(db)
	categoryRepo := repository.CreateCategoryRepository(db)
	userRepo := repository.CreateUserRepository(db)

	mux := controller.NewRouter(GoogleAuthConfig, OAuthStateString, categoryRepo, productRepo, userRepo)

	req := httptest.NewRequest("GET", "/api/v1/products", nil)
	res := httptest.NewRecorder()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	result_byte := res.Body.Bytes()

	var responseBody helpers.JsonResponse

	err = json.Unmarshal(result_byte, &responseBody)

	assert.NoError(t, err)

	listOfProduct := responseBody.Data.([]interface{})
	assert.Equal(t, 50, len(listOfProduct))
}

func TestGetProductByID(t *testing.T) {
	err := godotenv.Load("./../.env")

	if err != nil {
		assert.NoError(t, err, "Cannot load .env file")
	}

	GoogleAuthConfig = &oauth2.Config{
		RedirectURL:  GoogleRedirectURL,
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
		assert.NoError(t, err, "Cannot open database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		assert.NoError(t, err, "Cannot ping database")
	}
	fmt.Println("Database Connected!")

	helpers.MigrateDB(db)
	productRepo := repository.CreateProductRepository(db)
	categoryRepo := repository.CreateCategoryRepository(db)
	userRepo := repository.CreateUserRepository(db)

	mux := controller.NewRouter(GoogleAuthConfig, OAuthStateString, categoryRepo, productRepo, userRepo)

	req := httptest.NewRequest("GET", "/api/v1/products/40", nil)
	res := httptest.NewRecorder()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	result_byte := res.Body.Bytes()

	var responseBody helpers.JsonResponse

	err = json.Unmarshal(result_byte, &responseBody)

	assert.NoError(t, err)

	Product := responseBody.Data.(map[string]interface{})

	assert.Equal(t, 40, int(Product["id"].(float64)))
}

func TestGetAllCategory(t *testing.T) {
	err := godotenv.Load("./../.env")

	if err != nil {
		assert.NoError(t, err, "Cannot load .env file")
	}

	GoogleAuthConfig = &oauth2.Config{
		RedirectURL:  GoogleRedirectURL,
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
		assert.NoError(t, err, "Cannot open database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		assert.NoError(t, err, "Cannot ping database")
	}
	fmt.Println("Database Connected!")

	helpers.MigrateDB(db)
	productRepo := repository.CreateProductRepository(db)
	categoryRepo := repository.CreateCategoryRepository(db)
	userRepo := repository.CreateUserRepository(db)

	mux := controller.NewRouter(GoogleAuthConfig, OAuthStateString, categoryRepo, productRepo, userRepo)

	req := httptest.NewRequest("GET", "/api/v1/categories", nil)
	res := httptest.NewRecorder()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	result_byte := res.Body.Bytes()

	var responseBody helpers.JsonResponse

	err = json.Unmarshal(result_byte, &responseBody)

	assert.NoError(t, err)

	listOfCategory := responseBody.Data.([]interface{})
	assert.Equal(t, 3, len(listOfCategory))

}

func TestGetAllProductByCategory(t *testing.T) {
	err := godotenv.Load("./../.env")

	if err != nil {
		assert.NoError(t, err, "Cannot load .env file")
	}

	GoogleAuthConfig = &oauth2.Config{
		RedirectURL:  GoogleRedirectURL,
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
		assert.NoError(t, err, "Cannot open database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		assert.NoError(t, err, "Cannot ping database")
	}
	fmt.Println("Database Connected!")

	helpers.MigrateDB(db)
	productRepo := repository.CreateProductRepository(db)
	categoryRepo := repository.CreateCategoryRepository(db)
	userRepo := repository.CreateUserRepository(db)

	mux := controller.NewRouter(GoogleAuthConfig, OAuthStateString, categoryRepo, productRepo, userRepo)

	req := httptest.NewRequest("GET", "/api/v1/products?category=makanan", nil)
	res := httptest.NewRecorder()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	result_byte := res.Body.Bytes()

	var responseBody helpers.JsonResponse

	err = json.Unmarshal(result_byte, &responseBody)

	assert.NoError(t, err)

	listOfProduct := responseBody.Data.([]interface{})
	assert.Equal(t, 14, len(listOfProduct))
}
