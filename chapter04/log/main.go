package main

import (
	"fmt"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter04/internal/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
