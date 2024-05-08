package main

import "fmt"

var _ State = &VendingMachine{}

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemPrice: itemPrice,
		itemCount: itemCount,
	}
	hasItemState := &HasItemState{vendingMachine: v}
	noHasItemState := &NoHasItemState{vendingMachine: v}
	hasMoneyState := &HasMoneyState{vendingMachine: v}
	itemRequestedState := &ItemRequestedState{vendingMachine: v}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.noItem = noHasItemState
	v.hasMoney = hasMoneyState

	return v
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) addItem(i int) error {
	return v.currentState.addItem(i)
}

func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) insertMoney(money int) error {
	err := v.currentState.insertMoney(money)
	if err != nil {
		return err
	}
	return nil
}

func (v *VendingMachine) dispenseItem() error {
	err := v.currentState.dispenseItem()
	if err != nil {
		return err
	}
	return nil
}
func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
