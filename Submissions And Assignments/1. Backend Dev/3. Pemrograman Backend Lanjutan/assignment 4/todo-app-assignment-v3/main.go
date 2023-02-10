package main

import (
	"fmt"
	"net/http"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func Logout(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))
	// TODO: answer here
}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	db.Task = map[string][]model.Todo{}
	w.WriteHeader(http.StatusOK)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
	mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
	mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))

	// TODO: answer here

	mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
