package main

import (
	"fmt"
	"net/http"
	"os"
)

func MethodGet(r *http.Request) error {
	if r.Method != http.MethodGet {
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

func CheckDataRequest(r *http.Request) error {
	data := r.URL.Query().Get("data")
	if len(data) == 0 {
		return fmt.Errorf("Data not found")
	}
	return nil
}

func CheckOpenFile(r *http.Request) error {
	filename := r.URL.Query().Get("filename")
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("File not found")
	}
	return nil
}

func MethodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := MethodGet(r)
		fmt.Println(err) // TODO: replace this
	}
}

func DataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckDataRequest(r)
		fmt.Println(err) // TODO: replace this
	}
}

func OpenFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckOpenFile(r)
		fmt.Println(err) // TODO: replace this
	}
}
