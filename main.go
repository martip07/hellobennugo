package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type GeneralMessage struct {
	Message string `json:message`
}

func HelloHandler(ctx *fasthttp.RequestCtx) {
	respGeneral := GeneralMessage{
		Message: "Hello from bennu demo",
	}
	jsonBody, err := json.Marshal(respGeneral)
	if err != nil {
		ctx.Error("Ups, json didn't work", 500)
		return
	}
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetStatusCode(200)
	ctx.Response.SetBody(jsonBody)
	return
}

func main() {
	fmt.Println("Howdy from hellobennugo")
	r := router.New()
	r.GET("/api/hello", HelloHandler)
	log.Fatal(fasthttp.ListenAndServe(":8088", r.Handler))
}
