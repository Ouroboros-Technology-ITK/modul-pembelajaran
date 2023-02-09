package main

import (
	"net/http"
)

func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) // TODO: replace this
}

func main() {
	// TODO: answer here

	http.ListenAndServe("localhost:8080", nil)
}
