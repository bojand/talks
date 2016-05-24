const http = require('http');

const server = http.createServer(function (req, res) {
  res.end('<h1>Hello World</h1>');
});

server.listen(8001);
