package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/SkYNewZ/hello_protobuf/internal/greeter"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type greaterHandler struct {
	greeter.UnimplementedGreeterServer
}

func (g greaterHandler) SayHello(_ context.Context, request *greeter.HelloRequest) (*greeter.HelloReply, error) {
	log.Printf("greaterHandler.SayHello")
	return &greeter.HelloReply{Message: request.Name}, nil
}

func NewGreaterServiceServer() greeter.GreeterServer {
	return &greaterHandler{}
}

type server struct{}

func (s server) Serve() {
	gRPCAddr := "localhost:7777"
	httpAddr := "localhost:7778"

	// GRPC server
	go func() {
		listener, err := net.Listen("tcp", gRPCAddr)
		if err != nil {
			log.Fatalln(err)
		}

		srv := grpc.NewServer()
		serviceServer := NewGreaterServiceServer()
		greeter.RegisterGreeterServer(srv, serviceServer)

		log.Printf("starting gRPC server at %s", gRPCAddr)
		if err := srv.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}()

	// REST server
	go func() {
		mux := runtime.NewServeMux()

		err := greeter.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, gRPCAddr, []grpc.DialOption{
			grpc.WithInsecure(),
		})
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("starting REST server at %s", httpAddr)
		if err := http.ListenAndServe(httpAddr, mux); err != nil {
			log.Fatalln(err)
		}
	}()

	// infinite loop
	log.Printf("entering infinite loop")
	select {}
}

type Server interface {
	Serve()
}

func NewServer() Server {
	return &server{}
}
