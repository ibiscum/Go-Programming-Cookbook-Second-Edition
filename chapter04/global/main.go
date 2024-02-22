package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter04/internal/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
