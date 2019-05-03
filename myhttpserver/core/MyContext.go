package core

import "net/http"

type MyContext struct {
	request *http.Request
	http.ResponseWriter
}

func (this *MyContext) WriteString(str string) {
	this.Write([]byte(str))
}
