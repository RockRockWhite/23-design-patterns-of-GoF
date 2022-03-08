package main

import (
    "Decorator/Decorator"
    "fmt"
)

func main() {
    var pizza Decorator.Pizza = &Decorator.VeggieMania{}

    // 添加tomato
    pizza = &Decorator.TomatoTopping{Pizza: pizza}
    // 添加cheese
    pizza = &Decorator.CheeseTopping{Pizza: pizza}

    fmt.Println(pizza.GetPrice())
}
