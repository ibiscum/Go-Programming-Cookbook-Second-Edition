package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter07/internal/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
