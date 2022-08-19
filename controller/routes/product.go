package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetAllProductByCategory(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")

		if len(category) < 1 {
			log.Println("error product by category : invalid category \"", category, "\"")
			helpers.ErrorResponseJSON(w, fmt.Sprintf("error product by category : invalid category \"%s\"", category), http.StatusBadRequest)
			return
		}

		results, err := productRepo.GetProductByCategory(category)

		if err != nil {
			log.Println("error product by category : ", err.Error())
			helpers.ErrorResponseJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "success getting product by category", results)
	})
}
func GetAllProduct(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := productRepo.GetAllProduct()

		if err != nil {
			helpers.ErrorResponseJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success getting all product", result)
	})
}

func GetProductByID(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := productRepo.GetProductByID(id)

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success getting product", result)
	})
}
