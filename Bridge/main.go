package main

import "Bridge/Bridge"

func main() {
    lenovo := Bridge.NewComputer(Bridge.Lenovo{})
    lenovo.Info()

    apple := Bridge.NewComputer(Bridge.Apple{})
    apple.Info()
}
