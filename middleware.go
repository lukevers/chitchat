package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Gin middleware to check if the user is logged in or not.
// The user is then directed to the login page if they are
// not currently logged in.
func LoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get session
		session, err := store.Get(c.Request, "user")
		if err != nil {
			fmt.Println("Could not check for session: %s", err)
		}

		// Check to see if session exists or not
		if session.IsNew {
			c.Redirect(http.StatusFound, "/login")
		}

		// If we get this far then we are logged in
		c.Next()
	}
}
