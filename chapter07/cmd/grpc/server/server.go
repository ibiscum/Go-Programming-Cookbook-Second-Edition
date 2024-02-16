package main

import (
	"fmt"
	"net"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter07/internal/grpc/greeter"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServiceServer(grpcServer, &server{})
	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port :4444")
	grpcServer.Serve(lis)
}
