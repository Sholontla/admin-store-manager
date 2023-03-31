package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

const redirectUri = "http://localhost:8081/callback"

type app struct {
}

func main() {
	a := app{}

	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			a.index(ctx)

		case "/callback":
			a.callback(ctx)

		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	log.Println("Running Server")
	err := fasthttp.ListenAndServe(":8081", m)

	if err != nil {
		fmt.Printf("ListenAndServe error: %s\n", err)
	}
}

func (a *app) index(ctx *fasthttp.RequestCtx) {
}

func (a *app) callback(ctx *fasthttp.RequestCtx) {

}
