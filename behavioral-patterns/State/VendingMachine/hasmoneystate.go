package VendingMachine

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (s HasMoneyState) AddItem(i int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (s HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (s HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (s HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing item")
	s.vendingMachine.itemCount = s.vendingMachine.itemCount - 1
	if s.vendingMachine.itemCount == 0 {
		s.vendingMachine.SetState(s.vendingMachine.NoItem)
	} else {
		s.vendingMachine.SetState(s.vendingMachine.HasItem)
	}
	return nil
}
