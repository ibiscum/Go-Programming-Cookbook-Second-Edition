package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter07/internal/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
