package main

import (
	"fmt"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter03/internal/nulls"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.NullEncoding(); err != nil {
		panic(err)
	}
}
