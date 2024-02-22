package main

import (
	"fmt"
	"net/http"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter07/internal/twirp/greeter"
)

func main() {
	server := &Greeter{}
	twirpHandler := greeter.NewGreeterServiceServer(server, nil)

	fmt.Println("Listening on port :4444")
	err := http.ListenAndServe(":4444", twirpHandler)
	if err != nil {
		panic(err)
	}
}
