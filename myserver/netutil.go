package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func GetRequestPath(rq string) string {
	r := regexp.MustCompile(`^GET\s(.*?)\sHTTP`)
	if r.MatchString(rq) {
		return r.FindStringSubmatch(rq)[1]
	} else {
		return "/"
	}
}

func response(body []byte, status HttpStatus) string {
	str := `HTTP/1.1 %d %s
Server: myserver
Content-Type: text/html

%s
`
	ret := fmt.Sprintf(str, status.Code, status.Message, body)
	return ret
}

func ExistsFile(path string) (bool, error) {
	_, err := os.Stat("./web" + path)
	if err == nil {
		return true, nil
	} else {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
}

func ReadHtml(path string) string {
	exists, _ := ExistsFile(path)
	if exists {
		file, _ := ioutil.ReadFile("./web" + path)
		return response(file, newHttpStatus(200, "ok"))
	} else {
		return response([]byte("Not Found"), newHttpStatus(404, "Not Found"))
	}
}
