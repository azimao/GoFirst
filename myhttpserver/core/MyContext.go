package core

import (
	"github.com/pquerna/ffjson/ffjson"
	"net/http"
)

type MyContext struct {
	request *http.Request
	http.ResponseWriter
}

func (this *MyContext) WriteString(str string) {
	this.Write([]byte(str))
}

func (this *MyContext) WriteJson(m interface{}) {
	this.Header().Add("Content-type", "application/json")
	res, err := ffjson.Marshal(m)
	if err != nil {
		panic(err)
	}
	this.Write(res)
}
