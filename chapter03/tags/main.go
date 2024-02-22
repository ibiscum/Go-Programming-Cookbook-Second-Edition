package main

import (
	"fmt"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter03/internal/tags"
)

func main() {

	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
