package routes

import (
	"azura-lab-intern/study-case-1/repository"
	"net/http"
)

func Login(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func Register(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
