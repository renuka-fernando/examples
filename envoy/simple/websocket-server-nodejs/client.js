const WebSocket = require('ws');

// Create a WebSocket client and connect to the server
const ws = new WebSocket('ws://localhost:8000/ws');

// Event listener for when the connection is open
ws.on('open', function open() {
    console.log('Connected to the WebSocket server');

    // Send a message to the server
    ws.send('Hello Server!');

    // Sleep for 2 seconds and close the connection
    setTimeout(() => {
        ws.close(1000, 'Normal closure');
    }, 2000);
});

// Event listener for when the connection is closed
ws.on('close', function close(code, reason) {
    console.log(`WebSocket connection closed with code: ${code}, reason: ${reason}`);
});

// Event listener for any errors
ws.on('error', function error(error) {
    console.error(`WebSocket error: ${error.message}`);
});

// Event listener for incoming messages
ws.on('message', function incoming(message) {
    console.log(`Received: ${message}`);
});
