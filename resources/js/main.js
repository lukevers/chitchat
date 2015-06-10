// Store the url and websocket globally
var url, ws;

// Initialize the websockets
(function init() {
	// Figure out current url with `ws` instead of http/https
	url = location.href.replace(location.protocol.slice(0, -1), 'ws');

	// Create new websocket
	ws = new WebSocket(url+'ws');

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
			// Switch tabs
			// TODO
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
			if (message !== "") {
				// Then let's clear the input since we are sending it
				$(this).val("");
				// Send message to server
				send(message);
			}
		}
	});
}

// Build HTML message
function buildMessage(username, message) {
	return '<div class="msg"><div class="username">'+username+'</div><div class="message">'+message+'</div></div>';
}

// Send data
function send(data) {
	ws.send(data);
}

// Receive data
function receive(data) {
	data = JSON.parse(data.data);
	console.log(data);
	$('#messages').append(buildMessage(data.Username, data.Message));
}
