package middleware

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/models"
	"azura-lab-intern/study-case-1/repository"
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
)

func Method(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			helpers.ErrorResponseJSON(w, "Method is Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Authorization(googleOAuthConfig *oauth2.Config, userRepo *repository.UserRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearer := r.Header.Get("Authorization")

			if !strings.Contains(bearer, "Bearer") {
				log.Println("Error : Bearer Token Not Found")
				helpers.ErrorResponseJSON(w, "Bearer Token Not Found", http.StatusBadRequest)
				return
			}

			tokenString := strings.Replace(bearer, "Bearer ", "", -1)
			if len(strings.Split(tokenString, ".")) != 3 {
				response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tokenString)

				if response.StatusCode == 401 {
					log.Println("Invalid OAuth Token")
					helpers.ErrorResponseJSON(w, "Invalid Token", http.StatusUnauthorized)
					return

				}
				if err != nil {
					log.Println("Error on requesting user info: ", err.Error())
					helpers.ErrorResponseJSON(w, "Error on requesting user info: "+err.Error(), http.StatusUnauthorized)
					return
				}
				defer response.Body.Close()
				content, err := ioutil.ReadAll(response.Body)

				if err != nil {
					log.Println("Failed reading user info: ", err.Error())
					helpers.ErrorResponseJSON(w, "Failed user info: "+err.Error(), http.StatusInternalServerError)
					return
				}

				var content_map map[string]interface{}
				json.Unmarshal(content, &content_map)
				var user *models.User
				user, err = userRepo.GetUserByEmail(content_map["email"].(string))

				if err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						newUser := models.User{
							Name:  content_map["name"].(string),
							Email: content_map["email"].(string),
						}

						user, _ = userRepo.InsertUser(newUser)
					} else {
						log.Println("Error while getting user : ", err.Error())
						helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
						return
					}
				}

				ctx := context.WithValue(r.Context(), "user_data", user)
				r = r.WithContext(ctx)

			} else {
				tokenHeader := strings.Split(tokenString, ".")[0]
				decotedByte, err := base64.StdEncoding.DecodeString(tokenHeader)
				if err != nil {
					log.Println("Failed to decode base64: ", err.Error())
					helpers.ErrorResponseJSON(w, "Invalid Token", http.StatusUnauthorized)
					return
				}
				var tokenHeaderMap map[string]interface{}

				json.Unmarshal(decotedByte, &tokenHeaderMap)
				if tokenHeaderMap["typ"].(string) != "JWT" {
					log.Println("Token type doesn't match")
					helpers.ErrorResponseJSON(w, "Invalid Token", http.StatusUnauthorized)
					return
				}

				userClaims := helpers.UserClaims{}
				token, err := jwt.ParseWithClaims(tokenString, &userClaims, func(token *jwt.Token) (interface{}, error) {
					if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("signing method invalid")
					} else if method != jwt.SigningMethodHS256 {
						return nil, fmt.Errorf("signing method invalid")
					}

					return []byte(os.Getenv("JWT_SECRET_KEY")), nil
				})
				if err != nil {
					log.Println("Error on parsing claim: ", err.Error())
					helpers.ErrorResponseJSON(w, "Invalid Token", http.StatusUnauthorized)
					return
				}
				if !token.Valid {
					log.Println("Invalid Token")
					helpers.ErrorResponseJSON(w, "Invalid Token", http.StatusUnauthorized)
					return
				}
				user := userClaims.UserData

				ctx := context.WithValue(r.Context(), "user_data", user)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}
}
