package main

import (
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

var ws = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Username string
	Message  string
}

func handleWsRoute(c *gin.Context) {
	wshandler(c.Writer, c.Request)
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade websocket: %s", err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		session, err := store.Get(r, "user")
		if err != nil {
			fmt.Println("Error trying to get session: %s", err)
		}

		conn.WriteJSON(&Message{
			Username: session.Values["username"].(string),
			Message: string(msg),
		})
	}
}
