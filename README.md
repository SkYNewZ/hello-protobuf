# hello-protobuf

Hello world project to learn [protobuf](https://github.com/protocolbuffers/protobuf) and [gRPC](https://grpc.io).

## Requirements

You need these tools in your `$PATH`.

- Go environment
- [`protoc`](https://grpc.io/docs/protoc-installation/)
- [`protoc-gen-grpc-gateway`](https://github.com/grpc-ecosystem/grpc-gateway#installation)
- [`protoc-gen-go`](https://grpc.io/docs/languages/go/quickstart/)
- [`protoc-gen-go-grpc`](https://grpc.io/docs/languages/go/quickstart/)
- [`protoc-gen-openapiv2`](https://github.com/grpc-ecosystem/grpc-gateway#installation)

## Compile and run

```shell
$ protoc -I ./proto -I ../googleapis \
--go_out . \
--go-grpc_out . \
--grpc-gateway_out . \
--openapiv2_out . \
./proto/greeter.proto

$ go run .
```

This will start the gRPC server on `tcp://localhost:7777` and the REST server on `http://localhost:7778`.
You can use [bloomrpc](https://github.com/uw-labs/bloomrpc) for gRPC as Postman for HTTP.

## Resources

- https://blog.eleven-labs.com/fr/presentation-grpc/
- https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2
- https://venilnoronha.io/envoy-grpc-and-rate-limiting
- https://grpc.io
- https://developers.google.com/protocol-buffers/docs/proto3