package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter03/internal/dataconv"

func main() {
	dataconv.ShowConv()
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Interfaces()
}
