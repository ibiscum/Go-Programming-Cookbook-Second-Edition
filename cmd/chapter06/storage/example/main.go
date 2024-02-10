package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter06/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
