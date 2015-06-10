package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
	"time"
)

// Session store
var store *sessions.FilesystemStore

// Database structure of the users table
type User struct {
	Id        uint
	Username  string
	password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Get a User struct from the database by a username string
func GetUser(username string) (user User) {
	db.QueryRow("SELECT id, username, password, created_at, updated_at FROM users WHERE username=?", username).
		Scan(&user.Id, &user.Username, &user.password, &user.CreatedAt, &user.UpdatedAt)
	return
}

// Get all users
func GetAllUsers() (users []User) {
	rows, err := db.Query("SELECT id, username, password, created_at, updated_at FROM users")
	if err != nil {
		fmt.Println("Error trying to get all users: %s", err)
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Username, &user.password, &user.CreatedAt, &user.UpdatedAt); err == nil {
			users = append(users, user)
		}
	}

	return
}

// HashPassword is a func that takes a string password and hashes it
// with bcrypt.
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating hash from password: %s", err)
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
	// Clean up old sessions
	allSessions, err := filepath.Glob(*sessionsPath + "/session_*")
	if err != nil {
		fmt.Println("Old sessions were not cleared: %s", err)
	} else {
		for _, s := range allSessions {
			err = os.Remove(s)
			if err != nil {
				fmt.Println("Error removing old session: %s", err)
			}
		}
	}

	// Create new session store
	store = sessions.NewFilesystemStore(
		// Path to sessions
		*sessionsPath,

		// Secret key with strength set to 64
		[]byte(securecookie.GenerateRandomKey(64)),
	)
}
