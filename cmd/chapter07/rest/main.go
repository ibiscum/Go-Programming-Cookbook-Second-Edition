package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter07/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
