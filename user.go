package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Session store
var store *sessions.FilesystemStore

// Database structure of the users table
type User struct {
	id        uint
	username  string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

// HashPassword is a func that takes a string password and hashes it
// with bcrypt.
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error generating hash from password: %s\n", err)
	}

	return string(hash)
}

// PasswordAndHashMatch is a func that checks a password and a hash to
// see if they match.
func PasswordAndHashMatch(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// setUserSessionStorage sets the session store for users
func setUserSessionStorage() {
	store = sessions.NewFilesystemStore(
		// Path to sessions
		*sessionsPath,

		// Secret key with strength set to 64
		[]byte(securecookie.GenerateRandomKey(64)),
	)
}
