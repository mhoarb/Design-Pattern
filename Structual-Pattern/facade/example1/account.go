package main

import "fmt"

type Account struct {
	name string
}

func NewAccount(accountName string) *Account {
	return &Account{name: accountName}
}

func (a *Account) CheckAccount(name string) error {
	if a.name != name {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("account verified")
	return nil
}
