package main

import "adapter/Adapter"

func main() {
    adapter := Adapter.CreateAdapter(&Adapter.XmlServer{})
    adapter.Server()
}
