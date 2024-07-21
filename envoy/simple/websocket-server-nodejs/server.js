const WebSocket = require('ws');

const server = new WebSocket.Server({ port: 8080 });

server.on('connection', (ws) => {
    console.log('New client connected');

    // ws.send('Welcome to the WebSocket server!');

    ws.on('message', (message) => {
        console.log(`Received: ${message}`);
        ws.send(`Echo: ${message}`);
    });

    ws.on('close', function close(code, reason) {
        console.log(`Client disconnected with code: ${code}, reason: ${reason}`);
    });

    ws.on('error', function error(error) {
        console.error(`WebSocket error: ${error.message}`);
    });
});

console.log('WebSocket server is running on ws://localhost:8080');
