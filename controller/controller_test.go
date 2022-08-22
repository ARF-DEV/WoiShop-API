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
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "study_case_1"
)

func TestGetAllProduct(t *testing.T) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
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

	mux := controller.NewMux(categoryRepo, productRepo)

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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
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

	mux := controller.NewMux(categoryRepo, productRepo)

	req := httptest.NewRequest("GET", "/api/v1/product?id=40", nil)
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
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

	mux := controller.NewMux(categoryRepo, productRepo)

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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
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

	mux := controller.NewMux(categoryRepo, productRepo)

	req := httptest.NewRequest("GET", "/api/v1/product/by/category?category=makanan", nil)
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
