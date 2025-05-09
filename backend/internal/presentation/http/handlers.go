package handlers

import (
	"log"
	"net/http"
)

func HandlerTest(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler")
	w.WriteHeader(http.StatusOK)
}
