package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter01/templates"

func main() {
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}

	if err := templates.InitTemplates(); err != nil {
		panic(err)
	}

	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}
}
