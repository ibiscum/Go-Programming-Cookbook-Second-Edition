package main

import (
	"fmt"
	"net/http"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter08/internal/grpcjson/kvintern"
)

func main() {
	c := Controller{KeyValue: kvintern.NewKeyValue()}
	http.HandleFunc("/set", c.SetHandler)
	http.HandleFunc("/get", c.GetHandler)

	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
