package main

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (h *HasMoneyState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h *HasMoneyState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.itemRequested)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("ite requested\n")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("plesae select item first")
}

func (h *HasMoneyState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}
