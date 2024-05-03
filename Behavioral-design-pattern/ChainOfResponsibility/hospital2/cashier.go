package main

import "fmt"

var _ Departmant = &Cashier{}

type Cashier struct {
	next Departmant
}

func (c *Cashier) execute(patient *Patient) {
	if patient.paymentDone {
		fmt.Println("payment already done ")
		fmt.Println("finish")
	}
	fmt.Println("finish")
}

func (c *Cashier) setNext(next Departmant) {
	c.next = next
}
