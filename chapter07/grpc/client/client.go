package main

import (
	"context"
	"fmt"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter07/internal/grpc/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":4444", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := greeter.NewGreeterServiceClient(conn)

	ctx := context.Background()
	req := greeter.GreetRequest{Greeting: "Hello", Name: "Reader"}
	resp, err := client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	req.Greeting = "Goodbye"
	resp, err = client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
