package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter07/twirp/greeter"
)

func main() {
	// you can put in a custom client for tighter controls on timeouts etc.
	client := greeter.NewGreeterServiceProtobufClient("http://localhost:4444", &http.Client{})

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
