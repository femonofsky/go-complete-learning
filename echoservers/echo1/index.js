const express = require('express');
let app = express();
let config = require('./config')

app.all('*', (req, res) => {
    res.send("Alive !");
});

// app.get('/', (req, res) => res.send('Hello World!'))
app.listen(config.port, () => console.log('Server listening on ', config.port));