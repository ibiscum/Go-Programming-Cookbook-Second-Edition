package main

import (
	"net/http"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter11/internal/consensus"
)

func main() {
	consensus.Config(3)

	http.HandleFunc("/", consensus.Handler)
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
