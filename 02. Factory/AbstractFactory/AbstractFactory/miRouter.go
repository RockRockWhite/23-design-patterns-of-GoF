package AbstractFactory

import "fmt"

type MiRouter struct {
}

func (MiRouter) RouterName() {
    fmt.Println("MiRouter")
}
