var oboe = require('oboe');
var url = 'http://localhost:3000/users';
oboe(url).on('node', '{email firstName lastName}', function (user) {
  console.log(`USER: ${user.email} ${user.firstName} ${user.lastName}`)
});
