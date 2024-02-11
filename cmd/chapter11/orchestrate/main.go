package main

import (
	mongodb "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter11/orchestrate"
)

func main() {
	if err := mongodb.Exec("mongodb://mongodb:27017"); err != nil {
		panic(err)
	}
}
