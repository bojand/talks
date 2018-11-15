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

- Maybe the same mechanisms of how we served static content is perhaps not the best way for applications to communicate to each other?

---

class: center, middle

.title[RPC]

???

- We want the convenience of local function calls... but to be executed in distributed manner.
- That is if we commit into the RPC paradigm in the distributed systems context
- Alternative we could chose different architecture altogeher, such as asyncrhonous reactive systems.
- But even then, a lot of what we talk about here may still apply in some ways.

---

.center[![Issue](/img/grpc-logo.svg)]

<blockquote><strong>A high performance, open-source universal RPC framework</strong></blockquote>

???

---

# FEATURES

- HTTP/2
- RPC using Protocol Buffers (or JSON)
- Plugins to extend functionality
- Forwards / Backwards Compatible
- Self-Describing
- Streaming
- General appication framework (logging, security, monitoring, tracing)
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

# ARCHITECTURE

<br>

.center[![Architecture](/img/grpc-arch.svg)]

???

- Generated code provides client libraries and server stubs
- RPC Mechanisms
- Unary - simple request / response
- Streaming request and single response
- Single request and streaming response
- Duplex / bi-directional streaming

