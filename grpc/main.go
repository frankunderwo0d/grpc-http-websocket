package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	active "grpc/active"
	"grpc/controller"
	"net"
)

var (
	port = flag.Int64("p", 1111, "the port of grpc server")
)

func init() {
	flag.Parse()
}

func main() {
	//simpleVersion()
	customVersion()
}

func simpleVersion() {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	active.RegisterAServiceServer(grpcServer, &controller.AService{})
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func customVersion() {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	active.RegisterCommunicationServer(grpcServer,controller.Controller)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
