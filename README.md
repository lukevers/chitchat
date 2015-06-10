# chitchat

A basic websocket-based chat service written in Go

## Building

You need the following:

- [Go](http://golang.org/)
- [Node](https://nodejs.org/)
- [NPM](https://www.npmjs.com/)
- [Bower](http://bower.io/)
- [Gulp](http://gulpjs.com/)

```bash
# Build our application
go get
go build

# Install assets
npm install
bower install

# Compile/minify assets
gulp
```

## Running

With all defaults:

```bash
# Development mode
./chitchat

# Production mode
./chitchat --production
```

There are flags that you can pass, and will probably have to at least for your database.

You can view them all in `flags.go` or pass the `--help` flag to the application.

```bash
./chitchat --help
```

## Tasks

Conceptually, the service should have five high-level capabilties:

 - [x] Connect
 - [x] Login
 - [x] Send Message
 - [x] Receive Message (push)
 - [x] Fetch Message History

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
