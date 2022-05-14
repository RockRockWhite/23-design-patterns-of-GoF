package VendingMachine

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (n NoItemState) AddItem(i int) error {
	n.vendingMachine.IncrementItemCount(i)
	n.vendingMachine.SetState(n.vendingMachine.HasItem)
	return nil
}

func (n NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (n NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (n NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
