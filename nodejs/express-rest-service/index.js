const express = require('express');
const app = express();
const port = 3000;

// Middleware to parse JSON bodies
app.use(express.json());

// Route for /foo
app.get('/foo', (req, res) => {
  res.send('GET request to /foo');
});

app.post('/foo', (req, res) => {
  const data = req.body;
  res.send(`POST request to /foo with data: ${JSON.stringify(data)}`);
});

// Route for /foo/:id
app.get('/foo/:id', (req, res) => {
  const id = req.params.id;
  res.send(`GET request to /foo/${id}`);
});

app.put('/foo/:id', (req, res) => {
  const id = req.params.id;
  const data = req.body;
  res.send(`PUT request to /foo/${id} with data: ${JSON.stringify(data)}`);
});

app.delete('/foo/:id', (req, res) => {
  const id = req.params.id;
  res.send(`DELETE request to /foo/${id}`);
});

app.listen(port, () => {
  console.log(`App listening at http://localhost:${port}`);
});
