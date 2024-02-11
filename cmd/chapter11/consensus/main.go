package main

import (
	"net/http"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter11/consensus"
)

func main() {
	consensus.Config(3)

	http.HandleFunc("/", consensus.Handler)
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
