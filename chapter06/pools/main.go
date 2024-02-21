package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter06/internal/pools"

func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
