// Store the url and websocket globally
var url, ws;

// Initialize the websockets
(function init() {
	// Figure out current url with `ws` instead of http/https
	url = location.href.replace(location.protocol.slice(0, -1), 'ws');
	// Create new websocket
	ws = new WebSocket(url+'ws');

	// Bind enter key to send data
	bindEnterKeyToSend();

	// Bind on message
	ws.onmessage = receive;
})();

// Bind enter key to send data
function bindEnterKeyToSend() {
	$('#message').keypress(function(e) {
		// Only if the enter key was pressed
		if (e.which == 13) {
			// First let's get the message
			var message = $(this).val();
			// Then let's clear the input since we are sending it
			$(this).val("");
			// Send message to server
			send(message);
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
