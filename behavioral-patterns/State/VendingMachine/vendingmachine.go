package VendingMachine

import "fmt"

type VendingMachine struct {
	HasItem       State
	ItemRequested State
	HasMoney      State
	NoItem        State

	currentState State
	itemCount    int
	itemPrice    int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	vm := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	vm.HasItem = &HasItemState{vm}
	vm.ItemRequested = &ItemRequestedState{vm}
	vm.HasMoney = &HasMoneyState{vm}
	vm.NoItem = &NoItemState{vm}
	vm.SetState(vm.HasItem)

	return vm
}

func (vm *VendingMachine) RequestItem() error {
	return vm.currentState.RequestItem()
}

func (vm *VendingMachine) AddItem(count int) error {
	return vm.currentState.AddItem(count)
}

func (vm *VendingMachine) InsertMoney(money int) error {
	return vm.currentState.InsertMoney(money)
}

func (vm *VendingMachine) DispenseItem() error {
	return vm.currentState.DispenseItem()
}

func (vm *VendingMachine) SetState(s State) {
	vm.currentState = s
}

func (vm *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	vm.itemCount += count
}
