package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter14/internal/crypto"
)

func main() {

	http.HandleFunc("/guess", crypto.GuessHandler)
	fmt.Println("server started at localhost:8080")
	log.Panic(http.ListenAndServe("localhost:8080", nil))
}
