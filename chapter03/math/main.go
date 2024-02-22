package main

import (
	"fmt"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter03/internal/math"
)

func main() {
	math.Examples()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()
}
