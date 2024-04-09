package main

import "errors"

func NewAccount(accountType string) (Account, error) {
	switch accountType {
	case "checking":
		return &CheckingAccountBalance{balance: 1000}, nil
	case "savings":

		return &SavingAccount{balance: 2000}, nil
	default:
		return nil, errors.New("invalid account type")

	}
}
