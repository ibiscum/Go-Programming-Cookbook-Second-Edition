package main

import (
	"context"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter07/internal/oauthcli"
)

func main() {
	ctx := context.Background()
	conf := oauthcli.Setup()

	tok, err := oauthcli.GetToken(ctx, conf)
	if err != nil {
		panic(err)
	}
	client := conf.Client(ctx, tok)

	if err := oauthcli.GetUser(client); err != nil {
		panic(err)
	}

}
