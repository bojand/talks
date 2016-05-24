var express = require('express');
var bodyParser = require('body-parser');
var morgan = require('morgan');
var errorHander = require('errorhandler');
var addRequestId = require('express-request-id')();
var api = require('./15-user-api');

var app = express();

// set up id token for logging
morgan.token('id', function getId(req) {
  return req.id
});

app.use(bodyParser.json());
app.use(addRequestId);
app.use(morgan(':method :url :status :response-time ms - :id'));
if (process.env.NODE_ENV === 'development') {
  // only use in development
  app.use(errorhandler());
}

app.get('/users', api.list);
app.get('/users/:email', api.get);
app.post('/users', api.create);
app.put('/users/:email', api.update);
app.patch('/users/:email', api.update);
app.delete('/users/:email', api.delete);

app.listen(3000, function () {
  console.log('User API listening on port 3000!');
});
