package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type OTPTokenBody struct {
	Token string `json:"otp_token"`
}

type OAuthCodeBody struct {
	State string `json:"state"`
	Code  string `json:"code"`
}

func LoginOAuth(GoogleAuthConfig *oauth2.Config, OAuthStateString string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := GoogleAuthConfig.AuthCodeURL(OAuthStateString)

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})
}

func HandleOAuthCallBack(GoogleAuthConfig *oauth2.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Balikin Codenya dalam response

		helpers.SuccessResponseJSON(w, "Success to Login with Google", OAuthCodeBody{
			State: r.FormValue("state"),
			Code:  r.FormValue("code"),
		})
	})
}

func GetAccessToken(GoogleConfig *oauth2.Config, OAuthStateString string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var b OAuthCodeBody
		defer r.Body.Close()
		json.NewDecoder(r.Body).Decode(&b)

		if b.State != OAuthStateString {
			log.Println("Error : Invalid OAuth State")
			helpers.ErrorResponseJSON(w, "Invalid OAuth State", http.StatusBadRequest)
			return
		}

		if len(b.Code) < 1 {
			log.Println("Error : OAuth Code is not Found")
			helpers.ErrorResponseJSON(w, "OAuth Code is Not Found", http.StatusBadRequest)
		}

		token, err := GoogleConfig.Exchange(context.Background(), b.Code)

		if err != nil {
			log.Println("Code Exchange failed: ", err.Error())
			helpers.ErrorResponseJSON(w, "Code Exchage failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success requesting access token", token)
	})
}
