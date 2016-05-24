var level = require('level');

// 1) Create our database, supply location and options.
//    This will create or open the underlying LevelDB store.
var db = level('./mydb', { valueEncoding: 'json' });

// 2) put a key & value
db.put('key-123', { foo: 'bar' }, function (err) {
  if (err) return console.log('Ooops!', err); // some kind of I/O error
  // 3) fetch by key
  db.get('key-123', function (err, value) {
    if (err) return console.log('Ooops!', err); // likely the key was not found
    console.log(value);
  });
});
