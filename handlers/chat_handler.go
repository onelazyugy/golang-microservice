package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel, channel is use communicate btw go routines, msg send to this channel and it will send to other routine
// Configure the upgrader
var upgrader = websocket.Upgrader{}

// Message our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
	DateTime string `json:"dateTime"`
}

// HandleConnections handle chat connection
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

// HandleMessages handle incoming messages
func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast //BLOCKING CALL HERE
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg) //return what is send in by the user to the UI
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

/*
clients is a map of all connected clients
broadcast is a channel that receive messages and send it out to all connected clients
everytime there is a new message comes in, the HandleConnections function will send that message to the broadcast channel
the go routine will run HandleMessages function and get all of the messages from the broadcast channel and send it to all of the client
*/
