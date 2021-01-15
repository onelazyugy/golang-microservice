package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onelzyugy/projects/golang-microservice/handlers"
)

func main() {
	r := newRouter()
	err := http.ListenAndServe(":8181", r)

	if err != nil {
		panic(err.Error())
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/bird", handlers.GetBirdHandler).Methods("GET")
	r.HandleFunc("/bird", handlers.CreateBirdHandler).Methods("POST")
	r.HandleFunc("/add-todo", handlers.AddTodoHandler).Methods("POST")
	r.HandleFunc("/retrieve-todo", handlers.RetrieveTodoHandler).Methods("GET")
	return r
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
