package main

import (
	"context"
	"fmt"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter10/internal/pipeline"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	in, out := pipeline.NewPipeline(ctx, 10, 2)

	go func() {
		for i := 0; i < 20; i++ {
			in <- fmt.Sprint("Message ", i)
		}
	}()

	for i := 0; i < 20; i++ {
		<-out
	}
}
