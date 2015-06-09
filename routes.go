package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
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
	// TODO
	// - logout
	// - redirect to login
}

// Post Sign Up
func postSignUp(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username)
	fmt.Println(password)
}

// Post Login
func postLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username)
	fmt.Println(password)
}
