class: center, middle

# Beyond REST

## A guide to gRPC

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

# REST

* Well-supported
* Debuggable
* Cachable
* Scalable
* Standardized?
* Easy?
* Performant?

???
Talk about current state

---

# Example API

.left[<img src="/img/twilio-logo-red.svg" alt="Twilio Logo" height="42px">]

- `200 OK` for `GET`
- `201 OK` for `POST` and `PUT`
- `204 OK` for `DELETE`

.left[<img src="/img/Stripe logo - blue.svg" alt="Stripe Logo" height="56px">]

- `200 OK` for all successful requests

???
Notes

---

# REST API Considerations

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
Notes

---

# REST API Considerations

.left[![Issue](/img/language_issue.png)]
<!-- <img src="/img/language_issue.png" alt="Issue" width="640px"> -->

???
Notes

---

# REST API Considerations

- HTTP/1 is not performant
- Text-based protocol is developer-friendly but inefficient
- HTTP/2 better, but not as widely used
- Semantics
  * `POST`/`PUT`/`PATCH`
  * Status codes
  * Error responses
  * Single vs. plural resource names
  * Versioning

???
Notes

---

# gRPC

.center[![Issue](/img/grpc-logo.svg)]

- HTTP/2
- IDL with multi-language support
- Efficient binary srialization using Protocol Buffers
- Plugins to extend functionality
- Forwards / Backwards Compatible
- Self-Describing

???
Notes

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

# Protocol Buffers

- Interface Definition Language
- Efficient binary serialization format
- Language-neutral
- Flexible & Extensible
- `protoc` compiler with plugin support
- Language support: C++, C#, Dart, Go, Java, Android, Java, Node.js, Objective-C, PHP, Python, Ruby
- Compile `.proto` file to generate language-specific code

???
Notes

---

# Protocol Buffers gRPC Example

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

# Generating code

```sh
$ protoc -I helloworld/ \ 
  helloworld/helloworld.proto \
  --go_out=plugins=grpc:helloworld
```

```sh
$ npm install -g grpc-tools
$ grpc_tools_node_protoc \
  --js_out=import_style=commonjs,binary:../node/static_codegen/ \
  --grpc_out=../node/static_codegen \
  --plugin=protoc-gen-grpc=grpc_node_plugin helloworld.proto
```
