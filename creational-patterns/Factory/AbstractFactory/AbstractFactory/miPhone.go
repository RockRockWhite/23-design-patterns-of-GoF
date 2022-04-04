package AbstractFactory

import "fmt"

type MiPhone struct {
}

func (MiPhone) PhoneName() {
    fmt.Println("MiPhone")
}
