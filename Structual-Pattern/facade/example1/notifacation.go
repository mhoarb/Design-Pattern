package main

import "fmt"

type Notification struct {
}

func (n *Notification) WalletCreditNotif() {
	fmt.Println("sending wallet credit notification")
}
func (n *Notification) WalletDebitNotif() {
	fmt.Println("sending wallet debit notification")
}
