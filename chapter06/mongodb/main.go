package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter06/internal/mongodb"

func main() {
	if err := mongodb.Exec("mongodb://localhost"); err != nil {
		panic(err)
	}
}
