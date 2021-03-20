package main

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"http-service/handlers"
	"http-service/storage"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}	

	defer logger.Sync()
	var db = &storage.Storage{}
	// err = db.FillStorage()
	if err != nil {
		zap.L().Error(err.Error())
	}
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/upsert":
			handlers.Upsert(db, ctx)
		case "/delete":
			handlers.Delete(db, ctx)
		case "/get":
			handlers.Get(db, ctx)
		case "/list":
			handlers.List(db, ctx)
		}
	}

	fasthttp.ListenAndServe(":8082", m)
}
