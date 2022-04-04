package middlewares

// Handler 处理者接口 比如HttpHandler
type Handler interface {
	Execute()
	SetNext(handler Handler)
}
