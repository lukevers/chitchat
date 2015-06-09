package main

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"time"
)

var store = sessions.NewFilesystemStore(
	// Path to sessions
	*sessionsPath,

	// Secret key with strength set to 64
	[]byte(securecookie.GenerateRandomKey(64)),
)

type User struct {
	id        uint
	username  string
	password  string
	createdAt time.Time
	updatedAt time.Time
}
