package middleware

import (
	"log"
	"net/http"
)

func MiddlewareTest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware")
		next(w, r)
	}
}
