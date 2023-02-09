package main

import (
	"net/http"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) // TODO: replace this
}

func main() {
	// TODO: answer here
	http.ListenAndServe("localhost:8080", nil)
}
