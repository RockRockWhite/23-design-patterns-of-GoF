package AbstractFactory

import "fmt"

type HuaweiRouter struct {
}

func (HuaweiRouter) RouterName() {
    fmt.Println("HuaweiRouter")
}
