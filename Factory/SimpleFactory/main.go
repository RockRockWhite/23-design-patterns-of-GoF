package main

import "factory/SimpleFactory"

func main() {
    // 简单工厂模式
    car := SimpleFactory.NewBenz()
    car2 := SimpleFactory.NewTesla()

    car.Name()
    car2.Name()
}
