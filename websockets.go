package main

import (
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"fmt"
)

var ws = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connections []*websocket.Conn

func handleWsRoute(c *gin.Context) {
	// Upgrade the connection to a websocket connection
	conn, err := ws.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to upgrade websocket: %s", err)
		return
	}

	// Get the session for the current connection
	session, err := store.Get(c.Request, "user")
	if err != nil {
		fmt.Println("Error trying to get session: %s", err)
	}

	// Set a session value for our connection id
	session.Values["conn_id"] = len(connections)

	// Figure out who the user is and get their db row
	user := GetUser(session.Values["username"].(string))

	fmt.Println(user)

	// Add the connection to the list
	connections = append(connections, conn)

	// Prepare insert statements for new messages
	stmt, err := db.Prepare("INSERT INTO messages(sender, receiver, message) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println("Error preparing database statement: %s", err)
	}

	for {
		var event Event
		err := conn.ReadJSON(&event)
		if err != nil {
			// The connection is being closed
			i := session.Values["conn_id"].(int)
			// Remove the connection from our slice of connections
			connections = append(connections[:i], connections[i+1:]...)
			// Now break out of the for loop since we're done
			break
		}

		// Figure out who the receiver is
		receiver := GetUser(event.Message.Receiver)

		// Add a new message into the database
		stmt.Exec(user.Id, receiver.Id, event.Message.Message)

		// Broadcast to everyone currently.
		for _, c := range connections {
			c.WriteJSON(&Message{
				Sender: user.Username,
				Message: event.Message.Message,
				Receiver: event.Message.Receiver,
			})
		}
	}
}
