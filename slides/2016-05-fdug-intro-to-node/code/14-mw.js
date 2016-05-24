var express = require('express');
var uuid = require('uuid');
var app = express();

function requestId(req, res, next) {
	if(!req.get('X-Request-Id')) req.headers['x-request-id'] = uuid.v4();
	req.requestId = req.get('X-Request-Id');
	res.setHeader('X-Request-Id', req.requestId);
	next();
}

function logger(req, res, next) {
	console.log(`${req.method} ${req.url} Request ID: ${req.requestId}`);
	next();
}

app.use(requestId);
app.use(logger);

app.get('/hello/:name', function (req, res) {
  res.send(`Hello ${req.params.name}!`);
});

app.listen(3000, function () {
  console.log('Example app listening on port 3000!');
});
