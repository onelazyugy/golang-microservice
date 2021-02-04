package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/onelazyugy/golang-microservice/handlers"
)

const (
	port = "PORT"
)

func main() {
	log.Println("starting...")
	r := newRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	log.Println("router setup completed...")

	//
	// Start listening for incoming chat messages
	// this routine will wait for a message from the broadcast channel, once it receive a msg, it will send that message to all clients
	go handlers.HandleMessages() //a new thread to run the HandleMessages function, go scheduler handle which routine to run
	//

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
	r.HandleFunc("/health-check", handlers.HealthCheck).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/connect/{username}", handlers.HandleConnections).Methods(http.MethodGet, http.MethodOptions)
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
