package main

import (
	proxy "Proxy/Proxy"
	"fmt"
)

func main() {
	// 用户调用代理, 代理执行实际服务, 并且完成一些附加的操作
	server := proxy.CreateNginxServer()

	for i := 0; i != 5; i++ {
		code, res := server.Request("/app/status", "GET")
		fmt.Println(code, res)
	}

	for i := 0; i != 5; i++ {
		code, res := server.Request("/create/user", "POST")
		fmt.Println(code, res)
	}
}
