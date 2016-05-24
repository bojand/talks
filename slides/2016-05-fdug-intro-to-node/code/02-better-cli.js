var program = require('commander');

program
  .version('1.0.0')
  .option('-g, --greeting [greeting]', 'Specify greeting')
  .option('-n, --name [name]', 'Specify name')
  .parse(process.argv);

console.log(`${program.greeting} ${program.name}!`);
