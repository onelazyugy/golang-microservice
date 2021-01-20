package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/onelzyugy/projects/golang-microservice/services"
	"github.com/onelzyugy/projects/golang-microservice/types"
)

// AddTodoHandler add todo itmes
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var todo types.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Printf("todo: %v", todo)
}

// RetrieveTodoHandler retrieve all todo items
func RetrieveTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	services.AddTodoItem()
	todos := services.GetTodoItems()
	fmt.Println(todos)
	todoBytes, _ := json.Marshal(todos)
	w.Write(todoBytes)
}

// HealthCheckHandler health check
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintf(w, "{\"status\": \"Status Up!\"}")
}

// GetBirdHandler get all birds
func GetBirdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bird1 := types.Bird{Species: "bird 1", Description: "bird 1 description"}
	bird2 := types.Bird{Species: "bird 2", Description: "bird 2 description"}
	birds := []types.Bird{}
	birds = append(birds, bird1)
	birds = append(birds, bird2)

	birdList := services.RetrieveBirds(birds)

	birdListBytes, err := json.Marshal(birdList)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(birdListBytes)
}

// CreateBirdHandler create a bird
func CreateBirdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a new instance of Bird
	bird := types.Bird{}

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
	// birds = append(birds, bird)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
