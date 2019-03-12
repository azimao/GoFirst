package main

import (
	"fmt"
	"net"
	"regexp"
	"runtime"
	"time"
)

func response() string {
	str := `HTTP/1.1 200 OK
Server: myserver
Content-Type: text/html

this is body
`
	return str
}

func GetRequestPath(rq string) string {
	r := regexp.MustCompile(`^GET\s(.*?)\sHTTP`)
	if r.MatchString(rq) {
		return r.FindStringSubmatch(rq)[1]
	} else {
		return "/"
	}
}

func main() {
	list, err := net.Listen("tcp", "127.0.0.1:8099")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer list.Close()
	fmt.Println("监听成功，等待客户端连接")
	go func() {
		for {
			fmt.Printf("当前任务数:%d\n", runtime.NumGoroutine())
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		client, err := list.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go func(c net.Conn) {

			defer c.Close()
			buf := make([]byte, 4096)
			n, err := c.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("读取到%d, 内容是：%s", n, string(buf[:n]))
			if GetRequestPath(string(buf[:n])) == "/delay" {
				time.Sleep(time.Second * 5)
			}
			c.Write([]byte(response()))
		}(client)
	}

}
