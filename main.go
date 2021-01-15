package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onelzyugy/projects/golang-microservice/services"
)

type bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []bird

func main() {
	defer fmt.Println("server started at port 8181")
	r := newRouter()
	err := http.ListenAndServe(":8181", r)

	if err != nil {
		panic(err.Error())
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	r.HandleFunc("/add-todo", addTodoHandler).Methods("POST")
	r.HandleFunc("/retrieve-todo", retrieveTodoHandler).Methods("GET")
	return r
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {

}

func retrieveTodoHandler(w http.ResponseWriter, r *http.Request) {
	services.AddTodoItem()
	todos := services.GetTodoItems()
	fmt.Println(todos)
	todoBytes, _ := json.Marshal(todos)
	w.Write(todoBytes)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status Up!")
}

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "birds" variable to json
	bird1 := bird{Species: "bird 1", Description: "bird 1 description"}
	bird2 := bird{Species: "bird 2", Description: "bird 2 description"}
	fmt.Printf("bird1 %v", bird1)
	birds = append(birds, bird1)
	birds = append(birds, bird2)
	birdListBytes, err := json.Marshal(birds)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Bird
	bird := bird{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	birds = append(birds, bird)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
