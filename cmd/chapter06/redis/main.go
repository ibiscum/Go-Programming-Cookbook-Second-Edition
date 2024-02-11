package main

import "github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter06/redis"

func main() {
	if err := redis.Exec(); err != nil {
		panic(err)
	}

	if err := redis.Sort(); err != nil {
		panic(err)
	}
}
