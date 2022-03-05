package Bridge

import "fmt"

type Computer struct {
    brand IBrand // 桥接两个类
}

func NewComputer(brand IBrand) *Computer {
    return &Computer{brand: brand}
}

func (computer *Computer) Info() {
    computer.brand.Info()
    fmt.Println("This is a computer")
}
