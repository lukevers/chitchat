# chitchat

A basic websocket-based chat service written in Go

## The Task

ASAPP's server infrastructures must scale to support companies with hundreds of millions of customers. Scaling realtime communication across multiple servers and persisting their conversation logs involves a number of pretty interesting engineering challenges. In order to tackle these challenges we rely on our dev team members' ability to reason about and design such systems, both conceptually and concretely, from a bird's eye perspective as well as in great detail.

Your challenge is to design and implement a basic websocket-based chat service, and then prepare to talk through what the challenges would be in scaling it up and how you would tackle those challenges.

Take the time that you need, but fast is better than slow.

Conceptually, the service should have five high-level capabilties:

 - [ ] Connect
 - [ ] Login
 - [ ] Send Message
 - [ ] Receive Message (push)
 - [ ] Fetch Message History

You shouldn't expect to complete everything - just like in a startup. Choose wisely, use your time efficiently, and do just enough to get the job done: demonstrate that you're capable.

If we mutually agree to proceed then I will ask you to come in and continue work on this challenge alongside the rest of the team in our office.


### Misc thoughts and recommendations before we go into details

- Use the programming languages and tools that you're most familiar with.
- I recommend using open source libraries rather than reinvent the wheel. Examples that we use for our go server:
    github.com/gorilla/websocket, github.com/go-sql-driver/mysql, github.com/garyburd/redigo,
    code.google.com/p/go.crypto/bcrypt, https://github.com/marcuswestin/FunGo/tree/master/sql
- Version control. It's a plus if your results come with a commit history.
- We rely on tests to move fast and break nothing. I recommend that you do as well.
- Avoid hosted services that aren't open source. You can't reason about their behavior, and I will look for your ability to reason about the behavior of your designs.
- Have fun! If you don't think this project sounds like fun, then working at ASAPP may not be your cup of tea :)


#### Service requirements

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
