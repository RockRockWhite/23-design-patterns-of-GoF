package SimpleFactory

import "fmt"

type benz struct {
}

func (benz) Name() {
    fmt.Println("benz")
}

func NewBenz() ICar {
    return &benz{}
}
