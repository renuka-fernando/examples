const express = require('express');
const app = express();
const port = 3000;

app.use(express.static('public'));

// SSE endpoint
app.get('/events', (req, res) => {
    res.setHeader('Content-Type', 'text/event-stream');
    res.setHeader('Cache-Control', 'no-cache');
    res.setHeader('X-Accel-Buffering', 'no');
    res.setHeader('Connection', 'keep-alive');
    res.flushHeaders();

    // Send a message every 2 seconds
    const intervalId = setInterval(() => {
        const now = new Date();
        res.write(`data: ${now.toLocaleTimeString()}\n\n`);
    }, 2000);

    // Cleanup when client closes connection
    req.on('close', () => {
        clearInterval(intervalId);
        res.end();
    });
});

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});
