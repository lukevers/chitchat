package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// Parse flags
	kingpin.Parse()

	// Connect to database
	dbconnect()

	// Setup router and webserver
	route()
}
