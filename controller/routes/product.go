package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"net/http"
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
