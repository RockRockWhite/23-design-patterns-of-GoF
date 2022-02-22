package AbstractFactory

import "fmt"

type HuaweiPhone struct {
}

func (HuaweiPhone) PhoneName() {
    fmt.Println("HuaweiPhone")
}
