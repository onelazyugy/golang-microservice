package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onelzyugy/projects/golang-microservice/handlers"
)

func main() {
	r := newRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	err := http.ListenAndServe(":8181", r)

	if err != nil {
		panic(err.Error())
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods(http.MethodGet, http.MethodOptions)         //GET
	r.HandleFunc("/bird", handlers.GetBirdHandler).Methods(http.MethodGet, http.MethodOptions)               //GET
	r.HandleFunc("/bird", handlers.CreateBirdHandler).Methods(http.MethodPost, http.MethodOptions)           // POST
	r.HandleFunc("/add-todo", handlers.AddTodoHandler).Methods(http.MethodPost, http.MethodOptions)          // POSt
	r.HandleFunc("/retrieve-todo", handlers.RetrieveTodoHandler).Methods(http.MethodGet, http.MethodOptions) //GET
	return r
}
