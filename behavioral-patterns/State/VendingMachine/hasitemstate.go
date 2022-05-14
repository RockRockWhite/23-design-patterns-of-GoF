package VendingMachine

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (s HasItemState) AddItem(i int) error {
	s.vendingMachine.IncrementItemCount(i)
	return nil
}

func (s HasItemState) RequestItem() error {
	if s.vendingMachine.itemCount == 0 {
		s.vendingMachine.SetState(s.vendingMachine.NoItem)
		return fmt.Errorf("No item present")
	}
	fmt.Println("Item requested")
	s.vendingMachine.SetState(s.vendingMachine.ItemRequested)
	return nil
}

func (s HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (s HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
