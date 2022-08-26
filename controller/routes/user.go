package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/models"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func Login(userRepo *repository.UserRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			log.Println("Error on Login: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		user, err := userRepo.GetUserByEmail(u.Email)

		if err != nil {
			log.Println("Error on Login: ", err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				helpers.ErrorResponseJSON(w, "User Not Found", http.StatusOK)
				return
			}

			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		expTime := time.Now().Add(time.Minute * 30)
		tokenString, err := helpers.GenerateUserToken(*user, expTime)

		if err != nil {
			log.Println("Error on Login: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Login Success", TokenResponse{Token: tokenString})
	})
}

func Register(userRepo *repository.UserRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
