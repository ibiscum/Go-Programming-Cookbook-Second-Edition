package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter01/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
