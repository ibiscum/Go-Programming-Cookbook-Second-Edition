package main

import (
	"fmt"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter4/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
