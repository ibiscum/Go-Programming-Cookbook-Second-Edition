package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter05/internal/dns"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <address>\n", os.Args[0])
		os.Exit(1)
	}
	address := os.Args[1]
	lookup, err := dns.LookupAddress(address)
	if err != nil {
		log.Panicf("failed to lookup: %s", err.Error())
	}
	fmt.Println(lookup)
}
