package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter01/internal/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
