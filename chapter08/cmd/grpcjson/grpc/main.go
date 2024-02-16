package main

import (
	"fmt"
	"net"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter08/internal/grpcjson/keyvalue"
	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter08/internal/grpcjson/kvintern"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	keyvalue.RegisterKeyValueServer(grpcServer, kvintern.NewKeyValue())
	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port :4444")
	grpcServer.Serve(lis)
}
