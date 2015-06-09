package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Add all routes
func addRoutes() {
	// Public routes
	router.GET("/signup", getSignUp)
	router.POST("/signup", postSignUp)
	router.GET("/login", getLogin)
	router.POST("/login", postLogin)
	router.GET("/logout", getLogout)

	// Authorized user routes
	private := router.Group("/")
	private.Use(LoggedIn())
	{
		private.GET("/", getIndex)
	}
}

// Get Sign Up
func getSignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup", nil)
}

// Get Login
func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", nil)
}

// Get Index
func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index", nil)
}

// Get Logout
func getLogout(c *gin.Context) {
	// Remove cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "user",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	// Redirect to login page
	c.Redirect(http.StatusFound, "/login")
}

// Post Sign Up
func postSignUp(c *gin.Context) {
	// Check to see if `username` is already in the database.
	var user string
	err := db.QueryRow("SELECT username FROM users WHERE username=?", c.PostForm("username")).Scan(&user)
	if err != sql.ErrNoRows {
		// Username is taken! Redirect back with the bad news.
		c.Redirect(http.StatusFound, "/signup")
	} else {
		// Username is available! Prepare an insert statement.
		stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
		if err != nil {
			fmt.Printf("Error preparing database statement: %s\n", err)
		}

		// Hash the password.
		password := HashPassword(c.PostForm("password"))

		// Execute the insert.
		stmt.Exec(c.PostForm("username"), password)

		// Now that we have a username/password let's login.
		c.Redirect(http.StatusFound, "/login")
	}
}

// Post Login
func postLogin(c *gin.Context) {
	// Try and get the user from the database
	var user User
	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", c.PostForm("username")).Scan(&user.username, &user.password)
	if err == sql.ErrNoRows {
		// Username does not exist, so the login isn't happening.
		c.Redirect(http.StatusFound, "/login")
	} else {
		// Username does exist, so now we want to match the passwords.
		if !PasswordAndHashMatch(c.PostForm("password"), user.password) {
			// Password does not match, so the login isn't happening.
			c.Redirect(http.StatusFound, "/login")
		} else {
			// Password matches, so now we can login.
			// Create session
			session, _ := store.New(c.Request, "user")
			session.Values["username"] = user.username

			// Save session
			session.Save(c.Request, c.Writer)

			// Redirect to our main page
			c.Redirect(http.StatusFound, "/")
		}
	}
}
