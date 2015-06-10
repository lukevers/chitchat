package main

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
