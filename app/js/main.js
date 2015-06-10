// Store the url and websocket globally
var url, ws;

// Initialize the websockets
(function init() {
	// Figure out current url with `ws` instead of http/https
	url = location.href.replace(location.protocol.slice(0, -1), 'ws');

	// Create new websocket
	ws = new WebSocket(url+'ws');

	// Add own class to autged user
	AddClassToAuthedUser();

	// Load all old messages
	LoadMessages();

	// Bind changing message category
	bindMessageSwitch();

	// Bind enter key to send data
	bindEnterKeyToSend();

	// Bind on message
	ws.onmessage = receive;
})();

// Bind message window switch with other users
function bindMessageSwitch() {
	$('.users ul li').bind('click', function(e) {
		$this = $(this);
		// Don't do anything if you click the active tab
		if (!$this.hasClass('active')) {
			// Remove active class on other tabs
			$('.users ul li.active').removeClass('active');
			// Set the clicked one to active
			$this.addClass('active');
			// Remove the active class from the active message window
			$('.messages.active').removeClass('active');
			// Add the active class to our active message window
			$('#messages-'+$this.data('user')).addClass('active');
			// Set in our active object
			active.messages = $this.data('user');
		}
	});
}

// Bind enter key to send data
function bindEnterKeyToSend() {
	$('#message').keypress(function(e) {
		// Only if the enter key was pressed
		if (e.which == 13) {
			// First let's get the message
			var message = $(this).val().trim();
			// Make sure message isn't empty
			if (message !== '') {
				if (active.messages === '') {
					alert('You must be in a conversation to send messages');
				} else {
					// Then let's clear the input since we are sending it
					$(this).val('');
					// Send message to server
					sendMessage(message);
				}
			}
		}
	});
}

// Build HTML message
function buildMessage(username, message) {
	return '<div class="msg"><div class="username">'+username+'</div><div class="message">'+message+'</div></div>';
}

// Send data
function sendMessage(message) {
	ws.send(JSON.stringify({
		Type: "message",
		Message: {
			Sender: "",
			Receiver: active.messages,
			Message: message
		}
	}));
}

// Receive data
function receive(data) {
	data = JSON.parse(data.data);
	console.log(data);

	if (data.Original) {
		$('#messages-'+data.Receiver).append(buildMessage(data.Sender, data.Message));
	} else {
		$('#messages-'+data.Sender).append(buildMessage(data.Sender, data.Message));
	}
}

// Load Messages
function LoadMessages() {
	$('.users ul li.user').map(function(i, v) {
		var user = $(v).data('user');
		$.getJSON('/messages/'+active.user+'/'+user, function(data) {
			data.map(function(msg) {
				console.log(msg);
				if (msg.Sender === active.user) {
					$('#messages-'+msg.Receiver).append(buildMessage(msg.Sender, msg.Message));
				} else {
					$('#messages-'+msg.Sender).append(buildMessage(msg.Sender, msg.Message));
				}
			});
		});
	});
}

//
function AddClassToAuthedUser() {
	$('.users ul li.user[data-user="'+active.user+'"]').addClass('authed');
}
