package main

import (
	"miniapp/internal/middleware"
	handlers "miniapp/internal/presentation/http"
	"net/http"
)

func main() {
	http.HandleFunc("/", middleware.MiddlewareTest(handlers.HandlerTest))

	http.ListenAndServe(":8080", nil)
}
