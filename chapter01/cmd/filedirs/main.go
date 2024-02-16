package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter01/internal/filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}
