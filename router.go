package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var router *gin.Engine

func route() {
	// If we're in production mode don't run gin in develop
	if *production {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create our web server
	router = gin.Default()

	// Compile html templates
	router.LoadHTMLGlob("app/html/*.html")

	// Add routes
	addRoutes()

	// Add our static file routes
	router.Static("/", "./public/")

	// Figure out host:port
	addr := fmt.Sprintf("%s:%d", *host, *port)

	// Create our server based off of net/http so we can
	// control it more
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Listen and serve
	s.ListenAndServe()
}

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

	// Websocket routes
	router.GET("/ws", handleWsRoute)
}
