package main

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (h HasItemState) addItem(i int) error {
	fmt.Printf("&d items added \n", i)
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *HasItemState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("no titem present")
	}
	fmt.Printf("item requested\n")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (h *HasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}
