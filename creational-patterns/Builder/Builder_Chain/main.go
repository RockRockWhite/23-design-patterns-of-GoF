package main

import (
    "Builder/Builder"
    "fmt"
)

func main() {
    builder := Builder.Worker{}

    builder.BuildWindow("name1").
        BuildFloor("name2").
        BuildDoor("name3").
        BuildWindow("后来的修改")

    house := builder.GetHouse()
    fmt.Println(house.Door)
    fmt.Println(house.Floor)
    fmt.Println(house.Window)
}
