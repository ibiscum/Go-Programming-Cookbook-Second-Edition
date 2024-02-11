package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter07/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
