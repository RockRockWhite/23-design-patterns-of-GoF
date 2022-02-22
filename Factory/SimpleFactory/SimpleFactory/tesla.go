package SimpleFactory

import "fmt"

type tesla struct {
}

func (tesla) Name() {
    fmt.Println("tesla")
}

func NewTesla() ICar {
    return &tesla{}
}
