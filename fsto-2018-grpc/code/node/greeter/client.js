const PROTO_PATH = __dirname + '../../../protos/greeter.proto'

const async = require('async')
const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const packageDefinition = protoLoader.loadSync(PROTO_PATH)
const proto = grpc.loadPackageDefinition(packageDefinition).greeter

const client = new proto.Greeter('localhost:50051', grpc.credentials.createInsecure())

function sayHello(fn) {
  console.log('client:sayHello')
  client.sayHello({ name: 'world' }, (err, response) => {
    console.log('Greeting: ', response.message)
    fn()
  })
}

function sayHellos(fn) {
  console.log('client:sayHellos')
  const call = client.sayHellos({ name: 'world', count: 5 })
  call.on('data', ({ message }) => console.log('Greeting: ', message))
  call.on('end', fn)
}

function main() {
  async.series([sayHello, sayHellos])
}

main()
