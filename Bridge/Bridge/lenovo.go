package Bridge

import "fmt"

type Lenovo struct {
}

func (l Lenovo) Info() {
    fmt.Println("I'm Lenovo")
}
