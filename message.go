package main

type Message struct {
	Sender   string
	Receiver string
	Message  string
}

type Event struct {
	Type    string
	Message Message
}
