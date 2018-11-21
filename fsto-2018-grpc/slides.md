class: center, middle

<img src="/img/FSTO_400x400.png" alt="FSTO Logo" width="156px">

.title[Beyond REST]

## A guide to gRPC

---

# A SIMPLER TIME...

<img src="/img/the_internet.jpg" alt="Welcome to the internet" width="100%">

.footnote[Source: Gizmodo]

???

The internet was simple. Servers served static HTML content that was rendered in the browser.

---

# A NEW AGE

<img src="/img/programming.jpg" alt="programming" width="100%">

.footnote[Source: Pexels]

???

- Then we started programming the web and making the content more dynamic.
- Not only did our applications become more complex, but how we consumed the data changed.
- Browsers are not the only clients of today. We have different mobile devices and platforms consuming our data.

---

# MICROSERVICES!

.center[<img src="/img/microservices.png" alt="microservices">]

.footnote[Source: ArcGIC]

???

- As our systems grew and evolved, to overcome the challenges of scale we started breaking up the monolith into microservices.
- How systems and different application communicate and talk to each other has drastically changed over the past two decades.

---

# PRESENT DAY

- Complex distributed computing landscape
- Systems and applications need to communicate
- Almost everything provides an API

<br>

.center[

  **How do apps communicate?**
  
  **How do we build APIs?**
]

???

---

# REST

<blockquote><strong>HTTP + JSON is the de facto standard for REST communication</strong></blockquote>

```sh
$ curl https://api.stripe.com/v1/charges \
   -u sk_test_4eC39HqLyjWDarjtT1zdp7dc: \
   -d amount=2000 \
   -d currency=usd \
   -d source=tok_visa \
   -d description="Charge for jenny.rosen@example.com"
```

???
Talk about current state

---

# HTTP/REST IS GREAT

- Text-based and debuggable
- Tooling for testing & inspection
- Well-supported in most languages 
- Cachable
- Scalable
- Standardized?
- Easy?
- Performant?

???
Talk about current state

---

# EXAMPLES

.left[<img src="/img/twilio-logo-red.svg" alt="Twilio Logo" height="42px">]

- `200 OK` for `GET`
- `201 OK` for `POST` and `PUT`
- `204 OK` for `DELETE`

.left[<img src="/img/Stripe logo - blue.svg" alt="Stripe Logo" height="56px">]

- `200 OK` for all successful requests

???
Notes

---

# REST API CONSIDERATIONS

- Schema
- Authentication
- Documentation
- Versioning
- Root endpoint
- Status code & client errors
- Redirects
- HTTP verbs
- Hypermedia
- Pagination
- Conditional Requests
- CORS
- JSON-P

???

And even if you take the time and get it all right...

---

# CLIENT LIBRARIES

.left[![Issue](/img/language_issue.png)]
<!-- <img src="/img/language_issue.png" alt="Issue" width="640px"> -->

???
Notes

---

# REST API CONSIDERATIONS

- HTTP/1 is not performant
- Text-based protocol is developer-friendly but inefficient
- Streaming is difficult
- No bi-directional streaming
- Not all operations are "resource"-based
- Semantics
  * `POST`/`PUT`/`PATCH`
  * Status codes
  * Error responses
  * Single vs. plural resource names
  * Versioning

???

- Example of non-resourceful: 
    * Restart some application / host
    * Encrypt some text
    * Classify an image or a sentance
- Maybe the same mechanisms of how we served static content is perhaps not the best way for applications to communicate to each other?

---

class: center, middle

.title[RPC]

???

- We want the convenience of local function calls... but to be executed in distributed manner.
- That is if we commit into the RPC paradigm in the distributed systems context
- Alternative we could chose different architecture altogeher, such as asynchronous reactive systems.
- But even then, a lot of what we talk about here may still apply in some ways.

---

.center[![Issue](/img/grpc-logo.svg)]

<blockquote><strong>A high performance, open-source universal RPC framework</strong></blockquote>

???

- Originally a Google project internally called "Stubby"
- Open sourced, mainly developerd by Google employees

---

# FEATURES

- HTTP/2
- RPC using Protocol Buffers (or JSON)
- Plugins to extend functionality
- Forwards / Backwards Compatible on the wire
- Self-Describing
- Streaming call support
- Mobile: Android and Objective-C, Experimental Swift
- Polyglot: C++, Go, Java, Ruby, Node.js, Python, C#, PHP

???

- HTTP2 is binary, instead of textual
- is fully multiplexed, instead of ordered and blocking
- multiple requests can be serviced at the same time in one connection
- one long lived connection
- each message is an HTTP/2 request with its own headers
- uses header compression to reduce overhead
- allows servers to “push” responses proactively to clients
- General appication framework allows for logging, security, monitoring, tracing via middleware and interceptors

- gRPC core implementations in C++, Go and Java. All others based on C++ core.

---

# gRPC ?

- 1.0 'g' stands for ['gRPC'](https://github.com/grpc/grpc/tree/v1.0.x)
- 1.1 'g' stands for ['good'](https://github.com/grpc/grpc/tree/v1.1.x)
- 1.2 'g' stands for ['green'](https://github.com/grpc/grpc/tree/v1.2.x)
- 1.3 'g' stands for ['gentle'](https://github.com/grpc/grpc/tree/v1.3.x)
- 1.4 'g' stands for ['gregarious'](https://github.com/grpc/grpc/tree/v1.4.x)
- 1.6 'g' stands for ['garcia'](https://github.com/grpc/grpc/tree/v1.6.x)
- 1.7 'g' stands for ['gambit'](https://github.com/grpc/grpc/tree/v1.7.x)
- 1.8 'g' stands for ['generous'](https://github.com/grpc/grpc/tree/v1.8.x)
- 1.9 'g' stands for ['glossy'](https://github.com/grpc/grpc/tree/v1.9.x)
- 1.10 'g' stands for ['glamorous'](https://github.com/grpc/grpc/tree/v1.10.x)
- 1.11 'g' stands for ['gorgeous'](https://github.com/grpc/grpc/tree/v1.11.x)
- ... ["g stands for" version list](https://github.com/grpc/grpc/blob/master/doc/g_stands_for.md)

???
Notes

---

# PROTOCOL BUFFERS

- Interface Definition Language
- Machine readable
- Efficient binary serialization format
- Language-neutral
- Flexible & extensible
- `protoc` compiler with plugin support
- Compile `.proto` file to generate language-specific code

???
Notes

---

# SERVICE DEFINITION

```proto
// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

???
Notes

---

# CODE GENERATION

```sh
$ protoc -I helloworld/ \ 
  helloworld/helloworld.proto \
  --go_out=plugins=grpc:helloworld
```

```sh
$ npm install -g grpc-tools
$ grpc_tools_node_protoc \
  --js_out=import_style=commonjs,binary:../codegen/ \
  --grpc_out=../codegen \
  --plugin=protoc-gen-grpc=grpc_node_plugin \
  helloworld.proto
```

???

- Install `protoc` compiler
- Compile `.proto` file to generate language-specific code
- Generated code is not to be edited
- Generated code is not idomatic for the target language, and is optimized for performance
- Use generated code for serialization and deserialization of data 

---

# MECHANISM

<br>

.center[![Architecture](/img/grpc-arch.svg)]

???

- Generated code provides client libraries and server stubs
- RPC Mechanisms
- Unary - simple client request & server response
- Streaming request and single server response
- Single client request and streaming response 
- Duplex / bi-directional streaming
- Streaming allows for no / easier pagination mechanisms without need for a cursor or page number

---

# SERVER - GO

```go
type server struct{}

func (s *server) SayHello(ctx context.Context, 
    in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()
  pb.RegisterGreeterServer(grpcServer, &server{})
  reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
```

???

- SayHello is the implementation of our service
- Think of it as lamda function
- The code in main() is a bit of boiler plate
- Reflection is for introspection. 
  - The service can explain what services and methods this grpc server has
  - Client can connect and build the client without knowning what lives on the server

---

# CLIENT - GO

```go
func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(),
		10*time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
```

???

- Notes

---

# SERVER - NODE.JS

```js
const PROTO_PATH = __dirname + './protos/helloworld.proto'
const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const packageDefinition = protoLoader.loadSync(PROTO_PATH)
const proto = 
  grpc.loadPackageDefinition(packageDefinition).helloworld

function sayHello(call, callback) {
  callback(null, { message: 'Hello ' + call.request.name })
}

function main() {
  var server = new grpc.Server()
  server.addService(proto.Greeter.service, 
    { sayHello: sayHello })
  server.bind('0.0.0.0:50051',
    grpc.ServerCredentials.createInsecure())
  server.start()
}

main()
```

???

- Notes

---

# CLIENT - NODE.JS

```js
const PROTO_PATH = __dirname + './protos/helloworld.proto';
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const proto = 
  grpc.loadPackageDefinition(packageDefinition).helloworld;

function main() {
  const client = new proto.Greeter(
    'localhost:50051', grpc.credentials.createInsecure());
  
  client.sayHello({ name: 'world' }, (err, response) => {
    console.log('Greeting: ', response.message);
  });
}

main();
```

???

- Notes

---

# MORE COMPLEX

```proto
syntax = "proto3";

package greeter;

service Greeter {
    rpc SayHello (HelloReq) returns (HelloRes) {}
    rpc SayHellos (HelloReq) returns (stream HelloRes) {}
    rpc GreetMany (stream HelloReq) returns (HelloRes) {}
    rpc GreetChat (stream HelloReq) returns (stream HelloRes) {}
}

message HelloReq {
    string name = 1;
    int32 count = 2;
}

message HelloRes {
    string message = 1;
}
```

???

- Notes

---

# SERVER STREAMING - SERVER

```js
function sayHellos(call) {
  let n = 0
  const timer = setInterval(() => {
    if (n < call.request.count) {
        call.write({ message: 'Hello ' + call.request.name })
        n++
    } else {
        clearInterval(timer)
        call.end()
    }
  }, 200)
}
```

???

- Notes

---

# SERVER STREAMING - CLIENT

```js
const deadline = 
  new Date().setSeconds(new Date().getSeconds() + 5)

const call = client.sayHellos(
  { name: 'world', count: 5 }, 
  { deadline })

call.on('data', 
  ({ message }) => console.log('Greeting: ', message))

call.on('end', () => console.log('done'))
```

???

- Notes

---

# CLIENT STREAMING - SERVER

```go
func (s *server) GreetMany(stream pb.Greeter_GreetManyServer)
error {
	names := make([]string, 0, 5)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			msg := "Hello " + strings.Join(names, ", ")
			res := &pb.HelloRes{Message: msg}
			return stream.SendAndClose(res)
		}
		if err != nil {
			return err
		}
		names = append(names, in.Name)
	}
}
```

???

- Notes

---

# CLIENT STREAMING - CLIENT

```go
func greetMany(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 
		10*time.Second)
	defer cancel()
  
	stream, err := client.GreetMany(ctx)
	
	names := [4]string{"Bob", "Kate", "Jim", "Sara"}
  
	for _, name := range names {
		msg := &pb.HelloReq{Name: names[n]}
		stream.Send(msg)
	}

	reply, _ := stream.CloseAndRecv()
	log.Printf("Greeting: %v", reply.Message)
}
```

???

- Notes

---

# BIDI STREAMING - SERVER

```go
func (s *server) GreetChat(stream pb.Greeter_GreetChatServer)
error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		msg := &pb.HelloRes{Message: "Hello " + in.Name}
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}
```

???

- Notes

---

# BIDI STREAMING - CLIENT

```js
call = client.greetChat()
const NAMES = ['Bob', 'Kate', 'Jim', 'Sara']
let n = 0
const timer = setInterval(() => {
  if (n < NAMES.length) {
    call.write({ name: NAMES[n] })
    n++
  } else {
    clearInterval(timer)
    call.end()
  }
}, 200)

call.on('data',
  ({ message }) => console.log('Greeting:', message))

call.on('end', () => console.log('done'))
```

???

- Notes

---

# METADATA - CLIENT

```go
conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
defer conn.Close()

c := pb.NewGreeterClient(conn)

ctx, cancel := context.WithTimeout(context.Background(),
  time.Second)
defer cancel()

ctx = metadata.AppendToOutgoingContext(
  ctx, "token", "xyz", "request-id", "123")

res, _ := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})

log.Printf("Greeting: %s", res.Message)
```

???

- All implementations should support it
- With Node.js just add it as an additional parameter

---

# METADATA - SERVER

```go
func (s *server) SayHello(ctx context.Context, 
  in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	for k, v := range md {
		fmt.Printf("%s: %s\n", k, v)
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
```

```sh
:authority: [localhost:50051]
content-type: [application/grpc]
user-agent: [grpc-go/1.16.0]
token: [xyz]
request-id: [123]
```

???

- Notes

---

# INTERCEPTORS

```go
func clientInterceptor(ctx context.Context, method string,
	req interface{}, reply interface{},
  cc *grpc.ClientConn, invoker grpc.UnaryInvoker, 
  opts ...grpc.CallOption,
) error {
	start := time.Now()

	err := invoker(ctx, method, req, reply, cc, opts...)

  fmt.Printf("Invoked RPC method=%s; Duration=%s; Error=%v\n", 
    method, time.Since(start), err)

	return err
}
```

```go
conn, err := grpc.Dial(addr, grpc.WithInsecure()
  grpc.WithUnaryInterceptor(clientInterceptor))
```

???

- Server-side interceptors / middleware also supported
- Not all languages have the same level of support
  * For example there is no support for server-side middleware for Node.js

---

# TOOLING - CLI

```sh
$ grpc_cli ls localhost:50051
helloworld.Greeter
grpc.reflection.v1alpha.ServerReflection

$ grpc_cli ls localhost:50051 helloworld.Greeter -l
filename: helloworld.proto
package: helloworld;
service Greeter {
  rpc SayHello(helloworld.HelloRequest) returns (helloworld.HelloReply) {}
}

$ grpc_cli call localhost:50051 SayHello 'name: "john"'
connecting to localhost:50051
message: "Hello john"

Rpc succeeded with OK status
```

???

- `grpc_cli` is the official command line tool
- There are other options such as grpcurl

---

# WEB SUPPORT ?

.center[![gRPC-Web](/img/grpc-web.png)]

```sh
protoc helloworld.proto \
  --js_out=import_style=commonjs:./codegen \
  --grpc-web_out=import_style=commonjs:./codegen
```

???

- We generate types like normal using `protoc`
- We Web client using `protoc` using 
- Envoy must be used as a proxy for web clients to talk to
- Nginx can also work

---

# HTTP / JSON + gRPC

```proto
package helloworld;

*import "google/api/annotations.proto";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {
*    option (google.api.http) = {
*      get: "/say"
*    };
  }
}
```

- Use [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to build a REST API service
- Use Envoy's gRPC-JSON transcoder filter

???

- `grpc-gateway` can be used to generate a Go stub that you then can create a go service proxy
- `grpc-gateway` can be used to generate swagger deinition as well
- Maps streaming APIs to newline-delimited JSON streams
- No BiDi streaming support

- Envoy's gRPC-JSON transcoder filter allows a RESTful JSON API client to send requests to Envoy over HTTP and get proxied to a gRPC service. The HTTP mapping for the gRPC service has to be defined by custom options.
- for gRPC stream request parameters, Envoy expects an array of messages, and it returns an array of messages for stream response parameters.
- No BiDi streaming support

- There are projects in Go at least that let you serve both HTTP+JSON and gRPC from single service on different ports
