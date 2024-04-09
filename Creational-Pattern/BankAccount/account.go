package main

type Account interface {
	GetBalance() float64
	Withdraw(ampunt float64) error
	Deposit(amount float64) error
}
