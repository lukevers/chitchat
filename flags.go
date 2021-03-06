package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	host         = kingpin.Flag("host", "Host to bind webserver to").Default("127.0.0.1").IP()
	port         = kingpin.Flag("port", "Port to bind webserver to").Default("2015").Int()
	production   = kingpin.Flag("production", "Run in production mode").Default("false").Bool()
	database     = kingpin.Flag("database", "Database connection string").Default("homestead:secret@tcp(127.0.0.1:33060)/chitchat").String()
	sessionsPath = kingpin.Flag("sessions-path", "Folder for sessions").Default("app/sessions").String()
)
