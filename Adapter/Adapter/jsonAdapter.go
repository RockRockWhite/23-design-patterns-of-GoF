package Adapter

import "fmt"

// JsonAdapter 适配器实现客户端接口, 封装不可改动的类
type JsonAdapter struct {
    xmlServer *XmlServer
}

func (adapter JsonAdapter) Server() {
    fmt.Println("Json converts Xml.")
    adapter.xmlServer.Server()
}

func CreateAdapter(server *XmlServer) *JsonAdapter {
    return &JsonAdapter{xmlServer: server}
}
