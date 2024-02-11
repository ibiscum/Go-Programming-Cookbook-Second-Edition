package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter4/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
