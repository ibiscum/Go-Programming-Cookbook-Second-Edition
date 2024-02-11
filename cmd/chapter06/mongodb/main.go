package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter06/mongodb"

func main() {
	if err := mongodb.Exec("mongodb://localhost"); err != nil {
		panic(err)
	}
}
