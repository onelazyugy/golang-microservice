package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/onelzyugy/projects/golang-microservice/handlers"
)

const (
	port = "PORT"
)

func main() {
	log.Println("starting...")
	r := newRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	log.Println("router setup completed...")
	err := http.ListenAndServe(fmt.Sprintf(":%s", getPort()), r)

	if err != nil {
		panic(err.Error())
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods(http.MethodGet, http.MethodOptions)                      // GET
	r.HandleFunc("/bird", handlers.GetBirdHandler).Methods(http.MethodGet, http.MethodOptions)                            // GET
	r.HandleFunc("/bird", handlers.CreateBirdHandler).Methods(http.MethodPost, http.MethodOptions)                        // POST
	r.HandleFunc("/add-todo", handlers.AddTodoHandler).Methods(http.MethodPost, http.MethodOptions)                       // POST
	r.HandleFunc("/retrieve-todo", handlers.RetrieveTodoHandler).Methods(http.MethodGet, http.MethodOptions)              // GET
	r.HandleFunc("/order", handlers.OrderBubbleTeaHandler).Methods(http.MethodPost, http.MethodOptions)                   // POST
	r.HandleFunc("/retrieve-order", handlers.RetrieveOrderedBubbleTeaHandler).Methods(http.MethodGet, http.MethodOptions) // GET
	return r
}

func getPort() string {
	port, found := syscall.Getenv(port)
	if !found {
		port = "8080"
	}
	log.Println("port running at: ", port)
	return port
}
