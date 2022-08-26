package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"errors"
	"log"
	"net/http"
)

func GetAllCategory(categoryRepo *repository.CategoryRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		results, err := categoryRepo.GetAllCategory()

		if err != nil {
			log.Println("Error on Get All Category : ", err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				helpers.ErrorResponseJSON(w, "Not Found", http.StatusOK)
				return
			}
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success getting all category", results)
	})
}
