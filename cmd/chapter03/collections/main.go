package main

import (
	"fmt"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter03/collections"
)

func main() {
	ws := []collections.WorkWith{
		{"Example", 1},
		{"Example 2", 2},
	}

	fmt.Printf("Initial list: %#v\n", ws)

	// first lower case the list
	ws = collections.Map(ws, collections.LowerCaseData)
	fmt.Printf("After LowerCaseData Map: %#v\n", ws)

	// next increment all versions
	ws = collections.Map(ws, collections.IncrementVersion)
	fmt.Printf("After IncrementVersion Map: %#v\n", ws)

	// lastly remove all versions older than 3
	ws = collections.Filter(ws, collections.OldVersion(3))
	fmt.Printf("After OldVersion Filter: %#v\n", ws)
}
