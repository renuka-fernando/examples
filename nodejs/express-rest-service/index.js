const express = require('express');
const app = express();
app.use(express.json());
const a = require('./src/app');

// GET request to the root URL
app.get('/', (req, res) => {
    res.send('Hello, World!');
});

// Example of handling a POST request
app.post('/api/data', (req, res) => {
    // Process incoming data
    res.send('Data received');
});

app.post('/api/register/put', (req, res) => {
    app.put('/api/data', (req, res) => {
        res.send('Hello, World! - PUT');
    });
    res.send({"message": "registered"})
});

const PORT = process.env.PORT || 3000; // Use the port specified in environment variable or 3000 by default
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
