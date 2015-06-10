# chitchat

A basic websocket-based chat service written in Go

## Tasks

Conceptually, the service should have five high-level capabilties:

 - [x] Connect
 - [x] Login
 - [ ] Send Message
 - [ ] Receive Message (push)
 - [ ] Fetch Message History

### Service requirements

- Connect to server

    The server should accept HTTP requests and upgrade them to websocket connections.

- "Login" with username

    A connected client should be able to authenticate with a given username.
    Don't worry about passwords at this point.

- Send Message to another username

    An authenticated connection should be able to send a text message to a given username.
    Messages should be persisted to a database.

- Receive Messages in realtime

    An authenticated connection should receive messages sent to its authenticated username.
    (What happens if multiple connections are authenticated with the same username?)

- Fetch Message History

    Mobile apps go offline and online all the time.
    An authenticated connection should be able to fetch missed messages.
