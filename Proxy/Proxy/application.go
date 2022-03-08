package proxy

// 真实的服务
type Application struct {
}

func (app *Application) Requset(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok."
	}
	if url == "/create/user" && method == "POST" {
		return 201, "Created."
	}
	return 404, "Not Found."
}
