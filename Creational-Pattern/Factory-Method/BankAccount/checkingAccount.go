package main

import "errors"

type CheckingAccountBalance struct {
	balance float64
}

func (c *CheckingAccountBalance) GetBalance() float64 {
	return c.balance
}
func (c *CheckingAccountBalance) Withdraw(amount float64) error {
	if amount >= c.balance {
		return errors.New("Insufficient funds")
	}
	c.balance -= amount
	return nil
}

func (c *CheckingAccountBalance) Deposit(amount float64) error {
	c.balance = +amount
	return nil
}

type SavingAccount struct {
	balance float64
}

func (s *SavingAccount) GetBalance() float64 {
	return s.balance
}

func (s *SavingAccount) Withdraw(amount float64) error {
	if s.balance >= amount {
		return errors.New("Insufficient funds")
	}
	s.balance = -amount
	return nil
}

func (s *SavingAccount) Deposit(amount float64) error {
	s.balance += amount
	return nil
}
