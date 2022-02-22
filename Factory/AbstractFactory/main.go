package main

import "AbstractFactory/AbstractFactory"

func main() {
    var factory AbstractFactory.IFactory
    factory = AbstractFactory.HuaweiFactory{}
    factory.GetPhone().PhoneName()
    factory.GetRouter().RouterName()

    factory = AbstractFactory.MiFactory{}
    factory.GetPhone().PhoneName()
    factory.GetRouter().RouterName()
}
