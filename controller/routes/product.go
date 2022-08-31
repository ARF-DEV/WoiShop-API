package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAllProductByCategory(productRepo *repository.ProductRepository) http.HandlerFunc {
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

		if len(results) < 1 {
			helpers.ErrorResponseJSON(w, "Not Found", http.StatusOK)
			return
		}

		helpers.SuccessResponseJSON(w, "success getting product by category", results)
	})
}
func GetAllProduct(productRepo *repository.ProductRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")

		if len(category) < 1 {

			result, err := productRepo.GetAllProduct()

			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					helpers.ErrorResponseJSON(w, "Not Found", http.StatusOK)
					return
				}
				helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			helpers.SuccessResponseJSON(w, "Success getting all product", result)
		} else {
			results, err := productRepo.GetProductByCategory(category)

			if err != nil {
				log.Println("error product by category : ", err.Error())
				helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if len(results) < 1 {
				helpers.ErrorResponseJSON(w, "Not Found", http.StatusOK)
				return
			}
			helpers.SuccessResponseJSON(w, "success getting product by category", results)

		}

	})
}

func GetProductByID(productRepo *repository.ProductRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if len(id) < 1 {
			log.Println("Error product by id : id query not found")
			helpers.ErrorResponseJSON(w, "id query required", http.StatusBadRequest)
			return
		}

		id_int, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, "Invalid id query", http.StatusBadRequest)
			return
		}

		result, err := productRepo.GetProductByID(id_int)

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				helpers.ErrorResponseJSON(w, "Not Found", http.StatusOK)
				return
			}
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success getting product", result)
	})
}
