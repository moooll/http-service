package main

import (
	"math/rand"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"http-service/handlers"
	"http-service/storage"
)


func main() {
	var db = storage.Storage{}
	db = rand.New()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/upsert":
			handlers.Upsert(ctx)
		case "/delete":
			handlers.Delete(ctx)
		case "/get":
			handlers.Get(ctx)
		case "/list":
			handlers.List(ctx)
		}
	}

	fasthttp.ListenAndServe(":8082", m)
}
