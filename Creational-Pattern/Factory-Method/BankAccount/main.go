package main

import "fmt"

func main() {
	checkingAccount, err := NewAccount("checking")
	if err != nil {
		fmt.Println("Error creating checking account:", err)
		return
	}

	savingsAccount, err := NewAccount("savings")
	if err != nil {
		fmt.Println("Error creating savings account:", err)
		return
	}

	fmt.Println("Checking Account Balance:", checkingAccount.GetBalance())
	fmt.Println("Savings Account Balance:", savingsAccount.GetBalance())
}
