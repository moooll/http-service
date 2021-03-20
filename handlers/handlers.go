package handlers

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"http-service/dto"
	"http-service/storage"
)

func Upsert(db *storage.Storage, ctx *fasthttp.RequestCtx) {
	request := []dto.Request{}
	err := ffjson.Unmarshal(ctx.Request.Body(), &request)
	if err != nil {
		zap.L().Info("error unmarshalling request: ", zap.Error(err))
		ctx.WriteString("could not unmarshal request")
		ctx.SetStatusCode(400)
	} else {
		db.Upsert(request)
		ctx.SetBodyString("updated")
		ctx.SetStatusCode(200)
	}
}

func Delete(db *storage.Storage, ctx *fasthttp.RequestCtx) {
	var keys []string
	args := ctx.QueryArgs().PeekMulti("key")
	for _, v := range args {
		keys = append(keys, string(v))
	}
	err := db.Delete(keys)
	if err != nil {
		ctx.SetBodyString(err.Error())
	} else {
		zap.L().Info("deleted")
		ctx.SetBodyString("deleted")
		ctx.SetStatusCode(200)
	}
}

func Get(db *storage.Storage, ctx *fasthttp.RequestCtx) {
	var keys []string
	args := ctx.QueryArgs().PeekMulti("key")
	for _, v := range args {
		keys = append(keys, string(v))
	}
	values, err := db.Get(keys)
	if err != nil {
		ctx.SetBodyString(err.Error())
	} else {
		body, err := ffjson.Marshal(values)
		if err != nil {
			zap.L().Info("error marshalling response", zap.Error(err))
			ctx.SetBodyString("server error")
		} else {
			ctx.SetBody(body)
		}
	}
}

func List(db *storage.Storage, ctx *fasthttp.RequestCtx) {
	values, err := db.List()
	if err != nil {
		ctx.WriteString(err.Error())
	} else {
		body, err := ffjson.Marshal(values)
		if err != nil {
			zap.L().Info("error marshalling response", zap.Error(err))
			ctx.WriteString("server error")
		}
		ctx.SetBody(body)
	}
}
