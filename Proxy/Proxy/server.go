package proxy

// 真实服务和代理都实现一样的接口
type Server interface {
	Request(string, string) (int, string)
}
