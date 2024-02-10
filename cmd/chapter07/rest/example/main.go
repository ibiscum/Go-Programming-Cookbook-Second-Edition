package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter7/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
