package proxy

type Nginx struct {
	application       *Application // 代理包含被代理服务, 管理其生命周期
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func CreateNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       map[string]int{},
	}
}

func (nginx *Nginx) Request(url, method string) (int, string) {
	if nginx.checkRateLimiting(url) {
		return nginx.application.Requset(url, method)
	} else {
		return 403, "Forbidden."
	}
}

func (nginx *Nginx) checkRateLimiting(url string) bool {
	if nginx.rateLimiter[url] == 0 {
		nginx.rateLimiter[url] = 1
	}
	if nginx.rateLimiter[url] > nginx.maxAllowedRequest {
		return false
	}
	nginx.rateLimiter[url] = nginx.rateLimiter[url] + 1
	return true
}
