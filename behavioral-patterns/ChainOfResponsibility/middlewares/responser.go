package middlewares

import "fmt"

type Responser struct {
	next Handler
}

func (r *Responser) Execute() {
	fmt.Println("进入responser中间件")
	fmt.Println("回复用户消息")
	if r.next != nil {
		r.next.Execute()
	}
	fmt.Println("退出responser中间件")
}

func (r *Responser) SetNext(handler Handler) {
	r.next = handler
}
