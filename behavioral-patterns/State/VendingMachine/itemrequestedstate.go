package VendingMachine

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (s ItemRequestedState) AddItem(i int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (s ItemRequestedState) RequestItem() error {
	return fmt.Errorf("Item already requested")
}

func (s ItemRequestedState) InsertMoney(money int) error {
	if money < s.vendingMachine.itemPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", s.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	s.vendingMachine.SetState(s.vendingMachine.HasMoney)
	return nil
}

func (s ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
