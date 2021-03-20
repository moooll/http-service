package handlers

import (
	"fmt"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"http-service/dto"
	"http-service/storage"
)

/*Добрый день!
Предлагаем вам выполнить тестовое задание, подробности ниже.
Если будут любые вопросы, пишите, с радостью ответим)
В случае успешного выполнения тестового задания мы пригласим вас на собеседование по видеосвязи или лично (если проживаете в Москве)

Задание:

Необходимо создать Http сервис - key-value хранилище.
Сервис должен содержать четыре метода в апи:
- Upsert (вставка либо обновление)
- Delete
- Get
- List
Хранить данные можно просто в оперативной памяти при помощи map.

Результаты можно залить на git или скинуть здесь архивом.*/

func Upsert(db *storage.Storage, ctx *fasthttp.RequestCtx) {
	request := []dto.Request{}
	err := ffjson.Unmarshal(ctx.Request.Body(), &request)
	if err != nil {
		zap.L().Info("error unmarshalling request: ", zap.Error(err))
		fmt.Println(err.Error())
		ctx.WriteString("could not unmarshal request")
	}
	// w/ params from uri
	// params := []dto.Request{}
	// keys := ctx.QueryArgs().PeekMulti("key")
	// values := ctx.QueryArgs().PeekMulti("value")
	// for i, v := range keys {
	// 	params = append(params, dto.Request{string(v), string(values[i])})
	// }
	db.Upsert(request)
	// todo: проверка перед отправкой статус кода тут должна быть??
	ctx.SetBodyString("updated")
	ctx.SetStatusCode(200)
}

// "/delete" receives a collection of "keys" as an argument
func Delete(db *storage.Storage, ctx *fasthttp.RequestCtx) {
	var keys []string
	err := ffjson.Unmarshal(ctx.Request.Body(), &keys)
	if err != nil {
		zap.L().Info("error unmarshalling request", zap.Error(err))
	}

	err = db.Delete(keys)
	if err != nil {
		ctx.SetBodyString("could not delete keys")
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
			zap.L().Info("error marshalling request", zap.Error(err))
		}
		ctx.SetBody(body)
	}
}

func List(db *storage.Storage, ctx *fasthttp.RequestCtx) {
	values, err := db.List()
	if err != nil {
		ctx.WriteString("nothing found")
	} else {
		body, err := ffjson.Marshal(values)
		if err != nil {
			zap.L().Info("error marshalling request", zap.Error(err))
			ctx.WriteString("error marshalling request")
		}
		ctx.Write(body)
	}
}
