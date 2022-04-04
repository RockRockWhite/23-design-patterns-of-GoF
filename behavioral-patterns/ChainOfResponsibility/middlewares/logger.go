package middlewares

import "fmt"

type Logger struct {
	next Handler
}

func (l *Logger) Execute() {
	fmt.Println("进入logger中间件")
	if l.next != nil {
		l.next.Execute()
	}
	fmt.Println("退出logger中间件")
}

func (l *Logger) SetNext(handler Handler) {
	l.next = handler
}
