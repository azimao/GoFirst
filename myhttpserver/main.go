package main

import (
	"com.azimao/myhttpserver/core"
	"net/http"
)

type MyHandler struct {
}

func main() {
	router := core.DefaultRouter()

	router.Get("/", func(ctx *core.MyContext) {
		ctx.WriteString("my string GET")
	})

	router.Post("/", func(ctx *core.MyContext) {
		ctx.WriteString("my string POST")
	})
	http.ListenAndServe(":8099", router)
}
