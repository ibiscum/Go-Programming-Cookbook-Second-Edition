package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// GetItems will return our items object
func GetItems(ctx *fasthttp.RequestCtx) {
	enc := json.NewEncoder(ctx)
	items := ReadItems()
	err := enc.Encode(&items)
	if err != nil {
		panic(err)
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// AddItems modifies our array
func AddItems(ctx *fasthttp.RequestCtx) {
	item, ok := ctx.UserValue("item").(string)
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	AddItem(item)
	ctx.SetStatusCode(fasthttp.StatusOK)
}
