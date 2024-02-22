package main

import (
	mongodb "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter11/internal/orchestrate"
)

func main() {
	if err := mongodb.Exec("mongodb://mongodb:27017"); err != nil {
		panic(err)
	}
}
