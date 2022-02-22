package Adapter

import "fmt"

// 无法改变的类

type XmlServer struct {
}

// Server 使用xml的服务
func (XmlServer) Server() {
    fmt.Println("Xml Server")
}
