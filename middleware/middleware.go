package middleware

import (
	"azura-lab-intern/study-case-1/helpers"
	"net/http"
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
