package main

import (
	"net/http"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
