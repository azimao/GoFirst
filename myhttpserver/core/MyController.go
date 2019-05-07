package core

import (
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
)

type MyController struct {
	Ctx *MyContext
}

func (this *MyController) Init(ctx *MyContext) {
	this.Ctx = ctx
}

func (this *MyController) GetParam(key string, defValue ...string) string {
	res := this.Ctx.request.URL.Query().Get(key)
	if res == "" && len(defValue) > 0 {
		return defValue[0]
	}
	return res
}

func (this *MyController) PostParam(key string, defValue ...string) string {
	res := this.Ctx.request.PostFormValue(key)
	if res == "" && len(defValue) > 0 {
		return defValue[0]
	}
	return res
}

func (this *MyController) JsonParam(obj interface{}) {
	body, err := ioutil.ReadAll(this.Ctx.request.Body)
	if err != nil {
		panic(err)
	}
	err = ffjson.Unmarshal(body, obj)
	if err != nil {
		panic(err)
	}
}

type ControllerInterface interface {
	Init(ctx *MyContext)
	GET()
	POST()
}
