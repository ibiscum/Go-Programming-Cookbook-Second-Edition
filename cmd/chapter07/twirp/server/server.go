package main

import (
	"fmt"
	"net/http"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter07/twirp/greeter"
)

func main() {
	server := &Greeter{}
	twirpHandler := greeter.NewGreeterServiceServer(server, nil)

	fmt.Println("Listening on port :4444")
	http.ListenAndServe(":4444", twirpHandler)
}
