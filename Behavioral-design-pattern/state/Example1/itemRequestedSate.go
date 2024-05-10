package main

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("item deispense in progress")
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("money intered is ok")
	return nil
}

func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}
