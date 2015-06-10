package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type Message struct {
	Sender   string
	Receiver string
	Message  string
	Original bool
}

type Event struct {
	Type    string
	Message Message
}

func getOldMessages(c *gin.Context) {
	user1 := GetUser(c.Param("user1"))
	user2 := GetUser(c.Param("user2"))

	var messages []Message
	rows, err := db.Query("SELECT sender, receiver, message FROM messages WHERE sender=? AND receiver=? OR SENDER=? AND receiver=?", user1.Id, user2.Id, user2.Id, user1.Id)
	if err != nil {
		fmt.Println("Error trying to get messages: %s", err)
	}
	
	defer rows.Close()
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.Sender, &message.Receiver, &message.Message); err == nil {
			// Get a string representation of the sender
			if message.Sender == fmt.Sprintf("%d", user1.Id) {
				// Sender is user1
				message.Sender = user1.Username
			} else {
				// Sender is user2
				message.Sender = user2.Username
			}

			// Get a string representation of the receiver
			if message.Receiver == fmt.Sprintf("%d", user1.Id) {
				// Sender is user1
				message.Receiver = user1.Username
			} else {
				// Sender is user2
				message.Receiver = user2.Username
			}

			// Append to the slice of messages
			messages = append(messages, message)
		}
	}

	c.JSON(200, messages)
}
