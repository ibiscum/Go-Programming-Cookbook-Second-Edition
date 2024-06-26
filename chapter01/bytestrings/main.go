package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter01/internal/bytestrings"

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	// each of these print to stdout
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
