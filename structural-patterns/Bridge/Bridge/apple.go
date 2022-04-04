package Bridge

import "fmt"

type Apple struct {
}

func (a Apple) Info() {
    fmt.Println("I'm Apple.")
}
