package main

import (
	"fmt"
	"log"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter12/internal/goflow"
	flow "github.com/trustmaster/goflow"
)

func main() {

	net := goflow.NewEncodingApp()

	in := make(chan string)
	err := net.SetInPort("In", in)
	if err != nil {
		log.Fatal(err)
	}

	wait := flow.Run(net)

	for i := 0; i < 20; i++ {
		in <- fmt.Sprint("Message", i)
	}

	close(in)
	<-wait
}
