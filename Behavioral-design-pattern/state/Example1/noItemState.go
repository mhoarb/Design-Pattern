package main

import "fmt"

type NoHasItemState struct {
	vendingMachine *VendingMachine
}

func (n *NoHasItemState) addItem(i int) error {
	n.vendingMachine.incrementItemCount(i)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}

func (n *NoHasItemState) requestItem() error {
	return fmt.Errorf("item out of Stock")
}

func (n *NoHasItemState) insertMoney(money int) error {
	return fmt.Errorf("ItemOust of stock")
}

func (n *NoHasItemState) dispenseItem() error {
	return fmt.Errorf("ItemOust of stock")
}
