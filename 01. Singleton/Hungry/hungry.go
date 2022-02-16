package Hungry

import "fmt"

// 饿汉模式
type hungry struct {
}

var instance *hungry

// 导入包执行
func init() {
    instance = &hungry{}
    fmt.Println("init once")
}

func GetInstance() *hungry {
    return instance
}
