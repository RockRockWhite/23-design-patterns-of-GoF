package main

import (
    "Builder/Builder"
    "fmt"
)

func main() {
    director := Builder.Director{}
    house := director.Build(&Builder.Worker{})

    fmt.Println(house.Floor)
    fmt.Println(house.Door)
    fmt.Println(house.Window)
}
