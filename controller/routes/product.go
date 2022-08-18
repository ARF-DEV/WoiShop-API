package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"log"
	"net/http"
	"strconv"
)

func GetAllProduct(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := productRepo.GetAllProduct()

		if err != nil {
			helpers.ErrorResponseJSON(w, err.Message, http.StatusInternalServerError)
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

		result, res_err := productRepo.GetProductByID(id)

		if res_err != nil {
			log.Println("Error product by id : ", res_err.Message)
			helpers.ErrorResponseJSON(w, res_err.Message, http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success getting product", result)
	})
}
