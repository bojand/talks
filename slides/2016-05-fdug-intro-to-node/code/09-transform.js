var fs = require('fs');
var stream = require('stream');

var yell = new stream.Transform({
	transform: function(chunk, encoding, next) {
    this.push(chunk.toString().toUpperCase());
    next();
  }
});

var r = fs.createReadStream('file.txt');
var w = fs.createWriteStream('FILE2.txt');
r.pipe(yell).pipe(w);
