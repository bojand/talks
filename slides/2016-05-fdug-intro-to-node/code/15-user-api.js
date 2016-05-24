var hl = require('highland');
var level = require('level');
var _ = require('lodash');
var db = level('./api-db', { valueEncoding: 'json' });

exports.get = function (req, res, next) {
  db.get(req.params.email, function (err, value) {
    if (err) {
      if (err.notFound) {
        return res.sendStatus(404);
      }
      return res.status(500).send(err.message);
    }
    return res.send(value);
  });
}

exports.create = function (req, res, next) {
  var key = req.body.email;
  if (!key) {
    return res.status(400).send('Email required.');
  }

  db.put(key, req.body, function (err) {
    if (err) {
      return res.status(500).send(err);
    }
    return res.status(201).send(req.body);
  });
}

exports.update = function (req, res, next) {
  var key = req.params.email;
  db.get(req.params.email, function (err, value) {
    if (err) {
      if (err.notFound) {
        return res.sendStatus(404);
      }
      return res.status(500).send(err.message);
    }

    var data = req.method === 'PUT' ? req.body : _.merge({}, req.body, value);
    db.put(key, data, function (err) {
      if (err) {
        return res.status(500).send(err);
      }
      return res.status(200).send(data);
    });
  });
}

exports.delete = function (req, res, next) {
  db.del(req.params.email, function (err) {
    if (err) {
      if (err.notFound) {
        return res.sendStatus(404);
      }
      return res.status(500).send(err.message);
    }
    res.sendStatus(200);
  })
}

exports.list = function (req, res, next) {
  res.setHeader('content-type', 'application/json');
  var dbstream = db.createValueStream({ valueEncoding: 'utf8' });
  var dataStream = hl(dbstream).intersperse(',');
  hl(['[']).concat(dataStream).concat(hl([']'])).pipe(res);
}
