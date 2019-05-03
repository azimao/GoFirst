package core

import "net/http"

type MyHandlerFunc func(ctx *MyContext)
type MyRouter struct {
	Mapping map[string]map[string]MyHandlerFunc
	Ctx     *MyContext
}

func DefaultRouter() *MyRouter {
	return &MyRouter{make(map[string]map[string]MyHandlerFunc), &MyContext{}}
}

func (this *MyRouter) Get(path string, f MyHandlerFunc) {
	if this.Mapping["GET"] == nil {
		this.Mapping["GET"] = make(map[string]MyHandlerFunc)
	}
	this.Mapping["GET"][path] = f
}

func (this *MyRouter) Post(path string, f MyHandlerFunc) {
	if this.Mapping["Post"] == nil {
		this.Mapping["Post"] = make(map[string]MyHandlerFunc)
	}
	this.Mapping["Post"][path] = f
}

func (this *MyRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	this.Ctx.request = request
	this.Ctx.ResponseWriter = writer
	if f, OK := this.Mapping[request.Method][request.URL.Path]; OK {
		f(this.Ctx)
	}
}
