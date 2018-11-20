const path = require('path')
const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')

const PROTO_PATH = path.join(__dirname, '../../protos/greeter.proto')

const packageDefinition = protoLoader.loadSync(PROTO_PATH)
const proto = grpc.loadPackageDefinition(packageDefinition).greeter

const HOSTPORT = '0.0.0.0:50051'

function sayHello(call, callback) {
  console.log('server:sayHello')
  callback(null, { message: 'Hello ' + call.request.name })
}

function sayHellos(call) {
  console.log('server:sayHellos')
  const count = call.request.count
  let n = 0
  const timer = setInterval(() => {
    if (n < count) {
        call.write({ message: 'Hello ' + call.request.name })
        n++
    } else {
        clearInterval(timer)
        call.end()
    }
  }, 200)
}

function sayHelloCs(call, fn) {
  const rn = call.metadata.getMap().rn
  console.log(rn)

  let counter = 0
  call.on('data', d => {
    console.dir(d)
    counter++
  })

  call.on('end', () => {
    fn(null, { message: 'Hello ' + counter })
  })
}

function sayHelloBidi(call) {
  call.on('data', d => {
    call.write({ message: 'Hello ' + d.name })
  })

  call.on('end', () => {
    call.end()
  })
}

function main() {
  var server = new grpc.Server()

  server.addService(proto.Greeter.service, {
    sayHello,
    sayHellos,
    // sayHelloCs,
    // sayHelloBidi
  })

  server.bind(HOSTPORT, grpc.ServerCredentials.createInsecure())
  server.start()
}

main()
