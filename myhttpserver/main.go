package main

import (
	"com.azimao/myhttpserver/core"
	"net/http"
)

type MyHandler struct {
}

func main() {
	router := core.DefaultRouter()

	router.Add("/", &NewsController{})

	http.ListenAndServe(":8099", router)
}
