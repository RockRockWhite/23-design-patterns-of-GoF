package middlewares

import "fmt"

type Auth struct {
	next Handler
}

func (a *Auth) Execute() {
	fmt.Println("进入auth中间件")
	if a.next != nil {
		a.next.Execute()
	}
	fmt.Println("退出auth中间件")
}

func (a *Auth) SetNext(handler Handler) {
	a.next = handler
}
