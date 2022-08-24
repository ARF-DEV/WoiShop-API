package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

func GetAllProductByCategory(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")

		if len(category) < 1 {
			log.Println("error product by category : category query is not found")
			helpers.ErrorResponseJSON(w, "Category query is required", http.StatusBadRequest)
			return
		}

		results, err := productRepo.GetProductByCategory(category)

		if err != nil {
			log.Println("error product by category : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "success getting product by category", results)
	})
}
func GetAllProduct(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := productRepo.GetAllProduct()

		if err != nil && err != sql.ErrNoRows {
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success getting all product", result)
	})
}

func GetProductByID(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		if len(id) < 1 {
			log.Println("Error product by id : id query not found")
			helpers.ErrorResponseJSON(w, "id query required", http.StatusBadRequest)
			return
		}

		id_int, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, "Invalid id query", http.StatusBadRequest)
			return
		}

		result, err := productRepo.GetProductByID(id_int)

		if err != nil && err != sql.ErrNoRows {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success getting product", result)
	})
}
