package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/onelzyugy/projects/golang-microservice/handlers"

	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
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
	// swagger:operation GET /health general Health
	// Check if the service is running.
	// Returns 204 without content
	// ---
	// responses:
	//     '204':
	//         description: Service available
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods(http.MethodGet, http.MethodOptions) // GET
	r.HandleFunc("/bird", handlers.GetBirdHandler).Methods(http.MethodGet, http.MethodOptions)       // GET
	// swagger:operation POST /oauth/token auth GetToken
	// Authenticate and authorise with Windows credentials.
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: grant_type
	//   in: formData
	//   description: Grant type
	//   required: true
	//   type: string
	//   default: "password"
	// - name: username
	//   in: formData
	//   description: domain\username
	//   type: string
	//   required: true
	// - name: password
	//   in: formData
	//   description: Your password
	//   required: true
	//   type: string
	//   format: password
	// responses:
	//     '200':
	//         description: Authenticated
	//     '400':
	//         description: Bad request
	//     '500':
	//         description: Internal server error
	r.HandleFunc("/bird", handlers.CreateBirdHandler).Methods(http.MethodPost, http.MethodOptions)                        // POST
	r.HandleFunc("/add-todo", handlers.AddTodoHandler).Methods(http.MethodPost, http.MethodOptions)                       // POST
	r.HandleFunc("/retrieve-todo", handlers.RetrieveTodoHandler).Methods(http.MethodGet, http.MethodOptions)              // GET
	r.HandleFunc("/order", handlers.OrderBubbleTeaHandler).Methods(http.MethodPost, http.MethodOptions)                   // POST
	r.HandleFunc("/retrieve-order", handlers.RetrieveOrderedBubbleTeaHandler).Methods(http.MethodGet, http.MethodOptions) // GET
	r.HandleFunc("/health-check", handlers.HealthCheck).Methods(http.MethodGet, http.MethodOptions)
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
