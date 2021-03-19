package handlers

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"http-service/dto"
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

// no globals??

func Upsert(ctx *fasthttp.RequestCtx) {
	request := []dto.Request{}
	err := ffjson.Unmarshal(ctx.Request.Body(), &request)
	if err != nil {
		zap.L().Info("error unmarshalling request: " + err.Error())
		ctx.WriteString("could not unmarshal request")
	}
	db.Upsert(request)
	// todo: проверка перед отправкой статус кода тут должна быть??
	ctx.SetStatusCode(200)
}

// "/delete" receives a collection of "keys" as an argument
func Delete(ctx *fasthttp.RequestCtx) {
	var keys []string
	err := ffjson.Unmarshal(ctx.Request.Body(), &keys)
	if err != nil {
		zap.L().Info("error unmarshalling request" + err.Error())
	}

	err = db.Delete(keys)
	if err != nil {
		ctx.SetBodyString("could not delete keys")
	} else {
		ctx.SetStatusCode(200)
	}
}

func Get(ctx *fasthttp.RequestCtx) {
	var keys []string
	args := ctx.QueryArgs().PeekMulti("key")
	for _, v := range args {
		keys = append(keys, string(v))
	}
	// err := ffjson.Unmarshal(ctx.Request.Body(), &keys)
	// if err != nil {
	// 	zap.L().Info("error unmarshalling request")
	// }

	values, err := db.Get(keys)
	if err != nil {
		ctx.SetBodyString(err.Error())
	} else {
		body, err := ffjson.Marshal(values)
		if err != nil {
			zap.L().Info("error marshalling request" + err.Error())
		}
		ctx.SetBody(body)
	}
	// append result to body
}

func List(ctx *fasthttp.RequestCtx) {
	values, err := db.List()
	if err != nil {
		ctx.WriteString("nothing found")
	} else {
		body, err := ffjson.Marshal(values)
		if err != nil {
			zap.L().Info("error marshalling request" + err.Error())
		}
		ctx.SetBody(body)
	}
}
