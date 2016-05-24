var fs = require('fs');
console.log('reading file');
fs.readFile('file.txt', 'utf8', function (err, data) {
  if (err) throw err;
  console.log(data);
	fs.writeFile('file2.txt', data, function (err) {
	  if(err) throw err;
	  console.log('all done!');
	});
});
console.log('end of code');
