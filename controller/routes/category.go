package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"net/http"
)

func GetAllCategory(categoryRepo *repository.CategoryRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		results, err := categoryRepo.GetAllCategory()

		if err != nil {
			helpers.ErrorResponseJSON(w, err.Message, http.StatusInternalServerError)
		}

		helpers.SuccessResponseJSON(w, "Success getting all category", results)
	})
}
