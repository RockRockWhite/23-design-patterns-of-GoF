package main

import (
	"State/VendingMachine"
	"fmt"
)

func main() {
	vm := VendingMachine.NewVendingMachine(1, 10)

	err := vm.RequestItem()
	if err != nil {
		fmt.Println(err)
	}

	err = vm.InsertMoney(10)
	if err != nil {
		fmt.Println(err)
	}

	err = vm.DispenseItem()
	if err != nil {
		fmt.Println(err)
	}

	err = vm.RequestItem()
	if err != nil {
		fmt.Println(err)
	}
}
