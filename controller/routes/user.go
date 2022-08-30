package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/models"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenBody struct {
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

		helpers.SuccessResponseJSON(w, "Login Success", TokenBody{Token: tokenString})
	})
}

func Register(userRepo *repository.UserRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println("Error on Register: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		expTime := time.Now().Add(time.Minute * 10)
		tokenString, key, err := helpers.GenerateRegisterOTPClaims(u, expTime)

		if err != nil {
			log.Println("Error on Register: ", err.Error())
			helpers.ErrorResponseJSON(w, "Failed Generating Token", http.StatusInternalServerError)
			return
		}

		//Kirim kode OTP ke email
		err = helpers.SendOTPEmail(u.Email, key)

		if err != nil {
			log.Println("Error while sending Email", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success Generating OTP Token", OTPTokenBody{
			Token: tokenString,
		})
	})
}

func VerifyOTPRegister(userRepo *repository.UserRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var t OTPTokenBody

		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			log.Println("Error while parsing JSON: ", err.Error())
			helpers.ErrorResponseJSON(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if len(t.Token) < 1 {
			log.Println("Error : OTP Token Guess is Empty", http.StatusBadRequest)
			helpers.ErrorResponseJSON(w, "OTP Token Guess Must not be Empty", http.StatusBadRequest)
			return
		}
		AuthHeader := r.Header.Get("Authorization")

		if !strings.Contains(AuthHeader, "OTP") {
			log.Println("Error while getting OTP token : OTP Token Not Found")
			helpers.ErrorResponseJSON(w, "OTP Token Not Found", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(AuthHeader, "OTP ", "", -1)

		claims := helpers.RegisterOTPClaims{}

		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}

			return []byte(os.Getenv("JWT_OTP_KEY")), nil
		})

		if err != nil {
			log.Println("Error while parsing claims: ", err.Error())
			helpers.ErrorResponseJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Println("Error : Token is Invalid")
			helpers.ErrorResponseJSON(w, "Unautherized", http.StatusUnauthorized)
			return
		}

		if t.Token != claims.OTPSecret {
			log.Println("Error : OTP Token didn't Match")
			helpers.ErrorResponseJSON(w, "Invalid OTP Token", http.StatusUnauthorized)
			return
		}

		err = userRepo.InsertUser(claims.UserData)

		if err != nil {
			log.Println("Error while inserting user: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success Registering User Using OTP", claims.UserData)

	})
}
